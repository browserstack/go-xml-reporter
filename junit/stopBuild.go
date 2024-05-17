package junit

func (jr *JUnitReporter) StopBuild() (string, error) {

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
		jr.resetData()

		return respMessage, nil
	}

	// Delete the created file & reset the inmemory to empty(default) values
	jr.resetData()

	return "Build stop is successful", nil
}
