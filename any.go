package fngo

type Any interface{}

func NewAnyCh(values ...Any) chan Any {
	out := make(chan Any)

	if len(values) > 0 {
		SendAny(out, values...)
	}

	return out
}

func SendAny(ch chan Any, values ...Any) {
	go func() {
		for _, val := range values {
			ch <- val
		}

		close(ch)
	}()
}

func MapAnyToAny(in chan Any, mapper func(Any) Any) chan Any {
	out := make(chan Any)

	go func() {
		for val := range in {
			out <- mapper(val)
		}

		close(out)
	}()

	return out
}

func MapAnyToString(in chan Any, mapper func(Any) string) chan string {
	out := make(chan string)

	go func() {
		for val := range in {
			out <- mapper(val)
		}

		close(out)
	}()

	return out
}

func ForEachAny(in chan Any, handler func(Any)) {
	for val := range in {
		handler(val)
	}
}

func ToAnySlice(in chan Any) []Any {
	var ret = []Any{}

	for val := range in {
		ret = append(ret, val)
	}

	return ret
}
