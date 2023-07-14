package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestWithToken(t *testing.T) {
	url := "/"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RequestWithToken)
	handler.ServeHTTP(rr, req)
	// fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	url = "ws://localhost:8080/chat?token=_eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6ImVhNzhlNDhhLTM3MDMtNGFlMy1iMzI1LWRlMWY1YzhlNTIzOCIsImV4cCI6MTY4NjU2NzQ5MiwidXNlcm5hbWUiOiJsb2dpbjIifQ.he668j-vRSU-DsDcxh8o2mjMHKQhoPW3o4rCy8HRMME"
	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(RequestWithToken)
	handler.ServeHTTP(rr, req)
	// fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	url = "ws://localhost:8080/chat?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6ImVhNzhlNDhhLTM3MDMtNGFlMy1iMzI1LWRlMWY1YzhlNTIzOCIsImV4cCI6MTY4NjU2NzQ5MiwidXNlcm5hbWUiOiJsb2dpbjIifQ.he668j-vRSU-DsDcxh8o2mjMHKQhoPW3o4rCy8HRMME_"
	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(RequestWithToken)
	handler.ServeHTTP(rr, req)
	// fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

}
