package yaml

import "testing"

func TestGoletaYAML007(t *testing.T){
 d:=&Decoder{};_ = Strict()(d);if !d.disallowUnknownField{t.Fatal("strict mode did not reject unknown fields")}
}
