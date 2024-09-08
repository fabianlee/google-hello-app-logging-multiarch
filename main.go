/**
 * Copyright 2021 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// [START gke_hello_app]
// [START container_hello_app]
package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)
// ldflags passed via Docker build-arg
var Version = "n/a"

// wait time between logging messages
const SECONDS_SLEEP = 10

var logType = "json"
var loopIndex int64 = 0
var whoAmI string = "world"

// for unstructured output
var stdoutLog = log.New(os.Stdout, "", 1)
var stderrLog = log.New(os.Stderr, "", 1)

// for structured output
// slog by default sends to stderr, switching to stdout
var line_logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
var json_logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func main() {

	// override structured log output? (default=json)
	if os.Getenv("logType") != "" {
	   logType = os.Getenv("logType")
	}
	if logType == "json" {
          slog.SetDefault(json_logger)
	}else {
	  slog.SetDefault(line_logger)
	}

	// any override?
	if os.Getenv("whoAmI") != "" {
	   whoAmI = os.Getenv("whoAmI")
	}
        // send log messages once a second
	go runDataLoop()

	// register hello function to handle all web requests
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	// use PORT environment variable, or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// start the web server on port and accept requests
	log.Printf("Server %s listening on port %s", Version, port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

// hello responds to the request with a plain-text "Hello, world" message.
func hello(w http.ResponseWriter, r *http.Request) {
        // send request info to stdout and stderr as test of unstructured
	stdoutLog.Printf("stdout logging request: %s\n",r.URL.Path)
	stderrLog.Printf("stderr Serving request: %s\n",r.URL.Path)

	// web response to end user
	fmt.Fprintf(w, "Hello, %s!\n",whoAmI)
}

// infinite loop, log structured messages at different levels every 10 seconds
func runDataLoop() {
    for {
        loopIndex++
	//slog.Info("logging at info level", slog.Int64("loopIndex",loopIndex), slog.String("whoAmI",whoAmI))
	slog.Info("logging at info level", "severity","INFO", "loopIndex",loopIndex, "whoAmI",whoAmI)
	slog.Warn("logging at info level", "severity","WARN", "loopIndex",loopIndex, "whoAmI",whoAmI)
	slog.Error("logging at info level", "severity","ERROR", "loopIndex",loopIndex, "whoAmI",whoAmI)
        time.Sleep(SECONDS_SLEEP * time.Second)
    }
}

// [END container_hello_app]
// [END gke_hello_app]
