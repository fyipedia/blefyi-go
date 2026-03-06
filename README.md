# blefyi-go

[![Go Reference](https://pkg.go.dev/badge/github.com/fyipedia/blefyi-go.svg)](https://pkg.go.dev/github.com/fyipedia/blefyi-go)

Go client for the [BLEFYI](https://blefyi.com) API. Look up Bluetooth Low Energy chips, profiles, versions, beacons, use cases, and manufacturers. Zero dependencies — stdlib only.

## Install

```bash
go get github.com/fyipedia/blefyi-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    blefyi "github.com/fyipedia/blefyi-go"
)

func main() {
    client := blefyi.NewClient()

    result, err := client.Search("heart rate")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Found %d results for 'heart rate'\n", result.Total)

    chip, err := client.Chip("nrf52840")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s: %s\n", chip.Name, chip.Description)
}
```

## API Methods

| Method | Description |
|--------|-------------|
| `Search(query)` | Search chips, profiles, and glossary |
| `Chip(slug)` | Get BLE chip details |
| `Profile(slug)` | Get BLE profile details |
| `Version(slug)` | Get BLE version details |
| `Beacon(slug)` | Get beacon type details |
| `Usecase(slug)` | Get use case details |
| `Manufacturer(slug)` | Get manufacturer details |
| `GlossaryTerm(slug)` | Get glossary term definition |
| `Compare(slugA, slugB)` | Compare two BLE chips |
| `Random()` | Get a random BLE chip |

## REST API

Free, no authentication required. Base URL: `https://blefyi.com/api`

```bash
curl https://blefyi.com/api/search/?q=heart+rate
curl https://blefyi.com/api/chip/nrf52840/
curl https://blefyi.com/api/random/
```

Full OpenAPI spec: https://blefyi.com/api/openapi.json

## Also Available

| Language | Package | Install |
|----------|---------|---------|
| Python | [blefyi](https://pypi.org/project/blefyi/) | `pip install blefyi` |
| TypeScript | [blefyi](https://www.npmjs.com/package/blefyi) | `npm install blefyi` |
| Go | [blefyi-go](https://pkg.go.dev/github.com/fyipedia/blefyi-go) | `go get github.com/fyipedia/blefyi-go` |
| Rust | [blefyi](https://crates.io/crates/blefyi) | `cargo add blefyi` |
| Ruby | [blefyi](https://rubygems.org/gems/blefyi) | `gem install blefyi` |

## Code FYI Family

| Site | Domain | Focus |
|------|--------|-------|
| BarcodeFYI | [barcodefyi.com](https://barcodefyi.com) | Barcode symbologies and standards |
| QRCodeFYI | [qrcodefyi.com](https://qrcodefyi.com) | QR code types and encoding |
| NFCFYI | [nfcfyi.com](https://nfcfyi.com) | NFC chips and standards |
| BLEFYI | [blefyi.com](https://blefyi.com) | Bluetooth Low Energy profiles |
| RFIDFYI | [rfidfyi.com](https://rfidfyi.com) | RFID tags and frequencies |
| SmartCardFYI | [smartcardfyi.com](https://smartcardfyi.com) | Smart card platforms and standards |

## License

MIT
