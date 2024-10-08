// Code generated by sgen.exe,2022-05-13 10:42:25. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_open_system/model"
	proto "cloud_platform/iot_proto/protos/protosService"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func OpenCompanyAuthLogs_pb2db(src *proto.OpenCompanyAuthLogs) *model.TOpenCompanyAuthLogs {
	if src == nil {
		return nil
	}
	dbObj := model.TOpenCompanyAuthLogs{
		Id:         src.Id,
		CompanyId:  src.CompanyId,
		AuthResult: src.AuthResult,
		AuthName:   src.AuthName,
		AuthDate:   src.AuthDate.AsTime(),
		Why:        src.Why,
		CreatedBy:  src.CreatedBy,
		CreatedAt:  src.CreatedAt.AsTime(),
	}
	return &dbObj
}

func OpenCompanyAuthLogs_db2pb(src *model.TOpenCompanyAuthLogs) *proto.OpenCompanyAuthLogs {
	if src == nil {
		return nil
	}
	pbObj := proto.OpenCompanyAuthLogs{
		Id:         src.Id,
		CompanyId:  src.CompanyId,
		AuthResult: src.AuthResult,
		AuthName:   src.AuthName,
		AuthDate:   timestamppb.New(src.AuthDate),
		Why:        src.Why,
		CreatedBy:  src.CreatedBy,
		CreatedAt:  timestamppb.New(src.CreatedAt),
		UpdatedBy:  src.UpdatedBy,
		UpdatedAt:  timestamppb.New(src.UpdatedAt),
	}
	return &pbObj
}
