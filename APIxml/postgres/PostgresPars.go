package postgres

import (
	"database/sql"
	"log"
)

var conStat bool = false
var DB *sql.DB

func makeConnect() {
	connStr := "host=host.docker.internal port=49156 user=root password=12345 dbname=link sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	DB = db
	conStat = true
}

func CreateNewLink(value string) {

	if !conStat {
		makeConnect()
	}

	DB.Exec("INSERT INTO links(value) VALUES('xmlFiles//$1')", value)
}

func GetLink(ID int) (string, error) {
	if !conStat {
		makeConnect()
	}

	rows, err := DB.Query("SELECT value FROM links WHERE id=$1", ID)

	if err != nil {
		return "", err
	}

	if rows.Next() {
		var copyLink string

		errScan := rows.Scan(&copyLink)

		if errScan != nil {
			return "", err
		}

		return copyLink, nil
	} else {
		return "", nil
	}
}
