package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func CreateJsonFlexProduct(UserID string) ([]byte, error) {
    resp, err := http.Get("http://localhost:8081/api/products/avaliable")
    if err != nil {
        log.Fatalf("Error calling API to get available products: %v", err)
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        log.Fatal("Error: received non-200 response code: %d", resp.StatusCode)
        return nil ,err
    }

    var result struct {
        Products []map[string]interface{} `json:"products"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        log.Fatal("Error decoding API response: %v", err)
        return nil ,err
    }

    file, err := ioutil.ReadFile("./src/view/json/flex-product.json")
    if err != nil {
        log.Fatalf("Error reading flex-product.json file: %v", err)
        return nil ,err
    }

    var bubbles []map[string]interface{}

    for _, product := range result.Products {
        ProductName := product["name"].(string)
        Description := product["description"].(string)
        Price := "฿"+strconv.FormatFloat(product["price"].(float64), 'f', -1, 64)
  
        URL := product["url"].(string)

        flexTemplateStr := string(file)
        flexTemplateStr = strings.Replace(flexTemplateStr, "ProductName", ProductName, -1)
        flexTemplateStr = strings.Replace(flexTemplateStr, "Description", Description, -1)
        flexTemplateStr = strings.Replace(flexTemplateStr, "Price", Price, -1)
        flexTemplateStr = strings.Replace(flexTemplateStr, "UrlImg", URL, -1)

        var modifiedFlexTemplate map[string]interface{}
        if err := json.Unmarshal([]byte(flexTemplateStr), &modifiedFlexTemplate); err != nil {
            log.Fatal("Error unmarshalling modified flex-product.json: %v", err)
            continue
        }

        bubbles = append(bubbles, modifiedFlexTemplate)
    }
    flexMessage := map[string]interface{}{
        "to": UserID,
        "messages": []map[string]interface{}{
            {
                "type":     "flex", 
                "altText":  "Bubble Messages",
                "contents": map[string]interface{}{
                    "type": "carousel", 
                    "contents": bubbles,
                },
              
            },
        },
    }
    
  
    flexMessageJSON, err := json.Marshal(flexMessage)
    if err != nil {
        log.Fatal("Error marshalling flex message: %v", err)
        return nil, err
    }

    log.Printf("Flex Message with carousel" )

 

    return flexMessageJSON, nil
}