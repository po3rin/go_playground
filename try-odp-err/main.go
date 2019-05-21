package main

import (
	"fmt"

	"github.com/pkg/errors"
)

// impliment error interface

type CustomErr struct {
	Msg    string
	Fields map[string]interface{}
}

func (c *CustomErr) Error() string {
	return c.Msg
}

func (c *CustomErr) GetFields() map[string]interface{} {
	return c.Fields
}

//custom error functions

func NewWithFields(msg string, fields ...func(*CustomErr)) error {
	c := &CustomErr{
		Msg:    msg,
		Fields: make(map[string]interface{}),
	}
	for _, setField := range fields {
		setField(c)
	}
	return c
}

func Field(key string, val interface{}) func(*CustomErr) {
	return func(c *CustomErr) {
		c.Fields[key] = val
	}
}

func GetFields(err error) map[string]interface{} {
	type fieldsGetter interface {
		GetFields() map[string]interface{}
	}

	type causer interface {
		Cause() error
	}

	fieldsMap := make(map[string]interface{})
	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
		customErr, ok := err.(fieldsGetter)
		if !ok {
			continue
		}
		filelds := customErr.GetFields()
		for k, v := range filelds {
			fieldsMap[k] = v
		}
	}

	return fieldsMap
}

func main() {
	// new error with Fields
	err := NewWithFields(
		"got error",
		Field("app_id", "aaa"),
		Field("process_id", "bbb"),
	)

	// can wrap !!
	err = errors.Wrap(err, "wrap!!")

	// get fields
	// map[app_id:aaa process_id:bbb]
	fmt.Println(GetFields(err))
}
