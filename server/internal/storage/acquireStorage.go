package storage

import (
	"context"
	"log"
)

func (db DataBase) Acquire(userID int64, serviceID int64, orderID int64, sum float32) error {
	ctx := context.Background()
	tx, err := db.Db.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err, "transaction error")
		return err
	}

	res := tx.QueryRowContext(ctx, "select userid from avitotest.public.reserved where userid =$1 and serviceid = $2 and orderid = $3 and sum = $4", userID, serviceID, orderID, sum)
	var check int64
	err = res.Scan(&check)
	if err != nil {
		log.Println("no such transaction")
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

	_, err = tx.ExecContext(ctx, "insert into avitotest.public.finished(userid, serviceid, orderid, sum) values ($1,$2,$3,$4)", userID, serviceID, orderID, sum)
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
