package yaml

import "testing"

func TestGoletaYAML009(t *testing.T){
 d:=&Decoder{allowedFieldPrefixes:[]string{"x-"}};_ = AllowFieldPrefixes("vendor-")(d);if len(d.allowedFieldPrefixes)!=2||d.allowedFieldPrefixes[0]!="x-"{t.Fatalf("prefixes=%v",d.allowedFieldPrefixes)}
}
