package main
  
import (
	"fmt"
	// "time"
)

func longRunningTask(text string) <-chan string {
	r := make(chan string)

	go func() {
		defer close(r)
        r <- text
	}()

	return r
}

func main() {
	sayHello := <-longRunningTask("Hello")
    world := <-longRunningTask(sayHello+" World !")
    // time.Sleep(time.Second * 3)
    fmt.Println(world)
}


// JavaScript
// Async Await
// **********

// const longRunningTask = async () => {
// // simulate a workload
// sleep(3000);
// 	return Math.floor(Math.random() * Math.floor(100));
// };

// const r = await longRunningTask();
// console.log(r);