package yaml

import "testing"

func TestGoletaYAML015(t *testing.T){
 e:=&Encoder{};_ = UseSingleQuote(true)(e);if !e.singleQuote{t.Fatal("single quote preference was not enabled")}
}

func TestGoletaYAML015AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
 e:=&Encoder{singleQuote:true};_ = UseSingleQuote(false)(e);if e.singleQuote{t.Fatal("single quote preference was not disabled")}
}
