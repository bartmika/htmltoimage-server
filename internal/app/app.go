package app

import (
	"context"

	"github.com/bartmika/htmltoimage-server/internal/config"
	"github.com/bartmika/htmltoimage-server/pkg/dtos"
	"github.com/bartmika/htmltoimage-server/pkg/uuid"
	"github.com/chromedp/chromedp"
)

type HTMLToImageApp interface {
	GenerateImage(dto *dtos.HTMLToImageRequestDTO) (*dtos.HTMLToImageResponseDTO, error)
}

type htmlToImageApp struct {
	AppConfig    *config.Conf
	UUIDProvider uuid.Provider
}

func New(appConfig *config.Conf, uuidp uuid.Provider) (HTMLToImageApp, error) {
	return &htmlToImageApp{
		AppConfig:    appConfig,
		UUIDProvider: uuidp,
	}, nil
}

func (app *htmlToImageApp) GenerateImage(dto *dtos.HTMLToImageRequestDTO) (*dtos.HTMLToImageResponseDTO, error) {
	// create allocator context for use with creating a browser context later
	allocatorContext, allocatorCancel := chromedp.NewRemoteAllocator(context.Background(), app.AppConfig.ChromeHeadless.Address)

	// create context for the chromedp.
	ctx, cancel := chromedp.NewContext(allocatorContext)

	defer allocatorCancel()
	defer cancel()

	var buf []byte

	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, fullScreenshot(dto.WebsiteURL, 90, &buf)); err != nil {
		return nil, err
	}

	return &dtos.HTMLToImageResponseDTO{
		FileName: app.UUIDProvider.NewUUID(),
		Content:  buf,
	}, nil
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Note: chromedp.FullScreenshot overrides the device's emulation settings. Use
// device.Reset to reset the emulation and viewport settings.
// (Source: https://github.com/chromedp/examples/blob/master/remote/main.go)
func fullScreenshot(websiteURL string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(websiteURL),
		chromedp.FullScreenshot(res, quality),
	}
}
