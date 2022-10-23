package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (s *Server) DetailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("Not correct method, instead of Get")
		http.Error(w, "Only Get method", http.StatusMethodNotAllowed)
		return
	}

	inString := r.URL.String()
	inString = inString[(len(DetailEndPoint)):]
	str := strings.Split(inString, "-")
	if len(str) != 3 {
		log.Println("Not correct number of params")
		http.Error(w, "Not correct number of params", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseInt(str[0], 10, 64)
	if err != nil {
		log.Println("Not correct user ID")
		http.Error(w, "Not correct user ID", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(str[1])
	if err != nil {
		log.Println("Not correct 2-nd param")
		http.Error(w, "Not correct 2-nd param", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(str[2])
	if err != nil {
		log.Println("Not correct 3-rd param")
		http.Error(w, "Not correct 3-rd param", http.StatusBadRequest)
		return
	}

	detailSlice, err := s.Storage.Detailing(userID, limit, page)
	if err != nil {
		log.Println(err, "could not get report")
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(detailSlice)
	if err != nil {
		log.Println(err, "encoding")
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	return
}
