package goldy

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

const (
	defaultEnv  = "GOLDY_UPDATE"
	defaultFlag = "update"
)

func EnvConfig() Config {
	update := os.Getenv(defaultFlag) != ""
	return Config{
		Update: func() bool { return update },
		Hint:   defaultEnv + "=true go test",
	}.WithDefaults()
}

var (
	flagAdded bool
	flagValue bool
	flagM     sync.Mutex
)

func FlagConfig() Config {
	// Concurrent calls to this func might happen when testing multiple packages
	// at once, e.g. via go test ./...
	flagM.Lock()
	defer flagM.Unlock()
	// Adding our flag more than once would panic, so let's not do that.
	if !flagAdded {
		flag.BoolVar(&flagValue, defaultFlag, false, "Update goldy test fixtures")
	}

	return Config{
		Update: func() bool { return flagValue },
		Hint:   "go test -" + defaultFlag,
	}.WithDefaults()
}

type Config struct {
	// Base is the name of the directory for storing fixtures. Defaults to
	// "test-fixtures" after calling WithDefaults
	Base   string
	Update func() bool
	Hint   string
}

func (c Config) WithDefaults() Config {
	if c.Base == "" {
		c.Base = "test-fixtures"
	}
	return c
}

func (c Config) NewFixtures(name ...string) *Fixtures {
	return &Fixtures{name: append([]string{c.Base}, name...), c: c}
}

type Fixtures struct {
	name     []string
	fixtures []*fixture
	c        Config
}

func (f *Fixtures) Add(name string, got interface{}) {
	f.fixtures = append(f.fixtures, newFixture(append(f.name, name), got))
}

func (f *Fixtures) Test(t *testing.T) {
	update := f.c.Update()
	f.checkFiles(t, update)
	for _, fixture := range f.fixtures {
		fixture.Test(t, update, f.c.Hint)
	}
}

func (f *Fixtures) checkFiles(t *testing.T, update bool) {
	var newIndexM = map[string]struct{}{}
	//var newIndexS string
	for _, fixture := range f.fixtures {
		name := fixture.Name()
		newIndexM[name] = struct{}{}
		//newIndexS += name + "\n"
	}

	path := filepath.Join(f.name...)
	files, err := ioutil.ReadDir(path)
	if err != nil && !os.IsNotExist(err) {
		t.Error(err)
	}

	for _, file := range files {
		name := file.Name()
		filePath := filepath.Join(path, name)
		if _, ok := newIndexM[name]; !ok {
			if update {
				if err := os.Remove(filePath); err != nil {
					t.Error(err)
				}
			} else {
				t.Errorf("%s: unknown file: run `"+f.c.Hint+"` to remove", filePath)
			}
		}
	}

	//indexPath := filepath.Join(append(f.name, f.c.Index)...)
	//if update {
	//if err := os.MkdirAll(filepath.Dir(indexPath), 0700); err != nil {
	//t.Error(err)
	//} else if err := ioutil.WriteFile(indexPath, []byte(newIndexS), 0600); err != nil {
	//t.Error(err)
	//}
	//}
}

func newFixture(path []string, got interface{}) *fixture {
	return &fixture{
		Got:  []byte(fmt.Sprintf("%s", got)),
		Path: filepath.Join(path...),
	}
}

type fixture struct {
	Path string
	Got  []byte
}

func (f *fixture) Name() string {
	return filepath.Base(f.Path)
}

func (f *fixture) Test(t *testing.T, update bool, hint string) {
	var err error
	if update {
		err = f.update()
	}
	if err == nil {
		err = f.verify(hint)
	}
	if err != nil {
		t.Error(err)
	}
}

func (f *fixture) update() error {
	if err := os.MkdirAll(filepath.Dir(f.Path), 0700); err != nil {
		return err
	} else {
		return ioutil.WriteFile(f.Path, f.Got, 0600)
	}
}

func (f *fixture) verify(hint string) error {
	if want, err := ioutil.ReadFile(f.Path); os.IsNotExist(err) {
		return fmt.Errorf("%s: does not exist: run `"+hint+"` to update", f.Path)
	} else if err != nil {
		return err
	} else if !bytes.Equal(f.Got, want) {
		return fmt.Errorf("%s: does not match: run `"+hint+"` to update", f.Path)
	}
	return nil
}
