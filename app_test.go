package incrementor

import "testing"

func Test_Increment(t *testing.T) {
	x := 0
	Increment(&x)

	if x != 1 {
		t.Errorf("ooooops")
	}
}
