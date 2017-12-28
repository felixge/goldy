package example

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/jpeg"
	"image/png"
	"path/filepath"
	"testing"

	"github.com/felixge/goldy"
)

var gc = goldy.DefaultConfig()

func TestMarshalIndent(t *testing.T) {
	tests := []struct {
		Name  string
		Input string
	}{
		{
			"small_object", "{}",
		},
	}
	gf := gc.GoldenFixtures("out", "marshal_indent")
	for _, test := range tests {
		b := &bytes.Buffer{}
		if err := json.Indent(b, []byte(test.Input), "", "  "); err != nil {
			b.WriteString(err.Error())
		}
		gf.Add(b.Bytes(), test.Name+".json")
	}
	if err := gf.Test(); err != nil {
		t.Fatal(err)
	}
}

func TestGradient(t *testing.T) {
	tests := []struct {
		Width  int
		Height int
	}{
		{256, 256},
		{100, 100},
		{50, 100},
		{100, 50},
	}

	gf := gc.GoldenFixtures("out", "gradient")
	for _, test := range tests {
		b := &bytes.Buffer{}
		png.Encode(b, Gradient(test.Width, test.Height))
		name := fmt.Sprintf("%dx%x.png", test.Width, test.Height)
		gf.Add(b.Bytes(), name)
	}
	if err := gf.Test(); err != nil {
		t.Fatal(err)
	}
}

func TestRedOnly(t *testing.T) {
	in, err := gc.InputFixtures("in")
	if err != nil {
		t.Fatal(err)
	}
	out := gc.GoldenFixtures("out", "red-only")
	for path, data := range in {
		img, err := jpeg.Decode(bytes.NewReader(data))
		if err != nil {
			t.Fatal(err)
		}
		b := &bytes.Buffer{}
		png.Encode(b, RedOnly(img))
		out.Add(b.Bytes(), filepath.Base(path))
	}
	if err := out.Test(); err != nil {
		t.Fatal(err)
	}
}
