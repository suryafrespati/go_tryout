package array_slice_example

import "fmt"

func Init() {

	fmt.Println("Init array & slice")

	var nums []uint8

	nums = append(nums, 12)
	nums = append(nums, 34)
	nums = append(nums, 91)
	nums = append(nums, 73)
	nums = append(nums, 234)
	nums = append(nums, 120)

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("$")
		}
		fmt.Print("\n")
	}

	capExample()
	copyExample()
	slice3IndexesExample()
	appendExample()
}

func capExample() {
	var fruits = []string{"apple", "grape", "banana", "melon", "papaya"}
	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4

	var aFruits = fruits[0:3]
	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4

	var bFruits = fruits[1:4]
	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3
}

func copyExample() {
	dst := make([]string, 2)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	// n := copy(dst, src)
	copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	// fmt.Println(n)   // 3

	for _, d := range dst {
		print(d, "\n")
	}

	print("\n\n")
}

func slice3IndexesExample() {
	var fruits = []string{"apple", "grape", "banana"}
	var aFruits = fruits[0:2]
	var bFruits = fruits[0:2:2]

	fmt.Println(fruits)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits)) // len: 3
	fmt.Println(cap(fruits)) // cap: 3

	fmt.Println(aFruits)      // ["apple", "grape"]
	fmt.Println(len(aFruits)) // len: 2
	fmt.Println(cap(aFruits)) // cap: 3

	fmt.Println(bFruits)      // ["apple", "grape"]
	fmt.Println(len(bFruits)) // len: 2
	fmt.Println(cap(bFruits)) // cap: 2
}

func appendExample() {

	// var list = []int{}
	list := make([]int, 0)
	_ = list

	list = append(list, 4)
	list = append(list, 7)
	list = append(list, 9)
	list = append(list, 12)
	list = append(list, 33)
	list = append(list, 50)

	for _, item := range list[:] {
		fmt.Print(item, "\n")
	}
}
