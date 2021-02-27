package models

import "fmt"

var modelConfigs map[string]interface{}

func init() {
    fmt.Println("init models")

    // modelConfigs = map[string]interface{}{
    //     "initiated": []string{"test", "dude", "yo"},
    // }
}

