package fngo

import "sync"

func NewIntCh(values ...int) chan int {
	out := make(chan int)

	if len(values) > 0 {
		SendInts(out, values...)
	}

	return out
}

func SendInts(ch chan int, values ...int) {
	go func() {
		for _, val := range values {
			ch <- val
		}

		close(ch)
	}()
}

func MapIntToInt(in <-chan int, mapper func(int) int) chan int {
	out := make(chan int)

	go func() {
		for val := range in {
			out <- mapper(val)
		}

		close(out)
	}()

	return out
}

func MapIntToString(in <-chan int, mapper func(int) string) chan string {
	out := make(chan string)

	go func() {
		for val := range in {
			out <- mapper(val)
		}

		close(out)
	}()

	return out
}

func ForEachInt(in <-chan int, handler func(int)) {
	for val := range in {
		handler(val)
	}
}

func ToIntSlice(in <-chan int) []int {
	var ret = []int{}

	for val := range in {
		ret = append(ret, val)
	}

	return ret
}

func MergeIntChs(in ...chan int) chan int {
	var wg sync.WaitGroup
	wg.Add(len(in))

	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	for _, c := range in {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
