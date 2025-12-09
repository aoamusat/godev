// Package main provides employee management functionality
package structs

import "time"

// Employee represents an employee in the organization with their details
type Employee struct {
	ID        int       // Unique identifier for the employee
	Name      string    // Full name of the employee
	Address   string    // Current address of the employee
	DoB       time.Time // Date of birth of the employee
	Position  string    // Current job position/title
	Salary    int       // Current salary amount
	ManagerID int       // ID of the employee's manager
}

// Global employee instance
var eddy Employee

// EmployeeDB stores all employees mapped by their IDs
var EmployeeDB = make(map[int]Employee)

func main() {

	// Reduce Eddy's salary twice by 5000
	eddy.Salary -= 5000    // demerit for laboring without a break
	(&eddy).Salary -= 5000 // equivalent to the previous, but with explicit dereferencing

	// Promote Eddy to Senior position
	position := &eddy.Position
	*position = "Senior " + *position // promoted, for outsourcing to Elbonia

	// Update employee of the month's position
	var employeeOfTheMonth *Employee = &eddy
	employeeOfTheMonth.Position += " (proactive team player)"
}

// GetEmployeeByID retrieves an employee from the database by their ID
// Returns nil if employee is not found
func GetEmployeeByID(id int) *Employee {
	employee, exists := EmployeeDB[id]
	if !exists {
		return nil
	}
	return &employee
}
