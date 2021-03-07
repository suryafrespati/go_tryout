package map_example

import (
	"fmt"
	"strconv"
	"time"
)

func Init() {
	fmt.Printf("\n# --- map_example\n\n")
	tryMapOfMap()
	sliceOfMap()
	fmt.Printf("\n# --- end of map_example\n\n")
}

func createUser(name string) map[string]string {
	return map[string]string{
		"name": name,
	}
}

func sliceOfMap() {
	fmt.Printf("\nsliceOfMap()\n\n")

	users := []map[string]string{}

	users = append(users, createUser("Surya"))
	users = append(users, createUser("Martin"))
	users = append(users, createUser("Alexander"))

	for _, u := range users {
		fmt.Println(u["name"])
	}

	fmt.Printf("\n")

	var pets = []map[string]string{
		{
			"name":   "mueza",
			"animal": "cat",
			"gender": "male",
		},
		{
			"name":   "oscar",
			"animal": "otter",
			"gender": "male",
		},
		{
			"name":   "fae",
			"animal": "rabbit",
			"gender": "female",
		},
	}

	for _, pet := range pets {
		fmt.Println(pet["name"], pet["animal"], pet["gender"])
	}

	fmt.Printf("\n")

	const MAX_COMPONENT_GROUP uint = 4
	var carComponentForceThresholds [MAX_COMPONENT_GROUP]map[string]int

	carComponentForceThresholds[0] = map[string]int{
		"backTrunk":  5,
		"frontTrunk": 10,
	}

	carComponentForceThresholds[1] = map[string]int{
		"backBumper":  6,
		"frontBumper": 15,
	}

	carComponentForceThresholds[2] = map[string]int{
		"backWindow":  12,
		"frontWindow": 17,
		"leftWindow":  8,
		"rightWindow": 8,
	}

	for k, v := range carComponentForceThresholds {
		fmt.Printf("%v - %v \n", k, v)
	}

	fmt.Printf("\nend of sliceOfMap()\n\n")
}

func tryMapOfMap() {
	fmt.Printf("\ntryMapOfMap()\n\n")

	usersById := map[string]map[string]string{}

	usersById[generateUid()] = map[string]string{
		"email": "surya@test.com",
	}
	usersById[generateUid()] = map[string]string{
		"email": "orang@test.com",
	}
	usersById[generateUid()] = map[string]string{
		"email": "saykoji@test.com",
	}

	for k, v := range usersById {
		fmt.Printf("%v - %v \n", k, v)
	}

	fmt.Printf("\nend of tryMapOfMap()\n\n")
}

func generateUid() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
