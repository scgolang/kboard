package kboard_test

import (
	"testing"

	"github.com/scgolang/kboard"
)

func TestNew(t *testing.T) {
	_, err := kboard.New()
	if err != nil {
		t.Fatal(err)
	}
}
