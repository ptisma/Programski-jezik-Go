package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type webResponseAdd struct {
	Number1 int `json:"n1"` // input number1
	Number2 int `json:"n2"` // input number2
	Result  int `json:"r"`  // result of number1 + number2
}

func main() {
	http.HandleFunc("/add/", add)
	http.ListenAndServe(":8080", nil)
}

func add(w http.ResponseWriter, r *http.Request) {
	// tijelo funkcije
	rez := r.URL.Query()
	if len(rez) == 2 {
		if len(rez.Get("n1")) > 0 && len(rez.Get("n2")) > 0 {

			n1, err1 := strconv.Atoi(rez.Get("n1"))
			n2, err2 := strconv.Atoi(rez.Get("n2"))
			if err1 == nil && err2 == nil {
				w.Header().Set("Content-Type", "application/json")
				response := webResponseAdd{n1, n2, n1 + n2}
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(response)
				return
			}
		}
	}
	//Return error
	w.WriteHeader(http.StatusBadRequest)
	return

}
