package main

import (
	"bytes"
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/gofont/goregular"
)

type scene struct{
	font *common.Font
}

func (s *scene) Type() string {
	return "scene"
}

func (s *scene) Preload() {
	if err := engo.Files.LoadReaderData("goregular.ttf", bytes.NewReader(goregular.TTF)); err != nil {
		panic(err)
	}
	s.font = &common.Font{
		URL:  "goregular.ttf",
		FG:   colornames.White,
		Size: 24,
	}
	if err := s.font.CreatePreloaded(); err != nil {
		panic(err)
	}
}

type text struct {
	ecs.BasicEntity
	*common.RenderComponent
	*common.SpaceComponent
}

func (s *scene) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)
	w.AddSystem(&common.RenderSystem{})

	t := text{ecs.NewBasic(), &common.RenderComponent{
		Drawable: common.Text{
			Font: s.font,
			Text: "qwe",
		},
	}, &common.SpaceComponent{}}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&t.BasicEntity, t.RenderComponent, t.SpaceComponent)
		}
	}
}

func main() {

	engoOptions := engo.RunOptions{
		Title:    "test",
		Width:    800,
		Height:   600,
		FPSLimit: 60,
		VSync:    true,
	}
	engo.Run(engoOptions, &scene{})

}
