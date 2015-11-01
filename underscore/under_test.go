package main

import (
	"testing"
)

func TestBasic(t *testing.T) {
	output := SimplifyUnderscore("a__bz_bear")
	expected := "a_bz_bear"
	if output != expected {
		t.Errorf("error: expected is %s, output is %s\n", expected, output)
	}
}

func TestStartUnderscore(t *testing.T) {
	output := SimplifyUnderscore("____b_tiger")
	expected := "_b_tiger"
	if output != expected {
		t.Errorf("error: expected is %s, output is %s\n", expected, output)
	}
}

func TestEndingUnderscore(t *testing.T) {
	output := SimplifyUnderscore("__b_____d_f_z____")
	expected := "_b_d_f_z_"
	if output != expected {
		t.Errorf("error: expected is %s, output %s \n", expected, output)
	}
}
