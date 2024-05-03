package server

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	goodPb "shopping/api/good"
	"shopping/internal/data/model"
	"shopping/internal/data/query"
)

type GoodServer struct {
	goodPb.UnimplementedGoodServer
}

func NewGoodServer() *GoodServer {
	return &GoodServer{}
}

func (g *GoodServer) ListGoodByKeyword(ctx context.Context, req *goodPb.ListGoodByKeywordRequest) (*goodPb.ListGoodByKeywordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGoodByKeyword not implemented")
}
func (g *GoodServer) ListGoodByCategory(ctx context.Context, req *goodPb.ListGoodByCategoryRequest) (*goodPb.ListGoodByCategoryResponse, error) {
	good := query.Good
	goods, err := good.WithContext(ctx).Where(good.Category.Eq(req.Category)).Find()
	if err != nil {
		zap.S().Errorf("[ListGoodByCategory] find good failed, err: %v", err)
		return nil, status.Errorf(codes.NotFound, "good not found")
	}
	var resp goodPb.ListGoodByCategoryResponse
	for _, v := range goods {
		var postFree bool
		if v.Postfree == 1 {
			postFree = true
		}
		resp.GoodList = append(resp.GoodList, &goodPb.GoodStruct{
			Id:          v.ID,
			Price:       float32(v.Price),
			Description: v.Description,
			Category:    v.Category,
			Shop:        v.Shop,
			PostFree:    postFree,
			Location:    v.Location,
			Deal:        int64(v.Deal),
		})
	}
	return &resp, nil
}
func (g *GoodServer) AddGood(ctx context.Context, req *goodPb.AddGoodRequest) (*goodPb.AddGoodResponse, error) {
	good := query.Good
	var postFree int32
	if req.PostFree {
		postFree = 1
	}
	goo := &model.Good{
		Price:       float64(req.Price),
		Deal:        int32(req.Deal),
		Description: req.Description,
		Shop:        req.Shop,
		Location:    req.Location,
		Postfree:    postFree,
		Category:    req.Category,
	}
	err := good.WithContext(ctx).Create(goo)
	if err != nil {
		zap.S().Errorf("[AddGood] create good failed, err: %v", err)
		return nil, status.Errorf(codes.Internal, "create good failed")
	}
	return &goodPb.AddGoodResponse{Id: goo.ID}, nil
}
func (g *GoodServer) GetGood(ctx context.Context, req *goodPb.GetGoodRequest) (*goodPb.GetGoodResponse, error) {
	good := query.Good
	goo, err := good.WithContext(ctx).Where(good.ID.Eq(req.Id)).First()
	if err != nil {
		zap.S().Errorf("[GetGood] find good failed, err: %v", err)
		return nil, status.Errorf(codes.NotFound, "good not found")
	}
	var postFree bool
	if goo.Postfree == 1 {
		postFree = true

	}
	return &goodPb.GetGoodResponse{
		Good: &goodPb.GoodStruct{
			Id:          goo.ID,
			Description: goo.Description,
			Price:       float32(goo.Price),
			Shop:        goo.Shop,
			Category:    goo.Category,
			Deal:        int64(goo.Deal),
			PostFree:    postFree,
			Location:    goo.Location,
		},
	}, nil
}
func (g *GoodServer) UpdateGood(ctx context.Context, req *goodPb.UpdateGoodRequest) (*goodPb.UpdateGoodResponse, error) {
	good := query.Good
	var postFree int32
	if req.Good.PostFree {
		postFree = 1
	}
	_, err := good.WithContext(ctx).Where(good.ID.Eq(req.Good.Id)).Updates(&model.Good{
		Price:       float64(req.Good.Price),
		Deal:        int32(req.Good.Deal),
		Postfree:    postFree,
		Description: req.Good.Description,
		Shop:        req.Good.Shop,
		Location:    req.Good.Location,
		Category:    req.Good.Category,
	})
	if err != nil {
		zap.S().Errorf("[UpdateGood] update good failed, err: %v", err)
		return nil, status.Errorf(codes.Internal, "update good failed")
	}
	return &goodPb.UpdateGoodResponse{Id: req.Good.Id}, nil
}
func (g *GoodServer) DeleteGood(ctx context.Context, req *goodPb.DeleteGoodRequest) (*goodPb.DeleteGoodResponse, error) {
	good := query.Good
	_, err := good.WithContext(ctx).Where(good.ID.Eq(req.Id)).Delete()
	if err != nil {
		zap.S().Errorf("[DeleteGood] delete good failed, err: %v", err)
		return nil, status.Errorf(codes.NotFound, "good not found")
	}
	return &goodPb.DeleteGoodResponse{Id: req.Id}, nil
}
