package array_slice_example

import "fmt"

func Init() {
	fmt.Printf("\n# --- array_slice_example\n\n")

	capExample()
	copyExample()
	slice3IndexesExample()
	appendExample()
	triangleExample()

	fmt.Printf("\n# --- end of array_slice_example\n\n")
}

func capExample() {
	fmt.Printf("\ncapExample()\n\n")

	var letters = []string{"A", "B", "C", "D"}
	fmt.Printf("letters: %v \n", letters)
	fmt.Printf("cap(letters): %v \n\n", cap(letters))

	var lettersA = letters[0:3]
	fmt.Printf("len(lettersA): %v \n", len(lettersA))
	fmt.Printf("cap(lettersA): %v \n\n", cap(lettersA))

	var lettersB = letters[1:4]
	fmt.Printf("len(lettersB): %v \n", len(lettersB))
	fmt.Printf("cap(lettersB): %v \n\n", cap(lettersB))

	fmt.Printf("\nend of capExample()\n\n")
}

func slice3IndexesExample() {
	fmt.Printf("\nslice3IndexesExample()\n\n")

	var brands = []string{"samsung", "apple", "xiaomi", "motorola", "lenovo"}
	var brandsA = brands[0:4]
	var brandsB = brands[0:2:3]

	fmt.Printf("brands: %v \n", brands)
	fmt.Printf("len(brands): %v \n", len(brands))
	fmt.Printf("cap(brands): %v \n\n", cap(brands))

	fmt.Printf("brandsA: %v \n", brandsA)
	fmt.Printf("len(brandsA): %v \n", len(brandsA))
	fmt.Printf("cap(brandsA): %v \n\n", cap(brandsA))

	fmt.Printf("brandsB: %v \n", brandsB)
	fmt.Printf("len(brandsB): %v \n", len(brandsB))
	fmt.Printf("cap(brandsB); %v \n\n", cap(brandsB))

	fmt.Printf("\nend of slice3IndexesExample()\n\n")
}

func appendExample() {
	fmt.Printf("\nappendExample()\n\n")

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
		fmt.Printf("%v \n", item)
	}

	fmt.Printf("\n")

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

	fmt.Printf("\nend of appendExample()\n\n")
}

func triangleExample() {
	fmt.Printf("\ntriangleExample()\n\n")

	size := 7

	for i := 0; i < size; i++ {
		for j := size - i; j >= 0; j-- {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("$ ")
		}
		fmt.Print("\n")
	}

	fmt.Printf("\nend of triangleExample()\n\n")
}

func copyExample() {
	fmt.Printf("\ncopyExample()\n\n")

	// make slice of string with capacity of 8
	// if slice's size is not allocated, thus the copy() won't work
	dst := make([]string, 8)
	src := []string{"gibson", "fender", "lespaul", "ibanez", "cort"}
	n := copy(dst, src)
	// copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println("len of slice:", n) // print len of slice

	fmt.Printf("\n")

	for i, d := range dst {
		if d == "" {
			d = "-"
		}
		fmt.Printf("%d %v \n", i+1, d)
	}

	fmt.Printf("\n")

	fmt.Printf("\nend of copyExample()\n\n")
}
