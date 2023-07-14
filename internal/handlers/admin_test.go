package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	url := "/"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Index)
	handler.ServeHTTP(rr, req)
	// fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(Index)
	handler.ServeHTTP(rr, req)
	// fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetUsersAll(t *testing.T) {
	url := "/admin"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserAll)
	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(GetUserAll)
	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

}
