package yaml

import "testing"

func TestGoletaYAML010(t *testing.T){
 d:=&Decoder{};_ = AllowDuplicateMapKey()(d);if !d.allowDuplicateMapKey{t.Fatal("duplicate keys were not enabled")}
}

func TestGoletaYAML010AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
 d:=&Decoder{allowDuplicateMapKey:false};_ = AllowDuplicateMapKey()(d);if !d.allowDuplicateMapKey{t.Fatal("duplicate key option remained disabled")}
}
