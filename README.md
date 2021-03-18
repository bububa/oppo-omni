# OPPO移动广告Omni API Golang SDK

- 用户相关
  - 代理商客户查询[ api.communal.agency.CustomerList(clt *core.SDKClient, req *agency.CustomerListRequest) (*agency.CustomerListResult, error) ]
  - 代理商余额查询[ api.communal.agency.Balance(clt *core.SDKClient) ([]agency.BalanceAccount, error) ]
  - 客户余额查询[ api.communal.owner.Balance(clt *core.SDKClient, ownerID uint64) ([]owner.BalanceAccount, error) ]
  - 客户日预算设置[ api.communal.owner.SetAccDayBudget(clt *core.SDKClient, ownerID uint64, budget int64) error ]
  - 账务流水查询[ api.communal.finance.BillHis(clt *core.SDKClient, req *finance.BillHisRequest) (*finance.BillHisResult, error) ]
- 转化回传 
  - H5 数据回传转化[ api.clue.SendData(clt *core.SDKClient, req *clue.SendDataRequest) error ]
- 数据
  - 报表
    - 整体数据 [ api.data.QTodayTotal(clt *core.SDKClient, ownerID uint64, adID uint64) (*data.QTodayTotalResult, error) ]
    - TOP 榜 [ QTodayTop(clt *core.SDKClient, ownerID uint64, deminsion *enum.DataDemision) (*data.QTodayTotalResult, error) ] 
    - 广告计划-图表 [ QPlan(clt *core.SDKClient, req *data.QPlanRequest) (*data.QResult, error) ]
    - 广告计划-列表 [ QPlanList(clt *core.SDKClient, req *data.QPlanListRequest) (*data.QListResult, error) ]
    - 广告组-图表 [ QAdgroup(clt *core.SDKClient, req *data.QAdgroupRequest) (*data.QResult, error) ]
    - 广告组-列表 [ QAdgroupList(clt *core.SDKClient, req *data.QAdgroupListRequest) (*data.QListResult, error) ]
    - 广告-图表 [ QAd(clt *core.SDKClient, req *data.QAdRequest) (*data.QResult, error) ]
    - 广告-列表 [ QAdList(clt *core.SDKClient, req *data.QAdListRequest) (*data.QListResult, error) ]
    - 关键词-图表 [ QKw(clt *core.SDKClient, req *data.QKwRequest) (*data.QResult, error) ]
    - 关键词-列表 [ QKwList(clt *core.SDKClient, req *data.QKwListRequest) (*data.QListResult, error) ]
    - 游戏-图表 [ QGame(clt *core.SDKClient, req *data.QGameRequest) (*data.QGameResult, error) ]
    - 游戏-列表 [ QGameList(clt *core.SDKClient, req *data.QGameListRequest) (*data.QListResult, error) ]
    - 小游戏-图表 [ QQuickAppGame(clt *core.SDKClient, req *data.QQuickAppGameRequest) (*data.QQuickAppGameResult, error) ]
    - 小游戏-列表 [ QQuickAppGameList(clt *core.SDKClient, req *data.QQuickAppGameListRequest) (*data.QQuickAppGameListResult, error) ]
- 监测链接
  - 点击检测[ api.track.ClickUrl(baseUrl string, fields []string) string ]

