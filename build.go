package main

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"math"
	"os"
	"path"

	"github.com/fogleman/gg"
)

func main() {
	files, err := ioutil.ReadDir("./source")
	if err != nil {
		panic(err)

	}

	const iconSize = 32

	cols := math.Ceil(math.Sqrt(float64(len(files))))
	rows := math.Ceil(float64(len(files)) / cols)
	face := gg.NewContext(int(rows)*iconSize, int(cols)*iconSize)
	for i, f := range files {
		x := i % int(rows)
		y := int(math.Floor(float64(i) / rows))
		unit, err := gg.LoadPNG(path.Join("source", f.Name()))
		if err != nil {
			panic(err)
		}
		face.Push()
		face.Translate(float64(x)*iconSize, float64(y)*iconSize)
		face.DrawImage(unit, 0, 0)
		face.Pop()
	}

	const scale = 2
	face2 := gg.NewContext(face.Width()*scale, face.Height()*scale)
	for x := 0; x < face2.Width(); x++ {
		for y := 0; y < face2.Height(); y++ {
			face2.SetColor(face.Image().At(int(x/scale), int(y/scale)))
			face2.SetPixel(x, y)
		}
	}
	face2.SavePNG("./build/face.png")

	//
	zipFile, err := os.Create("./build/RPGCharactersPack.zip")
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, f := range files {
		w, err := zipWriter.Create(f.Name())
		if err != nil {
			panic(err)
		}
		r, err := os.Open(path.Join("source", f.Name()))
		if err != nil {
			panic(err)
		}
		io.Copy(w, r)
		r.Close()
	}
}
