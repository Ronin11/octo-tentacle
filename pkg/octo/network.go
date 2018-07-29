package octo

type NetworkType int

const (
	NATSNetwork NetworkType = 1 + iota
	MQTTNetwork
)

//Network ...
type Network struct {
	networkType NetworkType
	server string
}

// JoinNetwork ...
func JoinNetwork(serverAddress string, networkType NetworkType) *Network{
	return &Network{
		networkType: networkType,
		server: serverAddress,
	}
}

//GetServerAddress ...
func (n Network) GetServerAddress() string{
	return n.server
}

//GetNetworkType ...
func (n Network) GetNetworkType() NetworkType{
	return n.networkType
}

func (n Network) AddService(service Service) error{
	// CreateService(&n)
	return nil
}