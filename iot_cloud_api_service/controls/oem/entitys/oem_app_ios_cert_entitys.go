// Code generated by sgen.exe,2022-06-02 11:15:11. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	proto "cloud_platform/iot_proto/protos/protosService"
)

type OemAppIosCertReq struct {
	AppId            string `json:"appId"`
	DistProvision    string `json:"distProvision"`
	DistCert         string `json:"distCert"`
	DistCertSecret   string `json:"distCertSecret"`
	Version          string `json:"version"`
	DistCertOfficial string `json:"distCertOfficial"` //ios正式证书
}

func (s *OemAppIosCertReq) ToAppIosCertReq(res *OemAppIosCertRes) {
	s.AppId = res.AppId
	s.DistProvision = res.DistProvision
	s.DistCert = res.DistCert
	s.DistCertSecret = res.DistCertSecret
	s.Version = res.Version
	s.DistCertOfficial = res.DistCertOfficial
}

type OemAppIosCertRes struct {
	AppId            string `json:"appId"`
	DistProvision    string `json:"distProvision"`
	DistCert         string `json:"distCert"`
	DistCertSecret   string `json:"distCertSecret"`
	Version          string `json:"version"`
	DistCertOfficial string `json:"distCertOfficial"` //ios正式证书
}

// 增、删、改及查询返回
type OemAppIosCertEntitys struct {
	Id             int64  `json:"id,omitempty"`
	AppId          int64  `json:"appId,omitempty"`
	Version        string `json:"version,omitempty"`
	DistProvision  string `json:"distProvision,omitempty"`
	DistCert       string `json:"distCert,omitempty"`
	DistCertSecret string `json:"distCertSecret,omitempty"`
}

// 新增参数非空检查
func (s *OemAppIosCertEntitys) AddCheck() error {
	return nil
}

// 修改参数非空检查
func (s *OemAppIosCertEntitys) UpdateCheck() error {
	return nil
}

// 查询参数必填检查
func (*OemAppIosCertQuery) QueryCheck() error {
	return nil
}

// 查询条件
type OemAppIosCertQuery struct {
	Page      uint64               `json:"page,omitempty"`
	Limit     uint64               `json:"limit,omitempty"`
	Sort      string               `json:"sort,omitempty"`
	SortField string               `json:"sortField,omitempty"`
	SearchKey string               `json:"searchKey,omitempty"`
	Query     *OemAppIosCertFilter `json:"query,omitempty"`
}

// OemAppIosCertFilter，查询条件，字段请根据需要自行增减
type OemAppIosCertFilter struct {
	Id             int64  `json:"id,omitempty"`
	AppId          int64  `json:"appId,omitempty"`
	Version        string `json:"version,omitempty"`
	DistProvision  string `json:"distProvision,omitempty"`
	DistCert       string `json:"distCert,omitempty"`
	DistCertSecret string `json:"distCertSecret,omitempty"`
}

// 实体转pb对象
func OemAppIosCert_e2pb(src *OemAppIosCertEntitys) *proto.OemAppIosCert {
	if src == nil {
		return nil
	}
	pbObj := proto.OemAppIosCert{
		Id:             src.Id,
		AppId:          src.AppId,
		Version:        src.Version,
		DistProvision:  src.DistProvision,
		DistCert:       src.DistCert,
		DistCertSecret: src.DistCertSecret,
	}
	return &pbObj
}

// pb对象转实体
func OemAppIosCert_pb2e(src *proto.OemAppIosCert) *OemAppIosCertEntitys {
	if src == nil {
		return nil
	}
	entitysObj := OemAppIosCertEntitys{
		Id:             src.Id,
		AppId:          src.AppId,
		Version:        src.Version,
		DistProvision:  src.DistProvision,
		DistCert:       src.DistCert,
		DistCertSecret: src.DistCertSecret,
	}
	return &entitysObj
}
