package junit

import "errors"

func (jr *JUnitReporter) SendReport(buildIdentifier string) (string, error) {

	if len(jr.testSuites) > 0 {
		err := generateXMLFromTestSuites(jr.buildDetails.buildIdentifier, jr.testSuites)
		if err != nil {
			return "", err
		}
		createZipFolderForUploader(jr.buildDetails.buildIdentifier)

		// Send this xml file to o11y api
		respMessage, uploaderError := O11yJunitUploader(jr.buildDetails)
		if uploaderError != nil {
			return "", uploaderError
		}

		// Delete the created file & reset the inmemory to empty(default) values
		jr.resetTestSuites()

		// Create a folder to store xml files and attachments
		createBuildDirErr := createBuildDirectory(jr.buildDetails.buildIdentifier)
		if createBuildDirErr != nil {
			removeBuildAssets(jr.buildDetails.buildIdentifier)
		}

		return respMessage, nil
	}

	return "", errors.New("no tests are added to builder to process xml report")
}
