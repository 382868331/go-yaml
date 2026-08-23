package yaml

import "testing"

func TestGoletaYAML012(t *testing.T){
 d:=&Decoder{};_ = UseJSONUnmarshaler()(d);if !d.useJSONUnmarshaler{t.Fatal("JSON unmarshaler fallback was not enabled")}
}
