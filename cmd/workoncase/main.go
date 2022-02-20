package main

import (
  "fmt"
  "math"
	"math/rand"
	"time"
  "context"
  "github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
  Case string `json:"Case"`
  Message string `json:"Message"`
}

type SupportCase struct {
  MyCaseID string `json:"Case:"`
  MyCaseStatus int `json:"Status:"`
  MyMessage string `json:"Message:"`
}

func HandleRequest(ctx context.Context, e Event) (SupportCase, error) {
  min := 0.0
  max := 1.0
  var supportCase SupportCase
  supportCase.MyCaseID = e.Case
  rand.Seed(time.Now().Unix())
  supportCase.MyCaseStatus = int(math.Floor(rand.Float64() * (max - min + 1.0)))
  if supportCase.MyCaseStatus == 1 {
    supportCase.MyMessage = fmt.Sprintf("%sresolved...", e.Message)
  } else if supportCase.MyCaseStatus == 0 {
    supportCase.MyMessage = fmt.Sprintf("%sunresolved...", e.Message)
  }
  return supportCase, nil
}

func main() {
  lambda.Start(HandleRequest)
}
