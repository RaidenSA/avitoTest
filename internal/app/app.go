package app

import (
	"avitoTest/internal/storage"
)

const Addr = "localhost:8080"
const user = "avitotestuser"
const myPass = "avitotestpass"
const dbname = "avitotest"
const connStr = "user=" + user + " password=" + myPass + " dbname=" + dbname + " sslmode=disable"
const GetBalanceEndPoint = "/balance/"
const AddBalanceEndPoint = "/addBalance"
const ReserveEndPoint = "/reserve"
const AcquireOrderEndPoint = "/acquire"

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
