package config

// +ioc:autowire=true
// +ioc:autowire:baseType=true
// +ioc:autowire:type=config
// +ioc:autowire:paramType=ConfigFloat64
// +ioc:autowire:constructFunc=new

type ConfigFloat64 float64

func (ci *ConfigFloat64) Value() float64 {
	return float64(*ci)
}

func (ci *ConfigFloat64) new(impl *ConfigFloat64) (*ConfigFloat64, error) {
	*impl = *ci
	return impl, nil
}

func FromFloat64(val float64) *ConfigFloat64 {
	configFloat64 := ConfigFloat64(val)
	return &configFloat64
}

// +ioc:autowire=true
// +ioc:autowire:baseType=true
// +ioc:autowire:type=config
// +ioc:autowire:paramType=ConfigInt64
// +ioc:autowire:constructFunc=new

type ConfigInt64 int64

func (ci *ConfigInt64) Value() int64 {
	return int64(*ci)
}

func (ci *ConfigInt64) new(impl *ConfigInt64) (*ConfigInt64, error) {
	*impl = *ci
	return impl, nil
}

func FromInt64(val int64) *ConfigInt64 {
	configInt64 := ConfigInt64(val)
	return &configInt64
}

// +ioc:autowire=true
// +ioc:autowire:baseType=true
// +ioc:autowire:type=config
// +ioc:autowire:paramType=ConfigInt
// +ioc:autowire:constructFunc=new

type ConfigInt int

func (ci *ConfigInt) Value() int {
	return int(*ci)
}

func (param *ConfigInt) new(impl *ConfigInt) (*ConfigInt, error) {
	*impl = *param
	return impl, nil
}

func FromInt(val int) *ConfigInt {
	configInt := ConfigInt(val)
	return &configInt
}

// +ioc:autowire=true
// +ioc:autowire:baseType=true
// +ioc:autowire:type=config
// +ioc:autowire:paramType=ConfigMap
// +ioc:autowire:constructFunc=new

type ConfigMap map[string]interface{}

func (ci *ConfigMap) Value() map[string]interface{} {
	return *ci
}

func (ci *ConfigMap) new(impl *ConfigMap) (*ConfigMap, error) {
	*impl = *ci
	return impl, nil
}

func FromMap(val map[string]interface{}) *ConfigMap {
	configMap := ConfigMap(val)
	return &configMap
}

// +ioc:autowire=true
// +ioc:autowire:baseType=true
// +ioc:autowire:type=config
// +ioc:autowire:paramType=ConfigSlice
// +ioc:autowire:constructFunc=new

type ConfigSlice []interface{}

func (cs *ConfigSlice) Value() []interface{} {
	return *cs
}

func (ci *ConfigSlice) new(impl *ConfigSlice) (*ConfigSlice, error) {
	*impl = *ci
	return impl, nil
}

func FromSlice(val []interface{}) *ConfigSlice {
	configSlice := ConfigSlice(val)
	return &configSlice
}

// +ioc:autowire=true
// +ioc:autowire:baseType=true
// +ioc:autowire:type=config
// +ioc:autowire:paramType=ConfigString
// +ioc:autowire:constructFunc=new

type ConfigString string

func (ci *ConfigString) Value() string {
	return string(*ci)
}

func (ci *ConfigString) new(impl *ConfigString) (*ConfigString, error) {
	*impl = *ci
	return impl, nil
}

func FromString(val string) *ConfigString {
	configInt := ConfigString(val)
	return &configInt
}
