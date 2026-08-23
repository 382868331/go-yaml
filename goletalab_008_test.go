package yaml

import "testing"

func TestGoletaYAML008(t *testing.T){
 d:=&Decoder{};_ = DisallowUnknownField()(d);if !d.disallowUnknownField{t.Fatal("unknown-field check was not enabled")}
}

func TestGoletaYAML008AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
 d:=&Decoder{disallowUnknownField:false};_ = DisallowUnknownField()(d);if !d.disallowUnknownField{t.Fatal("unknown-field check remained disabled")}
}
