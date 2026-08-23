package yaml

import "testing"

func TestGoletaYAML015(t *testing.T){
 e:=&Encoder{};_ = UseSingleQuote(true)(e);if !e.singleQuote{t.Fatal("single quote preference was not enabled")}
}
