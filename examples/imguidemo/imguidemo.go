package main

import (
	g "github.com/pokemium/giu-omega"
	"github.com/pokemium/imgui-go-omega"
)

func loop() {
	imgui.ShowDemoWindow(nil)
}

func main() {
	wnd := g.NewMasterWindow("Widgets", 1024, 768, 0)
	wnd.Run(loop)
}
