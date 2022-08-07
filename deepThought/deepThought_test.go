package deepThought

import "testing"

func TestAskPostive(t *testing.T) {
	argumentToTestPositve_1 :="life"
	argumentToTestPositve_2 :="universe"
	argumentToTestPositve_3 :="everything"
	v1, v2 := Ask(argumentToTestPositve_1,argumentToTestPositve_2,argumentToTestPositve_3)
	if  v1 != 42 || v2 != "change" {
		t.Errorf("Expected 42 and \"change\", but got %v and %v.\n", v1, v2)
	}
}

func TestAskNegative(t *testing.T) {
        argumentToTestPositve_1 :="me"
        argumentToTestPositve_2 :="myself"
        argumentToTestPositve_3 :="I"
        v1, v2 := Ask(argumentToTestPositve_1,argumentToTestPositve_2,argumentToTestPositve_3)
        if  v1 != 0 || v2 != "nothing" {
                t.Errorf("Expected 0 and \"nothing\", but got %v and %v.\n", v1, v2)
        }
}
