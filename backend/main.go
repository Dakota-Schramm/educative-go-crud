package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	Routers()
}


func Routers() {
	InitDB()
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/users",
		GetUsers).Methods("GET")
	router.HandleFunc("/users",
		CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}",
		GetUser).Methods("GET")
	router.HandleFunc("/users/{id}",
		UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}",
		DeleteUser).Methods("DELETE")
	http.ListenAndServe(":3000",
		&CORSRouterDecorator{router})
}


func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

    var users []User
	rows, err := db.Query("SELECT * FROM users;")
	if err != nil {
		http.Error(w, err.error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.FirstName, &user.MiddleName, &user.LastName, &user.Email, &user.CivilStatus, &user.Birthday, &user.Contact, &user.Address, &user.Age)
		if err != nil {
			http.Error(w, err.error(), http.StatusInternalServerError)
			return
		}
		queriedUsers = append(queriedUsers, user)
	}

    json.NewEncoder(w).Encode(queriedUsers)
}
// Task 4: write your code for create user here
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stmt, err := db.Prepare("INSERT INTO " + 
		"users(id, firstName, lastName, email, gender, civilStatus, birthday, contact, address, age) " + 
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, err.error(), http.StatusInternalServerError)
		return
	}

	if pl, err := ioutil.ReadAll(r.Body); err != nil {
		http.Error(w, err.error(), http.StatusInternalServerError)
		return
	}

	var keyVal[string]string
	err = json.Unmarshall(pl, &keyVal)

	firstName := keyVal["firstName"]
    middleName := keyVal["middleName"]
    lastName := keyVal["lastName"]
    email := keyVal["email"]
    gender := keyVal["gender"]
    civilStatus := keyVal["civilStatus"]
    birthday := keyVal["birthday"]
    contact := keyVal["contact"]
    address := keyVal["address"]

    _, err = stmt.Exec(firstName, middleName, lastName, email, gender, civilStatus, birthday, contact, address)
	if err != nil {
		http.Error(w, err.error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "New user was created")
}

// Task 5: Write code for get user here
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	rows, err := db.Query("SELECT * FROM users WHERE id = ?", params["ID"])
	if err != nil {
		http.Error(w, err.error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var user User
	userFound := false
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.FirstName, &user.MiddleName, &user.LastName, &user.Email, &user.CivilStatus, &user.Birthday, &user.Contact, &user.Address, &user.Age)
		if err != nil {
			http.Error(w, err.error(), http.StatusInternalServerError)
			return
		}
		userFound = true
	}

	if !userFound {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "User not found with ID: %s", params["id"])
        return
    }

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// Task 6: write code for update user here
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	stmt, err := db.Prepare("UPDATE users " + 
		"SET id = ?, SET firstName = ?, SET lastName = ?, SET email = ?, " +
		"SET gender = ?, SET civilStatus = ?, SET birthday = ?, " +
		"SET contact = ?, SET address = ?, SET age = ?) " + 
		"WHERE id = ?")
	if err != nil {
		http.Error(w, err.error(), http.StatusInternalServerError)
		return
	}

	body := r.Body
	var keyVal[string]string
	keyVal = json.Unmarshall(body, &keyVal)

	firstName := keyVal["firstName"]
    middleName := keyVal["middleName"]
    lastName := keyVal["lastName"]
    email := keyVal["email"]
    gender := keyVal["gender"]
    civilStatus := keyVal["civilStatus"]
    birthday := keyVal["birthday"]
    contact := keyVal["contact"]
    address := keyVal["address"]

    res, err = stmt.Exec(firstName, middleName, lastName, email, gender, civilStatus, birthday, contact, address, params["id"])
	if err != nil {
		http.Error(w, err.error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
    if err != nil {
		http.Error(w, err.error(), http.StatusInternalServerError)
		return
    }

    if rowsAffected == 0 {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "No user found with ID = %s", params["id"])
    } else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "User %d was updated successfully!", params["id"])
	}
}

// Task 7: Write code for delete user here
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?", params["id"])
	if err != nil {
		http.Error(w, err.error(), http.StatusInternalServerError)
		return
	}

	res, err = stmt.Exec(params["id"])
	if err != nil {
		http.Error(w, err.error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
    if err != nil {
		http.Error(w, err.error(), http.StatusInternalServerError)
		return
    }

    if rowsAffected == 0 {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "No user found with ID = %s", params["id"])
    } else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "User %d was deleted successfully!", params["id"])
	}

}

type User struct {
    ID          string `json:"id"`
    FirstName   string `json:"firstName"`
    MiddleName  string `json:"middleName"`
    LastName    string `json:"lastName"`
    Email       string `json:"email"`
    Gender      string `json:"gender"`
    CivilStatus string `json:"civilStatus"`
    Birthday    string `json:"birthday"`
    Contact     string `json:"contact"`
    Address     string `json:"address"`
    Age         string `json:"age"`
}

var db *sql.DB
var err error

// Task 1: Write code for DB initialization here
func InitDB() {
	const driver_name = "mysql"
	connection_string := databaseConnection()
	connection, err := sql.Open(driver_name, connection_string)

	if err != nil {
		panic(err)
	}
}

func databaseConnection() string {
	var user string
	var password string
	var host string
	var port string
	var database string

	return fmt.Sprintf("%s:%s-@tcp(%s:%s)/%s", user, password, host, port, database)

}

type CORSRouterDecorator struct {
	R *mux.Router
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter,
	req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Accept-Language,"+
				" Content-Type, YourOwnHeader")
	}

	if req.Method == "OPTIONS" {
		return
	}

	c.R.ServeHTTP(rw, req)
}
