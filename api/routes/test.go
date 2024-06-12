package routes

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/constants"
)

func Test(w http.ResponseWriter, r *http.Request) {
	// Step 1: Define the base URL and endpoint
	baseURL := constants.DATABASE_URL

	// Step 2: Set up headers
	req, err := http.NewRequest("GET", baseURL, nil) // Change "GET" to your desired method
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	log.Print("Bearer " + constants.DATABASE_API_KEY)
	req.Header.Add("Authorization", "Bearer "+constants.DATABASE_API_KEY)
	req.Header.Add("Content-Type", "application/json")

	// Step 3 & 4: For a POST request, prepare JSON payload (not shown here)

	// Step 5: Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Step 6: Handle the response
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}
		bodyString := string(bodyBytes)
		fmt.Println("Response:", bodyString)
	} else {
		fmt.Println("Received non-OK response:", resp.Status)
	}
}
