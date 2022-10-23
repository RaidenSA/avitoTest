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
		http.Error(w, "Empty user", http.StatusNotFound)
		return
	}
	userID, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		log.Println("Not correct user ID")
		http.Error(w, "Not correct user ID", http.StatusNotFound)
		return
	}

	if value, ok, err := s.Storage.GetBalance(userID); !ok {
		log.Println("No such user")
		http.Error(w, "No such user", http.StatusNotFound)
		return
	} else {
		//we get balance in if statement
		//need to return marshalled user balance
		w.Header().Set("Content-Type", "application/json")
		//w.Header().Set("Location", string(value))
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
		//log.Println("HTTP GET request served. Got token:", q, " Sent URL:", value)
	}
	return
}
