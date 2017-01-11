package goldy

import "testing"

var g = FlagConfig()

func TestGoldy(t *testing.T) {
	//g := EnvConfig()
	//g := FlagConfig()
	f := g.NewFixtures()
	f.Add("hey", "you\n")
	f.Add("bar", "123\n")
	f.Test(t)
}
