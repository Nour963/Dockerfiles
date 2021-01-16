package main

import (
   "fmt"
   /* "log"
    "net/http"*/
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db,err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
    if db.Ping() != nil {
       fmt.Println("if block")
       fmt.Println(err)
    } else {
         fmt.Println("else block")
           }
}
