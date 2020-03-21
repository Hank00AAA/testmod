package testmod

import (
	"fmt"
	"errors"
)

func Hi(name string, lang string) (string, error) {
	switch(lang) {
	case "en":
		return fmt.Sprintf("Hi, %s!", name), nil
	case "ch":
		return fmt.Sprintf("你好, %s!", name), nil
	}
	return "", errors.New("unknow language!")
}

