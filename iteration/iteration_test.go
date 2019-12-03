package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	char := "a"
	const count = 5
	repeated := Repeat(char, count)
	repeatedCount := strings.Count(repeated, char)
	expected := "aaaaa"
	expectedCount := count

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
	if repeatedCount != expectedCount {
		t.Errorf("expectedCount %q but got %q", expectedCount, repeatedCount)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeatedStr := Repeat("ex_", 5)
	fmt.Println(repeatedStr)
	fmt.Println(strings.Count(repeatedStr, "ex_"))
	// Output:
	// ex_ex_ex_ex_ex_
	// 5
}
