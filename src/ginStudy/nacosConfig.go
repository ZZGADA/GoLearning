package ginstudy

import (
	"basicLearning/src/config"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/util"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"reflect"
	"time"
)

func InitConfigClient() (config_client.IConfigClient, error) {
	// Nacos服务器地址  可以多个地址
	// nacos context path 和 通信规则默认是/nacos和http
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848, constant.WithContextPath("/nacos")),
	}
	// 客户端配置
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId("public"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("../logs/nacos"),
		constant.WithCacheDir("../cache/nacos"),
		constant.WithLogLevel("debug"),
	)

	// 创建配置客户端
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {
		panic(err)
	}

	return configClient, nil
}

// 配置中心功能 -->动态配置
func InitNacosConfigClientTest() {
	//create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848, constant.WithContextPath("/nacos")),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		//constant.WithNamespaceId("public"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
		//constant.WithUsername("nacos"),
		//constant.WithPassword("nacos"),
	)

	// create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	fmt.Println("client  type is ", reflect.TypeOf(client).Kind())
	//fmt.Printf("class %#v", client)

	if err != nil {
		panic(err)
	}

	//publish config
	//config key=dataId+group+namespaceId
	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "test-data",
		Group:   "test-group",
		Content: "hello world!",
	})
	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "test-data-2",
		Group:   "test-group",
		Content: "hello world!",
	})
	if err != nil {
		fmt.Printf("PublishConfig err:%+v \n", err)
	}
	time.Sleep(1 * time.Second)
	//get config
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "test-data",
		Group:  "test-group",
	})
	fmt.Println("GetConfig,config :" + content)

	//Listen config change,key=dataId+group+namespaceId.
	// 初始化一个监听器 监听test-data-2的变化
	err = client.ListenConfig(vo.ConfigParam{
		DataId: "test-data",
		Group:  "test-group",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
		},
	})

	err = client.ListenConfig(vo.ConfigParam{
		DataId: "test-data-2",
		Group:  "test-group",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
		},
	})

	time.Sleep(1 * time.Second)

	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "test-data",
		Group:   "test-group",
		Content: "test-listen",
	})

	time.Sleep(1 * time.Second)

	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "test-data-2",
		Group:   "test-group",
		Content: "test-listen",
	})

	time.Sleep(2 * time.Second)

	time.Sleep(1 * time.Second)
	_, err = client.DeleteConfig(vo.ConfigParam{
		DataId: "test-data",
		Group:  "test-group",
	})
	time.Sleep(1 * time.Second)

	//cancel config change
	err = client.CancelListenConfig(vo.ConfigParam{
		DataId: "test-data",
		Group:  "test-group",
	})

	searchPage, _ := client.SearchConfig(vo.SearchConfigParam{
		Search:   "blur",
		DataId:   "",
		Group:    "",
		PageNo:   1,
		PageSize: 10,
	})
	fmt.Printf("Search config:%+v \n", searchPage)

	config_val, _ := client.GetConfig(vo.ConfigParam{
		DataId: "basic-learning-go.yaml",
		Group:  "GO",
	})

	fmt.Printf("get one config:%+v \n", config_val)
}

func ReadNacosConfigYaml() (string, error) {
	//create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848, constant.WithContextPath("/nacos")),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
		constant.WithUsername("nacos"),
		constant.WithPassword("nacos"),
	)

	// create config client
	clientRead, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	val, err := clientRead.GetConfig(vo.ConfigParam{
		DataId: "basic-learning-go.yaml",
		Group:  "GO",
	})
	fmt.Println("ioio")
	return val, err
}

// nacos 服务发现功能 包括服务注册和发现
func InitNacosServicesClientTest() {
	//create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848, constant.WithContextPath("/nacos")),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// create naming client
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	//Register
	registerServiceInstance(client, vo.RegisterInstanceParam{
		Ip:          "10.0.0.10",
		Port:        8848,
		ServiceName: "demo.go",
		GroupName:   "group-a",
		ClusterName: "cluster-a",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
	})

	//DeRegister
	deRegisterServiceInstance(client, vo.DeregisterInstanceParam{
		Ip:          "10.0.0.10",
		Port:        8848,
		ServiceName: "demo.go",
		GroupName:   "group-a",
		Cluster:     "cluster-a",
		Ephemeral:   true, //it must be true
	})

	time.Sleep(1 * time.Second)

	//BatchRegister
	batchRegisterServiceInstance(client, vo.BatchRegisterInstanceParam{
		ServiceName: "demo.go",
		GroupName:   "group-a",
		Instances: []vo.RegisterInstanceParam{{
			Ip:          "10.0.0.10",
			Port:        8848,
			Weight:      10,
			Enable:      true,
			Healthy:     true,
			Ephemeral:   true,
			ClusterName: "cluster-a",
			Metadata:    map[string]string{"idc": "shanghai"},
		}, {
			Ip:          "10.0.0.12",
			Port:        8848,
			Weight:      7,
			Enable:      true,
			Healthy:     true,
			Ephemeral:   true,
			ClusterName: "cluster-a",
			Metadata:    map[string]string{"idc": "shanghai"},
		}},
	})

	time.Sleep(1 * time.Second)

	//Get service with serviceName, groupName , clusters
	getService(client, vo.GetServiceParam{
		ServiceName: "demo.go",
		GroupName:   "group-a",
		Clusters:    []string{"cluster-a"},
	})

	//SelectAllInstance
	//GroupName=DEFAULT_GROUP
	selectAllInstances(client, vo.SelectAllInstancesParam{
		ServiceName: "demo.go",
		GroupName:   "group-a",
		Clusters:    []string{"cluster-a"},
	})

	//SelectInstances only return the instances of healthy=${HealthyOnly},enable=true and weight>0
	//ClusterName=DEFAULT,GroupName=DEFAULT_GROUP
	selectInstances(client, vo.SelectInstancesParam{
		ServiceName: "demo.go",
		GroupName:   "group-a",
		Clusters:    []string{"cluster-a"},
		HealthyOnly: true,
	})

	//SelectOneHealthyInstance return one instance by WRR strategy for load balance
	//And the instance should be health=true,enable=true and weight>0
	//ClusterName=DEFAULT,GroupName=DEFAULT_GROUP
	selectOneHealthyInstance(client, vo.SelectOneHealthInstanceParam{
		ServiceName: "demo.go",
		GroupName:   "group-a",
		Clusters:    []string{"cluster-a"},
	})

	//Subscribe key=serviceName+groupName+cluster
	//Note:We call add multiple SubscribeCallback with the same key.
	subscribeParam := &vo.SubscribeParam{
		ServiceName: "demo.go",
		GroupName:   "group-a",
		SubscribeCallback: func(services []model.Instance, err error) {
			fmt.Printf("callback return services:%s \n\n", util.ToJsonString(services))
		},
	}
	subscribe(client, subscribeParam)

	//wait for client pull change from server
	time.Sleep(3 * time.Second)

	updateServiceInstance(client, vo.UpdateInstanceParam{
		Ip:          "10.0.0.11", //update ip
		Port:        8848,
		ServiceName: "demo.go",
		GroupName:   "group-a",
		ClusterName: "cluster-a",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "beijing1"}, //update metadata
	})

	//wait for client pull change from server
	time.Sleep(3 * time.Second)
	// UnSubscribe
	unSubscribe(client, subscribeParam)

	//GeAllService will get the list of service name
	//NameSpace default value is public.If the client set the namespaceId, NameSpace will use it.
	//GroupName default value is DEFAULT_GROUP
	getAllService(client, vo.GetAllServiceInfoParam{
		GroupName: "group-a",
		PageNo:    1,
		PageSize:  10,
	})
}

func InitNacosConfigCenter(config *config.NacosConfig) config_client.IConfigClient {
	//create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(config.Ip, config.Port, constant.WithContextPath(config.ContextPath)),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithTimeoutMs(config.TimeoutMs),
		constant.WithNotLoadCacheAtStart(config.NotLoadCacheAtStart),
		constant.WithLogDir(config.Dir.Log),
		constant.WithCacheDir(config.Dir.Cache),
		constant.WithLogLevel(config.LogLevel),
		constant.WithUsername(config.Username),
		constant.WithPassword(config.Password),
	)

	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	return client
}
