package main

import (
	"fmt"
	"github.com/asarturas/dodin"
	"github.com/asarturas/dodin/cloud/do"
	"os"
)

func main() {
	configProvider := dodin.ConfigFileProvider{"dodin-digital-ocean.ini"}
	cloudProvider := do.Provider(os.Getenv("DO_API_TOKEN"))
	fmt.Print(dodin.Dodin(configProvider, cloudProvider))
}
