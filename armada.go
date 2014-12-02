package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	//var manifest armada.Manifest

	//if _, err := toml.DecodeFile("myservice.toml", &manifest); err != nil {
	//fmt.Println(err)
	//return
	//}

	//fmt.Printf("Name:  %+v\n", manifest)

	app := cli.NewApp()
	app.Version = releaseVersion
	app.Name = "armada"
	app.Run(os.Args)

}
