Gomad
-----
An experiment in monads.

Gomad includes a collection of basic "generic" monads (ones that wrap `interface{}` values).

It also includes a rewriter. This rewriter takes a file (like `maybe/maybe.go`), and produces a type-specific one:

````
$ go run main.go -file maybe/maybe.go -type int
package maybe_int

type Maybe interface {
        Bind(func(v int) Maybe) Maybe

        Otherwise(v int) int
}

type Nothing struct{}

func (n Nothing) Bind(f func(v int) Maybe) Maybe {
        return n
}

func (n Nothing) Otherwise(v int) int {
        return v
}

type Just struct {
        Value int
}

func (j Just) Bind(f func(v int) Maybe) Maybe {
        return f(j.Value)
}

func (j Just) Otherwise(v int) int {
        return j.Value
}
````

[![Build Status](https://secure.travis-ci.org/frankshearar/gomad.png?branch=master)](http://travis-ci.org/frankshearar/gomad)


Licence
-------

Copyright (C) 2012-2013 by Frank Shearar

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
