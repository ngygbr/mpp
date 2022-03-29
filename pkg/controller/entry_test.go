package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"mock-paymentprocessor/pkg/model"
	"mock-paymentprocessor/pkg/transaction"
	"reflect"
	"strings"
	"testing"
)

func TestCorrectResponseCode(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    uint64
	}{
		{"fraud detection", "fraud detected", transaction.FraudDetectedCode},
		{"limit exceeded", "limit exceeded", transaction.LimitExceededCode},
		{"daily limit exceeded", "daily limit exceeded", transaction.DailyLimitExceededCode},
		{"card blocked", "card blocked", transaction.CardBlockedCode},
		{"success", "success", transaction.SuccessCode},
		{"default", "any", 400},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := correctResponseCode(tt.message); got != tt.want {
				t.Errorf("correctResponseCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func convertToIOReader(testReqBody *model.TestRequest) io.Reader {
	var body io.Reader
	if testReqBody.Body != nil {
		typeOf := reflect.TypeOf(testReqBody.Body).String()
		if strings.Contains(typeOf, "map") {
			jsonValue, err := json.Marshal(testReqBody.Body)
			if err != nil {
				return nil
			}
			body = bytes.NewBuffer(jsonValue)
		} else if typeOf == "string" {
			body = bytes.NewBuffer([]byte(testReqBody.Body.(string)))
		} else if typeOf == "[]uint8" {
			body = bytes.NewBuffer(testReqBody.Body.([]byte))
		} else {
			jsonValue, err := json.Marshal(testReqBody.Body)
			if err != nil {
				return nil
			}
			body = bytes.NewBuffer(jsonValue)
		}
	}

	return body
}
