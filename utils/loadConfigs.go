package utils

import (
	"errors"
	"os"
	"reflect"
	"strings"
)

var missedENV []string

func LoadConfig(config interface{}) error {
	Deep(config)

	if missedENV != nil {
		return errors.New("missed ENV configs: " + strings.Join(missedENV, ", "))
	}

	return nil
}

func Deep(str interface{}) {
	v := reflect.ValueOf(str).Elem()
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.Struct:
			Deep(f.Addr().Interface())
		case reflect.String:
			tag := t.Field(i).Tag.Get("env")
			if tag != "" {
				envByTag := os.Getenv(tag)
				if envByTag == "" {
					missedENV = append(missedENV, tag)
					continue
				}
				f.SetString(envByTag)
			}
		case reflect.Ptr:
			if f.Elem().Kind() == reflect.Struct {
				Deep(f.Interface())
			}
		}
	}
}
