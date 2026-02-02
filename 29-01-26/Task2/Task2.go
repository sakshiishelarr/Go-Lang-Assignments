package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type employee struct {
	empId     int
	empName   string
	empAge    int
	empSalary int
}

type department struct {
	depName string
	empList []employee
}

func (d *department) addEmployeeMethod(reader *bufio.Reader) {
	emp := employee{}
	var err error

	fmt.Println("Enter employee id: ")
	_, err = fmt.Scan(&emp.empId)
	findError(err)
	fmt.Scanln()

	fmt.Println("Enter employee name: ")
	name, _ := reader.ReadString('\n')
	emp.empName = strings.TrimSpace(name)

	fmt.Println("Enter employee age: ")
	_, err = fmt.Scan(&emp.empAge)
	findError(err)

	fmt.Println("Enter employe salary: ")
	_, err = fmt.Scan(&emp.empSalary)
	findError(err)
	fmt.Scanln()

	d.empList = append(d.empList, emp)
	fmt.Println("Employee added successfully!")
	fmt.Println("Company Data: ")
	d.printDepartment()

}

func (d *department) printDepartment() {
	fmt.Println("\nDepartment:", d.depName)

	if len(d.empList) == 0 {
		fmt.Println("No employees")
	}

	for _, emp := range d.empList {
		fmt.Printf("\n  ID: %v | Name: %v | Age: %v | Salary: %v\n", emp.empId, emp.empName, emp.empAge, emp.empSalary)
	}
}

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
	fmt.Printf("Average salary of %s department: %d\n", d.depName, averageSalary)
}

func (d *department) removeEmployeeMethod(reader *bufio.Reader) {
	if len(d.empList) == 0 {
		fmt.Println("No employees to remove")
		return
	}

	fmt.Println("\nExisting data: ")
	d.printDepartment()

	var id int
	fmt.Println("Enter Employee Id: ")
	fmt.Scan(&id)
	reader.ReadString('\n')

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

func (d *department) giveRaise() {
	if len(d.empList) == 0 {
		fmt.Println("No employees in department")
		return
	}

	fmt.Printf("\nExisting data: ")
	d.printDepartment()

	var id int
	var raise int

	fmt.Println("Enter employee id: ")
	fmt.Scan(&id)

	fmt.Println("Enter raise amount: ")
	fmt.Scan(&raise)

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

	reader := bufio.NewReader(os.Stdin)
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
						department1[i].giveRaise()

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

				for i := range department1 {
					department1[i].printDepartment()
				}
			}
			break //did not move to next iteration
		}
	}
}
