# rocksgo

rocksgo is a golang wrapper for rocksdb.

The API has been godoc'ed and [is available on the
web](http://godoc.org/github.com/ananclub/rocksgo).



## Building

You'll need the shared library build of
[rocksdb](http://github.com/facebook/rocksdb/) installed on your machine. The
current rocksdb will build it by default.


Now, if you build rocksdb and put the shared library and headers in one of the
standard places for your OS, you'll be able to simply run:

    go get github.com/ananclub/rocksgo

But, suppose you put the shared rocksdb library somewhere weird like
/path/to/lib and the headers were installed in /path/to/include. To install
rocksgo remotely, you'll run:

    CGO_CFLAGS="-I/path/to/rocksdb/include " CGO_LDFLAGS="-L/path/to/rocksdb/lib -lrocksdb -lstdc++ -lz -lrt" go get github.com/ananclub/rocksgo
and there you go.


Of course, these same rules apply when doing `go build`, as well.

## Caveats

Comparators and WriteBatch iterators must be written in C in your own
library. This seems like a pain in the ass, but remember that you'll have the
rocksdb C API available to your in your client package when you import rocksgo.

