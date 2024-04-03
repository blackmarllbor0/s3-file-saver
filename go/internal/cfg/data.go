package cfg

type appCfg struct {
	Repo repoCfg `yaml:"Repo"`
}

type repoCfg struct {
	Minio    minioCfg `yaml:"Minio"`
	Postgres pgCfg    `yaml:"Postgres"`
}

type minioCfg struct {
	User   string `yaml:"User"`
	Pwd    string `yaml:"Pwd"`
	Host   string `yaml:"Host"`
	Bucket string `yaml:"Bucket"`
	Region string `yaml:"Region"`
	Port   uint   `yaml:"Port"`
}

type pgCfg struct {
	ConnString                 string `yaml:"ConnString"`
	MaxCons                    uint16 `yaml:"MaxCons"`
	MinCons                    uint16 `yaml:"MinCons"`
	MaxConnLifetimeInMinutes   uint32 `yaml:"MaxConnLifetimeInMinutes"`
	MaxConnIdleTimeInMinutes   uint32 `yaml:"MaxConnIdleTimeInMinutes"`
	HealthCheckPeriodInSeconds uint32 `yaml:"HealthCheckPeriodInSeconds"`
	ConnectTimeoutInSeconds    uint32 `yaml:"ConnectTimeoutInSeconds"`
}
