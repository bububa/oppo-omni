# OPPO移动广告Omni API Golang SDK

- 用户相关
  - 代理商客户查询[ api.communal.agency.CustomerList(clt *core.SDKClient, req *agency.CustomerListRequest) (*agency.CustomerListResult, error) ]
  - 代理商余额查询[ api.communal.agency.Balance(clt *core.SDKClient) ([]agency.BalanceAccount, error) ]
  - 客户余额查询[ api.communal.owner.Balance(clt *core.SDKClient, ownerID uint64) ([]owner.BalanceAccount, error) ]
  - 客户日预算设置[ api.communal.owner.SetAccDayBudget(clt *core.SDKClient, ownerID uint64, budget int64) error ]
  - 账务流水查询[ api.communal.finance.BillHis(clt *core.SDKClient, req *finance.BillHisRequest) (*finance.BillHisResult, error) ]
- 转化回传 
  - H5 数据回传转化[ api.clue.SendData(clt *core.SDKClient, req *clue.SendDataRequest) error ]
- 监测链接
  - 点击检测[ api.track.ClickUrl(baseUrl string, fields []string) string ]

