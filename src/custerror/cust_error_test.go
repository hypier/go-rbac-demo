package custerror

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	var err = New("abc")

	fmt.Printf("%+v", err)

	fmt.Println()
}
