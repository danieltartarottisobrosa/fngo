package fngo

import "sync"

func NewStringCh(values ...string) chan string {
	out := make(chan string)

	if len(values) > 0 {
		SendStrings(out, values...)
	}

	return out
}

func SendStrings(ch chan string, values ...string) {
	go func() {
		for _, val := range values {
			ch <- val
		}

		close(ch)
	}()
}

func MapStringToString(in <-chan string, mapper func(string) string) chan string {
	out := make(chan string)

	go func() {
		for val := range in {
			out <- mapper(val)
		}

		close(out)
	}()

	return out
}

func MapStringToInt(in <-chan string, mapper func(string) int) chan int {
	out := make(chan int)

	go func() {
		for val := range in {
			out <- mapper(val)
		}

		close(out)
	}()

	return out
}

func ForEachString(in <-chan string, handler func(string)) {
	for val := range in {
		handler(val)
	}
}

func ToStringSlice(in <-chan string) []string {
	var ret = []string{}

	for val := range in {
		ret = append(ret, val)
	}

	return ret
}

func MergeStringChs(in ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	wg.Add(len(in))

	out := make(chan string)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan string) {
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
