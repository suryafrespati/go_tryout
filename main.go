package main

import (
    "fmt"
    "tryout/array_slice_example"
    "tryout/func_example"
    "tryout/map_example"
    "tryout/pointer_example"
    "tryout/struct_example"
    "tryout/embedded_struct_example"
    "tryout/method_example"
    "tryout/interface_example"
    "tryout/reflect_example"
    "tryout/goroutine_example"
    "tryout/stack"
    "tryout/library"
    "tryout/models"
)

type Environment map[string]string
type BaseConfig struct {
    stage string
    host string
    port uint16
}

func main() {
    fmt.Println("learn Golang")
    print("\n")

    var user1 = models.User{
        Name: "Surya",
        Age: 23,
    }

    fmt.Println(user1)

    var baseEnv = BaseConfig{
        stage: "local",
        host: "localhost",
        port: 8080,
    }

    var env = Environment{
        "name": "go_dsa",
    }

    for _, v := range env {
        fmt.Println("env val", v)
    }

    _ = env
    _ = baseEnv

    library.Init()
    stack.Init()
    array_slice_example.Init()
    map_example.Init()
    func_example.Init()
    pointer_example.Init()
    struct_example.Init()
    embedded_struct_example.Init()
    method_example.Init()
    interface_example.Init()
    reflect_example.Init()
    goroutine_example.Init()

    fmt.Println("--- end of file")
}

