package saga

import (
	rpc "github.com/channel-io/ch-app-store/internal/rpc/domain"
)

type ContextInjectingInvoker[REQ rpc.ContextAware, RES any] struct {
	invoker rpc.ContextAwareInvokeSvc[REQ, RES]
}

type PreContext struct {
}

type Request[REQ rpc.ContextAware] struct {
	PreContext PreContext
	Req        REQ
}

func (i *ContextInjectingInvoker[REQ, RES]) Invoke() (RES, error) {

}
