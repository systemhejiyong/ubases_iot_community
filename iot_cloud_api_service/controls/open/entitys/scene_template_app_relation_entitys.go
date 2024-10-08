// Code generated by sgen.exe,2022-11-09 16:50:47. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	proto "cloud_platform/iot_proto/protos/protosService"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// 增、删、改及查询返回
type SceneTemplateAppRelationEntitys struct {
	Id              int64     `json:"id,string,omitempty"`
	SceneTemplateId int64     `json:"sceneTemplateId,string,omitempty"`
	AppKey          string    `json:"appKey,omitempty"`
	TenantId        string    `json:"tenantId,omitempty"`
	CreatedBy       int64     `json:"createdBy,string,omitempty"`
	CreatedAt       time.Time `json:"createdAt,omitempty"`
}

// 新增参数非空检查
func (s *SceneTemplateAppRelationEntitys) AddCheck() error {
	return nil
}

// 修改参数非空检查
func (s *SceneTemplateAppRelationEntitys) UpdateCheck() error {
	return nil
}

// 查询参数必填检查
func (*SceneTemplateAppRelationQuery) QueryCheck() error {
	return nil
}

// 查询条件
type SceneTemplateAppRelationQuery struct {
	Page      uint64                          `json:"page,omitempty"`
	Limit     uint64                          `json:"limit,omitempty"`
	Sort      string                          `json:"sort,omitempty"`
	SortField string                          `json:"sortField,omitempty"`
	SearchKey string                          `json:"searchKey,omitempty"`
	Query     *SceneTemplateAppRelationFilter `json:"query,omitempty"`
}

// SceneTemplateAppRelationFilter，查询条件，字段请根据需要自行增减
type SceneTemplateAppRelationFilter struct {
	Id              int64     `json:"id,string,omitempty"`
	SceneTemplateId int64     `json:"sceneTemplateId,string,omitempty"`
	AppKey          string    `json:"appKey,omitempty"`
	TenantId        string    `json:"tenantId,omitempty"`
	CreatedBy       int64     `json:"createdBy,string,omitempty"`
	CreatedAt       time.Time `json:"createdAt,omitempty"`
}

// 实体转pb对象
func SceneTemplateAppRelation_e2pb(src *SceneTemplateAppRelationEntitys) *proto.SceneTemplateAppRelation {
	if src == nil {
		return nil
	}
	pbObj := proto.SceneTemplateAppRelation{
		Id:              src.Id,
		SceneTemplateId: src.SceneTemplateId,
		AppKey:          src.AppKey,
		TenantId:        src.TenantId,
		CreatedBy:       src.CreatedBy,
		CreatedAt:       timestamppb.New(src.CreatedAt),
	}
	return &pbObj
}

// pb对象转实体
func SceneTemplateAppRelation_pb2e(src *proto.SceneTemplateAppRelation) *SceneTemplateAppRelationEntitys {
	if src == nil {
		return nil
	}
	entitysObj := SceneTemplateAppRelationEntitys{
		Id:              src.Id,
		SceneTemplateId: src.SceneTemplateId,
		AppKey:          src.AppKey,
		TenantId:        src.TenantId,
		CreatedBy:       src.CreatedBy,
		CreatedAt:       src.CreatedAt.AsTime(),
	}
	return &entitysObj
}
