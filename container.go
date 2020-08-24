package container

import (
	"errors"
	"reflect"
)

type Container struct {
	Objs map[string]interface{}
}

func NewContainer() *Container {
	c := &Container{}
	c.Objs = map[string]interface{}{}
	return c
}

func (p *Container) Add(name string, ins interface{}) {
	p.Objs[name] = ins
}

func (p *Container) Run() error {
	for key, v := range p.Objs {
		err := p.initIns(key, v)
		if err != nil {
			return err
		}
	}
	for _, v := range p.Objs {
		if ins, ok := v.(Obj); ok {
			ins.Init()
		}
	}
	return nil
}

func (p *Container) initIns(name string, ins interface{}) error {
	insValue := reflect.ValueOf(ins).Elem()
	insType := insValue.Type()
	for i := 0; i < insType.NumField(); i++ {
		tag := insType.Field(i).Tag.Get("inject")
		if v, ok := p.Objs[tag]; ok {
			insValue.Field(i).Set(reflect.ValueOf(v))
		} else {
			return errors.New("not found inject obj :" + tag)
		}
	}
	return nil
}
