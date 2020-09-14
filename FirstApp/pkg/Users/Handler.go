package Users

import (
	"encoding/json"
	"net/http"
	"strconv"
)

var (
	user User
)

func getHandler(s Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			json.NewEncoder(w).Encode("Something went wrong... Either user does not exists, or id is wrong")
			return
		}
		u, err := s.ReadUser(id)
		if err != nil {
			json.NewEncoder(w).Encode("Something went wrong... Either user does not exists, or id is wrong")
			return
		}
		json.NewEncoder(w).Encode(&u)
	}
}

func postHandler(s Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&user)
		if err := s.SaveUser(user); err != nil {
			json.NewEncoder(w).Encode("Something went wrong...")
		}

		println("just added user to db...")
		// json.NewEncoder(w).Encode("You just added user1 to the cassandra db")
		json.NewEncoder(w).Encode("You just added user1 to the mysql db")

	}
}
