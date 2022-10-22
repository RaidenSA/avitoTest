package storage

type BalanceStorage interface {
	GetBalance(string) (int, bool, error) // user ID, return money
	InsertBalance(string, int) error
	UpdateBalance(string, int) error // user ID, money
	//Reserve(string, string, string, Money) error // user ID, service ID, order ID, money
	//UnReserve(string,string,string,Money) error // user ID, service ID, order ID, money
	//Acquire(string,string,string,Money) error // user ID, service ID, order ID, money
}

type ReserveStorage interface {
	Reserve(string, string, string, int) error   // user ID, service ID, order ID, money
	UnReserve(string, string, string, int) error // user ID, service ID, order ID, money
}

type FinalStorage interface {
	Acquire(string, string, string, int) error // user ID, service ID, order ID, money
}
