package itertools

// Infinite Iterators
func CountStep(start, step int) <-chan int {
	// inexpensive operation, so buffer 10
	c := make(chan int, 10)
	go func() {
		i := start
		for {
			c <- i
			i += step
		}
	}()
	return c
}

func Count(start int) <-chan int {
	return CountStep(start, 1)
}

func CycleByte(s string) <-chan byte {
	c := make(chan byte, 10)
	size := len(s)
	go func() {
		i := 0
		for {
			c <- s[i]
			i = (i + 1) % size
		}
	}()
	return c
}

//func CycleRune(s string) <-chan rune {
//	c := make(chan rune, 10)
//	go func() {
//		for {
//			for _, r := range s {
//				c <- r
//			}
//		}
//	}()
//	return c
//}
