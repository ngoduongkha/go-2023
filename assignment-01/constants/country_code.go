package constants

import "fmt"

type CountryCode string

const (
	VN CountryCode = "VN"
	US CountryCode = "US"
	CN CountryCode = "CN"
	DE CountryCode = "DE"
)

func ParseCountryCode(code string) (CountryCode, error) {
	switch code {
	case "VN":
		return VN, nil
	case "US":
		return US, nil
	case "CN":
		return CN, nil
	case "DE":
		return DE, nil
	default:
		supportedCodes := []CountryCode{VN, US, CN, DE}
		return "", fmt.Errorf("%s is not valid. Supported codes are %v", code, supportedCodes)
	}
}
