# IOCee
### Extract potential IOCs from unstructured text

[![GoDoc](https://godoc.org/github.com/DCSO/iocee?status.svg)](http://godoc.org/github.com/DCSO/iocee)
[![Build Status](https://travis-ci.org/DCSO/iocee.svg?branch=master)](https://travis-ci.org/DCSO/iocee)

IOCee is a simple tool that reads a stream of text from standard input and searches for potential IOC values in it. When it encounters a value that could be an IOC (e.g. a domain name, IP address, URL or hash) it generates all potential variations of it and prints each of them on a separate line. The output data can then be used with other tools such as `bloom`, which expect one value per line.

# Usage

To extract IOCs from a file:

    cat filename.txt | iocee

To interactively extract IOCs (useful for testing):

    iocee --interactive

# Installation & Usage

To install the command line tool:

    make install

To run the tests:

    make test
