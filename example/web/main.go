package main

import (
	"encoding/json"
	"liangzi.local/sjqm"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", qm)
	http.ListenAndServe(":8080", nil)
}

func qm(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Local()
	y := t.Year()
	dgz := "壬申"
	hgz := "丙午"
	st := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, time.Local)
	g := sjqm.Result(y, dgz, hgz, st)
	json.NewEncoder(w).Encode(g)

}

//go run main.go
