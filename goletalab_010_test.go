package yaml

import "testing"

func TestGoletaYAML010(t *testing.T){
 d:=&Decoder{};_ = AllowDuplicateMapKey()(d);if !d.allowDuplicateMapKey{t.Fatal("duplicate keys were not enabled")}
}
