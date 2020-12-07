package happysooner_gRPC_client

import (
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/grpc"
	"github.com/unliar/happysooner-proto/v2/account"
	"github.com/unliar/happysooner-proto/v2/counter"
	"github.com/unliar/happysooner-proto/v2/pay"
	"github.com/unliar/happysooner-proto/v2/push"
	"github.com/unliar/happysooner-proto/v2/search"
	"github.com/unliar/happysooner-proto/v2/writing"
	"time"
)

var c = grpc.NewClient(func(options *client.Options) {
	options.CallOptions.RequestTimeout = 15 * time.Second
})

var AccountService = account.NewAccountSVService(AccountRpcName, c)

var PushService = push.NewPushSVService(PushRpcName, c)

var WritingService = writing.NewWritingSVService(WritingRpcName, c)

var SearchService = search.NewSearchSVService(SearchRpcName, c)

var PayService = pay.NewPaySVService(PayRpcName, c)

var CounterService = counter.NewCounterSVService(CounterRpcName, c)
