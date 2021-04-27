package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)

func Handler(_ context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	language := 0
	if l, ok := r.QueryStringParameters["lang"]; ok {
		s, err := strconv.Atoi(l)
		language = s
		data, _ := json.Marshal(map[string]interface{}{
			"error": "provide language with queryparam [Example: 'lang=1']",
		})
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 400,
				Body:     string(data)  ,
			}, nil
		}
	}
	data, _ := json.Marshal(map[string]interface{}{
		"message": message(Language(language)),
	})
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(data),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
