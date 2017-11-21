package api

import (
	"net/http"
	"time"

	"github.com/Koke/BC/bdservice/mongo"
	"github.com/Koke/BC/bdservice/utils/log"

	"github.com/gorilla/mux"
)

type Controller struct {
	logger       log.ILogger
	mongoService mongo.IMongoService
}

func (c *Controller) StartServer() {
	c.inject()

	r := mux.NewRouter()
	r.HandleFunc("/Users", c.handleUsers).Methods("GET", "POST")
	r.HandleFunc("/User/{id}", c.handleUser).Methods("GET")
	r.HandleFunc("/User/{id}/Wallet", c.handleWallet).Methods("GET", "POST")

	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8100",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	c.logger.WriteLog("The server Start in port "+srv.Addr, log.Info)
	srv.ListenAndServe()
}

//Injections
func (c *Controller) inject() {
	var injLogger log.Logger
	c.logger = &injLogger

	var injMongoService mongo.MongoService
	c.mongoService = &injMongoService
}
