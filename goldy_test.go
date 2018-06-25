package goldy

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

var gc = EnvConfig(DefaultEnvName)

func Test_parseFlags(t *testing.T) {
	tests := []struct {
		Flags   string
		Want    map[Flag]bool
		WantErr string
	}{
		{Want: map[Flag]bool{}},
		{Flags: "update", Want: map[Flag]bool{FlagUpdate: true}},
		{Flags: "invalid", WantErr: `unknown flag: "invalid"`},
		{Flags: "update,invalid", WantErr: `unknown flag: "invalid"`},
		{Flags: "invalid,update", WantErr: `unknown flag: "invalid"`},
		{Flags: "update,diff", Want: map[Flag]bool{FlagUpdate: true, FlagDiff: true}},
		{Flags: "diff,update", Want: map[Flag]bool{FlagUpdate: true, FlagDiff: true}},
	}

	for _, test := range tests {
		got, err := parseFlags(test.Flags)
		if err == nil {
			if test.WantErr != "" {
				t.Errorf("got err=nil want=%s", test.WantErr)
			} else if !reflect.DeepEqual(got, test.Want) {
				t.Errorf("got=%#v want=%#v", got, test.Want)
			}
		} else {
			if test.WantErr == "" {
				t.Errorf("got err=%s want=nil", err)
			} else if !strings.Contains(err.Error(), test.WantErr) {
				t.Errorf("got err=%s want=%s", err, test.WantErr)
			}

		}
	}
}

func TestInputFixtures(t *testing.T) {
	tests := []struct {
		Dir  []string
		Want Fixtures
	}{
		{
			Dir:  []string{"in", "empty"},
			Want: Fixtures{},
		},
		{
			Dir: []string{"in", "flat"},
			Want: Fixtures{
				"test-fixtures/in/flat/a.txt": []byte("file a\n"),
				"test-fixtures/in/flat/b.txt": []byte("file b\n"),
			},
		},
		{
			Dir: []string{"in", "nested"},
			Want: Fixtures{
				"test-fixtures/in/nested/a.txt":   []byte("file a\n"),
				"test-fixtures/in/nested/b.txt":   []byte("file b\n"),
				"test-fixtures/in/nested/c/d.txt": []byte("file d\n"),
			},
		},
	}
	for _, test := range tests {
		got, err := gc.InputFixtures(test.Dir...)
		if err != nil {
			t.Errorf("%s: %s", test.Dir, err)
		} else if !reflect.DeepEqual(got, test.Want) {
			t.Errorf("%s: does not match fixture: got=%#v want=%#v", test.Dir, got, test.Want)
		} else if diff := got.Diff(test.Want); len(diff) > 0 {
			t.Errorf("%s: does not match fixture: got=%#v", test.Dir, diff)
		}
	}
}

func TestGoldenFixtures(t *testing.T) {
	// There is a large number of test cases that need to be checked here, so
	// we break them down in a few individual states a GoldenFixture and the
	// filesystem can be in and then test all combinations of them.
	states := []string{
		"base",
		"changed",
		"ignore",
		"missing",
		"unexpected",
	}

	buf := &bytes.Buffer{}
	combinations := 1 << uint(len(states))
	for i := 0; i < combinations; i++ {
		combo := []string{}
		for j, state := range states {
			if i&(1<<uint(j)) > 0 {
				combo = append(combo, state)
			}
		}

		test := strings.Join(combo, "_")
		if test == "" {
			test = "noop"
		}
		t.Run(test, func(t *testing.T) {
			tmpDir := filepath.Join(gc.Dir, "tmp", test)
			if err := os.RemoveAll(tmpDir); err != nil {
				t.Fatal(err)
			} else if err := os.MkdirAll(tmpDir, 0700); err != nil {
				t.Fatal(err)
			}
			defer os.RemoveAll(tmpDir)

			c := DefaultConfig()
			c.Dir = filepath.Dir(tmpDir)
			testGf := c.GoldenFixtures(filepath.Base(tmpDir))
			testGf.Flags = ""

			comboM := map[string]bool{}
			for _, state := range combo {
				comboM[state] = true
				switch state {
				case "base":
					for _, name := range []string{"base.txt", ".hidden.txt"} {
						data := []byte("data for: " + name)
						if !IsDotfile(name) {
							testGf.Add(data, name)
						}
						if err := ioutil.WriteFile(filepath.Join(tmpDir, name), data, 0600); err != nil {
							t.Fatal(err)
						}
					}
				case "ignore":
					testGf.IgnoreUnexpected = true
				case "missing":
					name := "missing.txt"
					testGf.Add([]byte("data for: "+name), name)
				case "unexpected":
					name := "unexpected.txt"
					data := []byte("data for: " + name)
					if err := ioutil.WriteFile(filepath.Join(tmpDir, name), data, 0600); err != nil {
						t.Fatal(err)
					}
				case "changed":
					name := "changed.txt"
					data := []byte("data for: " + name)
					if err := ioutil.WriteFile(filepath.Join(tmpDir, name), data, 0600); err != nil {
						t.Fatal(err)
					}
					testGf.Add(append([]byte("changed "), data...), name)
				default:
					panic("BUG")
				}
			}

			expectErr := (comboM["changed"] ||
				comboM["missing"] ||
				(comboM["unexpected"] && !comboM["ignore"]))

			gotErr := testGf.Test()
			if (gotErr != nil) != expectErr {
				t.Fatalf("gotErr=%v expectErr=%t", gotErr, expectErr)
			}

			testGf.Flags = "update"
			if err := testGf.Test(); err != nil {
				t.Fatalf("update error: %v", err)
			}
			testGf.Flags = ""
			if err := testGf.Test(); err != nil {
				t.Fatalf("re-test error: %v", err)
			}

			fmt.Fprintf(
				buf,
				"# (%d/%d) %s\n\n",
				i+1,
				combinations,
				strings.Replace(test, "_", "\\_", -1),
			)
			fmt.Fprintf(buf, "```\n")
			fmt.Fprintf(buf, "%v\n", gotErr)
			fmt.Fprintf(buf, "```\n")
		})
	}

	gf := gc.GoldenFixtures("out")
	gf.Add(buf.Bytes(), "golden_fixtures.md")
	if err := gf.Test(); err != nil {
		t.Fatal(err)
	}
}
