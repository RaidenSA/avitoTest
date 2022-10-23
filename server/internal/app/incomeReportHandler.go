package app

import (
	"encoding/csv"
	"log"
	"net/http"
	"time"
)

func (s *Server) IncomeReportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("Not correct method, instead of Get")
		http.Error(w, "Only Get method", http.StatusMethodNotAllowed)
		return
	}

	period := r.URL.String()
	period = period[(len(ReportEndPoint)):]
	_, err := time.Parse("2006-01", period)
	if err != nil {
		log.Println("Not correct date")
		http.Error(w, "Not correct date", http.StatusBadRequest)
		return
	}

	reportMap, keyCounter, err := s.Storage.Report(period)
	if err != nil {
		log.Println(err, "could not get report")
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=CSVFile.csv")
	wr := csv.NewWriter(w)
	for i := 0; i <= keyCounter; i++ {
		err = wr.Write(reportMap[i])
		if err != nil {
			log.Println(err, "could not send scv")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	wr.Flush()
	return
}
