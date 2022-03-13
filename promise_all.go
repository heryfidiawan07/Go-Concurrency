package main

import (
    "fmt"
    "math/rand"
)

func task() <-chan int32 {
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
    one, two, three := task(), task(), task()
    a, b, c := <-one, <-two, <-three
    fmt.Println(a, b, c)
    result := allResult(a, b, c)
    fmt.Println(result)
}

// Javascript
// Promise.all()
// *************

// const task = async () => {
// 	// simulate a workload
// 	// sleep(3000);
// 	return Math.floor(Math.random() * Math.floor(100));
// };

// const [a, b, c] = await Promise.all(
// 	task(),
// 	task(),
// 	task()
// );
// console.log(a, b, c);