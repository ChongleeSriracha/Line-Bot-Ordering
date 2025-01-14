package models
import (
    "encoding/json"
    "io/ioutil"
    "log"
    "strings"
	 "net/http"
)

func CreateJsonFlexProduct() {
    resp, err := http.Get("http://localhost:8081/api/products/avaliable")
    if err != nil {
        log.Printf("Error calling API to get available products: %v", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        log.Printf("Error: received non-200 response code: %d", resp.StatusCode)
        return
    }

    var result struct {
        Products []map[string]interface{} `json:"products"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        log.Printf("Error decoding API response: %v", err)
        return
    }

    file, err := ioutil.ReadFile("./src/view/json/flex-product.json")
    if err != nil {
        log.Printf("Error reading flex-product.json file: %v", err)
        return
    }

    var bubbles []map[string]interface{}

    for _, product := range result.Products {
        productName := product["name"].(string)
        flexTemplateStr := string(file)
        flexTemplateStr = strings.Replace(flexTemplateStr, "Flower-1", productName, -1)

        var modifiedFlexTemplate map[string]interface{}
        if err := json.Unmarshal([]byte(flexTemplateStr), &modifiedFlexTemplate); err != nil {
            log.Printf("Error unmarshalling modified flex-product.json: %v", err)
            continue
        }

        bubbles = append(bubbles, modifiedFlexTemplate)
    }

    flexMessage := map[string]interface{}{
        "to":     "U6de5e84c204f582fcea6d9e426d913ce",
        "messages": bubbles,
    }

    flexMessageJSON, err := json.Marshal(flexMessage)
    if err != nil {
        log.Printf("Error marshalling flex message: %v", err)
        return
    }

    log.Printf("Flex Message with carousel: %s", string(flexMessageJSON))
}