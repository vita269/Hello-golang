package main
import "testing"

gunc TestHello(t *testing.T) {
want := "Hello Go"

got := hello()

if want != got {
t.Fatalf("want %s, got %s/n", want, got)
}
}
