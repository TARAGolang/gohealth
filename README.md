# gohealth

[![Build Status](https://travis-ci.org/smartdigits/gohealth.svg?branch=master)](https://travis-ci.org/smartdigits/gohealth)
[![Go report card](http://goreportcard.com/badge/smartdigits/gohealth)](https://goreportcard.com/report/smartdigits/gohealth)
[![GoDoc](https://godoc.org/github.com/smartdigits/gohealth?status.svg)](https://godoc.org/github.com/smartdigits/gohealth)

<sup>Tested for Go 1.5, 1.6, 1.7, tip</sup>

Monitoring and alarming for go servers in SmartDigits

<!-- MarkdownTOC autolink=true bracket=round depth=4 -->

- [Monitors](#monitors)
	- [Monitor example](#monitor-example)
- [Dependencies](#dependencies)
- [Testing](#testing)

<!-- /MarkdownTOC -->

## Monitors

A monitor is an object that is looking to a specific thing and eventually
can trigger one or several alarms.

All monitors implement the interface `Monitorer` with only two methods:

* `GetAlarms() []*Alarm` - Return all stored alarms. Alarms should be removed.
* `GetStatus() interface{}` - Return interesting status information, for example,
amount of memory and percentage of free memory. This method should be fast to
execute, with cached information (samples of whatever). This SHOULD NOT be slow
things like checking dabase connectivity or testing a backend web service.

### Monitor example

To illustrate a real monitor, we are going to use `MonitorBicounter`. This
monitor count "good" and "bad" events and will launch alarms when from last N
events the "bad" ones reach a limit.

We create the monitor:

```go
max_allowed_bad_events := 3
memorize_last_events := 100
m := NewMonitorBicounter(max_allowed_bad_events, memorize_last_events)
```

Then some events are injected:

```go
m.Ok()
m.Ok()
m.Error()
m.Error()
m.Error()
m.Ok()
m.Ok()
m.Error() // this event will generate alarm
m.Ok() // following Ok generate alarms until last 100 errors are under 4
m.Ok()
m.Ok()
```

Lets see some alarms:

```go
for _,alarm := range m.GetAlarms() {
	fmt.Println(alarm.Msg)
}
```

The output:

```
Limit at 4 out of 100
Limit at 4 out of 100
Limit at 4 out of 100
```



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
