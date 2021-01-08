# SSB

[![Kitabisa SecLab](https://img.shields.io/badge/kitabisa-security%20project-blue)](#)
[![License](https://img.shields.io/badge/License-Apache%202.0-yellowgreen)](https://github.com/kitabisa/ssb/blob/master/LICENSE)
[![contributions](https://img.shields.io/badge/contributions-welcome-magenta.svg?style=flat)](https://github.com/kitabisa/ssb/issues)
[![made with Go](https://img.shields.io/badge/made%20with-Go-brightgreen)](http://golang.org)
[![Version](https://img.shields.io/badge/version-0.1.0-blueviolet)](https://github.com/kitabisa/ssb/releases)

**S**_ecure_ **S**_hell_ **B**_ruteforcer_ — A faster & simpler way to bruteforce SSH server.

<img src="https://user-images.githubusercontent.com/25837540/102938182-61d96e00-44de-11eb-81c9-2e37cdcfba8d.png" width="680px;" alt="Kitabisa SSB Demo"/>

---

## Installation

### from Binary

Download a pre-built binary from [releases page](https://github.com/kitabisa/ssb/releases), unpack and run! Or:

```bash
▶ (sudo) curl -sSfL 'https://git.io/kitabisa-ssb' | sh -s -- -b /usr/local/bin
```

### from Source

Need [go1.14+](https://golang.org/doc/install#download) compiler installed and configured, then:

```bash
▶ GO111MODULE=on go get ktbs.dev/ssb
```

## Usage

```bash
▶ ssb [-p port] [-w wordlist.txt] [-t timeout]
      [-c concurrent] [-r retries] [-o output] [user@]hostname
```

### Options:

```txt
  -p port
     Port to connect to on the remote host (default 22).
  -w wordlist
     Path to wordlist file.
  -t timeout
     Connection timeout (default 30s).
  -c concurrent
     Concurrency/threads level (default 100).
  -r retries
     Specify the connection retries (default 1).
  -o output
     Save valid password to file.
  -v
     Verbose mode.
```

## License

This program is free software: you can redistribute it and/or modify it under the terms of the [Apache license](https://github.com/kitabisa/ssb/blob/master/LICENSE). Kitabisa SSB and any contributions are Copyright © by Dwi Siswanto 2020.