package ingo

import "reflect"

type Container struct {
	deps map[reflect.Type]reflect.Value
}

func NewContainer() *Container {
	return &Container{
		deps: make(map[reflect.Type]reflect.Value),
	}
}

func (c *Container) Register(t reflect.Type, v interface{}) {
	c.deps[t] = reflect.ValueOf(v)
}

func (c *Container) Execute(f interface{}) ([]interface{}, error) {
	fn := reflect.ValueOf(f)
	arguments := []reflect.Value{}
	for i := 0; i < fn.Type().NumIn(); i++ {
		t := fn.Type().In(i)
		value := c.deps[t]
		if value.Kind() == reflect.Func {
			r, err := c.Execute(value.Interface())
			if err != nil {
				return nil, err
			}
			value = reflect.ValueOf(r[0])
		}
		arguments = append(arguments, value)
	}

	results := fn.Call(arguments)

	ret := []interface{}{}
	for i := 0; i < fn.Type().NumOut(); i++ {
		ret = append(ret, results[i].Interface())
	}
	return ret, nil
}
