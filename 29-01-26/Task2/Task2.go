package main

import "fmt"

type employee struct {
	empId         int
	empName       string
	empAge        int
	empSalary     int
	empDepartment string
}

func (e *employee) giveRaise(raise int) {
	e.empSalary += raise
}

type department struct {
	depName string
	empList []employee
}

func (d department) addEmployeeMethod(e *employee) {
	d.empList = append(d.empList, *e)
}

func (d department) avgSalaryMethod(e *employee) {
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

func (d department) removeEmployeeMethod(e employee) {
	e = employee{}
}

func main() {
	emp := employee{}

	var raise int

	for {
		deptname := ""
		fmt.Println("Employe department(GoLang/DevOps/DotNet): ")
		fmt.Scan(&deptname)
		emp.empDepartment = deptname

		fmt.Println("Enter employee details: ")
		fmt.Println("Employee Id: ")
		fmt.Scan(&emp.empId)

		fmt.Println("Employe name: ")
		fmt.Scan(&emp.empName)

		fmt.Println("Employee age: ")
		fmt.Scan(&emp.empAge)

		fmt.Println("Employee Salary: ")
		fmt.Scan(&emp.empSalary)

		fmt.Printf("Hi my id is %v, name is %v, my age is %v, my salary is %v, my dept is %v", emp.empId, emp.empName, emp.empAge, emp.empSalary, emp.empDepartment)

		fmt.Println("\n\nWelcome to employee database & department management")

		// fmt.Println("\n\nEnter the operation u wanna perform: ")
		// fmt.Println("1.Add Employee \n2.Remove Employee \n3.Average salary of department \n4.Give raise to employee")
		// fmt.Scan(&chooseOperation)

		var employeeList = make([]employee, 0)

		dept := department{deptname, employeeList}

		employeeList = append(employeeList, emp)
		dept.addEmployeeMethod(&emp)
		fmt.Printf("Entire list: %v", employeeList)

		dept.avgSalaryMethod(&emp)
		dept.removeEmployeeMethod(emp)

		

		fmt.Println("Enter the raise amount: ")
		fmt.Scan(&raise)
		emp.giveRaise(raise)

	}
}
