package mvn

import (
	"github.com/antchfx/xmlquery"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestXMLParser(t *testing.T) {
	s := `<?xml version="1.0" encoding="utf-8"?><project version="2.0"><version>1.0.0</version></project>`
	doc, err := xmlquery.Parse(strings.NewReader(s))

	if err != nil {
		t.Fail()
	}

	version := xmlquery.FindOne(doc, "//version")

	t.Logf("%s", version.InnerText())
}

func TestUpPomVersion(t *testing.T) {

	conf := PomConfig{
		File:         "./pom_test.xml",
		Automatic:    true,
		KeepSnapshot: false,
		Mayor:        "",
		Minor:        "",
		Patch:        "",
	}
	ci := New(conf)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	cVersion, err := ci.UpVersionAutomatic(wg)
	wg.Wait()
	if err != nil {
		t.Fail()
	}

	t.Logf("New Version %s", cVersion)

	if cVersion != "1.0.1" {
		t.Fatalf("Expected 1.0.1 got %s", cVersion)
	}

	err = os.Remove("pom.xml-new 1.0.1")
	if err != nil {
		t.Fail()
	}
}

func TestUpPomSnapshotVersion(t *testing.T) {

	conf := PomConfig{
		File:         "./pom_test1.xml",
		Automatic:    true,
		KeepSnapshot: false,
		Mayor:        "",
		Minor:        "",
		Patch:        "",
	}

	ci := New(conf)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	cVersion, err := ci.UpVersionAutomatic(wg)
	wg.Wait()

	if err != nil {
		t.Fail()
	}

	t.Logf("New Version %s", cVersion)

	if cVersion != "1.0.0" {
		t.Fatalf("Expected 1.0.0 got %s", cVersion)
	}

	err = os.Remove("pom.xml-new 1.0.0")
	if err != nil {
		t.Fail()
	}
}

func TestUpPomKeepSnapshotVersion(t *testing.T) {

	conf := PomConfig{
		File:         "./pom_test1.xml",
		Automatic:    true,
		KeepSnapshot: true,
		Mayor:        "",
		Minor:        "",
		Patch:        "",
	}

	ci := New(conf)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	cVersion, err := ci.UpVersionAutomatic(wg)
	wg.Wait()

	if err != nil {
		t.Fail()
	}

	t.Logf("New Version %s", cVersion)

	if cVersion != "1.0.1-SNAPSHOT" {
		t.Fatalf("Expected 1.0.2-SNAPSHOT got %s", cVersion)
	}

	err = os.Remove("pom.xml-new 1.0.1-SNAPSHOT")
	if err != nil {
		t.Fail()
	}
}
