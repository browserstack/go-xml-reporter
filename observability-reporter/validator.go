package observabilityReporter

import (
	"errors"
	"strings"
)

func stringValidator(value string) bool {
	if value == "" || value == "null" || len(strings.TrimSpace(value)) == 0 {
		return false
	}
	return true
}

func BuildDetailsValidator(buildDetails BuildDetailsParams) error {
	username := buildDetails.username
	password := buildDetails.password
	buildName := buildDetails.buildName
	projectName := buildDetails.projectName

	if !stringValidator(username) {
		return errors.New("username is required")
	} else if !stringValidator(password) {
		return errors.New("password is required")
	} else if !stringValidator(buildName) {
		return errors.New("build name is required")
	} else if !stringValidator(projectName) {
		return errors.New("project name is required")
	}
	return nil
}

func AddTestValidator(testcaseDetails map[string]string) error {
	buildIdentifier := testcaseDetails["buildIdentifier"]
	name := testcaseDetails["name"]
	result := testcaseDetails["result"]
	startTime := testcaseDetails["startTime"]
	duration := testcaseDetails["duration"]
	className := testcaseDetails["className"]
	errorTrace := testcaseDetails["errorTrace"]

	if !stringValidator(buildIdentifier) {
		return errors.New("build identifier is required")
	} else if !stringValidator(name) {
		return errors.New("best name is required")
	} else if !stringValidator(result) {
		return errors.New("result is required")
	} else if !stringValidator(startTime) {
		return errors.New("start time is required")
	} else if !stringValidator(duration) {
		return errors.New("duration is required")
	} else if !stringValidator(className) {
		return errors.New("className is required")
	} else if result == "failed" && !stringValidator(errorTrace) {
		return errors.New("error stack trace is missing in params")
	}
	return nil
}
