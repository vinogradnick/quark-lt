package config

type QuarkLTConfig struct {
	Name       string        `json:"name"`
	ServerHost string        `json:"server-host"`
	SiteSetup  SiteSetupConf `json:"site-setup"`
}

type Helpers struct {
	SshAgent *SshAgentConf `json:"ssh-agent"`
}

//-----------------------------------------------------------------------------------------------------

type SshAgentConf struct {
	Host string `json:"host"`
	User string `json:"user"`
	Port int    `json:"port"`

	AuthMethod *AuthMethod `json:"auth-method"`
}

//-----------------------------------------------------------------------------------------------------

type AuthMethod struct {
	UserAuth *UserAuth `json:"user-auth"`
	KeyAuth  *KeyAuth  `json:"key-auth"`
}

//-----------------------------------------------------------------------------------------------------

type UserAuth struct {
	Password string `json:"password"`
}

//-----------------------------------------------------------------------------------------------------

type KeyAuth struct {
	Path string `json:"path"`
}

//Worker config -----------------------------------------

const (
	IN_MEMORY_TYPE = 1
	EXTERNAL_TYPE  = 2
	MICROVM_TYPE   = 3
)

/// Node config

