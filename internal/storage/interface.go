package storage

type BalanceStorage interface {
	//should change string to int64, int to real
	//may be unique constrained
	GetBalance(int64) (float32, bool, error) // user ID, return money
	InsertBalance(int64, float32) error
	UpdateBalance(int64, float32) error         // user ID, money
	Reserve(int64, int64, int64, float32) error // user ID, service ID, order ID, money
	//UnReserve(string, string, string, int) error // user ID, service ID, order ID, money
	Acquire(int64, int64, int64, float32) error // user ID, service ID, order ID, money
}
