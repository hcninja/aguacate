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
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/hcninja/aguacate/nmap"
)

var (
	version      = "0.1.0"
	basePathFlag = flag.String("path", "./", "Base path for project structure")
	nmapFileFlag = flag.String("nmap", "", "Nmap .xml scan result")
	versionFlag  = flag.Bool("version", false, "Shows the version")
)

func main() {
	flag.Parse()

	log.Printf("Preparing some nice guacamole 🥑 v%s", version)

	if *nmapFileFlag == "" {
		log.Fatal("❗️ '-nmap' flag is mandatory")
	}

	fBuffer, err := ioutil.ReadFile(*nmapFileFlag)
	if err != nil {
		errorErr(err)
	}

	nm, err := nmap.Parse(fBuffer)
	if err != nil {
		errorErr(err)
	}

	stat, err := os.Stat(*basePathFlag)
	if err != nil {
		if err != os.ErrNotExist {
			log.Fatal("❗️ The directory does not exist")
		}
		log.Fatal(err)
	}

	path := *basePathFlag
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	if !stat.IsDir() {
		log.Fatal("❗️ Path is not a directory")
	}

	log.Println("Base path: ", *basePathFlag)

	log.Printf("Hosts: %d", len(nm.Hosts))

	for _, host := range nm.Hosts {
		var data mdData
		if len(host.Addresses) >= 1 {
			data.Address = host.Addresses[0].Addr
		} else {
			continue
		}

		log.Printf("Generating folder for %s ✅", host.Addresses[0].Addr)

		data.ScanTime = host.StartTime.String()

		if len(host.Os.OsMatches) >= 1 {
			data.OsName = host.Os.OsMatches[0].Name
			data.OsAccuracy = host.Os.OsMatches[0].Accuracy
		}

		data.Services = host.Ports

		md := generateMD(data)

		assetPath := "./" + path + data.Address + "/"

		if err := os.Mkdir(assetPath, 0700); err != nil {
			log.Fatal(err)
		}

		if err := ioutil.WriteFile(assetPath+data.Address+".md", []byte(md), 0600); err != nil {
			log.Fatal(err)
		}
	}
}

func generateMD(host mdData) string {
	buf := new(bytes.Buffer)
	t := template.Must(template.New("md").Parse(mdTmpl))
	t.Execute(buf, host)

	return buf.String()
}

type mdData struct {
	Address    string
	ScanTime   string
	OsName     string
	OsAccuracy string
	Services   []nmap.Port
}

var mdTmpl = `# {{.Address}}:
- Scanned on: {{.ScanTime}}
- OS: {{.OsName}} ({{.OsAccuracy}}%)

## Nmap:
{{- range .Services}}
- {{.PortId}} {{.State.State}} {{.Service.Name}} [{{.Service.Product}} ({{.Service.Version}})]
{{- end}}

`
