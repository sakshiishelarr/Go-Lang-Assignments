package main

import (
	"fmt"
	"slices"
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
	for i,val := range(d.empList){
		if val.empId==id{
			d.empList = slices.Delete(d.empList,i,i+1)
			return
		}
	}
}
func (d *department) giveRaise(id int,raise int) {
	for i := range d.empList {
		if d.empList[i].empId == id {
			d.empList[i].empSalary += raise
			return
		}
	}
}


func main() {
	
	var chooseOperation int
	var raise int
	var employeeList = make([]employee, 0)
	dept := department{"", employeeList}
	var id int

	for {

		fmt.Println("\n\nWelcome to employee database & department management")

		fmt.Println("\n\nEnter the operation u wanna perform: ")
		fmt.Println("1.Add Employee \n2.Remove Employee \n3.Average salary of department \n4.Give raise to employee\n5.Exit")
		fmt.Scan(&chooseOperation)

		if chooseOperation==5{
			break
		}

		emp := employee{}
		switch chooseOperation{
		case 1:
			fmt.Println("Enter employee details: ")
			fmt.Println("Employee Id: ")
			fmt.Scan(&emp.empId)

			fmt.Println("Employe name: ")
			fmt.Scan(&emp.empName)

			fmt.Println("Employee age: ")
			fmt.Scan(&emp.empAge)

			fmt.Println("Employee Salary: ")
			fmt.Scan(&emp.empSalary)

			fmt.Println("Employe department(GoLang/DevOps/DotNet): ")
			fmt.Scan(&emp.empDepartment)

			fmt.Printf("Hi my id is %v, name is %v, my age is %v, my salary is %v, my dept is %v", emp.empId, emp.empName, emp.empAge, emp.empSalary, emp.empDepartment)
			dept.addEmployeeMethod(&emp)
			fmt.Printf("Entire list: %v", dept.empList)

		case 2:
			fmt.Println("Enter id of employee you want to remove: ")
			fmt.Scan(&id)
			dept.removeEmployeeMethod(id)
			fmt.Println("Employee successfully rmeoved")
			fmt.Printf("New list: %v",dept.empList)
		case 3:
			dept.avgSalaryMethod()
		case 4:
			fmt.Printf("Enter id of employee you want to give a raise to: ")
			fmt.Scan(&id)
			fmt.Println("Enter the raise amount: ")
			fmt.Scan(&raise)

			dept.giveRaise(id,raise)
			fmt.Printf("New list: %v",dept.empList)
		}
		
	}
}
