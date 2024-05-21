package observabilityReporter

func (jr *JUnitReporter) CreateBuild(username string, password string, buildName string, projectName string, buildIdentifier string, buildTags []string, ciDetails string, frameworkDetails string, vcsDetails map[string]string, additionalProperties map[string]string) (string, error) {

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
		return buildIdentifier, validationError
	}

	// Save build info in memory
	jr.buildDetails = buildDetailsObject

	// Create a folder to store xml files and attachments
	createBuildDirErr := createBuildDirectory(jr.buildDetails.buildIdentifier)
	if createBuildDirErr != nil {
		removeBuildAssets(jr.buildDetails.buildIdentifier)
	}

	return jr.buildDetails.buildIdentifier, nil
}
