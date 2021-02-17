package mvn

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/solo-io/go-utils/log"
	osu "github.com/solo-io/go-utils/osutils"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Pom interface {
	//Returns an error Or
	//The n
	UpVersionAutomatic(wg *sync.WaitGroup) (string, error)
}

type pom struct {
	osClient osu.OsClient
	Config   PomConfig
}

func New(config PomConfig) Pom {
	return &pom{
		osClient: osu.NewOsClient(),
		Config:   config,
	}
}

//Read the file and returns a XML node
func (c *pom) readFile(filePath string) (*xmlquery.Node, error) {
	content, err := c.osClient.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var doc *xmlquery.Node
	doc, err = xmlquery.Parse(bytes.NewReader(content))
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (c *pom) UpVersion(filePath string) (string, error) {
	doc, err := c.readFile(filePath)

	log.Debugf("doc %s", doc.InnerText())

	return "", err
}

// Updates the patch version only
// Will remove 'SNAPSHOT' optional update patch version and keep SNAPSHOT
// Returns the new version
func (c *pom) UpVersionAutomatic(wg *sync.WaitGroup) (string, error) {
	defer wg.Done()
	con := c.Config
	doc, err := c.readFile(con.File)
	if err != nil {
		return "", err
	}

	version := xmlquery.FindOne(doc, "//version")
	cVersion := version.InnerText()
	log.Debugf("Current version %s", cVersion)

	all := strings.Split(cVersion, ".")
	major := all[0]
	minor := all[1]
	patch := all[2]

	var npatch string
	if strings.Contains(patch, "SNAPSHOT") && con.KeepSnapshot {
		list := strings.Split(patch, "-")
		patch = list[0]

		p, err := strconv.Atoi(patch)
		if err != nil {
			return "", err
		}
		p = p + 1
		npatch = fmt.Sprint(p)
	}

	if strings.Contains(patch, "SNAPSHOT") && !con.KeepSnapshot {
		list := strings.Split(patch, "-")
		npatch = list[0]
	}

	if !strings.Contains(patch, "SNAPSHOT") {
		p, err := strconv.Atoi(patch)
		if err != nil {
			return "", err
		}
		p = p + 1
		npatch = fmt.Sprint(p)

	}
	newVersion := major + "." + minor + "." + npatch
	if con.KeepSnapshot {
		newVersion = newVersion + "-SNAPSHOT"
	}

	file, err := os.Create("./pom.xml-new " + newVersion)
	if err != nil {
		return "", err
	}

	version.FirstChild.Data = newVersion

	writer := bufio.NewWriter(file)
	writer.WriteString(doc.OutputXML(false))
	writer.Flush()

	return newVersion, nil
}

//Update semantic version
//func updateVersion(currVersion string, keepSnapshot bool, postVersion string)  (string, error){
//
//
//	all := strings.Split(currVersion, ".")
//	major := all[0]
//	minor := all[1]
//	patch := all[2]
//	if strings.Contains(patch, "SNAPSHOT") {
//		list := strings.Split(patch, "-")
//		patch = list[0]
//	}
//
//	version := major + minor + b
//
//	m,err := strconv.Atoi(major)
//	i := 0
//	i, err = strconv.Atoi(minor)
//	b := 0
//	b, err = strconv.Atoi(bug)
//
//	if err != nil{
//		return "", err
//	}

//}
