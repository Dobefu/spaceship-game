//go:build ignore

//kage:unit pixels

package shaders

const glowSize = 256.0 / 512.0
const glowStrength = 5.0
const colorBoost = 20.0

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	originalColor := imageSrc0At(texCoord)
	boostedColor := originalColor * colorBoost
	blur := vec4(0, 0, 0, 0)

	blur += imageSrc0At(vec2(texCoord.x-3.0*glowSize, texCoord.y)) * 0.1
	blur += imageSrc0At(vec2(texCoord.x-2.0*glowSize, texCoord.y)) * 0.15
	blur += imageSrc0At(vec2(texCoord.x-glowSize, texCoord.y)) * 0.25
	blur += imageSrc0At(vec2(texCoord.x+glowSize, texCoord.y)) * 0.25
	blur += imageSrc0At(vec2(texCoord.x+2.0*glowSize, texCoord.y)) * 0.15
	blur += imageSrc0At(vec2(texCoord.x+3.0*glowSize, texCoord.y)) * 0.1

	blur += imageSrc0At(vec2(texCoord.x, texCoord.y-3.0*glowSize)) * 0.1
	blur += imageSrc0At(vec2(texCoord.x, texCoord.y-2.0*glowSize)) * 0.15
	blur += imageSrc0At(vec2(texCoord.x, texCoord.y-glowSize)) * 0.25
	blur += imageSrc0At(vec2(texCoord.x, texCoord.y+glowSize)) * 0.25
	blur += imageSrc0At(vec2(texCoord.x, texCoord.y+2.0*glowSize)) * 0.15
	blur += imageSrc0At(vec2(texCoord.x, texCoord.y+3.0*glowSize)) * 0.1

	blur += imageSrc0At(vec2(texCoord.x-2.0*glowSize, texCoord.y-2.0*glowSize)) * 0.1
	blur += imageSrc0At(vec2(texCoord.x+2.0*glowSize, texCoord.y-2.0*glowSize)) * 0.1
	blur += imageSrc0At(vec2(texCoord.x-2.0*glowSize, texCoord.y+2.0*glowSize)) * 0.1
	blur += imageSrc0At(vec2(texCoord.x+2.0*glowSize, texCoord.y+2.0*glowSize)) * 0.1

	blur = blur / 2.5

	finalColor := boostedColor + blur*glowStrength

	finalColor = vec4(
		finalColor.r/(finalColor.r+1.0),
		finalColor.g/(finalColor.g+1.0),
		finalColor.b/(finalColor.b+1.0),
		finalColor.a,
	)

	return finalColor
}
