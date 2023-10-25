package lib

// unless we need more configuration later, we can just embed this into cmd/version
var (
	Version = "1.0.8"
)

type Config struct {
	APIKey      string   `json:"-"` //mandatory
	Config      string   `json:"-"`
	Endpoint    string   `json:"-"` //mandatory
	Environment string   `json:"environment,omitempty"`
	Event       string   `json:"event"` //mandatory
	Type        string   `json:"type,omitempty"`
	Group       string   `json:"group,omitempty"`
	Origin      string   `json:"origin,omitempty"`
	RawData     string   `json:"rawData,omitempty"`
	Resource    string   `json:"resource,"` //mandatory
	Severity    string   `json:"severity,omitempty"`
	Service     []string `json:"service,omitempty"`
	Tag         []string `json:"tag,omitempty"`
	//Attributes  []string `json:"attributes,omitempty"`  // server wants this to be a JSON object
	Text    string `json:"text,omitempty"`
	Timeout int    `json:"timeout,omitempty"`
	Value   int    `json:"value,omitempty"`
}
