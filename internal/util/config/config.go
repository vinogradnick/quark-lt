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
	SshAgent SshAgentConf `yaml:"ssh-agent"`
}
type SshAgentConf struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
type AuthMethod struct {
	UserAuth UserAuth `yaml:"user-auth"`
	KeyAuth  KeyAuth  `yaml:"key-auth"`
}
type UserAuth struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
type KeyAuth struct {
	Path string `yaml:"path"`
}
