package yaml

import "testing"

func TestGoletaYAML002(t *testing.T){
	got,err:=Marshal(map[string]int{"count":1});if err!=nil{t.Fatal(err)};if string(got)!="count: 1\n"{t.Fatalf("got %q",got)}
}
