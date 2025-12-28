package main

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello Root Route"))
	fmt.Println("Hello Root Route")
}

func teacherHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// fmt.Println(r.URL.Path)
		// path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		// userID := strings.TrimSuffix(path, "/")

		// fmt.Println(userID, "user id")
		fmt.Println("Query params:", r.URL.Query())

		w.Write([]byte("hello teachers Route GET"))
	case http.MethodPost:
		w.Write([]byte("hello teachers Route POST"))
	case http.MethodPatch:
		w.Write([]byte("hello teachers Route PATCH"))
	case http.MethodDelete:
		w.Write([]byte("hello teachers Route DELETE"))
	}
}

func studetsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("hello students Route GET"))
	case http.MethodPost:
		w.Write([]byte("hello students Route POST"))
	case http.MethodPatch:
		w.Write([]byte("hello students Route PATCH"))
	case http.MethodDelete:
		w.Write([]byte("hello students Route DELETE"))
	}
}

func execHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("hello students Route GET"))
	case http.MethodPost:
		w.Write([]byte("hello students Route POST"))
	case http.MethodPatch:
		w.Write([]byte("hello students Route PATCH"))
	case http.MethodDelete:
		w.Write([]byte("hello students Route DELETE"))
	}
}

func main() {
	port := ":3000"

	// http.HandleFunc("/", rootHandler)

	http.HandleFunc("/teachers/", teacherHandler)
	http.HandleFunc("/students", studetsHandler)
	http.HandleFunc("/execs", execHandler)

	fmt.Println("Server is running on port", port)

	// server listens
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalln("error starting the server", err)
	}
}
