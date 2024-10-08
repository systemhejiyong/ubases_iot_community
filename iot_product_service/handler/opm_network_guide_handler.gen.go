// Code generated by sgen.exe,2022-05-06 14:01:20. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package handler

import (
	"context"

	"cloud_platform/iot_product_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type OpmNetworkGuideHandler struct{}

// 设置配网类型
func (h *OpmNetworkGuideHandler) SetNetworkGuideTypes(ctx context.Context, req *proto.SetNetworkGuideTypeRequest, rsp *proto.OpmNetworkGuideResponse) error {
	s := service.OpmNetworkGuideSvc{Ctx: ctx}
	err := s.SetNetworkGuideTypes(req)
	if err != nil {
		rsp.Code = ERROR
		rsp.Message = err.Error()
	} else {
		rsp.Code = SUCCESS
		rsp.Message = "success"
	}
	return nil
}

// 删除
func (h *OpmNetworkGuideHandler) DeleteByProductId(ctx context.Context, req *proto.OpmNetworkGuideFilter, resp *proto.Response) error {
	s := service.OpmNetworkGuideSvc{Ctx: ctx}
	err := s.DeleteOpmNetworkGuideByProductId(req)
	SetResponse(resp, err)
	return nil
}

// 创建
func (h *OpmNetworkGuideHandler) CreateAndUpdate(ctx context.Context, req *proto.OpmNetworkGuide, resp *proto.Response) error {
	s := service.OpmNetworkGuideSvc{Ctx: ctx}
	_, err := s.SaveOpmNetworkGuide(req)
	SetResponse(resp, err)
	return nil
}

// 根据ID查找，返回单条数据
func (h *OpmNetworkGuideHandler) FindByProductId(ctx context.Context, req *proto.OpmNetworkGuideFilter, resp *proto.OpmNetworkGuideResponse) error {
	s := service.OpmNetworkGuideSvc{Ctx: ctx}
	data, err := s.FindOpmNetworkGuideByProductId(req)
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
	} else {
		resp.Code = SUCCESS
		resp.Message = "success"
		resp.Data = data
	}
	return nil
}

// 根据ID查找，返回单条数据
func (h *OpmNetworkGuideHandler) FindDefaultByProductId(ctx context.Context, req *proto.OpmNetworkGuideFilter, resp *proto.OpmNetworkGuideResponse) error {
	s := service.OpmNetworkGuideSvc{Ctx: ctx}
	data, err := s.FindDefaultNetworkGuideByBaseProductId(req)
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
	} else {
		resp.Code = SUCCESS
		resp.Message = "success"
		resp.Data = data
	}
	return nil
}

func (h *OpmNetworkGuideHandler) SetResponse(resp *proto.OpmNetworkGuideResponse, data *proto.OpmNetworkGuide, err error) {
	if err != nil {
		resp.Code = ERROR
		resp.Message = err.Error()
	} else {
		resp.Code = SUCCESS
		resp.Message = "success"
		if data != nil {
			resp.Data = append(resp.Data, data)
		}
	}
}
