How to install?
===============

```
$ psql -c 'CREATE DATABASE "upwork"' -d postgres
$ mkdir upwork
$ cd upwork
$ git clone --recursive git@github.com:mahendrakalkura/upwork.git .
$ cp settings.toml.sample settings.toml
$ go get
```

How to run?
===========

```
$ cd upwork
$ go build
$ ./upwork --action=bootstrap
$ ./upwork --action=categories
$ ./upwork --action=skills
$ ./upwork --action=jobs
$ ./upwork --action=ui
```
