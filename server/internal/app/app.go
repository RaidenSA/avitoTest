package app

import (
	"avitoTest/server/internal/storage"
	"database/sql"
	"log"
)

const (
	Addr               = "localhost:8080"
	user               = "avitotestuser"
	myPass             = "avitotestpass"
	dbname             = "avitotest"
	connStr            = "user=" + user + " password=" + myPass + " dbname=" + dbname + " sslmode=disable"
	GetBalanceEndPoint = "/balance/"
	AddBalanceEndPoint = "/addBalance"
	ReserveEndPoint    = "/reserve"
	AcquireEndPoint    = "/acquire"
	ReportEndPoint     = "/report/"
	DetailEndPoint     = "/detail/"
)

type balanceStruct struct {
	UserID  int64
	Balance float32
}

/*
{"userID" : 2,
 "Balance" : 400}
*/

type transactionStruct struct {
	UserID    int64
	ServiceID int64
	OrderID   int64
	Sum       float32
}

/*
{"userID" : 2,
"serviceID":1,
"orderID":1,
"sum" : 200}
*/

type Server struct {
	Storage storage.DataBase
}

func ConnDB(conStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func New() *Server {
	db, err := ConnDB(connStr)
	if err != nil {
		log.Fatal(err)
	}
	s := &Server{
		storage.DataBase{db},
	}
	return s
}

func (s *Server) Close() {
	err := s.Storage.Db.Close()
	if err != nil {
		log.Fatal(err, "defer error")
	}
	return
}
