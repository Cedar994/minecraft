package download

import (
	"io/ioutil"
	"testing"
)

func TestUnzip(t *testing.T) {
	b, err := ioutil.ReadFile("1.0.json")
	if err != nil {
		t.Fatal(err)
	}
	l, err := Newlibraries(b)
	if err != nil {
		t.Fatal(err)
	}
	err = l.Unzip("", 5)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(l)
}