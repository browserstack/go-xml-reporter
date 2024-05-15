package junit

import "errors"

// UpdateBuild updates the specified build with optional parameters.
func (jr *JUnitReporter) UpdateBuildDetails(buildIdentifier string, buildTags []string, ciDetails string, frameworkDetails string, vcsDetails map[string]string, additionalProperties map[string]string) (string, error) {

	// validation for required params
	if !stringValidator(buildIdentifier) {
		return "failed", errors.New("build identifier is required")
	}

	// If build identifier is not matching with current object stopping execution
	if jr.buildDetails.buildIdentifier != buildIdentifier {
		return "failed", errors.New("build identifier is not matching")
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
