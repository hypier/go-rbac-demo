package custerror

import (
	"errors"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	var err1 = fmt.Errorf("err1 %w", errors.New("用户已存在"))
	var err2 = fmt.Errorf("err2 %w", err1)

	var err3 = Errorf("err3", errors.New("用户已存在"))
	var err4 = Errorf("err4", err3)
	//PrintError(err)

	fmt.Printf("%q", err2)
	if errors.Is(err2, err1) {
		fmt.Println("Yes")
	}
	fmt.Println()

	fmt.Printf("%q", err4)
	fmt.Println()

	if errors.Is(err4, err3) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	fmt.Println()

}
