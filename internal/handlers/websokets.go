package handlers

import (
	"Generate/internal/db"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// websocket connection request
func RequestWithToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorlog(w, "Only GET method allowed", http.StatusBadRequest)
		return
	}

	reqToken := r.URL.Query().Get("token")
	username, err := parseToken(reqToken, []byte(tokenSecretKey))
	fmt.Println("Token", username)
	if err != nil {
		errorlog(w, "Internal Server Error (parse)", http.StatusBadRequest)
		return
	}

	_, exists := connTable[username]
	if !exists {
		errorlog(w, "User error", http.StatusBadRequest)
		return
	}

	if reqToken != connTable[username].Token {
		errorlog(w, "Invalid token", http.StatusBadRequest)
		return
	}
	mu := sync.Mutex{}
	mu.Lock()
	connTable[username].Token = ""
	mu.Unlock()
	// tokenTable[username].ExpireAt = 0
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		errorlog(w, "Internal Server Error (upgrade)", http.StatusInternalServerError)
		return
	}

	connTable[username].WS = conn
	defer delete(connTable, username)
	defer conn.Close()
	defer db.CloseTimeUpdate(username)
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			errorlog(w, "Internal Server Error (ws read)", http.StatusInternalServerError)
			break
		}
		log.Printf("recv: %s", message)
		status, err := db.AddMessage(username, string(message))
		if err != nil {
			errorlog(w, err.Error(), status)
			return
		}

		for key, val := range connTable {
			if key != username {
				tt := append([]byte(username+": "), message...)
				err = val.WS.WriteMessage(mt, tt)
				if err != nil {
					errorlog(w, "Internal Server Error (ws write all)", http.StatusInternalServerError)
					break
				}
			}
		}
	}
}

// parse token from a websocket user connection request
func parseToken(accessToken string, signingKey []byte) (string, error) {
	var username string
	token, _, err := new(jwt.Parser).ParseUnverified(accessToken, jwt.MapClaims{})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		username = fmt.Sprint(claims["username"])
	}

	if username == "" {
		return "", fmt.Errorf("invalid token payload")
	}
	return username, nil

}
