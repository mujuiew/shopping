package models

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_InsertProduct1(t *testing.T) {
	// Create a new instance of *MockDB
	mockDB := new(DB1)
	mockDB.Testerror = errors.New("insert error1")
	db := sql.DB{}

	// Call the function with mock DB
	er := InsertProduct(mockDB, &db, "ac001", "Cock", 20, 20, "cock.jpg")
	// Assert that the first parameter in the call was SQL query
	assert.Equal(t, "insert error1", er.Error())
	// assert.Equal(t, "insert error", er)
}

func Test_InsertProduct2(t *testing.T) {
	// Create a new instance of *MockDB
	mockDB := new(DB1)
	mockDB.Testerror2 = errors.New("insert error2")
	db := sql.DB{}

	// Call the function with mock DB
	er := InsertProduct(mockDB, &db, "ac001", "Cock", 20, 20, "cock.jpg")
	// Assert that the first parameter in the call was SQL query
	assert.Equal(t, "insert error2", er.Error())
	// assert.Equal(t, "insert error", er)
}
func Test_InsertProduct(t *testing.T) {
	// Create a new instance of *MockDB
	mockDB := new(DB1)
	db := sql.DB{}

	// Call the function with mock DB
	er := InsertProduct(mockDB, &db, "ac001", "Cock", 20, 20, "cock.jpg")
	// Assert that the first parameter in the call was SQL query
	assert.Equal(t, nil, er)
}

// DB1 ...
type DB1 struct {
	Testerror  error
	Testerror2 error
}

// Exec ...
func (bds DB1) Exec(db *sql.DB, sql string, args ...interface{}) error {

	if sql == `INSERT INTO public.product("Account_id", product_name, product_img, product_price, product_amount)	VALUES ($1, $2,$3,$4,$5)` {
		return bds.Testerror
	} else if sql == `INSERT INTO public.product("Account_id", product_name, product_img, product_price) VALUES ($1, $2,$3,$4)` {
		return bds.Testerror2
	} else {
		return nil
	}

}
