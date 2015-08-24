# Gotem - a simple HTML template tool #

Gotem is a very simple static HTML template tool. It allows you to
include other files in your HTML document. For example, you might have
a common header, footer and menu that you want to include on every
page. Gotem allows you to put these elements in their own files, and
then include them from another file.

## Usage ##

Put all of the elements to include into the same directory, containing
one file for each element ending in `.template`. For example, you
could create a directory called `includes`, and then create a file in
there called `html_start.template` containing:

    <html>
    <body>

And a file called `html_end.template` in the same directory containing:

    </body>
    </html>

Then create your main HTML document in a file called `index.template`
containing:

    {{include "html_start"}}
    <h1>My HTML contents goes here</h1>
    {{include "html_end"}}

Then you can compile the page to `index.html` by running:

    gotem -I includes index.template index.html

And that's it.

If the output file is missing or `-` then the compiled HTML file will
be printed to stdout. If the input file is missing or `-` then the
template will be read from stdin. You can run gotem without `-I` but
then you won't be able to include anything, so it might be kinda
pointless.

(Behind the scenes, gotem uses Go's templating system so maybe you can
do something fancier with the templates.)

## Installing ##

Binaries for Mac and Linux are available from the
[https://github.com/jcinnamond/gotem/releases/tag/v1.0.0](releases
page).

If you have [http://golang.org](Go) installed you can install from
source by running:

    go get github.com/jcinnamond/gotem

## Licence ##

The MIT License (MIT)

Copyright (c) 2015 John Cinnamond

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
