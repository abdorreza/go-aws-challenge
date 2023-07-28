package device

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/abdorreza/go-aws-challenge/db"
	"github.com/aws/aws-lambda-go/events"
)

func Get(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Retrieve the "deviceId" path parameter from the request
	deviceId := request.PathParameters["id"]

	// Your logic to fetch device data based on "deviceId"
	// Replace the below sample response with your actual data retrieval code
	deviceData := fmt.Sprintf(`{"deviceId": "%s", "name": "Sensor", "note": "Testing a sensor."}`, deviceId)

	// Create the HTT
	// Create the HTTP response with status code 200 and the device data
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       deviceData,
	}, nil
}

func Add(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Retrieve the "deviceId" path parameter from the request
	fmt.Println("*******************************************************")
	deviceId := request.PathParameters["id"]

	// Your logic to fetch device data based on "deviceId"
	// Replace the below sample response with your actual data retrieval code
	deviceData := fmt.Sprintf(`{"deviceId": "%s", "name": "Sensor", "note": "Testing a sensor."}`, deviceId)

	// Create the HTT
	// Create the HTTP response with status code 200 and the device data
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       deviceData,
	}, nil
}

// ////////////////////////////////////////////

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home Page")
}

func getPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1 style='color:blue;'>Welcome to the GetPage!</h1>")

	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) < 4 {
		fmt.Fprintf(w, "Not Exist ID for searching device!")
	}

	device, err := db.GetDevice(r.Context(), "/devices/id2")
	if err != nil {
		fmt.Fprintf(w, err.Error()+"<br>")
		fmt.Println(err.Error())
	}
	fmt.Fprintf(w, "<table border='1'>")
	fmt.Fprintf(w, "<tr>")
	fmt.Fprintf(w, "<th>ID</th>")
	fmt.Fprintf(w, "<th>MODEL</th>")
	fmt.Fprintf(w, "<th>NAME</th>")
	fmt.Fprintf(w, "<th>NOTE</th>")
	fmt.Fprintf(w, "<th>SERIAL</th>")
	fmt.Fprintf(w, "</tr>")
	fmt.Fprintf(w, "<tr>")
	fmt.Fprintf(w, "<td>"+device.Id+"</td>")
	fmt.Fprintf(w, "<td>"+device.DeviceModel+"</td>")
	fmt.Fprintf(w, "<td>"+device.Name+"</td>")
	fmt.Fprintf(w, "<td>"+device.Note+"</td>")
	fmt.Fprintf(w, "<td>"+device.Serial+"</td>")
	fmt.Fprintf(w, "</tr>")
	fmt.Fprintf(w, "</table>")
}

func postPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the PostPage!\n\n")
	db.InsertDevice(context.TODO())
	fmt.Println("Endpoint Hit: PostPage")
}

func HandleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/api/devices", getPage)
	http.HandleFunc("/api/devices/wrt", postPage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
