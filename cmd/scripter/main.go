package main

import (
	"fmt"
	"github.com/lessbutter/alloff-api/cmd"
	"github.com/lessbutter/alloff-api/cmd/scripter/scripts"
)

var (
	GitInfo   = "no info"
	BuildTime = "no datetime"
	Env       = "local"
)

func main() {
	fmt.Println("Git commit information: ", GitInfo)
	fmt.Println("Build date, time: ", BuildTime)

	cmd.SetBaseConfig(Env)
	scripts.AddBrandsSeeder()
	scripts.AddClassifyRules()
}
