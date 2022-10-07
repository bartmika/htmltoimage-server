package app

import (
	"context"
	"errors"

	"github.com/bartmika/htmltoimage-server/internal/config"
	"github.com/bartmika/htmltoimage-server/pkg/dtos"
	"github.com/bartmika/htmltoimage-server/pkg/uuid"
	"github.com/chromedp/chromedp"
)

type HTMLToImageApp interface {
	Screenshot(dto *dtos.ScreenshotRequestDTO) (*dtos.ScreenshotResponseDTO, error)
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

func (app *htmlToImageApp) Screenshot(dto *dtos.ScreenshotRequestDTO) (*dtos.ScreenshotResponseDTO, error) {
	// Defensive Code:
	if dto.ImageType != "png" {
		return nil, errors.New("only accepted image type is currently the value `png`")
	}
	if dto.ImageQuality < 0 {
		return nil, errors.New("image quality cannot be less then `0`")
	}
	if dto.ImageQuality > 100 {
		return nil, errors.New("image quality cannot be greater then `100`")
	}

	// create allocator context for use with creating a browser context later
	allocatorContext, allocatorCancel := chromedp.NewRemoteAllocator(context.Background(), app.AppConfig.ChromeHeadless.Address)

	// create context for the chromedp.
	ctx, cancel := chromedp.NewContext(allocatorContext)

	defer allocatorCancel()
	defer cancel()

	var buf []byte

	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, fullScreenshot(dto.WebsiteURL, dto.ImageQuality, &buf)); err != nil {
		return nil, err
	}

	return &dtos.ScreenshotResponseDTO{
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
