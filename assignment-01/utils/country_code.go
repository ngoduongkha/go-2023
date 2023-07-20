package utils

import (
	"fmt"
	"github.com/ngoduongkha/go-2023/assignment-01/constants"
)

type Reorder interface {
	Do(firstName, lastName string, middleName ...string) string
}

type ReorderFactory struct {
}

func New() *ReorderFactory {
	return &ReorderFactory{}
}

func (f *ReorderFactory) GetStrategy(code constants.CountryCode) (Reorder, error) {
	switch code {
	case constants.CN, constants.VN:
		return &NaughtyOrder{}, nil
	case constants.DE, constants.US:
		return &PreferredOrder{}, nil
	default:
		return nil, fmt.Errorf("%s is not handled", code)
	}
}

type PreferredOrder struct {
}

func (p *PreferredOrder) Do(firstName, lastName string, middleName ...string) string {
	if len(middleName) == 0 {
		return firstName + " " + lastName
	}

	res := firstName + " "
	for _, name := range middleName {
		res += name + " "
	}
	return res + lastName
}

type NaughtyOrder struct {
}

func (n *NaughtyOrder) Do(firstName, lastName string, middleName ...string) string {
	if len(middleName) == 0 {
		return lastName + " " + firstName
	}

	res := lastName + " "
	for _, name := range middleName {
		res += name + " "
	}
	return res + firstName
}
