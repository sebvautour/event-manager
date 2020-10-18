package config

const ConfigTypeScript = "script"

type Script struct {
	Config
	Metadata ScriptMetadata `json:"metadata"`
	Script   string         `json:"script"`
}

type ScriptMetadata struct {
	Description         string `json:"description"`
	Type                string `json:"type"`
	InitialAuthor       string `json:"initial_author"`
	PerMessageTimeoutMs int    `json:"per_message_timeout_ms"`
	RateLimitPerMinute  int    `json:"rate_limit_per_minute"`
}

const ConfigTypeScriptList = "script_list"

type ScriptList struct {
	Config
	Scripts []string `json:"scripts"`
}
