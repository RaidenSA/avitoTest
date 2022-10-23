package storage

import "database/sql"

type DataBase struct {
	Db *sql.DB
}

type BalanceStorage interface {
	//may be unique constrained
	GetBalance(int64) (float32, bool, error) // user ID, return money
	InsertBalance(int64, float32) error
	UpdateBalance(int64, float32) error         // user ID, money
	Reserve(int64, int64, int64, float32) error // user ID, service ID, order ID, money
	//UnReserve(string, string, string, int) error // user ID, service ID, order ID, money
	Acquire(int64, int64, int64, float32) error              // user ID, service ID, order ID, money
	Report(string) (map[int][]string, int, error)            //YYYY-MM
	Detailing(int64, int, int) (map[int]DetailStruct, error) //userID,limit,page
}

type DetailStruct struct {
	UserID          int64   `json:"user_id"`
	ServiceID       int64   `json:"service_id"`
	OrderID         int64   `json:"order_id"`
	Sum             float32 `json:"sum"`
	Created         string  `json:"created"`
	Source          string  `json:"source"`
	TransactionType string  `json:"transaction_type"`
}
