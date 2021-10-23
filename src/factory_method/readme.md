# 工厂方法模式

## 含义
工厂方法模式在父类（或者接口）中提供一个创建对象的方法，子类决定创建具体的对象类型，也就是一个工厂只生成一种对象，对应例子中，A工厂只能生成Acpu，B工厂只能生成Bcpu


## 模式结构
[![工厂方法模式](https://z3.ax1x.com/2021/10/23/52P1wF.md.png)](https://imgtu.com/i/52P1wF)

## 代码
```go
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


```


## 分析
### 优点
* 工厂方法将创建产品的过程与消费产品的过程进行分离，从而可以在不影响其他代码的情况下进行扩展产品类型（新建一种工厂和对应产品的类），所以具有扩展性强，低耦合等优点

### 缺点
* 如果产品种类特别多，需要定义很多个工厂类和产品类，系统中类的个数将成对增加，在一定程度上增加了系统的复杂性
* 创建对象的方法由于是每个工厂单独实现的，所以没有办法统一修改

## 适用场景
* 一个类不知道它所需要的对象的类：我们只需要知道创建对象的工厂类即可，不需要知道实例化对象的类，因为创建过程封装在了工厂的创建方法中
* 如果希望使用者能扩展你的库或者框架的内部组件，可以使用工厂方法模式

## 总结
工厂方法模式是对简单工厂模式的一个升级，简单工厂模式只有一个工厂，通过创建方法的参数在内部决定创建那种对象，而工厂方法模式一种对象对应一个工厂，即一种工厂只能创建一种对象，这将更有利于系统扩展，遵循了设计模式的开闭原则