package main

import (
	"github.com/lessbutter/alloff-api/cmd"
	"github.com/lessbutter/alloff-api/pkg/seeder/malls"
)

func main() {
	cmd.SetBaseConfig()

	malls.AddFlannels()
}
