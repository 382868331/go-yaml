package yaml

import "testing"

func TestGoletaYAML002(t *testing.T){
	got,err:=Marshal(map[string]int{"count":1});if err!=nil{t.Fatal(err)};if string(got)!="count: 1\n"{t.Fatalf("got %q",got)}
}

func TestGoletaYAML002AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	got,err:=Marshal([]string{"a","b"});if err!=nil{t.Fatal(err)};if len(got)==0||got[len(got)-1]!='\n'{t.Fatalf("missing document terminator: %q",got)}
}
