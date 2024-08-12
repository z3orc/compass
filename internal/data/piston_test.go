package data

import (
	"testing"

	"github.com/z3orc/compass/internal/util"
)

func TestNewPistonDataSource(t *testing.T) {
	src := NewPistonDataSource()

	if src == nil {
		t.Fatal("Expected *PistonDataSource but got nil")
	}

	if !util.CheckUrl(src.url) {
		t.Fatal("Expected datasource url but got invalid")
	}
}

func TestGetVersion(t *testing.T) {
	src := NewPistonDataSource()

	v, err := src.GetVersion("1.21")
	if v == nil {
		t.Fatal("Expected *Version but got nil")
	} else if err != nil {
		t.Fatalf("Expected err to be nil but got %v", err)
	}

	valid := v.IsValid()
	if valid != nil {
		t.Fatalf("Expected valid Version but got %e", valid)
	}
}
