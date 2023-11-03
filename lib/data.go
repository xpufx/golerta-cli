package lib

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
	Tags        []string `json:"tags,omitempty"` //alert wants 'tag' while heartbeat wants 'tags'
	//Attributes  []string `json:"attributes,omitempty"`  // server wants this to be a JSON object
	Text    string `json:"text,omitempty"`
	Timeout int    `json:"timeout,omitempty"`
	Value   int    `json:"value,omitempty"`
}

type Heartbeat struct {
	APIKey   string `json:"-"` //mandatory
	Config   string `json:"-"`
	Endpoint string `json:"-"` //mandatory
	Origin   string `json:"origin,omitempty"`
	Text     string `json:"text,omitempty"`
	Timeout  int    `json:"timeout,omitempty"`
}

type Golerta struct {
	Config
	ConfigFile string //  Config.Config should be moved here
	Debug      bool
	Curl       bool
	Dryrun     bool
}
