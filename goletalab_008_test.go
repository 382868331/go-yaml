package yaml

import "testing"

func TestGoletaYAML008(t *testing.T){
 d:=&Decoder{};_ = DisallowUnknownField()(d);if !d.disallowUnknownField{t.Fatal("unknown-field check was not enabled")}
}
