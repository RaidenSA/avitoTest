package app

import (
	"avitoTest/server/internal/storage"
	"database/sql"
	"log"
	"time"
)

const (
	Addr               = ":8080"
	user               = "avitotestuser"
	myPass             = "avitotestpass"
	dbname             = "avitotest"
	connStr            = "host=postgres" + " port=5432" + " user=" + user + " password=" + myPass + " dbname=" + dbname + " sslmode=disable"
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
		log.Println("db connection error")
		return nil, err
	}
	connectionHealthFlag := false
	for i := 0; i < 4; i++ {
		err = db.Ping()
		if err != nil {
			log.Println("db health check error")
			//return nil, err
			time.Sleep(3 * time.Second)
			continue
		}
		connectionHealthFlag = true
		break
	}
	if !connectionHealthFlag {
		log.Fatal("could not connect to db")
	}

	log.Println("connected to db successfully")
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
