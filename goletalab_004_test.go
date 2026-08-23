package yaml

import "testing"

func TestGoletaYAML004(t *testing.T){
 opt:=ReferenceFiles("base.yml","extra.yml");d:=&Decoder{};if err:=opt(d);err!=nil{t.Fatal(err)};if len(d.referenceFiles)!=2||d.referenceFiles[1]!="extra.yml"{t.Fatalf("files=%v",d.referenceFiles)}
}
