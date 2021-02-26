package func_example

import "time"
import "fmt"
import "math/rand"
import "math"
import "strings"

func Init() {
	fmt.Println("Init func_example")

	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	var avg = calculateAll(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)

    closureExample()
    callbackFunctionExample()
}

func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

// func calculate(d float64) (float64, float64) {
// 	// hitung luas
// 	var area = math.Pi * math.Pow(d/2, 2)
// 	// hitung keliling
// 	var circumference = math.Pi * d
//
// 	// kembalikan 2 nilai
// 	return area, circumference
// }

/*
	predefined return value
*/
func calculate(d float64) (area float64, circumference float64) {
	// hitung luas
	area = math.Pi * math.Pow(d/2, 2)
	// hitung keliling
	circumference = math.Pi * d

	// kembalikan 2 nilai
	return
}

/*
	variadic function
*/
func calculateAll(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

func closureExample() {
    var getMinMax = func(n []int) (int, int) {
        var min, max int
        for i, e := range n {
            switch {
            case i == 0:
                max, min = e, e
            case e > max:
                max = e
            case e < min:
                min = e
            }
        }
        return min, max
    }

    var numbers = []int{2, 3, 4, 3, 4, 2, 3}
    var min, max = getMinMax(numbers)
    fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)
}

type FilterCallback func(string) bool
type ArrayString []string

func filter(data ArrayString, callback FilterCallback) ArrayString {
    var result ArrayString
    for _, each := range data {
        if filtered := callback(each); filtered {
            result = append(result, each)
        }
    }
    return result
}

func callbackFunctionExample() {
    var data = []string{"wick", "jason", "ethan"}
    var dataContainsO = filter(data, func(each string) bool {
        return strings.Contains(each, "o")
    })
    var dataLenght5 = filter(data, func(each string) bool {
        return len(each) == 5
    })

    fmt.Println("data asli \t\t:", data)
    // data asli : [wick jason ethan]

    fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
    // filter ada huruf "o" : [jason]

    fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
    // filter jumlah huruf "5" : [jason ethan]
}


