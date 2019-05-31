package svr

import (
	"github.com/iopig/feed-scale/interface/grpc/go_out/fsapi"
	"golang.org/x/net/context"
)

type SvrApi struct {
}

func (c *SvrApi) PadLogin(ctx context.Context, in *fsapi.DevInfoReq) (*fsapi.PigstyInfoRes, error) {
	return nil, nil
}
func (c *SvrApi) LoadCmd(ctx context.Context, in *fsapi.LoadReq) (*fsapi.ResHeader, error) {
	return nil, nil
}
func (c *SvrApi) ChoosePigsty(ctx context.Context, in *fsapi.ChoosePigstyReq) (*fsapi.CurrentFedRes, error) {
	return nil, nil
}
func (c *SvrApi) UploadRawInfo(ctx context.Context, in *fsapi.ChoosePigstyReq) (*fsapi.CurrentFedRes, error) {
	return nil, nil
}
