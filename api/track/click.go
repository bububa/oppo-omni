package track

import (
	"fmt"
	"net/url"
	"strings"
)

var DEFAULT_CLICK_FIELDS = []string{
	"req_id",
	"ad_id",
	"ad_name",
	"os",
	"os_version",
	"model",
	"lang",
	"country",
	"width",
	"height",
	"pkg",
	"app_version",
	"useragent",
	"referer",
	"net_type",
	"carrier",
	"progress",
	"imei",
	"oaid",
}

var CLICK_FIELDS_MAP = map[string]string{
	"req_id":      "$req$",
	"ad_id":       "$ad$",
	"ad_name":     "$an$",
	"os":          "$os$",
	"os_version":  "$ov$",
	"model":       "$m$",
	"lang":        "$lan$",
	"country":     "$c$",
	"width":       "$w$",
	"height":      "$h$",
	"pkg":         "$pkg$",
	"app_version": "$av$",
	"useragent":   "$ua$",
	"referer":     "$rf$",
	"net_type":    "$nt$",
	"carrier":     "$ca$",
	"progress":    "$progress$",
	"imei":        "$im$",
	"oaid":        "__OAID__",
}

// 点击检测链接
func ClickUrl(baseUrl string, fields []string) string {
	if fields == nil {
		fields = DEFAULT_CLICK_FIELDS
	}
	parsedUrl, _ := url.Parse(baseUrl)
	parsedUrl.Fragment = ""
	var query []string
	for _, field := range fields {
		if val, found := CLICK_FIELDS_MAP[field]; found {
			query = append(query, fmt.Sprintf("%s=%s", field, val))
		}
	}
	values := parsedUrl.Query()
	if len(values) > 0 {
		return parsedUrl.String() + "&" + strings.Join(query, "&")
	}
	return parsedUrl.String() + "?" + strings.Join(query, "&")
}
