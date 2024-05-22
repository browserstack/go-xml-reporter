package example

import (
	"fmt"
	observabilityReporter "go-xml-reporter/observability-reporter"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Example() {

	// Load the environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Access the environment variables
	BSTACK_USERNAME := os.Getenv("BSTACK_USERNAME")
	BSTACK_PASSWORD := os.Getenv("BSTACK_PASSWORD")

	// Create a new reporter object
	build1, createBuildErr := observabilityReporter.CreateBuild(BSTACK_USERNAME, BSTACK_PASSWORD, "test-build", "Junit report uploads - Go XML Library", "", []string{"junit_upload", "regression"}, "http://localhost:8080/", "mocha, 10.0.0", map[string]string{}, map[string]string{})
	if createBuildErr != nil {
		panic(createBuildErr)
	}

	fmt.Println("***", build1.BuildIdentifier)

	// Updating build info
	_, updateBuildErr := build1.UpdateBuildDetails([]string{"xml"}, "", "", map[string]string{}, map[string]string{})
	if updateBuildErr != nil {
		panic(updateBuildErr)
	}

	// Adding a test
	testIdentifier1, addTestErr1 := build1.AddTestRun("Test 1", "failed",
		`Test 1 error stacktrace 1`, "2023-05-24T11:00:14", "3.133", "nil",
		`classname 1`,
		`/Users/testuser/work/samples/signIn.e2e.js`,
		map[string]interface{}{
			"browser":    "Google Chrome",
			"os":         "Windows",
			"os_version": "11",
			"devicename": "Samsung Galaxy S10 Plus",
			"author":     "Adrian",
			"attachment": []string{"example/screenshots/browserstack.png", "example/screenshots/observability.jpeg"},
			"tags":       []string{"p1", "must_pass", "sanity"},
		}, "log info", nil)

	if addTestErr1 != nil {
		panic(addTestErr1)
	}

	// Updating a test
	_, updateTestErr := build1.UpdateTestRunDetails(testIdentifier1, map[string]interface{}{
		"tag":        []string{"extra"},
		"attachment": []string{"example/screenshots/checklist.png"},
	}, "log info updated", map[string]string{})

	if updateTestErr != nil {
		panic(updateTestErr)
	}

	firstSplitResponse, firstSplitErr := build1.SendReport()
	if firstSplitErr != nil {
		panic(firstSplitErr)
	}

	fmt.Println("First split response message: ", firstSplitResponse)

	// Adding more tests
	build1.AddTestRun("Test 1", "passed",
		``, "2023-05-24T11:00:14", "3.133", "nil",
		`classname 1`,
		`/Users/testuser/work/samples/home.js`,
		nil, "log info", nil)
	build1.AddTestRun("Test 2",
		"failed", `first error stacktrace`,
		"2023-05-24T11:00:17",
		"2.343",
		"nil",
		`classname 2`,
		`/Users/testuser/work/samples/home.js`,
		nil, "", nil,
	)
	build1.AddTestRun(`BStack - Login fucntionality &quot;after all&quot; hook for &quot;Login with invalid credentials - Dynamic Skip&quot;`,
		"failed", `second error stacktrace`,
		"2023-05-24T11:00:17",
		"1.500",
		"nil",
		`classname 2`,
		`/Users/testuser/work/samples/home.js`,
		nil, "", nil,
	)

	secondSplitResponse, secondSplitErr := build1.SendReport()
	if secondSplitErr != nil {
		panic(secondSplitErr)
	}

	fmt.Println("Second split response message: ", secondSplitResponse)

	build1.AddTestRun("Test 1", "failed",
		`Test 1 error stacktrace 1`, "2023-05-24T11:00:14", "3.133", "parentsuite 1",
		`parentsuite 1 classname 1`,
		`/Users/testuser/work/samples/signIn.e2e.js`,
		nil, "", nil)
	build1.AddTestRun("Test 2",
		"failed", `first error stacktrace`,
		"2023-05-24T11:00:17",
		"2.343",
		"parentsuite 2",
		`parentsuite 2 classname 1`,
		`/Users/testuser/work/samples/signIn.e2e.js`,
		nil, "", nil,
	)
	build1.AddTestRun(`BStack - Login fucntionality &quot;after all&quot; hook for &quot;Login with invalid credentials - Dynamic Skip&quot;`,
		"failed", `second error stacktrace`,
		"2023-05-24T11:00:17",
		"1.500",
		"parentsuite 2",
		`parentsuite 2 classname 2`,
		`/Users/testuser/work/samples/signIn.e2e.js`,
		nil, "", nil,
	)

	// Stop a build ( this will generate xml report for tests added and upload it to browserstack observability. )
	respMessage, err := build1.StopBuild()
	if err != nil {
		panic(err)
	}

	fmt.Println("Response message: ", respMessage)
}
