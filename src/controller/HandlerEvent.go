package controller

import (
	"line-Bot-Ordering/src/component"
	"line-Bot-Ordering/src/view"
)

func HandleEventAction(action string, userID, channelAccessToken string) {
	if action == "Product" {
	
		err := component.FlexProduct(userID, channelAccessToken)
		if err != nil {
			view.LogError(err, "Error creating JSON flex product")
			return
		}
		view.LogMessage("Flex message sent successfully via push")
	}
}
