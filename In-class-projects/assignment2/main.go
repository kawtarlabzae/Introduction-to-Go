package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Salary    int    `json:"salary"`
	Education string `json:"education"`
}

type InputData struct {
	People []Person `json:"people"`
}

type OutputData struct {
	AverageAge               float64            `json:"average_age"`
	YoungestPersons          []string           `json:"youngest_persons"`
	OldestPersons            []string           `json:"oldest_persons"`
	AverageSalary            float64            `json:"average_salary"`
	HighestSalaryPersons     []string           `json:"highest_salary_persons"`
	LowestSalaryPersons      []string           `json:"lowest_salary_persons"`
	CountsByEducationLevel   map[string]int     `json:"counts_by_education_level"`
}

func main() {
	// Read input JSON using os.ReadFile
	inputFile, err := os.ReadFile("file-person.json")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	var inputData InputData
	err = json.Unmarshal(inputFile, &inputData)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	// Variables for calculations
	var totalAge, totalSalary int
	minAge, maxAge := inputData.People[0].Age, inputData.People[0].Age
	minSalary, maxSalary := inputData.People[0].Salary, inputData.People[0].Salary
	youngestPersons := []string{}
	oldestPersons := []string{}
	highestSalaryPersons := []string{}
	lowestSalaryPersons := []string{}
	countsByEducationLevel := make(map[string]int)

	// Process the data
	for _, person := range inputData.People {
		totalAge += person.Age
		totalSalary += person.Salary

		// Track youngest and oldest
		if person.Age < minAge {
			minAge = person.Age
			youngestPersons = []string{person.Name}
		} else if person.Age == minAge {
			youngestPersons = append(youngestPersons, person.Name)
		}

		if person.Age > maxAge {
			maxAge = person.Age
			oldestPersons = []string{person.Name}
		} else if person.Age == maxAge {
			oldestPersons = append(oldestPersons, person.Name)
		}

		// Track highest and lowest salary
		if person.Salary > maxSalary {
			maxSalary = person.Salary
			highestSalaryPersons = []string{person.Name}
		} else if person.Salary == maxSalary {
			highestSalaryPersons = append(highestSalaryPersons, person.Name)
		}

		if person.Salary < minSalary {
			minSalary = person.Salary
			lowestSalaryPersons = []string{person.Name}
		} else if person.Salary == minSalary {
			lowestSalaryPersons = append(lowestSalaryPersons, person.Name)
		}

		// Count education levels
		countsByEducationLevel[person.Education]++
	}

	// Calculate averages
	averageAge := float64(totalAge) / float64(len(inputData.People))
	averageSalary := float64(totalSalary) / float64(len(inputData.People))

	// Prepare output
	outputData := OutputData{
		AverageAge:               averageAge,
		YoungestPersons:          youngestPersons,
		OldestPersons:            oldestPersons,
		AverageSalary:            averageSalary,
		HighestSalaryPersons:     highestSalaryPersons,
		LowestSalaryPersons:      lowestSalaryPersons,
		CountsByEducationLevel:   countsByEducationLevel,
	}

	// Write output JSON
	outputFile, err := os.Create("output.json")
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer outputFile.Close()

	encoder := json.NewEncoder(outputFile)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(outputData)
	if err != nil {
		fmt.Printf("Error writing JSON: %v\n", err)
		return
	}

	fmt.Println("Output written to output.json")
}
