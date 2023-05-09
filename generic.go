func NewCh[T any](values ...T) chan T {
	out := make(chan T)

	if len(values) > 0 {
		SendCh(out, values...)
	}

	return out
}

func SendCh[T any](ch chan T, values ...T) {
	go func() {
		for _, val := range values {
			ch <- val
		}

		close(ch)
	}()
}

func MapCh[A any, B any](in <-chan A, mapper func(A) B) chan B {
	out := make(chan B)

	go func() {
		for val := range in {
			out <- mapper(val)
		}

		close(out)
	}()

	return out
}

func ForEachCh[T any](in <-chan T, handler func(T)) {
	for val := range in {
		handler(val)
	}
}

func ToSliceCh[T any](in <-chan T) []T {
	var ret = []T{}

	for val := range in {
		ret = append(ret, val)
	}

	return ret
}

func PrintCh[T any](in <-chan T) {
	ForEachCh(in, func(val T) {
		fmt.Println(val)
	})
}

func MergeChs[T any](in ...chan T) chan T {
	var wg sync.WaitGroup
	wg.Add(len(in))

	out := make(chan T)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan T) {
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
