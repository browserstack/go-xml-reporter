package junit

import (
	"archive/zip"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func generateUUID() string {
	id := uuid.New()

	return id.String()
}

func getCurrentBuildFolderPath(buildIdentifier string) (string, error) {
	// Specify the directory name
	dirName := buildIdentifier
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	buildFolder := filepath.Join(wd, "build")
	return filepath.Join(buildFolder, dirName), nil
}

func getCurrentBuildZipFilePath(buildIdentifier string) (string, error) {
	// Specify the directory name
	fileName := buildIdentifier + ".zip"
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	buildFolder := filepath.Join(wd, "build")
	return filepath.Join(buildFolder, fileName), nil
}

func createBuildDirectory(buildIdentifier string) error {

	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	buildFolder := filepath.Join(wd, "build")

	if _, err := os.Stat(buildFolder); os.IsNotExist(err) {
		// Create the folder
		err := os.MkdirAll(buildFolder, 0755)
		if err != nil {
			fmt.Println("Error creating folder:", err)
		}
	}

	// Create the full path
	fullPath, err := getCurrentBuildFolderPath(buildIdentifier)
	if err != nil {
		return err
	}

	// Create the directory
	err = os.Mkdir(fullPath, 0755)
	if err != nil {
		return err
	}

	return nil
}

func removeBuildAssets(buildIdentifier string) error {

	// Create the full path
	fullBuildAssetsPath, err := getCurrentBuildFolderPath(buildIdentifier)
	if err != nil {
		return err
	}
	// Specify the name of the zip file to be created
	zipFilePath, err := getCurrentBuildZipFilePath(buildIdentifier)
	if err != nil {
		return err
	}

	// Delete the directory
	err = os.RemoveAll(fullBuildAssetsPath)
	if err != nil {
		return err
	}

	// Delete the zip file
	err = os.Remove(zipFilePath)
	if err != nil {
		return err
	}

	return nil
}

func copyAttachmentsToBuildDir(attachmentPath string, buildIdentifier string) {

	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting current directory")
		return
	}

	// Create the source && destination full path
	sourceFilePath := filepath.Join(wd, attachmentPath)
	// Create the full path
	destinationFolderPath, err := getCurrentBuildFolderPath(buildIdentifier)
	if err != nil {
		return
	}

	// Open the source file
	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer sourceFile.Close()

	// Create the destination folder if it doesn't exist
	if _, err := os.Stat(destinationFolderPath); os.IsNotExist(err) {
		if err != nil {
			fmt.Println("Error creating destination folder:", err)
			return
		}
	}

	// Extract the file name from the source file path
	fileInfo, err := sourceFile.Stat()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fileName := fileInfo.Name()

	// Create the destination file
	destinationFilePath := destinationFolderPath + "/" + fileName
	destinationFile, err := os.Create(destinationFilePath)
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer destinationFile.Close()

	// Copy the contents of the source file to the destination file
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}

}

func testPropsCreator(testProps map[string]interface{}, buildIdentifier string) []Property {
	propertiesList := []Property{}

	for propName, propValue := range testProps {
		if propValueStr, ok := propValue.(string); ok {
			newProperty := Property{}
			newProperty.Name = propName
			newProperty.Value = propValueStr
			if propName == "attachment" {
				newProperty.Value = filepath.Base(propValueStr)
				copyAttachmentsToBuildDir(propValueStr, buildIdentifier)
			}
			propertiesList = append(propertiesList, newProperty)
		} else if propValueArray, ok := propValue.([]string); ok {
			for _, propValueStr := range propValueArray {
				newProperty := Property{}
				newProperty.Name = propName
				newProperty.Value = propValueStr
				if propName == "attachment" {
					newProperty.Value = filepath.Base(propValueStr)
					copyAttachmentsToBuildDir(propValueStr, buildIdentifier)
				}
				propertiesList = append(propertiesList, newProperty)
			}
		} else {
			continue
		}
	}

	return propertiesList
}

// base64Encode encodes a string to base64
func base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// ConvertToXML converts an array of test suites to an XML string
func convertToXML(testsuites TestSuites) (string, error) {
	xmlData, err := xml.MarshalIndent(testsuites, "", "  ")
	if err != nil {
		return "", err
	}
	return xml.Header + string(xmlData), nil
}

// Generate XML file from testsuites
func generateXMLFromTestSuites(buildIdentifier string, testSuites map[string]TestSuite) error {

	// Create the full path
	destinationFolderPath, err := getCurrentBuildFolderPath(buildIdentifier)
	if err != nil {
		return err
	}

	// Create the full path
	fullFilePath := filepath.Join(destinationFolderPath, "example-test.xml")

	// Create a dummy xml file using io
	file, err := os.Create(fullFilePath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	parsedTestSuites := []TestSuite{}

	for _, v := range testSuites {
		parsedTestSuites = append(parsedTestSuites, v)
	}

	// Generate xml string from remaining testsuits
	var actualTestSuite = TestSuites{
		TestSuites: parsedTestSuites,
	}
	xmlString, err := convertToXML(actualTestSuite)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// // Write the XML data to the file
	_, err = file.WriteString(xmlString)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func createZipFolderForUploader(buildIdentifier string) {

	// Create the full path
	sourceFolder, err := getCurrentBuildFolderPath(buildIdentifier)
	if err != nil {
		return
	}

	// Specify the name of the zip file to be created
	zipFilePath, err := getCurrentBuildZipFilePath(buildIdentifier)
	if err != nil {
		return
	}

	// Create a new zip file
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		fmt.Println("Error creating zip file:", err)
		return
	}
	defer zipFile.Close()

	// Create a new zip archive
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Walk through the directory and add files to the zip archive
	err = filepath.Walk(sourceFolder, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get the relative path for the file in the zip archive
		relativePath, err := filepath.Rel(sourceFolder, filePath)
		if err != nil {
			return err
		}

		// If the file is a directory, skip it
		if info.IsDir() {
			return nil
		}

		// Open the file
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		// Create a new file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = relativePath
		// Set the compression method to DEFLATE
		header.Method = zip.Deflate

		// Add the file header to the zip archive
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// Copy the file data to the zip archive
		_, err = io.Copy(writer, file)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		fmt.Println("Error walking through directory:", err)
		return
	}

}
