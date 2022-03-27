package main

import tg "github.com/borghop/terrain-generator"

func main() {
	tm := tg.NewDefault()

	//saves picture of terrain as "map.png"
	tm.Generate()
}
