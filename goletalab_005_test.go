package yaml

import "testing"

func TestGoletaYAML005(t *testing.T){
 opt:=ReferenceDirs("config","shared");d:=&Decoder{};_ = opt(d);if len(d.referenceDirs)!=2||d.referenceDirs[0]!="config"{t.Fatalf("dirs=%v",d.referenceDirs)}
}
