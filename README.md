# datetime

[![Test Status](https://github.com/keyhorn/datetime/workflows/test/badge.svg)](https://github.com/keyhorn/datetime/actions?query=workflow%3Atest)
[![Lint Status](https://github.com/keyhorn/datetime/workflows/lint/badge.svg)](https://github.com/keyhorn/datetime/actions?query=workflow%3Alint)
[![Coverage Status](https://coveralls.io/repos/github/keyhorn/datetime/badge.svg?branch=main)](https://coveralls.io/github/keyhorn/datetime?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/keyhorn/datetime)](https://goreportcard.com/report/github.com/keyhorn/datetime)
[![BCH compliance](https://bettercodehub.com/edge/badge/keyhorn/datetime?branch=main)](https://bettercodehub.com/)
[![Documentation](https://godoc.org/github.com/keyhorn/datetime?status.svg)](http://godoc.org/github.com/keyhorn/datetime)
[![License](https://img.shields.io/github/license/keyhorn/datetime.svg?maxAge=2592000)](https://github.com/keyhorn/datetime/LICENSE)
[![Release](https://img.shields.io/github/release/keyhorn/datetime.svg?label=Release)](https://github.com/keyhorn/datetime/releases)

Date and time library for Go

## Install

Use `go get` to install this library.

```shell
go get -u github.com/keyhorn/datetime
```

## API Document

See [GoDoc](https://godoc.org/github.com/keyhorn/datetime) for full doument.

## Usage

```golang
import "github.com/keyhorn/datetime"

func main() {
    now := datetime.Now()

    fmt.Println(now.Format("yyyy-MM-dd hh:mm:dd"))
}

```

## DateFormat

| Format    | Description                  | Example   |
|-----------|------------------------------|-----------|
| GG        | Common era                   | A.D.; B.C.|
| G         | Common era                   | AD; BC    |
| yyyy      | Year                         | 2020      |
| yy        | Year                         | 20        |
| MM        | Month in year                | 04; 12    |
| M         | Month in year                | 4; 12     |
| dd        | Day in month                 | 01; 31    |
| d         | Day in month                 | 1; 31     |
| E         | Day of the week              | Monday    |
| e         | Day of the week              | Mon       |
| HH        | Hour in day (1-24)           | 01; 24    |
| hh        | Hour in am/pm (1-12)         | 01; 12    |
| mm        | Minute in hour               | 01; 60    |
| m         | Minute in hour               | 1; 60     |
| ss        | Second in minute             | 01; 60    |
| s         | Second in minute             | 1; 60     |
| SSS       | Millisecond                  | 123       |
| SSSSSS    | Microsecond                  | 123456    |
| SSSSSSSSS | Nanosecond                   | 123456789 |
| a         | Am/pm marker                 | AM; PM    |

## License

This library is licensed under MIT license. See [LICENSE](https://github.com/keyhorn/datetime/LICENSE) for details.