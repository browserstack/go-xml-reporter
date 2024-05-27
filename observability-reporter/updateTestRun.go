package observabilityReporter

import (
	"errors"
)

func (jr *JUnitReporter) UpdateTestRunDetails(testIdentifier string, testProps map[string]interface{}, logs string, additionalProps map[string]string) (string, error) {
	if !stringValidator(testIdentifier) {
		return "failure", errors.New("mandatory parameter testIdentifier is missing")
	}

	matchedTestSuiteKey := ""
	matchedTestCasesList := []TestCase{}
	matchedTestCase := TestCase{}
	matchedTestCaseIdx := -1

	for testSuiteKey, testSuiteValue := range jr.testSuites {
		for testCaseIdx, testCaseValue := range testSuiteValue.TestCases {
			if testCaseValue.Id == testIdentifier {
				matchedTestCaseIdx = testCaseIdx
				matchedTestCase = testCaseValue
				break
			}
		}
		if stringValidator(matchedTestCase.Id) {
			matchedTestSuiteKey = testSuiteKey
			matchedTestCasesList = testSuiteValue.TestCases
			break
		}
	}

	if !stringValidator(matchedTestCase.Id) {
		return "failure", errors.New("no test run is matching with the given identifier")
	}

	// Logs section
	switch matchedTestCase.Result {
	case "failed":
		if stringValidator(logs) {
			newSystemErrorLog := ErrorLog{}
			newSystemErrorLog.Content = logs
			matchedTestCase.ErrorLog = &newSystemErrorLog
		}
	case "skipped":
		if stringValidator(logs) {
			newSystemConsoleLog := ConsoleLog{}
			newSystemConsoleLog.Content = logs
			matchedTestCase.ConsoleLog = &newSystemConsoleLog
		}
	default:
		if stringValidator(logs) {
			newSystemConsoleLog := ConsoleLog{}
			newSystemConsoleLog.Content = logs
			matchedTestCase.ConsoleLog = &newSystemConsoleLog
		}
	}

	// Test prop section
	propertiesObject := matchedTestCase.Properties
	newPropertiesList := testPropsCreator(testProps, jr.buildDetails.buildIdentifier)
	propertiesObject.Properties = append(propertiesObject.Properties, newPropertiesList...)
	if len(propertiesObject.Properties) > 0 {
		matchedTestCase.Properties = propertiesObject
	}

	matchedTestCasesList[matchedTestCaseIdx] = matchedTestCase
	jr.testSuites[matchedTestSuiteKey].TestCases[matchedTestCaseIdx] = matchedTestCase

	return "success", nil
}
