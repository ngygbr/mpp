package validpackage

import "regexp"

var (
	amountregex = regexp.MustCompile("^[0-9]*$")

	cardNumberRegex = regexp.MustCompile("^[0-9]{16}$")
	holderNameRegex = regexp.MustCompile("^[\\D\\s]*[\\D\\s\\-',]*$")
	expDateRegex    = regexp.MustCompile("[0-1][1-9]+/+[0-9]")
	cvcRegex        = regexp.MustCompile("^[0-9]{3}$")

	TransactionIdRegex     = regexp.MustCompile("^[\\w]{20}$")
	paymentMethodTypeRegex = regexp.MustCompile("creditcard|ach|apple_pay|google_pay|apm")

	firstNameRegex   = regexp.MustCompile("^[\\D\\s]*$")
	lastNameRegex    = regexp.MustCompile("^[\\D\\s]*$")
	postalCodeRegex  = regexp.MustCompile("^[0-9]{4,6}$")
	cityRegex        = regexp.MustCompile("^[\\w\\s]{2,}$")
	addressLineRegex = regexp.MustCompile("^[\\w\\s-._]{2,}$")
	emailRegex       = regexp.MustCompile("[\\w-._]+@([\\w-._]+\\.)+[\\w-]{2,4}$")
	phoneRegex       = regexp.MustCompile("^[0-9]{5,12}$")

	achSECRegex         = regexp.MustCompile("ccd|ppd|tel|web")
	achRoutingRegex     = regexp.MustCompile("\\A([0-9]{9})\\z")
	achAccountRegex     = regexp.MustCompile("\\A([0-9]{4,17})\\z")
	achAccountTypeRegex = regexp.MustCompile("checking|savings")

	keyAES256Regex = regexp.MustCompile("([0-9a-zA-Z]{64})")
)
