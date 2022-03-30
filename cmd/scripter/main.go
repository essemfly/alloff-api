package main

import (
	"fmt"
	"github.com/lessbutter/alloff-api/cmd"
	"github.com/lessbutter/alloff-api/cmd/scripter/scripts"
	"github.com/lessbutter/alloff-api/pkg/seeder"
)

var (
	GitInfo   = "no info"
	BuildTime = "no datetime"
	Env       = "dev"
)

func main() {
	fmt.Println("Git commit information: ", GitInfo)
	fmt.Println("Build date, time: ", BuildTime)

	cmd.SetBaseConfig(Env)
	seeder.MakeClassifyRules()
	scripts.AddBrandsSeeder()
	scripts.AddClassifyRules()
}
