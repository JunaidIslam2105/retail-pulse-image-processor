# Retail Pulse Image Processor

## Description

The **Retail Pulse Image Processor** is a Go-based microservice designed to handle the processing of images from retail stores. This service allows for the submission of jobs related to images (such as calculating perimeters), tracking the status of these jobs, and storing results. It is designed to be scalable, efficient, and easily integrable into a larger ecosystem of services aimed at supporting retail operations.

Key features of this service:
- Accepts image processing jobs via a REST API endpoint.
- Each job contains metadata about the store and images to be processed.
- Allows checking the status of submitted jobs.
- Provides a simple interface for managing store and job information.

## Assumptions

- A sample CSV file containing valid store IDs is assumed to be already present inside the application for testing purposes.
- The images' perimeter calculations are simulated, with each job processing having an artificial delay to mimic real-world processing times.
- The API expects valid payloads in a specific JSON format for job submissions.

## Installation & Setup

### Prerequisites

Before running the project, ensure that the following dependencies are installed on your system:

1. **Go** - Ensure that Go is installed (v1.18 or later recommended).
    - You can install Go from [here](https://golang.org/dl/).

2. **Docker** (for running the app in a container):
    - You can install Docker from [here](https://www.docker.com/products/docker-desktop).

### Steps to Run the Project

1. Clone the repository to your local machine:
    ```bash
    git clone https://github.com/yourusername/retail-pulse-image-processor.git
    cd retail-pulse-image-processor
    ```

2. Set up the project environment:
    ```bash
    go mod tidy
    ```

### Running the Application with Docker

To run the application using Docker, follow these steps:

1. Build the Docker image:
    ```bash
    docker build -t retail-pulse-image-processor .
    ```

2. Run the Docker container:
    ```bash
    docker run -p 8080:8080 retail-pulse-image-processor
    ```

   The application will now be running in a Docker container and accessible on `http://localhost:8080`.


### Testing the API with Postman

To test the **Retail Pulse Image Processor** API using Postman, follow the steps below:

1. **Install Postman**:
    - If you don't already have Postman installed, download and install it from [here](https://www.postman.com/downloads/).

2. **Start the Application**:
    - Make sure the application is running before testing the API. You can run the application using Docker or directly from the terminal.
        - **Using Docker**:
          ```bash
          docker run -p 8080:8080 retail-pulse-image-processor
          ```
        - **Directly (Without Docker)**:
          Run the following command:
          ```bash
          go run cmd/main.go
          ```

3. **Submit Job Request**:
    - Open Postman and follow these steps:
        - **Method**: `POST`
        - **URL**: `http://localhost:8080/api/submit`
        - **Body**: Select `raw` and choose `JSON` format. Then, paste the following JSON payload:
          ```json
          {
            "count": 1,
            "visits": [
              {
                "store_id": "RP00007",
                "image_url": ["https://www.gstatic.com/webp/gallery/2.jpg"],
                "visit_time": "2025-01-23T10:00:00Z"
              }
            ]
          }
          ```
        - **Expected Response**:
            - **Status Code**: `201 Created`
            - **Response Body**: Should return the job ID in the following format:
          ```json
          {
            "job_id": "some-unique-job-id"
          }
          ```

4. **Check Job Status**:
    - To check the status of the job you just submitted, make a `GET` request in Postman:
        - **Method**: `GET`
        - **URL**: `http://localhost:8080/api/status?job_id=some-unique-job-id`
        - Replace `some-unique-job-id` with the actual job ID returned from the `POST` request above.
        - **Expected Response**:
            - **Status Code**: `200 OK`
            - **Response Body**: Should return the current status of the job:
          ```json
          {
            "status": "ongoing"
          }
          ```

5. **Error Handling Tests**:
    - **Invalid Job ID Test**:
        - Send a `GET` request with an invalid `job_id`:
            - **Method**: `GET`
            - **URL**: `http://localhost:8080/api/status?job_id=invalid-job-id`
            - **Expected Response**:
                - **Status Code**: `404 Not Found`
                - **Response Body**:
              ```json
              {
                "error": "Job not found"
              }
              ```

    - **Count Mismatch Test**:
        - Submit a job where `count` does not match the length of `visits` array:
            - **Method**: `POST`
            - **URL**: `http://localhost:8080/api/submit`
            - **Body**: Use the following invalid payload:
              ```json
              {
                "count": 2,
                "visits": [
                  {
                    "store_id": "RP00007",
                    "image_urls": ["https://www.gstatic.com/webp/gallery/2.jpg"],
                    "visit_time": "2025-01-23T10:00:00Z"
                  }
                ]
              }
              ```
            - **Expected Response**:
                - **Status Code**: `400 Bad Request`
                - **Response Body**:
              ```json
              {
                "error": "count does not match the number of visits"
              }
              ```

By following these steps in Postman, you can test the job submission, status retrieval, and error handling features of the **Retail Pulse Image Processor** API.
    
## Work Environment

- **Operating System**: macOS, Linux, or Windows
- **Text Editor/IDE**: Visual Studio Code, GoLand, or any editor that supports Go programming.
- **Libraries/Frameworks**:
    - **Go**: v1.18+ (for all backend development)
    - **HTTP Testing**: `httptest` for simulating HTTP requests and responses
    - **JSON Encoding**: Standard Go `encoding/json` package for handling JSON data
- **Additional Tools**:
    - **Docker** for containerization
    - **Postman** for API testing

## Improvements (Future Work)

Given more time, here are some potential improvements for the project:

1. **Error Handling**: Improve error handling by introducing more descriptive error messages and custom error types. Implement centralized logging for better traceability of issues.

2. **Persistent Storage**: Integrate persistent storage such as MySQL to store job and store data, allowing for job persistence across service restarts.

3. **Job Queueing**: Introduce a job queueing mechanism (e.g., using Redis or RabbitMQ) for better job management and to avoid overloading the system with concurrent requests.

4. **API Rate Limiting**: Implement rate limiting to prevent abuse and ensure fair use of the service.



---

