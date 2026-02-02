package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//store details of one user
type Person struct {
	name string
	age  int
}


//takes name & age input. returns faslse when user enters exit.
func (p *Person) UserInput(reader *bufio.Reader) bool {

	fmt.Println("Enter your name(type 'EXIT' to exit application): ")
	name, _ := reader.ReadString('\n')
	p.name = strings.TrimSpace(name)


	if strings.EqualFold(p.name, "exit") {
		return false
	}


	fmt.Println("Enter your age: ")
	ageInput, _ := reader.ReadString('\n')
	ageInput = strings.TrimSpace(ageInput)


	age, err := strconv.Atoi(ageInput)
	if err != nil {
		fmt.Println("Age must be a number")
		return false
	}

	p.age = age
	return true
}


//checks if name & age are valid
func (p Person) validateInputs() (bool, bool) {

	isValidName := len(p.name) >= 2  
	isValidAge := p.age >= 1

	return isValidName, isValidAge
}

//prints user introduction
func (p Person) introduceMyself() {
	fmt.Printf("Hi my name is: %v and my age is: %v", p.name, p.age)
}


//to update user's age
func (p *Person) updateAge(reader *bufio.Reader) {
	fmt.Println("Enter your new age: ")
	inputNewAge, _ := reader.ReadString('\n')
	inputNewAge = strings.TrimSpace(inputNewAge)

	newAge, err := strconv.Atoi(inputNewAge)
	if err != nil {
		fmt.Println("Age must be a number")
		return
	}

	//new age must be greater than current age
	if newAge <= p.age {
		fmt.Println("Age must be greater than current age")
		return
	}

	p.age = newAge
	fmt.Println("Age updated succcessfully")

}

//to check user's eligibility for voting
func (p Person) voteEligibility() {
	if p.age >= 18 {
		fmt.Println("Congratulations! You are eligible to vote")
	} else {
		fmt.Println("Sorry you're not eligible to vote")
	}
}

func main() {

	user := Person{}						//create a new user
	reader := bufio.NewReader(os.Stdin)		//reader for taking users input

	//outer for loop to keep on asking for new users
	for {
		ok := user.UserInput(reader)
		if !ok {
			fmt.Println("Exiting Application")
			break
		}
		isValidName, isValidAge := user.validateInputs()

		if isValidName && isValidAge {

			//inner for loop to show menu for the same user
			for {
				fmt.Println("\n\nEnter a choice: \n1.Introduce yourself \n2.Update Age \n3.Check eligibility for voting\n4.Exit to enter new user")
				choiceInput, _ := reader.ReadString('\n')
				choiceInput = strings.TrimSpace(choiceInput)

				choice, err := strconv.Atoi(choiceInput)
				if err != nil {
					fmt.Println("Invalid choice. Enter a number")
					continue
				}

				//exits the current menu to take new user
				if choice == 4 {
					break
				}

				switch choice {
				case 1:
					user.introduceMyself()
				case 2:
					user.updateAge(reader)
				case 3:
					user.voteEligibility()
				default:
					fmt.Println("Invalid option. Choose between 1-4")
				}
			}
		} else {
	
			if !isValidName {
				fmt.Println("Your name is too short")
			}
			if !isValidAge {
				fmt.Println("Your age must be greater than 0")
			}
		}
	}

}
