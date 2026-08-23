package yaml

import "testing"

func TestGoletaYAML001(t *testing.T){
	s:=MapSlice{{Key:"name",Value:"alice"},{Key:"age",Value:7}};got:=s.ToMap();if got["name"]!="alice"||got["age"]!=7{t.Fatalf("unexpected map: %#v",got)}
}

func TestGoletaYAML001AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	s:=MapSlice{{Key:"enabled",Value:true}};got:=s.ToMap();if got["enabled"]!=true{t.Fatalf("unexpected map: %#v",got)}
}
