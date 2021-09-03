package main

import (
	"APIxml/postgres"
	"bytes"
	"fmt"

	"bufio"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/feed_example/", feedExample)
	http.ListenAndServe(":8081", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {

}

func feedExample(w http.ResponseWriter, r *http.Request) {
	link, err := postgres.GetLink(1)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(link)

	fileData, errData := os.Open("xmlFiles//feed_example.xml")

	if errData != nil {
		log.Fatal(errData)
	}

	wr := bytes.Buffer{}
	sc := bufio.NewScanner(fileData)

	for sc.Scan() {
		wr.WriteString(sc.Text())
	}

	fmt.Fprintln(w, wr.String())
}
