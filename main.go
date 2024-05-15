package main

import (
	"fmt"
	"net/http"
	"regexp"
)

type Person struct {
	Name  string `json:"name"`
	IIN   string `json:"IIN"`
	Phone string `json:"phone"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Errors  string `json:"errors"`
}

func main() {
	http.Handler("iin/check", iinCheckHanlder)
	http.Handler("people/info", peopleInfoHandler)
	http.Handler("people/info/iin", personinfoByIINHandler)
	http.Handler("people/info/phone", personInfoByNameHandler)

	fmt.Println("Server is running")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	}

}

func iinCheckHanlder(w http.ResponseWriter, r *http.Request) {
	iin := r.URL.Path[len("iin_check"):]

	if isValidIIN(iin) {
		sex := "male"
		if iin[6] >= '5' {
			sex = "female"
		}
	}

}

func isValidIIN(iin string) bool {
	match, _ := regexp.MatchString(`^d{12}$`, iin)
	if !match {
		return false
	}
	return true
}
