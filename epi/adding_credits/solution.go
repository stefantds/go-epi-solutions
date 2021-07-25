package adding_credits

type Solution interface {
	Insert(clientID string, c int)
	Remove(clientID string) bool
	Lookup(clientID string) int
	AddAll(c int)
	Max() string
}

type ClientsCreditsInfo struct {
	offset          int
	clientToCredit  map[string]int
	creditToClients *BSTMap
}

func NewClientsCreditsInfo() Solution {
	return &ClientsCreditsInfo{
		clientToCredit: make(map[string]int),
	}
}

func NewClientsCreditsInfo() Solution {
	// TODO - Add your code here
	return &ClientsCreditsInfo{}
}

func (cc *ClientsCreditsInfo) Insert(clientID string, c int) {
	cc.Remove(clientID)
	cc.clientToCredit[clientID] = c - cc.offset

	if tm := cc.creditToClients.Find(c - cc.offset); tm != nil {
		tm.Value[clientID] = true
	} else {
		newMap := map[string]bool{
			clientID: true,
		}
		cc.creditToClients = cc.creditToClients.Insert(c-cc.offset, newMap)
	}
}

func (cc *ClientsCreditsInfo) Remove(clientID string) bool {
	credit, ok := cc.clientToCredit[clientID]
	if ok {
		clients := cc.creditToClients.Find(credit)
		if clients != nil && len(clients.Value) > 0 {
			delete(clients.Value, clientID)
			if len(clients.Value) == 0 {
				cc.creditToClients = cc.creditToClients.Delete(credit)
			}
		}
		delete(cc.clientToCredit, clientID)
		return true
	}

	return false
}

func (cc *ClientsCreditsInfo) Lookup(clientID string) int {
	c, ok := cc.clientToCredit[clientID]
	if !ok {
		return -1
	}
	return c + cc.offset
}

func (cc *ClientsCreditsInfo) AddAll(c int) {
	cc.offset += c
}

func (cc *ClientsCreditsInfo) Max() string {
	if cc.creditToClients == nil {
		return ""
	}
	return getAnyKey(cc.creditToClients.Max().Value)
}

// getAnyKey returns some key from a map (if at least one key exists)
func getAnyKey(m map[string]bool) string {
	for k := range m {
		return k
	}
	return ""
}
