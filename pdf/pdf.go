package pdf

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

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

var rootPath string

const projectDirName = "server"

func init() {
	if os.Getenv("env") != "HEROKU" {
		re := regexp.MustCompile(`^(.*` + os.Getenv("DIRNAME") + `)`)
		cwd, _ := os.Getwd()
		rootPath = string(re.Find([]byte(cwd)))
	} else {
		re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
		cwd, _ := os.Getwd()
		rootPath = string(re.Find([]byte(cwd)))
	}
}

func CreatePdfFile(qrcode string, enterpriseName string) (*bytes.Buffer, error) {
	qrc, err := qr.NewWith(qrcode)
	if err != nil {
		return nil, err
	}

	fileID := rootPath + "/" + uuid.NewString() + ".png"

	w, err := standard.New(fileID)
	if err != nil {
		return nil, err
	}

	qrc.Save(w)

	bFile := toBase64(fileID)
	androidImage := toBase64(rootPath + "/assets/google-play-badge.jpg")
	iosImage := toBase64(rootPath + "/assets/Download_on_the_App_Store_Badge_FR_RGB_blk_100517.jpg")

	LogoImage := toBase64(rootPath + "/assets/swootte.jpg")

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
				_pdf.Text("Payez "+strings.ToTitle(enterpriseName)+" avec Tinda", props.Text{
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
		_pdf.Base64Image(LogoImage, consts.Jpg, props.Rect{
			Center: true,
		})
	})

	_pdf.RegisterFooter(func() {
		_pdf.Line(1)
		_pdf.Row(13, func() {
			_pdf.Col(10, func() {
				_pdf.Base64Image(androidImage, consts.Jpg, props.Rect{
					Percent: 57,
					Left:    50,
				})
			})

			_pdf.Col(7, func() {
				_pdf.Base64Image(iosImage, consts.Jpg, props.Rect{
					Percent: 39,
					Top:     1.2,
				})
			})
		})

		_pdf.Row(3, func() {
			_pdf.Col(12, func() {
				_pdf.Text("Swootte vous permet de créer votre solution de point de vente pour accepter les paiements en Franc CFA sur place ou sur le web. Commercez sans frontière même à l'international. nos frais de traitement de paiements sont de 1%, notre système est sécurisé à l'aide de chiffrements et vous offre une sécurité de paiement optimale.", props.Text{
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
