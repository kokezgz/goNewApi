package api

import (
	"encoding/json"
	"net/http"

	"../mongo"
	"github.com/gorilla/mux"
)

func (c *Controller) handleWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		s := c.mongoService.MongoSession()
		response, _ := c.mongoService.FindWallet(s, vars["id"])
		result, _ := json.Marshal(response)
		w.Write(result)

	case "POST":
		var wallet mongo.Wallet
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&wallet)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
		} else {
			vars := mux.Vars(r)
			s := c.mongoService.MongoSession()
			response, _ := c.mongoService.DepositWallet(s, vars["id"], wallet.Amount)
			result, _ := json.Marshal(response)
			w.Write(result)
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
