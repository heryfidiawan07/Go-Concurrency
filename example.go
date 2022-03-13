package main

import "fmt"
import "runtime"

func main() {
    runtime.GOMAXPROCS(2)

    messages := make(chan string)

    say := func(who string) {
        // var data = fmt.Sprintf("hello %s", who)
		data := "hello " + who
        messages <- data
    }

    go say("john wick")
    go say("ethan hunt")
    go say("jason bourne")

    message1 := <-messages
    fmt.Println(message1)

    message2 := <-messages
    fmt.Println(message2)

    message3 := <-messages
    fmt.Println(message3)
}