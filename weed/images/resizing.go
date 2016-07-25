package images

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"

	"github.com/chrislusf/seaweedfs/weed/glog"
	"github.com/disintegration/imaging"
)

func Resized_Fit(ext string, data []byte, size string) (resized []byte, w int, h int) {
	
	width,height := getSize(size)
	
	if width == 0 && height == 0 {
		return data, 0, 0
	}
	
	srcImage, _, err := image.Decode(bytes.NewReader(data))
	if err == nil {
		bounds := srcImage.Bounds()
		var dstImage *image.NRGBA
		if bounds.Dx() <= width && bounds.Dy() <= height {
			return data, bounds.Dx(), bounds.Dy()
		} else {
			dstImage = imaging.Fit(srcImage, width, height, imaging.Lanczos)
		}
		var buf bytes.Buffer
		switch ext {
		case ".png":
			png.Encode(&buf, dstImage)
		case ".jpg", ".jpeg":
			jpeg.Encode(&buf, dstImage, nil)
		case ".gif":
			gif.Encode(&buf, dstImage, nil)
		}
		return buf.Bytes(), dstImage.Bounds().Dx(), dstImage.Bounds().Dy()
	} else {
		glog.Error(err)
	}
	return data, 0, 0
}

func getSize(size string) (int, int) {
	switch size {
	case "pico":
		return 16, 16
	case "icon":
		return 32, 32
	case "thumb":
		return 50, 50
	case "small":
		return 100, 100
	case "compact":
		return 160, 160
	case "medium":
		return 240, 240
	case "large":
		return 480, 480
	case "grande":
		return 600, 600
	case "1024x1024":
		return 1024, 1024
	case "2048x2048":
		return 2048, 2048
	case "master":
		return 2048, 2048
	}
	return 0, 0
}

func Resized(ext string, data []byte, width, height int) (resized []byte, w int, h int) {
	if width == 0 && height == 0 {
		return data, 0, 0
	}
	srcImage, _, err := image.Decode(bytes.NewReader(data))
	if err == nil {
		bounds := srcImage.Bounds()
		var dstImage *image.NRGBA
		if bounds.Dx() > width && width != 0 || bounds.Dy() > height && height != 0 {
			if width == height && bounds.Dx() != bounds.Dy() {
				dstImage = imaging.Thumbnail(srcImage, width, height, imaging.Lanczos)
				w, h = width, height
			} else {
				dstImage = imaging.Resize(srcImage, width, height, imaging.Lanczos)
			}
		} else {
			return data, bounds.Dx(), bounds.Dy()
		}
		var buf bytes.Buffer
		switch ext {
		case ".png":
			png.Encode(&buf, dstImage)
		case ".jpg", ".jpeg":
			jpeg.Encode(&buf, dstImage, nil)
		case ".gif":
			gif.Encode(&buf, dstImage, nil)
		}
		return buf.Bytes(), dstImage.Bounds().Dx(), dstImage.Bounds().Dy()
	} else {
		glog.Error(err)
	}
	return data, 0, 0
}
