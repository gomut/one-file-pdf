// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-22 03:10:56 468B87                           [utest/draw_text.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// DrawText is the unit test for
// DrawText(s string) *PDF
func DrawText(t *testing.T) {
	fmt.Println("utest.DrawText")
	//
	func() {
		var ob = pdf.NewPDF("A4")
		ob.SetCompression(false).
			SetUnits("cm").
			SetColumnWidths(1, 4, 9).
			SetColor("#006B3C CadmiumGreen").
			SetFont("Helvetica-Bold", 10).
			SetX(5).
			SetY(5).
			DrawText("FIRST").
			DrawText("SECOND").
			DrawText("THIRD")
		//
		var expect = `
		%PDF-1.4
		1 0 obj<</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj<</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj<</Type/Page/Parent 2 0 R/Contents 4 0 R\
		/Resources<</Font <</FNT1 5 0 R>>>>>>
		endobj
		4 0 obj <</Length 143>>stream
		BT /FNT1 10 Tf ET
		 0.000 0.420 0.235 rg
		0.000 0.420 0.235 RG
		BT 0 700 Td (FIRST) Tj ET
		BT 28 700 Td (SECOND) Tj ET
		BT 141 700 Td (THIRD) Tj ET
		endstream
		5 0 obj<</Type/Font/Subtype/Type1/Name/F1/BaseFont/Helvetica-Bold\
		/Encoding/WinAnsiEncoding>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000009 00000 n
		0000000053 00000 n
		0000000125 00000 n
		0000000217 00000 n
		0000000401 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		501
		%%EOF
		`
		pdfCompare(t, ob.Bytes(), expect, pdfStreamsInText)
	}()

	func() {
		var ob = pdf.NewPDF("A4")
		ob.SetCompression(false).
			SetUnits("cm").
			SetFont("Ye-Olde-Scriptte", 10).
			SetXY(5, 5).
			SetHorizontalScaling(150).
			DrawText("Ye-Olde-Scriptte")
		//
		var expect = `
		%PDF-1.4
		1 0 obj<</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj<</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj<</Type/Page/Parent 2 0 R/Contents 4 0 R\
		/Resources<</Font <</FNT1 5 0 R>>>>>>
		endobj
		4 0 obj <</Length 113>>stream
		BT /FNT1 10 Tf ET
		BT 150 Tz ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 141 700 Td (Ye-Olde-Scriptte) Tj ET
		endstream
		5 0 obj<</Type/Font/Subtype/Type1/Name/F1/BaseFont/Helvetica\
		/Encoding/WinAnsiEncoding>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000009 00000 n
		0000000053 00000 n
		0000000125 00000 n
		0000000217 00000 n
		0000000371 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		466
		%%EOF
		`
		pdfCompare(t, ob.Bytes(), expect, pdfStreamsInText)
		TEqual(t, len(ob.Errors()), 1)
		TEqual(t, ob.PullError(),
			fmt.Errorf(`Invalid font name: "Ye-Olde-Scriptte" @DrawText`))
	}()

} //                                                                    DrawText

//end
