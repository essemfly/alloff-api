package main

import (
	"fmt"

	"github.com/lessbutter/alloff-api/cmd"
)

var (
	GitInfo   = "no info"
	BuildTime = "no datetime"
	Env       = "prod"
)

func main() {
	fmt.Println("Git commit information: ", GitInfo)
	fmt.Println("Build date, time: ", BuildTime)

	cmd.SetBaseConfig(Env)

}
