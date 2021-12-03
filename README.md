# plist toolkit
Convert Apple's `plist` file format to `JSON` or `YAML` (natural & high fidelity modes) effortlessly.
Often complex bundles and other files are
very hard to read or seamlessly port to other applications. This library
offers several integration options for the ingestion on these files.

**Here we can find:**
- A **CLI** tool for reading and converting PLIST that can be fully intgrated
with `shell` scripting.
- A **golang module** that can be imported directly into golang projects with
an easy to use API.
- A **native bundle** that can be imported to any C-compatible 
applications (e.g. Swift, C, C++, Python, etc).

## Quick Snapshot (CLI)
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
Output as **JSON (Natural conversion)**
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
Output as **YAML (Natural conversion)**
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
However, depending on your application, you might want a full, high fidelity
translation, defining each entry and explicitly exposing the data types. For
that, given the same input above, you'll have the following output:

Output as **JSON (High fidelity)**
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

## Getting Started

### CLI (From Source)
Make sure you have `golang 1.17` or greater, then:
````
git clone https://github.com/aurc/plist.git
````




