package incrementor

import "testing"

func Test_Increment(t *testing.T) {
	start := 0
	expected := 1
	Increment(&start)

	if start != expected {
		t.Errorf("Increment(&start) = ?, wanted %d.\n", expected)
	}
}

func Test_Decrement(t *testing.T){
	start := 1
	expected := 0
	Decrement(&start)

	if start != expected {
		t.Errorf("Decrement(&start) = ?, wanted %d.\n", expected)
	}
}




