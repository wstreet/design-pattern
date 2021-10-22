package factory_method

import (
	"fmt"
)

type Cpu interface {
	calculate() string
}

//-------Acpu类---------------
type Acpu struct {
	core int
}

func NewAcpu() *Acpu {
	return &Acpu{2}
}

func (a Acpu) calculate() string {
	result := fmt.Sprintf("Acpu has %d core calculate result", a.core)

	return result
}

//-------Bcpu类---------------
type Bcpu struct {
	core int
}

func NewBcpu() *Bcpu {
	return &Bcpu{4}
}

func (b Bcpu) calculate() string {
	result := fmt.Sprintf("Bcpu has %d core calculate result", b.core)
	return result
}

// 工厂抽象
type CpuFactory interface {
	create() Cpu
}

// Acpu工厂
type AcpuFactory struct {
}

func NewAcpuFactory() *AcpuFactory {
	return &AcpuFactory{}
}

func (a AcpuFactory) create() Cpu {
	return NewAcpu()
}

// Bcpu工厂
type BcpuFactory struct {
}

func NewBcpuFactory() *BcpuFactory {
	return &BcpuFactory{}
}

func (b BcpuFactory) create() Cpu {
	return NewBcpu()
}

func main() {
	af := NewAcpuFactory()
	a1 := af.create()
	a2 := af.create()
	a1.calculate()
	a2.calculate()

	bf := NewBcpuFactory()
	b1 := bf.create()
	b2 := bf.create()
	b1.calculate()
	b2.calculate()

}
