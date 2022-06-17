buntrepair
==========

These are data recovery and debugging tools for [BuntDB](https://github.com/tidwall/buntdb).

At this point, two warnings are salient:

* BACK UP YOUR DATABASE BEFORE USING ANY OF THESE TOOLS.
* DO NOT USE THESE TOOLS ON A DATABASE IN USE BY ANOTHER PROCESS.

These tools may destructively modify the data they operate on, even if that doesn't seem possible.

Building
--------

Obtain an [up-to-date distribution of the Go language for your OS and architecture](https://golang.org/dl/). Then run `make`, which will build executables named `buntrepair` and `buntdump`.

buntrepair
----------

`buntrepair` attempts to import a database, skipping over any damaged or uninterpretable sections that would otherwise be fatal errors. This involves fallible heuristics and could lead to arbitrary data loss, or possibly even associating keys with the wrong values. You've been warned. Then it uses [(*DB).Shrink](https://pkg.go.dev/github.com/tidwall/buntdb#DB.Shrink) to produce a syntactically valid database that can be read by normal BuntDB-based applications.

To use it, MAKE A SEPARATE COPY OF YOUR DATABASE, then run `./buntrepair copy.db`.

buntdump
--------

`buntdump` prints out all the key-value pairs from a database. This tool also risks destructively modifying the database, or losing data, so back up your database first and don't use it on a running database. You may also want to build `exe/buntdump.go` against an unmodified version of the BuntDB library that errors out properly when encountering corrupt sections.

To use it, MAKE A SEPARATE COPY OF YOUR DATABASE, then run `./buntdump copy.db`.
