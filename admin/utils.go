package admin

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2/rlog"
)

func FetchMasterAddrByClusterName(ctx context.Context, admin *admin, clusterName string) ([]string, error) {
	var brokerAddrs []string
	clusterInfo, err := admin.cli.GetBrokerClusterInfo(ctx)
	if err != nil {
		rlog.Error("Fail to get clusterInfo", nil)
		return nil, err
	}
	brokerNames := clusterInfo.ClusterAddrTable[clusterName]
	for _, brokerName := range brokerNames {
		brokerAddrs = append(brokerAddrs, clusterInfo.BrokerAddrTable[brokerName].BrokerAddresses[0])
	}
	return brokerAddrs, nil
}
