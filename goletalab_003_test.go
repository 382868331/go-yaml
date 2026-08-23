package yaml

import "testing"

func TestGoletaYAML003(t *testing.T){
	var dst map[string]any;if err:=Unmarshal(nil,&dst);err!=nil{t.Fatalf("empty document returned error: %v",err)}
}
