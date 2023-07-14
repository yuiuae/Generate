package db

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/yuiuae/Generate/internal/conf"
	"github.com/yuiuae/Generate/internal/structs"
	"github.com/yuiuae/Generate/pkg/hasher"

	_ "github.com/go-sql-driver/mysql"
)

func openDb(driverName string, dataSourceName string) *sql.DB {
	// db, err := sql.Open(driverName, dataSourceName)
	db, err := sql.Open("mysql", "root:yui1hsm2@/chat")

	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	return db
}

func GetAllUsers() ([]structs.UserInfo, error) {
	var users []structs.UserInfo
	db := openDb(conf.DriverName, conf.DataSourceName)
	defer db.Close()
	rows, err := db.Query("SELECT username,userpasshash,useruuid FROM users")
	if err != nil {
		return nil, fmt.Errorf("error - SELECT * FROM users, err = %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	// var userName, userPassHash, userUUID string
	for rows.Next() {
		var usr structs.UserInfo
		if err := rows.Scan(&usr.Username, &usr.UserPassHash, &usr.UserUUID); err != nil {
			return nil, fmt.Errorf("error - fetch FROM users, err = %v", err)
		}
		users = append(users, usr)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error2 - SELECT * FROM users, err = %v", err)
	}
	return users, nil
}

func getUserInfo(db *sql.DB, uname string) (structs.UserInfo, error) {
	ui := structs.UserInfo{}
	err := db.QueryRow("SELECT username,userpasshash, useruuid FROM users WHERE username = ?", uname).Scan(&ui.Username, &ui.UserPassHash, &ui.UserUUID)
	return ui, err
}

func AddNewUser(uname, uhashpass, uuuid string) (int, error) {
	db := openDb(conf.DriverName, conf.DataSourceName)
	defer db.Close()

	// checking that a user with this username exists in the db
	_, err := getUserInfo(db, uname)
	if err != nil {

		if err.Error() != "sql: no rows in result set" {
			return http.StatusInternalServerError, errors.New("internal Server Error (user existence check)")
		} else {
			query := "INSERT INTO users (username, userpasshash, useruuid, usertime) VALUES (?, ?, ?, NOW())"
			insert, err := db.Prepare(query)
			if err != nil {
				return http.StatusInternalServerError, errors.New("internal Server Error (insert into users)")
			}
			_, err = insert.Exec(uname, uhashpass, uuuid)
			insert.Close()
			if err != nil {
				return http.StatusInternalServerError, errors.New("internal Server Error (insert users 2)")
			}
			fmt.Println("user added")
			return 0, nil
		}
	}
	return http.StatusBadRequest, errors.New("user alredy exists")

}

func PassVerify(uname, upass string) (structs.UserInfo, bool, error) {
	db := openDb(conf.DriverName, conf.DataSourceName)
	defer db.Close()
	ui, err := getUserInfo(db, uname)
	if err != nil {
		return ui, false, err
	}
	ok := hasher.CheckPasswordHash(ui.UserPassHash, upass)
	return ui, ok, err
}

func AddMessage(uname, msgtext string) (int, error) {
	db := openDb(conf.DriverName, conf.DataSourceName)
	defer db.Close()

	query := "INSERT INTO messages (username, msgtext, msgtime) VALUES (?, ?, NOW())"
	insert, err := db.Prepare(query)
	if err != nil {
		fmt.Println("error = ", err)
		return http.StatusInternalServerError, errors.New("internal Server Error (insert into messages)")
	}
	_, err = insert.Exec(uname, msgtext)
	insert.Close()
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal Server Error (insert into messages 2)")
	}
	fmt.Println("messages added")
	return http.StatusOK, nil

}

func CloseTimeUpdate(uname string) (int, error) {
	db := openDb(conf.DriverName, conf.DataSourceName)
	defer db.Close()
	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()

	query := "UPDATE users SET usertime = NOW() where username = ?"
	insert, err := db.Prepare(query)
	if err != nil {
		fmt.Println("error = ", err)
		return http.StatusInternalServerError, errors.New("internal Server Error (update time)")
	}
	_, err = insert.Exec(uname)
	insert.Close()
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal Server Error (update time 2)")
	}
	fmt.Println("time updated")
	return http.StatusOK, nil

}
