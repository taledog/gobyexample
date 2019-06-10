package main

type Explodes interface {
	Bang()
	Boom()
}

// Type Bomb implements Explodes
type Bomb struct{}

func (*Bomb) Bang() {}
func (Bomb) Boom()  {}

func main() {
	var bomb *Bomb = nil
	var explodes Explodes = bomb
	println(bomb, explodes) // '0x0 (0x10a7060,0x0)'
	if explodes != nil {
		println("Not nil!") // 'Not nil!' 我们为什么会走到这里?!?!
		explodes.Bang()     // 运行正常
		explodes.Boom()     // panic: value method main.Bomb.Boom called using nil *Bomb pointer
	} else {
		println("nil!") // 为什么没有在这里结束？
	}
}
