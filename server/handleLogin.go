package api

import (
	"encoding/json"
	"net/http"

	"../mongo"
)

func (c *Controller) handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "POST":
		var user mongo.User
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		s := c.userServive.MongoSession()
		response := c.userServive.Login(s, user)
		result, _ := json.Marshal(response)
		w.Write(result)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
