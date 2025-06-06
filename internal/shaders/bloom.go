//go:build ignore

//kage:unit pixels

package shaders

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	posX := position.x
	posY := imageDstSize().y - position.y

	srcColor := imageSrc0At(vec2(posX, posY))

	bloomColour := srcColor
	numSamples := 1.0

	for x := -8.0; x <= 8.0; x += 1.0 {
		for y := -8.0; y <= 8.0; y += 1.0 {
			addColor := imageSrc0At(vec2(posX+x, posY+y))

			if max(addColor.r, max(addColor.g, addColor.b)) <= 0.3 {
				continue
			}

			dist := length(vec2(x, y)) + 1.0
			glowColor := max((addColor*128.0)/pow(dist, 2.0), vec4(0.0))

			if max(glowColor.r, max(glowColor.g, glowColor.b)) <= 0.0 {
				continue
			}

			bloomColour += glowColor
			numSamples += 1.0
		}
	}

	bloomColour = bloomColour / numSamples

	result := mix(srcColor, bloomColour, 0.05)
	result.a = 1.0

	return result
}
