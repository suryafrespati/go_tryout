package goroutine_example

import  "fmt"
import "runtime"

func Init() {
    runtime.GOMAXPROCS(2)

    // basicExample()
    channelExample()

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
        var data = fmt.Sprintf("hello %s", who)
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

    var pushToCh func(string, func(string)) = func (text string, cb func(string)) {
        go cb(text)
        chLen += 1
    }

    pushToCh("#1 message. learn golang", sayHelloTo)
    pushToCh("#2 message. write some code", sayHelloTo)
     pushToCh("#3 message. create product", sayHelloTo)

    for i := 0; i < chLen; i++ {
        fmt.Println(<- messages)
    }
}

