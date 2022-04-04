package controller

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCorrectResponseCode(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    uint64
	}{
		{"fraud detection", "fraud detected", FraudDetectedCode},
		{"limit exceeded", "limit exceeded", LimitExceededCode},
		{"daily limit exceeded", "daily limit exceeded", DailyLimitExceededCode},
		{"card blocked", "card blocked", CardBlockedCode},
		{"success", "success", SuccessCode},
		{"default", "any", ErrorOccurredCode},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseCode := correctResponseCode(tt.message)
			assert.Equal(t, tt.want, responseCode)
		})
	}
}
