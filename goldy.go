package goldy

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pmezard/go-difflib/difflib"
)

type Flag string

const (
	// FlagUpdate causes goldy to update any modified or missing fixtures and
	// to delete any fixtures that were removed.
	FlagUpdate Flag = "update"
	// FlagDiff causes goldly to print a diff for mismatching fixtures.
	FlagDiff Flag = "diff"
)

func parseFlags(flags string) (map[Flag]bool, error) {
	r := map[Flag]bool{}
	if flags == "" {
		return r, nil
	}
	for _, flag := range strings.Split(flags, ",") {
		switch f := Flag(flag); f {
		case FlagUpdate, FlagDiff:
			r[f] = true
		default:
			return nil, fmt.Errorf("unknown flag: %q", flag)
		}
	}
	return r, nil
}

const (
	// DefaultEnvName is the environment variable name used by DefaultConfig.
	DefaultEnvName = "GOLDY"
)

// DefaultConfig is a wrapper for EnvConfig(DefaultEnvName). It is the
// recommended method for integrating goldy. See package docs for more
// information.
func DefaultConfig() Config {
	return EnvConfig(DefaultEnvName)
}

// EnvConfig returns a new Config that uses the env variable with the given
// name name to determine if golden fixtures should be updated or compared
// when calling Test on them.
func EnvConfig(name string) Config {
	return Config{
		Flags: os.Getenv(name),
		Hint:  name + "=update go test",
	}.WithDefaults()
}

// FlagConfig returns a new Config that registers a flag with the given name
// with the global flag package to determine golden fixtures should be updated
// or compared when calling Test on them. This method exists because the
// author's hate for dogma exceeds his hate for global state. That being said,
// he'll still shake his head if you end up using this method.
func FlagConfig(name string) *Config {
	c := (Config{Hint: "go test -" + name}).WithDefaults()
	flag.StringVar(&c.Flags, name, "", "Goldy flags: update, diff")
	return &c
}

// Config allows you to customize your goldy integration. You're probably
// better off using DefaultConfig() instead.
type Config struct {
	// Dir is the base dir used for loading input or golden fixtures. Set to
	// "test-fixtures" by WithDefaults.
	Dir string
	// Flags is a comma separated string containing flags that control goldy's
	// behavior.
	Flags string
	// Hint is a message displayed when golden fixtures fail comparison. It's
	// intended to tell the user how to automatically update the fixtures.
	Hint string
	// IgnoreUnexpected is inherited by all GoldenFixtures created from this
	// Config.
	IgnoreUnexpected bool
	// Exclude is called for every file when loading input or golden fixtures and
	// allows to exclude it by returning false. Set to IsDotfile by WithDefaults.
	Exclude func(path string) bool
}

// WithDefaults returns a a copy of c that replaces zero values with default
// values as documented on the Config struct.
func (c Config) WithDefaults() Config {
	if c.Dir == "" {
		c.Dir = "test-fixtures"
	}
	if c.Exclude == nil {
		c.Exclude = IsDotfile
	}
	return c
}

// GoldenFixtures returns a new GoldenFixtures instance pointing to the given
// path inside c.Dir.
func (c Config) GoldenFixtures(path ...string) *GoldenFixtures {
	return &GoldenFixtures{
		Dir:              filepath.Join(append([]string{c.Dir}, path...)...),
		Fixtures:         Fixtures{},
		Flags:            c.Flags,
		Hint:             c.Hint,
		IgnoreUnexpected: c.IgnoreUnexpected,
		Exclude:          IsDotfile,
	}
}

// GoldenFixture returns an error if the fixture at the given path does not
// match the given data.
func (c Config) GoldenFixture(data []byte, path ...string) error {
	gf := c.GoldenFixtures(path...)
	gf.IgnoreUnexpected = true
	gf.Add(data)
	return gf.Test()
}

// InputFixtures loads Fixtures from the given path inside of c.Dir.
func (c Config) InputFixtures(path ...string) (Fixtures, error) {
	dir := filepath.Join(append([]string{c.Dir}, path...)...)
	return Load(dir, c.Exclude)
}

// InputFixture returns the data for the fixture at the given path or an error.
func (c Config) InputFixture(path ...string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(append([]string{c.Dir}, path...)...))
}

// GoldenFixtures is a set of fixture files that can be compared with files on
// disk, or used to update them. Most users probably want to create instances
// using a Config instance rather than initializing this struct directly.
type GoldenFixtures struct {
	// Dir is the disk directory to compare/update.
	Dir string
	// Fixtures is the in-memory set of fixture files created by the test.
	Fixtures Fixtures
	// Flags controls goldy's behavior.
	Flags string
	// Hint is displayed when comparing the in-memory fixtures with those on
	// disk shows differences.
	Hint string
	// IgnoreUnexpected determines if unexpected files found in Dir are ignored
	// when running Test().
	IgnoreUnexpected bool
	// Exclude allows to exclude on-disk files from the comparison/update.
	Exclude func(path string) bool
}

// Add adds a new fixture file with the given path relative to gf.Dir and data
// for being compared or updated when calling Test.
func (gf *GoldenFixtures) Add(data []byte, path ...string) {
	gf.Fixtures.Add(data, append([]string{gf.Dir}, path...)...)
}

// Diff returns the diff between gf.Fixtures and the golden fixtures from
// gf.Dir or an error.
func (gf *GoldenFixtures) Diff() (Diff, error) {
	want, err := Load(gf.Dir, gf.Exclude)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to load golden fixtures: %s", err)
	}
	diff := gf.Fixtures.Diff(want)
	if !gf.IgnoreUnexpected {
		return diff, nil
	}
	var newDiff Diff
	for _, d := range diff {
		if d.Kind != DiffUnexpected {
			newDiff = append(newDiff, d)
		}
	}
	return newDiff, nil
}

// Test returns an error if the comparison between gf.Fixtures and the golden
// fixtures in gf.Dir produced a diff. Or if gf.Flags[FlagUpdate] is true, it
// instead overwrites the golden fixtures in gf.Dir with those in gf.Fixtures
// and only returns an error if the update fails.
func (gf *GoldenFixtures) Test() error {
	flags, err := parseFlags(gf.Flags)
	if err != nil {
		return err
	}

	diff, err := gf.Diff()
	if err != nil {
		return err
	}

	if flags[FlagUpdate] {
		return gf.update(diff)
	} else {
		return gf.compare(diff, flags[FlagDiff])
	}
}

func (gf *GoldenFixtures) update(diff Diff) error {
	msg := make([]string, 0, len(diff))
	for _, d := range diff {
		switch d.Kind {
		case DiffUnexpected:
			if err := os.Remove(d.Path); err != nil {
				msg = append(msg, fmt.Sprintf("could not remove: %s: %s", d.Path, err))
			}
		case DiffMissing, DiffChanged:
			dir := filepath.Dir(d.Path)
			if err := os.MkdirAll(dir, 0700); err != nil {
				msg = append(msg, fmt.Sprintf("could not mkdir: %s: %s", dir, err))
			} else if err := ioutil.WriteFile(d.Path, gf.Fixtures[d.Path], 0600); err != nil {
				msg = append(msg, fmt.Sprintf("could not write: %s: %s", d.Path, err))
			}
		}
	}
	if len(msg) > 0 {
		return fmt.Errorf(
			"%d errors:\n%s",
			len(msg),
			strings.Join(msg, "\n"),
		)
	}
	return nil
}

func (gf *GoldenFixtures) compare(diff Diff, diffFlag bool) error {
	if len(diff) == 0 {
		return nil
	}
	var msg []string
	for _, d := range diff {
		switch d.Kind {
		case DiffUnexpected:
			msg = append(msg, fmt.Sprintf("unexpected file: %s", d.Path))
		case DiffMissing:
			msg = append(msg, fmt.Sprintf("missing file: %s", d.Path))
		case DiffChanged:
			msg = append(msg, fmt.Sprintf("changed file: %s", d.Path))
			if diffFlag {
				msg = append(msg, textDiff(d.A, d.B))
			}
		}
	}
	return fmt.Errorf(
		"%d errors:\n%s\n\nrun `%s` to automatically update all files above",
		len(diff),
		strings.Join(msg, "\n"),
		gf.Hint,
	)
}

func textDiff(a, b []byte) string {
	diff := difflib.UnifiedDiff{
		A:       difflib.SplitLines(string(a)),
		B:       difflib.SplitLines(string(b)),
		Context: 3,
	}
	text, _ := difflib.GetUnifiedDiffString(diff)
	return indent(strings.TrimRight(text, "\n"))
}

func indent(s string) string {
	return "  " + strings.Replace(s, "\n", "\n  ", -1)
}

// Load loads a Fixtures from the given path. The exclude func is called for every
// file and allows excluding paths by returning false.
func Load(path string, exclude func(path string) bool) (Fixtures, error) {
	s := Fixtures{}
	return s, filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		} else if info.IsDir() || exclude(path) {
			return nil
		} else if data, err := ioutil.ReadFile(path); err != nil {
			return err
		} else {
			s[path] = data
			return nil
		}
	})
}

// IsDotfile returns true if path starts with a ".". This is useful for
// excluding hidden files on Unix / Linux, e.g. vim undo files.
func IsDotfile(path string) bool {
	return strings.HasPrefix(filepath.Base(path), ".")
}

// Fixtures maps file paths to their file contents.
type Fixtures map[string][]byte

// Add adds the given path and file contents or panics if the path already
// exists.
func (f Fixtures) Add(data []byte, path ...string) {
	key := filepath.Join(path...)
	if _, ok := f[key]; ok {
		panic("set already has path: " + key)
	}
	f[key] = data
}

// Diff compares set a with set b and returns the diff. If a and b are equal,
// the returned len(diff) is 0. See Fixtures.Diff for for more details. The
// main caller of this func is GoldenFixtures.Diff, in that context b is the
// existing fixture on disk, and a is the results from the test and we want
// to show the change from b to a.
func (a Fixtures) Diff(b Fixtures) Diff {
	var diff Diff
	// First pass through a finds all paths that exist in a but not b or that
	// exist in both but hold different data.
	for aPath, aData := range a {
		d := &FileDiff{
			Path: aPath,
			B:    aData,
		}
		if bData, ok := b[aPath]; !ok {
			d.Kind = DiffMissing
			diff = append(diff, d)
		} else if !bytes.Equal(aData, bData) {
			d.Kind = DiffChanged
			d.A = bData
			diff = append(diff, d)
		}
	}
	// Second pass through b is needed for finding paths that exist in b, but not
	// in a.
	for bPath, _ := range b {
		if bData, ok := a[bPath]; !ok {
			diff = append(diff, &FileDiff{
				Path: bPath,
				A:    bData,
				Kind: DiffUnexpected,
			})
		}
	}

	sort.Slice(diff, func(i, j int) bool {
		return diff[i].Path < diff[j].Path
	})
	return diff
}

// Paths returns all path keys from f in ascending byte order.
func (f Fixtures) Paths() []string {
	sorted := make([]string, 0, len(f))
	for p, _ := range f {
		sorted = append(sorted, p)
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})
	return sorted
}

type Diff []*FileDiff

type FileDiff struct {
	Path string
	Kind DiffKind
	A    []byte
	B    []byte
}

// DiffKind describes how a file differs between fixture a and b. See
// Fixtures.Diff for more information.
type DiffKind string

var (
	// DiffMissing means that the file is only present in fixture a.
	DiffMissing DiffKind = "missing"
	// DiffUnexpected means that the file is only present in fixture b.
	DiffUnexpected DiffKind = "added"
	// DiffChanged means that the file content in fixture a is different from b.
	DiffChanged DiffKind = "changed"
)
