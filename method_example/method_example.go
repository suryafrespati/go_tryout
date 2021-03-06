package method_example

import "fmt"

func Init() {
	fmt.Println("Init method_example")

	var car1 Car = Car{
		name:      "Car 1",
		brand:     "Toyota",
		sellPrice: 25000,
	}

	car1.showInfo()

	car1.sellPrice = 47000
	car1.showInfo()

	car2 := Car{
		name:  "Car 2",
		brand: "Ford",
	}
	car2.showInfo()

	car1.setSellPrice(39000)
	car1.showInfo()
}

type Car struct {
	name      string
	brand     string
	sellPrice uint
}

func (c Car) getPrice() uint {
	return c.sellPrice
}

func (c *Car) setSellPrice(price uint) {
	c.sellPrice = price
}

func (c Car) showInfo() {
	fmt.Println(c.name, c.brand, c.sellPrice)
}
