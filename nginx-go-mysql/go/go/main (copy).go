package main

import (
    "fmt"
    "log"
    "net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/test")
    err = db.Ping()
   if err != nil {
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

    defer db.Close()

}
