package main

import (
	"avitoTest/server/internal/app"
	"log"
	"net/http"
)

func main() {
	s := app.New()
	log.Println("HTTP Running")
	defer s.Close()
	http.HandleFunc(app.GetBalanceEndPoint, s.GetBalanceHandler)
	http.HandleFunc(app.AddBalanceEndPoint, s.AddBalanceHandler)
	http.HandleFunc(app.ReserveEndPoint, s.ReserveHandler)
	http.HandleFunc(app.AcquireEndPoint, s.AcquireHandler)
	http.HandleFunc(app.ReportEndPoint, s.IncomeReportHandler)
	http.HandleFunc(app.DetailEndPoint, s.DetailHandler)
	log.Fatal(http.ListenAndServe(app.Addr, nil))
}
