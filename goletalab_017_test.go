package yaml

import "testing"

func TestGoletaYAML017(t *testing.T){
 e:=&Encoder{};_ = WithSmartAnchor()(e);if !e.enableSmartAnchor{t.Fatal("smart anchor was not enabled")}
}

func TestGoletaYAML017AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
 e:=&Encoder{enableSmartAnchor:false};_ = WithSmartAnchor()(e);if !e.enableSmartAnchor{t.Fatal("smart anchor remained disabled")}
}
