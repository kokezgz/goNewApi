package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Koke/BC/bdservice/mongo"
)

func (c *Controller) handleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		s := c.mongoService.MongoSession()
		response, _ := c.mongoService.AllUsers(s)
		users, _ := json.Marshal(response)
		w.Write(users)

	case "POST":
		var user mongo.User
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			s := c.mongoService.MongoSession()
			response := c.mongoService.NewUser(s, user)
			result, _ := json.Marshal(response)
			w.Write(result)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (c *Controller) handleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		s := c.mongoService.MongoSession()
		response, _ := c.mongoService.FindUser(s, vars["id"])
		result, _ := json.Marshal(response)
		w.Write(result)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
