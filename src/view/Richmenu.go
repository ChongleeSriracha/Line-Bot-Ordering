package view

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


type RichMenu struct {
	Size        Size       `json:"size"`
	Selected    bool       `json:"selected"`
	Name        string     `json:"name"`
	ChatBarText string     `json:"chatBarText"`
	Areas       []Area     `json:"areas"`
}


type Size struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}


type Area struct {
	Bounds Bound   `json:"bounds"`
	Action Action  `json:"action"`
}


type Bound struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}


type Action struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// RichMenu function reads richmenu.json and creates the rich menu
func CreateRichMenu(channelAccessToken string) {

	file, err := ioutil.ReadFile("./src/view/richmenu.json")
	if err != nil {
		log.Fatal("Error reading richmenu.json file:", err)
	}

	// Step 2: Parse the JSON into a RichMenu struct
	var richMenu RichMenu
	if err := json.Unmarshal(file, &richMenu); err != nil {
		log.Fatal("Error unmarshalling richmenu.json:", err)
	}


	url := "https://api.line.me/v2/bot/richmenu"
	reqBody, err := json.Marshal(richMenu)
	if err != nil {
		log.Fatal("Error marshalling rich menu JSON:", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal("Error creating POST request:", err)
	}

	req.Header.Set("Authorization", "Bearer "+channelAccessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request to LINE API:", err)
	}
	defer resp.Body.Close()


	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error creating rich menu: %s", string(respBody))
		return
	}

	var response map[string]interface{}
	if err := json.Unmarshal(respBody, &response); err != nil {
		log.Fatal("Error unmarshalling response:", err)
	}

	richMenuID := response["richMenuId"].(string)
	log.Printf("Rich menu created successfully with ID: %s", richMenuID)


	uploadRichMenuImage(richMenuID, channelAccessToken)


	setDefaultRichMenu(richMenuID, channelAccessToken)
}

// uploadRichMenuImage uploads the image for the rich menu
func uploadRichMenuImage(richMenuID, channelAccessToken string) {
	imagePath := "./src/view/richmenu-image.png"

	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("Error opening image file:", err)
	}
	defer file.Close()

	url := fmt.Sprintf("https://api-data.line.me/v2/bot/richmenu/%s/content", richMenuID)
	req, err := http.NewRequest("POST", url, file)
	if err != nil {
		log.Fatal("Error creating POST request for image upload:", err)
	}

	req.Header.Set("Authorization", "Bearer "+channelAccessToken)
	req.Header.Set("Content-Type", "image/png")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error uploading image:", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error uploading image: %s", string(respBody))
		return
	}

	log.Println("Rich menu image uploaded successfully")
}

// setDefaultRichMenu sets the rich menu as the default for all users
func setDefaultRichMenu(richMenuID, channelAccessToken string) {
	url := fmt.Sprintf("https://api.line.me/v2/bot/user/all/richmenu/%s", richMenuID)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatal("Error creating POST request to set default rich menu:", err)
	}

	req.Header.Set("Authorization", "Bearer "+channelAccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error setting default rich menu:", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error setting default rich menu: %s", string(respBody))
		return
	}

	log.Println("Default rich menu set successfully")
}
