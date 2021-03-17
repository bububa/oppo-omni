package model

const (
	BASE_URL             = "https://sapi.ads.heytapmobi.com"
	DEFAULT_VERSION      = "v1"
	DEFAULT_CONTENT_TYPE = "application/json"
)

type Request interface {
	GetOwnerID() uint64
	GetBaseUrl() string
	GetVersion() string
	GetContentType() string
	ResourceName() string
	ResourceAction() string
}

type BaseRequest struct {
	OwnerID           uint64 `json:"ownerId,omitempty"`
	apiBaseUrl        string `json:"-"`
	apiVersion        string `json:"-"`
	apiContentType    string `json:"-"`
	apiResourceName   string `json:"-"`
	apiResourceAction string `json:"-"`
}

func (r *BaseRequest) SetOwnerID(ownerID uint64) {
	r.OwnerID = ownerID
}

func (r BaseRequest) GetOwnerID() uint64 {
	return r.OwnerID
}

func (r *BaseRequest) SetResourceName(resourceName string) {
	r.apiResourceName = resourceName
}

func (r BaseRequest) ResourceName() string {
	return r.apiResourceName
}

func (r *BaseRequest) SetResourceAction(resourceAction string) {
	r.apiResourceAction = resourceAction
}

func (r BaseRequest) ResourceAction() string {
	return r.apiResourceAction
}

func (r *BaseRequest) SetBaseUrl(baseUrl string) {
	r.apiBaseUrl = baseUrl
}

func (r BaseRequest) GetBaseUrl() string {
	if r.apiBaseUrl == "" {
		return BASE_URL
	}
	return r.apiBaseUrl
}

func (r *BaseRequest) SetVersion(ver string) {
	r.apiVersion = ver
}

func (r BaseRequest) GetVersion() string {
	if r.apiVersion == "" {
		return DEFAULT_VERSION
	}
	return r.apiVersion
}

func (r *BaseRequest) SetContentType(contentType string) {
	r.apiContentType = contentType
}

func (r BaseRequest) GetContentType() string {
	if r.apiContentType == "" {
		return DEFAULT_CONTENT_TYPE
	}
	return r.apiContentType
}
