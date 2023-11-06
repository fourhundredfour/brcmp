package comparison

import (
	"fmt"
	"strings"
)

func CompareBinaries(firstBinary []byte, secondBinary []byte) string {
	var diff strings.Builder
	for i := 0; i < len(firstBinary) && i < len(secondBinary); i++ {
		if firstBinary[i] != secondBinary[i] {
			diff.WriteString(fmt.Sprintf("OFFSET: %d | FIRST BINARY: %x | SECOND BINARY: %x\n", i, firstBinary[i], secondBinary[i]))
		}
	}
	return diff.String()
}
