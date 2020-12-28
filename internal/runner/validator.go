package runner

import (
	"bufio"
	"errors"
	"flag"
	"os"
	"os/user"
	"strings"
)

func (opt *Options) validate() error {
	var errFile error

	if flag.Arg(0) != "" {
		uhost = strings.SplitN(flag.Arg(0), "@", 2)
	} else {
		return errors.New("Please define a target server")
	}

	if len(uhost) < 2 {
		usr, err := user.Current()
		if err != nil {
			opt.user = "root"
		} else {
			opt.user = usr.Username
		}
		opt.host = uhost[0]
	} else {
		opt.user = uhost[0]
		opt.host = uhost[1]
	}

	if opt.wordlist != "" {
		f, err := os.Open(opt.wordlist)
		if err != nil {
			return err
		}
		opt.list = bufio.NewScanner(f)
	} else {
		return errors.New("No wordlist file provided")
	}

	if opt.output != "" {
		opt.file, errFile = os.OpenFile(opt.output, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if errFile != nil {
			return errFile
		}
	}

	return nil
}
