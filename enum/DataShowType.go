package enum

type DataShowType int

const (
	DataShowType_BANNER        DataShowType = 1    // banner
	DataShowType_INTERSTITIAL  DataShowType = 2    // 插屏
	DataShowType_LAUNCH_SCREEN DataShowType = 3    // 开屏
	DataShowType_NATIVE        DataShowType = 8    // 原生（联盟）
	DataShowType_FEED          DataShowType = 16   // 信息流
	DataShowType_SEARCH        DataShowType = 32   // 搜索
	DataShowType_REWARD_VIDEO  DataShowType = 64   // 激励视频
	DataShowType_DISTRIBUTION  DataShowType = 128  // 分发推广
	DataShowType_SEARCH_CARD   DataShowType = 512  // 搜索卡片
	DataShowTYpe_UNLIMITED     DataShowType = 1024 // 不限
)
