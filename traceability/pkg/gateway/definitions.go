package gateway

import "strconv"

// CHANGE_HERE - Change the structures below to represent the log entry the agent is going to receive

// Headers - Type for request/response headers
type Headers map[string]string

// GwTransaction - Type for gateway transaction detail
type GwTransaction struct {
	ID              string  `json:"id"`
	SourceHost      string  `json:"srcHost"`
	SourcePort      int     `json:"srcPort,string"`
	DesHost         string  `json:"destHost"`
	DestPort        int     `json:"destPort,string"`
	URI             string  `json:"uri"`
	Method          string  `json:"method"`
	StatusCode      int     `json:"statusCode,string"`
	RequestHeaders  Headers `json:"requestHeaders"`
	ResponseHeaders Headers `json:"responseHeaders"`
	RequestBytes    int     `json:"requestByte"`
	ResponseBytes   int     `json:"responseByte"`
	BackendLatency  int     `json:"backendLatency"`
}

// GwTrafficLogEntry - Represents the structure of log entry the agent will receive
type GwTrafficLogEntry struct {
	TraceID             string        `json:"traceId"`
	APIName             string        `json:"apiName"`
	ResponseTime        int           `json:"respTime"`
	StartTime           int64         `json:"startTime,string"`
	InboundTransaction  GwTransaction `json:"inbound"`
	OutboundTransaction GwTransaction `json:"outbound"`
}

func (gt *GwTransaction) getDestPortString() string {
	if gt.DestPort > 0 {
		return strconv.Itoa(gt.DestPort)
	}
	return ""
}

func (gt *GwTransaction) getSourcePortString() string {
	if gt.SourcePort > 0 {
		return strconv.Itoa(gt.SourcePort)
	}
	return ""
}
