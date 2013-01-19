package main

import (
	"net/http"
	"os"
	"path"
	"strconv"
)

var rootpath string

// Serve starts an http server using net/http which listens on the given
// interface and port. It serves files using the given path as its root
// directory. If iface is blank, 0.0.0.0 is assumed by net/http.
//
// In order to pass information to the handler functions, it sets the
// rootpath variable. This means that invoking this function a second
// time will impair the behavior of the first call.
//
// This function will block until the http server encounters an error or
// otherwise stops.
func Serve(root, iface string, port int) (err error) {
	// Root the path, and clean it if necessary.

	// 18/01/2013 It might make sense to move this to a helper routine
	// or further up in the stack.
	if !path.IsAbs(root) {
		var wd string
		wd, err = os.Getwd()
		if err != nil {
			return
		}
		root = path.Join(wd, root)
	} else {
		root = path.Clean(root)
	}
	l.Printf("Starting http server %s:%d\nRoot path %q", iface, port, root)
	rootpath = root

	http.HandleFunc("/p/", handle)
	http.ListenAndServe(iface+":"+strconv.Itoa(port), nil)
	return
}

// handle is the "maintainence" handler. It manages requests (such as
// for the favicon) and may, in the future, redirect project requests to
// the handleProject() handler.
func handle(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, Forest."))
}

// handleProject is the handler for project requests. It uses the
// variable rootpath to determine where to serve content from.
func handleProject(w http.ResponseWriter, req *http.Request) {
	// 19/01/2013 I hope this doesn't turn into a god function.
	w.Write([]byte("This is a project directory."))
}
