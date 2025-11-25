package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/color"

	"github.com/curioswitch/go-build"
)

func main() {
	build.DefineTasks()
	Main()
}

// Main is an extension of goyek.Main which additionally
// defines flags and uses the most useful middlewares.
func Main() {
	flag.CommandLine.SetOutput(goyek.Output())
	flag.Usage = usage
	tasks, args := goyek.SplitTasks(os.Args[1:])
	if err := flag.CommandLine.Parse(args); err != nil {
		fmt.Fprintln(goyek.Output(), err)
		os.Exit(2)
	}

	goyek.UseExecutor(color.ReportFlow)

	goyek.Use(color.ReportStatus)

	var opts []goyek.Option

	goyek.SetUsage(usage)
	goyek.SetLogger(&color.CodeLineLogger{})
	goyek.Main(tasks, opts...)
}

func usage() {
	fmt.Println("Usage of build: [flags] [--] [tasks]")
	goyek.Print()
	fmt.Println("Flags:")
	flag.PrintDefaults()
}
