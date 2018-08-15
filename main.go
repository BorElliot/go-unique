// Reference: https://www.calhoun.io/creating-random-strings-in-go/
package main

import (
	"database/sql"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const charset = "0123456789"

var db *sql.DB

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/my_db")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1000000000; i++ {
		random_str := String(1)
		_, err = db.Exec("INSERT INTO codes(code) VALUES(?)", random_str)
		if err != nil {
			continue
		}
		log.Println(random_str)
	}
}
