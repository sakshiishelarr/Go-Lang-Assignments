package main

import (
	"fmt"
	"sync"
)

//sends jobs onto channel
func sendData(myChannel chan<- int, numberOfJobs int) { 
	for i := 0; i < numberOfJobs; i++ {
		myChannel <- i  
	}

	close(myChannel) 	
	fmt.Println("Jobs added to channel")
}

//processes jobs per worker
func worker(myChannel <-chan int, numberOfWorkers int, wg *sync.WaitGroup) { 

	//creates workers
	for i := 0; i < numberOfWorkers; i++ {		

		go func(workerID int) {					
			defer wg.Done()

			for value := range myChannel {		
				fmt.Printf("Job %d processed by worker %d\n", value, workerID)
			}
		}(i)
	}
}

func main() {
	var numberOfJobs int
	var numberOfWorkers int

	var wg sync.WaitGroup

	fmt.Println("Enter number of jobs:")
	fmt.Scan(&numberOfJobs)

	fmt.Println("Enter number of workers:")
	fmt.Scan(&numberOfWorkers)

	myChannel := make(chan int, numberOfJobs)		
	
	wg.Add(numberOfWorkers)						
	worker(myChannel, numberOfWorkers, &wg)

	sendData(myChannel, numberOfJobs)

	wg.Wait()
	fmt.Println("All jobs processed")
}
