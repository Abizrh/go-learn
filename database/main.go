package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "abhi"
	password = "270402"
	dbname   = "go-learn"
)

var (
	db  *sql.DB
	err error
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " + " password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("database connected!!!")

	CreateEmployee()

}

type Employee struct {
	ID        int
	Full_name string
	Email     string
	Age       int
	Division  string
}


func CreateEmployee (){
	var employee = Employee{}

	sqlQuery := `
	INSERT INTO employees (full_name, email, age, division)
	VALUES ($1, $2, $3, $4)
	Returning *
	`

	err = db.QueryRow(sqlQuery, "alwi", "alwi@gmail.com", 22, "IT").
	Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Employee Data : %+v\n", employee)
}