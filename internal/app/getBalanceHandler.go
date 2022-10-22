package app

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type balanceStruct struct {
	UserID  string
	Balance int
}

type transactionStruct struct {
	UserID    string
	ServiceID string
	OrderID   string
	Sum       int
}

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
	if value, ok, err := s.Storage.GetBalance(q); !ok {
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
			UserID:  q,
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
		if err != nil {
			log.Println(err, "db balance search")
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

func (s *Server) ReserveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("Not correct method, instead of Post")
		http.Error(w, "Only Post method", http.StatusMethodNotAllowed)
	}
	var res transactionStruct
	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		log.Println(err, "could not decode json")
		http.Error(w, "could not decode json", http.StatusBadRequest)
		return
	}
	log.Println(res)
	if value, ok, _ := s.Storage.GetBalance(res.UserID); !ok {
		log.Println("No such user")
		http.Error(w, "No such user", http.StatusNotFound)
		return
	} else {
		if (value - res.Sum) < 0 {
			//return not enough
			log.Println("Not enough funds")
			http.Error(w, "Not enough funds", http.StatusPaymentRequired)
			return
		}
		//proceed to update reserved table
		err2 := s.Storage.UpdateBalance(res.UserID, value-res.Sum)
		if err2 != nil {
			log.Println(err2, "could not update user")
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

	}
}
func (s *Server) AcquireHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("Not correct method, instead of Post")
		http.Error(w, "Only Post method", http.StatusMethodNotAllowed)
	}
	var res transactionStruct
	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		log.Println(err, "could not decode json")
		http.Error(w, "could not decode json", http.StatusBadRequest)
		return
	}
}
