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
  Event
  Status int `json:"Status"`
}

func HandleRequest(ctx context.Context, e Event) (SupportCase, error) {
  min := 0.0
  max := 1.0
  var supportCase SupportCase
  supportCase.Case = e.Case
  rand.Seed(time.Now().Unix())
  supportCase.Status = int(math.Floor(rand.Float64() * (max - min + 1.0)))
  if supportCase.Status == 1 {
    supportCase.Message = fmt.Sprintf("%sresolved...", e.Message)
  } else if supportCase.Status == 0 {
    supportCase.Message = fmt.Sprintf("%sunresolved...", e.Message)
  }
  return supportCase, nil
}

func main() {
  lambda.Start(HandleRequest)
}
