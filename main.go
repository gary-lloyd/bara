package main

import (
	"github.com/gary-lloyd-tessella/bara/pkg/template"
	"fmt"
	"github.com/fatih/color"
	flag "gopkg.in/alecthomas/kingpin.v2"
	"os"
	"os/exec"
)

var (
	debug       = flag.Flag("debug", "Enable debug mode.").Short('d').Bool()
	templateDir = flag.Flag("template", "Kubernetes deployment template file to use").Short('t').Required().String()
	configFile  = flag.Flag("config", "Environment configuration file to use").Short('c').Required().String()
	dryrun      = flag.Flag("dryrun", "Build templates without applying to cluster").Short('r').Bool()
)

func main() {
	flag.Version("0.0.1")
	flag.CommandLine.HelpFlag.Short('h')
	flag.Parse()
	color.Blue("Using template directory: %s\n", *templateDir)
	color.Blue("Using config: %s\n", *configFile)
	if *dryrun {
		color.Yellow("Dry run - Templates will no be applied")
	}

	outputDirectory := ".bara"
	createOutputDirectory(outputDirectory)
	template.ProcessTemplates(*templateDir, *configFile, outputDirectory)

	kubectlPath, _ := exec.LookPath("kubectl")
	fmt.Println(fmt.Sprintf("Using kubectl from path: %s", kubectlPath))

	templateDirToExecute := outputDirectory + "/" + *templateDir
	fmt.Println(templateDirToExecute)

	cmd := exec.Command(kubectlPath, "apply", "-f", templateDirToExecute)
	out, err := cmd.Output()

	if err != nil {
		// Log the error and continue as we want to apply all valid manifests
		fmt.Println(string(err.(*exec.ExitError).Stderr))
	}
	fmt.Println(string(out))
}

func createOutputDirectory(dirName string) bool {
	src, err := os.Stat(dirName)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirName, 0755)
		if errDir != nil {
			panic(fmt.Sprintf("Unable to create output direcory: %s", dirName))
		}
		return true
	} else {
		fmt.Println("Clearing existing output directory")
		os.Remove(dirName)
	}

	if src.Mode().IsRegular() {
		fmt.Println(dirName, "already exist as a file!")
		return false
	}

	return false
}