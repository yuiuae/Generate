package main

import (
	"Generate/internal/handlers"
	"Generate/internal/middleware"
	"log"
	"net/http"
	"os"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// fmt.Println("YUI TEST")
	// server := chat.NewServer("/entry")
	// go server.Listen()

	http.Handle("/", middleware.MiddleLog(http.HandlerFunc(handlers.Index)))
	http.Handle("/user", middleware.MiddleLog(http.HandlerFunc(handlers.UserCreate)))
	http.Handle("/user/login", middleware.MiddleLog(http.HandlerFunc(handlers.UserLogin)))
	http.Handle("/admin", middleware.MiddleLog(http.HandlerFunc(handlers.GetUserAll)))
	http.Handle("/actusers", middleware.MiddleLog(http.HandlerFunc(handlers.ActiveUsers)))

	http.Handle("/chat", middleware.MiddleLog(http.HandlerFunc(handlers.RequestWithToken)))

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
