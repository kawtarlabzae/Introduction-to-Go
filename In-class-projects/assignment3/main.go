package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State    string `json:"state"`
	PostalCode string `json:"postal-code"`
}
type Course struct {
	Code string `json:"code"`
	Name   string `json:"name"`
	Credit    int `json:"credit"`
}

type Student struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Age   int  `json:"age"`
	Major    string  `json:"major"`
	Address Address `json:"address"`
	Course []Course `json:"courses"`
}

type StudentsWrapper struct {
	Students []Student `json:"users"`
}


func handleStudentGet(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	students, err := readStudentsFromFile()
	if err != nil {
		http.Error(w, "Unable to read students", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}
func handleStudentPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	students, err := readStudentsFromFile()
	if err != nil {
		http.Error(w, "Unable to read students", http.StatusInternalServerError)
		return
	}
	var newStudent Student
	if err := json.NewDecoder(r.Body).Decode(&newStudent); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	newStudent.ID = getNextID(students)
	students = append(students, newStudent)
	if err := writeStudentsToFile(students); err != nil {
		http.Error(w, "Unable to save new student", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newStudent)
}
func readStudentsFromFile() ([]Student, error) {
	data, err := os.ReadFile("users.json") 
	if err != nil {
		return nil, err
	}
	var wrapper StudentsWrapper
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, err
	}
	return wrapper.Students, nil
}
func writeStudentsToFile(students []Student) error {
	wrapper := StudentsWrapper{Students: students}
	data, err := json.MarshalIndent(wrapper, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("users.json", data, 0644) 
}
func getNextID(students []Student) int {
	maxID := 0
	for _, student := range students {
		if student.ID > maxID {
			maxID = student.ID
		}
	}
	return maxID + 1
}
func handleGetStudentByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	students, err := readStudentsFromFile()
	if err != nil {
		http.Error(w, "Unable to read students", http.StatusInternalServerError)
		return
	}
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}
	for _, student := range students {
		if student.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(student)
			return
		}
	}
	http.Error(w, "Student not found", http.StatusNotFound)
}
func handleUpdateStudent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	students, err := readStudentsFromFile()
	if err != nil {
		http.Error(w, "Unable to read students", http.StatusInternalServerError)
		return
	}
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}
	var updatedStudent Student
	if err := json.NewDecoder(r.Body).Decode(&updatedStudent); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	updated := false
	for i, student := range students {
		if student.ID == id {
			students[i].Name = updatedStudent.Name
			students[i].Age = updatedStudent.Age
			students[i].Major = updatedStudent.Major
			students[i].Address = updatedStudent.Address
			students[i].Course = updatedStudent.Course
			updated = true
			break
		}
	}
	if !updated {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	if err := writeStudentsToFile(students); err != nil {
		http.Error(w, "Unable to save updated student", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedStudent)
}

func handleDeleteStudent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	students, err := readStudentsFromFile()
	if err != nil {
		http.Error(w, "Unable to read students", http.StatusInternalServerError)
		return
	}
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}
	deleted := false
	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)
			deleted = true
			break
		}
	}
	if !deleted {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	if err := writeStudentsToFile(students); err != nil {
		http.Error(w, "Unable to save updated student list", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent) 
}


func handleGetStudentByIDAndName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	students, err := readStudentsFromFile()
	if err != nil {
		http.Error(w, "Unable to read students", http.StatusInternalServerError)
		return
	}
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}
	name := r.URL.Query().Get("name")
	var studentFound *Student
	for _, student := range students {
		if student.ID == id {
			studentFound = &student
			break
		}
	}
	if studentFound == nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	if name != "" && !strings.EqualFold(studentFound.Name, name) {
		http.Error(w, "Student with the given name not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studentFound)
}




func main() {
	router := httprouter.New()
	router.PUT("/students/:id", handleUpdateStudent)   
	router.DELETE("/students/:id", handleDeleteStudent)
	router.GET("/students", handleStudentGet)
	router.POST("/students", handleStudentPost)
	router.GET("/students/:id", handleGetStudentByID)
	//Here you can try : GET /students/id?name=name
	router.GET("/students/:id", handleGetStudentByIDAndName)
	http.ListenAndServe(":8080", router)
}
