package yaml

import "testing"

func TestGoletaYAML016(t *testing.T){
 e:=&Encoder{};_ = Flow(true)(e);if !e.isFlowStyle{t.Fatal("flow style was not enabled")}
}

func TestGoletaYAML016AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
 e:=&Encoder{isFlowStyle:true};_ = Flow(false)(e);if e.isFlowStyle{t.Fatal("flow style was not disabled")}
}
