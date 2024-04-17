package grpc

import (
	"context"
	tagv1 "github.com/MuxiKeStack/be-api/gen/proto/tag/v1"
	"github.com/MuxiKeStack/be-tag/domain"
	"github.com/MuxiKeStack/be-tag/service"
	"github.com/ecodeclub/ekit/slice"
	"google.golang.org/grpc"
)

type TagServiceServer struct {
	tagv1.UnimplementedTagServiceServer
	svc service.GeneralTagService
}

func NewTagServiceServer(svc service.GeneralTagService) *TagServiceServer {
	return &TagServiceServer{svc: svc}
}

func (s *TagServiceServer) Register(server grpc.ServiceRegistrar) {
	tagv1.RegisterTagServiceServer(server, s)
}

func (s *TagServiceServer) CountAssessmentTagsByCourseTagger(ctx context.Context,
	request *tagv1.CountAssessmentTagsByCourseTaggerRequest) (*tagv1.CountAssessmentTagsByCourseTaggerResponse, error) {
	items, err := s.svc.CountTagsByBizTagger(ctx, tagv1.Biz_Course, request.GetCourseId(), request.GetTaggerIds(), domain.TagTypeAssessment)
	return &tagv1.CountAssessmentTagsByCourseTaggerResponse{
		Items: slice.Map(items, func(idx int, src domain.CountTagItem) *tagv1.CountAssessmentItem {
			return &tagv1.CountAssessmentItem{
				Tag:   tagv1.AssessmentTag(src.Tag),
				Count: src.Count,
			}
		}),
	}, err
}

func (s *TagServiceServer) CountFeatureTagsByCourseTagger(ctx context.Context,
	request *tagv1.CountFeatureTagsByCourseTaggerRequest) (*tagv1.CountFeatureTagsByCourseTaggerResponse, error) {
	items, err := s.svc.CountTagsByBizTagger(ctx, tagv1.Biz_Course, request.GetCourseId(), request.GetTaggerIds(), domain.TagTypeFeature)
	return &tagv1.CountFeatureTagsByCourseTaggerResponse{
		Items: slice.Map(items, func(idx int, src domain.CountTagItem) *tagv1.CountFeatureItem {
			return &tagv1.CountFeatureItem{
				Tag:   tagv1.FeatureTag(src.Tag),
				Count: src.Count,
			}
		}),
	}, err
}

func (s *TagServiceServer) GetAssessmentTagsByTaggerBiz(ctx context.Context, request *tagv1.GetAssessmentTagsByTaggerBizRequest) (*tagv1.GetAssessmentTagsByTaggerBizResponse, error) {
	tags, err := s.svc.GetTagsByTaggerBiz(ctx, request.GetTaggerId(), request.GetBiz(), request.GetBizId(), domain.TagTypeAssessment)
	return &tagv1.GetAssessmentTagsByTaggerBizResponse{
		Tags: slice.Map(tags, func(idx int, src int32) tagv1.AssessmentTag {
			return tagv1.AssessmentTag(src)
		}),
	}, err
}

func (s *TagServiceServer) GetFeatureTagsByTaggerBiz(ctx context.Context, request *tagv1.GetFeatureTagsByTaggerBizRequest) (*tagv1.GetFeatureTagsByTaggerBizResponse, error) {
	tags, err := s.svc.GetTagsByTaggerBiz(ctx, request.GetTaggerId(), request.GetBiz(), request.GetBizId(), domain.TagTypeFeature)
	return &tagv1.GetFeatureTagsByTaggerBizResponse{
		Tags: slice.Map(tags, func(idx int, src int32) tagv1.FeatureTag {
			return tagv1.FeatureTag(src)
		}),
	}, err
}

func (s *TagServiceServer) AttachAssessmentTags(ctx context.Context,
	request *tagv1.AttachAssessmentTagsRequest) (*tagv1.AttachAssessmentTagsResponse, error) {
	err := s.svc.AttachTags(ctx, request.GetTaggerId(), request.GetBiz(), request.GetBizId(),
		slice.Map(request.GetTags(), func(idx int, src tagv1.AssessmentTag) int32 {
			return int32(src)
		}), domain.TagTypeAssessment)
	return &tagv1.AttachAssessmentTagsResponse{}, err
}

func (s *TagServiceServer) AttachFeatureTags(ctx context.Context,
	request *tagv1.AttachFeatureTagsRequest) (*tagv1.AttachFeatureTagsResponse, error) {
	err := s.svc.AttachTags(ctx, request.GetTaggerId(), request.GetBiz(), request.GetBizId(),
		slice.Map(request.GetTags(), func(idx int, src tagv1.FeatureTag) int32 {
			return int32(src)
		}), domain.TagTypeFeature)
	return &tagv1.AttachFeatureTagsResponse{}, err
}
