package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	types "resful-server/datastruct"
	db "resful-server/userdata"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var jwtKey = []byte("my_secret_key")

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HandleGetLogin).Methods("GET")
	router.HandleFunc("/login", HandleLogin).Methods("POST")
	router.HandleFunc("/register", HandleGetRegister).Methods("GET")
	router.HandleFunc("/register", HandleRegister).Methods("POST")
	router.HandleFunc("/home", HandleHome).Methods("GET")

	router.PathPrefix("/js/").Handler(
		http.StripPrefix("/js/", http.FileServer(http.Dir("./public/js"))))

	fmt.Println("Server start at port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func HandleGetLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/login.html")
}

func HandleGetRegister(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/register.html")
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	// Lấy thông tin từ body của request
	var user types.User
	// Chuyển đổi json
	err := json.NewDecoder(r.Body).Decode(&user)
	httpErr(err, w)
	log.Printf("Sign in : %v", user)

	userData := db.SelectPasswords(user.Username)

	if userData.Passwords == user.Passwords {
		log.Printf("Sending : %v", userData)
		// Sinh token
		expirationTime := time.Now().Add(5 * time.Minute)

		claims := &types.Claims{
			Username: user.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString(jwtKey)
		httpErr(err, w)
		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   tokenString,
			Expires: expirationTime,
		})
		json.NewEncoder(w).Encode(userData)
	} else {
		log.Printf("Sending : %v", "Invalid username or password!")
		http.Error(w, "Invalid username or password", 500)
	}
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var user types.User
	// Lấy thông tin từ body của request
	err := json.NewDecoder(r.Body).Decode(&user)
	log.Printf("Sign up : %v", user)
	httpErr(err, w)

	str := db.InsertUser(user)

	if str == "Register successfully" {
		log.Printf("Sending : %v", "Register successfully!")
		fmt.Fprintf(w, "Register successfully!")
	} else {
		log.Printf("Sending : %v", "Register failure, please check your username!")
		http.Error(w, "Register failure, please check your username!", 500)
	}
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := cookie.Value
	claims := &types.Claims{}

	// Kiểm tra token
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	http.ServeFile(w, r, "./public/home.html")
}

func httpErr(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
