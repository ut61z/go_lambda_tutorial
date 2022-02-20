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

type SupportCase struct {
  MyCaseID string `json:"Case:"`
  MyCaseStatus int `json:"Status:"`
  MyMessage string `json:"Message:"`
}

func HandleRequest(ctx context.Context, e Event) (SupportCase, error) {
  return SupportCase{MyCaseID: e.Case, MyCaseStatus: e.Status, MyMessage: fmt.Sprintf("%sescalating.",  e.Message)}, nil
}

func main() {
  lambda.Start(HandleRequest)
}
