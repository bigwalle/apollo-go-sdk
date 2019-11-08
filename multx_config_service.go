package apollosdk

import (
	"github.com/welcome112s/apollo-go-sdk/core"
	"github.com/welcome112s/apollo-go-sdk/util"
	"sync"
)

const (
	NAMESPACE_APPLICATION       = "application"
	CLUSTER_NAME_DEFAULT        = "default"
	CLUSTER_NAMESPACE_SEPARATOR = "+"
)

var (
	lock       sync.Mutex
	once       sync.Once
)
type ApolloSdk struct {
	ConfitUtil core.ConfitUtil
	ConfigMap  map[string]*core.Config
}

func NewApolloSdk(c core.CConfig)*ApolloSdk{
	return &ApolloSdk{
		ConfitUtil: core.NewConfigWithApolloInitConfig(core.ApolloInitConfig{
			AppId:    c.AppId  ,
			Cluster:    c.Cluster,
			MetaServer: c.MmetaServer,
			DataCenter: c.DataCenter,
		}),
		ConfigMap:make(map[string]*core.Config, 10),
	}
}
//启动默认配置
func init() {
	util.SetDebug(false)
}

func (ApolloSdk)SetDebug(debug bool) {
	util.SetDebug(debug)
}

//自定义配置文件进行配置
func (a ApolloSdk)StartWithCusConfig(configFile string) {
	a.ConfitUtil = core.NewConfigWithConfigFile(configFile)
}

func (a *ApolloSdk)GetConfig(namespace string) core.Config {
	lock.Lock()
	defer lock.Unlock()
	config, ok := a.ConfigMap[namespace]

	if !ok {
		remoteRepository := core.NewRemoteConfigRepository(namespace, a.ConfitUtil)

		repository := core.ConfigRepository(remoteRepository)

		defaultConfig := core.NewDefaultConfig(namespace, repository,  a.ConfitUtil)

		config := core.Config(defaultConfig)
		a.ConfigMap[namespace] = &config
		return config
	}
	return *config
}

func (a *ApolloSdk)GetAppConfig() core.Config{
	return  a.GetConfig(NAMESPACE_APPLICATION)
}
