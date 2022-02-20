package main

import (
  "fmt"
  "context"
  "github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
  Case string `json:"Case"`
	Status int `json:"Status"`
  Message string `json:"Message"`
}

type SupportCase Event

func HandleRequest(ctx context.Context, e Event) (SupportCase, error) {
  return SupportCase{Case: e.Case, Status: e.Status, Message: fmt.Sprintf("%sescalating.",  e.Message)}, nil
}

func main() {
  lambda.Start(HandleRequest)
}
