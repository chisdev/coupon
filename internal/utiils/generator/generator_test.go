package generator

import (
	"fmt"
	"testing"
)

func Test_Gencode(t *testing.T) {
	code, err := GenCode(10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(code) != 10 {
		t.Fatalf("expected code length 10, got %d", len(code))
	}
	fmt.Printf("generated code: %s\n", code)
}
