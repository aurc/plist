# plist

[![GoDoc](https://godoc.org/github.com/aurc/plist?status.svg)](https://godoc.org/github.com/aurc/plist)
![example workflow](https://github.com/aurc/plist/actions/workflows/go.yml/badge.svg)
[![codecov](https://codecov.io/gh/aurc/plist/branch/main/graph/badge.svg?token=NZKNTEJ5YC)](https://codecov.io/gh/aurc/plist)
[![GitHub license](https://img.shields.io/github/license/aurc/plist)](https://github.com/aurc/plist/blob/main/LICENSE)


Convert Apple's `plist` file format to `JSON` or `YAML` (natural & high fidelity modes) effortlessly.
Often complex bundles and other files are very hard to read or seamlessly 
port to other applications. 

**This Package Provides:**
- A **CLI** tool for reading and converting PLIST that can be fully intgrated
  with `shell` scripting.
- A **golang module** that can be imported directly into golang projects with
  an easy to use API.
- A **native bundle** *(coming soon)* that can be imported to any C-compatible
  applications (e.g. Swift, C, C++, Python, etc).

## Installation

### Homebrew For the Inpatient

Consider using homebrew to install the utility from there:
````
brew install aurc/tap/plist
````
More information here: [plist homebrew tap](https://github.com/aurc/homebrew-tap)

### Golang

````bash
# As CLI Tool
go install github.com/aurc/plist@latest
# Then (assumung $GOPATH/bin set in the path)
plist --help
````
OR
````bash
# For working only with parser lib on your go code
go get github.com/aurc/plist/pkg/plistparser
# For entire package with CLI tool on your local (more below)
go get github.com/aurc/plist
````

## Basic Usage

### CLI Tool

To get started:
````
go install github.com/aurc/plist@latest
````

Assuming your `$GOPATH/bin` is in your path, then:

````
plist -h
````
Output:
````
This tool converts Apple's Property List (.plist) inputs into several useful
formats, such as JSON and YAML.

It supports both a file name as input and a piped ('|') input which might be useful
on more involved shell scripts.

For example:
    plist json -i myfile.plist
    cat myfile.plist | plist json

For individual commands instructions run:
        plist [command] -h
        plist json -h

Usage:
  plist [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  json        Converts plist into JSON
  yaml        Converts plist into YAML

Flags:
  -h, --help            help for plist
  -x, --high-fidelity   Specifies whether the output should be a one-to-one translation of the plist. 
                        Set to true, it's one-to-one. The default is false as it produces a more readable file.
  -i, --input string    Specifies a input file, e.g. --input myFile.plist

Use "plist [command] --help" for more information about a command.

````

Given the input

````xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
    <dict>
        <key>idle_wakeups</key><integer>103</integer>
        <key>idle_wakeups_per_s</key><real>20.533</real>
        <key>timer_wakeups</key>
        <array>
            <dict>
                <key>interval_ns</key><integer>2000000</integer>
                <key>wakeups</key><integer>270</integer>
                <key>wakeups_per_s</key><real>53.8244</real>
            </dict>
            <dict>
                <key>interval_ns</key><integer>1500000</integer>
                <key>wakeups</key><integer>170</integer>
                <key>wakeups_per_s</key><real>22.44</real>
            </dict>
        </array>
    </dict>
</plist>
````
#### Output as **JSON (Natural conversion)**
````
plist json -i docs/demo/demo.plist -p true
````
OR
````
cat docs/demo/demo.plist | plist json -p true
````
The natural conversion will hide the verbosity of datatypes, having types
to be easily inferred, for example a `<real>1.1</real>` becomes just `1.1` 
without quotes whereas a `<dict>` will look more like a `JSON` objects, with
the `key` entries listed as fields.
````json
{
  "idle_wakeups": 103,
  "idle_wakeups_per_s": 20.533,
  "timer_wakeups": [
    {
      "interval_ns": 2000000,
      "wakeups": 270,
      "wakeups_per_s": 53.8244
    },
    {
      "interval_ns": 1500000,
      "wakeups": 170,
      "wakeups_per_s": 22.44
    }
  ]
}
````
#### Output as **YAML (Natural conversion)**
````
plist yaml -i docs/demo/demo.plist
````
OR
````
cat docs/demo/demo.plist | plist yaml
````
Gives...
````yaml
idle_wakeups: 103
idle_wakeups_per_s: 20.533
timer_wakeups:
- interval_ns: 2000000
  wakeups: 270
  wakeups_per_s: 53.8244
- interval_ns: 1500000
  wakeups: 170
  wakeups_per_s: 22.44
````

#### Output as **JSON (High fidelity)**
````
plist json -i docs/demo/demo.plist -p true -x true
````
OR
````
cat docs/demo/demo.plist | plist json -p true -x true
````
Depending on your application, you might want a full, high fidelity
translation, defining each entry and explicitly exposing the data types. For
that, given the same input above, you'll have the following output:
````json
{
  "type": "dict",
  "value": [
    {
      "k": "idle_wakeups",
      "v": {
        "type": "integer",
        "value": "103"
      }
    },
    {
      "k": "idle_wakeups_per_s",
      "v": {
        "type": "real",
        "value": "20.533"
      }
    },
    {
      "k": "timer_wakeups",
      "v": {
        "type": "array",
        "value": [
          {
            "type": "dict",
            "value": [
              {
                "k": "interval_ns",
                "v": {
                  "type": "integer",
                  "value": "2000000"
                }
              },
              {
                "k": "wakeups",
                "v": {
                  "type": "integer",
                  "value": "270"
                }
              },
              {
                "k": "wakeups_per_s",
                "v": {
                  "type": "real",
                  "value": "53.8244"
                }
              }
            ]
          },
          {
            "type": "dict",
            "value": [
              {
                "k": "interval_ns",
                "v": {
                  "type": "integer",
                  "value": "1500000"
                }
              },
              {
                "k": "wakeups",
                "v": {
                  "type": "integer",
                  "value": "170"
                }
              },
              {
                "k": "wakeups_per_s",
                "v": {
                  "type": "real",
                  "value": "22.44"
                }
              }
            ]
          }
        ]
      }
    }
  ]
}
````
Output as **YAML (High fidelity)**
````
plist yaml -i docs/demo/demo.plist -x true
````
OR
````
cat docs/demo/demo.plist | plist yaml -x true
````
````yaml
type: dict
value:
- k: idle_wakeups
  v:
    type: integer
    value: "103"
- k: idle_wakeups_per_s
  v:
    type: real
    value: "20.533"
- k: timer_wakeups
  v:
    type: array
    value:
    - type: dict
      value:
      - k: interval_ns
        v:
          type: integer
          value: "2000000"
      - k: wakeups
        v:
          type: integer
          value: "270"
      - k: wakeups_per_s
        v:
          type: real
          value: "53.8244"
    - type: dict
      value:
      - k: interval_ns
        v:
          type: integer
          value: "1500000"
      - k: wakeups
        v:
          type: integer
          value: "170"
      - k: wakeups_per_s
        v:
          type: real
          value: "22.44"

````

## As a Package 

````bash
$ go get github.com/aurc/plist/pkg/plistparser
````

````go
package main

import (
    "fmt"
	
    "github.com/aurc/plist/pkg/plistparser"
)

func main() {
  plistFile := `<?xml version="1.0" encoding="UTF-8"?>
  <!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
  <plist version="1.0">
  <array>
      <string>String</string>
  </array>
  </plist>`

  out, _ := plistparser.Convert([]byte(plistFile), &plistparser.Config{
    Target:       plistparser.Json,
    HighFidelity: false,
    Beatify:      true,
  })
  
  fmt.Println(string(out))
}
````

## License

`plist` is released under the Apache 2.0 license. See [LICENSE](https://github.com/aurc/plist/blob/master/LICENSE)




