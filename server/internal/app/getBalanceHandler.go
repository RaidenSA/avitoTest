package app

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

func (s *Server) GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("Not correct method, instead of Get")
		http.Error(w, "Only Get method", http.StatusMethodNotAllowed)
		return
	}

	q := r.URL.String()
	q = q[(len(GetBalanceEndPoint)):]
	log.Println(q)
	if q == "" {
		log.Println("Empty user")
		http.Error(w, "Empty user", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		log.Println("Not correct user ID")
		http.Error(w, "Not correct user ID", http.StatusBadRequest)
		return
	}

	if value, ok, _ := s.Storage.GetBalance(userID); !ok {
		log.Println("No such user")
		http.Error(w, "No such user", http.StatusNotFound)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusFound)
		response := balanceStruct{
			UserID:  userID,
			Balance: value,
		}
		err = json.NewEncoder(w).Encode(&response)
		if err != nil {
			log.Println(err, "encoding")
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		log.Println("Get balance success")
	}
	return
}
