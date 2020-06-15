package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

type User struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

type Account struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Alias string `json:"alias,omitempty"`
	Owner User   `json:"owner,omitempty"`
}

type Transaction struct {
	ID          string  `json:"id,omitempty"`
	Amount      string  `json:"amount"`
	Source      Account `json:"source,omitempty"`
	Destination Account `json:"destination,omitempty"`
	DateTime    string  `json:"dateTime,omitempty"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/isalive", isAlive)

	router.HandleFunc("/api/v1/user/{userId}/account/{accountId}/transactions", doTranasaction).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/api/v1/user/{userId}/account/{accountId}/transactions", getTransactions).Methods(http.MethodGet, http.MethodOptions)

	spa := SpaHandler{StaticPath: "FRONTEND/build", IndexPath: "index.html"}
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		spa.ServeHTTP(w, r)
	})
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("FRONTEND/build"))))

	srv := &http.Server{
		Handler:      RequestLogger(router),
		Addr:         "127.0.0.1:5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println(srv.ListenAndServe())
}

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		targetMux.ServeHTTP(w, r)
		requesterIP := r.RemoteAddr
		println(
			"%s\t\t%s\t\t%s\t\t%v",
			r.Method,
			r.RequestURI,
			requesterIP,
			time.Since(start),
		)
	})
}

func isAlive(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ALIVE")
}

func allowCors(w http.ResponseWriter, r *http.Request) bool {

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*") //http://127.0.0.1:8000
	w.Header().Set("Access-Control-Allow-Methods", "POST,PUT,GET,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, Access-Control-Allow-Origin, access-control-allow-methods")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, PUT, OPTIONS")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return true
	}

	return false
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	if allowCors(w, r) {
		return
	}

	fmt.Println("GET")

	idUser, err := getParamFromPathUrl(r, "userId")
	idAccount, err := getParamFromPathUrl(r, "accountId")

	if err != nil {
		fmt.Println(err)
	}

	var ex Transaction
	fmt.Println(ex)

	fmt.Println(idUser)
	fmt.Println(idAccount)

	var sourcerUser User
	sourcerUser.ID = "1234"
	sourcerUser.Name = "Guillermo"

	var destinationUser User
	destinationUser.ID = "1234"
	destinationUser.Name = "Guillermo"

	var sourceAccount Account
	sourceAccount.ID = "132123"
	sourceAccount.Owner = sourcerUser
	sourceAccount.Alias = "HOUSE.TEA.MONTH"
	sourceAccount.Name = "CC"

	var destinationAccount Account
	destinationAccount.ID = "132123"
	destinationAccount.Owner = destinationUser
	destinationAccount.Alias = "HOUSE.TEA.MONTH"
	destinationAccount.Name = "CC"

	var trans Transaction
	trans.Amount = "100"
	trans.Source = sourceAccount
	trans.Destination = destinationAccount
	trans.ID = "432432"
	trans.DateTime = "Jun 14 10:00:00"

	var transactions []Transaction
	transactions = append(transactions, trans)
	transactions = append(transactions, trans)
	transactions = append(transactions, trans)
	transactions = append(transactions, trans)

	_ = json.NewEncoder(w).Encode(transactions)
}

func doTranasaction(w http.ResponseWriter, r *http.Request) {

	if allowCors(w, r) {
		return
	}

	fmt.Println("POST")

	idUser, err := getParamFromPathUrl(r, "userId")
	idAccount, err := getParamFromPathUrl(r, "accountId")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(idUser)
	fmt.Println(idAccount)

	result := make(map[string]interface{})
	result["result"] = "ok"

	_ = json.NewEncoder(w).Encode(result)

}

type SpaHandler struct {
	StaticPath string
	IndexPath  string
}

func (h SpaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.StaticPath, path)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.StaticPath, h.IndexPath))
		return
	} else if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.StaticPath)).ServeHTTP(w, r)
}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "json")
}

func getParamFromPathUrl(r *http.Request, paramId string) (string, error) {
	params := mux.Vars(r)
	sub := params[paramId]
	if sub == "" {
		return "", errors.New("Error: '" + paramId + "' param not found")
	}
	return sub, nil
}
