package main
  
import (
	"fmt"
	// "time"
)

func task(text string) <-chan string {
	r := make(chan string)

	go func() {
		defer close(r)
        r <- text
	}()

	return r
}

func main() {
	sayHello := <-task("Hello")
    world := <-task(sayHello+" World !")
    // time.Sleep(time.Second * 3)
    fmt.Println(world)
}


// JavaScript
// Async Await
// **********

// const task = async () => {
// // simulate a workload
// sleep(3000);
// 	return Math.floor(Math.random() * Math.floor(100));
// };

// const r = await task();
// console.log(r);