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
            Description: If any of the payload fields are missing. The response body will contain a descriptive
            error message for the client to identify the problem.
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
- Installing serverless [visit here](https://www.serverless.com/framework/docs/getting-started).
- using [AWS configure in terminal](https://aws.amazon.com/cli/) for setting 'AWS Access Key ID','AWS Secret Access Key',
  'Default region name' and 'Default output format' for connecting to AWS.
- Installing [Postman](https://www.postman.com/) for testing POST and GET requests.
- Using 'serverless deploy' for deploying on AWS

## Tests
In our test, we used mocking so that we don't need real dynamoDB data on AWS. In this way, we use local data to perform tests.
To run all the tests, we can run command 'go test ./...' from the root of the project.

## TODO
- Writing more test cases
- Fix deploy problem
