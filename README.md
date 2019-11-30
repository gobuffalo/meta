# github.com/gobuffalo/meta/v2

[![](https://github.com/gobuffalo/meta/workflows/Tests/badge.svg)](https://github.com/gobuffalo/meta/actions)
[![GoDoc](https://godoc.org/github.com/gobuffalo/meta?status.svg)](https://godoc.org/github.com/gobuffalo/meta)

### Requirements

* Go 1.13+
* Go Modules

### Installation

```bash
$ go get github.com/gobuffalo/meta/v2
```

---

## V2 Toml

```toml
bin = "bin/coke"
vcs = "git"

[with]
  nodejs = true
  pop = true
  sqlite = false
  webpack = true
  yarn = true

[as]
  api = false
  web = true
```

## V1 Toml

```toml
name = "coke"
bin = "bin/coke"
vcs = "git"
with_pop = true
with_sqlite = false
with_dep = false
with_webpack = true
with_nodejs = true
with_yarn = true
with_docker = true
with_grifts = true
as_web = true
as_api = false
```
