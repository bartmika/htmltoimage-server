package rpc

import (
	// "errors"

	"log"

	"github.com/bartmika/htmltoimage-server/pkg/dtos"
)

func (s *HTMLToImageService) Screenshot(websiteURL string) ([]byte, error) {
	// Create the request payload which we will send to the server.
	dto := &dtos.ScreenshotRequestDTO{
		WebsiteURL:   websiteURL,
		ImageType:    "png",
		ImageQuality: 90,
	}

	// Create the response payload that will be filled out by the server.
	var reply dtos.ScreenshotResponseDTO

	// Make the remote procedure call and handle the result.
	err := s.call("RPC.Screenshot", dto, &reply)
	if err != nil {
		log.Println("rpc_client | RPC.Screenshot | err", err)
		return nil, err
	}
	return reply.Content, nil
}
