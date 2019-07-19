# fs-bundler

:smile: fs-bundler allows to embed local files and generate JSON/YAML archives

[![CircleCI](https://circleci.com/gh/moul/fs-bundler.svg?style=shield)](https://circleci.com/gh/moul/fs-bundler)
[![GoDoc](https://godoc.org/moul.io/fs-bundler?status.svg)](https://godoc.org/moul.io/fs-bundler)
[![License](https://img.shields.io/github/license/moul/fs-bundler.svg)](https://github.com/moul/fs-bundler/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/moul/fs-bundler.svg)](https://github.com/moul/fs-bundler/releases)
[![Go Report Card](https://goreportcard.com/badge/moul.io/fs-bundler)](https://goreportcard.com/report/moul.io/fs-bundler)
[![CodeFactor](https://www.codefactor.io/repository/github/moul/fs-bundler/badge)](https://www.codefactor.io/repository/github/moul/fs-bundler)
[![Docker Metrics](https://images.microbadger.com/badges/image/moul/fs-bundler.svg)](https://microbadger.com/images/moul/fs-bundler)
[![Made by Manfred Touron](https://img.shields.io/badge/made%20by-Manfred%20Touron-blue.svg?style=flat)](https://manfred.life/)

## Example

```console
$ fs-bundler --indent ./.*.yml
{
  "files": [
    {
      "path": "./.golangci.yml",
      "name": ".golangci.yml",
      "content": "cnVuOgogIGRlYWRsaW5lOiAxbQogIHRlc3RzOiBmYWxzZQogICNza2lwLWZpbGVzOgogICMgIC0gIi4qXFwuZ2VuXFwuZ28iCgpsaW50ZXJzLXNldHRpbmdzOgogIGdvbGludDoKICAgIG1pbi1jb25maWRlbmNlOiAwCiAgbWFsaWduZWQ6CiAgICBzdWdnZXN0LW5ldzogdHJ1ZQogIGdvY29uc3Q6CiAgICBtaW4tbGVuOiA1CiAgICBtaW4tb2NjdXJyZW5jZXM6IDQKICBtaXNzcGVsbDoKICAgIGxvY2FsZTogVVMKCmxpbnRlcnM6CiAgZGlzYWJsZS1hbGw6IHRydWUKICBlbmFibGU6CiAgICAtIGdvY29uc3QKICAgIC0gbWlzc3BlbGwKICAgIC0gZGVhZGNvZGUKICAgIC0gbWlzc3BlbGwKICAgIC0gc3RydWN0Y2hlY2sKICAgIC0gZXJyY2hlY2sKICAgIC0gdW51c2VkCiAgICAtIHZhcmNoZWNrCiAgICAtIHN0YXRpY2NoZWNrCiAgICAtIHVuY29udmVydAogICAgLSBnb2ZtdAogICAgLSBnb2ltcG9ydHMKICAgIC0gZ29saW50CiAgICAtIGluZWZmYXNzaWduCg=="
    },
    {
      "path": "./.goreleaser.yml",
      "name": ".goreleaser.yml",
      "content": "YmVmb3JlOgogIGhvb2tzOgogICAgLSBnbyBtb2QgZG93bmxvYWQKYnVpbGRzOgogIC0KICAgIGdvb3M6IFtsaW51eCwgZGFyd2luLCB3aW5kb3dzXQogICAgZ29hcmNoOiBbMzg2LCBhbWQ2NCwgYXJtLCBhcm02NF0KYXJjaGl2ZXM6CiAgLSB3cmFwX2luX2RpcmVjdG9yeTogdHJ1ZQpjaGVja3N1bToKICBuYW1lX3RlbXBsYXRlOiAnY2hlY2tzdW1zLnR4dCcKc25hcHNob3Q6CiAgbmFtZV90ZW1wbGF0ZTogInt7IC5UYWcgfX0tbmV4dCIKY2hhbmdlbG9nOgogIHNvcnQ6IGFzYwogIGZpbHRlcnM6CiAgICBleGNsdWRlOgogICAgLSAnXmRvY3M6JwogICAgLSAnXnRlc3Q6JwpicmV3OgogIG5hbWU6IGZzLWJ1bmRsZXIKICBnaXRodWI6CiAgICBvd25lcjogbW91bAogICAgbmFtZTogaG9tZWJyZXctbW91bAogIGNvbW1pdF9hdXRob3I6CiAgICBuYW1lOiBtb3VsLWJvdAogICAgZW1haWw6ICJtK2JvdEA0Mi5hbSIKICBob21lcGFnZTogaHR0cHM6Ly9tYW5mcmVkLmxpZmUvCiAgZGVzY3JpcHRpb246ICJmcy1idW5kbGVyIgo="
    }
  ]
}
```

```console
$ fs-bundler -f yaml *.json Makefile
files:
  - path: Makefile
    name: Makefile
    content: [71, 79, 32, 63, 61, 32, 103, 111, 10, 10, 46, 80, 72, 79, 78, 89, 58,
        32, 105, 110, 115, 116, 97, 108, 108, 10, 105, 110, 115, 116, 97, 108, 108,
        58, 10, 9, 36, 40, 71, 79, 41, 32, 105, 110, 115, 116, 97, 108, 108, 32, 46,
        10, 10, 46, 80, 72, 79, 78, 89, 58, 32, 116, 101, 115, 116, 10, 116, 101,
        115, 116, 58, 10, 9, 36, 40, 71, 79, 41, 32, 116, 101, 115, 116, 32, 45, 99,
        111, 118, 101, 114, 32, 45, 118, 32, 46, 47, 46, 46, 46, 10, 10, 46, 80, 72,
        79, 78, 89, 58, 32, 108, 105, 110, 116, 10, 108, 105, 110, 116, 58, 10, 9,
        103, 111, 108, 97, 110, 103, 99, 105, 45, 108, 105, 110, 116, 32, 114, 117,
        110, 32, 45, 45, 118, 101, 114, 98, 111, 115, 101, 32, 46, 47, 46, 46, 46,
        10, 10, 46, 80, 72, 79, 78, 89, 58, 32, 114, 101, 108, 101, 97, 115, 101,
        10, 114, 101, 108, 101, 97, 115, 101, 58, 10, 9, 103, 111, 114, 101, 108,
        101, 97, 115, 101, 114, 32, 45, 45, 115, 110, 97, 112, 115, 104, 111, 116,
        32, 45, 45, 115, 107, 105, 112, 45, 112, 117, 98, 108, 105, 115, 104, 32,
        45, 45, 114, 109, 45, 100, 105, 115, 116, 10, 9, 64, 101, 99, 104, 111, 32,
        45, 110, 32, 34, 68, 111, 32, 121, 111, 117, 32, 119, 97, 110, 116, 32, 116,
        111, 32, 114, 101, 108, 101, 97, 115, 101, 63, 32, 91, 121, 47, 78, 93, 32,
        34, 32, 38, 38, 32, 114, 101, 97, 100, 32, 97, 110, 115, 32, 38, 38, 32, 91,
        32, 36, 36, 123, 97, 110, 115, 58, 45, 78, 125, 32, 61, 32, 121, 32, 93, 10,
        9, 103, 111, 114, 101, 108, 101, 97, 115, 101, 114, 32, 45, 45, 114, 109,
        45, 100, 105, 115, 116, 10]
  - path: renovate.json
    name: renovate.json
    content: [123, 10, 32, 32, 34, 101, 120, 116, 101, 110, 100, 115, 34, 58, 32,
        91, 10, 32, 32, 32, 32, 34, 99, 111, 110, 102, 105, 103, 58, 98, 97, 115,
        101, 34, 10, 32, 32, 93, 44, 10, 32, 32, 34, 103, 114, 111, 117, 112, 78,
        97, 109, 101, 34, 58, 32, 34, 97, 108, 108, 34, 10, 125, 10]
```

## Usage

```console
$ fs-bundler -h
NAME:
   fs-bundler - A new cli application

USAGE:
   fs-bundler [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --format value, -f value  output format ("json", "yaml") (default: "json")
   --indent, -i              use indented output (only for "json" format) (default: false)
   --help, -h                show help (default: false)
   --version, -v             print the version (default: false)
```

## Install

```console
$ go get -u moul.io/fs-bundler
```

## License

© 2019 [Manfred Touron](https://manfred.life) -
[Apache-2.0 License](https://github.com/moul/fs-bundler/blob/master/LICENSE)
