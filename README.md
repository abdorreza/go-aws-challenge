# go-aws-challenge
Welcome to the API documentation. This API allows you to manage devices and their corresponding information. The API accepts JSON requests and produces the corresponding HTTP responses.

## Examples

#### Create a Device

Request:

    HTTP Method: POST
    URL: https://<api-gateway-url>/api/devices
    Body (application/json):
    {
        "id": "/devices/id1",
        "deviceModel": "/devicemodels/id1",
        "name": "Sensor",
        "note": "Testing a sensor.",
        "serial": "A020000102"
    }

Responses:

        Success:
            HTTP Status: 201 Created
        Failure 1:
            HTTP Status: 400 Bad Request
            Description: If any of the payload fields are missing. The response body will contain a descriptive error message for the client to identify the problem.
        Failure 2:
            HTTP Status: 500 Internal Server Error
            Description: If any exceptional situation occurs on the server side.

#### Get Device Details

Request:

    HTTP Method: GET
    URL: https://<api-gateway-url>/api/devices/{id}
    Example: GET https://api123.amazonaws.com/api/devices/id1

Responses:

        Body (application/json):
        {
            "id": "/devices/id1",
            "deviceModel": "/devicemodels/id1",
            "name": "Sensor",
            "note": "Testing a sensor.",
            "serial": "A020000102"
        }

Responses:
            
    Success:
        HTTP Status: 200 OK
    Failure 1:
        HTTP Status: 404 Not Found
        Description: If the requested id does not exist in the database.
    Failure 2:
        HTTP Status: 500 Internal Server Error
        Description: If any exceptional situation occurs on the server side.

Feel free to use this API to interact with devices and retrieve their details. If you encounter any issues or have questions, please refer to the provided HTTP response codes and descriptions to diagnose the problem.

## Environments
    * Installing serverless [visit here](https://www.serverless.com/framework/docs/getting-started).
    * using [AWS configure in terminal](https://aws.amazon.com/cli/) for setting 'AWS Access Key ID','AWS Secret Access Key', 'Default region name' and 'Default output format' for connecting to AWS.
    * Installing [Postman](https://www.postman.com/) for testing POST and GET requests.

## Tests
    In testing, we used mocking the process of simulating the behavior of external
    dependencies or collaborating components during tests. Instead of relying on real
    implementations, we create "mocks" that mimic the expected behavior of these
    dependencies. This allows us to isolate the code under test and focus on specific
    scenarios without the complexities of real-world interactions.

    Testing is a crucial part of building reliable and bug-free software, and Golang
    provides a powerful testing framework to ensure the quality of our code. The command
    'go test ./...' is a convenient and efficient way to run all the tests in your
    project, including sub-packages. By executing this simple command from the root of
    your project, Go will automatically traverse all directories and execute tests in
    each package it discovers.

## TODO
    - Writing more test cases
    - Fix deploy problem