package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const otherWord = "*"

func main() {
	transforms := fetchDBTransforms()
	generate(transforms, os.Stdin, os.Stdout)
}

func fetchDBTransforms() []string {
	userName := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	addr := os.Getenv("MYSQL_ADDR")
	connString := fmt.Sprintf("%s:%s@tcp(%s)/sprinckle", userName, pass, addr)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Couldnt open sql connection", err)
	}
	rows, err := db.Query("select value,prefix,suffix from complements;")
	if err != nil {
		log.Fatal("couldnt perform query", err)
	}
	var transforms []string
	var value string
	var prefix bool
	var suffix bool
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&value, &prefix, &suffix)
		if err != nil {
			log.Fatal(err)
		}
		if suffix {
			transforms = append(transforms, fmt.Sprintf("%s %s", otherWord, value))
		}
		if prefix {
			transforms = append(transforms, fmt.Sprintf("%s %s", value, otherWord))
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return transforms
}

func generate(transforms []string, r io.Reader, w io.Writer) {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(r)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Fprintln(w, strings.Replace(t, otherWord, s.Text(), -1))
	}
}
