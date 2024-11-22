package sprites

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type CastleSprite struct {
	topRoof    *ebiten.Image
	tileCastle *ebiten.Image
}

func (sprite CastleSprite) GetCastle() *ebiten.Image {
	image := sprite.tileCastle
	// options := ebiten.DrawImageOptions{
	// 	Blend: ebiten.BlendCopy,
	// }
	// options.GeoM.Rotate(180)
	for x := range sprite.topRoof.Bounds().Dx() {
		for y := range sprite.topRoof.Bounds().Dy() {
			color := sprite.topRoof.At(x, y)
			// image.Set(sprite.topRoof.Bounds().Dx()-x, sprite.topRoof.Bounds().Dy()-y, color)
			image.Set(x, y, color)
		}
	}
	return image
}

var CastleSpriteValue = CastleSprite{
	topRoof: cropSprites(mustLoadImage("assets/kenney_block-pack/PNG/Default (64px)/tileCastle_topRoof.png"), &CropOptions{
		Rectangle: image.Rectangle{
			Min: image.Point{
				X: 0,
				Y: 100 - 54,
			},
			Max: image.Point{
				X: 64,
				Y: 100,
			},
		},
	}),
	tileCastle: cropSprites(mustLoadImage("assets/kenney_block-pack/PNG/Default (64px)/tileBuilding_stone.png"), nil),
}
