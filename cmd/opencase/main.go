package main

import (
  "fmt"
  "context"
  "github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
  InputCaseID string `json:"inputCaseID"`
}

type SupportCase struct {
  Case string `json:"Case"`
  Message string `json:"Message"`
}

func HandleRequest(ctx context.Context, e Event) (SupportCase, error) {
  return SupportCase{Case: e.InputCaseID, Message: fmt.Sprintf("Case %s opened...",  e.InputCaseID)}, nil
}

func main() {
  lambda.Start(HandleRequest)
}
