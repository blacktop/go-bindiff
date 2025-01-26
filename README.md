<p align="center">
  <a href="https://github.com/blacktop/go-bindiff"><img alt="Logo" src="https://github.com/blacktop/go-bindiff/raw/main/logo.png" height="300" /></a>
  <h1 align="center">go-bindiff</h1>
  <h4><p align="center">Go package to interface with bindiff</p></h4>
  <p align="center">
    <a href="https://github.com/blacktop/go-bindiff/actions" alt="Actions">
          <img src="https://github.com/blacktop/go-bindiff/actions/workflows/go.yml/badge.svg" /></a>
    <a href="https://pkg.go.dev/github.com/blacktop/go-bindiff" alt="Docs">
          <img src="https://pkg.go.dev/badge/github.com/blacktop/go-bindiff.svg" /></a>
    <a href="http://doge.mit-license.org" alt="LICENSE">
          <img src="https://img.shields.io/:license-mit-blue.svg" /></a>
</p>
<br>

## Getting Started

```bash
go get github.com/blacktop/go-bindiff
```

### `bindiff` demo tool

Install

```bash
go install github.com/blacktop/go-bindiff/cmd/bindiff@latest
```

Usage

```bash
❯ bindiff path/to/BinExport1 path/to/BinExport2
```
```bash
   • Running bindiff...
   • Reading bindiff results...
   • bindiff complete
   • memset_s                  confidence=0.99 similarity=1.00
   • timingsafe_bcmp           confidence=0.99 similarity=1.00
   • cc_clear                  confidence=0.99 similarity=1.00
   • cc_disable_dit            confidence=0.97 similarity=1.00
   • ccdigest_init             confidence=0.99 similarity=1.00
   • ccdigest_update           confidence=0.99 similarity=1.00
   • cchmac                    confidence=0.99 similarity=1.00
   • cchmac_init               confidence=0.99 similarity=1.00
   • cchmac_update             confidence=0.99 similarity=1.00
   • cchmac_final              confidence=0.99 similarity=1.00
   • ccdigest_final_64be       confidence=0.99 similarity=1.00
   <SNIP>
```

## License

MIT Copyright (c) 2025 **blacktop**