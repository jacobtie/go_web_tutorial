package main

import (
	"github.com/boltdb/bolt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"github.com/asdine/storm"
)

// Struct used by database
type Program struct {
	ID int 'storm:"id, increment"'
	User string 'storm:"unique"'
	Code string
	Date time.Time
}

// Storm/BoltDB initialization
func initDb() {
	db, err := storm.Open("savedData.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

// Save source code to file
func saveHandler(w http.ResponseWriter, r *http.Request){

}
