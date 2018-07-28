package octo

//Network ...
type Network struct {
	server string
}

// CreateNetwork ...
func CreateNetwork(serverAddress string) *Network{
	return &Network{server: serverAddress}
}

//GetServerAddress ...
func (n Network) GetServerAddress() string{
	return n.server
}