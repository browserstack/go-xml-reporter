package observabilityReporter

import (
	"fmt"
)

// Reporter defines the interface for generating JUnit XML reports.
type Reporter interface {
	AddTestRun(name string, result string, errorTrace string, startTime string, duration string, parentSuite string, className string, filePath string, testProps map[string]interface{}, logs string, additionalProps map[string]string) (string, error)
	UpdateBuildDetails(buildTags []string, ciDetails string, frameworkDetails string, vcsDetails map[string]string, additionalProperties map[string]string) (string, error)
	UpdateTestRunDetails(testIdentifier string, testProps map[string]interface{}, logs string, additionalProps map[string]string) (string, error)
	SendReport() (string, error)
	StopBuild() (string, error)
}

// JUnitReporter implements the Reporter interface.
type JUnitReporter struct {
	BuildIdentifier string
	buildDetails    BuildDetailsParams
	testSuites      map[string]TestSuite
}

func CreateBuild(username string, password string, buildName string, projectName string, buildIdentifier string, buildTags []string, ciDetails string, frameworkDetails string, vcsDetails map[string]string, additionalProperties map[string]string) (*JUnitReporter, error) {

	// Set build identifier if provided, otherwise generate it
	if !stringValidator(buildIdentifier) {
		buildIdentifier = generateUUID()
	}

	var buildDetailsObject BuildDetailsParams

	buildDetailsObject.username = username
	buildDetailsObject.password = password
	buildDetailsObject.buildName = buildName
	buildDetailsObject.projectName = projectName
	buildDetailsObject.buildIdentifier = buildIdentifier
	buildDetailsObject.buildTags = buildTags
	buildDetailsObject.ciDetails = ciDetails
	buildDetailsObject.frameworkDetails = frameworkDetails
	buildDetailsObject.vcsDetails = vcsDetails
	buildDetailsObject.additionalProperties = additionalProperties

	// validation for required params
	validationError := BuildDetailsValidator(buildDetailsObject)
	if validationError != nil {
		return nil, validationError
	}

	// Create a folder to store xml files and attachments
	createBuildDirErr := createBuildDirectory(buildIdentifier)
	if createBuildDirErr != nil {
		removeBuildAssets(buildIdentifier)
	}

	return &JUnitReporter{
		testSuites:      make(map[string]TestSuite),
		buildDetails:    buildDetailsObject,
		BuildIdentifier: buildIdentifier,
	}, nil
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
