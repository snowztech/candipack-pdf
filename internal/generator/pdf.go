package generator

import (
	"context"
	"fmt"
	"os"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
)

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) GeneratePDF(htmlFile string) ([]byte, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ExecPath(os.Getenv("CHROME_PATH")),
		chromedp.Headless,
		chromedp.NoSandbox,
		chromedp.DisableGPU,
	)

	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer allocCancel()
	chromeCtx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var pdf []byte
	err := chromedp.Run(chromeCtx,
		chromedp.Navigate(fmt.Sprintf("file://%s", htmlFile)),
		chromedp.WaitReady("body"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			data, _, err := page.PrintToPDF().
				WithMarginTop(0).
				WithMarginBottom(0).
				WithMarginLeft(0).
				WithMarginRight(0).
				WithPrintBackground(true).
				WithPaperWidth(8.27).
				WithPaperHeight(11.69).
				Do(ctx)
			if err != nil {
				return err
			}
			pdf = data
			return nil
		}),
	)

	if err != nil {
		return nil, errors.Wrap(err, "pdf generation failed")
	}

	return pdf, nil
}
