package fngo

import (
	"fmt"
	"testing"
)

func TestNewEmptyAnyCh(t *testing.T) {
	// Act
	ch := NewAnyCh()

	// Assert
	if ch == nil {
		t.Error("The channel is nil")
	}
}

func TestNewPopulatedAnyCh(t *testing.T) {
	// Act
	values := []Any{0, "a", 2, "c"}

	// Act
	ch := NewAnyCh(values...)

	// Assert
	for _, expected := range values {
		if actual := <-ch; actual != expected {
			t.Errorf("Expected: %v, Actual: %v", expected, actual)
		}
	}
}

func TestSendAny(t *testing.T) {
	// Arrange
	ch := make(chan Any)
	values := []Any{0, "a", 2, "c"}

	// Act
	SendAny(ch, values...)

	// Assert
	for _, expected := range values {
		if actual := <-ch; actual != expected {
			t.Errorf("Expected: %v, Actual: %v", expected, actual)
		}
	}
}

func TestMapAnyToAny(t *testing.T) {
	// Arrange
	ch := make(chan Any)
	expectedValues := []string{"0", "1", "2", "3"}

	change := func(val Any) Any {
		return fmt.Sprintf("%d", val)
	}

	SendAny(ch, 0, 1, 2, 3)

	// Act
	ch = MapAnyToAny(ch, change)

	// Assert
	for _, expected := range expectedValues {
		if actual := <-ch; actual != expected {
			t.Errorf("Expected: %v, Actual: %v", expected, actual)
		}
	}
}

func TestMapAnyToString(t *testing.T) {
	// Arrange
	anyCh := make(chan Any)
	expectedValues := []string{"Valor 0", "Valor 1", "Valor 2", "Valor 3"}

	toString := func(val Any) string {
		return fmt.Sprintf("Valor %d", val)
	}

	SendAny(anyCh, 0, 1, 2, 3)

	// Act
	strCh := MapAnyToString(anyCh, toString)

	// Assert
	for _, expected := range expectedValues {
		if actual := <-strCh; actual != expected {
			t.Errorf("Expected: %s, Actual: %s", expected, actual)
		}
	}
}

func TestForEachAny(t *testing.T) {
	// Arrange
	testValues := []int{0, 1, 2, 3}
	ch := NewIntCh(testValues...)
	actualValues := []int{}

	// Act
	ForEachInt(ch, func(val int) {
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
