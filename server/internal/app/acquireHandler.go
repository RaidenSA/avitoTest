package app

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) AcquireHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("Not correct method, instead of Post")
		http.Error(w, "Only Post method", http.StatusMethodNotAllowed)
		return
	}

	var res transactionStruct
	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		log.Println(err, "could not decode json")
		http.Error(w, "could not decode json", http.StatusBadRequest)
		return
	}

	err = s.Storage.Acquire(res.UserID, res.ServiceID, res.OrderID, res.Sum)
	if err != nil {
		log.Println(err, "could not reserve")
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	return
}
