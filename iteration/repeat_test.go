package iteration

import "testing"
import "fmt"

func TestRepeat(t *testing.T){
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B){
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	str := Repeat("wow", 5)
	fmt.Println(str)
	// Output: wowwowwowwowwow
}
