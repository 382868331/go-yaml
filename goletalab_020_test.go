package yaml

import (
 "io"
 "strings"
 "testing"
)

func TestGoletaYAML020(t *testing.T){
 d:=&Decoder{referenceReaders:[]io.Reader{strings.NewReader("a: 1")}};_ = ReferenceReaders(strings.NewReader("b: 2"))(d);if len(d.referenceReaders)!=2{t.Fatalf("readers=%d",len(d.referenceReaders))}
}

func TestGoletaYAML020AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
 d:=&Decoder{referenceReaders:[]io.Reader{strings.NewReader("a: 1")}};_ = ReferenceReaders(strings.NewReader("b: 2"),strings.NewReader("c: 3"))(d);if len(d.referenceReaders)!=3{t.Fatalf("readers=%d",len(d.referenceReaders))}
}
