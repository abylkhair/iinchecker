package main

import (
	"database/sql"
	"encoding/json"
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

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Создаем таблицу people в базе данных, если она не существует
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS people (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT,
            iin TEXT,
            phone TEXT
        )
    `)
	if err != nil {
		panic(err)
	}

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

		dateOfBirth := fmt.Sprintf("%s.%s.%s", iin[0:2], iin[2:4], iin[4:6])
		json.NewEncoder(w).Encode(map[string]interface{}{
			"correct":       true,
			"sex":           sex,
			"date_of_birth": dateOfBirth,
		})
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"correct": false,
		})
	}
}
func peopleInfoHandler(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Success: false,
			Errors:  err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

func personinfoByIINHandler(w http.ResponseWriter, r *http.Request) {
	iin := r.URL.Path[len("people/info/iin"):]
	if !isValidIIN(iin) {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Success: false,
			Errors:  "Invalid IIN",
		})
		return
	}
}

func personInfoByNameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/people/info/name/"):]
	rows, err := db.Query("SELECT name, iin, phone FROM people WHERE name LIKE '%' || ? || '%'", name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Success: false,
			Errors:  err.Error(),
		})
		return
	}
	defer rows.Close()

	var people []Person
	for rows.Next() {
		var person Person
		err := rows.Scan(&person.Name, &person.IIN, &person.Phone)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{
				Success: false,
				Errors:  err.Error(),
			})
			return
		}
		people = append(people, person)
	}

	json.NewEncoder(w).Encode(people)
}

func isValidIIN(iin string) bool {
	match, _ := regexp.MatchString(`^d{12}$`, iin)
	if !match {
		return false
	}
	return true
}
