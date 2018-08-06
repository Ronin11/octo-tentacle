package octo

// NetworkType ...
type NetworkType int

const (
	NATSNetwork NetworkType = iota
	MQTTNetwork
)

// Network ...
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

// GetServerAddress ...
func (n Network) GetServerAddress() string{
	return n.server
}

// GetNetworkType ...
func (n Network) GetNetworkType() NetworkType{
	return n.networkType
}

// AddService ...
func (n Network) AddService(service Service) error{
	service.AddToNetwork(&n)
	return nil
}