package component

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"line-Bot-Ordering/src/view"
	"line-Bot-Ordering/src/models"
	"net/http"
	"os"
)

func CreateRichMenu(channelAccessToken string) {
	file, err := ioutil.ReadFile("./src/view/json/richmenu.json")
	if err != nil {
		view.LogError(err, "Error reading richmenu.json file")
	}

	var richMenu models.RichMenu
	if err := json.Unmarshal(file, &richMenu); err != nil {
		view.LogError(err, "Error unmarshalling richmenu.json")
	}

	url := "https://api.line.me/v2/bot/richmenu"
	reqBody, err := json.Marshal(richMenu)
	if err != nil {
		view.LogError(err, "Error marshalling rich menu JSON")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		view.LogError(err, "Error creating POST request")
	}

	req.Header.Set("Authorization", "Bearer "+channelAccessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		view.LogError(err, "Error sending request to LINE API")
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		view.LogError(err, "Error reading response body")
	}

	if resp.StatusCode != http.StatusOK {
		view.LogMessage(fmt.Sprintf("Error creating rich menu: %s", string(respBody)))
		return
	}

	var response map[string]interface{}
	if err := json.Unmarshal(respBody, &response); err != nil {
		view.LogError(err, "Error unmarshalling response")
	}

	richMenuID := response["richMenuId"].(string)
	view.LogMessage(fmt.Sprintf("Rich menu created successfully with ID: %s", richMenuID))

	uploadRichMenuImage(richMenuID, channelAccessToken)
	setDefaultRichMenu(richMenuID, channelAccessToken)
}

func uploadRichMenuImage(richMenuID, channelAccessToken string) {
	imagePath := "./src/view//img/richmenu-image.png"
	file, err := os.Open(imagePath)
	if err != nil {
		view.LogError(err, "Error opening image file")
	}
	defer file.Close()

	url := fmt.Sprintf("https://api-data.line.me/v2/bot/richmenu/%s/content", richMenuID)
	req, err := http.NewRequest("POST", url, file)
	if err != nil {
		view.LogError(err, "Error creating POST request for image upload")
	}

	req.Header.Set("Authorization", "Bearer "+channelAccessToken)
	req.Header.Set("Content-Type", "image/png")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		view.LogError(err, "Error uploading image")
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		view.LogError(err, "Error reading response body")
	}

	if resp.StatusCode != http.StatusOK {
		view.LogMessage(fmt.Sprintf("Error uploading image: %s", string(respBody)))
		return
	}

	view.LogMessage("Rich menu image uploaded successfully")
}

func setDefaultRichMenu(richMenuID, channelAccessToken string) {
	url := fmt.Sprintf("https://api.line.me/v2/bot/user/all/richmenu/%s", richMenuID)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		view.LogError(err, "Error creating POST request to set default rich menu")
	}

	req.Header.Set("Authorization", "Bearer "+channelAccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		view.LogError(err, "Error setting default rich menu")
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		view.LogError(err, "Error reading response body")
	}

	if resp.StatusCode != http.StatusOK {
		view.LogMessage(fmt.Sprintf("Error setting default rich menu: %s", string(respBody)))
		return
	}

	view.LogMessage("Default rich menu set successfully")
}
