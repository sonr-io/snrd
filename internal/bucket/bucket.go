package bucket

import (
	"errors"
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/pkg/client"
	mt "github.com/sonr-io/sonr/third_party/types/motor"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

var (
	errContentNotFound = func(id string) error {
		if id != "" {
			return fmt.Errorf("could not find content with id: %s", id)
		}

		return errors.New("could not find content")
	}
)

type bucketImpl struct {
	address      string
	whereIs      *bt.WhereIs
	contentCache map[string]*mt.BucketContent
	bucketCache  map[string]Bucket
	shell        *shell.Shell
	rpcClient    *client.Client
}

func New(
	address string,
	whereIs *bt.WhereIs,
	shell *shell.Shell,
	rpcClient *client.Client) *bucketImpl {

	return &bucketImpl{
		address:      address,
		whereIs:      whereIs,
		shell:        shell,
		contentCache: make(map[string]*mt.BucketContent),
		rpcClient:    rpcClient,
	}
}
