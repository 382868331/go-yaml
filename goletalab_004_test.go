package yaml

import "testing"

func TestGoletaYAML004(t *testing.T){
 opt:=ReferenceFiles("base.yml","extra.yml");d:=&Decoder{};if err:=opt(d);err!=nil{t.Fatal(err)};if len(d.referenceFiles)!=2||d.referenceFiles[1]!="extra.yml"{t.Fatalf("files=%v",d.referenceFiles)}
}

func TestGoletaYAML004AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
 opt:=ReferenceFiles("single.yml");d:=&Decoder{};_ = opt(d);if len(d.referenceFiles)!=1||d.referenceFiles[0]!="single.yml"{t.Fatalf("files=%v",d.referenceFiles)}
}
