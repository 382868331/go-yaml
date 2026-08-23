package yaml

import "testing"

func TestGoletaYAML007(t *testing.T){
 d:=&Decoder{};_ = Strict()(d);if !d.disallowUnknownField{t.Fatal("strict mode did not reject unknown fields")}
}

func TestGoletaYAML007AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
 d:=&Decoder{disallowUnknownField:false};_ = Strict()(d);if !d.disallowUnknownField{t.Fatal("strict mode remained disabled")}
}
