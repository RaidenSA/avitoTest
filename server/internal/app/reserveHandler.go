package app

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) ReserveHandler(w http.ResponseWriter, r *http.Request) {
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

	if value, ok, _ := s.Storage.GetBalance(res.UserID); !ok {
		log.Println("No such user")
		http.Error(w, "No such user", http.StatusNotFound)
		return
	} else {
		if (value - res.Sum) < 0 {
			log.Println("Not enough funds")
			http.Error(w, "Not enough funds", http.StatusPaymentRequired)
			return
		}
		err2 := s.Storage.Reserve(res.UserID, res.ServiceID, res.OrderID, res.Sum)
		if err2 != nil {
			log.Println(err2, "could not reserve")
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
	}
}
