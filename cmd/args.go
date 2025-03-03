package cmd

type Args struct {
	Header            []string
	P                 []string
	IgnoreParams      []string
	Config            string
	Cookie            string
	Data              string
	CustomPayload     string
	CustomAlertValue  string
	CustomAlertType   string
	UserAgent         string
	Blind             string
	Output            string
	Format            string
	FoundAction       string
	FoundActionShell  string
	Proxy             string
	Grep              string
	IgnoreReturn      string
	MiningWord        string
	Method            string
	CookieFromRaw     string
	RemotePayloads    string
	RemoteWordlists   string
	OnlyPoC           string
	PoCType           string
	ReportFormat      string
	HarFilePath       string
	Timeout           int
	Delay             int
	Concurrence       int
	MaxCPU            int
	OnlyDiscovery     bool
	Silence           bool
	Mining            bool
	FindingDOM        bool
	FollowRedirect    bool
	NoColor           bool
	NoSpinner         bool
	UseBAV            bool
	SkipBAV           bool
	SkipMiningDom     bool
	SkipMiningDict    bool
	SkipMiningAll     bool
	SkipXSSScan       bool
	OnlyCustomPayload bool
	SkipGrep          bool
	Debug             bool
	SkipHeadless      bool
	UseDeepDXSS       bool
	OutputAll         bool
	WAFEvasion        bool
	ReportBool        bool
	OutputRequest     bool
	OutputResponse    bool
}
