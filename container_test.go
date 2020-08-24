package container

import (
	"fmt"
	"testing"
)

type Subnet struct {
	Port *Port `inject:"port"`
}

func (p *Subnet) Init() error {
	fmt.Print("subnet init success")
	return nil
}

type Router struct {
	Port *Port `inject:"port"`
}

func (p *Router) Init() error {
	fmt.Print("router init success")
	return nil
}

type Port struct {
}

func (p *Port) Init() error {
	fmt.Print("port init success")
	return nil
}

type Network struct {
	Subnet *Subnet `inject:"subnet"`
	Router *Router `inject:"router"`
	Port   *Port   `inject:"port"`
}

func (p *Network) Init() error {
	fmt.Print("network init success")
	return nil
}

func TestContainer_Run(t *testing.T) {
	c := NewContainer()
	c.Add("port", &Port{})
	c.Add("subnet", &Subnet{})
	c.Add("router", &Router{})
	c.Add("network", &Network{})
	c.Run()
}
