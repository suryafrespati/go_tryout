package map_example

import "fmt"

func Init() {
	fmt.Println("Init map")

	//	tryMapOfMap()
	sliceOfMap()
}

func sliceOfMap() {
	users := []map[string]string{}
	users = append(users, createUser("Surya"))
	users = append(users, createUser("Firdaus"))
	users = append(users, createUser("Respati"))

	for _, u := range users {
		fmt.Println(u["name"])
	}

	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	var products [11]map[string]int
	products[0] = map[string]int{
		"Macbook": 0,
	}
}

func createUser(name string) map[string]string {
	return map[string]string{
		"name": name,
	}
}

func tryMapOfMap() {

	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50

	usersById := map[string]map[string]string{}

	usersById["uuid1"] = map[string]string{
		"email": "suryafrespati@gmail.com",
	}
}
