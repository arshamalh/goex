package concurrency

func Sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

// Main

// s := []int{7, 2, 8, -9, 4, 0}

// c := make(chan int)
// middle := len(s) / 2
// go sum(s[:middle], c)
// go sum(s[middle:], c)
// x := <-c // receive from c
// y := <-c

// fmt.Println(x, y, x+y)