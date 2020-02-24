package config

import (
	"github.com/quark_lt/cmd/quark_worker"
	"go.uber.org/zap"
)
//-----------------------------------------------------------------------------------------------------

type StepConf struct {
	Start    int32  `yaml:"start"`
	End      int32  `yaml:"end"`
	Duration string `yaml:"duration"`
	Step     int32  `yaml:"step"`
}

func (s StepConf) validate() bool {
	var errorValue int
	if PositiveValidate(s.Start) != nil {
		zap.L().Error("[StepConfig]  Start value has Error value")
		errorValue++
	}
	if PositiveValidate(s.End) != nil {
		zap.L().Error("[StepConfig]  End value has Error value")
		errorValue++

	}
	if PositiveValidate(int32(quark_worker.DurationConvertation(s.Duration))) != nil {
		zap.L().Error("[StepConfig] Duration value has Error value")
		errorValue++

	}
	if PositiveValidate(s.Step) != nil {
		zap.L().Error("[StepConfig] Step value has Error value")
		errorValue++

	}
	return errorValue==0

}
//-----------------------------------------------------------------------------------------------------

type ConstConf struct {
	Value    int32  `yaml:"value"`
	Duration string `yaml:"duration"`
}

func (c ConstConf) validate() bool{
	var errorValue int
	if PositiveValidate(c.Value) != nil {
		zap.L().Error("[ConstConfig]  Value  has Error value")
		errorValue++
	}
	if PositiveValidate(int32(quark_worker.DurationConvertation(c.Duration))) != nil {
		zap.L().Error("[ConstConfig]  Duration value has Error value")
		errorValue++

	}
	return errorValue==0
}
//-----------------------------------------------------------------------------------------------------
type LinearConf struct {
	Start    int32  `yaml:"start"`
	End      int32  `yaml:"end"`
	Duration string `yaml:"duration"`
}

func (l LinearConf) validate() bool {
	var errorValue int
	if PositiveValidate(l.Start) != nil {
		zap.L().Error("[LinearConf]  Start value has Error value")
		errorValue++
	}
	if PositiveValidate(l.End) != nil {
		zap.L().Error("[LinearConf]  End value has Error value")
		errorValue++

	}
	if PositiveValidate(int32(quark_worker.DurationConvertation(l.Duration))) != nil {
		zap.L().Error("[LinearConf] Duration value has Error value")
		errorValue++

	}

	return errorValue==0
}
//-----------------------------------------------------------------------------------------------------

type ExpConf struct {
	Value    int32  `yaml:"value"`
	Duration string `yaml:"duration"`
}
func (e ExpConf) validate() bool {
	var errorValue int
	if PositiveValidate(e.Value) != nil {
			zap.L().Error("[ExpConfig]  Value  has Error value")
		errorValue++
	}
	if PositiveValidate(int32(quark_worker.DurationConvertation(e.Duration))) != nil {
		zap.L().Error("[ExpConfig]  Duration value has Error value")
		errorValue++

	}
	return errorValue==0
}

//-----------------------------------------------------------------------------------------------------

type Helpers struct {
	SshAgent *SshAgentConf `yaml:"ssh-agent"`
}



//-----------------------------------------------------------------------------------------------------

type SshAgentConf struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Port string `yaml:"port"`

	AuthMethod *AuthMethod `yaml:"auth-method"`
}




//-----------------------------------------------------------------------------------------------------

type AuthMethod struct {
	UserAuth *UserAuth `yaml:"user-auth"`
	KeyAuth  *KeyAuth  `yaml:"key-auth"`
}



//-----------------------------------------------------------------------------------------------------

type UserAuth struct {
	Password string `yaml:"password"`
}


//-----------------------------------------------------------------------------------------------------

type KeyAuth struct {
	Path string `yaml:"path"`
}


//-----------------------------------------------------------------------------------------------------



//-----------------------------------------------------------------------------------------------------
