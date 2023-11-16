package main

import (
	"fmt"
)

type Employee struct {
	Name       string
	Age        int
	Department string
}

func addEmployee(employees *[]Employee, emp Employee) {
	*employees = append(*employees, emp)
}

func listEmployees(employees []Employee) {
	fmt.Println("List of Employees:")
	for _, emp := range employees {
		fmt.Printf("Name: %s, Age: %d, Department: %s\n", emp.Name, emp.Age, emp.Department)
	}
}

func main() {
	var employees []Employee

	adam := Employee{Name: "Adam Elmaghraby", Age: 17, Department: "Marketing"}
	chris := Employee{Name: "Chris Elkins", Age: 41, Department: "Programming"}
	alan := Employee{Name: "Alan E", Age: 18, Department: "Programming"}

	addEmployee(&employees, adam)
	addEmployee(&employees, chris)
	addEmployee(&employees, alan)

	listEmployees(employees)

}
