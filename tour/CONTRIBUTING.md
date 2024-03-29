# Contributing to Go

Go is an open source project.

It is the work of hundreds of contributors. We appreciate your help!


## Filing issues

When [filing an issue](https://golang.org/issue/new), make sure to answer these five questions:

1. What version of Go are you using (`go version`)?
2. What operating system and processor architecture are you using?
3. What did you do?
4. What did you expect to see?
5. What did you see instead?

General questions should go to the [golang-nuts mailing list](https://groups.google.com/group/golang-nuts) instead of the issue tracker.
The gophers there will answer or ask you to file an issue if you've tripped over a bug.

## Verifying changes during development

In order to verify changes to the slides or code examples while developing
locally compile with your local toolchain:

	$ go install github.com/gophertrain/tour/gotour
	$ $GOPATH/bin/gotour

## Running the App Engine version locally

To view the App Engine version of the slides while developing locally, install
the [Go App Engine SDK](https://cloud.google.com/appengine/downloads?hl=en)
and then:

	$ cd $GOPATH/src/golang.org/x/tour
	$ $SDK_PATH/dev_appserver.py .

The App Engine version runs code examples against the service at play.golang.org.
To verify changes to the code examples you must use your local toolchain to compile
and run `gotour` locally.

## Contributing code

Please read the [Contribution Guidelines](https://golang.org/doc/contribute.html)
before sending patches.

**We do not accept GitHub pull requests**
(we use [Gerrit](https://code.google.com/p/gerrit/) instead for code review).

Unless otherwise noted, the Go source files are distributed under
the BSD-style license found in the LICENSE file.

