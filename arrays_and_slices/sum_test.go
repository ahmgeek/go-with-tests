package sum

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	want := 6

	if want != got {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}

}

func TestSumAllTails(t *testing.T){
	got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 4})
	want := []int{5, 13}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}
