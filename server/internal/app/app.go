package app

import (
	"avitoTest/server/internal/storage"
)

const Addr = "localhost:8080"
const user = "avitotestuser"
const myPass = "avitotestpass"
const dbname = "avitotest"
const connStr = "user=" + user + " password=" + myPass + " dbname=" + dbname + " sslmode=disable"
const GetBalanceEndPoint = "/balance/"
const AddBalanceEndPoint = "/addBalance"
const ReserveEndPoint = "/reserve"
const AcquireEndPoint = "/acquire"

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
	Storage storage.BalanceStorage
}

func New() *Server {
	var stor storage.BalanceStorage
	stor = storage.DataBase{
		ConnStr: connStr,
	}
	s := &Server{
		Storage: stor,
	}

	return s
}
