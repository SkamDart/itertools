package itertools

const (
	DEFAULT_CHANNEL_SIZE = 10
)

// Infinite Iterators
func CountStep(start, step int) <-chan int {
	// inexpensive operation, so buffer DEFAULT_CHANNEL_SIZE
	ch := make(chan int, DEFAULT_CHANNEL_SIZE)
	go func() {
		i := start
		for {
			ch <- i
			i += step
		}
		close(ch)
	}()
	return ch
}

func Count(start int) <-chan int {
	return CountStep(start, 1)
}

func CycleByte(s string) <-chan byte {
	ch := make(chan byte, DEFAULT_CHANNEL_SIZE)
	size := len(s)
	go func() {
		i := 0
		for {
			ch <- s[i]
			i = (i + 1) % size
		}
		close(ch)
	}()
	return ch
}

//func CycleRune(s string) <-chan rune {
//	c := make(chan rune, DEFAULT_CHANNEL_SIZE)
//	go func() {
//		for {
//			for _, r := range s {
//				c <- r
//			}
//		}
//	}()
//	return c
//}

// Terminating (by shortest arg) Iterators
func TakeWhileString(pred func(rune) bool, seq string) <-chan rune {
	ch := make(chan rune, DEFAULT_CHANNEL_SIZE)
	go func() {
		for _, s := range seq {
			if !pred(s) {
				break
			}
			ch <- s
		}
		close(ch)
	}()
	return ch
}

func TakeWhileInt(pred func(int) bool, seq []int) <-chan int {
	ch := make(chan int, DEFAULT_CHANNEL_SIZE)
	go func() {
		for _, c := range seq {
			if !pred(c) {
				break
			}
			ch <- c
		}
		close(ch)
	}()
	return ch
}

func FilterString(pred func(rune) bool, seq string) <-chan rune {
	ch := make(chan rune, DEFAULT_CHANNEL_SIZE)
	go func() {
		for _, s := range seq {
			if pred(s) {
				ch <- s
			}
		}
		close(ch)
	}()
	return ch
}

func FilterInt(pred func(int) bool, seq []int) <-chan int {
	ch := make(chan int, DEFAULT_CHANNEL_SIZE)
	go func() {
		for _, c := range seq {
			if pred(c) {
				ch <- c
			}
		}
		close(ch)
	}()
	return ch
}
