package main

import (
	"avitoTest/server/internal/app"
	"log"
	"net/http"
)

func main() {
	s := app.New()
	defer app.Close(s)
	log.Println("HTTP Running")
	/*
		go func(){
			http.HandleFunc(app.GetBalanceEndPoint, s.GetBalanceHandler)
			http.HandleFunc(app.AddBalanceEndPoint, s.AddBalanceHandler)
			http.HandleFunc(app.ReserveEndPoint, s.ReserveHandler)
			http.HandleFunc(app.AcquireEndPoint, s.AcquireHandler)
			http.HandleFunc(app.ReportEndPoint, s.IncomeReportHandler)
			http.HandleFunc(app.DetailEndPoint, s.DetailHandler)
			log.Fatal(http.ListenAndServe(app.Addr, nil))
		}()

	*/
	defer app.Close(s)
	http.HandleFunc(app.GetBalanceEndPoint, s.GetBalanceHandler)
	http.HandleFunc(app.AddBalanceEndPoint, s.AddBalanceHandler)
	http.HandleFunc(app.ReserveEndPoint, s.ReserveHandler)
	http.HandleFunc(app.AcquireEndPoint, s.AcquireHandler)
	http.HandleFunc(app.ReportEndPoint, s.IncomeReportHandler)
	http.HandleFunc(app.DetailEndPoint, s.DetailHandler)
	log.Fatal(http.ListenAndServe(app.Addr, nil))
}
