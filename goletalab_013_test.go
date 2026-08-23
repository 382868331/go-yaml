package yaml

import "testing"

func TestGoletaYAML013(t *testing.T){
 e:=&Encoder{};_ = Indent(4)(e);if e.indentNum!=4{t.Fatalf("indent=%d",e.indentNum)}
}

func TestGoletaYAML013AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
 e:=&Encoder{};_ = Indent(2)(e);if e.indentNum!=2{t.Fatalf("indent=%d",e.indentNum)}
}
