package fngo

import (
	"fmt"
	"testing"
)

func TestNewEmptyIntCh(t *testing.T) {
	// Act
	ch := NewCh[int]()

	// Assert
	if ch == nil {
		t.Error("The channel is nil")
	}
}

func TestNewPopulatedIntCh(t *testing.T) {
	// Arrange
	values := []int{0, 1, 2, 3}

	// Act
	ch := NewCh(values...)

	// Assert
	for _, expected := range values {
		if actual := <-ch; actual != expected {
			t.Errorf("Expected: %d, Actual: %d", expected, actual)
		}
	}
}

func TestSendInts(t *testing.T) {
	// Arrange
	ch := make(chan int)
	values := []int{0, 1, 2, 3}

	// Act
	SendCh(ch, values...)

	// Assert
	for _, expected := range values {
		if actual := <-ch; actual != expected {
			t.Errorf("Expected: %d, Actual: %d", expected, actual)
		}
	}
}

func TestMapIntToInt(t *testing.T) {
	// Arrange
	ch := make(chan int)
	expectedValues := []int{0, 10, 20, 30}

	tenTimes := func(val int) int {
		return val * 10
	}

	SendCh(ch, 0, 1, 2, 3)

	// Act
	ch = MapCh(ch, tenTimes)

	// Assert
	for _, expected := range expectedValues {
		if actual := <-ch; actual != expected {
			t.Errorf("Expected: %d, Actual: %d", expected, actual)
		}
	}
}

func TestMapIntToString(t *testing.T) {
	// Arrange
	intCh := make(chan int)
	expectedValues := []string{"Valor 0", "Valor 1", "Valor 2", "Valor 3"}

	toString := func(val int) string {
		return fmt.Sprintf("Valor %d", val)
	}

	SendCh(intCh, 0, 1, 2, 3)

	// Act
	strCh := MapCh(intCh, toString)

	// Assert
	for _, expected := range expectedValues {
		if actual := <-strCh; actual != expected {
			t.Errorf("Expected: %s, Actual: %s", expected, actual)
		}
	}
}

func TestForEachInt(t *testing.T) {
	// Arrange
	testValues := []int{0, 1, 2, 3}
	ch := NewCh(testValues...)
	actualValues := []int{}

	// Act
	ForEachCh(ch, func(val int) {
		actualValues = append(actualValues, val)
	})

	// Assert
	testLength := len(testValues)
	actualLength := len(actualValues)

	if testLength != actualLength {
		t.Fatalf("Expected length: %d, Actual length: %d", testLength, actualLength)
	}

	for i := 0; i < len(testValues); i++ {
		testVal := testValues[i]
		actualVal := actualValues[i]

		if testVal != actualVal {
			t.Errorf("Expected %d, Actual: %d", testVal, actualVal)
		}
	}
}

func TestMergeIntChs(t *testing.T) {
	// Arrange
	ch1 := make(chan int)
	ch2 := make(chan int)

	testValues := []int{1, 2, 3, 4, 5, 6}

	SendCh(ch1, testValues[:3]...)
	SendCh(ch2, testValues[3:]...)

	// Act
	ch := MergeChs(ch1, ch2)

	// Assert
	actualValues := ToSliceCh(ch)

	testLength := len(testValues)
	actualLength := len(actualValues)

	if testLength != actualLength {
		t.Fatalf("Expected length: %d, Actual length: %d", testLength, actualLength)
	}

	sort.Ints(actualValues)

	for i := 0; i < len(testValues); i++ {
		testVal := testValues[i]
		actualVal := actualValues[i]

		if testVal != actualVal {
			t.Errorf("Expected %d, Actual: %d", testVal, actualVal)
		}
	}
}
