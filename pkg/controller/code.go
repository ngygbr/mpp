package controller

const (
	SuccessCode            = 100
	LimitExceededCode      = 201
	CardBlockedCode        = 202
	DailyLimitExceededCode = 203
	FraudDetectedCode      = 204
	ErrorOccurredCode      = 206
)

func correctResponseCode(message string) uint64 {
	switch message {
	case "fraud detected":
		return FraudDetectedCode
	case "limit exceeded":
		return LimitExceededCode
	case "daily limit exceeded":
		return DailyLimitExceededCode
	case "card blocked":
		return CardBlockedCode
	case "success":
		return SuccessCode
	}
	return ErrorOccurredCode
}
