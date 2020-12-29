# Redsync

WARNING: THIS IS FORK OF ORIGINAL REDSYNC (https://github.com/go-redsync/redsync) THAT MAY OUT OF DATE

This version doesn't have hard dependenct to go-redis. The reason to do this is to remove protobuf dependency in source code. See original repo for additional info and latest update.

Redsync provides a Redis-based distributed mutual exclusion lock implementation for Go as described in [this post](http://redis.io/topics/distlock). A reference library (by [antirez](https://github.com/antirez)) for Ruby is available at [github.com/antirez/redlock-rb](https://github.com/antirez/redlock-rb).

## Installation

Install Redsync using the go get command:

    $ go get github.com/ridwanadhip/redsync/v4

There is one driver that supported by this library:

 * [Redigo](https://github.com/gomodule/redigo)

See the [examples](examples) folder for usage of each driver.

## Documentation

- [Reference](https://godoc.org/github.com/go-redsync/redsync)

## Contributing

Contributions are welcome.

## License

Redsync is available under the [BSD (3-Clause) License](https://opensource.org/licenses/BSD-3-Clause).

## Disclaimer

This code implements an algorithm which is currently a proposal, it was not formally analyzed. Make sure to understand how it works before using it in production environments.
