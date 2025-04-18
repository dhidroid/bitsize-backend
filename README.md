# Go Backend API Service

This is a minimal Go backend API service, built as a learning project and demonstration. It's based on the principles outlined in the Google Cloud Run quickstart guide for Go services: [https://cloud.google.com/run/docs/quickstarts/build-and-deploy/deploy-go-service](https://cloud.google.com/run/docs/quickstarts/build-and-deploy/deploy-go-service)

## Description

This project demonstrates a simple API server written in Go. It's designed to be easily deployable and provides a basic framework for building more complex API services. The goal of this project is to provide a minimal service, so that the steps to create and run a simple service can be understood.

## How to Run

### Automatic Startup

When you start a workspace (e.g., using a development environment like Gitpod or a similar service), the server should automatically start running.

### Manual Startup

If you need to run the server manually, follow these steps:

1.  **Navigate to the project directory:** `cd <your-project-directory>`
2.  **Run the server:** `go run main.go`

This will start the server, and you can then access the API endpoints.

## Other relevant information
* **Dependencies**: The dependencies used in this project are managed by Go modules. Use `go mod tidy` to ensure all dependencies are correctly managed.
* **Error Logs**: In the event of a build error, see the `tmp/build-errors.log` file for details.
* **Binary**: When build successfuly, you can run the server by using the `tmp/main` binary.


