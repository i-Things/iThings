// Code generated by goctl. DO NOT EDIT.
// Source: job.proto

package job

import (
	"context"

	"github.com/i-Things/things/src/timedqueuesvr/job/internal/svc"
	"github.com/i-Things/things/src/timedqueuesvr/job/pb/job"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = job.Request
	Response = job.Response

	Job interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultJob struct {
		cli zrpc.Client
	}

	directJob struct {
		svcCtx *svc.ServiceContext
		svr    job.JobServer
	}
)

func NewJob(cli zrpc.Client) Job {
	return &defaultJob{
		cli: cli,
	}
}

func NewDirectJob(svcCtx *svc.ServiceContext, svr job.JobServer) Job {
	return &directJob{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultJob) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := job.NewJobClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (d *directJob) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.Ping(ctx, in)
}
