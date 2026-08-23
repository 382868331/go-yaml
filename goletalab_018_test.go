package yaml

import "testing"

func TestGoletaYAML018(t *testing.T){
 e:=&Encoder{};_ = UseLiteralStyleIfMultiline(true)(e);if !e.useLiteralStyleIfMultiline{t.Fatal("literal multiline mode was not enabled")}
}
