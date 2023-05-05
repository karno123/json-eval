package main

import (
	"encoding/json"
	"fmt"
	"time"

	jsoneval "github.com/karno123/json-eval"
)

func main() {

	jsonStr := `{
		"id": true,
		"stream_timestamp": 2000,
		"payload": {
			"event_type": "PostCashOut",
			"event_id": "432ab66e-16d5-40af-a74b-4d584715a76l3-karno-61",
			"entities": [
				{
					"entity_type": "customer",
					"entity_id": "111222333",
					"phone_number": "+6285878982122"
				},
				{
					"entity_type": "merchant",
					"entity_id": "201100000002713936"
				}
			],
			"recommendations": [
				"notify",
				"ban",
				"create_ticket",
				"create_ticket_merchant"
			],
			"rule_results": [
				{
					"treatments": [
						"ban",
						"notify",
						"create_ticket"
					],
					"public_description": "testing public description",
					"rule_id": 20,
					"version_id": 245,
					"value_collection": {
						"notification_title": "Account Disabled",
						"notification_message": "Sorry - we detect suspicious activity in our account, and need to put your account under banning",
						"notification_user_type": "customer",
						"notification_user_id": "111222333",
						"ticket_description": "testing ticket description",
						"fraud_type": "Social Engineering"
					}
				},
				{
					"treatments": [
						"ban",
						"notify",
						"create_ticket"
					],
					"rule_id": 21,
					"public_description": "testing public description 11",
					"version_id": 245,
					"value_collection": {
						"notification_title": "Account Disabled",
						"notification_message": "Sorry - we detect suspicious activity in our account, and need to put your account under banning",
						"notification_user_type": "customer",
						"notification_user_id": "111222333",
						"ticket_description": "testing ticket description 11",
						"fraud_type": "Social Engineering"
					}
				},
				{
					"treatments": [
						"ban",
						"notify",
						"create_ticket_merchant"
					],
					"rule_id": 22,
					"version_id": 245,
					"public_description": "testing public description 2",
					"value_collection": {
						"notification_title": "Account Disabled",
						"notification_message": "Sorry - we detect suspicious activity in our account, and need to put your account under banning",
						"notification_user_type": "customer",
						"notification_user_id": "111222333",
						"ticket_description": "testing ticket description 2",
						"fraud_type": "Social Engineering"
					}
				},
				{
					"treatments": [
						"ban",
						"notify",
						"create_ticket_merchant"
					],
					"rule_id": 23,
					"public_description": "testing public description 12",
					"version_id": 245,
					"value_collection": {
						"notification_title": "Account Disabled",
						"notification_message": "Sorry - we detect suspicious activity in our account, and need to put your account under banning",
						"notification_user_type": "customer",
						"notification_user_id": "111222333",
						"ticket_description": "testing ticket description 2",
						"fraud_type": "Social Engineering"
					}
				}
			],
			"errors": []
		}
	}`

	var x map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &x)

	currTime := time.Now().UnixMilli()
	// expression := `payload.recommendations[0] != "notify" && (stream_timestamp < 100) || (payload.event_type == "PostCashOut")`
	expression := `"abc" && "abc"`
	mapEval := jsoneval.NewJsonEvaluator()
	result, err := mapEval.Evaluate(expression, x)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(time.Now().UnixMilli() - currTime)

	jsonStr = `
    {
        "glossary": {
            "title": "example glossary",
            "GlossDiv": {
                "total": 1000
            }
        }
    }    
`
	expression = `glossary.title == "example glossary" && ( glossary.GlossDiv.total > 100  )    `
	jsonEval := jsoneval.NewJsonEvaluator()
	result, err = jsonEval.EvaluateJson(expression, jsonStr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
