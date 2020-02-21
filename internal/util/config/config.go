package config

type StepConf struct {
	Start    int32  `yaml:"start"`
	End      int32  `yaml:"end"`
	Duration string `yaml:"duration"`
	Step     int32  `yaml:"step"`
}
type ConstConf struct {
	Value    int32  `yaml:"value"`
	Duration string `yaml:"duration"`
}
type LinearConf struct {
	Start    int32  `yaml:"start"`
	End      int32  `yaml:"end"`
	Duration string `yaml:"duration"`
}
type ExpConf struct {
	Value    int32  `yaml:"value"`
	Duration string `yaml:"duration"`
}

type Helpers struct {
	SshAgent *SshAgentConf `yaml:"ssh-agent"`
}
type SshAgentConf struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Port string `yaml:"port"`

	AuthMethod *AuthMethod `yaml:"auth-method"`
}
type AuthMethod struct {
	UserAuth *UserAuth `yaml:"user-auth"`
	KeyAuth  *KeyAuth  `yaml:"key-auth"`
}
type UserAuth struct {
	Password string `yaml:"password"`
}
type KeyAuth struct {
	Path string `yaml:"path"`
}
