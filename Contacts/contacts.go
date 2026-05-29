package main

import "fmt"

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}

func (person Person) FullName() string {
	return fmt.Sprintf("%s %s", person.FirstName, person.LastName)
}

func (person Person) Presentation() string {
	return fmt.Sprintf("I am %s, %d years old.", person.FullName(), person.Age)
}

type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
}

func (address Address) Format() string {
	return fmt.Sprintf("%s, %s %s", address.Street, address.PostalCode, address.City)
}

type Employee struct {
	Person
	Address
	Job    string  `json:"job"`
	Salary float64 `json:"salary"`
}

func (employee Employee) EmployeeRecord() string {
	return fmt.Sprintf(
		"--- Employee ---\n%s\nEmail: %s\nJob: %s\nSalary: $%.2f\nAddress: %s",
		employee.Person.Presentation(),
		employee.Email,
		employee.Job,
		employee.Salary,
		employee.Address.Format(),
	)
}

func (employee *Employee) RaiseSalary(percentage float64) {
	employee.Salary *= 1 + percentage/100
}

type Student struct {
	Person
	Promo   string  `json:"promo"`
	Average float64 `json:"average"`
}

func (student Student) Mention() string {
	switch {
	case student.Average >= 16:
		return "Very good"
	case student.Average >= 14:
		return "Good"
	case student.Average >= 12:
		return "Fine"
	case student.Average >= 10:
		return "Pass"
	default:
		return "None"
	}
}

func (student Student) StudentRecord() string {
	return fmt.Sprintf(
		"--- Student ---\n%s\nEmail: %s\nPromo: %s\nAverage: %.2f\nMention: %s",
		student.Person.Presentation(),
		student.Email,
		student.Promo,
		student.Average,
		student.Mention(),
	)
}

func main() {
	employees := map[string]Employee{
		"alice.martin@company.com": {
			Person: Person{
				FirstName: "Alice",
				LastName:  "Martin",
				Age:       32,
				Email:     "alice.martin@company.com",
			},
			Address: Address{
				Street:     "12 Peace Street",
				City:       "Paris",
				PostalCode: "75002",
			},
			Job:    "Developer",
			Salary: 42000,
		},
		"bob.smith@company.com": {
			Person: Person{
				FirstName: "Bob",
				LastName:  "Smith",
				Age:       45,
				Email:     "bob.smith@company.com",
			},
			Address: Address{
				Street:     "5 Victor Hugo Avenue",
				City:       "Lyon",
				PostalCode: "69003",
			},
			Job:    "Project Manager",
			Salary: 55000,
		},
	}

	students := map[string]Student{
		"claire.brown@school.edu": {
			Person: Person{
				FirstName: "Claire",
				LastName:  "Brown",
				Age:       20,
				Email:     "claire.brown@school.edu",
			},
			Promo:   "M1 Dev 2026",
			Average: 16.5,
		},
		"david.lee@school.edu": {
			Person: Person{
				FirstName: "David",
				LastName:  "Lee",
				Age:       19,
				Email:     "david.lee@school.edu",
			},
			Promo:   "M1 Dev 2026",
			Average: 11.2,
		},
	}

	for email, emp := range employees {
		fmt.Println(emp.EmployeeRecord())
		emp.RaiseSalary(5)
		employees[email] = emp
		fmt.Printf("After +5%%: salary = $%.2f\n", employees[email].Salary)
		fmt.Println()
	}

	for _, student := range students {
		fmt.Println(student.StudentRecord())
		fmt.Println()
	}
}
