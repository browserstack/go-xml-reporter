// Contents of junit.go

package junit

import "encoding/xml"

// TestSuites represents the top-level <testsuites> element in JUnit XML.
type TestSuites struct {
	XMLName    xml.Name    `xml:"testsuites"`
	Text       string      `xml:",chardata"`
	Name       string      `xml:"name,attr,omitempty"`
	Time       string      `xml:"time,attr,omitempty"`
	Tests      string      `xml:"tests,attr,omitempty"`
	Failures   string      `xml:"failures,attr,omitempty"`
	Errors     string      `xml:"errors,attr,omitempty"`
	Skipped    string      `xml:"skipped,attr,omitempty"`
	Assertions string      `xml:"assertions,attr,omitempty"`
	TestSuites []TestSuite `xml:"testsuite"`
}

// TestSuite represents the <testsuite> element in JUnit XML.
type TestSuite struct {
	Name       string     `xml:"name,attr"`
	Tests      string     `xml:"tests,attr,omitempty"`
	Failures   string     `xml:"failures,attr,omitempty"`
	Errors     string     `xml:"errors,attr,omitempty"`
	Skipped    string     `xml:"skipped,attr,omitempty"`
	Assertions string     `xml:"assertions,attr,omitempty"`
	Time       string     `xml:"time,attr,omitempty"`
	Timestamp  string     `xml:"timestamp,attr,omitempty"`
	File       string     `xml:"file,attr"`
	Text       string     `xml:",chardata"`
	HostName   string     `xml:"hostname,attr,omitempty"`
	TestCases  []TestCase `xml:"testcase"`
}

// TestCase represents the <testcase> element in JUnit XML.
type TestCase struct {
	Id         string      `xml:"-"`
	Name       string      `xml:"name,attr"`
	ClassName  string      `xml:"classname,attr"`
	Assertions string      `xml:"assertions,attr,omitempty"`
	Time       string      `xml:"time,attr"`
	File       string      `xml:"file,attr"`
	Result     string      `xml:"result,attr,omitempty"`
	Failure    *Failure    `xml:"failure,omitempty"`
	Skipped    *Skipped    `xml:"skipped,omitempty"`
	Error      *Error      `xml:"error,omitempty"`
	ErrorLog   *ErrorLog   `xml:"system-err,omitempty"`
	ConsoleLog *ConsoleLog `xml:"system-out,omitempty"`
	Properties *Properties `xml:"properties,omitempty"`
}

// Failure represents the <failure> element in JUnit XML.
type Failure struct {
	Message string `xml:"message,attr,omitempty"`
	Type    string `xml:"type,attr,omitempty"`
	Content string `xml:",chardata"`
}

type Skipped struct {
	Message string `xml:"message,attr"`
}

type Error struct {
	Message string `xml:"message,attr,omitempty"`
	Type    string `xml:"type,attr,omitempty"`
	Content string `xml:",chardata"`
}

type ErrorLog struct {
	Content string `xml:",chardata"`
}

type ConsoleLog struct {
	Content string `xml:",chardata"`
}

type Properties struct {
	Properties []Property `xml:"property"`
}

type Property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type BuildDetailsParams struct {
	username             string
	password             string
	buildName            string
	projectName          string
	buildIdentifier      string
	buildTags            []string
	ciDetails            string
	frameworkDetails     string
	vcsDetails           map[string]string
	additionalProperties map[string]string
}
