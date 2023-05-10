package scratchpad

import "testing"

func Test_ReadPeopleCSV(t *testing.T) {
	exp := 1000

	// call method ReadPeopleCSV
	people, _ := ReadPeopleCSV()

	act := len(people)

	for _, p := range people {
		t.Log(p.String())
	}

	if act != exp {
		t.Errorf("expected: %d got: %d", exp, act)
	}
}
