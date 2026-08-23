package yaml

import "testing"

func TestGoletaYAML019(t *testing.T){
 e:=&Encoder{};_ = JSON()(e);if !e.isJSONStyle||!e.isFlowStyle{t.Fatalf("json=%v flow=%v",e.isJSONStyle,e.isFlowStyle)}
}

func TestGoletaYAML019AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
 e:=&Encoder{isJSONStyle:false,isFlowStyle:false};_ = JSON()(e);if !e.isJSONStyle||!e.isFlowStyle{t.Fatalf("json=%v flow=%v",e.isJSONStyle,e.isFlowStyle)}
}
