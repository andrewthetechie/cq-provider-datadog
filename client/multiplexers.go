package client

import (
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func AccountMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for _, account := range client.AccountConfigs {
		l = append(l, client.withAccount(account))
	}
	return l
}
