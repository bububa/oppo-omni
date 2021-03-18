package enum

type DataChnID int

const (
	DataChnID_APP_STORE     DataChnID = 1 // 软件商店
	DataChnID_NOT_APP_STORE DataChnID = 2 // 非软件商店
	DataChnID_UNION         DataChnID = 3 // 联盟
)
