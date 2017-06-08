package goldy

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

var gc = EnvConfig(DefaultEnvName)

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
	tmpDir := filepath.Join("test-fixtures", "out", "tmp")
	if err := os.RemoveAll(tmpDir); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// There is a bit inception going on here because we're using the code we're
	// testing to test itself. Since compilers consider self hosting a major
	// achievement, we'll consider test code that is testing itself to be great
	// here as well : p.
	gf := gc.GoldenFixtures("out", "golden_fixtures")
	for _, test := range []string{
		"compare_missing",
		"update_missing",
		"update_changed_added",
		"compare_good",
		"compare_added",
		"compare_modified",
	} {
		func(test string) {
			testGf := &GoldenFixtures{
				Dir:      tmpDir,
				Hint:     "EXAMPLE_HINT",
				Fixtures: Fixtures{},
				Exclude:  IsDotfile,
			}

			testGf.Add([]byte("data_a"), "file_a")
			testGf.Add([]byte("data_b"), "file_b")
			testGf.Add([]byte("data_d"), "dir_c", "file_d")

			var err error
			switch test {
			case "compare_missing":
				// Initially our tmpDir doesn't exist, so we expect Test to produce
				// errors for the missing files.
			case "update_missing":
				// Next do an update that creates the missing files. This should not
				// produce an error.
				testGf.Update = true
			case "update_changed_added":
				// Now create a new file and modify another one, then perform an update
				// again. This should remove the new file and restore the modified one.
				// This should not produce an error.
				paths := []string{
					filepath.Join(testGf.Dir, "remove_me"),
					testGf.Fixtures.Paths()[0],
				}
				for _, path := range paths {
					if err := ioutil.WriteFile(path, []byte(path), 0600); err != nil {
						t.Fatal(err)
					}
				}
				testGf.Update = true
			case "compare_good":
				// Now we perform a compare to make sure the previous update worked.
				// This should not produce an error.
			case "compare_added":
				// Here we add two new files. One of them should be ignored because it
				// starts with a dot, the other should produce an error.
				for _, file := range []string{"new", ".hidden"} {
					path := filepath.Join(testGf.Dir, file)
					if err := ioutil.WriteFile(path, []byte(path), 0600); err != nil {
						t.Fatal(err)
					}
					// Clean up before the next step.
					defer os.Remove(path)
				}
			case "compare_modified":
				// Here we modify the first file in the fixture. We expect this to
				// produce an error message telling us that this file doesn't match.
				path := testGf.Fixtures.Paths()[0]
				if err := ioutil.WriteFile(path, []byte(test), 0600); err != nil {
					t.Fatal(err)
				}
				// Clean up before the next step.
				defer ioutil.WriteFile(path, testGf.Fixtures[path], 0600)
			default:
				panic("unknown test: " + test)
			}

			err = testGf.Test()
			gf.Add([]byte(fmt.Sprintf("%v", err)), test+".txt")
		}(test)
	}

	if err := gf.Test(); err != nil {
		t.Fatal(err)
	}
}
