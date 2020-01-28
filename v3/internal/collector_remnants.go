package internal

import (
	"encoding/json"
	"fmt"
)

const (
	// Methods used in collector communication.
	cmdMetrics      = "metric_data"
	cmdCustomEvents = "custom_event_data"
	cmdTxnEvents    = "analytic_event_data"
	cmdErrorEvents  = "error_event_data"
	cmdErrorData    = "error_data"
	cmdTxnTraces    = "transaction_sample_data"
	cmdSlowSQLs     = "sql_trace_data"
	cmdSpanEvents   = "span_event_data"
)

// ConstructConnectReply takes the body of a Connect reply, in the form of bytes, and a
// PreconnectReply, and converts it into a *ConnectReply
func ConstructConnectReply(body []byte, preconnect PreconnectReply) (*ConnectReply, error) {
	var reply struct {
		Reply *ConnectReply `json:"return_value"`
	}
	reply.Reply = ConnectReplyDefaults()
	err := json.Unmarshal(body, &reply)
	if nil != err {
		return nil, fmt.Errorf("unable to parse connect reply: %v", err)
	}

	reply.Reply.PreconnectReply = preconnect

	reply.Reply.rulesCache = newRulesCache(txnNameCacheLimit)

	return reply.Reply, nil
}
