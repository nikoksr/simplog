<div align="center">
<img
    width=40%
    src="assets/gopher-bug.svg"
    alt="simplog logo"
/>

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/nikoksr/simplog?color=success&label=version&sort=semver)
[![codecov](https://codecov.io/gh/nikoksr/simplog/branch/main/graph/badge.svg?token=NY51VEB9GZ)](https://codecov.io/gh/nikoksr/simplog)
[![Go Report Card](https://goreportcard.com/badge/github.com/nikoksr/simplog)](https://goreportcard.com/report/github.com/nikoksr/simplog)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/2d2f6bfb58834346b790dd35657f1a33)](https://www.codacy.com/gh/nikoksr/simplog/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=nikoksr/simplog&amp;utm_campaign=Badge_Grade)
[![Maintainability](https://api.codeclimate.com/v1/badges/c9295422ae29fb489503/maintainability)](https://codeclimate.com/github/nikoksr/simplog/maintainability)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/nikoksr/simplog)

</div>

<p align="center">A simple and opinionated library that lets you set up and use <a href="https://github.com/uber-go/zap">zap</a> quickly.

<h1></h1>

#### About

Simplog is a small library that sets [zap](https://github.com/uber-go/zap) up in a way that is easy to use and provides
some additional features. Simplog is opinionated and tries to provide a good default configuration for most use cases.
If you need more features, you may use the resulting zap logger directly.

#### Features

- Client and server modes
- Visualized log levels
- Optional context-binding support
- Easy to use

#### Install <a id="install"></a>

```bash
go get -u github.com/nikoksr/simplog
```

&nbsp;

#### Example Code

> For more examples, see the [examples](examples) directory.

```go
package main

import "github.com/nikoksr/simplog"

func main() {
  // Using the manual configuration; alternatively you can use NewClientLogger() or NewServerLogger().
  logger := simplog.NewWithOptions(&simplog.Options{
    Debug:             false,
    IsServer:          true,
  })

  // At this point, you're using a zap.SugaredLogger and can use it as you would normally do.
  logger.Info("You're awesome!")
  logger.Warn("Coffee is almost empty!")
  logger.Error("Unable to operate, caffein levels too low.")
}
```

&nbsp;

#### Example Outputs

##### Client & server mode in debug

> In debug mode, independent of the mode, the logger will print all messages greater-equal than the debug-level in a human readable format.

```bash
2022-10-23T14:25:15.537+0200	INFO	simplog	test-simplog/main.go:12	You're awesome!
2022-10-23T14:25:15.537+0200	WARN	simplog	test-simplog/main.go:13	Coffee is almost empty!
2022-10-23T14:25:15.537+0200	ERROR	simplog	test-simplog/main.go:14	Unable to operate, caffein levels too low.
```

&nbsp;

##### Client mode in production

> In production mode, the client logger will print all messages greater-equal than the info-level in a human readable format and replace the log level with a colored emoji.
>
> The symbols are configurable and can be set to any string.

```bash
💡 You're awesome!
⚠️ Coffee is almost empty!
🔥 Unable to operate, caffein levels too low.
```

&nbsp;

##### Server mode in production

> In production mode, the server logger will print all messages greater-equal than the info-level in a structured format.

```bash
{"level":"info","ts":1666528089.4873903,"logger":"simplog","caller":"test-simplog/main.go:12","msg":"You're awesome!"}
{"level":"warn","ts":1666528089.4874253,"logger":"simplog","caller":"test-simplog/main.go:13","msg":"Coffee is almost empty!"}
{"level":"error","ts":1666528089.487434,"logger":"simplog","caller":"test-simplog/main.go:14","msg":"Unable to operate, caffein levels too low."}
```

&nbsp;

#### Credits

- Logo by the amazing [MariaLetta/free-gophers-pack](https://github.com/MariaLetta/free-gophers-pack)
