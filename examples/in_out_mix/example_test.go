package example

import (
	"crypto/sha1"
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/felixge/goldy"
)

var gc = goldy.DefaultConfig()

func TestSha1(t *testing.T) {
	input, err := gc.InputFixtures()
	if err != nil {
		t.Fatal(err)
	}
	gf := gc.GoldenFixtures()
	gf.IgnoreUnexpected = true
	for path, data := range input {
		if strings.HasSuffix(path, ".golden.txt") {
			continue
		}
		outPath := strings.Replace(filepath.Base(path), "input", "golden", -1)
		gf.Add([]byte(fmt.Sprintf("%x", sha1.Sum(data))), outPath)
	}
	if err := gf.Test(); err != nil {
		t.Fatal(err)
	}
}
