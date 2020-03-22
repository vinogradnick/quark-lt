package reports

import "src/github.com/jung-kurt/gofpdf"

type ReportGenerator struct {
}

func (rg *ReportGenerator) CreateDocument() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
}
