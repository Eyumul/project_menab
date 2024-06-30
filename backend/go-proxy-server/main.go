package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func handleInitializeTransaction(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Handle preflight requests
	if r.Method == http.MethodOptions {
		return
	}

	url := "https://api.chapa.co/v1/transaction/initialize"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	req.Header.Set("Content-Type", r.Header.Get("Content-Type"))

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(res.StatusCode)
	w.Write(body)
}

func main() {
	http.HandleFunc("/initialize-transaction", handleInitializeTransaction)
	log.Println("Server is running on port 3001")
	log.Fatal(http.ListenAndServe(":3001", nil))
}
