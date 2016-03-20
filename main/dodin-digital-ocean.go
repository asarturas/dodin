package main

import (
	"github.com/asarturas/dodin"
	"os"
	"github.com/asarturas/dodin/cloud/do"
	"fmt"
)

func main() {
	configProvider := dodin.ConfigFileProvider{"dodin-digital-ocean.ini"}
	cloudProvider := do.Provider(os.Getenv("DO_API_TOKEN"))
	fmt.Print(dodin.Dodin(configProvider, cloudProvider))
}
