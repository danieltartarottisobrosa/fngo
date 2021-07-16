package fngo

import (
	"sort"
	"testing"
)

func TestNewEmptyStringCh(t *testing.T) {
	// Act
	ch := NewStringCh()

	// Assert
	if ch == nil {
		t.Error("The channel is nil")
	}
}

func TestNewPopulatedStringCh(t *testing.T) {
	// Act
	values := []string{"a", "b", "c", "d"}

	// Act
	ch := NewStringCh(values...)

	// Assert
	for _, expected := range values {
		if actual := <-ch; actual != expected {
			t.Errorf("Expected: %s, Actual: %s", expected, actual)
		}
	}
}

func TestSendStrings(t *testing.T) {
	// Arrange
	ch := make(chan string)
	values := []string{"a", "b", "c", "d"}

	// Act
	SendStrings(ch, values...)

	// Assert
	for _, expected := range values {
		if actual := <-ch; actual != expected {
			t.Errorf("Expected: %s, Actual: %s", expected, actual)
		}
	}
}

func TestMapStringToString(t *testing.T) {
	// Arrange
	ch := make(chan string)
	expectedValues := []string{"Item a", "Item b", "Item c", "Item d"}

	concatItem := func(val string) string {
		return "Item " + val
	}

	SendStrings(ch, "a", "b", "c", "d")

	// Act
	ch = MapStringToString(ch, concatItem)

	// Assert
	for _, expected := range expectedValues {
		if actual := <-ch; actual != expected {
			t.Errorf("Expected: %s, Actual: %s", expected, actual)
		}
	}
}

func TestMapStringToInt(t *testing.T) {
	// Arrange
	strCh := make(chan string)
	expectedValues := []int{97, 98, 99, 100}

	toInt := func(val string) int {
		return int([]rune(val)[0])
	}

	SendStrings(strCh, "a", "b", "c", "d")

	// Act
	intCh := MapStringToInt(strCh, toInt)

	// Assert
	for _, expected := range expectedValues {
		if actual := <-intCh; actual != expected {
			t.Errorf("Expected: %d, Actual: %d", expected, actual)
		}
	}
}

func TestForEachString(t *testing.T) {
	// Arrange
	testValues := []string{"a", "b", "c", "d"}
	ch := NewStringCh(testValues...)
	actualValues := []string{}

	// Act
	ForEachString(ch, func(val string) {
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
			t.Errorf("Expected %s, Actual: %s", testVal, actualVal)
		}
	}
}

func TestMergeStringChs(t *testing.T) {
	// Arrange
	ch1 := make(chan string)
	ch2 := make(chan string)

	testValues := []string{"a", "b", "c", "d", "e", "f"}

	SendStrings(ch1, testValues[:3]...)
	SendStrings(ch2, testValues[3:]...)

	// Act
	ch := MergeStringChs(ch1, ch2)

	// Assert
	actualValues := ToStringSlice(ch)

	testLength := len(testValues)
	actualLength := len(actualValues)

	if testLength != actualLength {
		t.Fatalf("Expected length: %d, Actual length: %d", testLength, actualLength)
	}

	sort.Strings(actualValues)

	for i := 0; i < len(testValues); i++ {
		testVal := testValues[i]
		actualVal := actualValues[i]

		if testVal != actualVal {
			t.Errorf("Expected %s, Actual: %s", testVal, actualVal)
		}
	}
}
