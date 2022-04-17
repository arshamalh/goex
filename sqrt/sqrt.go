package sqrt

import "fmt"

type ErrorNegativeSqrt float64

func (err ErrorNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(err))
}

func Sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, ErrorNegativeSqrt(num)
	}
	result := num
	for i := 0; i < 9; i++ {
		// We can also make it while true and break it whenever result was equal to num^2
		result -= (result*result - num) / (2 * result)
	}
	return result, nil
}

// Sqrt exercise main
// fmt.Println(sqrt.Sqrt(-2))
// fmt.Println(math.Sqrt(-2))
