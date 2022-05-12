package main

import (
	"planetscale-witn-ent/cmd/seed/seed"
	"planetscale-witn-ent/config"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})
	seed.Seed()
}
