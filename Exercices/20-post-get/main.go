package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Person struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Salary    int    `json:"salary"`
	Education string `json:"education"`
}

var people []Person

func loadPeopleFromFile(filename string) error {
	data, err := os.ReadFile(filename) // Use os.ReadFile instead of ioutil.ReadFile
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &people)
	if err != nil {
		return err
	}
	return nil
}

func savePeopleToFile(filename string) error {
	file, err := os.Create(filename) // Use os.Create to write to the file
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(people)
	if err != nil {
		return err
	}
	return nil
}

func getPeopleHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func postPersonHandler(w http.ResponseWriter, r *http.Request) {
	var newPerson Person
	err := json.NewDecoder(r.Body).Decode(&newPerson)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	people = append(people, newPerson)

	err = savePeopleToFile("people.json")
	if err != nil {
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}

	fmt.Printf("New person added: %+v\n", newPerson)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPerson)
}

func putPersonHandler(w http.ResponseWriter, r *http.Request) {
	var updatedPerson Person
	err := json.NewDecoder(r.Body).Decode(&updatedPerson)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	found := false
	for i, person := range people {
		if person.Name == updatedPerson.Name {
			people[i] = updatedPerson
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	err = savePeopleToFile("people.json")
	if err != nil {
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Person updated: %+v\n", updatedPerson)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPerson)
}

func deletePersonHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	index := -1
	for i, person := range people {
		if person.Name == name {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	people = append(people[:index], people[index+1:]...)

	err := savePeopleToFile("people.json")
	if err != nil {
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Person deleted: %s\n", name)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	err := loadPeopleFromFile("people.json")
	if err != nil {
		fmt.Printf("Error loading data: %s\n", err)
		return
	}

	http.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getPeopleHandler(w, r)
		case "POST":
			postPersonHandler(w, r)
		case "PUT":
			putPersonHandler(w, r)
		case "DELETE":
			deletePersonHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	port := ":8080"
	fmt.Printf("Server running on http://localhost%s\n", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}





/* package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Person struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Salary    int    `json:"salary"`
	Education string `json:"education"`
}

var people []Person

func loadPeopleFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &people)
	if err != nil {
		return err
	}
	return nil
}

func savePeopleToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(people)
	if err != nil {
		return err
	}
	return nil
}

func getPeopleHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		// Return all people if no id is provided
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(people)
		return
	}

	// If id is provided, find the person by ID
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, person := range people {
		if person.ID == idInt {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(person)
			return
		}
	}

	http.Error(w, "Person not found", http.StatusNotFound)
}

func postPersonHandler(w http.ResponseWriter, r *http.Request) {
	var newPerson Person
	err := json.NewDecoder(r.Body).Decode(&newPerson)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Generate a new ID for the person
	newID := 1
	if len(people) > 0 {
		newID = people[len(people)-1].ID + 1
	}
	newPerson.ID = newID

	people = append(people, newPerson)

	err = savePeopleToFile("people.json")
	if err != nil {
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}

	fmt.Printf("New person added: %+v\n", newPerson)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPerson)
}

func deletePersonHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	index := -1
	for i, person := range people {
		if person.ID == idInt {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	people = append(people[:index], people[index+1:]...)

	err = savePeopleToFile("people.json")
	if err != nil {
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Person with ID %d deleted\n", idInt)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	err := loadPeopleFromFile("people.json")
	if err != nil {
		fmt.Printf("Error loading data: %s\n", err)
		return
	}

	http.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getPeopleHandler(w, r)
		case "POST":
			postPersonHandler(w, r)
		case "DELETE":
			deletePersonHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	port := ":8080"
	fmt.Printf("Server running on http://localhost%s\n", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
*/