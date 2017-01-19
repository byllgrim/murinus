package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

const (
	MenuXOffset   int32 = 64
	MenuArrowSize int32 = 1280 / 2
)

var font *ttf.Font

type Menu struct {
	menuItems       []*MenuItem
	selectedElement int
}

type MenuItem struct {
	texture  *sdl.Texture
	src, dst *sdl.Rect
}

func (menu *Menu) Display(renderer *sdl.Renderer) {

	renderer.SetRenderTarget(nil)
	renderer.SetDrawColor(25, 25, 112, 255)
	renderer.Clear()

	first, last := menu.menuItems[0].dst.Y,
		menu.menuItems[len(menu.menuItems)-1].dst.Y
	last = newScreenHeight - last

	diff := (last + first) / 2
	diff -= first

	for i := range menu.menuItems {
		item := menu.menuItems[i]
		item.dst.X = newScreenWidth/2 - MenuXOffset
		item.dst.Y += diff
		renderer.Copy(item.texture, item.src, item.dst)
	}

	sel := menu.selectedElement % len(menu.menuItems)
	x := menu.menuItems[sel].dst.X - MenuXOffset/38
	y := menu.menuItems[sel].dst.Y + menu.menuItems[sel].dst.H/2 - 2
	renderer.SetDrawColor(255, 255, 255, 255)
	for i := int32(0); i < MenuArrowSize/85; i++ {
		renderer.DrawLine(int(x-i), int(y-i), int(x-i), int(y+i))
	}

	renderer.Present()
}

func GetMenus(renderer *sdl.Renderer) []*Menu {
	var err error
	if ttf.WasInit() || font != nil {
		panic("Should only be called once!")
	}
	e(ttf.Init())
	font, err = ttf.OpenFont("./font/AverageMono.ttf", 20)
	e(err)
	ret := make([]*Menu, 3)

	ret[0] = &Menu{[]*MenuItem{
		GetMenuItem("1 Player", screenHeight/2-80, renderer),
		GetMenuItem("2 Players", screenHeight/2-40, renderer),
		GetMenuItem("High-Scores", screenHeight/2, renderer),
		GetMenuItem("Options", screenHeight/2+40, renderer),
		GetMenuItem("Quit", screenHeight/2+80, renderer),
	}, 0}
	ret[1] = &Menu{[]*MenuItem{
		GetMenuItem("Beginner", screenHeight/2-40, renderer),
		GetMenuItem("Intermediate", screenHeight/2, renderer),
		GetMenuItem("Advanced", screenHeight/2+40, renderer),
	}, 0}
	ret[2] = &Menu{[]*MenuItem{
		GetMenuItem("Set Name", screenHeight/2-80, renderer),
		GetMenuItem("Highscores", screenHeight/2-40, renderer),
		GetMenuItem("Continue", screenHeight/2, renderer),
		GetMenuItem("Retry", screenHeight/2+40, renderer),
		GetMenuItem("Exit to menu", screenHeight/2+80, renderer),
	}, 0}

	return ret
}

func GetMenuItem(text string, y int32, renderer *sdl.Renderer) *MenuItem {
	texture, src, dst := GetText(text, sdl.Color{0, 190, 0, 255},
		screenWidth/2-screenWidth/8, y, renderer)
	return &MenuItem{texture, src, dst}
}

func GetText(text string, color sdl.Color, x, y int32,
	renderer *sdl.Renderer) (*sdl.Texture, *sdl.Rect, *sdl.Rect) {
	textSurface, err := font.RenderUTF8_Solid(text, color)
	e(err)
	defer textSurface.Free()

	texture, err := renderer.CreateTextureFromSurface(textSurface)
	e(err)
	src := &sdl.Rect{0, 0, textSurface.W, textSurface.H}
	dst := &sdl.Rect{x, y - src.H/2, src.W, src.H}
	return texture, src, dst
}

func (menu *Menu) Run(renderer *sdl.Renderer, input *Input) int {
	prevVal := int32(0)
	step := 0
	repeat := true
	input.mono.a.down = false
	input.mono.b.down = false
	for !input.mono.a.down && !quit {
		if input.mono.b.down {
			return -1
		}
		menu.Display(renderer)
		input.Poll()
		val := input.mono.upDown.Val()
		if val != prevVal || repeat {
			repeat = false
			if prevVal != val {
				step = 0
			}
			prevVal = val
			if val > 0 {
				menu.selectedElement = (menu.selectedElement +
					1) % len(menu.menuItems)
			} else if val < 0 {
				menu.selectedElement = (menu.selectedElement +
					len(menu.menuItems) - 1) % len(menu.menuItems)
			}
		}
		step++
		if step > 20 && step%3 == 0 {
			repeat = true
		}
	}
	return menu.selectedElement
}

func (menu *Menu) Free() {
	for i := 0; i < len(menu.menuItems); i++ {
		if menu.menuItems[i].texture != nil {
			menu.menuItems[i].texture.Destroy()
		}
	}
	menu.menuItems = nil
}
