package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"strconv"
)

// stores data of an empoloyee
type employee struct {
	empId     int
	empName   string
	empAge    int
	empSalary int
}

// name of department and employee list
type department struct {
	depName string
	empList []employee
}

// adds new employee
func (d *department) addEmployeeMethod(reader *bufio.Reader) {
	emp := employee{} //instance of employee struct

	fmt.Println("Enter unique employee id: ")
	inputId,_ := reader.ReadString('\n')
	inputId = strings.TrimSpace(inputId)
	id, err := strconv.Atoi(inputId)
	if err!=nil || id<=0 {
		fmt.Println("Id should be a positive number")
		return
	}
	for i:= range d.empList{
		if d.empList[i].empId == id{
			fmt.Println("Id already exists. Cannot add 2 employees with same id!")
			return
		}
	}
	emp.empId = id

	
	fmt.Println("Enter employee name: ")
	name, err := reader.ReadString('\n')
	emp.empName = strings.TrimSpace(name)
	if findError(err) {
		return
	}

	fmt.Println("Enter employee age: ")
	inputAge,_ := reader.ReadString('\n')
	inputAge = strings.TrimSpace(inputAge)
	age, err := strconv.Atoi(inputAge)
	if err!=nil || age<=0{
		fmt.Println("Age should be a positive number")
		return
	}
	emp.empAge = age

	fmt.Println("Enter employee salary: ")
	inputSalary,_ := reader.ReadString('\n')
	inputSalary = strings.TrimSpace(inputSalary)
	salary, err := strconv.Atoi(inputSalary)
	if err!=nil || salary<=0{
		fmt.Println("Salary should be a positive number")
		return
	}
	emp.empSalary = salary

	
	d.empList = append(d.empList, emp)
	fmt.Println("Employee added successfully!")
	fmt.Println("Company Data: ")
	d.printDepartment()

}

// prints entire company data
func (d *department) printDepartment() {
	fmt.Println("\nDepartment:", d.depName)

	if len(d.empList) == 0 {
		fmt.Println("No employees")
	}

	for _, emp := range d.empList {
		fmt.Printf("\n  ID: %v | Name: %v | Age: %v | Salary: %v\n", emp.empId, emp.empName, emp.empAge, emp.empSalary)
	}
}

// remove employee based on ID
func (d *department) removeEmployeeMethod(reader *bufio.Reader) {
	if len(d.empList) == 0 {
		fmt.Println("No employees to remove")
		return
	}

	fmt.Println("\nExisting data: ")
	d.printDepartment()

	var id int
	fmt.Println("Enter Employee Id: ")
	inputId,_:=reader.ReadString('\n')
	inputId = strings.TrimSpace(inputId)

	id,err := strconv.Atoi(inputId)
	if findError(err){
		return
	}



	for i, emp := range d.empList {
		if emp.empId == id {
			d.empList = slices.Delete(d.empList, i, i+1)
			fmt.Println("\nEmployee removed successfully")
			fmt.Printf("\nNew data: ")
			d.printDepartment()
			return
		}
	}

	fmt.Println("Employee not found ")
}

// average salary of one department
func (d *department) avgSalaryMethod() {
	fmt.Println("\nExisting data: ")
	d.printDepartment()

	if len(d.empList) == 0 {
		fmt.Println("No employees in this department")
		return
	}

	sum := 0
	for _, emp := range d.empList {
		sum += emp.empSalary
	}

	averageSalary := sum / len(d.empList)
	fmt.Printf("Average salary of %v department: %v\n", d.depName, averageSalary)
}

// give raise to employee based on id
func (d *department) giveRaise(reader *bufio.Reader) {
	if len(d.empList) == 0 {
		fmt.Println("No employees in department")
		return
	}

	fmt.Printf("\nExisting data: ")
	d.printDepartment()

	

	fmt.Println("Enter employee id: ")
	inputId,_:=reader.ReadString('\n')
	inputId = strings.TrimSpace(inputId)
	id,err := strconv.Atoi(inputId)
	if findError(err){
		return
	}

	fmt.Println("Enter raise amount: ")
	inputRaise,_:=reader.ReadString('\n')
	inputRaise = strings.TrimSpace(inputRaise)
	raise,err := strconv.Atoi(inputRaise)
	if findError(err){
		return
	}

	for i := range d.empList {
		if d.empList[i].empId == id {
			d.empList[i].empSalary += raise
			fmt.Printf("Raise given to employee %v ", id)
			fmt.Println("New data: ")
			d.printDepartment()
			return
		}
	}
	fmt.Println("Employee not found")
}

// generic error fucntion
func findError(err error) bool {
	if err != nil {
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			fmt.Println("Input finished or Unexpected end of input", err)
		} else {
			fmt.Println("Error reading input ", err)
		}
		return true
	}
	return false
}

func main() {

	var chooseOperation int
	var err error

	reader := bufio.NewReader(os.Stdin) //instance of bufio struct

	//instance of department struct
	department1 := make([]department, 0)

	for {

		fmt.Println("\n\nWelcome to employee database & department management")
		fmt.Println("\n\nEnter the operation u wanna perform: ")
		fmt.Println("1.Add Employee \n2.Remove Employee \n3.Average salary of department \n4.Give raise to employee\n5.Exit")

		_, err = fmt.Scan(&chooseOperation)
		reader.ReadString('\n') // clears \n or else we wont move ahead to next input
		findError(err)

		if chooseOperation == 5 {
			fmt.Println("Exiting Application")
			break
		}

		for chooseOperation != 5 {
			fmt.Println("Enter department name: ")
			inputDeptName, _ := reader.ReadString('\n')
			inputDeptName = strings.TrimSpace(inputDeptName)

			found := false

			for i := range department1 {
				if department1[i].depName == inputDeptName {
					found = true

					switch chooseOperation {
					case 1:
						department1[i].addEmployeeMethod(reader)

					case 2:
						department1[i].removeEmployeeMethod(reader)

					case 3:
						department1[i].avgSalaryMethod()

					case 4:
						department1[i].giveRaise(reader)

					default:
						fmt.Println("Invalid option. Choose between 1-5")
					}
					break

				}
			}

			if !found && chooseOperation == 1 {
				newDept := department{
					depName: inputDeptName,
					empList: make([]employee, 0),
				}
				newDept.addEmployeeMethod(reader)

				department1 = append(department1, newDept)

			} else if !found {
				fmt.Println("Department not found")

			} else {
				if len(department1) == 0 {
					fmt.Println("No departments available")
					continue
				}
			}
			break //did not move to next iteration
		}
	}
}
