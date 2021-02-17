package main

import (
	"com.roxorgaming/ja/app"
	"com.roxorgaming/ja/mvn"
	"flag"
	"fmt"
	"os"
)

var command string
var fileType string
var automatic bool
var snapshot bool
var major string
var minor string
var patch string
var help string

// main runs the command-line parsing and validations. This function will also start the application logic execution.
func main() {

	const (

		short = " (shorthand)"

		defaultCmd = ""
		usageCmd = "command to run. mvn | gocd"

		defaultFile = ""
		usageFile   = "path of the maven pom file to change"

		defaultAuto = false
		usageAuto = "automatically increase the version number"

		defaultSnapShot = false
		usageSnapShot = "increase the version number but keep the SNAPSHOT if present"

		defaultMajor = "0"
		usageMajor  = "semantic versioning the major number"

		defaultMinor = "0"
		usageMinor   = "semantic versioning the minor number"

		defaultPatch= "0"
		usagePatch   = "semantic versioning the patch number"

		defaultHelp = ""
		usageHelp = "show help for mvn or gocd"
	)

	flag.StringVar(&command, "command", defaultCmd, usageCmd)
	flag.StringVar(&command, "-c", defaultCmd, usageCmd + short)
	// Parse command-line arguments
	// Register file flag.

	flag.StringVar(&fileType, "file", defaultFile, usageFile)
	flag.StringVar(&fileType, "-f", defaultFile, usageFile + short)

	flag.BoolVar(&automatic, "automatic", defaultAuto, usageAuto)
	flag.BoolVar(&automatic, "-a", defaultAuto, usageAuto + short)

	flag.BoolVar(&snapshot, "snapshot", defaultSnapShot, usageSnapShot)
	flag.BoolVar(&snapshot, "-s", defaultSnapShot, usageSnapShot + short)

	flag.StringVar(&major, "major", defaultMajor, usageMajor)
	flag.StringVar(&major, "-mj", defaultMajor, usageMajor + short)

	flag.StringVar(&minor, "minor", defaultMinor, usageMinor)
	flag.StringVar(&minor, "-mi", defaultMinor, usageMinor + short)

	flag.StringVar(&patch, "patch", defaultPatch, usagePatch)
	flag.StringVar(&patch, "-p", defaultPatch, usagePatch + short)

	flag.StringVar(&help, "help", defaultHelp, usageHelp)
	flag.StringVar(&help, "-h", defaultHelp, usageHelp + short)

	// Parse the flags.
	flag.Parse()

	if help != ""{
		fmt.Sprint("")
	}

	switch command {
	case "mvn":{
		if fileType == "" {
			fmt.Print("File (-f) is required")
			os.Exit(1)
		}
	}
	case "gocd":{

	}
	default:



	}

	// Convert to internal config
	cfg := mvn.pom_config{fileType, automatic, snapshot, major, minor, patch}

	// Run the App
	err := app.Run(command, cfg)
	if err != nil {
		// do stuff
		os.Exit(1)
	}
}
