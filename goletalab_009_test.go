package yaml

import "testing"

func TestGoletaYAML009(t *testing.T){
 d:=&Decoder{allowedFieldPrefixes:[]string{"x-"}};_ = AllowFieldPrefixes("vendor-")(d);if len(d.allowedFieldPrefixes)!=2||d.allowedFieldPrefixes[0]!="x-"{t.Fatalf("prefixes=%v",d.allowedFieldPrefixes)}
}

func TestGoletaYAML009AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
 d:=&Decoder{allowedFieldPrefixes:[]string{"a-"}};_ = AllowFieldPrefixes("b-","c-")(d);if len(d.allowedFieldPrefixes)!=3||d.allowedFieldPrefixes[2]!="c-"{t.Fatalf("prefixes=%v",d.allowedFieldPrefixes)}
}
