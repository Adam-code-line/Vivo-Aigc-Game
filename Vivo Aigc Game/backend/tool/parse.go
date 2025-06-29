package tool

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseQuery(target interface{}, ctx *gin.Context) error {
	v := reflect.ValueOf(target).Elem()
	t := v.Type()
	if t.Kind() != reflect.Struct {
		return errors.New("target must be a struct")
	}
	n := t.NumField()
	for i := 0; i < n; i++ {
		name := t.Field(i).Tag.Get("json")
		must := t.Field(i).Tag.Get("bg") == "must"
		setType := t.Field(i).Type.Kind()
		setValue := ctx.Query(name)
		if setValue == "" {
			if must {
				return errors.New(name + " must be set")
			}
			continue
		}
		switch setType {
		case reflect.String:
			v.Field(i).SetString(setValue)
		case reflect.Bool:
			if setValue == "true" {
				v.Field(i).SetBool(true)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			setNumber, _ := strconv.ParseInt(setValue, 10, 64)
			v.Field(i).SetInt(int64(setNumber))
		case reflect.Float32:
			setNumber, _ := strconv.ParseFloat(setValue, 64)
			v.Field(i).SetFloat(setNumber)
		}
	}
	return nil
}

func ParseForm(target interface{}, ctx *gin.Context) error {
	v := reflect.ValueOf(target).Elem()
	t := v.Type()
	if t.Kind() != reflect.Struct {
		return errors.New("target must be a struct")
	}
	n := t.NumField()
	for i := 0; i < n; i++ {
		name := t.Field(i).Tag.Get("json")
		must := t.Field(i).Tag.Get("bg") == "must"
		setType := t.Field(i).Type.Kind()
		setValue := ctx.PostForm(name)
		if setValue == "" {
			if must {
				return errors.New(name + " must be set")
			}
			continue
		}
		switch setType {
		case reflect.String:
			v.Field(i).SetString(setValue)
		case reflect.Bool:
			if setValue == "true" {
				v.Field(i).SetBool(true)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			setNumber, _ := strconv.ParseInt(setValue, 10, 64)
			v.Field(i).SetInt(int64(setNumber))
		case reflect.Float32:
			setNumber, _ := strconv.ParseFloat(setValue, 64)
			v.Field(i).SetFloat(setNumber)
		}
	}
	return nil
}
