package yaml

import "testing"

func TestGoletaYAML012(t *testing.T){
 d:=&Decoder{};_ = UseJSONUnmarshaler()(d);if !d.useJSONUnmarshaler{t.Fatal("JSON unmarshaler fallback was not enabled")}
}

func TestGoletaYAML012AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
 d:=&Decoder{useJSONUnmarshaler:false};_ = UseJSONUnmarshaler()(d);if !d.useJSONUnmarshaler{t.Fatal("JSON fallback remained disabled")}
}
