package main

import (
	"fmt"
	"sync"
)

func sendData(myChannel chan<- int, numberOfJobs int) { //send only channel
	for i := 0; i < numberOfJobs; i++ {
		myChannel <- i  //sends jobs onto the channel.Each channel waits until channel has space then puts job
	}
	close(myChannel) 	//closed channel or else workers will wait forever
	fmt.Println("Jobs added to channel")
}

func worker(myChannel <-chan int, numberOfWorkers int, wg *sync.WaitGroup) { //receive only channel. uses wg to indicate compleiton

	
	for i := 0; i < numberOfWorkers; i++ {		//creates workers

		go func(workerID int) {					//new go routine per worker
			defer wg.Done()

			for value := range myChannel {		//keep receiving from channel
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

	myChannel := make(chan int, numberOfJobs)		//buffered channel so sender doesnt blocck immediately

	
	wg.Add(numberOfWorkers)						
	go worker(myChannel, numberOfWorkers, &wg)

	go sendData(myChannel, numberOfJobs)

	wg.Wait()
	fmt.Println("All jobs processed")
}
