package pdf

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	qr "github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func toBase64(filename string) string {
	// Read the entire file into a byte slice
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return base64.StdEncoding.EncodeToString(bytes)
}

func CreatePdfFile(qrcode string) (*bytes.Buffer, error) {
	qrc, err := qr.NewWith(qrcode)
	if err != nil {
		return nil, err
	}

	pwd, _ := os.Getwd()

	fileID := pwd + "/" + uuid.NewString() + ".png"

	w, err := standard.New(fileID)
	if err != nil {
		return nil, err
	}

	qrc.Save(w)

	bFile := toBase64(fileID)

	_pdf := pdf.NewMarotoCustomSize(consts.Portrait, "A6", "mm", 105.0, 148.0)
	redColor := color.Color{
		Red:   255,
		Green: 0,
		Blue:  0,
	}
	_pdf.RegisterHeader(func() {
		_pdf.Row(25, func() {
			_pdf.Col(13, func() {
				_pdf.SetBackgroundColor(redColor)
				_pdf.Text("HEADER", props.Text{
					Size: 10,
				})
			})
		})
	})

	_pdf.Row(55, func() {
		_pdf.Base64Image(bFile, consts.Jpg, props.Rect{
			Center: true,
		})
	})

	_pdf.Row(5, func() {
		_pdf.Text("Scanner le qrcode ou le code barre pour payer", props.Text{
			Align: consts.Center,
			Size:  6,
		})
	})

	_pdf.Row(5, func() {
		_pdf.Barcode(qrcode, props.Barcode{
			Center: true,
		})
	})

	_pdf.RegisterFooter(func() {
		_pdf.Line(1)
		_pdf.Row(13, func() {
			_pdf.Col(10, func() {
				_pdf.FileImage(pwd+"/assets/google-play-badge.png", props.Rect{
					Percent: 57,
					Left:    50,
				})
			})

			_pdf.Col(7, func() {
				_pdf.FileImage(pwd+"/assets/Download_on_the_App_Store_Badge_FR_RGB_blk_100517.png", props.Rect{
					Percent: 39,
					Top:     1.2,
				})
			})
		})

		_pdf.Row(3, func() {
			_pdf.Col(12, func() {
				_pdf.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.", props.Text{
					Align: consts.Center,
					Size:  6,
				})
			})
		})

	})

	os.Remove(fileID)

	buff, err := _pdf.Output()
	if err != nil {
		return nil, err
	}

	return &buff, err

}
