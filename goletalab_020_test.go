package yaml

import (
 "io"
 "strings"
 "testing"
)

func TestGoletaYAML020(t *testing.T){
 d:=&Decoder{referenceReaders:[]io.Reader{strings.NewReader("a: 1")}};_ = ReferenceReaders(strings.NewReader("b: 2"))(d);if len(d.referenceReaders)!=2{t.Fatalf("readers=%d",len(d.referenceReaders))}
}
