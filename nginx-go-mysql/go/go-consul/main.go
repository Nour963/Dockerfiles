package main

import (
    "fmt"
    "log"
    "net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, _ := sql.Open("mysql", "nour:hello123@tcp(localhost:3306)/test")
    if db.Ping() != nil {
   fmt.Println("Couldn't connect to the database!")
   fs := http.FileServer(http.Dir("no"))
   http.Handle("/", fs)
   log.Fatal(http.ListenAndServe(":8080", nil))
    } else {
     fmt.Println("connect to the database!")
     fs := http.FileServer(http.Dir("yes"))
     http.Handle("/", fs)
     log.Fatal(http.ListenAndServe(":8080", nil))
           }
}
