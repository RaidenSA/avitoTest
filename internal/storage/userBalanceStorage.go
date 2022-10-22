package storage

import (
	"database/sql"
	"log"
)

type DataBase struct {
	ConnStr string
}

func (db DataBase) GetBalance(userID string) (int, bool, error) {
	//open connection
	connection, err := sql.Open("postgres", db.ConnStr)
	if err != nil {
		return 0, false, err
	}
	defer func(connection *sql.DB) {
		err := connection.Close()
		if err != nil {
			log.Fatal(err, "defer error")
		}
	}(connection)
	row := connection.QueryRow("select userBalance from avitotest.public.balance where userID = $1", userID)
	balance := 0
	err = row.Scan(&balance)
	if err == sql.ErrNoRows {
		return 0, false, err
	}
	return balance, true, nil
}

func (db DataBase) InsertBalance(userID string, newBalance int) error {
	connection, err := sql.Open("postgres", db.ConnStr)
	if err != nil {
		return err
	}
	defer func(connection *sql.DB) {
		err := connection.Close()
		if err != nil {
			log.Fatal(err, "defer error")
		}
	}(connection)
	_, err = connection.Exec("insert into avitotest.public.balance (userID, userBalance) values ($1, $2)", userID, newBalance)
	if err != nil {
		log.Fatal(err, "connection insert error")
		return err
	}
	return nil
}

func (db DataBase) UpdateBalance(userID string, newBalance int) error {
	connection, err := sql.Open("postgres", db.ConnStr)
	if err != nil {
		return err
	}
	defer func(connection *sql.DB) {
		err := connection.Close()
		if err != nil {
			log.Fatal(err, "defer error")
		}
	}(connection)
	_, err = connection.Exec("update avitotest.public.balance set userBalance = $2 where userID = $1", userID, newBalance)
	if err != nil {
		log.Fatal(err, "connection insert error")
		return err
	}
	return nil
}
