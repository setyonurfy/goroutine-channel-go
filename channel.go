package goroutine_go

func Kuadrat(angka []int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range angka {
			out <- n * n
		}
	}()
	return out
}

func Genap(angka []int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for _, n := range angka {
			if n%2 == 0 {
				ch <- n
			}
		}

	}()
	return ch
}
