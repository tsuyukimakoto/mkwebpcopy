package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"golang.org/x/image/draw"

	"github.com/harukasan/go-libwebp/webp"
)

func contains(s string, targets []string) bool {
	for _, v := range targets {
		if v == s {
			return true
		}
	}
	return false
}

func glob(dir string, exts []string) ([]string, error) {

	files := []string{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if contains(filepath.Ext(path), exts) {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func ConvertToRGBA(img image.Image) (outimage image.Image) {
	bounds := img.Bounds()
	dest := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.CatmullRom.Scale(dest, dest.Bounds(), img, bounds, draw.Over, nil)
	return dest
}

func ConvertToRGBAIfNeed(img image.Image) (outimage image.Image) {
	switch img.(type) {
	case *image.RGBA:
		return img
	default:
		return ConvertToRGBA(img)
	}
}

func ReadImage(path string) (img image.Image) {
	io, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	ext := filepath.Ext(path)
	switch ext {
	case ".jpg":
		img, err := jpeg.Decode(io)
		if err != nil {
			panic(err)
		}
		return ConvertToRGBAIfNeed(img)
	case ".png":
		img, err := png.Decode(io)
		if err != nil {
			panic(err)
		}
		return ConvertToRGBAIfNeed(img)
	default:
		panic(fmt.Sprintf("%s is not acceptable(%s).", ext, path))
	}
}

func ConvertToWebP(path string) {
	dir_path, file_name := filepath.Split(path)
	file_name_len := len(file_name) - len(filepath.Ext(file_name))
	base_file_name := file_name[0:file_name_len]
	fmt.Println(path)

	img := ReadImage(path)

	f, err := os.Create(filepath.Join(dir_path, base_file_name) + ".webp")
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)
	defer func() {
		w.Flush()
		f.Close()
	}()
	config, err := webp.ConfigPreset(webp.PresetDefault, 90)
	if err := webp.EncodeRGBA(w, img.(*image.RGBA), config); err != nil {
		panic(err)
	}
}

func main() {
	root := flag.String("root", ".", "Root path")
	flag.Parse()
	exts := []string{".jpg", ".png"}
	files, err := glob(*root, exts)
	fmt.Println("Path:" + *root)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		ConvertToWebP(file)
	}
}
