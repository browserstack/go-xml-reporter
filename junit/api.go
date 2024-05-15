package junit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

const (
	OBSERVABILITY_UPLOADER_ENDPOINT         = "https://upload-observability.browserstack.com/upload"
	OBSERVABILITY_STAGING_UPLOADER_ENDPOINT = "http://upload-observability-devtestops.bsstag.com/upload"
)

func O11yJunitUploader(buildDetails BuildDetailsParams) (string, error) {

	zipFilePath, err := getCurrentBuildZipFilePath(buildDetails.buildIdentifier)
	if err != nil {
		return "", err
	}

	// Open the file to be uploaded
	file, err := os.Open(zipFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create a new multipart writer to handle the file and other fields
	var requestBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBody)

	// Add file to the request
	fileWriter, err := multipartWriter.CreateFormFile("data", zipFilePath)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return "", err
	}

	// Add other form fields
	formFields := map[string]string{
		"projectName":      buildDetails.projectName,
		"buildName":        buildDetails.buildName,
		"buildIdentifier":  buildDetails.buildIdentifier,
		"tags":             strings.Join(buildDetails.buildTags, ", "),
		"ci":               buildDetails.ciDetails,
		"frameworkVersion": buildDetails.frameworkDetails,
	}

	for key, value := range formFields {
		err := multipartWriter.WriteField(key, value)
		if err != nil {
			fmt.Println("Error writing form field:", err)
			return "", err
		}
	}

	// Close multipart writer to finalize the request body
	err = multipartWriter.Close()
	if err != nil {
		return "", err
	}

	// Create a new POST request with the multipart body
	req, err := http.NewRequest("POST", OBSERVABILITY_STAGING_UPLOADER_ENDPOINT, &requestBody)
	if err != nil {
		return "", err
	}

	// Define the Basic Auth token
	username := buildDetails.username
	password := buildDetails.password
	authToken := base64Encode(fmt.Sprintf("%s:%s", username, password))
	req.Header.Set("Authorization", "Basic "+authToken)

	// Set Content-Type header
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code", resp.StatusCode)
		return "nil", errors.New("error: unexpected status code " + (string)(resp.StatusCode))
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		// return nil
	}

	// Unmarshal the JSON data into the struct
	var responseData UploaderResponseData
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("Error unmarshalling response body:", err)
		return "", err
	}

	// Access the desired key
	message := responseData.Message

	return message, nil

}
