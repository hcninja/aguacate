// Copyright [2020] Jose Gonzalez-Krause <contact@hackercat.ninja>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"os"

	"github.com/fatih/color"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("")
	log.SetOutput(os.Stdout)
}

var (
	yellow = color.New(color.FgYellow).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()

	remarkText = color.New(color.FgMagenta, color.Bold).SprintFunc()
	info       = color.New(color.FgWhite, color.BgGreen).SprintFunc()
	warn       = color.New(color.FgRed, color.Bold).SprintFunc()
	errp       = color.New(color.FgWhite, color.BgRed).SprintFunc()
)

func fatalErr(err error) {
	if err != nil {
		log.Fatalf("[!] Fatal: %s", err.Error())
	}
}

func errorErr(err error) {
	if err != nil {
		log.Printf("[!] Error: %s", err.Error())
	}
}
