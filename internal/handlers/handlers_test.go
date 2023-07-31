package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserCreate(t *testing.T) {
	url := "/user"
	payload := `
	{
		"password":""
	}
`
	req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UserCreate)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	payload = `
	{
		"password":"
	}
`
	req, err = http.NewRequest(http.MethodPost, url, strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(UserCreate)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	payload = `
	{
		"password":""
	}
`
	req, err = http.NewRequest(http.MethodPost, url, strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("2handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	payload = `
		{
			"username":"login1",
			"password":""
		}
		`
	req, err = http.NewRequest(http.MethodPost, url, strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	fmt.Println("Code = ", rr.Code)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("3handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	payload = `
	{
		"username":"login1",
		"password":"password1"
	}
	`
	req, err = http.NewRequest(http.MethodPost, url, strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	fmt.Println("Code = ", rr.Code)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("4handler return wrong status code: got %v want %v", status, http.StatusOK)
	}
	payload = `
	{
		"username":"login1",
		"password":"password1"
	}
	`
	req, err = http.NewRequest(http.MethodPost, url, strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	fmt.Println("Code = ", rr.Code)
	if status := rr.Code; status != http.StatusConflict {
		t.Errorf("5handler return wrong status code: got %v want %v", status, http.StatusConflict)
	}

}

func TestUserLogin(t *testing.T) {
	url := "/user/login"

	payload := `
			{
				"username":"login1",
				"password":"password1"
			}
		`
	req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UserLogin)
	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	payload = `
			{
				"username":"
			}
		`
	req, err = http.NewRequest(http.MethodPost, url, strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(UserLogin)
	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	payload = `
			{
				"username":"login1",
				"password":"password2"
			}
		`
	req, err = http.NewRequest(http.MethodPost, url, strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(UserLogin)
	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("1handler return wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	payload = `
			{
				"username":"login1",
				"password":"password1"
			}
		`
	req, err = http.NewRequest(http.MethodPost, url, strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(UserLogin)
	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler return wrong status code: got #{status} want #{http,statusOk}")
	}

	// req, err = http.NewRequest(http.MethodGet, url, strings.NewReader(payload))
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// rr = httptest.NewRecorder()
	// handler = http.HandlerFunc(UserLogin)
	// handler.ServeHTTP(rr, req)
	// fmt.Println("code = ", rr.Code)
	// if status := rr.Code; status != http.StatusBadRequest {
	// 	t.Errorf("handler return wrong status code: got #{status} want #{http,statusOk}")
	// }
}

