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
	Zip    string `json:"zip"`
}

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

type UsersWrapper struct {
	Users []User `json:"users"`
}

func handleUserGet(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	users, err := readUsersFromFile()
	if err != nil {
		http.Error(w, "Unable to read users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
func handleUserPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := readUsersFromFile()
	if err != nil {
		http.Error(w, "Unable to read users", http.StatusInternalServerError)
		return
	}

	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	newUser.ID = getNextID(users)
	users = append(users, newUser)

	if err := writeUsersToFile(users); err != nil {
		http.Error(w, "Unable to save new user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
func readUsersFromFile() ([]User, error) {
	data, err := os.ReadFile("users.json") // Use os.ReadFile
	if err != nil {
		return nil, err
	}

	var wrapper UsersWrapper
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, err
	}
	return wrapper.Users, nil
}
func writeUsersToFile(users []User) error {
	wrapper := UsersWrapper{Users: users}
	data, err := json.MarshalIndent(wrapper, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("users.json", data, 0644) // Use os.WriteFile
}
func getNextID(users []User) int {
	maxID := 0
	for _, user := range users {
		if user.ID > maxID {
			maxID = user.ID
		}
	}
	return maxID + 1
}
func handleGetUserByIDAndName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	users, err := readUsersFromFile()
	if err != nil {
		http.Error(w, "Unable to read users", http.StatusInternalServerError)
		return
	}

	// Parse path parameter
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Parse query parameter
	name := r.URL.Query().Get("name")

	// Find user by ID
	var userFound *User
	for _, user := range users {
		if user.ID == id {
			userFound = &user
			break
		}
	}

	// Handle user not found
	if userFound == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Filter by name if the query parameter is provided
	if name != "" && !strings.EqualFold(userFound.Name, name) {
		http.Error(w, "User with the given name not found", http.StatusNotFound)
		return
	}

	// Return the user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userFound)
}
func handleGetUserByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	users, err := readUsersFromFile()
	if err != nil {
		http.Error(w, "Unable to read users", http.StatusInternalServerError)
		return
	}

	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	//strconv.ParseFloat(str, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}
func handleUpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	users, err := readUsersFromFile()
	if err != nil {
		http.Error(w, "Unable to read users", http.StatusInternalServerError)
		return
	}

	// Get the ID from the path parameter
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Decode the updated user details from the request body
	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Update the user with the given ID
	updated := false
	for i, user := range users {
		if user.ID == id {
			// Update user details
			users[i].Name = updatedUser.Name
			users[i].Email = updatedUser.Email
			users[i].Address = updatedUser.Address
			updated = true
			break
		}
	}

	if !updated {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Save updated user list to the file
	if err := writeUsersToFile(users); err != nil {
		http.Error(w, "Unable to save updated user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	users, err := readUsersFromFile()
	if err != nil {
		http.Error(w, "Unable to read users", http.StatusInternalServerError)
		return
	}

	// Get the ID from the path parameter
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Find and delete the user with the given ID
	deleted := false
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...) // Remove user
			deleted = true
			break
		}
	}

	if !deleted {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Save updated user list to the file
	if err := writeUsersToFile(users); err != nil {
		http.Error(w, "Unable to save updated user list", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) 
}
func handleUserByIdAndName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse path parameters
	idStr := ps.ByName("id")
	name := ps.ByName("name")

	// Convert id to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Read users from the file
	users, err := readUsersFromFile()
	if err != nil {
		http.Error(w, "Unable to read users", http.StatusInternalServerError)
		return
	}

	// Search for all users with the matching ID and Name
	var matchingUsers []User
	for _, user := range users {
		if user.ID == id && strings.EqualFold(user.Name, name) {
			matchingUsers = append(matchingUsers, user)
		}
	}

	// If no users are found, return 404
	if len(matchingUsers) == 0 {
		http.Error(w, "No users found with the provided ID and name", http.StatusNotFound)
		return
	}

	// Return all matching users as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matchingUsers)
}
func handleGetUsersByMultipleIds(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Parse the query parameter for IDs
	idsParam := r.URL.Query().Get("ids")
	if idsParam == "" {
		http.Error(w, "IDs query parameter is required", http.StatusBadRequest)
		return
	}

	// Split the IDs by comma and convert to integers
	idStrings := strings.Split(idsParam, ",")
	var ids []int
	for _, idStr := range idStrings {
		id, err := strconv.Atoi(strings.TrimSpace(idStr))
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}
		ids = append(ids, id)
	}

	// Read users from the file
	users, err := readUsersFromFile()
	if err != nil {
		http.Error(w, "Unable to read users", http.StatusInternalServerError)
		return
	}

	// Find users with matching IDs
	var matchingUsers []User
	for _, user := range users {
		for _, id := range ids {
			if user.ID == id {
				matchingUsers = append(matchingUsers, user)
				break
			}
		}
	}

	// If no users are found, return 404
	if len(matchingUsers) == 0 {
		http.Error(w, "No users found with the provided IDs", http.StatusNotFound)
		return
	}

	// Return the matching users as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matchingUsers)
}
func handleGetUsersSelective(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Read users from the file
	users, err := readUsersFromFile()
	if err != nil {
		http.Error(w, "Unable to read users", http.StatusInternalServerError)
		return
	}

	// Parse the fields query parameter
	fieldsParam := r.URL.Query().Get("fields")
	fields := strings.Split(fieldsParam, ",")

	// Build a response dynamically based on requested fields
	var selectiveUsers []map[string]interface{}
	for _, user := range users {
		userMap := make(map[string]interface{})
		for _, field := range fields {
			switch field {
			case "id":
				userMap["id"] = user.ID
			case "name":
				userMap["name"] = user.Name
			case "email":
				userMap["email"] = user.Email
			case "address":
				userMap["address"] = user.Address
			}
		}
		selectiveUsers = append(selectiveUsers, userMap)
	}

	// Return the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(selectiveUsers)
}
func handleGetAverageID(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Read users from the file
	users, err := readUsersFromFile()
	if err != nil {
		http.Error(w, "Unable to read users", http.StatusInternalServerError)
		return
	}

	// Compute the average of the IDs
	var total int
	for _, user := range users {
		total += user.ID
	}
	average := float64(total) / float64(len(users))

	// Return the average as a JSON response
	response := map[string]float64{
		"average_id": average,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


func main() {
	router := httprouter.New()
	router.PUT("/users/:id", handleUpdateUser)   
	router.DELETE("/users/:id", handleDeleteUser)
	router.GET("/users", handleUserGet)
	router.POST("/users", handleUserPost)
	router.GET("/users/:id", handleGetUserByID)
	router.GET("/users/:id/:name", handleUserByIdAndName)
	router.GET("/users", handleGetUsersByMultipleIds)
//GET /users?ids=1,2

	router.GET("/users/:id", handleGetUserByIDAndName)
	http.ListenAndServe(":8080", router)
}
