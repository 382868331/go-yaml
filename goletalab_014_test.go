package yaml

import "testing"

func TestGoletaYAML014(t *testing.T){
 e:=&Encoder{};_ = IndentSequence(true)(e);if !e.indentSequence{t.Fatal("sequence indentation was not enabled")}
}
