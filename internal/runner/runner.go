package runner

import (
	"fmt"
	"os"
	"strconv"

	"github.com/logrusorgru/aurora"
	"github.com/projectdiscovery/gologger"
	"github.com/remeh/sizedwaitgroup"
	"ktbs.dev/ssb/pkg/ssb"
)

// New execute bruteforces
func New(opt *Options) {
	opt.showInfo()

	defer opt.Close()
	job := make(chan string)
	cur := opt.concurrent
	swg := sizedwaitgroup.New(cur)

	for i := 0; i < cur; i++ {
		swg.Add()
		go func() {
			defer swg.Done()
			for pass := range job {
				for i := 0; i < opt.retries; i++ {
					if opt.run(pass) {
						break
					}
				}
			}
		}()
	}

	for scanner := opt.listScanner(); scanner.Scan(); {
		job <- scanner.Text()
	}

	close(job)
	swg.Wait()
	gologger.Infof("Done!")
}

func (opt *Options) run(password string) bool {
	cfg := ssb.New(opt.user, password, opt.timeout)

	con, err := ssb.Connect(opt.host, opt.port, cfg)
	if err != nil {
		if opt.verbose {
			gologger.Errorf("Failed '%s': %s.", password, err.Error())
		}
	}

	if con {
		fmt.Printf("[%s] Connected with '%s'.\n", aurora.Green("VLD"), aurora.Magenta(password))

		if opt.file != nil {
			fmt.Fprintf(opt.file, "%s\n", password)
		}

		vld = true
	}

	return vld
}

func (opt *Options) showInfo() {
	info := "________________________\n"
	info += "\n :: Username: " + opt.user
	info += "\n :: Hostname: " + opt.host
	info += "\n :: Port    : " + strconv.Itoa(opt.port)
	info += "\n :: Wordlist: " + opt.wordlist
	info += "\n :: Threads : " + strconv.Itoa(opt.concurrent)
	info += "\n :: Timeout : " + opt.timeout.String()
	info += "\n________________________\n\n"

	fmt.Fprint(os.Stderr, info)
}
