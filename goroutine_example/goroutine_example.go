package goroutine_example

import  "fmt"
import "runtime"
import "time"

func Init() {
    runtime.GOMAXPROCS(2)

    // basicExample()
    channelExample()
    bufferedChannelExample()

    fmt.Println()
}

func basicExample() {
    fmt.Println("basicExample()")

    var alert = func(text string, limit int) {
        for i := 0; i < limit; i++ {
            fmt.Println(text)
        }
    }

    fmt.Println("Init goroutine_example")

    var limit int = 7

    go alert("goroutine A!", limit)
    go alert("goroutine B!", limit)

    alert("alert", limit)

    var input string
    fmt.Scanln(&input)
}

func channelExample() {
    fmt.Println("channelExample()")

    var messages chan string = make(chan string)
    var chLen int = 0

    var sayHelloTo = func(who string) {
        var data = fmt.Sprintf("# %s", who)
        messages <- data
    }

    // go sayHelloTo("john wick")
    // go sayHelloTo("ethan hunt")
    // go sayHelloTo("jason bourne")

    // var message1 = <-messages
    // fmt.Println(message1)

    // var message2 = <-messages
    // fmt.Println(message2)

    // var message3 = <-messages
    // fmt.Println(message3)

    time.Sleep(1 * time.Millisecond)

    var pushToCh func(string, func(string)) = func (text string, cb func(string)) {
        go cb(text)

        // fmt.Println(<- messages)

        // to make sure each goroutine executed in correct order.
        time.Sleep(10 * time.Millisecond)

        chLen += 1
    }

    pushToCh("1 message. learn golang", sayHelloTo)
    pushToCh("2 message. write some code", sayHelloTo)
    pushToCh("3 message. create product", sayHelloTo)
    pushToCh("4 message. test product", sayHelloTo)
    pushToCh("5 message. push to market", sayHelloTo)

    for i := 0; i < chLen; i++ {
        fmt.Println(<- messages)
    }
}

func bufferedChannelExample() {
    fmt.Println("bufferedChannelExample()")

    messages := make(chan int, 2)

    go func() {
        for {
            i := <-messages
            fmt.Println("receive data", i)
        }
    }()

    for i := 0; i < 5; i++ {
        fmt.Println("send data", i)
        messages <- i
    }
}

