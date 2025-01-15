package component

import (
	"bytes"
	
	"fmt"
	"io"
	
	"log"

	"line-Bot-Ordering/src/models"
	"line-Bot-Ordering/src/view"
	"net/http"
)

func FlexProduct(userID,channelAccessToken string)(error){
	
	flexMessageJSON , err := models.CreateJsonFlexProduct(userID)
	
	if err != nil {
		view.LogError(err, "Error creating JSON flex product")
		return err
	}

	log.Printf("Flex Message JSON: %s", string(flexMessageJSON))


	url := "https://api.line.me/v2/bot/message/push"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(flexMessageJSON))
	if err != nil {
		view.LogError(err, "Error creating POST request")
		return err
	}

	req.Header.Set("Authorization", "Bearer "+channelAccessToken)
	req.Header.Set("Content-Type", "application/json")

	clientL:= &http.Client{}
	resp, err := clientL.Do(req)
	if err != nil {
		view.LogError(err, "Error sending request to LINE API")
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		view.LogError(err, "Error reading response body")
		return err
	}

	log.Printf("Response Status: %s", resp.Status)
	log.Printf("Response Body: %s", string(respBody))

	if resp.StatusCode != http.StatusOK {
		view.LogError(fmt.Errorf(string(respBody)), "Error in response body")
		view.LogMessage(fmt.Sprintf("Error creating  flex bubble: %s", string(respBody)))
		return err
	}

	return nil

}