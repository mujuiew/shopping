package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/mujuiew/api-shopping/structtype"
)

const (
	// host = "192.168.88.193"
	host     = "172.20.0.2"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "shopping"
)

// InsertAccount ...
func InsertAccount(db *sql.DB, fname string, lname string, mail string, phone string) (string, string) {

	rows, err := db.Query(`SELECT "ID" FROM Account ORDER BY "ID" DESC LIMIT 1`)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		if err := rows.Scan(&structtype.LastID.AccountID); err != nil {
			log.Fatal(err)
		}
	}
	id := ""
	if structtype.LastID.AccountID == "" {
		id = "ac001"
	} else {
		str1 := structtype.LastID.AccountID[2:5]
		intVar1, _ := strconv.ParseInt(str1, 10, 32)
		int := fmt.Sprintf("%v", intVar1+1)

		if len(int) == 1 {
			id = "ac00" + int
		} else if len(int) == 2 {
			id = "ac0" + int
		} else if len(int) == 3 {
			id = "ac" + int
		}
	}

	queryStr := `INSERT INTO Account("ID", account_fname, account_lname, account_email, account_phone)
		VALUES ($1, $2,$3,$4,$5)`
	_, err = db.Exec(queryStr, id, fname, lname, mail, phone)
	if err != nil {
		// panic(err)
		return "", "Fail"
	}
	return id, "Successfully"

}

// GetStore ...
func GetStore(db *sql.DB, id string) (string, string, string, int, int) {
	rows, err := db.Query(`SELECT "ID", product_name, product_img, product_price, product_amount FROM product WHERE "Account_id" = '` + id + `'`)

	if err != nil {
		panic(err)
	}
	for rows.Next() {
		if err := rows.Scan(&structtype.PROID.ProductID, &structtype.PROID.ProductName, &structtype.PROID.ProductImg, &structtype.PROID.ProductPrice, &structtype.PROID.ProductAmount); err != nil {
			log.Fatal(err)
		}
	}
	// fmt.Println(structtype.PROID.ProductID)

	ProductID := structtype.PROID.ProductID
	ProductName := structtype.PROID.ProductName
	ProductImg := structtype.PROID.ProductImg
	ProductPrice := structtype.PROID.ProductPrice
	ProductAmount := structtype.PROID.ProductAmount

	return ProductID, ProductName, ProductImg, ProductPrice, ProductAmount

}

// API ...
type API interface {
	Exec(db *sql.DB, sql string, args ...interface{}) error
}

// InsertProduct ...
func InsertProduct(api API, db *sql.DB, ac string, pname string, pprice int, pamount int, pimg string) error {

	queryStr := `INSERT INTO public.product("Account_id", product_name, product_img, product_price, product_amount)	VALUES ($1, $2,$3,$4,$5)`
	err := api.Exec(db, queryStr, ac, pname, pprice, pamount, pimg)
	if err != nil {
		// panic(err)
		return err
	}
	queryStr2 := `INSERT INTO public.product("Account_id", product_name, product_img, product_price) VALUES ($1, $2,$3,$4)`
	err2 := api.Exec(db, queryStr2, ac, pname, pprice, pamount)

	if err2 != nil {
		// panic(err)
		return err2
	}

	return nil

}

// DBs ...
type DBs struct {
}

// Exec ...
func (bds DBs) Exec(db *sql.DB, sql string, args ...interface{}) error {

	// queryStr := sql
	_, err := db.Exec(sql, args)
	if err != nil {
		// panic(err)
		return err
	}
	return nil

}
