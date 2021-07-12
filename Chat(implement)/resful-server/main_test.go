package main

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var cookie []*http.Cookie

func TestHandleGetLogin(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleGetLogin)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestHandleGetRegister(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", HandleGetRegister)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/register", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestHandleLogin(t *testing.T) {
	// initial function HandleLogin()
	mux := http.NewServeMux()
	mux.HandleFunc("/login", HandleLogin)

	// Testcase 1: Username or password true
	payload1 := strings.NewReader("{\"username\": \"lam\",\"passwords\": \"3026972\"}")
	writer1 := httptest.NewRecorder()
	request1, _ := http.NewRequest("POST", "/login", payload1)
	request1.Header.Add("Content-Type", "application/json")
	mux.ServeHTTP(writer1, request1)

	if writer1.Code != 200 {
		t.Errorf("Response code is %v", writer1.Code)
	}

	// Testcase 2: Username or password failure
	payload2 := strings.NewReader("{\"username\": \"lam@\",\"passwords\": \"3026972\"}")
	writer2 := httptest.NewRecorder()
	request2, _ := http.NewRequest("POST", "/login", payload2)
	request2.Header.Add("Content-Type", "application/json")
	mux.ServeHTTP(writer2, request2)

	if writer2.Code != 500 {
		t.Errorf("Response code is %v", writer2.Code)
	}

	cookie = writer1.Result().Cookies()

}

func TestHandleRegister(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", HandleRegister)

	// Testcase 1: Username has been registered on database
	payload := strings.NewReader("{\"id\":123,\"username\":\"kienzcv\",\"passwords\":\"12345\",\"firstname\":\"kjzxvkjas\",\"lastname\":\"lkjbzxlv\"}")
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/register", payload)
	request.Header.Add("Content-Type", "application/json")
	mux.ServeHTTP(writer, request)

	if writer.Code != 500 {
		t.Errorf("Response code is %v", writer.Code)
	}

	// Testcase 2: Username has not been registered on database
	payload2 := strings.NewReader("{\"id\":123,\"username\":\"" + randomString(10) + "\",\"passwords\":\"12345\",\"firstname\":\"kjzxvkjas\",\"lastname\":\"lkjbzxlv\"}")
	writer2 := httptest.NewRecorder()
	request2, _ := http.NewRequest("POST", "/register", payload2)
	request2.Header.Add("Content-Type", "application/json")
	mux.ServeHTTP(writer2, request2)
	if writer2.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func TestHandleHome(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", HandleHome)

	// Testcase 3: Send true coockie
	writer2 := httptest.NewRecorder()
	request2, _ := http.NewRequest("GET", "/home", nil)
	request2.AddCookie(cookie[0])
	mux.ServeHTTP(writer2, request2)

	if writer2.Code != 200 {
		t.Errorf("Response code is %v", writer2.Code)
	}


	// Testcase 1: Do not send cookie
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/home", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 401 && writer.Code != 400 {
		t.Errorf("Response code is %v", writer.Code)
	}

	// Testcase 2: Send empty or error token
	writer1 := httptest.NewRecorder()
	request1, _ := http.NewRequest("GET", "/home", nil)
	cookie1 := http.Cookie{Name: "session_token", Value: "John", Path: "/", Expires: time.Now().Add(10 * time.Minute) , MaxAge: 90000}
	request1.AddCookie(&cookie1)
	mux.ServeHTTP(writer1, request1)

	if writer.Code != 401 && writer.Code != 400 {
		t.Errorf("Response code is %v", writer.Code)
	}
}

func randomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
