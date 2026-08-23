package yaml

import "testing"

func TestGoletaYAML003(t *testing.T){
	var dst map[string]any;if err:=Unmarshal(nil,&dst);err!=nil{t.Fatalf("empty document returned error: %v",err)}
}

func TestGoletaYAML003AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	var dst struct{ Name string `yaml:"name"` };if err:=Unmarshal([]byte{},&dst);err!=nil{t.Fatalf("zero-length document returned error: %v",err)}
}
