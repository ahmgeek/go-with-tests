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

func TestAllTails(t *testing.T){
	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := AllTails([]int{1,2}, []int{0,9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := AllTails([]int{}, []int{3, 4, 5})
		want :=[]int{0, 9}
		checkSums(t, got, want)
	})
}
