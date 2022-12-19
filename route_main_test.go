package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Get_Index(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	body := writer.Body.String()
	if strings.Contains(body, "スレッドを開始するか") == false {
		t.Errorf("Body does not contain スレッドを開始するか")
	}
}

func Test_Get_Err(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/err", err)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/err?msg=test", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	body := writer.Body.String()
	if strings.Contains(body, "test") == false {
		t.Errorf("Body does not contain parameter")
	}
}
