package config

type QuarkLTConfig struct {
	Name       string        `json:"name"`
	ServerHost string        `json:"server-host",yaml:"server-host"`
	SiteSetup  SiteSetupConf `json:"site-setup",yaml:"site-setup"`
}

type Helpers struct {
	SshAgent *SshAgentConf `json:"ssh-agent",yaml:"ssh-agent"`
}

//-----------------------------------------------------------------------------------------------------

type SshAgentConf struct {
	Host string `json:"host"`
	User string `json:"user"`
	Port int    `json:"port"`

	AuthMethod *AuthMethod `json:"auth-method",yaml:"auth-method"`
}

func ValidateSshAgentConf(agentConf *SshAgentConf) bool {
	if agentConf == nil {
		return false
	}
	if len(agentConf.Host) == 0 {
		return false
	}
	if len(agentConf.User) == 0 {
		return false
	}
	if agentConf.Port == 0 {
		return false
	}
	if agentConf.AuthMethod == nil {
		return false
	}
	return true
}

//-----------------------------------------------------------------------------------------------------

type AuthMethod struct {
	UserAuth *UserAuth `json:"user-auth",yaml:"user-auth"`
	KeyAuth  *KeyAuth  `json:"key-auth",yaml:"key-auth"`
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
