package plugin

type MetricTemplate struct {
	WarnSwitch  bool    `bson:"WarnSwitch"`
	AlertSwitch bool    `bson:"AlertSwitch"`
	SMSSwitch   bool    `bson:"SMSSwitch"`
	Min         int     `bson:"Min"`
	Max         int     `bson:"Max"`
	Scale       int     `bson:"Scale"`
	Warn        float32 `bson:"Warn"`
	Alert       float32 `bson:"Alert"`
}

type CurrentTemplate struct {
	MetricTemplate
	SPL int `bson:"SPL"`
}

type TemperatureTemplate struct {
	MetricTemplate
}

type VoltageTemplate struct {
	MetricTemplate
	WarnSMSSwitch bool `bson:"WarnSMSSwitch"`
}

type IOTemplate struct {
	AlertSwitch bool `bson:"AlertSwitch"`
	SMSSwitch   bool `bson:"SMSSwitch"`
}

type IRTemplate struct {
	WarnSwitch  bool    `bson:"WarnSwitch"`
	AlertSwitch bool    `bson:"AlertSwitch"`
	SMSSwitch   bool    `bson:"SMSSwitch"`
	Warn        float32 `bson:"Warn"`
	Alert       float32 `bson:"Alert"`
	SPL         int     `bson:"SPL"`
}

type T4Template struct {
	WarnSwitch  bool    `bson:"WarnSwitch"`
	AlertSwitch bool    `bson:"AlertSwitch"`
	SMSSwitch   bool    `bson:"SMSSwitch"`
	Warn        float32 `bson:"Warn"`
	Alert       float32 `bson:"Alert"`
}
