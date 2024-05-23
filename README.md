# go-xml-reporter
This repository is being created to assist client(s) to generate and upload xml reports to existing JUnit XML Uploader feature available on Observability


## Installation

Run this command in your workspace terminal

```shell
go get github.com/browserstack/go-xml-reporter
```

After installation is successful, you can verify the installation in your go.mod file like this

```go
module your-module

go 1.22.2

require (
    ...
	github.com/browserstack/go-xml-reporter v1.0.0 // version might differ as per your installation
)
```

## Usage

For complete usage, refer this example file. [example.go](https://github.com/browserstack/go-xml-reporter/blob/main/example/example.go)

## Contributing

Before getting started, ensure that you have installed the following requirements:

-   Go (v1.22.2)

Follow the steps below to set up the repository on your local machine:

1.  Fork the repository: 

    `https://github.com/browserstack/go-xml-reporter.git` 

2.  Clone the forked repository: 

    `git clone https://github.com/<---your-github-username--->/go-xml-reporter.git` 

3.  Navigate to the Repository: 

    `cd go-xml-reporter`

3.  Checkout to the main branch (or any branch of your choice):
   
    `git checkout main`

4.  Install Dependencies

    `go mod tidy`

### How to Setup and Upload a Example report

1.  Create a `.env` file 

    Add the following snippet to your .env file and replace the values with your actual credentials:

    ```
    BSTACK_USERNAME=<---YOUR_BROWSERSTACK_USERNAME--->
    BSTACK_PASSWORD=<---YOUR_BROWSERSTACK_PASSWORD--->
    ```

3. To get your username and password in your BrowserStack account, visit this [page](https://observability.browserstack.com/get-started/junit-reports). (Note: You must be logged in to get your credentials.)
    

4.  Run the main Go file:
    
    `go run main.go`

    This will upload a JUnit XML report to your BrowserStack's Test Observability product.

5.  Feel free to customize the values in the example.go file located inside the example folder according to your requirements. After making changes, you can rerun the command mentioned above to test different scenarios.



