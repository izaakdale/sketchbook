package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type DBConn interface {
	Ping() error
	Insert(key string, value []byte) error
}

func New(db DBConn) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", pingDB(db))
	mux.HandleFunc("/somedata", postDB(db))

	return mux
}

func pingDB(db DBConn) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			http.Error(w, fmt.Sprintf("error attempting ping to db: %v", err), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("pong!"))
	}
}

func postDB(db DBConn) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, err := json.Marshal(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("body error: %v", err), http.StatusBadRequest)
			return
		}
		id := uuid.NewString()
		db.Insert(id, bodyBytes)

		w.Write([]byte(id))
	}
}
