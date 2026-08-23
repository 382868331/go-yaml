package yaml

import "testing"

func TestGoletaYAML006(t *testing.T){
 d:=&Decoder{};_ = RecursiveDir(true)(d);if !d.isRecursiveDir{t.Fatal("recursive mode was not enabled")}
}
