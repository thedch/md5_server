package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/thedch/md5_server/sums"
)

// Reading files requires checking most calls for errors.
// This helper will streamline the error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	host   = "postgres"
	port   = 5432
	user   = "postgres"
	dbname = "postgres"
)

func handler(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Path[1:]
	// fmt.Printf("%s", r.Header)
	hash := sums.GetMD5Hash(r.URL.Path[1:])
	// Print it for the user
	fmt.Fprintf(w, "Hello there! The md5 hash of %s is %x\n", input, hash)

	// Open the database, check for errors
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	check(err)
	defer db.Close()

	// Ping the database to ensure an actual connection
	err = db.Ping()
	check(err)

	// Add the user input to the database
	if input != "favicon.ico" {
		sqlInsert := `
    INSERT INTO md5 (input, hash)
    VALUES ($1, 'test')`

		_, err = db.Exec(sqlInsert, input) // Underscore?
		check(err)
	}

	// Query the database for infomation
	sqlStatement := "SELECT input FROM md5;"
	rows, err := db.Query(sqlStatement)
	check(err)
	defer rows.Close()
	for rows.Next() {
		var hash string
		err = rows.Scan(&hash)
		check(err)
		fmt.Fprintln(w, hash)
	}
	err = rows.Err()
	check(err)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
