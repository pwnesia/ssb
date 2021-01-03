package runner

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/projectdiscovery/gologger"
)

// Options declare its options
type Options struct {
	concurrent int
	wordlist   string
	retries    int
	verbose    bool
	timeout    time.Duration
	output     string
	port       int
	list       io.ReadCloser
	file       *os.File

	user string
	host string
}

// Parse arguments
func Parse() *Options {
	opt = &Options{}
	opt.timeout = timeout

	flag.IntVar(&opt.port, "p", 22, "")
	flag.IntVar(&opt.retries, "r", 1, "")
	flag.IntVar(&opt.concurrent, "c", 100, "")
	flag.BoolVar(&opt.verbose, "v", false, "")
	flag.StringVar(&opt.output, "o", "", "")
	flag.StringVar(&opt.wordlist, "w", "", "")
	flag.DurationVar(&opt.timeout, "t", opt.timeout, "")

	flag.Usage = func() {
		showBanner()
		fmt.Fprint(os.Stderr, usage)
	}
	flag.Parse()

	showBanner()

	if err := opt.validate(); err != nil {
		gologger.Fatalf("Error! %s.", err.Error())
	}

	return opt
}

func showBanner() {
	fmt.Fprint(os.Stderr, aurora.Bold(aurora.Cyan(banner)))
}
