# gohealth

[![Build Status](https://travis-ci.org/smartdigits/gohealth.svg?branch=master)](https://travis-ci.org/smartdigits/gohealth)
[![Go report card](http://goreportcard.com/badge/smartdigits/gohealth)](https://goreportcard.com/report/smartdigits/gohealth)
[![GoDoc](https://godoc.org/github.com/smartdigits/gohealth?status.svg)](https://godoc.org/github.com/smartdigits/gohealth)

<sup>Tested for Go 1.5, 1.6, 1.7, tip</sup>

Monitoring and alarming for go servers in SmartDigits

<!-- MarkdownTOC autolink=true bracket=round depth=4 -->

- [Dependencies](#dependencies)
- [Testing](#testing)

<!-- /MarkdownTOC -->

## Dependencies

Dependencies for testing are:

* github.com/fulldump/golax

Transitive dependencies for runtime are:

* github.com/fulldump/golax [optional]
* gopkg.in/mgo.v2 [optional]


## Testing

As simple as:

```sh
git clone "<this-repo>"
make setup && make dependencies
make test
```
