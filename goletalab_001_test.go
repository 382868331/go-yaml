package yaml

import "testing"

func TestGoletaYAML001(t *testing.T){
	s:=MapSlice{{Key:"name",Value:"alice"},{Key:"age",Value:7}};got:=s.ToMap();if got["name"]!="alice"||got["age"]!=7{t.Fatalf("unexpected map: %#v",got)}
}
