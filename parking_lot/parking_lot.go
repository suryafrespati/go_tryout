package main

import "fmt"

func main() {

	fmt.Println("Init Parking Lot")

}

func init(size int) {
	var lotSize uint8 = size
}
