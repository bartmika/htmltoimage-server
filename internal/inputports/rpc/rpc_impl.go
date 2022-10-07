package rpc

import (
	"github.com/bartmika/htmltoimage-server/pkg/dtos"
)

func (rpc *RPC) Screenshot(req *dtos.ScreenshotRequestDTO, res *dtos.ScreenshotResponseDTO) error {

	pdf, err := rpc.App.Screenshot(req)
	if err != nil {
		return err
	}

	*res = *pdf
	return nil
}
