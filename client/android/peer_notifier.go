package android

// PeerInfo describe information about the peers. It designed for the UI usage
type PeerInfo struct {
	IP                         string
	FQDN                       string
	ConnStatus                 string // Todo replace to enum
	PubKey                     string
	LocalIceCandidateType      string
	RemoteIceCandidateType     string
	LocalIceCandidateEndpoint  string
	RemoteIceCandidateEndpoint string
	BytesRx                    int64
	BytesTx                    int64
	Latency                    int64
	Relayed                    bool
	Direct                     bool
	ConnStatusUpdate           string
	LastWireguardHandshake     string
	RosenpassEnabled           bool
}

// PeerInfoArray is a wrapper of []PeerInfo
type PeerInfoArray struct {
	items []PeerInfo
}

// Add new PeerInfo to the collection
func (array *PeerInfoArray) Add(s PeerInfo) *PeerInfoArray {
	array.items = append(array.items, s)
	return array
}

// Get return an element of the collection
func (array *PeerInfoArray) Get(i int) *PeerInfo {
	return &array.items[i]
}

// Size return with the size of the collection
func (array *PeerInfoArray) Size() int {
	return len(array.items)
}
