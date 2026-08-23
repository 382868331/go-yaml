package yaml

import "testing"

func TestGoletaYAML018(t *testing.T){
 e:=&Encoder{};_ = UseLiteralStyleIfMultiline(true)(e);if !e.useLiteralStyleIfMultiline{t.Fatal("literal multiline mode was not enabled")}
}

func TestGoletaYAML018AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
 e:=&Encoder{useLiteralStyleIfMultiline:true};_ = UseLiteralStyleIfMultiline(false)(e);if e.useLiteralStyleIfMultiline{t.Fatal("literal multiline mode was not disabled")}
}
