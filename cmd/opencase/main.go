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
  MyCaseID string `json:"Case:"`
  MyMessage string `json:"Message:"`
}

func HandleRequest(ctx context.Context, e Event) (SupportCase, error) {
  return SupportCase{MyCaseID: e.InputCaseID, MyMessage: fmt.Sprintf("Case %s opened...",  e.InputCaseID)}, nil
}

func main() {
  lambda.Start(HandleRequest)
}
