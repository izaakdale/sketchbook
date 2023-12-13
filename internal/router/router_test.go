package router_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/izaakdale/sketchbook/internal/router"
)

type sadTestDB struct{}

// Insert implements router.DBConn.
func (*sadTestDB) Insert(key string, value []byte) error {
	panic("unimplemented")
}

type happyTestDB struct{}

// Insert implements router.DBConn.
func (*happyTestDB) Insert(key string, value []byte) error {
	panic("unimplemented")
}

func (t *sadTestDB) Ping() error {
	return errors.New("db is down")
}

func (t *happyTestDB) Ping() error {
	return nil
}

func TestRouter(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		mux := router.New(&happyTestDB{})
		rec := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/ping", nil)
		if err != nil {
			t.Fail()
		}

		mux.ServeHTTP(rec, req)
		if rec.Result().StatusCode != http.StatusOK {
			t.Fail()
		}

		rec = httptest.NewRecorder()
		req, err = http.NewRequest(http.MethodGet, "/somethingelse", nil)
		if err != nil {
			t.Fail()
		}

		mux.ServeHTTP(rec, req)
		if rec.Result().StatusCode != http.StatusNotFound {
			t.Fail()
		}
	})

	t.Run("sad path", func(t *testing.T) {
		mux := router.New(&sadTestDB{})
		rec := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/ping", nil)
		if err != nil {
			t.Fail()
		}

		mux.ServeHTTP(rec, req)

		if rec.Result().StatusCode != http.StatusInternalServerError {
			t.Error("failing DB does not cause http failure")
		}
	})
}
