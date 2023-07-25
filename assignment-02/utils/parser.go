package utils

import (
	"fmt"
	"strconv"
)

type Parser interface {
	ParseInt([]string) ([]int, error)
	ParseFloat([]string) ([]float64, error)
}

type parser struct{}

func NewParser() Parser {
	return &parser{}
}

func (p *parser) ParseInt(args []string) ([]int, error) {
	var arr = make([]int, len(args))
	for i, item := range args {
		parsedValue, err := strconv.Atoi(item)
		if err != nil {
			return nil, fmt.Errorf("unable to convert '%s' to int", item)
		} else {
			arr[i] = parsedValue
		}
	}
	return arr, nil
}

func (p *parser) ParseFloat(args []string) ([]float64, error) {
	var arr = make([]float64, len(args))
	for i, item := range args {
		parsedValue, err := strconv.ParseFloat(item, 64)
		if err != nil {
			return nil, fmt.Errorf("unable to convert '%s' to float", item)
		} else {
			arr[i] = parsedValue
		}
	}
	return arr, nil
}
