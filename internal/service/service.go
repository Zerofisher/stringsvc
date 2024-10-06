package service

import (
	"strings"

	"github.com/Zerofisher/stringsvc/pkg/stringsvc"
)

type stringService struct{}

func NewStringService() stringsvc.StringService {
	return &stringService{}
}

func (s stringService) Uppercase(str string) (string, error) {
	if str == "" {
		return "", stringsvc.ErrEmpty
	}
	return strings.ToUpper(str), nil
}

func (s stringService) Count(str string) int {
	return len(str)
}
