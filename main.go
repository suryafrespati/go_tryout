package main

import (
	"array_slice_example"
	"fmt"
	"func_example"
	"map_example"
    "pointer_example"
    "struct_example"
    "embedded_struct_example"
	"stack"
)

type Environment map[string]string
type BaseConfig struct {
    stage string
    host string
    port uint16
}

func main() {
	fmt.Println("learn Golang")
	print("\n\n\n")

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

	stack.Init()
	array_slice_example.Init()
	map_example.Init()
	func_example.Init()
	pointer_example.Init()
	struct_example.Init()
	embedded_struct_example.Init()

    fmt.Println("--- end of file")
}

