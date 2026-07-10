package config

type JWTConf struct {
	AccessSecret         string
	RefreshSecret        string
	AccessExpireSeconds  int64
	RefreshExpireSeconds int64
}

type MySQLConf struct {
	DataSource string
}

type RedisConf struct {
	Addr     string
	Password string
	DB       int
}

type LLMProviderConf struct {
	APIKey  string
	BaseURL string
}

type LLMConf struct {
	OpenAI LLMProviderConf
}
