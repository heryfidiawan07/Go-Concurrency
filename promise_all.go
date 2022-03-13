package main

import (
    "fmt"
    "math/rand"
)

func longRunningTask() <-chan int32 {
    r := make(chan int32)

    go func() {
        defer close(r)
        r <- rand.Int31n(100)
    }()

    return r
}

func allResult(a int32, b int32, c int32) int32 {
    return a + b + c
}

func main() {
    one, two, three := longRunningTask(), longRunningTask(), longRunningTask()
    a, b, c := <-one, <-two, <-three
    fmt.Println(a, b, c)
    result := allResult(a, b, c)
    fmt.Println(result)
}

// Javascript
// Promise.all()
// *************

// const longRunningTask = async () => {
// 	// simulate a workload
// 	// sleep(3000);
// 	return Math.floor(Math.random() * Math.floor(100));
// };

// const [a, b, c] = await Promise.all(
// 	longRunningTask(),
// 	longRunningTask(),
// 	longRunningTask()
// );
// console.log(a, b, c);