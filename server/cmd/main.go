package main

import (
	"avitoTest/server/internal/app"
	"log"
	"net/http"
)

func main() {
	s := app.New()
	log.Println("HTTP Running")
	http.HandleFunc(app.GetBalanceEndPoint, s.GetBalanceHandler)
	http.HandleFunc(app.AddBalanceEndPoint, s.AddBalanceHandler)
	http.HandleFunc(app.ReserveEndPoint, s.ReserveHandler)
	http.HandleFunc(app.AcquireEndPoint, s.AcquireHandler)
	log.Fatal(http.ListenAndServe(app.Addr, nil))
}