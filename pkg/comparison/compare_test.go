package comparison

import "testing"

func TestCompareBinariesShouldReturnDifferenceBetweenTwoBinaries(t *testing.T) {
	firstBinary := []byte("Hello!")
	secondBinary := []byte("Hello.")
	expectedOutput := "OFFSET: 5 | FIRST BINARY: 21 | SECOND BINARY: 2e\n"
	if output := CompareBinaries(firstBinary, secondBinary); expectedOutput != output {
		t.Errorf("CompareBinaries doesnt match the expected output. Expected: %v  got: %v", expectedOutput, output)
	}

}
