package storage

import (
	"context"
	"log"
)

func (db DataBase) GetBalance(userID int64) (float32, bool, error) {
	//open connection
	/*
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

	*/
	row := db.Db.QueryRow("select userBalance from avitotest.public.balance where userID = $1", userID)
	var balance float32
	err := row.Scan(&balance)
	if err != nil {
		return 0, false, err
	}
	return balance, true, nil
}

func (db DataBase) InsertBalance(userID int64, newBalance float32) error {
	/*
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
	*/
	ctx := context.Background()
	tx, err := db.Db.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err, "transaction error")
		return err
	}
	_, err = tx.ExecContext(ctx, "insert into avitotest.public.balance (userID, userBalance) values ($1, $2)", userID, newBalance)
	if err != nil {
		log.Println(err, " insert error")
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}
	_, err = tx.ExecContext(ctx, "insert into avitotest.public.transactions (userID, sum, type) values ($1, $2, $3)", userID, newBalance, "NEW USER")
	if err != nil {
		log.Println(err, " log error")
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}
	err2 := tx.Commit()
	if err2 != nil {
		return err2
	}
	return nil
}

func (db DataBase) UpdateBalance(userID int64, newBalance float32) error {
	/*connection, err := sql.Open("postgres", db.ConnStr)
	if err != nil {
		return err
	}
	defer func(connection *sql.DB) {
		err := connection.Close()
		if err != nil {
			log.Fatal(err, "defer error")
		}
	}(connection)
	*/
	ctx := context.Background()
	tx, err := db.Db.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err, "transaction error")
		return err
	}
	_, err = tx.ExecContext(ctx, "update avitotest.public.balance set userBalance = $2 where userID = $1", userID, newBalance)
	if err != nil {
		log.Println(err, " insert error")
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}
	_, err = tx.ExecContext(ctx, "insert into avitotest.public.transactions (userID, sum, type) values ($1, $2,$3)", userID, newBalance, "CREDITING")
	if err != nil {
		log.Println(err, " log error")
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}
	err2 := tx.Commit()
	if err2 != nil {
		return err2
	}
	return nil
}
