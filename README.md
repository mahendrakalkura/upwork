How to install?
===============

```
$ mkdir upwork
$ cd upwork
$ git clone --recursive git@github.com:mahendrakalkura/upwork.git .
$ cp config.json.sample config.json
$ go get
```

How to run?
===========

```
$ cd upwork
$ go build
$ ./upwork --action=categories | jq
$ ./upwork --action=jobs       | jq
```
