package yaml

import "testing"

func TestGoletaYAML017(t *testing.T){
 e:=&Encoder{};_ = WithSmartAnchor()(e);if !e.enableSmartAnchor{t.Fatal("smart anchor was not enabled")}
}
