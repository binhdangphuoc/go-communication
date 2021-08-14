package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	action "server/actionMethod"
	//config "server/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
	//// if there is an error opening the connection, handle it
	//if err != nil {
	//	panic(err.Error())
	//}
	//defer db.Close()

	//fmt.Println(db)
	router := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.Methods(http.MethodGet).Path("/*").HandlerFunc(action.Welcome)
	router.HandleFunc("/api/students", action.GetStudents).Methods("GET")
	router.HandleFunc("/api/students/{id}", action.GetStudent).Methods("GET")
	router.HandleFunc("/api/students", action.CreateStudent).Methods("POST")
	router.HandleFunc("/api/students/{id}", action.UpdateStudents).Methods("PUT")
	router.HandleFunc("/api/students/{id}", action.DeleteStudents).Methods("DELETE")

	err := http.ListenAndServe(":8888", handlers.CORS(headers, methods, origins)(router))

	if err != nil {
		panic(err)
		os.Exit(1)
	}

}
