<div align="center">
<img
    width=40%
    src="images/gopher-bug.svg"
    alt="proji logo"
/>

[![Go Report Card](https://goreportcard.com/badge/github.com/nikoksr/simplog)](https://goreportcard.com/report/github.com/nikoksr/simplog)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/643b7cce9fd2491e9fde38de6e1c58ad)](https://www.codacy.com/manual/nikoksr/proji?utm_source=github.com&utm_medium=referral&utm_content=nikoksr/proji&utm_campaign=Badge_Grade)
[![Maintainability](https://api.codeclimate.com/v1/badges/c9295422ae29fb489503/maintainability)](https://codeclimate.com/github/nikoksr/simplog/maintainability)
[![Actions Status](https://github.com/nikoksr/simplog/workflows/Go-Test/badge.svg)](https://github.com/{owner}/{repo}/actions)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/nikoksr/simplog)

</div>

<p align="center">A simple logger. No dependencies, no special features, just logging.</p>

<h1></h1>

#### About

<p>Simplog was created due to the unsuccessful search for a small and simple logging system. I needed a system that's exactly in the middle between 'does-what-it-should' and 'no-overkill'. Besides that I just wanted to write a logging library myself for fun.

If you like it and/or if it is useful for you, it has already more than fulfilled its purpose. Contributions are still welcome but just remember that this logger is intentionally kept simple.

</p>

#### Install <a id="install"></a>

`go get -u github.com/nikoksr/simplog`

#### Example Code

```go
// Open log file
logPath := "simple.log"

f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
	log.Fatalf("could not open log file %s. %v", logPath, err)
}

// Create logger
var log = simplog.New("Core", false, f)
if err != nil {
    log.Fatalf("could not create logger %v", err)
}

// Set log level to warning
log.SetLevel(simplog.Warning)

// fmt.Print style
log.Info("You're awesome!\n")

// fmt.Println style
log.Warningln("Coffee is almost empty.")

// fmt.Printf style
log.Debugf("%d should be equal to %d.\n", aNumber, anotherNumber)
```

#### Example Output

    CORE 2020/03/18 17:13:24 main.go:24: INFO You're awesome!
    CORE 2020/03/18 17:13:24 main.go:26: WARN Coffee is almost empty.
    CORE 2020/03/18 17:13:25 main.go:28: DEBUG 17 should be equal to 9.

#### Credits

-   Logo from [MariaLetta/free-gophers-pack](https://github.com/MariaLetta/free-gophers-pack)
