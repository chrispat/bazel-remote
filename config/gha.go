package config

type GitHubActionsCacheConfig struct {
	AuthToken string `yaml:"auth_token"`
	CacheUrl  string `yaml:"cache_url"`
}
