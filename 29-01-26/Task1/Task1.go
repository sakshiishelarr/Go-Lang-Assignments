package main

import "fmt"

type Person struct{
	name string
	age int
}

func (p *Person) UserInput(){
	fmt.Println("Enter your name(type 'EXIT' to exit application): ")
	fmt.Scan(&p.name)


	fmt.Println("Enter your age: ")
	fmt.Scan(&p.age)


}

func (p Person) validateInputs() (bool, bool){
	
	isValidName:= len(p.name)>=2
	isValidAge:= p.age>0

	return isValidName, isValidAge
}

func (p Person) introduceMyself(){
	fmt.Printf("Hi my name is: %v and my age is: %v", p.name, p.age)
}

func(p *Person) updateAge(newAge int){
	p.age = newAge
}

func (p Person) voteEligibility(){
	if p.age>=18{
		fmt.Println("Congratulations! You are eligible to vote")
	} else {
		fmt.Println("Sorry you're not eligible to vote")
	}
}

func main(){
	user:= Person{}
	var choice int
	var newAge int
	for {
		user.UserInput()
		if user.name=="EXIT"{
			break
		}
		isValidName,isValidAge := user.validateInputs()


		if isValidName && isValidAge{
			for {
				fmt.Println("\n\nEnter a choice: \n1.Introduce yourself \n2.Update Age \n3.Check eligibility for voting\n4.Exit to enter new user")
				fmt.Scan(&choice)

				if choice==4 {
					break
				}
				switch choice{
				case 1: 
					user.introduceMyself()
				case 2: 
					fmt.Println("Enter your new age: ")
					fmt.Scan(&newAge)
					user.updateAge(newAge)
				case 3:
					user.voteEligibility()
				}
			}
		} else {
			if !isValidName{
				fmt.Println("Your name is too short")
			}
			if !isValidAge{
				fmt.Println("Your age is negative")
			}
		}
	}
	


}