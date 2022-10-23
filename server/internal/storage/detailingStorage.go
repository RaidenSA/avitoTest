package storage

import (
	"database/sql"
	"log"
)

func (db DataBase) Detailing(userID int64, limit int, page int) (map[int]DetailStruct, error) {
	//open connection
	connection, err := sql.Open("postgres", db.ConnStr)
	if err != nil {
		return nil, err
	}
	defer func(connection *sql.DB) {
		err := connection.Close()
		if err != nil {
			log.Fatal(err, "defer error")
		}
	}(connection)
	rows, err := connection.Query("select * from avitotest.public.transactions where userid=$1 order by created desc, sum desc limit $2 offset $3;", userID, limit, limit*(page-1))
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err, "defer error")
		}
	}(rows)

	res := make(map[int]DetailStruct)
	detail := DetailStruct{}
	i := 1 + (page-1)*limit
	for ; rows.Next(); i += 1 {
		err = rows.Scan(&detail.UserID, &detail.ServiceID, &detail.OrderID, &detail.Sum, &detail.Created, &detail.Source, &detail.TransactionType)
		if err != nil {
			log.Println("scan error", err)
			return nil, err
		}
		//res[i]=[]string{fmt.Sprintf("serviceID: %d",serviceID), fmt.Sprintf("income: %.2f",sum)}
		res[i] = detail
	}
	err = rows.Err()
	if err != nil {
		log.Println("rows error", err)
		return nil, err
	}
	return res, nil
}
