package observabilityReporter

import (
	"fmt"
)

// Reporter defines the interface for generating JUnit XML reports.
type Reporter interface {
	CreateBuild(username string, password string, buildName string, projectName string, buildIdentifier string, buildTags []string, ciDetails string, frameworkDetails string, vcsDetails map[string]string, additionalProperties map[string]string) (string, error)
	AddTestRun(name string, result string, errorTrace string, startTime string, duration string, parentSuite string, className string, filePath string, testProps map[string]interface{}, logs string, additionalProps map[string]string) (string, error)
	UpdateBuildDetails(buildTags []string, ciDetails string, frameworkDetails string, vcsDetails map[string]string, additionalProperties map[string]string) (string, error)
	UpdateTestRunDetails(testIdentifier string, testProps map[string]interface{}, logs string, additionalProps map[string]string) (string, error)
	SendReport() (string, error)
	StopBuild() (string, error)
}

// JUnitReporter implements the Reporter interface.
type JUnitReporter struct {
	buildDetails BuildDetailsParams
	testSuites   map[string]TestSuite
}

func NewJUnitReporter() *JUnitReporter {
	return &JUnitReporter{
		testSuites: make(map[string]TestSuite),
	}
}

func (jr *JUnitReporter) resetData() {

	removeBuildDirErr := removeBuildAssets(jr.buildDetails.buildIdentifier)
	if removeBuildDirErr != nil {
		fmt.Println("xml build directory deletion failed")
	}

	// Resetting buildDetails data
	jr.buildDetails.username = ""
	jr.buildDetails.password = ""
	jr.buildDetails.buildName = ""
	jr.buildDetails.projectName = ""
	jr.buildDetails.buildIdentifier = ""
	jr.buildDetails.buildTags = []string{}
	jr.buildDetails.ciDetails = ""
	jr.buildDetails.frameworkDetails = ""
	jr.buildDetails.vcsDetails = make(map[string]string)
	jr.buildDetails.additionalProperties = make(map[string]string)

	// Resetting testSuites data
	jr.testSuites = make(map[string]TestSuite)
}

func (jr *JUnitReporter) resetTestSuites() {

	removeBuildDirErr := removeBuildAssets(jr.buildDetails.buildIdentifier)
	if removeBuildDirErr != nil {
		fmt.Println("xml build directory deletion failed")
	}

	// Resetting testSuites data
	jr.testSuites = make(map[string]TestSuite)
}
