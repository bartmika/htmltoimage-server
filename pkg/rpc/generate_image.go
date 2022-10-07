package rpc

import (
	// "errors"

	"log"

	"github.com/bartmika/htmltoimage-server/pkg/dtos"
)

func (s *HTMLToImageService) GenerateImage(dto *dtos.HTMLToImageRequestDTO) (*dtos.HTMLToImageResponseDTO, error) {
	var reply dtos.HTMLToImageResponseDTO
	err := s.call("RPC.GenerateImage", dto, &reply)
	if err != nil {
		log.Println("rpc_client | RPC.GenerateImage | err", err)
		return nil, err
	}
	return &reply, nil
}
