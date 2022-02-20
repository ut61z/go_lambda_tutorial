package main

import (
  "fmt"
  "context"
  "github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
  Case string `json:"Case"`
	Message string `json:"Message"`
}

type SupportCase struct {
  MyCaseID string `json:"Case:"`
  MyMessage string `json:"Message:"`
}

func HandleRequest(ctx context.Context, e Event) (SupportCase, error) {
  return SupportCase{MyCaseID: e.Case, MyMessage: fmt.Sprintf("%sassigned...",  e.Message)}, nil
}

func main() {
  lambda.Start(HandleRequest)
}
