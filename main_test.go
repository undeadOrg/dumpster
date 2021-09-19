package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	// Initialize Router
	m := Router()

	//fmt.Println(m.Get)

	// Setup Testserver
	ts := httptest.NewServer(m)
	defer ts.Close()

	// GET /ping - Should be successful
		res, err := http.Get(ts.URL + "/ping")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code for /ping is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}

	// POST /ping - Should Fail ie: negative test
	res, err = http.Post(ts.URL+"/ping", "text/plain", nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status code for /ping is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusMethodNotAllowed)
	}
}

