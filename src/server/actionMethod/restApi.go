package actionMethod

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	logrus "github.com/sirupsen/logrus"
	"net/http"
	"server/model"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Welcome ...")
	json.NewEncoder(w).Encode(http.StatusOK)
}
func GetStudents(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Request get list Students")
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db1")
	if err != nil {
		logrus.Error("Fault to connect database")
		return
	}
	defer db.Close()
	res, err := db.Query("SELECT student_id,student_name,student_phone,student_birthday FROM students")
	if err != nil {
		logrus.Error("Statment query is mistake")
		return
	}
	var listStudents []model.Student
	for res.Next() {
		var (
			id       int
			name     string
			phone    string
			birthday string
		)
		err = res.Scan(&id, &name, &phone, &birthday)
		if err != nil {
			logrus.Error("Scan result query is mistake")
		}
		student := model.Student{Id: id, Name: name, Phone: phone, BirthDay: birthday}
		listStudents = append(listStudents, student)
	}
	logrus.Info("Responded list Students")
	err = json.NewEncoder(w).Encode(listStudents)
	if err != nil {
		logrus.Error("respond json list students is mistake")
	}
}
func GetStudent(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Request get a Student")
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db1")
	if err != nil {
		logrus.Error("Fault to connect database")
		return
	}
	defer db.Close()
	para := mux.Vars(r)
	res, err := db.Query("SELECT student_id,student_name,student_phone,student_birthday FROM students WHERE student_id=?", para["id"])
	if err != nil {
		logrus.Error("Statment get query is mistake")
		return
	}
	var student model.Student
	if res.Next() {
		var (
			id       int
			name     string
			phone    string
			birthday string
		)
		err = res.Scan(&id, &name, &phone, &birthday)
		if err != nil {
			logrus.Error("Scan result query is mistake")
		}
		student = model.Student{Id: id, Name: name, Phone: phone, BirthDay: birthday}
	} else {
		logrus.Info("No found from request")
		_ = json.NewEncoder(w).Encode("No Found student id")
		return
	}
	logrus.Info("Responded a Students")
	_ = json.NewEncoder(w).Encode(student)

}
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Request create a Student")
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db1")
	if err != nil {
		logrus.Error("Fault to connect database")
		return
	}
	defer db.Close()
	var newStudent model.Student
	err = json.NewDecoder(r.Body).Decode(&newStudent)
	if err != nil {
		logrus.Error("Change data to student obj fail")
		return
	}

	_, err = db.Query("INSERT INTO students (student_name,student_phone,student_birthday)VALUES (?,?,?)", newStudent.Name, newStudent.Phone, newStudent.BirthDay)
	if err != nil {
		logrus.Error("Statment query create is mistake")
		return
	}
	logrus.Info("Create a new Student cuccess")
	_ = json.NewEncoder(w).Encode(http.StatusOK)

}
func UpdateStudents(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Request update a Student")
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db1")
	if err != nil {
		logrus.Error("Fault to connect database")
		return
	}
	defer db.Close()
	para := mux.Vars(r)
	res, err := db.Query("SELECT student_name,student_phone,student_birthday FROM students WHERE student_id=?", para["id"])
	if err != nil {
		logrus.Error("Statment query update is mistake")
		return
	}

	if res.Next() {
		var (
			name     string
			phone    string
			birthday string
		)
		err = res.Scan(&name, &phone, &birthday)
		if err != nil {
			logrus.Error("Scan result query is mistake")
			_ = json.NewEncoder(w).Encode("Can't update student")
			return
		}
		_, err := db.Query("UPDATE students SET student_name=?, student_phone=?, student_birthday=? WHERE student_id=?", name, phone, birthday, para["id"])
		if err != nil {
			logrus.Error("Statment query update is mistake")
			return
		}
	} else {
		logrus.Info("No found student id from request")
		_ = json.NewEncoder(w).Encode("No Found student id")
		return
	}
	logrus.Info("Responded update Student cuccess ")
	_ = json.NewEncoder(w).Encode(http.StatusCreated)
}
func DeleteStudents(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Request delete a Student")
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db1")
	if err != nil {
		logrus.Error("Fault to connect database")
		return
	}
	defer db.Close()
	para := mux.Vars(r)
	_, err = db.Query("DELETE FROM students WHERE student_id=?", para["id"])
	if err != nil {
		logrus.Error("Statment query delete is mistake")
		return
	}

	logrus.Info("Responded delete student cuccess ")
	_ = json.NewEncoder(w).Encode(http.StatusAccepted)

}
