package enum

type DataDemision string

const (
	DataDemision_EXPOSE         DataDemision = "EXPOSE"
	DataDemision_CLICK          DataDemision = "CLICK"
	DataDemision_DOWNLOAD       DataDemision = "DOWNLOAD"
	DataDemision_COST           DataDemision = "COST"
	DataDemision_CLICK_PRICE    DataDemision = "CLICK_PRICE"
	DataDemision_DOWNLOAD_PRICE DataDemision = "DOWNLOAD_PRICE"
	DataDemision_CTR            DataDemision = "CTR"
	DataDemision_ECPM           DataDemision = "ECPM"
	DataDemision_DOWNLOAD_RATE  DataDemision = "DOWNLOAD_RATE"
)
