package tg

import (
	"math/rand"
	"time"

	"github.com/fogleman/gg"
	"github.com/ojrac/opensimplex-go"
)

func init() {
	//set rand seed based on time
	rand.Seed(time.Now().UnixNano())
}

type color struct {
	r int
	g int
	b int
}

type TerrainMap struct {
	scale        float64
	lacunarity   float64
	persistance  float64
	octaves      int
	deepWater    color
	mediumWater  color
	shallowWater color
	sand         color
	lowGrass     color
	highGrass    color
	rock         color
	snow         color
}

func NewDefault() *TerrainMap {
	//default values for terrain map generation
	return &TerrainMap{
		scale:        125.0,
		lacunarity:   2.0,
		persistance:  0.5,
		octaves:      5,
		deepWater:    color{r: 8, g: 50, b: 201},
		mediumWater:  color{r: 8, g: 66, b: 201},
		shallowWater: color{r: 8, g: 114, b: 201},
		sand:         color{r: 255, g: 228, b: 110},
		lowGrass:     color{r: 23, g: 140, b: 22},
		highGrass:    color{r: 23, g: 120, b: 22},
		rock:         color{r: 55, g: 63, b: 66},
		snow:         color{r: 230, g: 232, b: 237},
	}
}

//generate a terrain map
func (tm *TerrainMap) Generate() {
	//init noise and new graphic context
	noise := opensimplex.New(rand.Int63())
	dc := gg.NewContext(750, 500)

	for y := 0; y < dc.Height(); y++ {
		for x := 0; x < dc.Width(); x++ {
			//sample x and y and apply scale
			sampleX := float64(x) / tm.scale
			sampleY := float64(y) / tm.scale

			//init values for octave calculation
			frequency := 1.0
			amplitude := 1.0
			normalizeOctaves := 0.0
			total := 0.0

			//octave calculation
			for i := 0; i < tm.octaves; i++ {
				total += noise.Eval2(sampleX*frequency, sampleY*frequency) * amplitude
				normalizeOctaves += amplitude
				amplitude *= tm.persistance
				frequency *= tm.lacunarity
			}

			//normalize to -1 to 1, and then from 0 to 1 (this is for the ability to use grayscale, if using colors could keep from -1 to 1)
			depth := (total/normalizeOctaves + 1) / 2

			//color terrain
			//deep water -> medium water -> shallow water -> sand -> low grass -> high grass -> rock -> snow
			if depth > 0.70 {
				dc.SetRGB255(tm.deepWater.r, tm.deepWater.g, tm.deepWater.b)
			} else if depth > 0.56 {
				dc.SetRGB255(tm.mediumWater.r, tm.mediumWater.g, tm.mediumWater.b)
			} else if depth > 0.52 {
				dc.SetRGB255(tm.shallowWater.r, tm.shallowWater.g, tm.shallowWater.b)
			} else if depth > 0.50 {
				dc.SetRGB255(tm.sand.r, tm.sand.g, tm.sand.b)
			} else if depth > 0.41 {
				dc.SetRGB255(tm.lowGrass.r, tm.lowGrass.g, tm.lowGrass.b)
			} else if depth > 0.31 {
				dc.SetRGB255(tm.highGrass.r, tm.highGrass.g, tm.highGrass.b)
			} else if depth > 0.23 {
				dc.SetRGB255(tm.rock.r, tm.rock.g, tm.rock.b)
			} else {
				dc.SetRGB255(tm.snow.r, tm.snow.g, tm.snow.b)
			}

			//set current pixel on image to assigned color
			dc.SetPixel(x, y)
		}
	}
	//save terrain map image
	dc.SavePNG("map.png")
}

func (tm *TerrainMap) SetGenerationModifiers(scale float64, lacunarity float64, persistance float64, octaves int) {
	tm.scale = scale
	tm.lacunarity = lacunarity
	tm.persistance = persistance
	tm.octaves = octaves
}

func (tm *TerrainMap) SetDeepWaterColor(r int, g int, b int) {
	tm.deepWater.r = r
	tm.deepWater.g = g
	tm.deepWater.b = b
}

func (tm *TerrainMap) SetMediumWaterColor(r int, g int, b int) {
	tm.mediumWater.r = r
	tm.mediumWater.g = g
	tm.mediumWater.b = b
}

func (tm *TerrainMap) SetShallowWaterColor(r int, g int, b int) {
	tm.shallowWater.r = r
	tm.shallowWater.g = g
	tm.shallowWater.b = b
}

func (tm *TerrainMap) SetSandColor(r int, g int, b int) {
	tm.sand.r = r
	tm.sand.g = g
	tm.sand.b = b
}

func (tm *TerrainMap) SetLowGrassColor(r int, g int, b int) {
	tm.lowGrass.r = r
	tm.lowGrass.g = g
	tm.lowGrass.b = b
}

func (tm *TerrainMap) SetHighGrassColor(r int, g int, b int) {
	tm.highGrass.r = r
	tm.highGrass.g = g
	tm.highGrass.b = b
}

func (tm *TerrainMap) SetRockColor(r int, g int, b int) {
	tm.rock.r = r
	tm.rock.g = g
	tm.rock.b = b
}

func (tm *TerrainMap) SetSnowColor(r int, g int, b int) {
	tm.snow.r = r
	tm.snow.g = g
	tm.snow.b = b
}
