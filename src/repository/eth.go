package repository

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/openvino/openvino-api/src/config"
	ens "github.com/wealdtech/go-ens/v3"
)

var Eth *ethclient.Client

func SetupETH(config config.EthereumConfig) (*ethclient.Client, error) {

	var InfuraHost string = "https://" + config.Network + ".infura.io/v3/" + config.InfuraSecretKey
	Eth, ethError := ethclient.Dial(InfuraHost)

	if ethError != nil {
		return nil, ethError
	}

	return Eth, nil

}

func GetDomain(address string) (string, error) {

	EthAddress := common.HexToAddress(address)

	domain, err := ens.ReverseResolve(Eth, EthAddress)
	if err != nil {
		return "", err
	} else {
		return domain, nil
	}

}
