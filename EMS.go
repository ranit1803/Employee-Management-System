package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Employee struct {
	Name        string
	ID          int
	Dept        string
	Salary      int
	Age         int
	Address     string
	Contact     string
	Email       string
	Designation string
	Gender      string
	Status      string
}

var employees = make(map[int]Employee)
var nextID int = 1

// clearInputBuffer clears any leftover input in the buffer
func clearInputBuffer() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n') // Discard input buffer
}

func addEmployee(name, dept string, salary, age int, address, contact, email, designation, gender, status string) {
	employees[nextID] = Employee{
		Name:        name,
		ID:          nextID,
		Dept:        dept,
		Salary:      salary,
		Age:         age,
		Address:     address,
		Contact:     contact,
		Email:       email,
		Designation: designation,
		Gender:      gender,
		Status:      status,
	}
	nextID++
	fmt.Println("Employee Added Successfully!")
}

func updateEmployee(id int, name, dept string, salary, age int, address, contact, email, designation, gender, status string) {
	if emp, exists := employees[id]; exists {
		emp.Name = name
		emp.Dept = dept
		emp.Salary = salary
		emp.Age = age
		emp.Address = address
		emp.Contact = contact
		emp.Email = email
		emp.Designation = designation
		emp.Gender = gender
		emp.Status = status
		employees[id] = emp
		fmt.Println("Employee Updated Successfully!")
	} else {
		fmt.Println("Employee Not Found!")
	}
}

func removeEmployee(id int) {
	if _, exists := employees[id]; exists {
		delete(employees, id)
		fmt.Printf("Employee with ID %d Removed Successfully!\n", id)
	} else {
		fmt.Println("Employee Not Found!")
	}
}

func displayEmployeeDetails() {
	if len(employees) == 0 {
		fmt.Println("No Employees Found!")
		return
	}
	for _, emp := range employees {
		fmt.Printf("Employee ID: %d\n", emp.ID)
		fmt.Printf("Employee Name: %s\n", emp.Name)
		fmt.Printf("Employee Department: %s\n", emp.Dept)
		fmt.Printf("Employee Salary: %d\n", emp.Salary)
		fmt.Printf("Employee Age: %d\n", emp.Age)
		fmt.Printf("Employee Address: %s\n", emp.Address)
		fmt.Printf("Employee Contact: %s\n", emp.Contact)
		fmt.Printf("Employee Email: %s\n", emp.Email)
		fmt.Printf("Employee Designation: %s\n", emp.Designation)
		fmt.Printf("Employee Gender: %s\n", emp.Gender)
		fmt.Printf("Employee Status: %s\n", emp.Status)
		fmt.Println("-------------------------------------------------")
	}
}

func main() {
	fmt.Println("Welcome to Employee Management System")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n1. Add Employee")
		fmt.Println("2. Update Employee")
		fmt.Println("3. Remove Employee")
		fmt.Println("4. Display Employee Details")
		fmt.Println("5. Exit")
		fmt.Print("Enter Your Choice: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")
			clearInputBuffer()
			continue
		}

		switch choice {
		case 1:
			var name, dept, address, contact, email, designation, gender, status string
			var salary, age int

			fmt.Print("Enter Employee Name: ")
			name, _ = reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter Employee Department: ")
			dept, _ = reader.ReadString('\n')
			dept = strings.TrimSpace(dept)

			fmt.Print("Enter Employee Salary: ")
			_, err := fmt.Scan(&salary)
			if err != nil {
				fmt.Println("Invalid input for salary! Please enter a number.")
				clearInputBuffer()
				continue
			}
			clearInputBuffer() // Clear the input buffer after reading salary

			fmt.Print("Enter Employee Age: ")
			_, err = fmt.Scan(&age)
			if err != nil {
				fmt.Println("Invalid input for age! Please enter a number.")
				clearInputBuffer()
				continue
			}
			clearInputBuffer() // Clear the input buffer after reading age

			fmt.Print("Enter Employee Address: ")
			address, _ = reader.ReadString('\n')
			address = strings.TrimSpace(address)

			fmt.Print("Enter Employee Contact: ")
			contact, _ = reader.ReadString('\n')
			contact = strings.TrimSpace(contact)

			fmt.Print("Enter Employee Email: ")
			email, _ = reader.ReadString('\n')
			email = strings.TrimSpace(email)

			fmt.Print("Enter Employee Designation: ")
			designation, _ = reader.ReadString('\n')
			designation = strings.TrimSpace(designation)

			fmt.Print("Enter Employee Gender: ")
			gender, _ = reader.ReadString('\n')
			gender = strings.TrimSpace(gender)

			fmt.Print("Enter Employee Status: ")
			status, _ = reader.ReadString('\n')
			status = strings.TrimSpace(status)

			addEmployee(name, dept, salary, age, address, contact, email, designation, gender, status)

		case 2:
			var id int
			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&id)
			clearInputBuffer() // Clear the input buffer after reading ID

			if _, exists := employees[id]; !exists {
				fmt.Println("Employee Not Found!")
				continue
			}

			fmt.Println("Updating Employee ID:", id)
			var name, dept, address, contact, email, designation, gender, status string
			var salary, age int
			fmt.Print("Enter Employee Name: ")
			name, _ = reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter Employee Department: ")
			dept, _ = reader.ReadString('\n')
			dept = strings.TrimSpace(dept)

			fmt.Print("Enter Employee Salary: ")
			fmt.Scan(&salary)
			clearInputBuffer() // Clear the input buffer after reading salary

			fmt.Print("Enter Employee Age: ")
			fmt.Scan(&age)
			clearInputBuffer() // Clear the input buffer after reading age

			fmt.Print("Enter Employee Address: ")
			address, _ = reader.ReadString('\n')
			address = strings.TrimSpace(address)

			fmt.Print("Enter Employee Contact: ")
			contact, _ = reader.ReadString('\n')
			contact = strings.TrimSpace(contact)

			fmt.Print("Enter Employee Email: ")
			email, _ = reader.ReadString('\n')
			email = strings.TrimSpace(email)

			fmt.Print("Enter Employee Designation: ")
			designation, _ = reader.ReadString('\n')
			designation = strings.TrimSpace(designation)

			fmt.Print("Enter Employee Gender: ")
			gender, _ = reader.ReadString('\n')
			gender = strings.TrimSpace(gender)

			fmt.Print("Enter Employee Status: ")
			status, _ = reader.ReadString('\n')
			status = strings.TrimSpace(status)

			updateEmployee(id, name, dept, salary, age, address, contact, email, designation, gender, status)

		case 3:
			var id int
			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&id)
			clearInputBuffer() // Clear the input buffer after reading ID
			removeEmployee(id)

		case 4:
			displayEmployeeDetails()

		case 5:
			fmt.Println("Thank You for Using Employee Management System")
			return

		default:
			fmt.Println("Invalid Choice!")
			clearInputBuffer()
		}
	}
}