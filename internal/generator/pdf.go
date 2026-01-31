package generator

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
)

const (
	userAgentOverride   = "WebScraper 1.0"
	htmlSelector        = "body"
	networkReadyTimeOut = 15 * time.Second
)

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) GeneratePDF(content string) ([]byte, error) {
	var pdf []byte
	url := getFilePathAsURL(content)

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	err := chromedp.Run(ctx, g.saveURLAsPDF(url, &pdf))
	if err != nil {
		return nil, errors.Wrap(err, "failed to run chromedp tasks")
	}

	return pdf, nil
}

func (g *Generator) saveURLAsPDF(url string, pdf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		emulation.SetUserAgentOverride(userAgentOverride),
		chromedp.Navigate(url),
		chromedp.WaitVisible(htmlSelector, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			if err := waitForNetworkIdle(ctx, networkReadyTimeOut); err != nil {
				slog.Warn(err.Error())
			}
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			data, _, err := page.
				PrintToPDF().
				WithMarginLeft(0).
				WithMarginTop(0.4).
				WithMarginRight(0).
				WithMarginBottom(0.4).
				WithPaperWidth(8.3).
				WithPaperHeight(11.7).
				WithPrintBackground(true).
				Do(ctx)
			if err != nil {
				return errors.Wrap(err, "failed to generate pdf")
			}
			*pdf = data
			return nil
		}),
	}
}

func waitForNetworkIdle(ctx context.Context, timeout time.Duration) error {
	idleChan := make(chan struct{})

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if event, ok := ev.(*page.EventLifecycleEvent); ok {
			if event.Name == "networkIdle" {
				close(idleChan)
			}
		}
	})

	select {
	case <-idleChan:
		// Network is idle
		return nil
	case <-time.After(timeout):
		return fmt.Errorf("timeout %.0f seconds waiting for network idle", timeout.Seconds())
	}
}


func getFilePathAsURL(filename string) string {
	path, err := os.Getwd()
	if err != nil {
		slog.Error("Failed to get working directory", "error", err)
	}

	return fmt.Sprintf("file://%s/%s", path, filename)
}