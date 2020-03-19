package main

import (
	"fmt"
	"os/exec"
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
)

// Chart struct fields must be public in order for unmarshal to
// correctly populate the data.
type Chart struct{
	APIVersion string
	Name string
	Version string
	AppVersion string
	Description string
	Keywords [3]string
	Home string
	Icon string
	Sources [1]string
	Maintainers []Maintainer
	Engine string
}

//Maintainer stuff
type Maintainer struct{
	Name string
	Email string
}

func main() {
	var chart Chart
	chartName := os.Args[1]
	filePath := fmt.Sprintf("%v/Chart.yaml", chartName)
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal([]byte(fileData), &chart)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for major := 0; major < 5; major++ {
		for minor := 0; minor < 5; minor++{
			for patch := 1; patch < 10; patch++{
				newVersion := fmt.Sprintf("%v.%v.%v", major, minor, patch)
				chart.Modify(newVersion)
				newData, err := yaml.Marshal(&chart)
				if err != nil {
					log.Fatalf("error: %v", err)
				}
				err = ioutil.WriteFile(filePath, []byte(string(newData)), 0644)
				if err != nil {
					log.Fatalf("error: %v", err)
				}
				cmd := exec.Command("helm", "package", chartName)
				cmd.Run()
				packageName := fmt.Sprintf("@%v-%v.tgz", chartName, newVersion)
				cmd = exec.Command("curl", "--data-binary", packageName, "http://localhost:8080/api/charts")
				cmd.Run()
			}
		}
	}
}

// Modify is good
func (c *Chart) Modify(version string) {
  c.Version = version
}
