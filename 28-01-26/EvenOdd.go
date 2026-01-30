package main

import (
	"sync"
	"fmt"
	"time"
)

var wg sync.WaitGroup

func addEven(evenNumbers []int,wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(5*time.Second)
	
	sum:=0
	for _,val:=range evenNumbers{
		sum += val
	}
	fmt.Println("Even Sum: ", sum)
}

func addOdd(oddNumbers []int,wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(5*time.Second)

	sum:=0
	for _,val:=range oddNumbers{
		sum+=val
	}
	fmt.Println("Odd Sum: ",sum)
}

func main(){
	evenNumbers:=[]int{2,4,6,8,10}
	oddNumbers:=[]int{1,3,5,7,9}

	wg.Add(2)

	go addEven(evenNumbers,&wg)
	go addOdd(oddNumbers,&wg)


	fmt.Println("Hi")
	wg.Wait()
}