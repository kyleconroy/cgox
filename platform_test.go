package main

import (
	"reflect"
	"testing"
)

func TestSupportedPlatforms(t *testing.T) {
	var ps []Platform

	ps = SupportedPlatforms("go1.0")
	if !reflect.DeepEqual(ps, Platforms_1_0) {
		t.Fatalf("bad: %#v", ps)
	}

	ps = SupportedPlatforms("go1.1")
	if !reflect.DeepEqual(ps, Platforms_1_1) {
		t.Fatalf("bad: %#v", ps)
	}
}
