package owner

import "github.com/bububa/oppo-omni/model"

type SetAccDayBudgetRequest struct {
	model.BaseRequest
	AccDayBudget int64 `json:"accDayBudget,omitempty"`
}
