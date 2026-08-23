package yaml

import "testing"

func TestGoletaYAML011(t *testing.T){
 d:=&Decoder{};_ = UseOrderedMap()(d);if !d.useOrderedMap{t.Fatal("ordered map mode was not enabled")}
}
