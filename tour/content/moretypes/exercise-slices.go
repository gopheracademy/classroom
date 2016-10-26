// +build no-build OMIT

package main

import "github.com/gophertrain/tour/pic"

func Pic(dx, dy int) [][]uint8 {
}

func main() {
	pic.Show(Pic)
}
