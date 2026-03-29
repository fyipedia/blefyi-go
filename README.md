# blefyi-go

[![Go Reference](https://pkg.go.dev/badge/github.com/fyipedia/blefyi-go.svg)](https://pkg.go.dev/github.com/fyipedia/blefyi-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Go client for the [BLEFYI](https://blefyi.com) REST API. Bluetooth Low Energy, GATT, beacons. Zero external dependencies beyond stdlib.

> **Try the interactive tools at [blefyi.com](https://blefyi.com)** — explore, search, and discover.

## Install

```bash
go get github.com/fyipedia/blefyi-go
```

## Quick Start

```go
package main

import (
    "fmt"
    blefyi "github.com/fyipedia/blefyi-go"
)

func main() {
    client := blefyi.NewClient()

    // Search across all content
    result, err := client.Search("query")
    if err != nil {
        panic(err)
    }
    fmt.Println(result)
}
```

## API Methods

| Method | Description |
|--------|-------------|
| `Beacons()` | List beacons |
| `Chips()` | List chips |
| `Faqs()` | List faqs |
| `Glossary()` | List glossary |
| `Guides()` | List guides |
| `Manufacturers()` | List manufacturers |
| `Mesh()` | List mesh |
| `Profiles()` | List profiles |
| `Tools()` | List tools |
| `UseCases()` | List use cases |
| `Versions()` | List versions |
| `Search(query)` | Search across all content |

## Also Available

| Platform | Package | Link |
|----------|---------|------|
| **Python** | `pip install blefyi` | [PyPI](https://pypi.org/project/blefyi/) |
| **npm** | `npm install blefyi` | [npm](https://www.npmjs.com/package/@fyipedia/blefyi) |
| **Go** | `go get github.com/fyipedia/blefyi-go` | [pkg.go.dev](https://pkg.go.dev/github.com/fyipedia/blefyi-go) |
| **Rust** | `cargo add blefyi` | [crates.io](https://crates.io/crates/blefyi) |
| **Ruby** | `gem install blefyi` | [rubygems](https://rubygems.org/gems/blefyi) |

## Tag FYI Family

Part of the [FYIPedia](https://fyipedia.com) open-source developer tools ecosystem.

| Site | Domain | Focus |
|------|--------|-------|
| BarcodeFYI | [barcodefyi.com](https://barcodefyi.com) | Barcode formats, EAN, UPC, ISBN standards |
| **BLEFYI** | [blefyi.com](https://blefyi.com) | Bluetooth Low Energy, GATT, beacons |
| NFCFYI | [nfcfyi.com](https://nfcfyi.com) | NFC chips, tag types, NDEF, standards |
| QRCodeFYI | [qrcodefyi.com](https://qrcodefyi.com) | QR code types, versions, encoding modes |
| RFIDFYI | [rfidfyi.com](https://rfidfyi.com) | RFID tags, frequency bands, standards |
| SmartCardFYI | [smartcardfyi.com](https://smartcardfyi.com) | Smart cards, EMV, APDU, Java Card |

## Embed Widget

Embed [BLEFYI](https://blefyi.com) widgets on any website with [blefyi-embed](https://widget.blefyi.com):

```html
<script src="https://cdn.jsdelivr.net/npm/blefyi-embed@1/dist/embed.min.js"></script>
<div data-blefyi="entity" data-slug="example"></div>
```

Zero dependencies · Shadow DOM · 4 themes (light/dark/sepia/auto) · [Widget docs](https://widget.blefyi.com)

## License

MIT
