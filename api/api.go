package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/mujuiew/api-shopping/models"
	"github.com/mujuiew/api-shopping/structtype"
)

const (
	// host     = "172.17.106.172"
	host     = "192.168.88.193"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "shopping"
)

// Reister ...
func Reister(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()

	w.Header().Set("content-type", "application/json; charset=UTF-8;")
	w.Header().Set("x-request-id", r.Header.Get("x-request-id"))
	w.Header().Set("x-job-id", r.Header.Get("x-job-id"))
	w.Header().Set("datetime", time.Now().Format(time.RFC3339))
	w.Header().Set("x-roundtrip", "0.03")
	var in structtype.Input
	_ = json.NewDecoder(r.Body).Decode(&in)

	id, txt := models.InsertAccount(db, in.IAccountFirsteName, in.IAccountLastName, in.IAccountEmail, in.IAccountPhone)

	output := structtype.Output{id, txt}
	js, err := json.Marshal(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

// Store ...
func Store(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()

	w.Header().Set("content-type", "application/json; charset=UTF-8;")
	w.Header().Set("x-request-id", r.Header.Get("x-request-id"))
	w.Header().Set("x-job-id", r.Header.Get("x-job-id"))
	w.Header().Set("datetime", time.Now().Format(time.RFC3339))
	w.Header().Set("x-roundtrip", "0.03")

	query := r.URL.Query()
	store, present := query["id"]
	Instore := query.Get("id")
	if !present || len(store) == 0 {
		fmt.Println("store not present")
		w.WriteHeader(400)
		w.Write([]byte("store not present"))
	}
	p1, p2, p3, p4, p5 := models.GetStore(db, Instore)

	output := structtype.OutputStore{p1, p2, p3, p4, p5}
	js, err := json.Marshal(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

// AddProduct ...
func AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=UTF-8;")
	w.Header().Set("x-request-id", r.Header.Get("x-request-id"))
	w.Header().Set("x-job-id", r.Header.Get("x-job-id"))
	w.Header().Set("datetime", time.Now().Format(time.RFC3339))
	w.Header().Set("x-roundtrip", "0.03")

	var in structtype.InputProdut
	_ = json.NewDecoder(r.Body).Decode(&in)
	dbcon := models.DB{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBname:   dbname,
	}
	db := models.Connect(dbcon)
	obj := models.DBs{}                                                                                         //struct
	text := models.InsertProduct(obj, db, in.IProAccount, in.IProName, in.IProPrice, in.IProAmount, in.IProImg) //โยน struct

	output := structtype.OutputPro{text}
	js, err := json.Marshal(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
