package func_example

import "time"
import "fmt"
import "math/rand"
import "math"

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

