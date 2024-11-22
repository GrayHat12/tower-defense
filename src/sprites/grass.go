package sprites

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type GrassSprite struct {
	flat    *ebiten.Image
	partial *ebiten.Image
}

func (sprite GrassSprite) Flat() (*ebiten.Image, ebiten.DrawImageOptions) {
	return sprite.flat, ebiten.DrawImageOptions{
		GeoM: ebiten.GeoM{},
	}
}

func (sprite GrassSprite) PartialUp() (*ebiten.Image, ebiten.DrawImageOptions) {
	geometry := ebiten.GeoM{}
	geometry.Rotate(0)
	drawOptions := ebiten.DrawImageOptions{
		GeoM: geometry,
	}
	return sprite.partial, drawOptions
}

func (sprite GrassSprite) PartialRight() (*ebiten.Image, ebiten.DrawImageOptions) {
	geometry := ebiten.GeoM{}
	geometry.Rotate(90)
	drawOptions := ebiten.DrawImageOptions{
		GeoM: geometry,
	}
	return sprite.partial, drawOptions
}

func (sprite GrassSprite) PartialDown() (*ebiten.Image, ebiten.DrawImageOptions) {
	geometry := ebiten.GeoM{}
	geometry.Rotate(180)
	drawOptions := ebiten.DrawImageOptions{
		GeoM: geometry,
	}
	return sprite.partial, drawOptions
}

func (sprite GrassSprite) PartialLeft() (*ebiten.Image, ebiten.DrawImageOptions) {
	geometry := ebiten.GeoM{}
	geometry.Rotate(270)
	drawOptions := ebiten.DrawImageOptions{
		GeoM: geometry,
	}
	return sprite.partial, drawOptions
}

var GrassSpriteValue = GrassSprite{
	flat:    cropSprites(mustLoadImage("assets/kenney_block-pack/PNG/Default (64px)/tileGrass_slope.png"), nil),
	partial: cropSprites(mustLoadImage("assets/kenney_block-pack/PNG/Default (64px)/tileGrass.png"), nil),
}

var GrassSpriteValyes = [7]*ebiten.Image{
	mustLoadImage("assets/src/sprites/assets/kenney_hexagon-pack/PNG/Tiles/Terrain/Grass/grass_10.png"),
	mustLoadImage("assets/src/sprites/assets/kenney_hexagon-pack/PNG/Tiles/Terrain/Grass/grass_11.png"),
	mustLoadImage("assets/src/sprites/assets/kenney_hexagon-pack/PNG/Tiles/Terrain/Grass/grass_12.png"),
	mustLoadImage("assets/src/sprites/assets/kenney_hexagon-pack/PNG/Tiles/Terrain/Grass/grass_13.png"),
	mustLoadImage("assets/src/sprites/assets/kenney_hexagon-pack/PNG/Tiles/Terrain/Grass/grass_14.png"),
	mustLoadImage("assets/src/sprites/assets/kenney_hexagon-pack/PNG/Tiles/Terrain/Grass/grass_15.png"),
	mustLoadImage("assets/src/sprites/assets/kenney_hexagon-pack/PNG/Tiles/Terrain/Grass/grass_16.png"),
}
