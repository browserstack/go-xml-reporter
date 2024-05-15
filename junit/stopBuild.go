package junit

func (jr *JUnitReporter) StopBuild() (string, error) {

	err := generateXMLFromTestSuites(jr.buildDetails.buildIdentifier, jr.testSuites)
	if err != nil {
		return "", err
	}

	createZipFolderForUploader(jr.buildDetails.buildIdentifier)

	// // Send this xml file to o11y api
	_, uploaderError := O11yJunitUploader(jr.buildDetails)
	if uploaderError != nil {
		return "", uploaderError
	}

	// Delete the created file & reset the inmemory to empty(default) values
	jr.resetData()

	return "", nil
}
