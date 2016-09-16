package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
)

var (
	flags = pflag.NewFlagSet("", pflag.ExitOnError)

	dir     = flags.StringP("dir", "d", "/var/lib/docker/containers", "The directory where the files are stored")
	pattern = flags.StringP("pattern", "p", "*-json.log", "The pattern that identifies the files that should be truncated")
)

func emptyLog(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Print(err)
		return nil
	}

	_, filename := filepath.Split(path)

	if matched, _ := filepath.Match(*pattern, filename); matched {
		fmt.Println(path, "Truncating")
		os.Truncate(path, 0)
	}

	return nil
}

func main() {
	flags.AddGoFlagSet(flag.CommandLine)
	flags.Parse(os.Args)
	err := filepath.Walk(*dir, emptyLog)
	if err != nil {
		log.Fatal(err)
	}
}
