package pdf_test

import (
	"server/pdf"
	"testing"
)

func TestPdf(t *testing.T) {
	t.Run("create pdf", func(t *testing.T) {
		pdf.CreatePdfFile("eth, net, web3, admin")
	})
}
