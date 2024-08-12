package data

import (
	"errors"
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

func TestGetVersionCorrectVersion(t *testing.T) {
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

func TestGetVersionInvalidVersion(t *testing.T) {
	src := NewPistonDataSource()

	v, err := src.GetVersion("1.2112312")
	if v != nil {
		t.Fatalf("Expected nil pointer but got %v", v)
	}

	target := &UnknownVersionError{}
	if !errors.As(err, &target) {
		t.Fatalf("Expected '%s' but got '%s'", target, err)
	}

}
