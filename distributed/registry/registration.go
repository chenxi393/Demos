package registry

type Registration struct {
	ServiceName      ServiceName
	ServiceURL       string
	RequiredServices []ServiceName //所依赖的服务
	// 我们并不需要保存服务所依赖服务的所有URL 当然服务名是需要的
	ServiceUpdateURL string        //动态更新所依赖服务的字段 
	// 这里就是一个string 不是数组 对应了每个服务 /services URL
	// 用于更新服务所需的服务的URL
	// 用于心态监测的URL
	HeartbeatURL string
}

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
	PortalService = ServiceName("PortalService")
)

type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
