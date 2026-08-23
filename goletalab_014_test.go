package yaml

import "testing"

func TestGoletaYAML014(t *testing.T){
 e:=&Encoder{};_ = IndentSequence(true)(e);if !e.indentSequence{t.Fatal("sequence indentation was not enabled")}
}

func TestGoletaYAML014AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
 e:=&Encoder{indentSequence:true};_ = IndentSequence(false)(e);if e.indentSequence{t.Fatal("sequence indentation was not disabled")}
}
