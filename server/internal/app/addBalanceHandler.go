package app

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) AddBalanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("Not correct method, instead of Post")
		http.Error(w, "Only Post method", http.StatusMethodNotAllowed)
	}
	var res balanceStruct
	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		log.Println(err, "could not decode json")
		http.Error(w, "could not decode json", http.StatusBadRequest)
		return
	}
	if value, ok, err := s.Storage.GetBalance(res.UserID); !ok {
		//err != nil not only because there is no such user. need to update to 1 more param
		if err != nil && err != sql.ErrNoRows {
			log.Println(err, "db balance search error")
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		err2 := s.Storage.InsertBalance(res.UserID, res.Balance)
		if err2 != nil {
			log.Println(err2, "could not insert new user")
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	} else {
		err2 := s.Storage.UpdateBalance(res.UserID, res.Balance+value)
		if err2 != nil {
			log.Println(err2, "could not update user")
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
