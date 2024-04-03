package cfg

type appCfg struct {
	Repo repoCfg
}

type repoCfg struct {
	Minio minioCfg
}

type minioCfg struct {
	User   string `yaml:"User"`
	Pwd    string `yaml:"Pwd"`
	Host   string `yaml:"Host"`
	Bucket string `yaml:"Bucket"`
	Region string `yaml:"Region"`
	Port   uint   `yaml:"Port"`
}
