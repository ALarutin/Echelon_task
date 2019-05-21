package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/ALarutin/Echelon_task/router"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestCase struct {
	Number       uint64   `json:"-"`
	Sites        []string `json:"sites"`
	SearchText   string   `json:"search-text"`
	WrongJSON    string   `json:"-"`
	FoundAtSites string   `json:"-"`
	Error        string   `json:"-"`
}

func TestCheckText_Successful(t *testing.T) {
	tests := []TestCase{
		{
			Number:       1,
			Sites:        []string{"https://google.com", "https://yahoo.com"},
			SearchText:   "Google",
			FoundAtSites: `{"found-at-sites":["https://google.com"]}`,
		},
		{
			Number:       2,
			Sites:        []string{"https://google.com", "https://yahoo.com"},
			SearchText:   "yAhOO",
			FoundAtSites: `{"found-at-sites":["https://yahoo.com"]}`,
		},
		{
			Number:       3,
			Sites:        []string{"https://google.com", "https://yahoo.com", "https://google.music.com", "https://mail.google.com"},
			SearchText:   "GOOGLE",
			FoundAtSites: `{"found-at-sites":["https://google.com","https://google.music.com","https://mail.google.com"]}`,
		},
	}

	r := router.GetRouter()

	for _, test := range tests {
		data, err := json.Marshal(test)
		if err != nil {
			t.Errorf("Failed when marshiling test case: %s", err.Error())
			return
		}
		req, err := http.NewRequest("POST", "/checkText", bytes.NewBuffer(data))
		if err != nil {
			t.Errorf("Failed when creating new request: %s", err.Error())
			return
		}
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		if status := rec.Code; status != http.StatusOK {
			t.Errorf("Hander returned wrong status in %d case!\nwanted: %v\ngot: %v",
				test.Number, http.StatusOK, status)
		}
		if response := rec.Body.String(); response != test.FoundAtSites {
			t.Errorf("Hander returned wrong string in %d case!\nwanted: %v\ngot: %v",
				test.Number, test.FoundAtSites, response)
		}
	}
}

func TestCheckText_Unsuccessful_BadRequest(t *testing.T) {
	tests := []TestCase{
		{
			Number:    1,
			WrongJSON: `{"sites:[https://google.com","https://yahoo.com"],"search-text":"Google"}`,
			Error:     `{"error":"invalid character ',' after object key"}`,
		},
		{
			Number:       2,
			WrongJSON: `{"sites"[https://google.com","https://yahoo.com"],"search-text":"Google"}`,
			Error:     `{"error":"invalid character '[' after object key"}`,
		},
		{
			Number:       3,
			WrongJSON: `{"sites":["https://google.com","https://yahoo.com"],"search-text":"Google"`,
			Error:     `{"error":"unexpected EOF"}`,
		},
	}

	r := router.GetRouter()

	for _, test := range tests {
		req, err := http.NewRequest("POST", "/checkText", bytes.NewBuffer([]byte(test.WrongJSON)))
		if err != nil {
			t.Errorf("Failed when creating new request: %s", err.Error())
			return
		}
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		if status := rec.Code; status != http.StatusBadRequest {
			t.Errorf("Hander returned wrong status in %d case!\nwanted: %v\ngot: %v",
				test.Number, http.StatusBadRequest, status)
		}
		if response := rec.Body.String(); response != test.Error {
			t.Errorf("Hander returned wrong error in %d case!\nwanted: %v\ngot: %v",
				test.Number, test.Error, response)
		}
	}
}

func TestCheckText_Unsuccessful_NoContent(t *testing.T) {
	tests := []TestCase{
		{
			Number:       1,
			Sites:        []string{"https://google.com", "https://yahoo.com"},
			SearchText:   "Yandex",
		},
		{
			Number:       2,
			Sites:        []string{"https://google.com", "https://yahoo.com"},
			SearchText:   "Googel",
		},
	}

	r := router.GetRouter()

	for _, test := range tests {
		data, err := json.Marshal(test)
		if err != nil {
			t.Errorf("Failed when marshiling test case: %s", err.Error())
			return
		}
		req, err := http.NewRequest("POST", "/checkText", bytes.NewBuffer(data))
		if err != nil {
			t.Errorf("Failed when creating new request: %s", err.Error())
			return
		}
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		if status := rec.Code; status != http.StatusNoContent {
			t.Errorf("Hander returned wrong status in %d case!\nwanted: %v\ngot: %v",
				test.Number, http.StatusNoContent, status)
		}
	}
}
