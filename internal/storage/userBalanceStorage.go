package storage

import (
	"context"
	"database/sql"
	"log"
)

type DataBase struct {
	ConnStr string
}

func (db DataBase) GetBalance(userID int64) (float32, bool, error) {
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
	var balance float32
	err = row.Scan(&balance)
	if err == sql.ErrNoRows {
		return 0, false, err
	}
	return balance, true, nil
}

func (db DataBase) InsertBalance(userID int64, newBalance float32) error {
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
		log.Println(err, "connection insert error")
		return err
	}
	return nil
}

func (db DataBase) UpdateBalance(userID int64, newBalance float32) error {
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
		log.Println(err, "connection insert error")
		return err
	}
	return nil
}

func (db DataBase) Reserve(userID int64, serviceID int64, orderID int64, sum float32) error {
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
	ctx := context.Background()
	tx, err := connection.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err, "transaction error")
		return err
	}

	_, err = tx.ExecContext(ctx, "update avitotest.public.balance set userBalance = userBalance - $2 where userID = $1", userID, sum)
	if err != nil {
		log.Println(err, "transaction update error")
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}
	_, err = tx.ExecContext(ctx, "insert into avitotest.public.reserved values ($1,$2,$3,$4)", userID, serviceID, orderID, sum)
	if err != nil {
		log.Println(err, "transaction update error")
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err, "transaction error")
		return err
	}
	return nil
}

func (db DataBase) Acquire(userID int64, serviceID int64, orderID int64, sum float32) error {
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
	ctx := context.Background()
	tx, err := connection.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err, "transaction error")
		return err
	}
	_, err = tx.ExecContext(ctx, "update avitotest.public.reserved set sum = sum- $4 where userID = $1 and serviceID = $2 and orderID =$3 ", userID, serviceID, orderID, sum)
	if err != nil {
		log.Println(err, "transaction update error")
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}
	_, err = tx.ExecContext(ctx, "insert into avitotest.public.finished values ($1,$2,$3,$4)", userID, serviceID, orderID, sum)
	if err != nil {
		log.Println(err, "transaction update error")
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err, "transaction error")
		return err
	}
	return nil
}
