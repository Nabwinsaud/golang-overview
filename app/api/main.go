package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"reflect"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "nabin"
	dbname   = "godb"
)

var db *sql.DB

func retrieveData(w http.ResponseWriter, r *http.Request) {
	data := []string{"nabin", "hari", "sita"}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	// for _, record := range data {
	// 	fmt.Fprint(w, record)

	// }\

	response := map[string]interface{}{
		"success": true,
		"message": "data retrieved successfully",
		"data":    data,
	}

	jsonRespnse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("error marshalling json", err)
	}

	w.Write(jsonRespnse)
}

type Student struct {
	Name  string `json:name`
	Age   int    `json:age`
	Email string `json:email`
	Id    int    `json:id`
}

func main() {

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	CheckError(err)
	// close database after main function ends
	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Successfully connected to database!")

	// insert data inside the student table
	// hardcode data
	// insertData := `INSERT INTO students (name,age,email) VALUES('Nabin',21,'nabin@gmail.com')`

	// _, e := db.Exec(insertData)
	// CheckError((e))

	// dynamic and secure way to insert data
	// insertDataSecureWay := `INSERT INTO students(name,age,email) VALUES($1,$2,$3)`

	// _, e := db.Exec(insertDataSecureWay, "Shayam", 21, "shyam@gmail.com")

	// CheckError(e)

	// query the result from the database

	rows, e := db.Query(`SELECT * FROM students`)

	CheckError(e)

	defer rows.Close()
	fmt.Printf("%v\n", (rows))

	// iterate over the rows in while loop
	for rows.Next() {

		var id int
		var name string
		var age int
		var email string
		err = rows.Scan(&name, &age, &email, &id) //* if the order of the columns is changed then the order of the scan should also be changed other wise it will throw an error
		CheckError(err)

		fmt.Printf("id is %d name is %s and age is %d and email is %s\n", id, name, age, email)
	}

	readTemplateEngine()
	http.HandleFunc("/api", retrieveData)
	http.HandleFunc("/api/create", CreateUser)
	http.HandleFunc("/api/update", updateUser)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	parsedData, err := json.Marshal(body)
	if err != nil {
		panic(err)

	}

	response := map[string]interface{}{
		"success": true,
		"data":    parsedData,
	}

	// jsonResponse, er := json.Marshaler(response)
	// if er != nil {
	// 	panic(er)
	// }

	if r.Method == "POST" {
		fmt.Println("data in body is", string(body))

	} else {
		panic(`Route not Found`)
	}

	fmt.Printf("parseed data in body and type is", string(body), reflect.TypeOf(body), response)

	// w.Write(jsonResponse)

}

// !issues in the updateUser while updating the data
func updateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)

	}
	defer r.Body.Close()
	fmt.Println("body is", string(body))

	var formattedData Student
	err = json.Unmarshal(body, &formattedData)
	if err != nil {
		http.Error(w, "failed to parse the data", http.StatusBadRequest)
	}

	fmt.Println("formatted data is ", formattedData, reflect.TypeOf(formattedData))
	if r.Method == "PUT" {
		// updateData, err := db.Exec(`UPDATE students SET name=$1,age=$2 WHERE id=$3`, "Nabin", 21, 1)
		updateData, err := db.Exec(`UPDATE students SET name=$1,age=$2 WHERE id=$3`, formattedData.Name, formattedData.Age, formattedData.Id)

		if err != nil {
			fmt.Println("Failed to execute the query", err)
			http.Error(w, "Failed to execute the query", http.StatusBadRequest)
		}
		fmt.Println("data updated successfully", updateData)
		w.WriteHeader(http.StatusOK)
	}

}

type User struct {
	Name  string
	Age   int
	Email string
}

func readTemplateEngine() {
	templ, err := template.ParseFiles("user.html")

	if err != nil {
		panic(err)
	}
	data := []User{{Name: "Nabin", Age: 21, Email: "nabin@gmail.com"}, {Name: "Hari", Age: 22, Email: "har@gmail.com"}}
	var tmplBuffer bytes.Buffer
	templErr := templ.Execute(&tmplBuffer, data)

	if templErr != nil {
		panic(templErr)
	}
	fmt.Println(tmplBuffer.String(), "the template is parsed successfully")

}
