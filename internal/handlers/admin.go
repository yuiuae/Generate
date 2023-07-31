// Copyright 2023 Serhii Khrystenko. All rights reserved.

/*
Package hasher implements user password verification.

This package is designed as an example of the Godoc
documentation and does not have any functionality:)
*/

package handlers

import (
	"Generate/internal/db"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorlog(w, "Only GET method allowed", http.StatusBadRequest)
		return
	}

	var msg string = `
	<html>
	<body>
	<h1>Welcome on main page!</h1>
	</body>
	</html>	
	`
	w.Write([]byte(msg))
}

// GetUserAll lists all existing users
//
//	@Summary      User List
//	@Description  get users
//	@Tags         users
//	@Produce      json
//	@Success      200
//	@Router       /users [get]
func GetUserAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TEST")
	if r.Method != "GET" {
		errorlog(w, "Only GET method allowed", http.StatusBadRequest)
		return
	}
	users, err := db.GetAllUsers()
	if err != nil {
		errorlog(w, "Internal Server Error (GetAllUsers)", http.StatusInternalServerError)
		return

	}
	for key, val := range users {
		fmt.Fprintln(w, key, "- ", val.Username, val.UserPassHash, val.UserUUID)
	}

}
