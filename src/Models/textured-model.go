package Models

import "go-game/src/Textures"

type TexturedModel struct {
	RawModel *RawModel
	Texture  Textures.ModelTexture
}
