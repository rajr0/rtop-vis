/*

rtop-vis - ad hoc cluster monitoring over SSH

Copyright (c) 2015 RapidLoop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
        "html/template"
        "log"
        "net/http"
	"os"
	"io/ioutil"
)

var tmpl *template.Template

func startWeb() {
	file, err := os.Open("rtop-vis.html")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
        tmpl = template.Must(template.New(".").Parse(string(data)))
        http.HandleFunc("/", webServer)
        log.Fatal(http.ListenAndServe(DEFAULT_WEB_ADDR, nil))
}

func webServer(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.Execute(w, allStats); err != nil {
		log.Print(err)
	}
}
