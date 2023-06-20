package main

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/exp/slices"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{}
var counter = 0

func usersHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/users" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		stringify, _ := json.Marshal(users)
		w.Header().Set("Content-Type", "application/json")
		w.Write(stringify)
		return
	}

	if r.Method == "POST" {

		var u User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		u.Id = counter
		counter++
		users = append(users, u)
		stringify, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(stringify)
		return
	}

	if r.Method == "PUT" {
		var u User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for i := 0; i < len(users); i++ {
			if users[i].Id == u.Id {
				users[i] = u
				break
			}
		}

		stringify, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "application/json")
		w.Write(stringify)
		return
	}

	if r.Method == "DELETE" {
		var u User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		stringify, _ := json.Marshal(u)

		for i := 0; i < len(users); i++ {
			if users[i].Id == u.Id {
				users = slices.Delete(users, i, i+1)
				break
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(stringify)
		return
	}
}

func main() {
	http.HandleFunc("/users", usersHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
