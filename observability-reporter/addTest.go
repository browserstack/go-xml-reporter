package observabilityReporter

import (
	"errors"
)

func (jr *JUnitReporter) AddTestRun(buildIdentifier string, name string, result string, errorTrace string, startTime string, duration string, parentSuite string, className string, filePath string, testProps map[string]interface{}, logs string, additionalProps map[string]string) (string, error) {

	// If build identifier is not matching with current object stopping execution
	if jr.buildDetails.buildIdentifier != buildIdentifier {
		return "", errors.New("build identifier is not matching")
	}

	testDetailsValidatorObject := map[string]string{
		"buildIdentifier": buildIdentifier,
		"name":            name,
		"result":          result,
		"startTime":       startTime,
		"duration":        duration,
		"className":       className,
		"errorTrace":      errorTrace,
	}

	// validation for required params
	validationError := AddTestValidator(testDetailsValidatorObject)
	if validationError != nil {
		return "", validationError
	}

	testIdentifier := generateUUID()

	if !stringValidator(parentSuite) {
		// className & parentSuite both
		matchedTestSuite := jr.testSuites[parentSuite]
		if stringValidator(matchedTestSuite.Name) {
			matchedTestSuite.TestCases = append(matchedTestSuite.TestCases, generateNewTestCase(testIdentifier, name, className, duration, filePath, result, errorTrace, logs, testProps, jr.buildDetails.buildIdentifier))
			jr.testSuites[parentSuite] = matchedTestSuite
		} else {
			newTestSuite := generateNewTestSuite(parentSuite, filePath)
			newTestSuite.TestCases = []TestCase{
				generateNewTestCase(testIdentifier, name, className, duration, filePath, result, errorTrace, logs, testProps, jr.buildDetails.buildIdentifier),
			}
			jr.testSuites[parentSuite] = newTestSuite
		}
	} else {
		// className only
		matchedTestSuite := jr.testSuites[className]
		if stringValidator(matchedTestSuite.Name) {
			matchedTestSuite.TestCases = append(matchedTestSuite.TestCases, generateNewTestCase(testIdentifier, name, className, duration, filePath, result, errorTrace, logs, testProps, jr.buildDetails.buildIdentifier))
			jr.testSuites[className] = matchedTestSuite
		} else {
			newTestSuite := generateNewTestSuite(className, filePath)
			newTestSuite.TestCases = []TestCase{
				generateNewTestCase(testIdentifier, name, className, duration, filePath, result, errorTrace, logs, testProps, jr.buildDetails.buildIdentifier),
			}
			jr.testSuites[className] = newTestSuite
		}
	}

	return testIdentifier, nil // TODO: return test identifier
}

func generateNewTestSuite(parentSuite string, filePath string) TestSuite {
	var newTestSuite TestSuite
	newTestSuite.Name = parentSuite
	newTestSuite.File = filePath

	return newTestSuite
}

func generateNewTestCase(
	testIdentifier string,
	name string,
	className string,
	duration string,
	filePath string,
	result string,
	errorTrace string,
	logs string,
	testProps map[string]interface{},
	buildIdentifier string) TestCase {

	newTestCase := TestCase{}
	newTestCase.Id = testIdentifier
	newTestCase.Name = name
	newTestCase.ClassName = className
	newTestCase.Time = duration
	newTestCase.File = filePath
	newTestCase.Result = result

	switch result {
	case "failed":
		newFailedTestCase := Failure{}
		newFailedTestCase.Content = errorTrace
		newTestCase.Failure = &newFailedTestCase
		if stringValidator(logs) {
			newSystemErrorLog := ErrorLog{}
			newSystemErrorLog.Content = logs
			newTestCase.ErrorLog = &newSystemErrorLog
		}
	case "skipped":
		newSkippedTestCase := Skipped{}
		newSkippedTestCase.Message = errorTrace
		newTestCase.Skipped = &newSkippedTestCase
		if stringValidator(logs) {
			newSystemConsoleLog := ConsoleLog{}
			newSystemConsoleLog.Content = logs
			newTestCase.ConsoleLog = &newSystemConsoleLog
		}
	default:
		if stringValidator(logs) {
			newSystemConsoleLog := ConsoleLog{}
			newSystemConsoleLog.Content = logs
			newTestCase.ConsoleLog = &newSystemConsoleLog
		}
	}

	propertiesObject := Properties{}

	newPropertiesList := testPropsCreator(testProps, buildIdentifier)

	propertiesObject.Properties = append(propertiesObject.Properties, newPropertiesList...)

	if len(newPropertiesList) > 0 {
		newTestCase.Properties = &propertiesObject
	}

	return newTestCase
}
