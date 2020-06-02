package pdf

// import (
// 	"fmt"
// 	"log"
// 	"strconv"

// 	"github.com/jung-kurt/gofpdf"
// 	"github.com/mmanjoura/pppr/report"
// )

// // GeneratePdf ...
// func GeneratePdf(pyts []report.Payment, hdr []string) {

// 	data := make([][]string, len(pyts))
// 	for i, pyt := range pyts {
// 		row := make([]string, 0)
// 		row = append(row, pyt.TransactionID)
// 		row = append(row, pyt.MerchantID)
// 		row = append(row, pyt.TerminalID)
// 		row = append(row, pyt.CardNumberMasked[10:])
// 		row = append(row, strconv.FormatFloat(pyt.OriginalTransactionAmount, 'f', 6, 64))
// 		row = append(row, pyt.LocalCurrency)
// 		row = append(row, pyt.TransactionDate)
// 		row = append(row, strconv.FormatFloat(pyt.AcquirerFee, 'f', 6, 64))
// 		row = append(row, pyt.MarginRate)
// 		row = append(row, pyt.IsCardPresent)

// 		data[i] = row

// 	}

// 	pdf := table(data, hdr)

// 	if pdf.Err() {
// 		log.Fatalf("Failed creating PDF report: %s\n", pdf.Error())
// 	}
// 	err := savePDF(pdf)
// 	if err != nil {
// 		log.Fatalf("Cannot save PDF: %s|n", err)
// 	}

// }

// // func loadCSV(path string) [][]string {
// // 	f, err := os.Open(path)
// // 	if err != nil {
// // 		log.Fatalf("Cannot open '%s': %s\n", path, err.Error())
// // 	}
// // 	defer f.Close()
// // 	r := csv.NewReader(f)
// // 	rows, err := r.ReadAll()
// // 	if err != nil {
// // 		log.Fatalln("Cannot read CSV data:", err.Error())
// // 	}
// // 	return rows
// // }

// // func path() string {
// // 	if len(os.Args) < 2 {
// // 		return "ordersReport.csv"
// // 	}
// // 	return os.Args[1]
// // }

// // func newReport() *gofpdf.Fpdf {

// // 	pdf := gofpdf.New("P", "mm", "A4", "")

// // 	// We start by adding a new page to the document.
// // 	pdf.AddPage()

// // 	// Now we set the font to "Times", the style to "bold", and the size to 20 points.
// // 	pdf.SetFont("Arial", "B", 6)

// // 	// Then we write a text cell of length 40 and height 10. There are no
// // 	// starting coordinates used here; instead, the `Cell()` method moves
// // 	// the current position to the end of the cell so that the next call
// // 	// to `Cell()` continues after the previous cell.
// // 	pdf.Cell(40, 10, "Daily Report")

// // 	// The `Ln()` function moves the current position to a new line, with
// // 	// an optional line height parameter.
// // 	pdf.Ln(12)

// // 	pdf.SetFont("Arial", "", 18)
// // 	pdf.Cell(40, 10, time.Now().Format("Mon Jan 2, 2006"))
// // 	pdf.Ln(20)

// // 	return pdf
// // }

// /* ### How Cell() and Ln() advance the output position

// As mentioned in the comments, the `Cell()` method takes no coordinates.
// Instead, the PDF document maintains the current output position internally,
// and advances it to the right by the length of the cell being written.

// Method `Ln()` moves the output position back to the left border and down
// by the provided value. (Passing `-1` uses the height of the recently written cell.)

// */

// // Having created the initial document, we can now create the table header.
// // This time, we generate a formatted cell with a light grey as the
// // background color.
// func header(pdf *gofpdf.Fpdf, hdr []string) (*gofpdf.Fpdf, map[int]float64) {
// 	//pdf.SetFont("Arial", "B", 5)
// 	//pdf.SetFillColor(240, 240, 240)
// 	m := make(map[int]float64)
// 	for i, str := range hdr {
// 		// The `CellFormat()` method takes a couple of parameters to format
// 		// the cell. We make use of this to create a visible border around
// 		// the cell, and to enable the background fill.
// 		cellWidth := float64(len(str))
// 		//headerWidth = append(headerWidth, cellWidth)
// 		m[i] = cellWidth

// 		pdf.CellFormat(float64(240/12), 7, str, "0", 0, "", false, 0, "")
// 	}

// 	// Passing `-1` to `Ln()` uses the height of the last printed cell as
// 	// the line height.
// 	pdf.Ln(-1)
// 	return pdf, m
// }

// func table(tbl [][]string, hdr []string) *gofpdf.Fpdf {

// 	// Every column gets aligned according to its contents.
// 	align := []string{"L", "L", "L", "L", "L", "L", "L", "L", "L", "L"}

// 	//template, pdf := CreateTemplate(hdr)
// 	template, pdf := CreateTemplate(hdr)

// 	pdf.AddPage()
// 	pdf.UseTemplate(template)
// 	pdf.Ln(44)
// 	for j, line := range tbl {
// 		fmt.Println("J" + strconv.Itoa(j))

// 		if j > 0 && j%25 == 0 {
// 			pdf.AddPage()
// 			pdf.UseTemplate(template)
// 			pdf.Ln(44)
// 		}

// 		if j == 200 {
// 			break
// 		}

// 		for i, str := range line {
// 			fmt.Println("i" + strconv.Itoa(i))
// 			// Again, we need the `CellFormat()` method to create a visible
// 			// border around the cell. We also use the `alignStr` parameter
// 			// here to print the cell content either left-aligned or
// 			// right-aligned.

// 			// Noice the widht of each line "240/12"
// 			pdf.CellFormat(float64(240/12), 7, str, "0", 0, align[i], false, 0, "")

// 		}

// 		pdf.Ln(-1)

// 	}
// 	return pdf
// }

// // Next, let's not forget to impress our boss by adding a fancy image.
// func image(pdf *gofpdf.Fpdf) *gofpdf.Fpdf {
// 	// The `ImageOptions` method takes a file path, x, y, width, and height
// 	// parameters, and an `ImageOptions` struct to specify a couple of options.
// 	pdf.ImageOptions("stats.png", 225, 10, 25, 25, false, gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}, 0, "")
// 	return pdf
// }

// func savePDF(pdf *gofpdf.Fpdf) error {
// 	return pdf.OutputFileAndClose("report.pdf")
// }

// // CreateTemplate ...
// func CreateTemplate(hd []string) (gofpdf.Template, *gofpdf.Fpdf) {
// 	pdf := gofpdf.New("P", "mm", "A4", "")
// 	pdf.SetCompression(false)
// 	pdf.SetFont("Arial", "", 5)
// 	pdf.AliasNbPages("")

// 	pdf.SetFooterFunc(func() {
// 		// pdf.SetY(-15)
// 		// pdf.SetX(165)
// 		pdf.Image("visamastercard.png", 6, 280, 20, 5, false, "", 0, "")
// 		//pdf.SetFont("Arial", "I", 8)

// 		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d/{nb}", pdf.PageNo()),
// 			"", 0, "C", false, 0, "")

// 	})

// 	template := pdf.CreateTemplate(func(tpl *gofpdf.Tpl) {
// 		tpl.Image("planetlogo.png", 6, 6, 10, 0, false, "", 0, "")

// 		// Planet Address
// 		tpl.SetY(10)
// 		tpl.Cell(10, 0, "")
// 		tpl.CellFormat(30, 10, "Planet Address", "1", 0, "C", false, 0, "")

// 		// Invoice Title
// 		tpl.SetY(10)
// 		tpl.Cell(80, 0, "")
// 		tpl.CellFormat(30, 10, "Invoice Title", "1", 0, "C", false, 0, "")

// 		// customer Address
// 		tpl.SetY(10)
// 		tpl.Cell(150, 0, "")
// 		tpl.CellFormat(30, 10, "Planet Address", "1", 0, "C", false, 0, "")

// 		tpl.Ln(40)
// 		for _, str := range hd {
// 			// The `CellFormat()` method takes a couple of parameters to format
// 			// the cell. We make use of this to create a visible border around
// 			// the cell, and to enable the background fill.
// 			//cellWidth := float64(len(str))
// 			//headerWidth = append(headerWidth, cellWidth)
// 			//m[i] = cellWidth

// 			// Noice the widht of each line "240/12" this need to match table
// 			tpl.CellFormat(float64(240/12), 7, str, "0", 0, "", false, 0, "")

// 		}

// 		// Passing `-1` to `Ln()` uses the height of the last printed cell as
// 		// the line height.
// 		//pdf.Ln(-1)
// 		//tpl.Ln(-1)

// 	})

// 	return template, pdf

// }
