package yaml

import "testing"

func TestGoletaYAML006(t *testing.T){
 d:=&Decoder{};_ = RecursiveDir(true)(d);if !d.isRecursiveDir{t.Fatal("recursive mode was not enabled")}
}

func TestGoletaYAML006AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
 d:=&Decoder{isRecursiveDir:true};_ = RecursiveDir(false)(d);if d.isRecursiveDir{t.Fatal("recursive mode was not disabled")}
}
