package grpc

import (
	"context"
	tagv1 "github.com/MuxiKeStack/be-api/gen/proto/tag/v1"
	"github.com/MuxiKeStack/be-tag/service"
	"google.golang.org/grpc"
)

type TagServiceServer struct {
	tagv1.UnimplementedTagServiceServer
	svc service.TagService
}

func NewTagServiceServer(svc service.TagService) *TagServiceServer {
	return &TagServiceServer{svc: svc}
}

func (s *TagServiceServer) Register(server grpc.ServiceRegistrar) {
	tagv1.RegisterTagServiceServer(server, s)
}

func (s *TagServiceServer) AttachAssessmentTags(ctx context.Context,
	request *tagv1.AttachAssessmentTagsRequest) (*tagv1.AttachAssessmentTagsResponse, error) {
	err := s.svc.AttachAssessmentTags(ctx, request.GetTaggerId(), request.GetBiz(), request.GetBizId(), request.GetTags())
	return &tagv1.AttachAssessmentTagsResponse{}, err
}

func (s *TagServiceServer) AttachFeatureTags(ctx context.Context,
	request *tagv1.AttachFeatureTagsRequest) (*tagv1.AttachFeatureTagsResponse, error) {
	err := s.svc.AttachFeatureTags(ctx, request.GetTaggerId(), request.GetBiz(), request.GetBizId(), request.GetTags())
	return &tagv1.AttachFeatureTagsResponse{}, err
}
