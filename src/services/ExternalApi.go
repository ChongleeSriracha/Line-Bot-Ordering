package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"line-Bot-Ordering/src/models"
	"log"
	"net/http"
	"os"
)


// Richmenu
func CreateRichMenu(channelAccessToken string) {
    file, err := ioutil.ReadFile("./src/view/json/richmenu.json")
    if err != nil {
        log.Fatal("Error reading richmenu.json file: ", err)
    }

    var richMenu models.RichMenu
    if err := json.Unmarshal(file, &richMenu); err != nil {
        log.Fatal("Error unmarshalling richmenu.json: ", err)
    }

    url := "https://api.line.me/v2/bot/richmenu"
    reqBody, err := json.Marshal(richMenu)
    if err != nil {
        log.Fatal("Error marshalling rich menu JSON: ", err)
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
    if err != nil {
        log.Fatal("Error creating POST request: ", err)
    }

    req.Header.Set("Authorization", "Bearer "+channelAccessToken)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal("Error sending request to LINE API: ", err)
    }
    defer resp.Body.Close()

    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal("Error reading response body: ", err)
    }

    if resp.StatusCode != http.StatusOK {
        log.Fatal(fmt.Sprintf("Error creating rich menu: %s", string(respBody)))
        return
    }

    var response map[string]interface{}
    if err := json.Unmarshal(respBody, &response); err != nil {
        log.Fatal("Error unmarshalling response: ", err)
    }

    richMenuID := response["richMenuId"].(string)
    log.Println(fmt.Sprintf("Rich menu created successfully with ID: %s", richMenuID))

    uploadRichMenuImage(richMenuID, channelAccessToken)
    setDefaultRichMenu(richMenuID, channelAccessToken)
}

func uploadRichMenuImage(richMenuID, channelAccessToken string) {
    imagePath := "./src/view/img/richmenu-image.png"
    file, err := os.Open(imagePath)
    if err != nil {
        log.Fatal("Error opening image file: ", err)
    }
    defer file.Close()

    url := fmt.Sprintf("https://api-data.line.me/v2/bot/richmenu/%s/content", richMenuID)
    req, err := http.NewRequest("POST", url, file)
    if err != nil {
        log.Fatal("Error creating POST request for image upload: ", err)
    }

    req.Header.Set("Authorization", "Bearer "+channelAccessToken)
    req.Header.Set("Content-Type", "image/png")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal("Error uploading image: ", err)
    }
    defer resp.Body.Close()

    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal("Error reading response body: ", err)
    }

    if resp.StatusCode != http.StatusOK {
        log.Fatal(fmt.Sprintf("Error uploading image: %s", string(respBody)))
        return
    }

    log.Println("Rich menu image uploaded successfully")
}

func setDefaultRichMenu(richMenuID, channelAccessToken string) {
    url := fmt.Sprintf("https://api.line.me/v2/bot/user/all/richmenu/%s", richMenuID)
    req, err := http.NewRequest("POST", url, nil)
    if err != nil {
        log.Fatal("Error creating POST request to set default rich menu: ", err)
    }

    req.Header.Set("Authorization", "Bearer "+channelAccessToken)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal("Error setting default rich menu: ", err)
    }
    defer resp.Body.Close()

    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal("Error reading response body: ", err)
    }

    if resp.StatusCode != http.StatusOK {
        log.Fatal(fmt.Sprintf("Error setting default rich menu: %s", string(respBody)))
        return
    }

    log.Println("Default rich menu set successfully")
}


//Flex bubble product

func FlexProduct(userID,channelAccessToken string)(error){
	
	flexMessageJSON , err := models.CreateJsonFlexProduct(userID)
	
	if err != nil {
		log.Fatal("Error creating JSON flex product")
		return err
	}

	log.Printf("Flex Message JSON: %s", string(flexMessageJSON))


	url := "https://api.line.me/v2/bot/message/push"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(flexMessageJSON))
	if err != nil {
		log.Fatal("Error creating POST request")
		return err
	}

	req.Header.Set("Authorization", "Bearer "+channelAccessToken)
	req.Header.Set("Content-Type", "application/json")

	clientL:= &http.Client{}
	resp, err := clientL.Do(req)
	if err != nil {
		log.Fatal("Error sending request to LINE API")
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		
		log.Fatal("Error reading response body")
		return err
	}

	log.Printf("Response Status: %s", resp.Status)
	log.Printf("Response Body: %s", string(respBody))

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Error creating flex bubble")
		return err
	}

	return nil

}