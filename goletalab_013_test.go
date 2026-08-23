package yaml

import "testing"

func TestGoletaYAML013(t *testing.T){
 e:=&Encoder{};_ = Indent(4)(e);if e.indentNum!=4{t.Fatalf("indent=%d",e.indentNum)}
}
