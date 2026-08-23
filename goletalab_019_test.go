package yaml

import "testing"

func TestGoletaYAML019(t *testing.T){
 e:=&Encoder{};_ = JSON()(e);if !e.isJSONStyle||!e.isFlowStyle{t.Fatalf("json=%v flow=%v",e.isJSONStyle,e.isFlowStyle)}
}
