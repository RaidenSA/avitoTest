package storage

import (
	"database/sql"
	"fmt"
	"log"
)

func (db DataBase) Report(period string) (map[int][]string, int, error) {
	/*
		connection, err := sql.Open("postgres", db.ConnStr)
		if err != nil {
			return nil, 0, err
		}
		defer func(connection *sql.DB) {
			err := connection.Close()
			if err != nil {
				log.Fatal(err, "defer error")
			}
		}(connection)
	*/
	rows, err := db.Db.Query("select serviceid, sum(sum) from avitotest.public.finished where updated < to_date($1,'YYYY-MM') and updated >to_date($1,'YYYY-MM')-interval '1 month' group by serviceid", period)
	if err != nil {
		return nil, 0, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err, "defer error")
		}
	}(rows)

	res := make(map[int][]string)
	var serviceID int64
	var sum float32
	i := 0
	for ; rows.Next(); i += 1 {
		err = rows.Scan(&serviceID, &sum)
		if err != nil {
			log.Println("scan error", err)
			return nil, i, err
		}
		res[i] = []string{fmt.Sprintf("serviceID: %d", serviceID), fmt.Sprintf("income: %.2f", sum)}
	}
	err = rows.Err()
	if err != nil {
		log.Println("rows error", err)
		return nil, i, err
	}
	return res, i, nil
}
