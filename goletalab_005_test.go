package yaml

import "testing"

func TestGoletaYAML005(t *testing.T){
 opt:=ReferenceDirs("config","shared");d:=&Decoder{};_ = opt(d);if len(d.referenceDirs)!=2||d.referenceDirs[0]!="config"{t.Fatalf("dirs=%v",d.referenceDirs)}
}

func TestGoletaYAML005AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
 opt:=ReferenceDirs("fixtures");d:=&Decoder{};_ = opt(d);if len(d.referenceDirs)!=1||d.referenceDirs[0]!="fixtures"{t.Fatalf("dirs=%v",d.referenceDirs)}
}
