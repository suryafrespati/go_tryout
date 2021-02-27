package models

import "fmt"

type User struct {
    Name string
    Age uint8
}

func init() {
    fmt.Println("init user model")

    // fmt.Println(modelConfigs["initiated"])

    // for _, v := range modelConfigs["initiated"] {
    //     fmt.Println(v)
    // }

    // var temp []string = modelConfigs["initiated"]
    // temp = append(temp, "user")
    // modelConfigs["initiated"] = temp
}

