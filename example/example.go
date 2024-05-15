package example

import (
	"go-xml-reporter/junit"
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

	reporter1 := junit.NewJUnitReporter()

	buildIdentifier, err := reporter1.CreateBuild(BSTACK_USERNAME, BSTACK_PASSWORD, "Ryomen Sukuna", "Siva - Junit report uploads", "", []string{"junit_upload", "regression"}, "http://localhost:8080/", "mocha, 10.0.0", map[string]string{}, map[string]string{})

	if err != nil {
		panic(err)
	}

	// -- START --
	// First test suite with classname

	testIdentifier1, addTestErr1 := reporter1.AddTestRun(buildIdentifier, "Test 1", "failed",
		`Test 1 error stacktrace 1`, "2023-05-24T11:00:14", "3.133", "nil",
		`classname 1`,
		`/Users/akhilcherian/Repo/test-observability-samples/test-samples/nodejs/mocha/specs/e2e/bstack-demo/bdd/signIn/signIn.e2e.js`,
		map[string]interface{}{
			"browser":    "Google Chrome",
			"os":         "Andriod",
			"device":     "Samsung Galaxy s24 Ultra",
			"author":     "Adrian",
			"attachment": []string{"example/screenshots/browserstack.png", "example/screenshots/observability.jpeg"},
			"tags":       []string{"p1", "must_pass", "sanity"},
		}, "error logger 1", nil)

	if addTestErr1 != nil {
		panic(addTestErr1)
	}

	_, updateTestErr := reporter1.UpdateTestRunDetails(testIdentifier1, map[string]interface{}{
		"tag":        []string{"extra"},
		"attachment": []string{"example/screenshots/checklist.png"},
	}, "error logger 1 updated", map[string]string{})

	if updateTestErr != nil {
		panic(updateTestErr)
	}

	reporter1.AddTestRun(buildIdentifier, "Test 1", "passed",
		``, "2023-05-24T11:00:14", "3.133", "nil",
		`classname 1`,
		`/Users/akhilcherian/Repo/test-observability-samples/test-samples/nodejs/mocha/specs/e2e/bstack-demo/bdd/signIn/signIn.e2e.js`,
		nil, "system logger 1", nil)

	// Second test suite with classname
	reporter1.AddTestRun(buildIdentifier, "Test 2",
		"failed", `first error stacktrace`,
		"2023-05-24T11:00:17",
		"2.343",
		"nil",
		`classname 2`,
		"/Users/akhilcherian/Repo/test-observability-samples/test-samples/nodejs/mocha/specs/e2e/bstack-demo/bdd/signIn/signIn.e2e.js",
		nil, "", nil,
	)
	reporter1.AddTestRun(buildIdentifier, `BStack - Login fucntionality &quot;after all&quot; hook for &quot;Login with invalid credentials - Dynamic Skip&quot;`,
		"failed", `second error stacktrace`,
		"2023-05-24T11:00:17",
		"1.500",
		"nil",
		`classname 2`,
		"/Users/akhilcherian/Repo/test-observability-samples/test-samples/nodejs/mocha/specs/e2e/bstack-demo/bdd/signIn/signIn.e2e.js",
		nil, "", nil,
	)
	// --- END --

	// First test suite with parentSuite
	reporter1.AddTestRun(buildIdentifier, "Test 1", "failed",
		`Test 1 error stacktrace 1`, "2023-05-24T11:00:14", "3.133", "parentsuite 1",
		`parentsuite 1 classname 1`,
		`/Users/akhilcherian/Repo/test-observability-samples/test-samples/nodejs/mocha/specs/e2e/bstack-demo/bdd/signIn/signIn.e2e.js`,
		nil, "", nil)

	// Second test suite with parentSuite
	reporter1.AddTestRun(buildIdentifier, "Test 2",
		"failed", `first error stacktrace`,
		"2023-05-24T11:00:17",
		"2.343",
		"parentsuite 2",
		`parentsuite 2 classname 1`,
		"/Users/akhilcherian/Repo/test-observability-samples/test-samples/nodejs/mocha/specs/e2e/bstack-demo/bdd/signIn/signIn.e2e.js",
		nil, "", nil,
	)
	reporter1.AddTestRun(buildIdentifier, `BStack - Login fucntionality &quot;after all&quot; hook for &quot;Login with invalid credentials - Dynamic Skip&quot;`,
		"failed", `second error stacktrace`,
		"2023-05-24T11:00:17",
		"1.500",
		"parentsuite 2",
		`parentsuite 2 classname 2`,
		"/Users/akhilcherian/Repo/test-observability-samples/test-samples/nodejs/mocha/specs/e2e/bstack-demo/bdd/signIn/signIn.e2e.js",
		nil, "", nil,
	)
	// --- END --

	_, updateBuildErr := reporter1.UpdateBuildDetails(buildIdentifier, []string{"jjk", "regression", "mahoraga"}, "", "", map[string]string{}, map[string]string{})
	if updateBuildErr != nil {
		panic(updateBuildErr)
	}

	reporter1.StopBuild()
}
