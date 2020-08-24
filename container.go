package container

import (
	"reflect"
)

type Container struct {
	Instances map[string]interface{}
}

func NewContainer() *Container {
	c := &Container{}
	c.Instances = map[string]interface{}{}
	return c
}

func (p *Container) Add(name string, ins interface{}) {
	p.Instances[name] = ins
}

func (p *Container) Run() error {
	for key, v := range p.Instances {
		err := p.initIns(key, v)
		if err != nil {
			return err
		}
	}
	for _, v := range p.Instances {
		if bType, ok := v.(Instance); ok {
			bType.Init()
		}
	}
	return nil
}

func (p *Container) initIns(name string, ins interface{}) error {
	insValue := reflect.ValueOf(ins).Elem()
	insType2 := insValue.Type()
	for i := 0; i < insType2.NumField(); i++ {
		tag := insType2.Field(i).Tag.Get("inject")
		if v, ok := p.Instances[tag]; ok {
			insValue.Field(i).Set(reflect.ValueOf(v))
		}
	}
	return nil
}
