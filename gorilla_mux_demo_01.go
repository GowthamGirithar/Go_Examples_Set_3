package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	mux "github.com/gorilla/mux"
	"log"
	"net/http"
)

type BookCatalog struct {
	//exported field since it begins with a capital letter
	Name   string
	Author string
}


var bookCatalogs =make([]BookCatalog,0)
var db *sql.DB

func main() {

	router:=mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books", bookOperation)
	db,err:=getDbConnection()
	if err != nil{
		fmt.Println("error in getting db" , err)
	}
	db.Exec(`
	create table if not exists user5_books (
		name varchar(200),
		author  varchar(200)
	);
`)
	fmt.Println("the tables created successfully")
	log.Fatal(http.ListenAndServe(":8000", router))


	
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT name , author FROM user5_books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next(){
		bk := new(BookCatalog)
		err := rows.Scan(&bk.Name, &bk.Author)
		if err != nil {
			log.Fatal("error while retrieving rows")
		}
		bookCatalogs = append(bookCatalogs, *bk)
	}
	data, err:=json.Marshal(bookCatalogs)
	if err != nil{
		log.Fatal(err)
	}
	w.Write(data)

}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bookCatalog BookCatalog
	_ = json.NewDecoder(r.Body).Decode(&bookCatalog)
	stmt, err:=db.Prepare("INSERT INTO user5_books VALUES (?,?)")
	if err != nil {
		log.Fatal("err in prepared statement" ,err)
	}
	res, err:=stmt.Exec(bookCatalog.Name,bookCatalog.Author)
	if err != nil{
		log.Fatal("error while executing" , err)
	}
	fmt.Println(res)
	w.Write([]byte("Created successfully!\n"))
}

func bookOperation(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func getDbConnection() (*sql.DB,error){
	var err error
	db, err = sql.Open("mysql", "user5:user5@tcp(34.93.106.67:3306)/booksdb")
	if err != nil{
		fmt.Println("error in connecting to db" , err)
		return nil, err
	}
	return db, nil
}
