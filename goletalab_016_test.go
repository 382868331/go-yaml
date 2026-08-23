package yaml

import "testing"

func TestGoletaYAML016(t *testing.T){
 e:=&Encoder{};_ = Flow(true)(e);if !e.isFlowStyle{t.Fatal("flow style was not enabled")}
}
