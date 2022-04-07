module github.com/sonr-io/core

go 1.16

require (
	git.mills.io/prologic/bitcask v1.0.0
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/duo-labs/webauthn v0.0.0-20220330035159-03696f3d4499
	github.com/duo-labs/webauthn.io v0.0.0-20200929144140-c031a3e0f95d
	github.com/fxamacker/cbor/v2 v2.4.0
	github.com/gabriel-vasile/mimetype v1.3.1
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang/protobuf v1.5.2
	github.com/google/open-location-code/go v0.0.0-20210504205230-1796878d947c
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/kataras/golog v0.1.7
	github.com/lestrrat-go/jwx v1.2.20
	github.com/libp2p/go-libp2p v0.15.1
	github.com/libp2p/go-libp2p-connmgr v0.2.4
	github.com/libp2p/go-libp2p-core v0.9.0
	github.com/libp2p/go-libp2p-discovery v0.5.1
	github.com/libp2p/go-libp2p-kad-dht v0.13.1
	github.com/libp2p/go-libp2p-pubsub v0.5.4
	github.com/libp2p/go-msgio v0.0.6
	github.com/multiformats/go-multiaddr v0.4.0
	github.com/ockam-network/did v0.1.4-0.20210103172416-02ae01ce06d8
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.12.0
	github.com/shengdoushi/base58 v1.0.0
	github.com/sonr-io/blockchain v0.0.7
	github.com/spf13/viper v1.10.1
	github.com/stretchr/testify v1.7.1
	github.com/tendermint/starport v0.19.5
	go.buf.build/grpc/go/sonr-io/core v1.3.9
	go.buf.build/grpc/go/sonr-io/sonr v1.3.17
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
