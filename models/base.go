package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/mujuiew/api-shopping/structtype"
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
