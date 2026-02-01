package main

import (
	"fmt"
	"io"
	"slices"
	"bufio"
	"os"
	"strings"
)

type employee struct {
	empId         int
	empName       string
	empAge        int
	empSalary     int
	empDepartment string
}

type department struct {
	depName string
	empList []employee
}


func (d *department) addEmployeeMethod(e *employee) {
	d.empList = append(d.empList, *e)
}

func (d department) avgSalaryMethod() {
	if len(d.empList) == 0{
		fmt.Println("No employees in department")
		return
	}

	sum := 0
	count := 0
	fmt.Println("length    ", len(d.empList))
	for _, value := range d.empList {
		sum += value.empSalary
		count++
	}

	fmt.Println("count:::", count)
	averageSalary := sum / count
	fmt.Printf("The average salary of all employees is: %d", averageSalary)

}

func (d *department) removeEmployeeMethod(id int) {
	for i, val := range d.empList {
		if val.empId == id {
			d.empList = slices.Delete(d.empList, i, i+1)
			fmt.Println("Employee deleted successfully!")
			return
		}
	}
	fmt.Println("Employe with given id not found")
}
func (d *department) giveRaise(id int, raise int) {
	for i := range d.empList {
		if d.empList[i].empId == id {
			d.empList[i].empSalary += raise
			fmt.Printf("Raise %v given to employee with id %v",raise,id)
			return
		}
	}
}

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
	var raise int
	var id int

	reader := bufio.NewReader(os.Stdin)
	dept := department{"", make([]employee, 0)}
	

	for {

		fmt.Println("\n\nWelcome to employee database & department management")
		fmt.Println("\n\nEnter the operation u wanna perform: ")
		fmt.Println("1.Add Employee \n2.Remove Employee \n3.Average salary of department \n4.Give raise to employee\n5.Exit")
		
		if _,err:=fmt.Scan(&chooseOperation); findError(err){
			fmt.Scanln()
			continue
		}

		if chooseOperation == 5 {
			fmt.Println("Exiting Application")
			break
		}

		
		switch chooseOperation {
		case 1:
			emp := employee{}

			fmt.Println("Enter employee details: ")
			fmt.Println("Employee Id: ")
			if _,err:= fmt.Scan(&emp.empId); findError(err){
				fmt.Scanln()
				continue
			}
			fmt.Scanln() //clear buffer before reading string


			fmt.Println("Employe name: ")
			nameInput,_ := reader.ReadString('\n')
			emp.empName = strings.TrimSpace(nameInput)

			fmt.Println("Employee age: ")
			if _, err := fmt.Scan(&emp.empAge); findError(err) || emp.empAge <=0 {
				fmt.Println("Age must be greater than 0")
				fmt.Scanln()
				continue
			}
		


			fmt.Println("Employee Salary: ")
			if _, err := fmt.Scan(&emp.empSalary); findError(err){
				fmt.Println("Salary must be greater than 0")
				fmt.Scanln()
				continue
			}

			fmt.Scanln() //clear buffer to take string

		
			fmt.Println("Employe department(GoLang/DevOps/DotNet): ")
			deptInput,_ := reader.ReadString('\n')
			emp.empDepartment = strings.TrimSpace(deptInput)

			fmt.Printf("Hi my id is %v, name is %v, my age is %v, my salary is %v, my dept is %v", emp.empId, emp.empName, emp.empAge, emp.empSalary, emp.empDepartment)
			
			dept.addEmployeeMethod(&emp)
			fmt.Printf("Entire list: %v", dept.empList)

		case 2:
			fmt.Println("Enter id of employee you want to remove: ")
			if _, err := fmt.Scan(&id); findError(err){
				fmt.Scanln()
				continue
			}

			dept.removeEmployeeMethod(id)
			fmt.Println("Employee successfully rmeoved")
			fmt.Printf("New list: %v", dept.empList)

		case 3:
			dept.avgSalaryMethod()

		case 4:
			fmt.Printf("Enter id of employee you want to give a raise to: ")
			if _, err := fmt.Scan(&id); findError(err){
				fmt.Scanln()
				continue
			}
			
			fmt.Println("Enter the raise amount: ")
			if _, err := fmt.Scan(&raise); findError(err) || raise<=0{
				fmt.Println("Raise must be greater than 0")
				fmt.Scanln()
				continue
			}
			
			dept.giveRaise(id, raise)
			fmt.Printf("New list: %v", dept.empList)
			
		default:
			fmt.Println("Invalid option. Choose between 1-5")
		}

	}
}
