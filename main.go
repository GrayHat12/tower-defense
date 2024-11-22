package main

import (
	_ "image/png"

	"github.com/GrayHat12/tower-defense/src/logic"
	"github.com/GrayHat12/tower-defense/src/sprites"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	GameLogic logic.TowerDefenseGame
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for row := range logic.EXPECTED_DIM {
		for col := range logic.EXPECTED_DIM {
			switch g.GameLogic.InitialBoard[row*logic.EXPECTED_DIM+col] {
			case logic.DIRT:
				geom := ebiten.GeoM{}
				geom.Translate(float64(col*sprites.DirtSpriteValue.Bounds().Dx()), float64(row*sprites.DirtSpriteValue.Bounds().Dy()))
				screen.DrawImage(sprites.DirtSpriteValue, &ebiten.DrawImageOptions{
					GeoM: geom,
				})
				break
			case logic.D_GRASS:
				sprite, options := sprites.GrassSpriteValue.PartialDown()
				options.GeoM.Translate(float64(col*sprite.Bounds().Dx()), float64(row*sprite.Bounds().Dy()))
				screen.DrawImage(sprite, &options)
				break
			case logic.F_GRASS:
				sprite, options := sprites.GrassSpriteValue.Flat()
				options.GeoM.Translate(float64(col*sprite.Bounds().Dx()), float64(row*sprite.Bounds().Dy()))
				screen.DrawImage(sprite, &options)
				break
			case logic.L_GRASS:
				sprite, options := sprites.GrassSpriteValue.PartialLeft()
				options.GeoM.Translate(float64(col*sprite.Bounds().Dx()), float64(row*sprite.Bounds().Dy()))
				screen.DrawImage(sprite, &options)
				break
			case logic.R_GRASS:
				sprite, options := sprites.GrassSpriteValue.PartialRight()
				options.GeoM.Translate(float64(col*sprite.Bounds().Dx()), float64(row*sprite.Bounds().Dy()))
				screen.DrawImage(sprite, &options)
				break
			case logic.TOWER:
				break
			case logic.U_GRASS:
				sprite, options := sprites.GrassSpriteValue.PartialUp()
				options.GeoM.Translate(float64(col*sprite.Bounds().Dx()), float64(row*sprite.Bounds().Dy()))
				screen.DrawImage(sprite, &options)
				break
			case logic.CASTLE:
				geom := ebiten.GeoM{}
				geom.Translate(float64(col*sprites.CastleSpriteValue.GetCastle().Bounds().Dx()), float64(row*sprites.CastleSpriteValue.GetCastle().Bounds().Dy()))
				screen.DrawImage(sprites.CastleSpriteValue.GetCastle(), &ebiten.DrawImageOptions{
					GeoM: geom,
				})
			default:
				break
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(logic.EXPECTED_DIM*sprites.DirtSpriteValue.Bounds().Dx(), logic.EXPECTED_DIM*sprites.DirtSpriteValue.Bounds().Dy())
	g := &Game{}
	g.GameLogic.ProcedurallyGenerateLevel()
	// outsideWidth, outsideHeight = sprites.DirtSpriteValue.Bounds().Dx()*10, sprites.DirtSpriteValue.Bounds().Dy()*10
	// screenWidth, screenHeight = g.Layout(outsideWidth, outsideHeight)

	// scalingFactor = float64(screenWidth) / float64(outsideWidth)

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
