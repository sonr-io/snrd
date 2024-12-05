package producer

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common"
)

type ProducerContext struct {
	echo.Context
	// TokenParser is the attentuations assigned to the producer service
	TokenParser common.UCANParser

	// IPFSClient is the IPFS client used to resolve the UCAN
	IPFSClient common.IPFSClient
}
