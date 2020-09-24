package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mujuiew/api-shopping/models"
	"github.com/mujuiew/api-shopping/structtype"
)

// Reister ...
func Reister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=UTF-8;")
	w.Header().Set("x-request-id", r.Header.Get("x-request-id"))
	w.Header().Set("x-job-id", r.Header.Get("x-job-id"))
	w.Header().Set("datetime", time.Now().Format(time.RFC3339))
	w.Header().Set("x-roundtrip", "0.03")
	var in structtype.Input
	_ = json.NewDecoder(r.Body).Decode(&in)

	id, txt := models.InsertAccount(in.IAccountFirsteName, in.IAccountLastName, in.IAccountEmail, in.IAccountPhone)

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
	p1, p2, p3, p4, p5 := models.GetStore(Instore)

	output := structtype.OutputStore{p1, p2, p3, p4, p5}
	js, err := json.Marshal(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
