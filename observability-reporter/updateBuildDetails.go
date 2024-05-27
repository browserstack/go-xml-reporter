package observabilityReporter

import "errors"

func (jr *JUnitReporter) UpdateBuildDetails(buildTags []string, ciDetails string, frameworkDetails string, vcsDetails map[string]string, additionalProperties map[string]string) (string, error) {

	// If build identifier is not matching with current object stopping execution
	if !stringValidator(jr.buildDetails.buildIdentifier) {
		return "", errors.New("build identifier is not exist")
	}

	currentBuildDetails := jr.buildDetails

	if len(buildTags) > 0 {
		existingBuildTags := currentBuildDetails.buildTags
		buildTagsMap := make(map[string]string)
		for _, tag := range existingBuildTags {
			buildTagsMap[tag] = tag
		}
		for _, tag := range buildTags {
			buildTagsMap[tag] = tag
		}
		newBuildTags := []string{}
		for key := range buildTagsMap {
			newBuildTags = append(newBuildTags, key)
		}
		jr.buildDetails.buildTags = newBuildTags
	}

	if stringValidator(ciDetails) {
		jr.buildDetails.ciDetails = ciDetails
	}

	if stringValidator(frameworkDetails) {
		jr.buildDetails.frameworkDetails = frameworkDetails
	}

	return "success", nil
}
