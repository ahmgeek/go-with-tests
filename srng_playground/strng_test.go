package strng

import(
	"testing"
)

func TestContain(t *testing.T){
	result := Contain("seafood", "foo")
	expected := true

	if result != true {
		t.Errorf("expected '%t' got '%t'", expected, result)
	}
}
