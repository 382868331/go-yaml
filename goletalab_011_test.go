package yaml

import "testing"

func TestGoletaYAML011(t *testing.T){
 d:=&Decoder{};_ = UseOrderedMap()(d);if !d.useOrderedMap{t.Fatal("ordered map mode was not enabled")}
}

func TestGoletaYAML011AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
 d:=&Decoder{useOrderedMap:false};_ = UseOrderedMap()(d);if !d.useOrderedMap{t.Fatal("ordered map mode remained disabled")}
}
