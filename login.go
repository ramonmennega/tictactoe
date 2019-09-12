package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)
var username string
var password string
var users []string
var win string
var execute string
func logWin(a string) {
	win = a

	fmt.Printf(a)
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/tictactoego")
	if err != nil {
	} else {
		fmt.Println("db is connected")
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
	}
	if a == "O"{
		fmt.Println(users[1])
		execute = fmt.Sprintf("INSERT INTO wins (username) VALUES ('%s')", users[1])

	} else if a == "X" {
		fmt.Println(users[0])
		execute = fmt.Sprintf("INSERT INTO wins (username) VALUES ('%s')", users[0])

	}
	fmt.Printf(execute)

	_, _ = db.Query(execute)


}


func login(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	username = r.PostFormValue("username")
	password = r.PostFormValue("password")
	CreateCon(w,r)

}

func CreateCon(w http.ResponseWriter, r *http.Request) *sql.DB {

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/tictactoego")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("MySQL db is not connected")
		fmt.Println(err.Error())
	}
	execute := fmt.Sprintf("SELECT username,password FROM users WHERE username = '%s'", username)

	rows, err := db.Query(execute)

	for rows.Next() {
		var password2 string
		err = rows.Scan(&username, &password2)
		fmt.Println(password2)
		if password == password2 {
			users = append(users, username)
			for k,_ := range users{
				if k == 1 {
					KKK = Tictactoe{
						Beurt: "1",
						Win:   "",
						P1: 0,
						P2: 0,
						Username1: users[0],
						Username2: users[1],
						Bord: Speelveld{
							[9]string{"", "", "", "", "", "", "", "", ""},
						},
					}
				}
			}



			fmt.Println(users)

			http.Redirect(w, r, "/game.html", http.StatusSeeOther)

		}


	}
	http.Redirect(w, r, "/index.html", http.StatusSeeOther)


	return db
}