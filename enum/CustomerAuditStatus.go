package enum

type CustomerAuditStatus int

const (
	CustomerAuditStatus_PENDDING CustomerAuditStatus = 0 // 未审核
	CustomerAuditStatus_PASS     CustomerAuditStatus = 1 //  审核通过
	CustomerAuditStatus_REJECT   CustomerAuditStatus = 2 //  审核不通过
)
