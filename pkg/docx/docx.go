package docx

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/carmel/gooxml/color"
	"github.com/carmel/gooxml/common"
	"github.com/carmel/gooxml/document"
	"github.com/carmel/gooxml/measurement"
	"github.com/carmel/gooxml/schema/soo/wml"
)

func Test() {
	doc := document.New()

	para := doc.AddParagraph()
	para.SetStyle("Title")
	para.Properties().SetAlignment(wml.ST_JcCenter)

	run := para.AddRun()
	run.AddText("Simple Document Formatting")

	para = doc.AddParagraph()
	para.SetStyle("Heading1")
	run = para.AddRun()
	run.AddText("Some Heading Text")

	para = doc.AddParagraph()
	para.SetStyle("Heading2")
	run = para.AddRun()
	run.AddText("Some Heading Text")

	para = doc.AddParagraph()
	para.SetStyle("Heading3")
	run = para.AddRun()
	run.AddText("Some Heading Text")

	para = doc.AddParagraph()
	para.Properties().SetFirstLineIndent(0.5 * measurement.Inch)

	run = para.AddRun()
	run.AddText("A run is a string of characters with the same formatting. ")

	run = para.AddRun()
	run.Properties().SetBold(true)
	run.Properties().SetFontFamily("Courier")
	run.Properties().SetSize(15)
	run.Properties().SetColor(color.Red)
	run.AddText("Multiple runs with different formatting can exist in the same paragraph. ")

	run = para.AddRun()
	run.AddText("Adding breaks to a run will insert line breaks after the run. ")
	run.AddBreak()
	run.AddBreak()

	createParaRun(doc, "Runs support styling options:")

	run = createParaRun(doc, "small caps")
	run.Properties().SetSmallCaps(true)

	run = createParaRun(doc, "strike")
	run.Properties().SetStrikeThrough(true)

	run = createParaRun(doc, "double strike")
	run.Properties().SetDoubleStrikeThrough(true)

	run = createParaRun(doc, "outline")
	run.Properties().SetOutline(true)

	run = createParaRun(doc, "emboss")
	run.Properties().SetEmboss(true)

	run = createParaRun(doc, "shadow")
	run.Properties().SetShadow(true)

	run = createParaRun(doc, "imprint")
	run.Properties().SetImprint(true)

	run = createParaRun(doc, "highlighting")
	run.Properties().SetHighlight(wml.ST_HighlightColorYellow)

	run = createParaRun(doc, "underline")
	run.Properties().SetUnderline(wml.ST_UnderlineWavyDouble, color.Red)

	run = createParaRun(doc, "text effects")
	run.Properties().SetEffect(wml.ST_TextEffectAntsRed)

	nd := doc.Numbering.Definitions()[0]

	for i := 1; i < 5; i++ {
		p := doc.AddParagraph()
		p.SetNumberingLevel(i - 1)
		p.SetNumberingDefinition(nd)
		run := p.AddRun()
		run.AddText(fmt.Sprintf("Level %d", i))
	}

	var mapList = []map[string][]string{
		map[string][]string{"Preferred programming language": []string{"Go", "Java", "PHP"}},
		map[string][]string{"Which sport you love to play": []string{"Football", "Basketball", "Diving"}},
		map[string][]string{"Another information": []string{"This A", "This B", "This C"}},
	}

	AddNumbering(doc)
	newDef := doc.Numbering.Definitions()[1]
	for i := 0; i < len(mapList); i++ {
		for key, val := range mapList[i] {
			para := doc.AddParagraph()
			para.SetNumberingDefinition(newDef)
			para.SetNumberingLevel(0)
			para.AddRun().AddText(key)

			for i := 0; i < len(val); i++ {
				para := doc.AddParagraph()
				para.SetNumberingDefinition(newDef)
				para.SetNumberingLevel(1)
				para.AddRun().AddText(val[i])
			}
		}
	}

	// Adding an Image
	img2data, err := ioutil.ReadFile("test.png")
	if err != nil {
		log.Fatalf("unable to read file: %s", err)
	}
	img2, err := common.ImageFromBytes(img2data)
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}

	img2ref, err := doc.AddImage(img2)
	if err != nil {
		log.Fatalf("unable to add image to document: %s", err)
	}
	para = doc.AddParagraph()
	para.Properties().SetAlignment(wml.ST_JcCenter)
	// anchored, err := para.AddRun().AddDrawingAnchored(img2ref)
	// if err != nil {
	// 	log.Fatalf("unable to add anchored image: %s", err)
	// }

	run = para.AddRun()
	inl, err := run.AddDrawingInline(img2ref)
	if err != nil {
		log.Fatalf("unable to add inline image: %s", err)
	}
	inl.SetSize(1*measurement.Inch, 1*measurement.Inch)

	doc.SaveToFile("simple.docx")
}

func createParaRun(doc *document.Document, s string) document.Run {
	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText(s)
	return run
}

func AddNumbering(doc *document.Document) {
	// Create numbering definition.
	newDef := doc.Numbering.AddDefinition()

	// Add level to number definition with decimal format.
	lvl := newDef.AddLevel()
	lvl.SetFormat(wml.ST_NumberFormatDecimal)
	lvl.SetAlignment(wml.ST_JcLeft)
	lvl.Properties().SetLeftIndent(0.5 * measurement.Distance(1) * measurement.Inch)

	// Sets the numbering level format.
	lvl.SetText("%1.")

	slvl := newDef.AddLevel()
	slvl.SetFormat(wml.ST_NumberFormatDecimal)
	slvl.SetAlignment(wml.ST_JcLeft)
	slvl.Properties().SetLeftIndent(0.5 * measurement.Distance(2) * measurement.Inch)
	slvl.SetText("%1.%2")
}
