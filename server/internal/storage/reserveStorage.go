package storage

import (
	"context"
	"log"
)

func (db DataBase) Reserve(userID int64, serviceID int64, orderID int64, sum float32) error {
	ctx := context.Background()
	tx, err := db.Db.BeginTx(ctx, nil)
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

	_, err = tx.ExecContext(ctx, "insert into avitotest.public.reserved (userid, serviceid, orderid, sum) values ($1,$2,$3,$4)", userID, serviceID, orderID, sum)
	if err != nil {
		log.Println(err, "transaction insert error")
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	_, err = tx.ExecContext(ctx, "insert into avitotest.public.transactions (userid,serviceid,orderid, sum) values ($1,$2,$3,$4)", userID, serviceID, orderID, sum)
	if err != nil {
		log.Println(err, "transaction insert error")
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
