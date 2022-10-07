package rpc

import (
	"github.com/bartmika/htmltoimage-server/pkg/dtos"
)

func (rpc *RPC) GenerateImage(req *dtos.HTMLToImageRequestDTO, res *dtos.HTMLToImageResponseDTO) error {

	pdf, err := rpc.App.GenerateImage(req)
	if err != nil {
		return err
	}

	*res = *pdf
	return nil
}
