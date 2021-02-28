package models

import "fmt"

type User struct {
    Name string
    Age uint8
}

func init() {
    fmt.Println("init user model")

    // try appending slice inside of an interface{}
    var initiatedModels []string =
        modelConfigs["initiated"].([]string)
    initiatedModels = append(initiatedModels, "user")
    modelConfigs["initiated"] = initiatedModels

    fmt.Println("initiated model", modelConfigs["initiated"])
}

