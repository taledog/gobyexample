// _接口(Interfaces)_ 是命名了的方法签名(signatures)的集合。

package main

import "fmt"
import "math"

// 这里是一个几何体的基本接口。
type geometry interface {
	area() float64
	perim() float64
}

// 在我们的例子中，我们将在类型 `rect` 和 `circle` 上实现
// 这个接口
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

type round struct {
	width float64
}

// 要在 Go 中实现一个接口，我们就需要实现接口中的所有
// 方法。这里我们在 `rect` 上实现了 `geometry` 接口。
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// `circle` 的实现。
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// round 实现
func (r round) perim() float64 {
	return r.width * r.width
}

// 如果一个变量具有接口类型，那么我们可以调用指定接口中的
// 方法。这里有一个通用的 `measure` 函数，利用它来在任
// 何的 `geometry` 上工作。
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// ============================================
type value interface {
	Show()
}

type test_value struct{}

// 如果这里注释掉，下面的下面的语句会报错。
func (value *test_value) Show() { fmt.Printf("test \n") }

// 使用类似方法处理，编译的时候会报错，由于test_value没有实现value接口所有的方法
var _ value = new(test_value)

// ============================================

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// 由于round没有实现geometry接口所有的方法，所以编译时就会报错
	//var _ geometry = round{}

	// 由于round实现geometry接口所有的方法，所以编译时不会报错
	var _ geometry = circle{}

	// 结构体类型 `circle` 和 `rect` 都实现了 `geometry`
	// 接口，所以我们可以使用它们的实例作为 `measure` 的参数。
	measure(r)
	measure(c)
}
