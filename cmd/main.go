package main

import (
	"avitoTest/internal/app"
	"log"
	"net/http"
)

func main() {
	s := app.New()
	log.Println("HTTP Running")
	http.HandleFunc(app.GetBalanceEndPoint, s.GetBalanceHandler)
	http.HandleFunc(app.AddBalanceEndPoint, s.AddBalanceHandler)
	http.HandleFunc(app.ReserveEndPoint, s.ReserveHandler)
	log.Fatal(http.ListenAndServe(app.Addr, nil))
}
