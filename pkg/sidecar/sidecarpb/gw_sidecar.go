
//gw_sidecar.go

// Avesha LLC
// Feb 2022
//
// Copyright (c) Avesha LLC. 2022
//
// Module: Avesha Sidecar - Slice Controller Commnunication GRPC Module


// GwSidecar represents the GRPC Gateway sidecar

package sidecar

import(
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GwSidecar struct {
	UnimplementedGwSidecarServiceServer
}

// GetStatus get the status of sidecar.
func (s *GwSidecar) GetStatus(ctx context.Context, in *empty.Empty) (*GwPodStatus, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "Client cancelled, abandoning.")
	}
	podStatus, err := getGwPodStatus()
	return podStatus, err
}

func (s *GwSidecar) UpdateConnectionContext(ctx context.Context, in *SliceGwConnectionContext) (*SidecarResponse, error){
	return &SidecarResponse{},nil
}

func (s *GwSidecar) UpdateSliceQosProfile(ctx context.Context, in *SliceQosProfile) (*SidecarResponse, error){
	return &SidecarResponse{},nil
}