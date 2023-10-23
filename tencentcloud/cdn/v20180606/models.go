// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20180606

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AccessControl struct {
	// 启用请求头部及请求url访问控制开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 请求头部及请求url访问规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessControlRules []*AccessControlRule `json:"AccessControlRules,omitnil" name:"AccessControlRules"`

	// 返回状态码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnCode *int64 `json:"ReturnCode,omitnil" name:"ReturnCode"`
}

type AccessControlRule struct {
	// requestHeader ：对请求头部进行访问控制
	// url ： 对访问url进行访问控制
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *string `json:"RuleType,omitnil" name:"RuleType"`

	// 封禁内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleContent *string `json:"RuleContent,omitnil" name:"RuleContent"`

	// on ：正则匹配
	// off ：字面匹配
	// 注意：此字段可能返回 null，表示取不到有效值。
	Regex *string `json:"Regex,omitnil" name:"Regex"`

	// RuleType为requestHeader时必填，否则不需要填
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleHeader *string `json:"RuleHeader,omitnil" name:"RuleHeader"`
}

// Predefined struct for user
type AddCLSTopicDomainsRequestParams struct {
	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil" name:"LogsetId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 域名区域配置
	DomainAreaConfigs []*DomainAreaConfig `json:"DomainAreaConfigs,omitnil" name:"DomainAreaConfigs"`

	// 接入渠道，cdn或者ecdn，默认值为cdn
	Channel *string `json:"Channel,omitnil" name:"Channel"`
}

type AddCLSTopicDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil" name:"LogsetId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 域名区域配置
	DomainAreaConfigs []*DomainAreaConfig `json:"DomainAreaConfigs,omitnil" name:"DomainAreaConfigs"`

	// 接入渠道，cdn或者ecdn，默认值为cdn
	Channel *string `json:"Channel,omitnil" name:"Channel"`
}

func (r *AddCLSTopicDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCLSTopicDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "TopicId")
	delete(f, "DomainAreaConfigs")
	delete(f, "Channel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCLSTopicDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCLSTopicDomainsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddCLSTopicDomainsResponse struct {
	*tchttp.BaseResponse
	Response *AddCLSTopicDomainsResponseParams `json:"Response"`
}

func (r *AddCLSTopicDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCLSTopicDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCdnDomainRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 加速域名业务类型
	// web：网页小文件
	// download：下载大文件
	// media：音视频点播
	// hybrid:  动静加速
	// dynamic:  动态加速
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 源站配置
	Origin *Origin `json:"Origin,omitnil" name:"Origin"`

	// 项目 ID，默认为 0，代表【默认项目】
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// IP 黑白名单配置
	IpFilter *IpFilter `json:"IpFilter,omitnil" name:"IpFilter"`

	// IP 限频配置
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitnil" name:"IpFreqLimit"`

	// 状态码缓存配置
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitnil" name:"StatusCodeCache"`

	// 智能压缩配置
	Compression *Compression `json:"Compression,omitnil" name:"Compression"`

	// 带宽封顶配置
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitnil" name:"BandwidthAlert"`

	// Range 回源配置
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitnil" name:"RangeOriginPull"`

	// 301/302 回源跟随配置。
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitnil" name:"FollowRedirect"`

	// 错误码重定向配置（功能灰度中，尚未全量）
	ErrorPage *ErrorPage `json:"ErrorPage,omitnil" name:"ErrorPage"`

	// 请求头部配置
	RequestHeader *RequestHeader `json:"RequestHeader,omitnil" name:"RequestHeader"`

	// 响应头部配置
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitnil" name:"ResponseHeader"`

	// 下载速度配置
	DownstreamCapping *DownstreamCapping `json:"DownstreamCapping,omitnil" name:"DownstreamCapping"`

	// 节点缓存键配置
	CacheKey *CacheKey `json:"CacheKey,omitnil" name:"CacheKey"`

	// 头部缓存配置
	ResponseHeaderCache *ResponseHeaderCache `json:"ResponseHeaderCache,omitnil" name:"ResponseHeaderCache"`

	// 视频拖拽配置
	VideoSeek *VideoSeek `json:"VideoSeek,omitnil" name:"VideoSeek"`

	// 缓存过期时间配置
	Cache *Cache `json:"Cache,omitnil" name:"Cache"`

	// 跨国链路优化配置
	OriginPullOptimization *OriginPullOptimization `json:"OriginPullOptimization,omitnil" name:"OriginPullOptimization"`

	// Https 加速配置
	Https *Https `json:"Https,omitnil" name:"Https"`

	// 时间戳防盗链配置
	Authentication *Authentication `json:"Authentication,omitnil" name:"Authentication"`

	// SEO 优化配置
	Seo *Seo `json:"Seo,omitnil" name:"Seo"`

	// 访问协议强制跳转配置
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitnil" name:"ForceRedirect"`

	// Referer 防盗链配置
	Referer *Referer `json:"Referer,omitnil" name:"Referer"`

	// 浏览器缓存配置（功能灰度中，尚未全量）
	MaxAge *MaxAge `json:"MaxAge,omitnil" name:"MaxAge"`

	// Ipv6 配置（功能灰度中，尚未全量）
	Ipv6 *Ipv6 `json:"Ipv6,omitnil" name:"Ipv6"`

	// 地域属性特殊配置
	// 适用于域名境内加速、境外加速配置不一致场景
	SpecificConfig *SpecificConfig `json:"SpecificConfig,omitnil" name:"SpecificConfig"`

	// 域名加速区域
	// mainland：中国境内加速
	// overseas：中国境外加速
	// global：全球加速
	// 使用中国境外加速、全球加速时，需要先开通中国境外加速服务
	Area *string `json:"Area,omitnil" name:"Area"`

	// 回源超时配置
	OriginPullTimeout *OriginPullTimeout `json:"OriginPullTimeout,omitnil" name:"OriginPullTimeout"`

	// 标签配置
	Tag []*Tag `json:"Tag,omitnil" name:"Tag"`

	// Ipv6 访问配置
	Ipv6Access *Ipv6Access `json:"Ipv6Access,omitnil" name:"Ipv6Access"`

	// 离线缓存
	OfflineCache *OfflineCache `json:"OfflineCache,omitnil" name:"OfflineCache"`

	// Quic访问（收费服务，详见计费说明和产品文档）
	Quic *Quic `json:"Quic,omitnil" name:"Quic"`

	// 回源S3私有鉴权
	AwsPrivateAccess *AwsPrivateAccess `json:"AwsPrivateAccess,omitnil" name:"AwsPrivateAccess"`

	// 回源OSS私有鉴权
	OssPrivateAccess *OssPrivateAccess `json:"OssPrivateAccess,omitnil" name:"OssPrivateAccess"`

	// 华为云对象存储回源鉴权
	HwPrivateAccess *HwPrivateAccess `json:"HwPrivateAccess,omitnil" name:"HwPrivateAccess"`

	// 七牛云对象存储回源鉴权
	QnPrivateAccess *QnPrivateAccess `json:"QnPrivateAccess,omitnil" name:"QnPrivateAccess"`

	// 其他厂商对象存储回源鉴权
	OthersPrivateAccess *OthersPrivateAccess `json:"OthersPrivateAccess,omitnil" name:"OthersPrivateAccess"`

	// HTTPS服务，默认开启（收费服务，详见计费说明和产品文档）
	HttpsBilling *HttpsBilling `json:"HttpsBilling,omitnil" name:"HttpsBilling"`
}

type AddCdnDomainRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 加速域名业务类型
	// web：网页小文件
	// download：下载大文件
	// media：音视频点播
	// hybrid:  动静加速
	// dynamic:  动态加速
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 源站配置
	Origin *Origin `json:"Origin,omitnil" name:"Origin"`

	// 项目 ID，默认为 0，代表【默认项目】
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// IP 黑白名单配置
	IpFilter *IpFilter `json:"IpFilter,omitnil" name:"IpFilter"`

	// IP 限频配置
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitnil" name:"IpFreqLimit"`

	// 状态码缓存配置
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitnil" name:"StatusCodeCache"`

	// 智能压缩配置
	Compression *Compression `json:"Compression,omitnil" name:"Compression"`

	// 带宽封顶配置
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitnil" name:"BandwidthAlert"`

	// Range 回源配置
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitnil" name:"RangeOriginPull"`

	// 301/302 回源跟随配置。
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitnil" name:"FollowRedirect"`

	// 错误码重定向配置（功能灰度中，尚未全量）
	ErrorPage *ErrorPage `json:"ErrorPage,omitnil" name:"ErrorPage"`

	// 请求头部配置
	RequestHeader *RequestHeader `json:"RequestHeader,omitnil" name:"RequestHeader"`

	// 响应头部配置
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitnil" name:"ResponseHeader"`

	// 下载速度配置
	DownstreamCapping *DownstreamCapping `json:"DownstreamCapping,omitnil" name:"DownstreamCapping"`

	// 节点缓存键配置
	CacheKey *CacheKey `json:"CacheKey,omitnil" name:"CacheKey"`

	// 头部缓存配置
	ResponseHeaderCache *ResponseHeaderCache `json:"ResponseHeaderCache,omitnil" name:"ResponseHeaderCache"`

	// 视频拖拽配置
	VideoSeek *VideoSeek `json:"VideoSeek,omitnil" name:"VideoSeek"`

	// 缓存过期时间配置
	Cache *Cache `json:"Cache,omitnil" name:"Cache"`

	// 跨国链路优化配置
	OriginPullOptimization *OriginPullOptimization `json:"OriginPullOptimization,omitnil" name:"OriginPullOptimization"`

	// Https 加速配置
	Https *Https `json:"Https,omitnil" name:"Https"`

	// 时间戳防盗链配置
	Authentication *Authentication `json:"Authentication,omitnil" name:"Authentication"`

	// SEO 优化配置
	Seo *Seo `json:"Seo,omitnil" name:"Seo"`

	// 访问协议强制跳转配置
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitnil" name:"ForceRedirect"`

	// Referer 防盗链配置
	Referer *Referer `json:"Referer,omitnil" name:"Referer"`

	// 浏览器缓存配置（功能灰度中，尚未全量）
	MaxAge *MaxAge `json:"MaxAge,omitnil" name:"MaxAge"`

	// Ipv6 配置（功能灰度中，尚未全量）
	Ipv6 *Ipv6 `json:"Ipv6,omitnil" name:"Ipv6"`

	// 地域属性特殊配置
	// 适用于域名境内加速、境外加速配置不一致场景
	SpecificConfig *SpecificConfig `json:"SpecificConfig,omitnil" name:"SpecificConfig"`

	// 域名加速区域
	// mainland：中国境内加速
	// overseas：中国境外加速
	// global：全球加速
	// 使用中国境外加速、全球加速时，需要先开通中国境外加速服务
	Area *string `json:"Area,omitnil" name:"Area"`

	// 回源超时配置
	OriginPullTimeout *OriginPullTimeout `json:"OriginPullTimeout,omitnil" name:"OriginPullTimeout"`

	// 标签配置
	Tag []*Tag `json:"Tag,omitnil" name:"Tag"`

	// Ipv6 访问配置
	Ipv6Access *Ipv6Access `json:"Ipv6Access,omitnil" name:"Ipv6Access"`

	// 离线缓存
	OfflineCache *OfflineCache `json:"OfflineCache,omitnil" name:"OfflineCache"`

	// Quic访问（收费服务，详见计费说明和产品文档）
	Quic *Quic `json:"Quic,omitnil" name:"Quic"`

	// 回源S3私有鉴权
	AwsPrivateAccess *AwsPrivateAccess `json:"AwsPrivateAccess,omitnil" name:"AwsPrivateAccess"`

	// 回源OSS私有鉴权
	OssPrivateAccess *OssPrivateAccess `json:"OssPrivateAccess,omitnil" name:"OssPrivateAccess"`

	// 华为云对象存储回源鉴权
	HwPrivateAccess *HwPrivateAccess `json:"HwPrivateAccess,omitnil" name:"HwPrivateAccess"`

	// 七牛云对象存储回源鉴权
	QnPrivateAccess *QnPrivateAccess `json:"QnPrivateAccess,omitnil" name:"QnPrivateAccess"`

	// 其他厂商对象存储回源鉴权
	OthersPrivateAccess *OthersPrivateAccess `json:"OthersPrivateAccess,omitnil" name:"OthersPrivateAccess"`

	// HTTPS服务，默认开启（收费服务，详见计费说明和产品文档）
	HttpsBilling *HttpsBilling `json:"HttpsBilling,omitnil" name:"HttpsBilling"`
}

func (r *AddCdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCdnDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "ServiceType")
	delete(f, "Origin")
	delete(f, "ProjectId")
	delete(f, "IpFilter")
	delete(f, "IpFreqLimit")
	delete(f, "StatusCodeCache")
	delete(f, "Compression")
	delete(f, "BandwidthAlert")
	delete(f, "RangeOriginPull")
	delete(f, "FollowRedirect")
	delete(f, "ErrorPage")
	delete(f, "RequestHeader")
	delete(f, "ResponseHeader")
	delete(f, "DownstreamCapping")
	delete(f, "CacheKey")
	delete(f, "ResponseHeaderCache")
	delete(f, "VideoSeek")
	delete(f, "Cache")
	delete(f, "OriginPullOptimization")
	delete(f, "Https")
	delete(f, "Authentication")
	delete(f, "Seo")
	delete(f, "ForceRedirect")
	delete(f, "Referer")
	delete(f, "MaxAge")
	delete(f, "Ipv6")
	delete(f, "SpecificConfig")
	delete(f, "Area")
	delete(f, "OriginPullTimeout")
	delete(f, "Tag")
	delete(f, "Ipv6Access")
	delete(f, "OfflineCache")
	delete(f, "Quic")
	delete(f, "AwsPrivateAccess")
	delete(f, "OssPrivateAccess")
	delete(f, "HwPrivateAccess")
	delete(f, "QnPrivateAccess")
	delete(f, "OthersPrivateAccess")
	delete(f, "HttpsBilling")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCdnDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCdnDomainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddCdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *AddCdnDomainResponseParams `json:"Response"`
}

func (r *AddCdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCdnDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AdvanceCacheRule struct {
	// 规则类型：
	// all：所有文件生效
	// file：指定文件后缀生效
	// directory：指定路径生效
	// path：指定绝对路径生效
	// default：源站未返回 max-age 情况下的缓存规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheType *string `json:"CacheType,omitnil" name:"CacheType"`

	// 对应类型下的匹配内容：
	// all 时填充 *
	// file 时填充后缀名，如 jpg、txt
	// directory 时填充路径，如 /xxx/test/
	// path 时填充绝对路径，如 /xxx/test.html
	// default 时填充 "no max-age"
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheContents []*string `json:"CacheContents,omitnil" name:"CacheContents"`

	// 缓存过期时间
	// 单位为秒，最大可设置为 365 天
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheTime *int64 `json:"CacheTime,omitnil" name:"CacheTime"`
}

type AdvanceConfig struct {
	// 高级配置名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 是否支持高级配置，
	// on：支持
	// off：不支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type AdvanceHttps struct {
	// 自定义Tls数据开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomTlsStatus *string `json:"CustomTlsStatus,omitnil" name:"CustomTlsStatus"`

	// Tls版本列表，支持设置 TLSv1, TLSV1.1, TLSV1.2, TLSv1.3，修改时必须开启连续的版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	TlsVersion []*string `json:"TlsVersion,omitnil" name:"TlsVersion"`

	// 自定义加密套件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cipher *string `json:"Cipher,omitnil" name:"Cipher"`

	// 回源双向校验开启状态
	// off - 关闭校验
	// oneWay - 校验源站
	// twoWay - 双向校验
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyOriginType *string `json:"VerifyOriginType,omitnil" name:"VerifyOriginType"`

	// 回源层证书配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertInfo *ServerCert `json:"CertInfo,omitnil" name:"CertInfo"`

	// 源站证书配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginCertInfo *ClientCert `json:"OriginCertInfo,omitnil" name:"OriginCertInfo"`
}

type AdvancedAuthentication struct {
	// 防盗链配置开关，取值有：
	// on：开启
	// off：关闭
	// 开启时必须且只配置一种模式，其余模式需要设置为 null
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 时间戳防盗链高级版模式A配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeA *AdvancedAuthenticationTypeA `json:"TypeA,omitnil" name:"TypeA"`

	// 时间戳防盗链高级版模式B配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeB *AdvancedAuthenticationTypeB `json:"TypeB,omitnil" name:"TypeB"`

	// 时间戳防盗链高级版模式C配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeC *AdvancedAuthenticationTypeC `json:"TypeC,omitnil" name:"TypeC"`

	// 时间戳防盗链高级版模式D配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeD *AdvancedAuthenticationTypeD `json:"TypeD,omitnil" name:"TypeD"`

	// 时间戳防盗链高级版模式E配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeE *AdvancedAuthenticationTypeE `json:"TypeE,omitnil" name:"TypeE"`

	// 时间戳防盗链高级版模式F配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeF *AdvancedAuthenticationTypeF `json:"TypeF,omitnil" name:"TypeF"`
}

type AdvancedAuthenticationTypeA struct {
	// 用于计算签名的密钥，只允许字母和数字，长度6-32字节。
	SecretKey *string `json:"SecretKey,omitnil" name:"SecretKey"`

	// uri串中签名的字段名，字母，数字或下划线构成，同时必须以字母开头。
	SignParam *string `json:"SignParam,omitnil" name:"SignParam"`

	// uri串中时间的字段名，字母，数字或下划线构成，同时必须以字母开头。
	TimeParam *string `json:"TimeParam,omitnil" name:"TimeParam"`

	// 过期时间，单位秒。
	ExpireTime *int64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 是否必须提供过期时间参数。
	ExpireTimeRequired *bool `json:"ExpireTimeRequired,omitnil" name:"ExpireTimeRequired"`

	// URL 组成格式，如：${private_key}${schema}${host}${full_uri}。
	Format *string `json:"Format,omitnil" name:"Format"`

	// 时间格式，dec，hex分别表示十进制，十六进制。
	TimeFormat *string `json:"TimeFormat,omitnil" name:"TimeFormat"`

	// 鉴权失败时返回的状态码。
	FailCode *int64 `json:"FailCode,omitnil" name:"FailCode"`

	// 链接过期时返回的状态码。
	ExpireCode *int64 `json:"ExpireCode,omitnil" name:"ExpireCode"`

	// 需要鉴权的url路径列表。
	RulePaths []*string `json:"RulePaths,omitnil" name:"RulePaths"`

	// 保留字段。
	Transformation *int64 `json:"Transformation,omitnil" name:"Transformation"`
}

type AdvancedAuthenticationTypeB struct {
	// alpha键名。
	KeyAlpha *string `json:"KeyAlpha,omitnil" name:"KeyAlpha"`

	// beta键名。
	KeyBeta *string `json:"KeyBeta,omitnil" name:"KeyBeta"`

	// gamma键名。
	KeyGamma *string `json:"KeyGamma,omitnil" name:"KeyGamma"`

	// uri串中签名的字段名，字母，数字或下划线构成，同时必须以字母开头。
	SignParam *string `json:"SignParam,omitnil" name:"SignParam"`

	// uri串中时间的字段名，字母，数字或下划线构成，同时必须以字母开头。
	TimeParam *string `json:"TimeParam,omitnil" name:"TimeParam"`

	// 过期时间，单位秒。
	ExpireTime *int64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 时间格式，dec，hex分别表示十进制，十六进制。
	TimeFormat *string `json:"TimeFormat,omitnil" name:"TimeFormat"`

	// 鉴权失败时返回的状态码。
	FailCode *int64 `json:"FailCode,omitnil" name:"FailCode"`

	// 链接过期时返回的状态码。
	ExpireCode *int64 `json:"ExpireCode,omitnil" name:"ExpireCode"`

	// 需要鉴权的url路径列表。
	RulePaths []*string `json:"RulePaths,omitnil" name:"RulePaths"`
}

type AdvancedAuthenticationTypeC struct {
	// 访问密钥。
	AccessKey *string `json:"AccessKey,omitnil" name:"AccessKey"`

	// 鉴权密钥。
	SecretKey *string `json:"SecretKey,omitnil" name:"SecretKey"`
}

type AdvancedAuthenticationTypeD struct {
	// 用于计算签名的密钥，只允许字母和数字，长度6-32字节。
	SecretKey *string `json:"SecretKey,omitnil" name:"SecretKey"`

	// 备份密钥，当使用SecretKey鉴权失败时会使用该密钥重新鉴权。
	BackupSecretKey *string `json:"BackupSecretKey,omitnil" name:"BackupSecretKey"`

	// uri串中签名的字段名，字母，数字或下划线构成，同时必须以字母开头。
	SignParam *string `json:"SignParam,omitnil" name:"SignParam"`

	// uri串中时间的字段名，字母，数字或下划线构成，同时必须以字母开头。
	TimeParam *string `json:"TimeParam,omitnil" name:"TimeParam"`

	// 过期时间，单位秒。
	ExpireTime *int64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 时间格式，dec，hex分别表示十进制，十六进制。
	TimeFormat *string `json:"TimeFormat,omitnil" name:"TimeFormat"`
}

type AdvancedAuthenticationTypeE struct {
	// 用于计算签名的密钥，只允许字母和数字，长度6-32字节。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitnil" name:"SecretKey"`

	// uri串中签名的字段名，字母，数字或下划线构成，同时必须以字母开头。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignParam *string `json:"SignParam,omitnil" name:"SignParam"`

	// uri串中Acl签名的字段名，字母，数字或下划线构成，同时必须以字母开头。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AclSignParam *string `json:"AclSignParam,omitnil" name:"AclSignParam"`

	// uri串中开始时间字段名，字母，数字或下划线构成，同时必须以字母开头。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTimeParam *string `json:"StartTimeParam,omitnil" name:"StartTimeParam"`

	// uri串中过期时间字段名，字母，数字或下划线构成，同时必须以字母开头。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTimeParam *string `json:"ExpireTimeParam,omitnil" name:"ExpireTimeParam"`

	// 时间格式，dec
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeFormat *string `json:"TimeFormat,omitnil" name:"TimeFormat"`
}

type AdvancedAuthenticationTypeF struct {
	// uri串中签名的字段名，字母，数字或下划线构成，同时必须以字母开头。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignParam *string `json:"SignParam,omitnil" name:"SignParam"`

	// uri串中时间的字段名，字母，数字或下划线构成，同时必须以字母开头。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeParam *string `json:"TimeParam,omitnil" name:"TimeParam"`

	// uri串中Transaction字段名，字母，数字或下划线构成，同时必须以字母开头。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransactionParam *string `json:"TransactionParam,omitnil" name:"TransactionParam"`

	// 用于计算签名的主密钥，只允许字母和数字，长度6-32字节。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitnil" name:"SecretKey"`

	// 用于计算签名的备选密钥，主密钥校验失败后再次尝试备选密钥，只允许字母和数字，长度6-32字节。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupSecretKey *string `json:"BackupSecretKey,omitnil" name:"BackupSecretKey"`
}

type AdvancedCCRules struct {
	// 规则名称
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 探测时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectionTime *uint64 `json:"DetectionTime,omitnil" name:"DetectionTime"`

	// 限频阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrequencyLimit *uint64 `json:"FrequencyLimit,omitnil" name:"FrequencyLimit"`

	// IP 惩罚配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	PunishmentSwitch *string `json:"PunishmentSwitch,omitnil" name:"PunishmentSwitch"`

	// IP 惩罚时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	PunishmentTime *uint64 `json:"PunishmentTime,omitnil" name:"PunishmentTime"`

	// 执行动作，intercept|redirect
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitnil" name:"Action"`

	// 动作为 redirect 时，重定向的url
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectUrl *string `json:"RedirectUrl,omitnil" name:"RedirectUrl"`

	// 七层限频具体配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Configure []*ScdnSevenLayerRules `json:"Configure,omitnil" name:"Configure"`

	// 自定义cc规则配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type AdvancedCache struct {
	// 缓存过期规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheRules []*AdvanceCacheRule `json:"CacheRules,omitnil" name:"CacheRules"`

	// 强制缓存配置
	// on：开启
	// off：关闭
	// 开启时，源站返回 no-cache、no-store 头部时，仍按照缓存过期规则进行节点缓存
	// 默认为关闭状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreCacheControl *string `json:"IgnoreCacheControl,omitnil" name:"IgnoreCacheControl"`

	// 当源站返回Set-Cookie头部时，节点是否缓存该头部及body
	// on：开启，不缓存该头部及body
	// off：关闭，遵循用户自定义的节点缓存规则
	// 默认为关闭状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreSetCookie *string `json:"IgnoreSetCookie,omitnil" name:"IgnoreSetCookie"`
}

type AdvancedScdnAclGroup struct {
	// 规则名称
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 具体配置
	Configure []*AdvancedScdnAclRule `json:"Configure,omitnil" name:"Configure"`

	// 执行动作，intercept|redirect
	Result *string `json:"Result,omitnil" name:"Result"`

	// 规则是否生效，active|inactive
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误页面配置
	ErrorPage *ScdnErrorPage `json:"ErrorPage,omitnil" name:"ErrorPage"`
}

type AdvancedScdnAclRule struct {
	// 匹配关键字，可取值有：
	// <li>protocol：HTTP协议</li>
	// <li>httpVersion：HTTP版本</li>
	// <li>method：请求方法</li>
	// <li>ip：请求源IP</li>
	// <li>ipAsn：请求源IP自治域号</li>
	// <li>ipCountry：请求源IP所在国家</li>
	// <li>ipArea：请求源IP所在大区</li>
	// <li>xForwardFor：请求头X-Forwarded-For</li>
	// <li>directory：路径</li>
	// <li>index：首页</li>
	// <li>path：文件全路径</li>
	// <li>file：文件扩展名</li>
	// <li>param：请求参数</li>
	// <li>referer：请求头Referer</li>
	// <li>cookie：请求头Cookie</li>
	// <li>userAgent：请求头User-Agent</li>
	// <li>head：自定义请求头</li>
	MatchKey *string `json:"MatchKey,omitnil" name:"MatchKey"`

	// 逻辑操作符，取值如下：
	// 不包含：exclude
	// 包含：include
	// 不等于：notequal
	// 等于：equal
	// 前缀匹配：matching
	// 内容为空或不存在：null
	LogicOperator *string `json:"LogicOperator,omitnil" name:"LogicOperator"`

	// 匹配值。
	// 当MatchKey为protocol时
	// 取值：HTTP、HTTPS
	// 
	// 当MatchKey为httpVersion时
	// 取值：HTTP/1.0、HTTP/1.1、HTTP/1.2、HTTP/2、HTTP/3
	// 
	// 当MatchKey为method时
	// 取值：HEAD、GET、POST、PUT、OPTIONS、TRACE、DELETE、PATCH、CONNECT
	// 
	// 当MatchKey为ipCountry时，取值为：
	// 其他：OTHER
	// 委内瑞拉：VE
	// 乌拉圭：UY
	// 苏里南：SR
	// 巴拉圭：PY
	// 秘鲁：PE
	// 圭亚那：GY
	// 厄瓜多尔：EC
	// 哥伦比亚：CO
	// 智利：CL
	// 巴西：BR
	// 玻利维亚：BO
	// 阿根廷：AR
	// 新西兰：NZ
	// 萨摩亚：WS
	// 瓦努阿图：VU
	// 图瓦卢：TV
	// 汤加：TO
	// 托克劳：TK
	// 帕劳：PW
	// 纽埃：NU
	// 瑙鲁：NR
	// 基里巴斯：KI
	// 关岛：GU
	// 密克罗尼西亚：FM
	// 澳大利亚：AU
	// 美国：US
	// 波多黎各：PR
	// 多米尼加共和国：DO
	// 哥斯达黎加：CR
	// 东萨摩亚：AS
	// 安提瓜和巴布达：AG
	// 巴拿马：PA
	// 尼加拉瓜：NI
	// 墨西哥：MX
	// 牙买加：JM
	// 海地：HT
	// 洪都拉斯：HN
	// 危地马拉：GT
	// 瓜德罗普岛：GP
	// 格陵兰：GL
	// 格林纳达：GD
	// 古巴：CU
	// 加拿大：CA
	// 伯利兹：BZ
	// 巴哈马：BS
	// 百慕大：BM
	// 巴巴多斯：BB
	// 阿鲁巴：AW
	// 安圭拉：AI
	// 梵蒂冈：VA
	// 斯洛伐克：SK
	// 英国：GB
	// 捷克共和国：CZ
	// 乌克兰：UA
	// 土耳其：TR
	// 斯洛文尼亚：SI
	// 瑞典：SE
	// 塞尔维亚：RS
	// 罗马尼亚：RO
	// 葡萄牙：PT
	// 波兰：PL
	// 挪威：NO
	// 荷兰：NL
	// 马耳他：MT
	// 马其顿：MK
	// 黑山：ME
	// 摩尔多瓦：MD
	// 摩纳哥：MC
	// 拉脱维亚：LV
	// 卢森堡：LU
	// 立陶宛：LT
	// 列支敦士登：LI
	// 哈萨克斯坦：KZ
	// 意大利：IT
	// 冰岛：IS
	// 爱尔兰：IE
	// 匈牙利：HU
	// 克罗地亚：HR
	// 希腊：GR
	// 直布罗陀：GI
	// 根西岛：GG
	// 格鲁吉亚：GE
	// 法国：FR
	// 芬兰：FI
	// 西班牙：ES
	// 爱沙尼亚：EE
	// 丹麦：DK
	// 德国：DE
	// 塞浦路斯：CY
	// 瑞士：CH
	// 白俄罗斯：BY
	// 保加利亚：BG
	// 比利时：BE
	// 阿塞拜疆：AZ
	// 奥地利：AT
	// 亚美尼亚：AM
	// 阿尔巴尼亚：AL
	// 安道尔：AD
	// 东帝汶：TL
	// 叙利亚：SY
	// 沙特阿拉伯：SA
	// 巴勒斯坦：PS
	// 斯里兰卡：LK
	// 斯里兰卡：LK
	// 朝鲜：KP
	// 吉尔吉斯斯坦：KG
	// 中国香港：HK
	// 文莱：BN
	// 孟加拉：BD
	// 阿联酋：AE
	// 也门：YE
	// 越南：VN
	// 乌兹别克斯坦：UZ
	// 中国台湾：TW
	// 土库曼斯坦：TM
	// 塔吉克斯坦：TJ
	// 泰国：TH
	// 新加坡：SG
	// 卡塔尔：QA
	// 巴基斯坦：PK
	// 菲律宾：PH
	// 阿曼：OM
	// 尼泊尔：NP
	// 马来西亚：MY
	// 马尔代夫：MV
	// 中国澳门：MO
	// 蒙古：MN
	// 缅甸：MM
	// 黎巴嫩：LB
	// 科威特：KW
	// 韩国：KR
	// 柬埔寨：KH
	// 日本：JP
	// 约旦：JO
	// 伊朗：IR
	// 伊拉克：IQ
	// 印度：IN
	// 以色列：IL
	// 印度尼西亚：ID
	// 中国：CN
	// 不丹：BT
	// 巴林：BH
	// 阿富汗：AF
	// 利比亚：LY
	// 刚果金：CG
	// 留尼汪岛：RE
	// 斯威士兰：SZ
	// 津巴布韦：ZW
	// 赞比亚：ZM
	// 马约特：YT
	// 乌干达：UG
	// 坦桑尼亚：TZ
	// 突尼斯：TN
	// 多哥：TG
	// 乍得：TD
	// 索马里：SO
	// 塞内加尔：SN
	// 苏丹：SD
	// 塞舌尔：SC
	// 卢旺达：RW
	// 尼日利亚：NG
	// 尼日尔：NE
	// 纳米比亚：NA
	// 莫桑比克：MZ
	// 马拉维：MW
	// 毛里求斯：MU
	// 毛里塔尼亚：MR
	// 马里：ML
	// 马达加斯加：MG
	// 摩洛哥：MA
	// 莱索托：LS
	// 利比里亚：LR
	// 科摩罗：KM
	// 肯尼亚：KE
	// 几内亚：GN
	// 冈比亚：GM
	// 加纳：GH
	// 加蓬：GA
	// 埃塞俄比亚：ET
	// 厄立特里亚：ER
	// 埃及：EG
	// 阿尔及利亚：DZ
	// 吉布提：DJ
	// 喀麦隆：CM
	// 刚果：CG
	// 博茨瓦纳：BW
	// 贝宁：BJ
	// 布隆迪：BI
	// 安哥拉：AO
	// 
	// 当MatchKey为ipArea时，取值为：
	// 其他：OTHER
	// 亚洲：AS
	// 欧洲：EU
	// 南极洲：AN
	// 非洲：AF
	// 大洋洲：OC
	// 北美洲：NA
	// 南美洲：SA
	// 
	// 当MatchKey为index时
	// 取值为：/;/index.html
	MatchValue []*string `json:"MatchValue,omitnil" name:"MatchValue"`

	// 是否区分大小写 true：区分 false：不区分
	CaseSensitive *bool `json:"CaseSensitive,omitnil" name:"CaseSensitive"`

	// 当MatchKey为param时必填：表示请求参数Key 当MatchKey为cookie时必填：表示请求头Cookie中参数的
	MatchKeyParam *string `json:"MatchKeyParam,omitnil" name:"MatchKeyParam"`
}

type Authentication struct {
	// 防盗链配置开关，取值有：
	// on：开启
	// off：关闭
	// 开启时必须且只配置一种模式，其余模式需要设置为 null
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 鉴权算法，取值有：
	// md5：按MD5算法取hash值
	// sha256：按SHA-256算法取hash值
	// 默认为 md5
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthAlgorithm *string `json:"AuthAlgorithm,omitnil" name:"AuthAlgorithm"`

	// 时间戳防盗链模式 A 配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeA *AuthenticationTypeA `json:"TypeA,omitnil" name:"TypeA"`

	// 时间戳防盗链模式 B 配置（模式 B 后台升级中，暂时不支持配置）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeB *AuthenticationTypeB `json:"TypeB,omitnil" name:"TypeB"`

	// 时间戳防盗链模式 C 配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeC *AuthenticationTypeC `json:"TypeC,omitnil" name:"TypeC"`

	// 时间戳防盗链模式 D 配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeD *AuthenticationTypeD `json:"TypeD,omitnil" name:"TypeD"`
}

type AuthenticationTypeA struct {
	// 计算签名的密钥
	// 仅允许大小写字母与数字，长度 6~32 位
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitnil" name:"SecretKey"`

	// 签名参数名设置
	// 仅允许大小写字母、数字或下划线，长度 1~100 位，不能以数字开头
	SignParam *string `json:"SignParam,omitnil" name:"SignParam"`

	// 签名过期时间设置
	// 单位为秒，最大可设置为 630720000
	ExpireTime *int64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 鉴权/不做鉴权的文件扩展名列表设置
	// 如果包含字符 *  则表示所有文件
	FileExtensions []*string `json:"FileExtensions,omitnil" name:"FileExtensions"`

	// whitelist：白名单，表示对除了 FileExtensions 列表之外的所有类型进行鉴权
	// blacklist：黑名单，表示仅对 FileExtensions 中的类型进行鉴权
	FilterType *string `json:"FilterType,omitnil" name:"FilterType"`

	// 计算签名的备用密钥
	// 仅允许大小写字母与数字，长度 6~32 位
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupSecretKey *string `json:"BackupSecretKey,omitnil" name:"BackupSecretKey"`
}

type AuthenticationTypeB struct {
	// 计算签名的密钥
	// 仅允许大小写字母与数字，长度 6~32 位
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitnil" name:"SecretKey"`

	// 签名过期时间设置
	// 单位为秒，最大可设置为 630720000
	ExpireTime *int64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 鉴权/不做鉴权的文件扩展名列表设置
	// 如果包含字符 *  则表示所有文件
	FileExtensions []*string `json:"FileExtensions,omitnil" name:"FileExtensions"`

	// whitelist：白名单，表示对除了 FileExtensions 列表之外的所有类型进行鉴权
	// blacklist：黑名单，表示仅对 FileExtensions 中的类型进行鉴权
	FilterType *string `json:"FilterType,omitnil" name:"FilterType"`

	// 计算签名的备用密钥
	// 仅允许大小写字母与数字，长度 6~32 位
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupSecretKey *string `json:"BackupSecretKey,omitnil" name:"BackupSecretKey"`
}

type AuthenticationTypeC struct {
	// 计算签名的密钥
	// 仅允许大小写字母与数字，长度 6~32 位
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitnil" name:"SecretKey"`

	// 签名过期时间设置
	// 单位为秒，最大可设置为 630720000
	ExpireTime *int64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 鉴权/不做鉴权的文件扩展名列表设置
	// 如果包含字符 *  则表示所有文件
	FileExtensions []*string `json:"FileExtensions,omitnil" name:"FileExtensions"`

	// whitelist：白名单，表示对除了 FileExtensions 列表之外的所有类型进行鉴权
	// blacklist：黑名单，表示仅对 FileExtensions 中的类型进行鉴权
	FilterType *string `json:"FilterType,omitnil" name:"FilterType"`

	// 时间戳进制设置
	// dec：十进制
	// hex：十六进制
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeFormat *string `json:"TimeFormat,omitnil" name:"TimeFormat"`

	// 计算签名的备用密钥
	// 仅允许大小写字母与数字，长度 6~32 位
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupSecretKey *string `json:"BackupSecretKey,omitnil" name:"BackupSecretKey"`
}

type AuthenticationTypeD struct {
	// 计算签名的密钥
	// 仅允许大小写字母与数字，长度 6~32 位
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitnil" name:"SecretKey"`

	// 签名过期时间设置
	// 单位为秒，最大可设置为 630720000
	ExpireTime *int64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 鉴权/不做鉴权的文件扩展名列表设置
	// 如果包含字符 *  则表示所有文件
	FileExtensions []*string `json:"FileExtensions,omitnil" name:"FileExtensions"`

	// whitelist：白名单，表示对除了 FileExtensions 列表之外的所有类型进行鉴权
	// blacklist：黑名单，表示仅对 FileExtensions 中的类型进行鉴权
	FilterType *string `json:"FilterType,omitnil" name:"FilterType"`

	// 签名参数名设置
	// 仅允许大小写字母、数字或下划线，长度 1~100 位，不能以数字开头
	SignParam *string `json:"SignParam,omitnil" name:"SignParam"`

	// 时间戳参数名设置
	// 仅允许大小写字母、数字或下划线，长度 1~100 位，不能以数字开头
	TimeParam *string `json:"TimeParam,omitnil" name:"TimeParam"`

	// 时间戳进制设置
	// dec：十进制
	// hex：十六进制
	TimeFormat *string `json:"TimeFormat,omitnil" name:"TimeFormat"`

	// 计算签名的备用密钥
	// 仅允许大小写字母与数字，长度 6~32 位
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupSecretKey *string `json:"BackupSecretKey,omitnil" name:"BackupSecretKey"`
}

type AvifAdapter struct {
	// 图片优化AvifAdapter配置项开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type AwsPrivateAccess struct {
	// s3源站回源鉴权配置项开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 访问ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessKey *string `json:"AccessKey,omitnil" name:"AccessKey"`

	// 密钥。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitnil" name:"SecretKey"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// Bucketname
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitnil" name:"Bucket"`
}

type BandwidthAlert struct {
	// 用量封顶配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 用量封顶阈值，带宽单位为bps，流量单位byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	BpsThreshold *int64 `json:"BpsThreshold,omitnil" name:"BpsThreshold"`

	// 达到阈值后的操作
	// RETURN_404：全部请求返回 404
	// 注意：此字段可能返回 null，表示取不到有效值。
	CounterMeasure *string `json:"CounterMeasure,omitnil" name:"CounterMeasure"`

	// 境内区域上次触发用量封顶阈值的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTriggerTime *string `json:"LastTriggerTime,omitnil" name:"LastTriggerTime"`

	// 用量封顶提醒配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertSwitch *string `json:"AlertSwitch,omitnil" name:"AlertSwitch"`

	// 用量封顶阈值提醒百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertPercentage *int64 `json:"AlertPercentage,omitnil" name:"AlertPercentage"`

	// 海外区域上次触发用量封顶阈值的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTriggerTimeOverseas *string `json:"LastTriggerTimeOverseas,omitnil" name:"LastTriggerTimeOverseas"`

	// 用量阈值触发的维度
	// 带宽：bandwidth
	// 流量：flux
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 累计用量配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatisticItems []*StatisticItem `json:"StatisticItems,omitnil" name:"StatisticItems"`
}

type BotCookie struct {
	// Bot cookie策略配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 规则类型，当前只有all
	RuleType *string `json:"RuleType,omitnil" name:"RuleType"`

	// 规则值，['*']
	RuleValue []*string `json:"RuleValue,omitnil" name:"RuleValue"`

	// 执行动作，monitor|intercept|redirect|captcha
	Action *string `json:"Action,omitnil" name:"Action"`

	// 重定向时设置的重定向页面
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectUrl *string `json:"RedirectUrl,omitnil" name:"RedirectUrl"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type BotJavaScript struct {
	// Bot js策略配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 规则类型，当前只有file
	RuleType *string `json:"RuleType,omitnil" name:"RuleType"`

	// 规则值，['html', 'htm']
	RuleValue []*string `json:"RuleValue,omitnil" name:"RuleValue"`

	// 执行动作，monitor|intercept|redirect|captcha
	Action *string `json:"Action,omitnil" name:"Action"`

	// 重定向时设置的重定向页面
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectUrl *string `json:"RedirectUrl,omitnil" name:"RedirectUrl"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type BotRecord struct {
	// 动作，取值为以为3个类型中的一个："intercept","permit","monitor"，分别表示： 拦截， 放行，监控
	Action *string `json:"Action,omitnil" name:"Action"`

	// 会话总次数
	Nums *int64 `json:"Nums,omitnil" name:"Nums"`

	// BotType=UB时，表示预测标签，取值如下：
	//                 "crawler_unregular",
	//                 "crawler_regular",
	//                 "request_repeat",
	//                 "credential_miss_user",
	//                 "credential_without_user",
	//                 "credential_only_action",
	//                 "credential_user_password",
	//                 "credential_cracking",
	//                 "credential_stuffing",
	//                 "brush_sms",
	//                 "brush_captcha",
	//                 "reg_malicious"
	// BotType=TCB时，表示Bot分类，取值如下：
	//                 "Uncategorised",
	//                 "Search engine bot",
	//                 "Site monitor",
	//                 "Screenshot creator",
	//                 "Link checker",
	//                 "Web scraper",
	//                 "Vulnerability scanner",
	//                 "Virus scanner",
	//                 "Speed tester",
	//                 "Feed Fetcher",
	//                 "Tool",
	//                 "Marketing"
	// BotType=UCB时，为二期接口，暂时未定义内容
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 会话持续时间
	SessionDuration *float64 `json:"SessionDuration,omitnil" name:"SessionDuration"`

	// 访问源IP
	SrcIp *string `json:"SrcIp,omitnil" name:"SrcIp"`

	// 异常特征
	BotFeature []*string `json:"BotFeature,omitnil" name:"BotFeature"`

	// 最新检测时间
	Time *string `json:"Time,omitnil" name:"Time"`

	// BOT得分
	Score *int64 `json:"Score,omitnil" name:"Score"`

	// 平均速率
	AvgSpeed *float64 `json:"AvgSpeed,omitnil" name:"AvgSpeed"`

	// BotType=TCB，表示TCB名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TcbDetail *string `json:"TcbDetail,omitnil" name:"TcbDetail"`

	// BOT记录唯一ID，用于查询访问详情
	Id *string `json:"Id,omitnil" name:"Id"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type BotSortBy struct {
	// 排序参数名称， 取值为：timestamp， nums， session_duration，score.total，stat.avg_speed分别表示按照：最新检测时间，会话总次数，会话持续时间，BOT得分，平均速率排序
	Key *string `json:"Key,omitnil" name:"Key"`

	// asc/desc
	Sequence *string `json:"Sequence,omitnil" name:"Sequence"`
}

type BotStatisticsCount struct {
	// BOT次数
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// Top指标值,如果是ip维度就是ip如果是session维度就是域名
	Value *string `json:"Value,omitnil" name:"Value"`

	// ip所在国家
	Country *string `json:"Country,omitnil" name:"Country"`

	// ip所在省份
	Province *string `json:"Province,omitnil" name:"Province"`

	// ip归属的idc
	Isp *string `json:"Isp,omitnil" name:"Isp"`
}

type BotStats struct {
	// 指标名称
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 指标详细数据
	DetailData []*BotStatsDetailData `json:"DetailData,omitnil" name:"DetailData"`
}

type BotStatsDetailData struct {
	// 时间
	Time *string `json:"Time,omitnil" name:"Time"`

	// 数据值
	Value *int64 `json:"Value,omitnil" name:"Value"`
}

type BriefDomain struct {
	// 域名 ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 腾讯云账号 ID
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// 加速域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名对应的 CNAME 地址
	Cname *string `json:"Cname,omitnil" name:"Cname"`

	// 加速服务状态
	// rejected：域名审核未通过，域名备案过期/被注销导致
	// processing：部署中
	// closing：关闭中
	// online：已启动
	// offline：已关闭
	Status *string `json:"Status,omitnil" name:"Status"`

	// 项目 ID，可前往腾讯云项目管理页面查看
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 域名业务类型
	// web：静态加速
	// download：下载加速
	// media：流媒体点播加速
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 域名创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 域名更新时间
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 源站配置详情
	Origin *Origin `json:"Origin,omitnil" name:"Origin"`

	// 域名封禁状态
	// normal：正常状态
	// overdue：账号欠费导致域名关闭，充值完成后可自行启动加速服务
	// malicious：域名出现恶意行为，强制关闭加速服务
	// ddos：域名被大规模 DDoS 攻击，关闭加速服务
	// idle：域名超过 90 天内无任何操作、数据产生，判定为不活跃域名自动关闭加速服务，可自行启动加速服务
	// unlicensed：域名未备案/备案注销，自动关闭加速服务，备案完成后可自行启动加速服务
	// capping：触发配置的带宽阈值上限
	// readonly：域名存在特殊配置，被锁定
	Disable *string `json:"Disable,omitnil" name:"Disable"`

	// 加速区域
	// mainland：中国境内加速
	// overseas：中国境外加速
	// global：全球加速
	Area *string `json:"Area,omitnil" name:"Area"`

	// 域名锁定状态
	// normal：未锁定
	// mainland：中国境内锁定
	// overseas：中国境外锁定
	// global：全球锁定
	Readonly *string `json:"Readonly,omitnil" name:"Readonly"`

	// 域名所属产品，cdn/ecdn
	Product *string `json:"Product,omitnil" name:"Product"`

	// 主域名
	ParentHost *string `json:"ParentHost,omitnil" name:"ParentHost"`
}

type Cache struct {
	// 基础缓存过期时间配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	SimpleCache *SimpleCache `json:"SimpleCache,omitnil" name:"SimpleCache"`

	// 高级缓存过期时间配置（已弃用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedCache *AdvancedCache `json:"AdvancedCache,omitnil" name:"AdvancedCache"`

	// 高级路径缓存配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleCache []*RuleCache `json:"RuleCache,omitnil" name:"RuleCache"`
}

type CacheConfig struct {
	// 启发式自定义时间缓存配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeuristicCacheTimeSwitch *string `json:"HeuristicCacheTimeSwitch,omitnil" name:"HeuristicCacheTimeSwitch"`

	// 单位 秒.
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeuristicCacheTime *int64 `json:"HeuristicCacheTime,omitnil" name:"HeuristicCacheTime"`
}

type CacheConfigCache struct {
	// 路径缓存配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 缓存过期时间设置
	// 单位为秒，最大可设置为 365 天
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheTime *int64 `json:"CacheTime,omitnil" name:"CacheTime"`

	// 高级缓存过期配置，开启时会对比源站返回的 max-age 值与 CacheRules 中设置的缓存过期时间，取最小值在节点进行缓存
	// on：开启
	// off：关闭
	// 默认为关闭状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareMaxAge *string `json:"CompareMaxAge,omitnil" name:"CompareMaxAge"`

	// 强制缓存
	// on：开启
	// off：关闭
	// 默认为关闭状态，开启后，源站返回的 no-store、no-cache 资源，也将按照 CacheRules 规则进行缓存
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreCacheControl *string `json:"IgnoreCacheControl,omitnil" name:"IgnoreCacheControl"`

	// 当源站返回Set-Cookie头部时，节点是否缓存该头部及body
	// on：开启，不缓存该头部及body
	// off：关闭，遵循用户自定义的节点缓存规则
	// 默认为关闭状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreSetCookie *string `json:"IgnoreSetCookie,omitnil" name:"IgnoreSetCookie"`
}

type CacheConfigFollowOrigin struct {
	// 路径缓存遵循源站配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 启发式缓存配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeuristicCache *HeuristicCache `json:"HeuristicCache,omitnil" name:"HeuristicCache"`
}

type CacheConfigNoCache struct {
	// 路径缓存不缓存配置配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 总是回源站校验
	// on：开启
	// off：关闭
	// 默认为关闭状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Revalidate *string `json:"Revalidate,omitnil" name:"Revalidate"`
}

type CacheKey struct {
	// 是否开启全路径缓存
	// on：开启全路径缓存（即关闭参数忽略）
	// off：关闭全路径缓存（即开启参数忽略）
	FullUrlCache *string `json:"FullUrlCache,omitnil" name:"FullUrlCache"`

	// 是否忽略大小写缓存
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreCase *string `json:"IgnoreCase,omitnil" name:"IgnoreCase"`

	// CacheKey中包含请求参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryString *QueryStringKey `json:"QueryString,omitnil" name:"QueryString"`

	// CacheKey中包含Cookie
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cookie *CookieKey `json:"Cookie,omitnil" name:"Cookie"`

	// CacheKey中包含请求头部
	// 注意：此字段可能返回 null，表示取不到有效值。
	Header *HeaderKey `json:"Header,omitnil" name:"Header"`

	// CacheKey中包含自定义字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheTag *CacheTagKey `json:"CacheTag,omitnil" name:"CacheTag"`

	// CacheKey中包含请求协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scheme *SchemeKey `json:"Scheme,omitnil" name:"Scheme"`

	// 分路径缓存键配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyRules []*KeyRule `json:"KeyRules,omitnil" name:"KeyRules"`
}

type CacheOptResult struct {
	// 成功的url列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessUrls []*string `json:"SuccessUrls,omitnil" name:"SuccessUrls"`

	// 失败的url列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailUrls []*string `json:"FailUrls,omitnil" name:"FailUrls"`
}

type CacheTagKey struct {
	// 使用CacheTag作为CacheKey的一部分配置开关，取值有
	// on：开启，使用CacheTag作为CacheKey的一部分
	// off：关闭，不使用CacheTag作为CacheKey的一部分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 自定义CacheTag的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type CappingRule struct {
	// 规则类型：
	// all：所有文件生效
	// file：指定文件后缀生效
	// directory：指定路径生效
	// path：指定绝对路径生效
	RuleType *string `json:"RuleType,omitnil" name:"RuleType"`

	// RuleType 对应类型下的匹配内容： 
	// all 时填充 *
	// file 时填充后缀名，如 jpg、txt
	// directory 时填充路径，如 /xxx/test/
	// path 时填充绝对路径，如 /xxx/test.html
	RulePaths []*string `json:"RulePaths,omitnil" name:"RulePaths"`

	// 下行速度值设置，单位为 KB/s
	KBpsThreshold *int64 `json:"KBpsThreshold,omitnil" name:"KBpsThreshold"`
}

type CcTopData struct {
	// 客户端Ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 访问URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 客户端UserAgent
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAgent *string `json:"UserAgent,omitnil" name:"UserAgent"`

	// 请求数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *uint64 `json:"Value,omitnil" name:"Value"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type CdnData struct {
	// 查询指定的指标名称：
	// flux：流量，单位为 byte
	// bandwidth：带宽，单位为 bps
	// request：请求数，单位为 次
	// fluxHitRate：流量命中率，单位为 %
	// statusCode：状态码，返回 2XX、3XX、4XX、5XX 汇总数据，单位为 个
	// 2XX：返回 2XX 状态码汇总及各 2 开头状态码数据，单位为 个
	// 3XX：返回 3XX 状态码汇总及各 3 开头状态码数据，单位为 个
	// 4XX：返回 4XX 状态码汇总及各 4 开头状态码数据，单位为 个
	// 5XX：返回 5XX 状态码汇总及各 5 开头状态码数据，单位为 个
	// 或指定查询的某一具体状态码
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 明细数据组合
	DetailData []*TimestampData `json:"DetailData,omitnil" name:"DetailData"`

	// 汇总数据组合
	SummarizedData *SummarizedData `json:"SummarizedData,omitnil" name:"SummarizedData"`
}

type CdnIp struct {
	// 指定查询的 IP
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// IP 归属：
	// yes：节点归属于腾讯云 CDN
	// no：节点不属于腾讯云 CDN
	Platform *string `json:"Platform,omitnil" name:"Platform"`

	// 节点所处的省份/国家
	// unknown 表示节点位置未知
	Location *string `json:"Location,omitnil" name:"Location"`

	// 节点上下线历史记录
	History []*CdnIpHistory `json:"History,omitnil" name:"History"`

	// 节点的所在区域
	// mainland：中国境内加速节点
	// overseas：中国境外加速节点
	// unknown：服务地域无法获取
	Area *string `json:"Area,omitnil" name:"Area"`

	// 节点的所在城市
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *string `json:"City,omitnil" name:"City"`
}

type CdnIpHistory struct {
	// 操作类型
	// online：节点上线
	// offline：节点下线
	Status *string `json:"Status,omitnil" name:"Status"`

	// 操作类型对应的操作时间
	// 当该值为 null 时表示无历史状态变更记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Datetime *string `json:"Datetime,omitnil" name:"Datetime"`
}

type ClientCert struct {
	// 客户端证书
	// PEM 格式，需要进行 Base 64 编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Certificate *string `json:"Certificate,omitnil" name:"Certificate"`

	// 客户端证书名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertName *string `json:"CertName,omitnil" name:"CertName"`

	// 证书过期时间
	// 作为入参时无需填充
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 证书颁发时间
	// 作为入参时无需填充
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployTime *string `json:"DeployTime,omitnil" name:"DeployTime"`
}

type ClientInfo struct {
	// 省份。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProvName *string `json:"ProvName,omitnil" name:"ProvName"`

	// 国家。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Country *string `json:"Country,omitnil" name:"Country"`

	// 运营商。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IspName *string `json:"IspName,omitnil" name:"IspName"`

	// 客户端IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil" name:"Ip"`
}

type ClsLogIpData struct {
	// IP
	ClientIp *string `json:"ClientIp,omitnil" name:"ClientIp"`

	// 在给定的时间段中，1秒内的最大请求量
	Request *uint64 `json:"Request,omitnil" name:"Request"`

	// 在获取的Top信息中，IP出现的次数
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 在给定的时间段中，1秒内的最大请求量对应的时间
	Time *string `json:"Time,omitnil" name:"Time"`
}

type ClsLogObject struct {
	// 主题ID
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 主题名字
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// 日志时间
	Timestamp *string `json:"Timestamp,omitnil" name:"Timestamp"`

	// 日志内容
	Content *string `json:"Content,omitnil" name:"Content"`

	// 采集路径
	Filename *string `json:"Filename,omitnil" name:"Filename"`

	// 日志来源设备
	Source *string `json:"Source,omitnil" name:"Source"`
}

type ClsSearchLogs struct {
	// 获取更多检索结果的游标
	Context *string `json:"Context,omitnil" name:"Context"`

	// 搜索结果是否已经全部返回
	Listover *bool `json:"Listover,omitnil" name:"Listover"`

	// 日志内容信息
	Results []*ClsLogObject `json:"Results,omitnil" name:"Results"`
}

type Compatibility struct {
	// 兼容标志状态码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *int64 `json:"Code,omitnil" name:"Code"`
}

type Compression struct {
	// 智能压缩配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 压缩规则数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompressionRules []*CompressionRule `json:"CompressionRules,omitnil" name:"CompressionRules"`
}

type CompressionRule struct {
	// true：需要设置为 ture，启用压缩
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compress *bool `json:"Compress,omitnil" name:"Compress"`

	// 触发压缩的文件长度最小值，单位为字节数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinLength *int64 `json:"MinLength,omitnil" name:"MinLength"`

	// 触发压缩的文件长度最大值，单位为字节数
	// 最大可设置为 30MB
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxLength *int64 `json:"MaxLength,omitnil" name:"MaxLength"`

	// 文件压缩算法
	// gzip：指定 GZIP 压缩
	// brotli：指定Brotli压缩
	// 注意：此字段可能返回 null，表示取不到有效值。
	Algorithms []*string `json:"Algorithms,omitnil" name:"Algorithms"`

	// 根据文件后缀类型压缩
	// 例如 jpg、txt
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileExtensions []*string `json:"FileExtensions,omitnil" name:"FileExtensions"`

	// 规则类型：
	// all：所有文件生效
	// file：指定文件后缀生效
	// directory：指定路径生效
	// path：指定绝对路径生效
	// contentType：指定Content-Type头为特定值时生效
	// 当指定了此字段时，FileExtensions字段不生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *string `json:"RuleType,omitnil" name:"RuleType"`

	// CacheType 对应类型下的匹配内容：
	// all 时填充 *
	// file 时填充后缀名，如 jpg、txt
	// directory 时填充路径，如 /xxx/test
	// path 时填充绝对路径，如 /xxx/test.html
	// contentType 时填充 text/html
	// 注意：此字段可能返回 null，表示取不到有效值。
	RulePaths []*string `json:"RulePaths,omitnil" name:"RulePaths"`
}

type CookieKey struct {
	// 使用Cookie作为Cache的一部分配置开关，取值有：
	// on：开启，使用Cookie作为Cache的一部分
	// off：关闭，不使用Cookie作为Cache的一部分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 使用的cookie，';' 分割
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

// Predefined struct for user
type CreateClsLogTopicRequestParams struct {
	// 日志主题名称
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil" name:"LogsetId"`

	// 接入渠道，cdn或者ecdn，默认值为cdn
	Channel *string `json:"Channel,omitnil" name:"Channel"`

	// 域名区域信息
	DomainAreaConfigs []*DomainAreaConfig `json:"DomainAreaConfigs,omitnil" name:"DomainAreaConfigs"`
}

type CreateClsLogTopicRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题名称
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil" name:"LogsetId"`

	// 接入渠道，cdn或者ecdn，默认值为cdn
	Channel *string `json:"Channel,omitnil" name:"Channel"`

	// 域名区域信息
	DomainAreaConfigs []*DomainAreaConfig `json:"DomainAreaConfigs,omitnil" name:"DomainAreaConfigs"`
}

func (r *CreateClsLogTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClsLogTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "LogsetId")
	delete(f, "Channel")
	delete(f, "DomainAreaConfigs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClsLogTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClsLogTopicResponseParams struct {
	// 主题ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateClsLogTopicResponse struct {
	*tchttp.BaseResponse
	Response *CreateClsLogTopicResponseParams `json:"Response"`
}

func (r *CreateClsLogTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClsLogTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDiagnoseUrlRequestParams struct {
	// 需诊断的url，形如：http://www.test.com/test.txt。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 请求源带协议头，形如：https://console.cloud.tencent.com
	Origin *string `json:"Origin,omitnil" name:"Origin"`
}

type CreateDiagnoseUrlRequest struct {
	*tchttp.BaseRequest
	
	// 需诊断的url，形如：http://www.test.com/test.txt。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 请求源带协议头，形如：https://console.cloud.tencent.com
	Origin *string `json:"Origin,omitnil" name:"Origin"`
}

func (r *CreateDiagnoseUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDiagnoseUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Url")
	delete(f, "Origin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDiagnoseUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDiagnoseUrlResponseParams struct {
	// 系统生成的诊断链接，一个诊断链接最多可访问10次，有效期为24h。
	DiagnoseLink *string `json:"DiagnoseLink,omitnil" name:"DiagnoseLink"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDiagnoseUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateDiagnoseUrlResponseParams `json:"Response"`
}

func (r *CreateDiagnoseUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDiagnoseUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgePackTaskRequestParams struct {
	// apk 所在的 cos 存储桶, 如 edgepack-xxxxxxxx
	CosBucket *string `json:"CosBucket,omitnil" name:"CosBucket"`

	// apk 源文件的存储路径, 如 /apk/xxxx.apk
	CosUriFrom *string `json:"CosUriFrom,omitnil" name:"CosUriFrom"`

	// BlockID 的值, WALLE为1903654775(0x71777777)，VasDolly为2282837503(0x881155ff),传0或不传时默认为 WALLE 方案
	BlockID *uint64 `json:"BlockID,omitnil" name:"BlockID"`

	// 拓展之后的 apk 目标存储路径,如 /out/xxxx.apk
	CosUriTo *string `json:"CosUriTo,omitnil" name:"CosUriTo"`
}

type CreateEdgePackTaskRequest struct {
	*tchttp.BaseRequest
	
	// apk 所在的 cos 存储桶, 如 edgepack-xxxxxxxx
	CosBucket *string `json:"CosBucket,omitnil" name:"CosBucket"`

	// apk 源文件的存储路径, 如 /apk/xxxx.apk
	CosUriFrom *string `json:"CosUriFrom,omitnil" name:"CosUriFrom"`

	// BlockID 的值, WALLE为1903654775(0x71777777)，VasDolly为2282837503(0x881155ff),传0或不传时默认为 WALLE 方案
	BlockID *uint64 `json:"BlockID,omitnil" name:"BlockID"`

	// 拓展之后的 apk 目标存储路径,如 /out/xxxx.apk
	CosUriTo *string `json:"CosUriTo,omitnil" name:"CosUriTo"`
}

func (r *CreateEdgePackTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgePackTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CosBucket")
	delete(f, "CosUriFrom")
	delete(f, "BlockID")
	delete(f, "CosUriTo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEdgePackTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgePackTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateEdgePackTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateEdgePackTaskResponseParams `json:"Response"`
}

func (r *CreateEdgePackTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgePackTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScdnDomainRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Web 攻击防护（WAF）配置
	Waf *ScdnWafConfig `json:"Waf,omitnil" name:"Waf"`

	// 自定义防护策略配置
	Acl *ScdnAclConfig `json:"Acl,omitnil" name:"Acl"`

	// CC 防护配置，目前 CC 防护默认开启
	CC *ScdnConfig `json:"CC,omitnil" name:"CC"`

	// DDOS 防护配置，目前 DDoS 防护默认开启
	Ddos *ScdnDdosConfig `json:"Ddos,omitnil" name:"Ddos"`

	// BOT 防护配置
	Bot *ScdnBotConfig `json:"Bot,omitnil" name:"Bot"`
}

type CreateScdnDomainRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Web 攻击防护（WAF）配置
	Waf *ScdnWafConfig `json:"Waf,omitnil" name:"Waf"`

	// 自定义防护策略配置
	Acl *ScdnAclConfig `json:"Acl,omitnil" name:"Acl"`

	// CC 防护配置，目前 CC 防护默认开启
	CC *ScdnConfig `json:"CC,omitnil" name:"CC"`

	// DDOS 防护配置，目前 DDoS 防护默认开启
	Ddos *ScdnDdosConfig `json:"Ddos,omitnil" name:"Ddos"`

	// BOT 防护配置
	Bot *ScdnBotConfig `json:"Bot,omitnil" name:"Bot"`
}

func (r *CreateScdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScdnDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Waf")
	delete(f, "Acl")
	delete(f, "CC")
	delete(f, "Ddos")
	delete(f, "Bot")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateScdnDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScdnDomainResponseParams struct {
	// 创建结果，Success表示成功
	Result *string `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateScdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateScdnDomainResponseParams `json:"Response"`
}

func (r *CreateScdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScdnDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScdnFailedLogTaskRequestParams struct {
	// 重试失败任务的taskID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 地域：mainland或overseas
	Area *string `json:"Area,omitnil" name:"Area"`
}

type CreateScdnFailedLogTaskRequest struct {
	*tchttp.BaseRequest
	
	// 重试失败任务的taskID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 地域：mainland或overseas
	Area *string `json:"Area,omitnil" name:"Area"`
}

func (r *CreateScdnFailedLogTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScdnFailedLogTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateScdnFailedLogTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScdnFailedLogTaskResponseParams struct {
	// 创建结果, 
	// "0" -> 创建成功
	Result *string `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateScdnFailedLogTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateScdnFailedLogTaskResponseParams `json:"Response"`
}

func (r *CreateScdnFailedLogTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScdnFailedLogTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScdnLogTaskRequestParams struct {
	// 防护类型
	// Mode 映射如下：
	//   waf = "Web攻击"
	//   cc = "CC攻击"
	//   bot = "Bot攻击"
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 指定域名查询, 不填默认查询全部域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 指定攻击类型, 不填默认查询全部攻击类型
	// AttackType 映射如下:
	//   other = '未知类型'
	//   malicious_scan = "恶意扫描"
	//   sql_inject = "SQL注入攻击"
	//   xss = "XSS攻击"
	//   cmd_inject = "命令注入攻击"
	//   ldap_inject = "LDAP注入攻击"
	//   ssi_inject = "SSI注入攻击"
	//   xml_inject = "XML注入攻击"
	//   web_service = "WEB服务漏洞攻击"
	//   web_app = "WEB应用漏洞攻击"
	//   path_traversal = "路径跨越攻击"
	//   illegal_access_core_file = "核心文件非法访问"
	//   trojan_horse = "木马后门攻击"
	//   csrf = "CSRF攻击"
	//   malicious_file_upload= '恶意文件上传'
	//   js = "JS主动探测"
	//   cookie = "Cookie指纹"
	AttackType *string `json:"AttackType,omitnil" name:"AttackType"`

	// 指定执行动作, 不填默认查询全部执行动作
	// DefenceMode 映射如下：
	//   observe = '观察模式'
	//   intercept = '拦截模式'
	//   captcha = "验证码"
	//   redirect = "重定向"
	DefenceMode *string `json:"DefenceMode,omitnil" name:"DefenceMode"`

	// 不填为全部ip
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 指定域名查询, 与 Domain 参数同时有值时使用 Domains 参数，不填默认查询全部域名，指定域名查询时最多支持同时选择 5 个域名查询
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 指定攻击类型查询, 与 AttackType 参数同时有值时使用 AttackTypes 参数，不填默认查询全部攻击类型
	AttackTypes []*string `json:"AttackTypes,omitnil" name:"AttackTypes"`

	// 查询条件
	Conditions []*ScdnEventLogConditions `json:"Conditions,omitnil" name:"Conditions"`

	// 来源产品 cdn ecdn
	Source *string `json:"Source,omitnil" name:"Source"`

	// 地域：mainland 或 overseas
	Area *string `json:"Area,omitnil" name:"Area"`
}

type CreateScdnLogTaskRequest struct {
	*tchttp.BaseRequest
	
	// 防护类型
	// Mode 映射如下：
	//   waf = "Web攻击"
	//   cc = "CC攻击"
	//   bot = "Bot攻击"
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 指定域名查询, 不填默认查询全部域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 指定攻击类型, 不填默认查询全部攻击类型
	// AttackType 映射如下:
	//   other = '未知类型'
	//   malicious_scan = "恶意扫描"
	//   sql_inject = "SQL注入攻击"
	//   xss = "XSS攻击"
	//   cmd_inject = "命令注入攻击"
	//   ldap_inject = "LDAP注入攻击"
	//   ssi_inject = "SSI注入攻击"
	//   xml_inject = "XML注入攻击"
	//   web_service = "WEB服务漏洞攻击"
	//   web_app = "WEB应用漏洞攻击"
	//   path_traversal = "路径跨越攻击"
	//   illegal_access_core_file = "核心文件非法访问"
	//   trojan_horse = "木马后门攻击"
	//   csrf = "CSRF攻击"
	//   malicious_file_upload= '恶意文件上传'
	//   js = "JS主动探测"
	//   cookie = "Cookie指纹"
	AttackType *string `json:"AttackType,omitnil" name:"AttackType"`

	// 指定执行动作, 不填默认查询全部执行动作
	// DefenceMode 映射如下：
	//   observe = '观察模式'
	//   intercept = '拦截模式'
	//   captcha = "验证码"
	//   redirect = "重定向"
	DefenceMode *string `json:"DefenceMode,omitnil" name:"DefenceMode"`

	// 不填为全部ip
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 指定域名查询, 与 Domain 参数同时有值时使用 Domains 参数，不填默认查询全部域名，指定域名查询时最多支持同时选择 5 个域名查询
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 指定攻击类型查询, 与 AttackType 参数同时有值时使用 AttackTypes 参数，不填默认查询全部攻击类型
	AttackTypes []*string `json:"AttackTypes,omitnil" name:"AttackTypes"`

	// 查询条件
	Conditions []*ScdnEventLogConditions `json:"Conditions,omitnil" name:"Conditions"`

	// 来源产品 cdn ecdn
	Source *string `json:"Source,omitnil" name:"Source"`

	// 地域：mainland 或 overseas
	Area *string `json:"Area,omitnil" name:"Area"`
}

func (r *CreateScdnLogTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScdnLogTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mode")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Domain")
	delete(f, "AttackType")
	delete(f, "DefenceMode")
	delete(f, "Ip")
	delete(f, "Domains")
	delete(f, "AttackTypes")
	delete(f, "Conditions")
	delete(f, "Source")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateScdnLogTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScdnLogTaskResponseParams struct {
	// 创建结果, 
	// "0" -> 创建成功
	Result *string `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateScdnLogTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateScdnLogTaskResponseParams `json:"Response"`
}

func (r *CreateScdnLogTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScdnLogTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVerifyRecordRequestParams struct {
	// 要取回的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type CreateVerifyRecordRequest struct {
	*tchttp.BaseRequest
	
	// 要取回的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *CreateVerifyRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVerifyRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVerifyRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVerifyRecordResponseParams struct {
	// 子解析
	SubDomain *string `json:"SubDomain,omitnil" name:"SubDomain"`

	// 解析值
	Record *string `json:"Record,omitnil" name:"Record"`

	// 解析类型
	RecordType *string `json:"RecordType,omitnil" name:"RecordType"`

	// 文件验证 URL 指引
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileVerifyUrl *string `json:"FileVerifyUrl,omitnil" name:"FileVerifyUrl"`

	// 文件校验域名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileVerifyDomains []*string `json:"FileVerifyDomains,omitnil" name:"FileVerifyDomains"`

	// 文件校验文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileVerifyName *string `json:"FileVerifyName,omitnil" name:"FileVerifyName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateVerifyRecordResponse struct {
	*tchttp.BaseResponse
	Response *CreateVerifyRecordResponseParams `json:"Response"`
}

func (r *CreateVerifyRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVerifyRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DDoSAttackBandwidthData struct {
	// ddos攻击类型，当值为all的时候表示所有的攻击类型的总带宽峰值
	AttackType *string `json:"AttackType,omitnil" name:"AttackType"`

	// ddos攻击带宽大小
	Value *float64 `json:"Value,omitnil" name:"Value"`

	// 攻击时间点
	Time *string `json:"Time,omitnil" name:"Time"`
}

type DDoSAttackIPTopData struct {
	// 攻击ip
	AttackIP *string `json:"AttackIP,omitnil" name:"AttackIP"`

	// 攻击ip所在省份
	Province *string `json:"Province,omitnil" name:"Province"`

	// 攻击ip所在国家
	Country *string `json:"Country,omitnil" name:"Country"`

	// 红果电信
	Isp *string `json:"Isp,omitnil" name:"Isp"`

	// 攻击次数
	AttackCount *float64 `json:"AttackCount,omitnil" name:"AttackCount"`
}

type DDoSStatsData struct {
	// 时间
	Time *string `json:"Time,omitnil" name:"Time"`

	// 带宽数值，单位bps
	Value *float64 `json:"Value,omitnil" name:"Value"`
}

type DDoSTopData struct {
	// 攻击类型
	AttackType *string `json:"AttackType,omitnil" name:"AttackType"`

	// 攻击带宽，单位：bps
	Value *uint64 `json:"Value,omitnil" name:"Value"`
}

// Predefined struct for user
type DeleteCdnDomainRequestParams struct {
	// 域名
	// 域名状态需要为【已停用】
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type DeleteCdnDomainRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	// 域名状态需要为【已停用】
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *DeleteCdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCdnDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCdnDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCdnDomainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCdnDomainResponseParams `json:"Response"`
}

func (r *DeleteCdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCdnDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClsLogTopicRequestParams struct {
	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil" name:"LogsetId"`

	// 接入渠道，cdn或者ecdn，默认值为cdn
	Channel *string `json:"Channel,omitnil" name:"Channel"`
}

type DeleteClsLogTopicRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil" name:"LogsetId"`

	// 接入渠道，cdn或者ecdn，默认值为cdn
	Channel *string `json:"Channel,omitnil" name:"Channel"`
}

func (r *DeleteClsLogTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClsLogTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "LogsetId")
	delete(f, "Channel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClsLogTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClsLogTopicResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteClsLogTopicResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClsLogTopicResponseParams `json:"Response"`
}

func (r *DeleteClsLogTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClsLogTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScdnDomainRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type DeleteScdnDomainRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *DeleteScdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScdnDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteScdnDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScdnDomainResponseParams struct {
	// 创建结果，Success表示成功
	Result *string `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteScdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteScdnDomainResponseParams `json:"Response"`
}

func (r *DeleteScdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScdnDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingDataRequestParams struct {
	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	// 根据指定时间粒度参数不同，会进行向前取整，如指定起始时间为 2018-09-04 10:40:00 按小时粒度查询，返回的第一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	// 根据指定时间粒度参数不同，会进行向前取整，如指定结束时间为  2018-09-04 10:40:00 按小时粒度查询时，返回的最后一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 时间粒度，支持模式如下：
	// min：1 分钟粒度，查询区间需要小于等于 24 小时
	// 5min：5 分钟粒度，查询区间需要小于等于 31 天(计费数据粒度)
	// hour：1 小时粒度，查询区间需要小于等于 31 天内
	// day：天粒度，查询区间需要大于 31 天
	// 
	// Area 字段为 overseas 时暂不支持1分钟粒度数据查询
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 指定域名查询计费数据
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 指定项目 ID 查询，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 若 Domain 参数填充了具体域名信息，则返回该域名的计费数据，而非指定项目计费数据
	Project *int64 `json:"Project,omitnil" name:"Project"`

	// 指定加速区域查询计费数据：
	// mainland：中国境内
	// overseas：中国境外
	// 不填充时，默认为 mainland
	Area *string `json:"Area,omitnil" name:"Area"`

	// Area 为 overseas 时，指定国家/地区查询
	// 省份、国家/地区编码可以查看 [省份编码映射](https://cloud.tencent.com/document/product/228/6316#.E7.9C.81.E4.BB.BD.E6.98.A0.E5.B0.84)
	// 不填充时，查询所有国家/地区
	District *int64 `json:"District,omitnil" name:"District"`

	// 计费统计类型
	// flux：计费流量
	// bandwidth：计费带宽
	// 默认为 bandwidth
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 指定查询的产品数据，可选为cdn或者ecdn，默认为cdn
	Product *string `json:"Product,omitnil" name:"Product"`

	// 指定查询时间的时区，默认UTC+08:00
	TimeZone *string `json:"TimeZone,omitnil" name:"TimeZone"`
}

type DescribeBillingDataRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	// 根据指定时间粒度参数不同，会进行向前取整，如指定起始时间为 2018-09-04 10:40:00 按小时粒度查询，返回的第一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	// 根据指定时间粒度参数不同，会进行向前取整，如指定结束时间为  2018-09-04 10:40:00 按小时粒度查询时，返回的最后一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 时间粒度，支持模式如下：
	// min：1 分钟粒度，查询区间需要小于等于 24 小时
	// 5min：5 分钟粒度，查询区间需要小于等于 31 天(计费数据粒度)
	// hour：1 小时粒度，查询区间需要小于等于 31 天内
	// day：天粒度，查询区间需要大于 31 天
	// 
	// Area 字段为 overseas 时暂不支持1分钟粒度数据查询
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 指定域名查询计费数据
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 指定项目 ID 查询，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 若 Domain 参数填充了具体域名信息，则返回该域名的计费数据，而非指定项目计费数据
	Project *int64 `json:"Project,omitnil" name:"Project"`

	// 指定加速区域查询计费数据：
	// mainland：中国境内
	// overseas：中国境外
	// 不填充时，默认为 mainland
	Area *string `json:"Area,omitnil" name:"Area"`

	// Area 为 overseas 时，指定国家/地区查询
	// 省份、国家/地区编码可以查看 [省份编码映射](https://cloud.tencent.com/document/product/228/6316#.E7.9C.81.E4.BB.BD.E6.98.A0.E5.B0.84)
	// 不填充时，查询所有国家/地区
	District *int64 `json:"District,omitnil" name:"District"`

	// 计费统计类型
	// flux：计费流量
	// bandwidth：计费带宽
	// 默认为 bandwidth
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 指定查询的产品数据，可选为cdn或者ecdn，默认为cdn
	Product *string `json:"Product,omitnil" name:"Product"`

	// 指定查询时间的时区，默认UTC+08:00
	TimeZone *string `json:"TimeZone,omitnil" name:"TimeZone"`
}

func (r *DescribeBillingDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Interval")
	delete(f, "Domain")
	delete(f, "Project")
	delete(f, "Area")
	delete(f, "District")
	delete(f, "Metric")
	delete(f, "Product")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillingDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingDataResponseParams struct {
	// 时间粒度，根据查询时传递参数指定：
	// min：1 分钟粒度
	// 5min：5 分钟粒度
	// hour：1 小时粒度
	// day：天粒度
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 数据明细
	Data []*ResourceBillingData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBillingDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillingDataResponseParams `json:"Response"`
}

func (r *DescribeBillingDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcDataRequestParams struct {
	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 时间粒度，支持以下几种模式：
	// min：1 分钟粒度，指定查询区间 24 小时内（含 24 小时），可返回 1 分钟粒度明细数据
	// 5min：5 分钟粒度，指定查询区间 31 天内（含 31 天），可返回 5 分钟粒度明细数据
	// hour：1 小时粒度，指定查询区间 31 天内（含 31 天），可返回 1 小时粒度明细数据
	// day：天粒度，指定查询区间大于 31 天，可返回天粒度明细数据
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 指定域名查询，为空时，表示查询账号级别数据
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 执行动作，取值为：intercept/redirect/observe
	// 分别表示：拦截/重定向/观察
	// 为空时，表示所有执行动作
	ActionName *string `json:"ActionName,omitnil" name:"ActionName"`

	// 指定域名列表查询，为空时，表示查询账号级别数据
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// cdn表示CDN数据，默认值
	// ecdn表示ECDN数据
	Source *string `json:"Source,omitnil" name:"Source"`

	// 地域：mainland或overseas，表示国内或海外，不填写默认表示国内
	Area *string `json:"Area,omitnil" name:"Area"`
}

type DescribeCcDataRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 时间粒度，支持以下几种模式：
	// min：1 分钟粒度，指定查询区间 24 小时内（含 24 小时），可返回 1 分钟粒度明细数据
	// 5min：5 分钟粒度，指定查询区间 31 天内（含 31 天），可返回 5 分钟粒度明细数据
	// hour：1 小时粒度，指定查询区间 31 天内（含 31 天），可返回 1 小时粒度明细数据
	// day：天粒度，指定查询区间大于 31 天，可返回天粒度明细数据
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 指定域名查询，为空时，表示查询账号级别数据
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 执行动作，取值为：intercept/redirect/observe
	// 分别表示：拦截/重定向/观察
	// 为空时，表示所有执行动作
	ActionName *string `json:"ActionName,omitnil" name:"ActionName"`

	// 指定域名列表查询，为空时，表示查询账号级别数据
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// cdn表示CDN数据，默认值
	// ecdn表示ECDN数据
	Source *string `json:"Source,omitnil" name:"Source"`

	// 地域：mainland或overseas，表示国内或海外，不填写默认表示国内
	Area *string `json:"Area,omitnil" name:"Area"`
}

func (r *DescribeCcDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Interval")
	delete(f, "Domain")
	delete(f, "ActionName")
	delete(f, "Domains")
	delete(f, "Source")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCcDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcDataResponseParams struct {
	// 指定执行动作的请求数数据，如果指定类型为空，表示所有类型的请求总数
	Data []*TimestampData `json:"Data,omitnil" name:"Data"`

	// 粒度
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 执行动作为拦截类型QPS统计数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	InterceptQpsData []*TimestampData `json:"InterceptQpsData,omitnil" name:"InterceptQpsData"`

	// 执行动作为重定向类型QPS统计数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectQpsData []*TimestampData `json:"RedirectQpsData,omitnil" name:"RedirectQpsData"`

	// 执行动作为观察类型QPS统计数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObserveQpsData []*TimestampData `json:"ObserveQpsData,omitnil" name:"ObserveQpsData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCcDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCcDataResponseParams `json:"Response"`
}

func (r *DescribeCcDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdnDataRequestParams struct {
	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:00 在按 1 小时的时间粒度查询时，返回的第一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:00 在按 1 小时的时间粒度查询时，返回的最后一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 指定查询指标，支持的类型有：
	// flux：流量，单位为 byte
	// fluxIn：上行流量，单位为 byte，该指标仅ecdn支持查询
	// fluxOut：下行流量，单位为 byte，该指标仅ecdn支持查询
	// bandwidth：带宽，单位为 bps
	// bandwidthIn：上行带宽，单位为 bps，该指标仅ecdn支持查询
	// bandwidthOut：下行带宽，单位为 bps，该指标仅ecdn支持查询
	// request：请求数，单位为 次
	// hitRequest：命中请求数，单位为 次
	// requestHitRate：请求命中率，单位为 %，保留小数点后两位
	// hitFlux：命中流量，单位为byte
	// fluxHitRate：流量命中率，单位为 %，保留小数点后两位
	// statusCode：状态码，返回 2xx、3xx、4xx、5xx 汇总数据，单位为 个
	// 2xx：返回 2xx 状态码汇总及各 2 开头状态码数据，单位为 个
	// 3xx：返回 3xx 状态码汇总及各 3 开头状态码数据，单位为 个
	// 4xx：返回 4xx 状态码汇总及各 4 开头状态码数据，单位为 个
	// 5xx：返回 5xx 状态码汇总及各 5 开头状态码数据，单位为 个
	// 支持指定具体状态码查询，若未产生过，则返回为空
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 指定查询域名列表
	// 查询单域名：指定单个域名
	// 查询多个域名：指定多个域名，最多可一次性查询 30 个
	// 查询账号下所有域名：不传参，默认查询账号维度
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 指定要查询的项目 ID，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 未填充域名情况下，指定项目查询，若填充了具体域名信息，以域名为主
	Project *int64 `json:"Project,omitnil" name:"Project"`

	// 时间粒度，支持以下几种模式：
	// min：1 分钟粒度，指定查询区间 24 小时内（含 24 小时），可返回 1 分钟粒度明细数据（指定查询服务地域为中国境外时不支持 1 分钟粒度）
	// 5min：5 分钟粒度，指定查询区间 31 天内（含 31 天），可返回 5 分钟粒度明细数据
	// hour：1 小时粒度，指定查询区间 31 天内（含 31 天），可返回 1 小时粒度明细数据
	// day：天粒度，指定查询区间大于 31 天，可返回天粒度明细数据
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 多域名查询时，默认（false)返回多个域名的汇总数据
	// 可按需指定为 true，返回每一个 Domain 的明细数据（statusCode 指标暂不支持）
	Detail *bool `json:"Detail,omitnil" name:"Detail"`

	// 查询中国境内CDN数据时，指定运营商查询，不填充表示查询所有运营商
	// 运营商编码可以查看 [运营商编码映射](https://cloud.tencent.com/document/product/228/6316#.E5.8C.BA.E5.9F.9F-.2F-.E8.BF.90.E8.90.A5.E5.95.86.E6.98.A0.E5.B0.84.E8.A1.A8)
	// 指定运营商查询时，不可同时指定省份、IP协议查询
	Isp *int64 `json:"Isp,omitnil" name:"Isp"`

	// 查询中国境内CDN数据时，指定省份查询，不填充表示查询所有省份
	// 查询中国境外CDN数据时，指定国家/地区查询，不填充表示查询所有国家/地区
	// 省份、国家/地区编码可以查看 [省份编码映射](https://cloud.tencent.com/document/product/228/6316#.E5.8C.BA.E5.9F.9F-.2F-.E8.BF.90.E8.90.A5.E5.95.86.E6.98.A0.E5.B0.84.E8.A1.A8)
	// 指定（中国境内）省份查询时，不可同时指定运营商、IP协议查询
	District *int64 `json:"District,omitnil" name:"District"`

	// 指定协议查询，不填充表示查询所有协议
	// all：所有协议
	// http：指定查询 HTTP 对应指标
	// https：指定查询 HTTPS 对应指标
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 指定数据源查询，白名单功能
	DataSource *string `json:"DataSource,omitnil" name:"DataSource"`

	// 指定IP协议查询，不填充表示查询所有协议
	// all：所有协议
	// ipv4：指定查询 ipv4 对应指标
	// ipv6：指定查询 ipv6 对应指标
	// 指定IP协议查询时，不可同时指定省份、运营商查询
	// 注意：非IPv6白名单用户不可指定ipv4、ipv6进行查询
	IpProtocol *string `json:"IpProtocol,omitnil" name:"IpProtocol"`

	// 指定服务地域查询，不填充表示查询中国境内CDN数据
	// mainland：指定查询中国境内 CDN 数据
	// overseas：指定查询中国境外 CDN 数据
	Area *string `json:"Area,omitnil" name:"Area"`

	// 查询中国境外CDN数据时，可指定地区类型查询，不填充表示查询服务地区数据（仅在 Area 为 overseas 时可用）
	// server：指定查询服务地区（腾讯云 CDN 节点服务器所在地区）数据
	// client：指定查询客户端地区（用户请求终端所在地区）数据
	AreaType *string `json:"AreaType,omitnil" name:"AreaType"`

	// 指定查询的产品数据，可选为cdn或者ecdn，默认为cdn
	Product *string `json:"Product,omitnil" name:"Product"`

	// 指定查询时间的时区，默认UTC+08:00
	TimeZone *string `json:"TimeZone,omitnil" name:"TimeZone"`
}

type DescribeCdnDataRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:00 在按 1 小时的时间粒度查询时，返回的第一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:00 在按 1 小时的时间粒度查询时，返回的最后一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 指定查询指标，支持的类型有：
	// flux：流量，单位为 byte
	// fluxIn：上行流量，单位为 byte，该指标仅ecdn支持查询
	// fluxOut：下行流量，单位为 byte，该指标仅ecdn支持查询
	// bandwidth：带宽，单位为 bps
	// bandwidthIn：上行带宽，单位为 bps，该指标仅ecdn支持查询
	// bandwidthOut：下行带宽，单位为 bps，该指标仅ecdn支持查询
	// request：请求数，单位为 次
	// hitRequest：命中请求数，单位为 次
	// requestHitRate：请求命中率，单位为 %，保留小数点后两位
	// hitFlux：命中流量，单位为byte
	// fluxHitRate：流量命中率，单位为 %，保留小数点后两位
	// statusCode：状态码，返回 2xx、3xx、4xx、5xx 汇总数据，单位为 个
	// 2xx：返回 2xx 状态码汇总及各 2 开头状态码数据，单位为 个
	// 3xx：返回 3xx 状态码汇总及各 3 开头状态码数据，单位为 个
	// 4xx：返回 4xx 状态码汇总及各 4 开头状态码数据，单位为 个
	// 5xx：返回 5xx 状态码汇总及各 5 开头状态码数据，单位为 个
	// 支持指定具体状态码查询，若未产生过，则返回为空
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 指定查询域名列表
	// 查询单域名：指定单个域名
	// 查询多个域名：指定多个域名，最多可一次性查询 30 个
	// 查询账号下所有域名：不传参，默认查询账号维度
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 指定要查询的项目 ID，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 未填充域名情况下，指定项目查询，若填充了具体域名信息，以域名为主
	Project *int64 `json:"Project,omitnil" name:"Project"`

	// 时间粒度，支持以下几种模式：
	// min：1 分钟粒度，指定查询区间 24 小时内（含 24 小时），可返回 1 分钟粒度明细数据（指定查询服务地域为中国境外时不支持 1 分钟粒度）
	// 5min：5 分钟粒度，指定查询区间 31 天内（含 31 天），可返回 5 分钟粒度明细数据
	// hour：1 小时粒度，指定查询区间 31 天内（含 31 天），可返回 1 小时粒度明细数据
	// day：天粒度，指定查询区间大于 31 天，可返回天粒度明细数据
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 多域名查询时，默认（false)返回多个域名的汇总数据
	// 可按需指定为 true，返回每一个 Domain 的明细数据（statusCode 指标暂不支持）
	Detail *bool `json:"Detail,omitnil" name:"Detail"`

	// 查询中国境内CDN数据时，指定运营商查询，不填充表示查询所有运营商
	// 运营商编码可以查看 [运营商编码映射](https://cloud.tencent.com/document/product/228/6316#.E5.8C.BA.E5.9F.9F-.2F-.E8.BF.90.E8.90.A5.E5.95.86.E6.98.A0.E5.B0.84.E8.A1.A8)
	// 指定运营商查询时，不可同时指定省份、IP协议查询
	Isp *int64 `json:"Isp,omitnil" name:"Isp"`

	// 查询中国境内CDN数据时，指定省份查询，不填充表示查询所有省份
	// 查询中国境外CDN数据时，指定国家/地区查询，不填充表示查询所有国家/地区
	// 省份、国家/地区编码可以查看 [省份编码映射](https://cloud.tencent.com/document/product/228/6316#.E5.8C.BA.E5.9F.9F-.2F-.E8.BF.90.E8.90.A5.E5.95.86.E6.98.A0.E5.B0.84.E8.A1.A8)
	// 指定（中国境内）省份查询时，不可同时指定运营商、IP协议查询
	District *int64 `json:"District,omitnil" name:"District"`

	// 指定协议查询，不填充表示查询所有协议
	// all：所有协议
	// http：指定查询 HTTP 对应指标
	// https：指定查询 HTTPS 对应指标
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 指定数据源查询，白名单功能
	DataSource *string `json:"DataSource,omitnil" name:"DataSource"`

	// 指定IP协议查询，不填充表示查询所有协议
	// all：所有协议
	// ipv4：指定查询 ipv4 对应指标
	// ipv6：指定查询 ipv6 对应指标
	// 指定IP协议查询时，不可同时指定省份、运营商查询
	// 注意：非IPv6白名单用户不可指定ipv4、ipv6进行查询
	IpProtocol *string `json:"IpProtocol,omitnil" name:"IpProtocol"`

	// 指定服务地域查询，不填充表示查询中国境内CDN数据
	// mainland：指定查询中国境内 CDN 数据
	// overseas：指定查询中国境外 CDN 数据
	Area *string `json:"Area,omitnil" name:"Area"`

	// 查询中国境外CDN数据时，可指定地区类型查询，不填充表示查询服务地区数据（仅在 Area 为 overseas 时可用）
	// server：指定查询服务地区（腾讯云 CDN 节点服务器所在地区）数据
	// client：指定查询客户端地区（用户请求终端所在地区）数据
	AreaType *string `json:"AreaType,omitnil" name:"AreaType"`

	// 指定查询的产品数据，可选为cdn或者ecdn，默认为cdn
	Product *string `json:"Product,omitnil" name:"Product"`

	// 指定查询时间的时区，默认UTC+08:00
	TimeZone *string `json:"TimeZone,omitnil" name:"TimeZone"`
}

func (r *DescribeCdnDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Metric")
	delete(f, "Domains")
	delete(f, "Project")
	delete(f, "Interval")
	delete(f, "Detail")
	delete(f, "Isp")
	delete(f, "District")
	delete(f, "Protocol")
	delete(f, "DataSource")
	delete(f, "IpProtocol")
	delete(f, "Area")
	delete(f, "AreaType")
	delete(f, "Product")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCdnDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdnDataResponseParams struct {
	// 返回数据的时间粒度，查询时指定：
	// min：1 分钟粒度
	// 5min：5 分钟粒度
	// hour：1 小时粒度
	// day：天粒度
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 指定条件查询得到的数据明细
	Data []*ResourceData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCdnDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCdnDataResponseParams `json:"Response"`
}

func (r *DescribeCdnDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdnDomainLogsRequestParams struct {
	// 指定域名查询
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目，默认为 100，最大为 1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 指定区域下载日志，默认为 mainland，可取值有：
	// <li>mainland：获取境内加速日志包下载链接</li>
	// <li>overseas：获取境外加速日志包下载链接</li>
	// <li>global：同时获取境内、境外加速日志包下载链接（分开打包）</li>
	Area *string `json:"Area,omitnil" name:"Area"`

	// 指定下载日志的类型，可取值有：
	// <li>access：访问日志</li>
	LogType *string `json:"LogType,omitnil" name:"LogType"`
}

type DescribeCdnDomainLogsRequest struct {
	*tchttp.BaseRequest
	
	// 指定域名查询
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目，默认为 100，最大为 1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 指定区域下载日志，默认为 mainland，可取值有：
	// <li>mainland：获取境内加速日志包下载链接</li>
	// <li>overseas：获取境外加速日志包下载链接</li>
	// <li>global：同时获取境内、境外加速日志包下载链接（分开打包）</li>
	Area *string `json:"Area,omitnil" name:"Area"`

	// 指定下载日志的类型，可取值有：
	// <li>access：访问日志</li>
	LogType *string `json:"LogType,omitnil" name:"LogType"`
}

func (r *DescribeCdnDomainLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnDomainLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Area")
	delete(f, "LogType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCdnDomainLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdnDomainLogsResponseParams struct {
	// 日志包下载链接。
	// 下载内容是gz后缀的压缩包，解压后是无扩展名的文本文件。
	DomainLogs []*DomainLog `json:"DomainLogs,omitnil" name:"DomainLogs"`

	// 查询到的总条数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCdnDomainLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCdnDomainLogsResponseParams `json:"Response"`
}

func (r *DescribeCdnDomainLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnDomainLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdnIpRequestParams struct {
	// 需要查询的 IP 列表
	Ips []*string `json:"Ips,omitnil" name:"Ips"`
}

type DescribeCdnIpRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的 IP 列表
	Ips []*string `json:"Ips,omitnil" name:"Ips"`
}

func (r *DescribeCdnIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ips")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCdnIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdnIpResponseParams struct {
	// 查询的节点归属详情。
	Ips []*CdnIp `json:"Ips,omitnil" name:"Ips"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCdnIpResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCdnIpResponseParams `json:"Response"`
}

func (r *DescribeCdnIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdnOriginIpRequestParams struct {

}

type DescribeCdnOriginIpRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCdnOriginIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnOriginIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCdnOriginIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCdnOriginIpResponseParams struct {
	// 回源节点IP详情。
	Ips []*OriginIp `json:"Ips,omitnil" name:"Ips"`

	// 回源节点IP总个数。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCdnOriginIpResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCdnOriginIpResponseParams `json:"Response"`
}

func (r *DescribeCdnOriginIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCdnOriginIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertDomainsRequestParams struct {
	// PEM格式证书Base64编码后的字符串
	Cert *string `json:"Cert,omitnil" name:"Cert"`

	// 托管证书ID，Cert和CertId不能均未空，都填写时以CerId为准。
	CertId *string `json:"CertId,omitnil" name:"CertId"`

	// 域名所属产品，cdn或ecdn，默认cdn。
	Product *string `json:"Product,omitnil" name:"Product"`
}

type DescribeCertDomainsRequest struct {
	*tchttp.BaseRequest
	
	// PEM格式证书Base64编码后的字符串
	Cert *string `json:"Cert,omitnil" name:"Cert"`

	// 托管证书ID，Cert和CertId不能均未空，都填写时以CerId为准。
	CertId *string `json:"CertId,omitnil" name:"CertId"`

	// 域名所属产品，cdn或ecdn，默认cdn。
	Product *string `json:"Product,omitnil" name:"Product"`
}

func (r *DescribeCertDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Cert")
	delete(f, "CertId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertDomainsResponseParams struct {
	// 已接入CDN的域名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 已配置证书的CDN域名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertifiedDomains []*string `json:"CertifiedDomains,omitnil" name:"CertifiedDomains"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCertDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCertDomainsResponseParams `json:"Response"`
}

func (r *DescribeCertDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSDataRequestParams struct {
	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 时间粒度，支持以下几种模式：
	// min：1 分钟粒度，指定查询区间 24 小时内（含 24 小时），可返回 1 分钟粒度明细数据
	// 5min：5 分钟粒度，指定查询区间 31 天内（含 31 天），可返回 5 分钟粒度明细数据
	// hour：1 小时粒度，指定查询区间 31 天内（含 31 天），可返回 1 小时粒度明细数据
	// day：天粒度，指定查询区间大于 31 天，可返回天粒度明细数据
	Interval *string `json:"Interval,omitnil" name:"Interval"`
}

type DescribeDDoSDataRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 时间粒度，支持以下几种模式：
	// min：1 分钟粒度，指定查询区间 24 小时内（含 24 小时），可返回 1 分钟粒度明细数据
	// 5min：5 分钟粒度，指定查询区间 31 天内（含 31 天），可返回 5 分钟粒度明细数据
	// hour：1 小时粒度，指定查询区间 31 天内（含 31 天），可返回 1 小时粒度明细数据
	// day：天粒度，指定查询区间大于 31 天，可返回天粒度明细数据
	Interval *string `json:"Interval,omitnil" name:"Interval"`
}

func (r *DescribeDDoSDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Interval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSDataResponseParams struct {
	// DDoS统计数据数组
	Data []*DDoSStatsData `json:"Data,omitnil" name:"Data"`

	// 时间粒度：
	// min：1 分钟粒度
	// 5min：5 分钟粒度
	// hour：1 小时粒度
	// day：天粒度
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// DDoS统计攻击带宽峰值数组
	AttackBandwidthData []*DDoSAttackBandwidthData `json:"AttackBandwidthData,omitnil" name:"AttackBandwidthData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDDoSDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSDataResponseParams `json:"Response"`
}

func (r *DescribeDDoSDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiagnoseReportRequestParams struct {
	// 报告ID
	ReportId *string `json:"ReportId,omitnil" name:"ReportId"`
}

type DescribeDiagnoseReportRequest struct {
	*tchttp.BaseRequest
	
	// 报告ID
	ReportId *string `json:"ReportId,omitnil" name:"ReportId"`
}

func (r *DescribeDiagnoseReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiagnoseReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReportId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiagnoseReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiagnoseReportResponseParams struct {
	// 诊断报告基础信息
	BaskInfo *DiagnoseData `json:"BaskInfo,omitnil" name:"BaskInfo"`

	// CNAME检测信息
	CnameInfo *DiagnoseData `json:"CnameInfo,omitnil" name:"CnameInfo"`

	// 客户端检测信息
	ClientInfo *DiagnoseData `json:"ClientInfo,omitnil" name:"ClientInfo"`

	// DNS检测信息
	DnsInfo *DiagnoseData `json:"DnsInfo,omitnil" name:"DnsInfo"`

	// 网络检测信息
	NetworkInfo *DiagnoseData `json:"NetworkInfo,omitnil" name:"NetworkInfo"`

	// 边缘节点检测信息
	OcNodeInfo *DiagnoseData `json:"OcNodeInfo,omitnil" name:"OcNodeInfo"`

	// 中间源节点检测信息
	MidNodeInfo *DiagnoseData `json:"MidNodeInfo,omitnil" name:"MidNodeInfo"`

	// 源站检测信息
	OriginInfo *DiagnoseData `json:"OriginInfo,omitnil" name:"OriginInfo"`

	// 刷新检测信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PurgeInfo *DiagnoseData `json:"PurgeInfo,omitnil" name:"PurgeInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDiagnoseReportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiagnoseReportResponseParams `json:"Response"`
}

func (r *DescribeDiagnoseReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiagnoseReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDistrictIspDataRequestParams struct {
	// 域名列表，最多支持20个域名
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	// 支持近 60 天内的数据查询，每次查询时间区间为 3 小时
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	// 结束时间与起始时间区间最大为 3 小时
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 指定查询指标，支持:
	// bandwidth：带宽，单位为 bps
	// flux：流量，单位为 byte
	// request：请求数，单位为 次
	// statusCode：状态码，返回 0、2xx、3xx、4xx、5xx 汇总数据，单位为 次
	// 2xx：返回 2xx 状态码汇总及各 2 开头状态码数据，单位为 次
	// 3xx：返回 3xx 状态码汇总及各 3 开头状态码数据，单位为 次
	// 4xx：返回 4xx 状态码汇总及各 4 开头状态码数据，单位为 次
	// 5xx：返回 5xx 状态码汇总及各 5 开头状态码数据，单位为 次
	// 支持指定具体状态码查询，若未产生过，则返回为空
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 指定省份查询，不填充表示查询所有省份
	// 省份、国家/地区编码可以查看 [省份编码映射](https://cloud.tencent.com/document/product/228/6316#.E5.8C.BA.E5.9F.9F-.2F-.E8.BF.90.E8.90.A5.E5.95.86.E6.98.A0.E5.B0.84.E8.A1.A8)
	Districts []*int64 `json:"Districts,omitnil" name:"Districts"`

	// 指定运营商查询，不填充表示查询所有运营商
	// 运营商编码可以查看 [运营商编码映射](https://cloud.tencent.com/document/product/228/6316#.E5.8C.BA.E5.9F.9F-.2F-.E8.BF.90.E8.90.A5.E5.95.86.E6.98.A0.E5.B0.84.E8.A1.A8)
	Isps []*int64 `json:"Isps,omitnil" name:"Isps"`

	// 指定协议查询，不填充表示查询所有协议
	// all：所有协议
	// http：指定查询 HTTP 对应指标
	// https：指定查询 HTTPS 对应指标
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 指定IP协议查询，不填充表示查询所有协议
	// all：所有协议
	// ipv4：指定查询 ipv4 对应指标
	// ipv6：指定查询 ipv6 对应指标
	// 指定IP协议查询时，不可同时指定省份、运营商查询
	IpProtocol *string `json:"IpProtocol,omitnil" name:"IpProtocol"`

	// 时间粒度，支持以下几种模式（默认5min）：
	// min：1 分钟粒度，支持近 60 天内的数据查询，每次查询时间区间不超过10分钟，可返回 1 分钟粒度明细数据
	// 5min：5 分钟粒度，支持近 60 天内的数据查询，每次查询时间区间不超过3 小时，可返回 5 分钟粒度明细数据
	Interval *string `json:"Interval,omitnil" name:"Interval"`
}

type DescribeDistrictIspDataRequest struct {
	*tchttp.BaseRequest
	
	// 域名列表，最多支持20个域名
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	// 支持近 60 天内的数据查询，每次查询时间区间为 3 小时
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	// 结束时间与起始时间区间最大为 3 小时
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 指定查询指标，支持:
	// bandwidth：带宽，单位为 bps
	// flux：流量，单位为 byte
	// request：请求数，单位为 次
	// statusCode：状态码，返回 0、2xx、3xx、4xx、5xx 汇总数据，单位为 次
	// 2xx：返回 2xx 状态码汇总及各 2 开头状态码数据，单位为 次
	// 3xx：返回 3xx 状态码汇总及各 3 开头状态码数据，单位为 次
	// 4xx：返回 4xx 状态码汇总及各 4 开头状态码数据，单位为 次
	// 5xx：返回 5xx 状态码汇总及各 5 开头状态码数据，单位为 次
	// 支持指定具体状态码查询，若未产生过，则返回为空
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 指定省份查询，不填充表示查询所有省份
	// 省份、国家/地区编码可以查看 [省份编码映射](https://cloud.tencent.com/document/product/228/6316#.E5.8C.BA.E5.9F.9F-.2F-.E8.BF.90.E8.90.A5.E5.95.86.E6.98.A0.E5.B0.84.E8.A1.A8)
	Districts []*int64 `json:"Districts,omitnil" name:"Districts"`

	// 指定运营商查询，不填充表示查询所有运营商
	// 运营商编码可以查看 [运营商编码映射](https://cloud.tencent.com/document/product/228/6316#.E5.8C.BA.E5.9F.9F-.2F-.E8.BF.90.E8.90.A5.E5.95.86.E6.98.A0.E5.B0.84.E8.A1.A8)
	Isps []*int64 `json:"Isps,omitnil" name:"Isps"`

	// 指定协议查询，不填充表示查询所有协议
	// all：所有协议
	// http：指定查询 HTTP 对应指标
	// https：指定查询 HTTPS 对应指标
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 指定IP协议查询，不填充表示查询所有协议
	// all：所有协议
	// ipv4：指定查询 ipv4 对应指标
	// ipv6：指定查询 ipv6 对应指标
	// 指定IP协议查询时，不可同时指定省份、运营商查询
	IpProtocol *string `json:"IpProtocol,omitnil" name:"IpProtocol"`

	// 时间粒度，支持以下几种模式（默认5min）：
	// min：1 分钟粒度，支持近 60 天内的数据查询，每次查询时间区间不超过10分钟，可返回 1 分钟粒度明细数据
	// 5min：5 分钟粒度，支持近 60 天内的数据查询，每次查询时间区间不超过3 小时，可返回 5 分钟粒度明细数据
	Interval *string `json:"Interval,omitnil" name:"Interval"`
}

func (r *DescribeDistrictIspDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDistrictIspDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Metric")
	delete(f, "Districts")
	delete(f, "Isps")
	delete(f, "Protocol")
	delete(f, "IpProtocol")
	delete(f, "Interval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDistrictIspDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDistrictIspDataResponseParams struct {
	// 地区运营商数据明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DistrictIspInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDistrictIspDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDistrictIspDataResponseParams `json:"Response"`
}

func (r *DescribeDistrictIspDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDistrictIspDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainsConfigRequestParams struct {
	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目，默认为 100，最大可设置为 1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询条件过滤器，复杂类型
	Filters []*DomainFilter `json:"Filters,omitnil" name:"Filters"`

	// 排序规则
	Sort *Sort `json:"Sort,omitnil" name:"Sort"`
}

type DescribeDomainsConfigRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目，默认为 100，最大可设置为 1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询条件过滤器，复杂类型
	Filters []*DomainFilter `json:"Filters,omitnil" name:"Filters"`

	// 排序规则
	Sort *Sort `json:"Sort,omitnil" name:"Sort"`
}

func (r *DescribeDomainsConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainsConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainsConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainsConfigResponseParams struct {
	// 域名列表
	Domains []*DetailDomain `json:"Domains,omitnil" name:"Domains"`

	// 符合查询条件的域名总数
	// 用于分页查询
	TotalNumber *int64 `json:"TotalNumber,omitnil" name:"TotalNumber"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDomainsConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainsConfigResponseParams `json:"Response"`
}

func (r *DescribeDomainsConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainsConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainsRequestParams struct {
	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目，默认为 100，最大可设置为 1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询条件过滤器，复杂类型
	Filters []*DomainFilter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目，默认为 100，最大可设置为 1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询条件过滤器，复杂类型
	Filters []*DomainFilter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainsResponseParams struct {
	// 域名列表
	Domains []*BriefDomain `json:"Domains,omitnil" name:"Domains"`

	// 符合查询条件的域名总数
	// 用于分页查询
	TotalNumber *int64 `json:"TotalNumber,omitnil" name:"TotalNumber"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainsResponseParams `json:"Response"`
}

func (r *DescribeDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgePackTaskStatusRequestParams struct {
	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分页查询限制数目，默认为 100，最大可设置为 1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询条件过滤器
	Filters []*EdgePackTaskFilter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeEdgePackTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分页查询限制数目，默认为 100，最大可设置为 1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询条件过滤器
	Filters []*EdgePackTaskFilter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeEdgePackTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgePackTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgePackTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgePackTaskStatusResponseParams struct {
	// 动态打包任务状态列表
	EdgePackTaskStatusSet []*EdgePackTaskStatus `json:"EdgePackTaskStatusSet,omitnil" name:"EdgePackTaskStatusSet"`

	// 总数，用于分页查询
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEdgePackTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgePackTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeEdgePackTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgePackTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventLogDataRequestParams struct {
	// 防护类型，映射如下：
	//   waf = "Web攻击"
	//   cc = "CC攻击"
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间，最长跨度为30分钟
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 执行动作，取值为：intercept/redirect/observe
	// 分别表示：拦截/重定向/观察
	// 参数放空，表示查询全部动作数据
	ActionName *string `json:"ActionName,omitnil" name:"ActionName"`

	// 请求URL，支持URL开头和结尾使用\*表示通配
	// 如：
	// /files/* 表示所有以/files/开头的请求
	// *.jpg 表示所有以.jpg结尾的请求
	Url *string `json:"Url,omitnil" name:"Url"`

	// 地域 mainland 或者 overseas，为空时默认 mainland
	Area *string `json:"Area,omitnil" name:"Area"`

	// 来源产品，cdn 或者 ecdn，为空时默认 cdn
	Source *string `json:"Source,omitnil" name:"Source"`
}

type DescribeEventLogDataRequest struct {
	*tchttp.BaseRequest
	
	// 防护类型，映射如下：
	//   waf = "Web攻击"
	//   cc = "CC攻击"
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间，最长跨度为30分钟
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 执行动作，取值为：intercept/redirect/observe
	// 分别表示：拦截/重定向/观察
	// 参数放空，表示查询全部动作数据
	ActionName *string `json:"ActionName,omitnil" name:"ActionName"`

	// 请求URL，支持URL开头和结尾使用\*表示通配
	// 如：
	// /files/* 表示所有以/files/开头的请求
	// *.jpg 表示所有以.jpg结尾的请求
	Url *string `json:"Url,omitnil" name:"Url"`

	// 地域 mainland 或者 overseas，为空时默认 mainland
	Area *string `json:"Area,omitnil" name:"Area"`

	// 来源产品，cdn 或者 ecdn，为空时默认 cdn
	Source *string `json:"Source,omitnil" name:"Source"`
}

func (r *DescribeEventLogDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventLogDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mode")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Domain")
	delete(f, "ActionName")
	delete(f, "Url")
	delete(f, "Area")
	delete(f, "Source")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEventLogDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventLogDataResponseParams struct {
	// 统计曲线结果
	Results []*EventLogStatsData `json:"Results,omitnil" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEventLogDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEventLogDataResponseParams `json:"Response"`
}

func (r *DescribeEventLogDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventLogDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHttpsPackagesRequestParams struct {
	// 分页查询起始地址，默认 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询记录个数，默认100，最大1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeHttpsPackagesRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询起始地址，默认 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询记录个数，默认100，最大1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeHttpsPackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHttpsPackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHttpsPackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHttpsPackagesResponseParams struct {
	// HTTPS请求包总个数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// HTTPS请求包详情
	HttpsPackages []*HttpsPackage `json:"HttpsPackages,omitnil" name:"HttpsPackages"`

	// 即将过期的HTTPS请求包个数（7天内）
	ExpiringCount *int64 `json:"ExpiringCount,omitnil" name:"ExpiringCount"`

	// 有效HTTPS请求包个数
	EnabledCount *int64 `json:"EnabledCount,omitnil" name:"EnabledCount"`

	// 付费HTTPS请求包个数
	PaidCount *int64 `json:"PaidCount,omitnil" name:"PaidCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeHttpsPackagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHttpsPackagesResponseParams `json:"Response"`
}

func (r *DescribeHttpsPackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHttpsPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageConfigRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type DescribeImageConfigRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *DescribeImageConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageConfigResponseParams struct {
	// WebpAdapter配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebpAdapter *WebpAdapter `json:"WebpAdapter,omitnil" name:"WebpAdapter"`

	// TpgAdapter配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TpgAdapter *TpgAdapter `json:"TpgAdapter,omitnil" name:"TpgAdapter"`

	// GuetzliAdapter配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	GuetzliAdapter *GuetzliAdapter `json:"GuetzliAdapter,omitnil" name:"GuetzliAdapter"`

	// AvifAdapter配置项
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvifAdapter *AvifAdapter `json:"AvifAdapter,omitnil" name:"AvifAdapter"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeImageConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageConfigResponseParams `json:"Response"`
}

func (r *DescribeImageConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpStatusRequestParams struct {
	// 加速域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 节点类型：
	// edge：表示边缘节点
	// last：表示回源层节点
	// 不填充情况下，默认返回边缘节点信息
	Layer *string `json:"Layer,omitnil" name:"Layer"`

	// 查询区域：
	// mainland: 国内节点
	// overseas: 海外节点
	// global: 全球节点
	Area *string `json:"Area,omitnil" name:"Area"`

	// 是否以IP段的格式返回。
	Segment *bool `json:"Segment,omitnil" name:"Segment"`

	// 是否查询节点 IPV6 信息。
	ShowIpv6 *bool `json:"ShowIpv6,omitnil" name:"ShowIpv6"`

	// 是否对IPV6进行缩写。
	AbbreviationIpv6 *bool `json:"AbbreviationIpv6,omitnil" name:"AbbreviationIpv6"`
}

type DescribeIpStatusRequest struct {
	*tchttp.BaseRequest
	
	// 加速域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 节点类型：
	// edge：表示边缘节点
	// last：表示回源层节点
	// 不填充情况下，默认返回边缘节点信息
	Layer *string `json:"Layer,omitnil" name:"Layer"`

	// 查询区域：
	// mainland: 国内节点
	// overseas: 海外节点
	// global: 全球节点
	Area *string `json:"Area,omitnil" name:"Area"`

	// 是否以IP段的格式返回。
	Segment *bool `json:"Segment,omitnil" name:"Segment"`

	// 是否查询节点 IPV6 信息。
	ShowIpv6 *bool `json:"ShowIpv6,omitnil" name:"ShowIpv6"`

	// 是否对IPV6进行缩写。
	AbbreviationIpv6 *bool `json:"AbbreviationIpv6,omitnil" name:"AbbreviationIpv6"`
}

func (r *DescribeIpStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Layer")
	delete(f, "Area")
	delete(f, "Segment")
	delete(f, "ShowIpv6")
	delete(f, "AbbreviationIpv6")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIpStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpStatusResponseParams struct {
	// 节点列表
	Ips []*IpStatus `json:"Ips,omitnil" name:"Ips"`

	// 节点总个数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeIpStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIpStatusResponseParams `json:"Response"`
}

func (r *DescribeIpStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpVisitRequestParams struct {
	// 查询起始时间，如：2018-09-04 10:40:10，返回结果大于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:10 在按 5 分钟的时间粒度查询时，返回的第一个数据对应时间点为 2018-09-04 10:40:00
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:10，返回结果小于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:10 在按 5 分钟的时间粒度查询时，返回的最后一个数据对应时间点为 2018-09-04 10:40:00
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 指定查询域名列表，最多可一次性查询 30 个加速域名明细
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 指定要查询的项目 ID，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 未填充域名情况下，指定项目查询，若填充了具体域名信息，以域名为主
	Project *int64 `json:"Project,omitnil" name:"Project"`

	// 时间粒度，支持以下几种模式：
	// 5min：5 分钟粒度，查询时间区间 24 小时内，默认返回 5 分钟粒度活跃用户数
	// day：天粒度，查询时间区间大于 1 天时，默认返回天粒度活跃用户数
	Interval *string `json:"Interval,omitnil" name:"Interval"`
}

type DescribeIpVisitRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间，如：2018-09-04 10:40:10，返回结果大于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:10 在按 5 分钟的时间粒度查询时，返回的第一个数据对应时间点为 2018-09-04 10:40:00
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:10，返回结果小于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:10 在按 5 分钟的时间粒度查询时，返回的最后一个数据对应时间点为 2018-09-04 10:40:00
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 指定查询域名列表，最多可一次性查询 30 个加速域名明细
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 指定要查询的项目 ID，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 未填充域名情况下，指定项目查询，若填充了具体域名信息，以域名为主
	Project *int64 `json:"Project,omitnil" name:"Project"`

	// 时间粒度，支持以下几种模式：
	// 5min：5 分钟粒度，查询时间区间 24 小时内，默认返回 5 分钟粒度活跃用户数
	// day：天粒度，查询时间区间大于 1 天时，默认返回天粒度活跃用户数
	Interval *string `json:"Interval,omitnil" name:"Interval"`
}

func (r *DescribeIpVisitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpVisitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Domains")
	delete(f, "Project")
	delete(f, "Interval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIpVisitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpVisitResponseParams struct {
	// 数据统计的时间粒度，支持5min,  day，分别表示5分钟，1天的时间粒度。
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 各个资源的回源数据详情。
	Data []*ResourceData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeIpVisitResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIpVisitResponseParams `json:"Response"`
}

func (r *DescribeIpVisitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpVisitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMapInfoRequestParams struct {
	// 映射查询类别：
	// isp：运营商映射查询
	// district：省份（中国境内）、国家/地区（中国境外）映射查询
	Name *string `json:"Name,omitnil" name:"Name"`
}

type DescribeMapInfoRequest struct {
	*tchttp.BaseRequest
	
	// 映射查询类别：
	// isp：运营商映射查询
	// district：省份（中国境内）、国家/地区（中国境外）映射查询
	Name *string `json:"Name,omitnil" name:"Name"`
}

func (r *DescribeMapInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMapInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMapInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMapInfoResponseParams struct {
	// 映射关系数组。
	MapInfoList []*MapInfo `json:"MapInfoList,omitnil" name:"MapInfoList"`

	// 服务端区域id和子区域id的映射关系。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerRegionRelation []*RegionMapRelation `json:"ServerRegionRelation,omitnil" name:"ServerRegionRelation"`

	// 客户端区域id和子区域id的映射关系。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientRegionRelation []*RegionMapRelation `json:"ClientRegionRelation,omitnil" name:"ClientRegionRelation"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMapInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMapInfoResponseParams `json:"Response"`
}

func (r *DescribeMapInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMapInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginDataRequestParams struct {
	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:00 在按 1 小时的时间粒度查询时，返回的第一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:00 在按 1 小时的时间粒度查询时，返回的最后一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 指定查询指标，支持的类型有：
	// flux：回源流量，单位为 byte
	// bandwidth：回源带宽，单位为 bps
	// request：回源请求数，单位为 次
	// failRequest：回源失败请求数，单位为 次
	// failRate：回源失败率，单位为 %
	// statusCode：回源状态码，返回 2xx、3xx、4xx、5xx 汇总数据，单位为 个
	// 2xx：返回 2xx 回源状态码汇总及各 2 开头回源状态码数据，单位为 个
	// 3xx：返回 3xx 回源状态码汇总及各 3 开头回源状态码数据，单位为 个
	// 4xx：返回 4xx 回源状态码汇总及各 4 开头回源状态码数据，单位为 个
	// 5xx：返回 5xx 回源状态码汇总及各 5 开头回源状态码数据，单位为 个
	// 支持指定具体状态码查询，若未产生过，则返回为空
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 指定查询域名列表，最多可一次性查询 30 个加速域名明细
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 指定要查询的项目 ID，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 未填充域名情况下，指定项目查询，最多可一次性查询 30 个加速域名明细
	// 若填充了具体域名信息，以域名为主
	Project *int64 `json:"Project,omitnil" name:"Project"`

	// 时间粒度，支持以下几种模式：
	// min：1 分钟粒度，指定查询区间 24 小时内（含 24 小时），可返回 1 分钟粒度明细数据（指定查询服务地域为中国境外时不支持 1 分钟粒度）
	// 5min：5 分钟粒度，指定查询区间 31 天内（含 31 天），可返回 5 分钟粒度明细数据
	// hour：1 小时粒度，指定查询区间 31 天内（含 31 天），可返回 1 小时粒度明细数据
	// day：天粒度，指定查询区间大于 31 天，可返回天粒度明细数据
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// Domains 传入多个时，默认（false)返回多个域名的汇总数据
	// 可按需指定为 true，返回每一个 Domain 的明细数据（statusCode 指标暂不支持）
	Detail *bool `json:"Detail,omitnil" name:"Detail"`

	// 指定服务地域查询，不填充表示查询中国境内 CDN 数据
	// mainland：指定查询中国境内 CDN 数据
	// overseas：指定查询中国境外 CDN 数据
	Area *string `json:"Area,omitnil" name:"Area"`

	// 指定查询时间的时区，默认UTC+08:00
	TimeZone *string `json:"TimeZone,omitnil" name:"TimeZone"`
}

type DescribeOriginDataRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:00 在按 1 小时的时间粒度查询时，返回的第一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:00 在按 1 小时的时间粒度查询时，返回的最后一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 指定查询指标，支持的类型有：
	// flux：回源流量，单位为 byte
	// bandwidth：回源带宽，单位为 bps
	// request：回源请求数，单位为 次
	// failRequest：回源失败请求数，单位为 次
	// failRate：回源失败率，单位为 %
	// statusCode：回源状态码，返回 2xx、3xx、4xx、5xx 汇总数据，单位为 个
	// 2xx：返回 2xx 回源状态码汇总及各 2 开头回源状态码数据，单位为 个
	// 3xx：返回 3xx 回源状态码汇总及各 3 开头回源状态码数据，单位为 个
	// 4xx：返回 4xx 回源状态码汇总及各 4 开头回源状态码数据，单位为 个
	// 5xx：返回 5xx 回源状态码汇总及各 5 开头回源状态码数据，单位为 个
	// 支持指定具体状态码查询，若未产生过，则返回为空
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 指定查询域名列表，最多可一次性查询 30 个加速域名明细
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 指定要查询的项目 ID，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 未填充域名情况下，指定项目查询，最多可一次性查询 30 个加速域名明细
	// 若填充了具体域名信息，以域名为主
	Project *int64 `json:"Project,omitnil" name:"Project"`

	// 时间粒度，支持以下几种模式：
	// min：1 分钟粒度，指定查询区间 24 小时内（含 24 小时），可返回 1 分钟粒度明细数据（指定查询服务地域为中国境外时不支持 1 分钟粒度）
	// 5min：5 分钟粒度，指定查询区间 31 天内（含 31 天），可返回 5 分钟粒度明细数据
	// hour：1 小时粒度，指定查询区间 31 天内（含 31 天），可返回 1 小时粒度明细数据
	// day：天粒度，指定查询区间大于 31 天，可返回天粒度明细数据
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// Domains 传入多个时，默认（false)返回多个域名的汇总数据
	// 可按需指定为 true，返回每一个 Domain 的明细数据（statusCode 指标暂不支持）
	Detail *bool `json:"Detail,omitnil" name:"Detail"`

	// 指定服务地域查询，不填充表示查询中国境内 CDN 数据
	// mainland：指定查询中国境内 CDN 数据
	// overseas：指定查询中国境外 CDN 数据
	Area *string `json:"Area,omitnil" name:"Area"`

	// 指定查询时间的时区，默认UTC+08:00
	TimeZone *string `json:"TimeZone,omitnil" name:"TimeZone"`
}

func (r *DescribeOriginDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Metric")
	delete(f, "Domains")
	delete(f, "Project")
	delete(f, "Interval")
	delete(f, "Detail")
	delete(f, "Area")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOriginDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginDataResponseParams struct {
	// 数据统计的时间粒度，支持min, 5min, hour, day，分别表示1分钟，5分钟，1小时和1天的时间粒度。
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 各个资源的回源数据详情。
	Data []*ResourceOriginData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeOriginDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOriginDataResponseParams `json:"Response"`
}

func (r *DescribeOriginDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePayTypeRequestParams struct {
	// 指定服务地域查询
	// mainland：境内计费方式查询
	// overseas：境外计费方式查询
	// global：全球计费方式查询
	// 未填充时，默认为 mainland
	Area *string `json:"Area,omitnil" name:"Area"`

	// 指定查询的产品数据，可选为cdn或者ecdn，默认为cdn
	Product *string `json:"Product,omitnil" name:"Product"`

	// 指定资源包查询
	// flux：流量包
	// https：HTTPS请求包
	// 未填充时，默认为 flux
	Type *string `json:"Type,omitnil" name:"Type"`
}

type DescribePayTypeRequest struct {
	*tchttp.BaseRequest
	
	// 指定服务地域查询
	// mainland：境内计费方式查询
	// overseas：境外计费方式查询
	// global：全球计费方式查询
	// 未填充时，默认为 mainland
	Area *string `json:"Area,omitnil" name:"Area"`

	// 指定查询的产品数据，可选为cdn或者ecdn，默认为cdn
	Product *string `json:"Product,omitnil" name:"Product"`

	// 指定资源包查询
	// flux：流量包
	// https：HTTPS请求包
	// 未填充时，默认为 flux
	Type *string `json:"Type,omitnil" name:"Type"`
}

func (r *DescribePayTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePayTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Area")
	delete(f, "Product")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePayTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePayTypeResponseParams struct {
	// 计费类型
	// flux：流量计费
	// bandwidth：带宽计费
	// request：请求数计费
	// flux_sep：动静分离流量计费
	// bandwidth_sep：动静分离带宽计费
	// 日结计费方式切换时，若当日产生消耗，则此字段表示第二天即将生效的计费方式，若未产生消耗，则表示已经生效的计费方式。
	PayType *string `json:"PayType,omitnil" name:"PayType"`

	// 计费周期
	// day：日结计费
	// month：月结计费
	// hour：小时结计费
	BillingCycle *string `json:"BillingCycle,omitnil" name:"BillingCycle"`

	// 统计类型
	// monthMax：日峰值月平均，月结模式
	// day95：日 95 带宽，月结模式
	// month95：月95带宽，月结模式
	// sum：总流量/总请求数，日结或月结模式
	// max：峰值带宽，日结模式
	StatType *string `json:"StatType,omitnil" name:"StatType"`

	// 计费区域
	// all：全地区统一计费
	// multiple：分地区计费
	RegionType *string `json:"RegionType,omitnil" name:"RegionType"`

	// 当前生效计费类型
	// flux：流量计费
	// bandwidth：带宽计费
	// request：请求数计费
	// flux_sep：动静分离流量计费
	// bandwidth_sep：动静分离带宽计费
	CurrentPayType *string `json:"CurrentPayType,omitnil" name:"CurrentPayType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePayTypeResponse struct {
	*tchttp.BaseResponse
	Response *DescribePayTypeResponseParams `json:"Response"`
}

func (r *DescribePayTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePayTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurgeQuotaRequestParams struct {

}

type DescribePurgeQuotaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribePurgeQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurgeQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePurgeQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurgeQuotaResponseParams struct {
	// URL刷新用量及配额。
	UrlPurge []*Quota `json:"UrlPurge,omitnil" name:"UrlPurge"`

	// 目录刷新用量及配额。
	PathPurge []*Quota `json:"PathPurge,omitnil" name:"PathPurge"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePurgeQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribePurgeQuotaResponseParams `json:"Response"`
}

func (r *DescribePurgeQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurgeQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurgeTasksRequestParams struct {
	// 指定刷新类型查询
	// url：url 刷新记录
	// path：目录刷新记录
	PurgeType *string `json:"PurgeType,omitnil" name:"PurgeType"`

	// 根据时间区间查询时，填充开始时间，如 2018-08-08 00:00:00
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 根据时间区间查询时，填充结束时间，如 2018-08-08 23:59:59
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 根据任务 ID 查询时，填充任务 ID
	// 查询时任务 ID 与起始时间必须填充一项
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目，默认为 20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 支持域名过滤，或 http(s):// 开头完整 URL 过滤
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 指定任务状态查询
	// fail：刷新失败
	// done：刷新成功
	// process：刷新中
	Status *string `json:"Status,omitnil" name:"Status"`

	// 指定刷新地域查询
	// mainland：境内
	// overseas：境外
	// global：全球
	Area *string `json:"Area,omitnil" name:"Area"`
}

type DescribePurgeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 指定刷新类型查询
	// url：url 刷新记录
	// path：目录刷新记录
	PurgeType *string `json:"PurgeType,omitnil" name:"PurgeType"`

	// 根据时间区间查询时，填充开始时间，如 2018-08-08 00:00:00
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 根据时间区间查询时，填充结束时间，如 2018-08-08 23:59:59
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 根据任务 ID 查询时，填充任务 ID
	// 查询时任务 ID 与起始时间必须填充一项
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目，默认为 20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 支持域名过滤，或 http(s):// 开头完整 URL 过滤
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 指定任务状态查询
	// fail：刷新失败
	// done：刷新成功
	// process：刷新中
	Status *string `json:"Status,omitnil" name:"Status"`

	// 指定刷新地域查询
	// mainland：境内
	// overseas：境外
	// global：全球
	Area *string `json:"Area,omitnil" name:"Area"`
}

func (r *DescribePurgeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurgeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PurgeType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TaskId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Keyword")
	delete(f, "Status")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePurgeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurgeTasksResponseParams struct {
	// 详细刷新记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	PurgeLogs []*PurgeTask `json:"PurgeLogs,omitnil" name:"PurgeLogs"`

	// 任务总数，用于分页
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePurgeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribePurgeTasksResponseParams `json:"Response"`
}

func (r *DescribePurgeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurgeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePushQuotaRequestParams struct {

}

type DescribePushQuotaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribePushQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePushQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePushQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePushQuotaResponseParams struct {
	// Url预热用量及配额。
	UrlPush []*Quota `json:"UrlPush,omitnil" name:"UrlPush"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePushQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribePushQuotaResponseParams `json:"Response"`
}

func (r *DescribePushQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePushQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePushTasksRequestParams struct {
	// 开始时间，如2018-08-08 00:00:00。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间，如2018-08-08 23:59:59。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 指定任务 ID 查询
	// TaskId 和起始时间必须指定一项
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 查询关键字，请输入域名或 http(s):// 开头完整 URL
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目，默认为 20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 指定地区查询预热记录
	// mainland：境内
	// overseas：境外
	// global：全球
	Area *string `json:"Area,omitnil" name:"Area"`

	// 指定任务状态查询
	// fail：预热失败
	// done：预热成功
	// process：预热中
	// invalid: 预热无效(源站返回4xx或5xx状态码)
	Status *string `json:"Status,omitnil" name:"Status"`
}

type DescribePushTasksRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，如2018-08-08 00:00:00。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间，如2018-08-08 23:59:59。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 指定任务 ID 查询
	// TaskId 和起始时间必须指定一项
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 查询关键字，请输入域名或 http(s):// 开头完整 URL
	Keyword *string `json:"Keyword,omitnil" name:"Keyword"`

	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目，默认为 20
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 指定地区查询预热记录
	// mainland：境内
	// overseas：境外
	// global：全球
	Area *string `json:"Area,omitnil" name:"Area"`

	// 指定任务状态查询
	// fail：预热失败
	// done：预热成功
	// process：预热中
	// invalid: 预热无效(源站返回4xx或5xx状态码)
	Status *string `json:"Status,omitnil" name:"Status"`
}

func (r *DescribePushTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePushTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TaskId")
	delete(f, "Keyword")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Area")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePushTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePushTasksResponseParams struct {
	// 预热历史记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	PushLogs []*PushTask `json:"PushLogs,omitnil" name:"PushLogs"`

	// 任务总数，用于分页
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePushTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribePushTasksResponseParams `json:"Response"`
}

func (r *DescribePushTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePushTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReportDataRequestParams struct {
	// 查询起始时间：yyyy-MM-dd
	// 当报表类型为daily，起始时间和结束时间必须为同一天
	// 当报表类型为weekly，起始时间须为周一，结束时间须为同一周的周日
	// 当报表类型为monthly，起始时间须为自然月第一天，即1号，结束时间须为该自然月最后一天
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间：yyyy-MM-dd
	// 当报表类型为daily，起始时间和结束时间必须为同一天
	// 当报表类型为weekly，起始时间须为周一，结束时间须为同一周的周日
	// 当报表类型为monthly，起始时间须为自然月第一天，即1号，结束时间须为该自然月最后一天
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 报表类型
	// daily：日报表
	// weekly：周报表（周一至周日）
	// monthly：月报表（自然月）
	ReportType *string `json:"ReportType,omitnil" name:"ReportType"`

	// 域名加速区域
	// mainland：中国境内
	// overseas：中国境外
	Area *string `json:"Area,omitnil" name:"Area"`

	// 偏移量，默认0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 数据个数，默认1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 按项目ID筛选
	Project *int64 `json:"Project,omitnil" name:"Project"`
}

type DescribeReportDataRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间：yyyy-MM-dd
	// 当报表类型为daily，起始时间和结束时间必须为同一天
	// 当报表类型为weekly，起始时间须为周一，结束时间须为同一周的周日
	// 当报表类型为monthly，起始时间须为自然月第一天，即1号，结束时间须为该自然月最后一天
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间：yyyy-MM-dd
	// 当报表类型为daily，起始时间和结束时间必须为同一天
	// 当报表类型为weekly，起始时间须为周一，结束时间须为同一周的周日
	// 当报表类型为monthly，起始时间须为自然月第一天，即1号，结束时间须为该自然月最后一天
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 报表类型
	// daily：日报表
	// weekly：周报表（周一至周日）
	// monthly：月报表（自然月）
	ReportType *string `json:"ReportType,omitnil" name:"ReportType"`

	// 域名加速区域
	// mainland：中国境内
	// overseas：中国境外
	Area *string `json:"Area,omitnil" name:"Area"`

	// 偏移量，默认0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 数据个数，默认1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 按项目ID筛选
	Project *int64 `json:"Project,omitnil" name:"Project"`
}

func (r *DescribeReportDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReportDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ReportType")
	delete(f, "Area")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Project")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReportDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReportDataResponseParams struct {
	// 域名维度数据详情
	DomainReport []*ReportData `json:"DomainReport,omitnil" name:"DomainReport"`

	// 项目维度数据详情
	ProjectReport []*ReportData `json:"ProjectReport,omitnil" name:"ProjectReport"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeReportDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReportDataResponseParams `json:"Response"`
}

func (r *DescribeReportDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReportDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScdnBotDataRequestParams struct {
	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// mainland 大陆地区 overseas境外地区
	Area *string `json:"Area,omitnil" name:"Area"`

	// 取值："2min"或者"hour"，表示查询2分钟或者1小时粒度的数据，如果查询时间范围>1天，则强制返回1小时粒度数据
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 域名数组，多选域名时，使用此参数,不填写表示查询所有域名的数据（AppID维度数据）
	Domains []*string `json:"Domains,omitnil" name:"Domains"`
}

type DescribeScdnBotDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// mainland 大陆地区 overseas境外地区
	Area *string `json:"Area,omitnil" name:"Area"`

	// 取值："2min"或者"hour"，表示查询2分钟或者1小时粒度的数据，如果查询时间范围>1天，则强制返回1小时粒度数据
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 域名数组，多选域名时，使用此参数,不填写表示查询所有域名的数据（AppID维度数据）
	Domains []*string `json:"Domains,omitnil" name:"Domains"`
}

func (r *DescribeScdnBotDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScdnBotDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Area")
	delete(f, "Interval")
	delete(f, "Domains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScdnBotDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScdnBotDataResponseParams struct {
	// 统计信息详细数据
	Data []*BotStats `json:"Data,omitnil" name:"Data"`

	// 当前返回数据的粒度，取值："2min"或者"hour"，分别表示2分钟或者1小时粒度
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeScdnBotDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScdnBotDataResponseParams `json:"Response"`
}

func (r *DescribeScdnBotDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScdnBotDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScdnBotRecordsRequestParams struct {
	// BOT类型，取值为"UB","UCB","TCB"，分别表示：未知类型，自定义类型，公开类型
	BotType *string `json:"BotType,omitnil" name:"BotType"`

	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分页参数
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页参数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// mainland 大陆地区 overseas境外地区
	Area *string `json:"Area,omitnil" name:"Area"`

	// 排序参数
	SortBy []*BotSortBy `json:"SortBy,omitnil" name:"SortBy"`

	// BotType=UB时，表示需要过滤的预测标签，取值如下：
	//                 "crawler_unregular",
	//                 "crawler_regular",
	//                 "request_repeat",
	//                 "credential_miss_user",
	//                 "credential_without_user",
	//                 "credential_only_action",
	//                 "credential_user_password",
	//                 "credential_cracking",
	//                 "credential_stuffing",
	//                 "brush_sms",
	//                 "brush_captcha",
	//                 "reg_malicious"
	// BotType=TCB时，表示需要过滤的Bot分类，取值如下：
	//                 "Uncategorised",
	//                 "Search engine bot",
	//                 "Site monitor",
	//                 "Screenshot creator",
	//                 "Link checker",
	//                 "Web scraper",
	//                 "Vulnerability scanner",
	//                 "Virus scanner",
	//                 "Speed tester",
	//                 "Feed Fetcher",
	//                 "Tool",
	//                 "Marketing"
	// BotType=UCB时，取值如下：
	// User-Agent为空或不存在
	// User-Agent类型为BOT
	// User-Agent类型为HTTP Library
	// User-Agent类型为Framework
	// User-Agent类型为Tools
	// User-Agent类型为Unkonwn BOT
	// User-Agent类型为Scanner
	// Referer空或不存在
	// Referer滥用(多个UA使用相同Referer)
	// Cookie滥用(多个UA使用相同Cookie)
	// Cookie空或不存在
	// Connection空或不存在
	// Accept空或不存在
	// Accept-Language空或不存在
	// Accept-Enconding空或不存在
	// 使用HTTP HEAD方法
	// HTTP协议为1.0或者更低
	// IDC-IP 腾讯云
	// IDC-IP 阿里云
	// IDC-IP 华为云
	// IDC-IP 金山云
	// IDC-IP UCloud
	// IDC-IP 百度云
	// IDC-IP 京东云
	// IDC-IP 青云
	// IDC-IP Aws
	// IDC-IP Azure
	// IDC-IP Google
	// 
	// 以上所有类型，FilterName为空时，表示不过滤，获取所有内容
	FilterName *string `json:"FilterName,omitnil" name:"FilterName"`

	// 目前支持的Action
	// "intercept" 拦截
	// "monitor"，监控
	// "permit" 放行
	// "redirect" 重定向
	// 
	// 尚未支持的Action
	// "captcha" 验证码
	FilterAction *string `json:"FilterAction,omitnil" name:"FilterAction"`

	// 过滤的IP
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`

	// 域名列表，为空表示查询AppID维度数据
	Domains []*string `json:"Domains,omitnil" name:"Domains"`
}

type DescribeScdnBotRecordsRequest struct {
	*tchttp.BaseRequest
	
	// BOT类型，取值为"UB","UCB","TCB"，分别表示：未知类型，自定义类型，公开类型
	BotType *string `json:"BotType,omitnil" name:"BotType"`

	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分页参数
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页参数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// mainland 大陆地区 overseas境外地区
	Area *string `json:"Area,omitnil" name:"Area"`

	// 排序参数
	SortBy []*BotSortBy `json:"SortBy,omitnil" name:"SortBy"`

	// BotType=UB时，表示需要过滤的预测标签，取值如下：
	//                 "crawler_unregular",
	//                 "crawler_regular",
	//                 "request_repeat",
	//                 "credential_miss_user",
	//                 "credential_without_user",
	//                 "credential_only_action",
	//                 "credential_user_password",
	//                 "credential_cracking",
	//                 "credential_stuffing",
	//                 "brush_sms",
	//                 "brush_captcha",
	//                 "reg_malicious"
	// BotType=TCB时，表示需要过滤的Bot分类，取值如下：
	//                 "Uncategorised",
	//                 "Search engine bot",
	//                 "Site monitor",
	//                 "Screenshot creator",
	//                 "Link checker",
	//                 "Web scraper",
	//                 "Vulnerability scanner",
	//                 "Virus scanner",
	//                 "Speed tester",
	//                 "Feed Fetcher",
	//                 "Tool",
	//                 "Marketing"
	// BotType=UCB时，取值如下：
	// User-Agent为空或不存在
	// User-Agent类型为BOT
	// User-Agent类型为HTTP Library
	// User-Agent类型为Framework
	// User-Agent类型为Tools
	// User-Agent类型为Unkonwn BOT
	// User-Agent类型为Scanner
	// Referer空或不存在
	// Referer滥用(多个UA使用相同Referer)
	// Cookie滥用(多个UA使用相同Cookie)
	// Cookie空或不存在
	// Connection空或不存在
	// Accept空或不存在
	// Accept-Language空或不存在
	// Accept-Enconding空或不存在
	// 使用HTTP HEAD方法
	// HTTP协议为1.0或者更低
	// IDC-IP 腾讯云
	// IDC-IP 阿里云
	// IDC-IP 华为云
	// IDC-IP 金山云
	// IDC-IP UCloud
	// IDC-IP 百度云
	// IDC-IP 京东云
	// IDC-IP 青云
	// IDC-IP Aws
	// IDC-IP Azure
	// IDC-IP Google
	// 
	// 以上所有类型，FilterName为空时，表示不过滤，获取所有内容
	FilterName *string `json:"FilterName,omitnil" name:"FilterName"`

	// 目前支持的Action
	// "intercept" 拦截
	// "monitor"，监控
	// "permit" 放行
	// "redirect" 重定向
	// 
	// 尚未支持的Action
	// "captcha" 验证码
	FilterAction *string `json:"FilterAction,omitnil" name:"FilterAction"`

	// 过滤的IP
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`

	// 域名列表，为空表示查询AppID维度数据
	Domains []*string `json:"Domains,omitnil" name:"Domains"`
}

func (r *DescribeScdnBotRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScdnBotRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotType")
	delete(f, "Domain")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Area")
	delete(f, "SortBy")
	delete(f, "FilterName")
	delete(f, "FilterAction")
	delete(f, "FilterIp")
	delete(f, "Domains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScdnBotRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScdnBotRecordsResponseParams struct {
	// BOT拦截结果数组
	Data []*BotRecord `json:"Data,omitnil" name:"Data"`

	// 记录数量
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeScdnBotRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScdnBotRecordsResponseParams `json:"Response"`
}

func (r *DescribeScdnBotRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScdnBotRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScdnConfigRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type DescribeScdnConfigRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *DescribeScdnConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScdnConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScdnConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScdnConfigResponseParams struct {
	// 自定义防护策略配置
	Acl *ScdnAclConfig `json:"Acl,omitnil" name:"Acl"`

	// Web 攻击防护（WAF）配置
	Waf *ScdnWafConfig `json:"Waf,omitnil" name:"Waf"`

	// CC 防护配置
	CC *ScdnConfig `json:"CC,omitnil" name:"CC"`

	// DDOS 防护配置
	Ddos *ScdnDdosConfig `json:"Ddos,omitnil" name:"Ddos"`

	// BOT 防护配置
	Bot *ScdnBotConfig `json:"Bot,omitnil" name:"Bot"`

	// 当前状态，取值online | offline
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeScdnConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScdnConfigResponseParams `json:"Response"`
}

func (r *DescribeScdnConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScdnConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScdnIpStrategyRequestParams struct {
	// 分页起始地址
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 列表分页记录条数，最大1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询条件过滤器
	Filters []*ScdnIpStrategyFilter `json:"Filters,omitnil" name:"Filters"`

	// 指定查询返回结果的排序字段，支持domain，update_time
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式，支持asc，desc
	Sequence *string `json:"Sequence,omitnil" name:"Sequence"`
}

type DescribeScdnIpStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 分页起始地址
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 列表分页记录条数，最大1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询条件过滤器
	Filters []*ScdnIpStrategyFilter `json:"Filters,omitnil" name:"Filters"`

	// 指定查询返回结果的排序字段，支持domain，update_time
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式，支持asc，desc
	Sequence *string `json:"Sequence,omitnil" name:"Sequence"`
}

func (r *DescribeScdnIpStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScdnIpStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "Sequence")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScdnIpStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScdnIpStrategyResponseParams struct {
	// IP策略列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpStrategyList []*ScdnIpStrategy `json:"IpStrategyList,omitnil" name:"IpStrategyList"`

	// 配置的策略条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeScdnIpStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScdnIpStrategyResponseParams `json:"Response"`
}

func (r *DescribeScdnIpStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScdnIpStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScdnTopDataRequestParams struct {
	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询的SCDN TOP攻击数据类型：
	// waf：Web 攻击防护TOP数据
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 排序对象，支持以下几种形式：
	// url：攻击目标 url 排序
	// ip：攻击源 IP 排序
	// attackType：攻击类型排序
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 排序使用的指标名称：
	// request：请求次数
	Filter *string `json:"Filter,omitnil" name:"Filter"`

	// 指定域名查询
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 指定攻击类型, 仅 Mode=waf 时有效
	// 不填则查询所有攻击类型的数据总和
	// AttackType 映射如下:
	//   other = '未知类型'
	//   malicious_scan = "恶意扫描"
	//   sql_inject = "SQL注入攻击"
	//   xss = "XSS攻击"
	//   cmd_inject = "命令注入攻击"
	//   ldap_inject = "LDAP注入攻击"
	//   ssi_inject = "SSI注入攻击"
	//   xml_inject = "XML注入攻击"
	//   web_service = "WEB服务漏洞攻击"
	//   web_app = "WEB应用漏洞攻击"
	//   path_traversal = "路径跨越攻击"
	//   illegal_access_core_file = "核心文件非法访问"
	//   trojan_horse = "木马后门攻击"
	//   csrf = "CSRF攻击"
	//   malicious_file_upload= '恶意文件上传'
	AttackType *string `json:"AttackType,omitnil" name:"AttackType"`

	// 指定防御模式,仅 Mode=waf 时有效
	// 不填则查询所有防御模式的数据总和
	// DefenceMode 映射如下：
	//   observe = '观察模式'
	//   intercept = '拦截模式'
	DefenceMode *string `json:"DefenceMode,omitnil" name:"DefenceMode"`
}

type DescribeScdnTopDataRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询的SCDN TOP攻击数据类型：
	// waf：Web 攻击防护TOP数据
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 排序对象，支持以下几种形式：
	// url：攻击目标 url 排序
	// ip：攻击源 IP 排序
	// attackType：攻击类型排序
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 排序使用的指标名称：
	// request：请求次数
	Filter *string `json:"Filter,omitnil" name:"Filter"`

	// 指定域名查询
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 指定攻击类型, 仅 Mode=waf 时有效
	// 不填则查询所有攻击类型的数据总和
	// AttackType 映射如下:
	//   other = '未知类型'
	//   malicious_scan = "恶意扫描"
	//   sql_inject = "SQL注入攻击"
	//   xss = "XSS攻击"
	//   cmd_inject = "命令注入攻击"
	//   ldap_inject = "LDAP注入攻击"
	//   ssi_inject = "SSI注入攻击"
	//   xml_inject = "XML注入攻击"
	//   web_service = "WEB服务漏洞攻击"
	//   web_app = "WEB应用漏洞攻击"
	//   path_traversal = "路径跨越攻击"
	//   illegal_access_core_file = "核心文件非法访问"
	//   trojan_horse = "木马后门攻击"
	//   csrf = "CSRF攻击"
	//   malicious_file_upload= '恶意文件上传'
	AttackType *string `json:"AttackType,omitnil" name:"AttackType"`

	// 指定防御模式,仅 Mode=waf 时有效
	// 不填则查询所有防御模式的数据总和
	// DefenceMode 映射如下：
	//   observe = '观察模式'
	//   intercept = '拦截模式'
	DefenceMode *string `json:"DefenceMode,omitnil" name:"DefenceMode"`
}

func (r *DescribeScdnTopDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScdnTopDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Mode")
	delete(f, "Metric")
	delete(f, "Filter")
	delete(f, "Domain")
	delete(f, "AttackType")
	delete(f, "DefenceMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScdnTopDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScdnTopDataResponseParams struct {
	// WAF 攻击类型统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopTypeData []*ScdnTypeData `json:"TopTypeData,omitnil" name:"TopTypeData"`

	// TOP 攻击源 IP 统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopIpData []*ScdnTopData `json:"TopIpData,omitnil" name:"TopIpData"`

	// 查询的SCDN类型，当前仅支持 waf
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// TOP URL 统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopUrlData []*ScdnTopUrlData `json:"TopUrlData,omitnil" name:"TopUrlData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeScdnTopDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScdnTopDataResponseParams `json:"Response"`
}

func (r *DescribeScdnTopDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScdnTopDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopDataRequestParams struct {
	// 查询起始日期：yyyy-MM-dd HH:mm:ss
	// 仅支持按天粒度的数据查询，取入参中的天信息作为起始日期
	// 返回大于等于起始日期当天 00:00:00 点产生的数据，如 StartTime为2018-09-04 10:40:00，返回数据的起始时间为2018-09-04 00:00:00
	// 仅支持 90 天内数据查询
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束日期：yyyy-MM-dd HH:mm:ss
	// 仅支持按天粒度的数据查询，取入参中的天信息作为结束日期
	// 返回小于等于结束日期当天 23:59:59 产生的数据，如EndTime为2018-09-05 22:40:00，返回数据的结束时间为2018-09-05 23:59:59
	// EndTime 需要大于等于 StartTime
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 排序对象，支持以下几种形式：
	// ip、ua_device、ua_browser、ua_os、referer
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 排序使用的指标名称：
	// flux：Metric 为 host 时指代访问流量
	// request：Metric 为 host 时指代访问请求数
	Filter *string `json:"Filter,omitnil" name:"Filter"`

	// 指定查询域名列表，最多可一次性查询 30 个加速域名明细
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 未填充域名情况下，指定项目查询，若填充了具体域名信息，以域名为主
	Project *int64 `json:"Project,omitnil" name:"Project"`

	// 是否详细显示每个域名的的具体数值
	Detail *bool `json:"Detail,omitnil" name:"Detail"`

	// 指定服务地域查询，不填充表示查询中国境内 CDN 数据
	// mainland：指定查询中国境内 CDN 数据
	// overseas：指定查询中国境外 CDN 数据
	Area *string `json:"Area,omitnil" name:"Area"`

	// 指定查询的产品数据，目前仅可使用cdn
	Product *string `json:"Product,omitnil" name:"Product"`
}

type DescribeTopDataRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始日期：yyyy-MM-dd HH:mm:ss
	// 仅支持按天粒度的数据查询，取入参中的天信息作为起始日期
	// 返回大于等于起始日期当天 00:00:00 点产生的数据，如 StartTime为2018-09-04 10:40:00，返回数据的起始时间为2018-09-04 00:00:00
	// 仅支持 90 天内数据查询
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束日期：yyyy-MM-dd HH:mm:ss
	// 仅支持按天粒度的数据查询，取入参中的天信息作为结束日期
	// 返回小于等于结束日期当天 23:59:59 产生的数据，如EndTime为2018-09-05 22:40:00，返回数据的结束时间为2018-09-05 23:59:59
	// EndTime 需要大于等于 StartTime
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 排序对象，支持以下几种形式：
	// ip、ua_device、ua_browser、ua_os、referer
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 排序使用的指标名称：
	// flux：Metric 为 host 时指代访问流量
	// request：Metric 为 host 时指代访问请求数
	Filter *string `json:"Filter,omitnil" name:"Filter"`

	// 指定查询域名列表，最多可一次性查询 30 个加速域名明细
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 未填充域名情况下，指定项目查询，若填充了具体域名信息，以域名为主
	Project *int64 `json:"Project,omitnil" name:"Project"`

	// 是否详细显示每个域名的的具体数值
	Detail *bool `json:"Detail,omitnil" name:"Detail"`

	// 指定服务地域查询，不填充表示查询中国境内 CDN 数据
	// mainland：指定查询中国境内 CDN 数据
	// overseas：指定查询中国境外 CDN 数据
	Area *string `json:"Area,omitnil" name:"Area"`

	// 指定查询的产品数据，目前仅可使用cdn
	Product *string `json:"Product,omitnil" name:"Product"`
}

func (r *DescribeTopDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Metric")
	delete(f, "Filter")
	delete(f, "Domains")
	delete(f, "Project")
	delete(f, "Detail")
	delete(f, "Area")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopDataResponseParams struct {
	// 各个资源的Top 访问数据详情。
	Data []*TopDataMore `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTopDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopDataResponseParams `json:"Response"`
}

func (r *DescribeTopDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrafficPackagesRequestParams struct {
	// 分页查询起始地址，默认 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询记录个数，默认100，最大1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 流量包排序方式，支持以下值：
	// expireTimeDesc：默认值，按过期时间倒序
	// expireTimeAsc：按过期时间正序
	// createTimeDesc：按创建时间倒序
	// createTimeAsc：按创建时间正序
	// status：按状态排序，正常抵扣>未生效>已用尽>已过期
	// channel：按来源排序，主动购买>自动续订>CDN赠送
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`
}

type DescribeTrafficPackagesRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询起始地址，默认 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询记录个数，默认100，最大1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 流量包排序方式，支持以下值：
	// expireTimeDesc：默认值，按过期时间倒序
	// expireTimeAsc：按过期时间正序
	// createTimeDesc：按创建时间倒序
	// createTimeAsc：按创建时间正序
	// status：按状态排序，正常抵扣>未生效>已用尽>已过期
	// channel：按来源排序，主动购买>自动续订>CDN赠送
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`
}

func (r *DescribeTrafficPackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrafficPackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SortBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrafficPackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrafficPackagesResponseParams struct {
	// 流量包总个数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 流量包详情
	TrafficPackages []*TrafficPackage `json:"TrafficPackages,omitnil" name:"TrafficPackages"`

	// 即将过期的流量包个数（7天内）
	ExpiringCount *int64 `json:"ExpiringCount,omitnil" name:"ExpiringCount"`

	// 有效流量包个数
	EnabledCount *int64 `json:"EnabledCount,omitnil" name:"EnabledCount"`

	// 付费流量包个数
	PaidCount *int64 `json:"PaidCount,omitnil" name:"PaidCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTrafficPackagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrafficPackagesResponseParams `json:"Response"`
}

func (r *DescribeTrafficPackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrafficPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUrlViolationsRequestParams struct {
	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目，默认为 100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 指定的域名查询
	Domains []*string `json:"Domains,omitnil" name:"Domains"`
}

type DescribeUrlViolationsRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目，默认为 100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 指定的域名查询
	Domains []*string `json:"Domains,omitnil" name:"Domains"`
}

func (r *DescribeUrlViolationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUrlViolationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Domains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUrlViolationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUrlViolationsResponseParams struct {
	// 违规 URL 详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	UrlRecordList []*ViolationUrl `json:"UrlRecordList,omitnil" name:"UrlRecordList"`

	// 记录总数，用于分页
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUrlViolationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUrlViolationsResponseParams `json:"Response"`
}

func (r *DescribeUrlViolationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUrlViolationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWafDataRequestParams struct {
	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 时间粒度，支持以下几种模式：
	// min：1 分钟粒度，指定查询区间 24 小时内（含 24 小时），可返回 1 分钟粒度明细数据
	// 5min：5 分钟粒度，指定查询区间 31 天内（含 31 天），可返回 5 分钟粒度明细数据
	// hour：1 小时粒度，指定查询区间 31 天内（含 31 天），可返回 1 小时粒度明细数据
	// day：天粒度，指定查询区间大于 31 天，可返回天粒度明细数据
	// 
	// 仅支持30天内数据查询，且查询时间范围在 7 到 30 天最小粒度是 hour。
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 指定域名查询
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 指定攻击类型
	// 不填则查询所有攻击类型的数据分布
	// AttackType 映射如下:
	// "webshell" : Webshell检测防护
	// "oa" : 常见OA漏洞防护
	// "xss" : XSS跨站脚本攻击防护
	// "xxe" : XXE攻击防护
	// "webscan" : 扫描器攻击漏洞防护
	// "cms" : 常见CMS漏洞防护
	// "upload" : 恶意文件上传攻击防护
	// "sql" : SQL注入攻击防护
	// "cmd_inject": 命令/代码注入攻击防护
	// "osc" : 开源组件漏洞防护
	// "file_read" : 任意文件读取
	// "ldap" : LDAP注入攻击防护
	// "other" : 其它漏洞防护
	AttackType *string `json:"AttackType,omitnil" name:"AttackType"`

	// 指定防御模式
	// 不填则查询所有防御模式的数据总和
	// DefenceMode映射如下：
	//   observe = '观察模式'
	//   intercept = '拦截模式'
	DefenceMode *string `json:"DefenceMode,omitnil" name:"DefenceMode"`

	// 地域：mainland 或 overseas
	Area *string `json:"Area,omitnil" name:"Area"`

	// 指定多个攻击类型，取值参考AttackType
	AttackTypes []*string `json:"AttackTypes,omitnil" name:"AttackTypes"`

	// 指定域名列表查询
	Domains []*string `json:"Domains,omitnil" name:"Domains"`
}

type DescribeWafDataRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 时间粒度，支持以下几种模式：
	// min：1 分钟粒度，指定查询区间 24 小时内（含 24 小时），可返回 1 分钟粒度明细数据
	// 5min：5 分钟粒度，指定查询区间 31 天内（含 31 天），可返回 5 分钟粒度明细数据
	// hour：1 小时粒度，指定查询区间 31 天内（含 31 天），可返回 1 小时粒度明细数据
	// day：天粒度，指定查询区间大于 31 天，可返回天粒度明细数据
	// 
	// 仅支持30天内数据查询，且查询时间范围在 7 到 30 天最小粒度是 hour。
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 指定域名查询
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 指定攻击类型
	// 不填则查询所有攻击类型的数据分布
	// AttackType 映射如下:
	// "webshell" : Webshell检测防护
	// "oa" : 常见OA漏洞防护
	// "xss" : XSS跨站脚本攻击防护
	// "xxe" : XXE攻击防护
	// "webscan" : 扫描器攻击漏洞防护
	// "cms" : 常见CMS漏洞防护
	// "upload" : 恶意文件上传攻击防护
	// "sql" : SQL注入攻击防护
	// "cmd_inject": 命令/代码注入攻击防护
	// "osc" : 开源组件漏洞防护
	// "file_read" : 任意文件读取
	// "ldap" : LDAP注入攻击防护
	// "other" : 其它漏洞防护
	AttackType *string `json:"AttackType,omitnil" name:"AttackType"`

	// 指定防御模式
	// 不填则查询所有防御模式的数据总和
	// DefenceMode映射如下：
	//   observe = '观察模式'
	//   intercept = '拦截模式'
	DefenceMode *string `json:"DefenceMode,omitnil" name:"DefenceMode"`

	// 地域：mainland 或 overseas
	Area *string `json:"Area,omitnil" name:"Area"`

	// 指定多个攻击类型，取值参考AttackType
	AttackTypes []*string `json:"AttackTypes,omitnil" name:"AttackTypes"`

	// 指定域名列表查询
	Domains []*string `json:"Domains,omitnil" name:"Domains"`
}

func (r *DescribeWafDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Interval")
	delete(f, "Domain")
	delete(f, "AttackType")
	delete(f, "DefenceMode")
	delete(f, "Area")
	delete(f, "AttackTypes")
	delete(f, "Domains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWafDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWafDataResponseParams struct {
	// 粒度数据
	Data []*TimestampData `json:"Data,omitnil" name:"Data"`

	// 粒度
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWafDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWafDataResponseParams `json:"Response"`
}

func (r *DescribeWafDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailDomain struct {
	// 域名 ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 腾讯云账号ID
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// 加速域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名对应的 CNAME 地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cname *string `json:"Cname,omitnil" name:"Cname"`

	// 加速服务状态
	// rejected：域名审核未通过，域名备案过期/被注销导致
	// processing：部署中
	// closing：关闭中
	// online：已启动
	// offline：已关闭
	Status *string `json:"Status,omitnil" name:"Status"`

	// 项目 ID，可前往腾讯云项目管理页面查看
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 加速域名业务类型
	// web：网页小文件
	// download：下载大文件
	// media：音视频点播
	// hybrid:  动静加速
	// dynamic:  动态加速
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 域名创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 域名更新时间
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 源站配置
	Origin *Origin `json:"Origin,omitnil" name:"Origin"`

	// IP 黑白名单配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpFilter *IpFilter `json:"IpFilter,omitnil" name:"IpFilter"`

	// IP 访问限频配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitnil" name:"IpFreqLimit"`

	// 状态码缓存配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitnil" name:"StatusCodeCache"`

	// 智能压缩配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compression *Compression `json:"Compression,omitnil" name:"Compression"`

	// 带宽封顶配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitnil" name:"BandwidthAlert"`

	// Range 回源配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitnil" name:"RangeOriginPull"`

	// 301/302 回源自动跟随配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitnil" name:"FollowRedirect"`

	// 自定义错误页面配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorPage *ErrorPage `json:"ErrorPage,omitnil" name:"ErrorPage"`

	// 自定义请求头部配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestHeader *RequestHeader `json:"RequestHeader,omitnil" name:"RequestHeader"`

	// 自定义响应头部配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitnil" name:"ResponseHeader"`

	// 单链接下行限速配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownstreamCapping *DownstreamCapping `json:"DownstreamCapping,omitnil" name:"DownstreamCapping"`

	// 带参/不带参缓存配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheKey *CacheKey `json:"CacheKey,omitnil" name:"CacheKey"`

	// 源站头部缓存配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseHeaderCache *ResponseHeaderCache `json:"ResponseHeaderCache,omitnil" name:"ResponseHeaderCache"`

	// 视频拖拽配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoSeek *VideoSeek `json:"VideoSeek,omitnil" name:"VideoSeek"`

	// 节点缓存过期规则配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cache *Cache `json:"Cache,omitnil" name:"Cache"`

	// 跨国链路优化配置（功能灰度中，敬请期待）
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginPullOptimization *OriginPullOptimization `json:"OriginPullOptimization,omitnil" name:"OriginPullOptimization"`

	// Https 加速相关配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Https *Https `json:"Https,omitnil" name:"Https"`

	// 时间戳防盗链配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Authentication *Authentication `json:"Authentication,omitnil" name:"Authentication"`

	// SEO 优化配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Seo *Seo `json:"Seo,omitnil" name:"Seo"`

	// 域名封禁状态
	// normal：正常状态
	// overdue：账号欠费导致域名关闭，充值完成后可自行启动加速服务
	// malicious：域名出现恶意行为，强制关闭加速服务
	// ddos：域名被大规模 DDoS 攻击，关闭加速服务
	// idle：域名超过 90 天内无任何操作、数据产生，判定为不活跃域名自动关闭加速服务，可自行启动加速服务
	// unlicensed：域名未备案/备案注销，自动关闭加速服务，备案完成后可自行启动加速服务
	// capping：触发配置的带宽阈值上限
	// readonly：域名存在特殊配置，被锁定
	// 注意：此字段可能返回 null，表示取不到有效值。
	Disable *string `json:"Disable,omitnil" name:"Disable"`

	// 访问协议强制跳转配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitnil" name:"ForceRedirect"`

	// Referer 防盗链配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Referer *Referer `json:"Referer,omitnil" name:"Referer"`

	// 浏览器缓存过期规则配置（功能灰度中，敬请期待）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAge *MaxAge `json:"MaxAge,omitnil" name:"MaxAge"`

	// Ipv6 回源配置（功能灰度中，敬请期待）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6 *Ipv6 `json:"Ipv6,omitnil" name:"Ipv6"`

	// 是否兼容旧版本配置（内部兼容性字段）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compatibility *Compatibility `json:"Compatibility,omitnil" name:"Compatibility"`

	// 区域特殊配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecificConfig *SpecificConfig `json:"SpecificConfig,omitnil" name:"SpecificConfig"`

	// 加速区域
	// mainland：中国境内加速
	// overseas：中国境外加速
	// global：全球加速
	// 注意：此字段可能返回 null，表示取不到有效值。
	Area *string `json:"Area,omitnil" name:"Area"`

	// 域名锁定状态
	// normal：未锁定
	// mainland：中国境内锁定
	// overseas：中国境外锁定
	// global：全球锁定
	// 注意：此字段可能返回 null，表示取不到有效值。
	Readonly *string `json:"Readonly,omitnil" name:"Readonly"`

	// 回源超时配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginPullTimeout *OriginPullTimeout `json:"OriginPullTimeout,omitnil" name:"OriginPullTimeout"`

	// 回源S3鉴权配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	AwsPrivateAccess *AwsPrivateAccess `json:"AwsPrivateAccess,omitnil" name:"AwsPrivateAccess"`

	// Scdn配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityConfig *SecurityConfig `json:"SecurityConfig,omitnil" name:"SecurityConfig"`

	// ImageOptimization配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageOptimization *ImageOptimization `json:"ImageOptimization,omitnil" name:"ImageOptimization"`

	// UA黑白名单配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAgentFilter *UserAgentFilter `json:"UserAgentFilter,omitnil" name:"UserAgentFilter"`

	// 访问控制
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessControl *AccessControl `json:"AccessControl,omitnil" name:"AccessControl"`

	// 是否支持高级配置项
	// on：支持
	// off：不支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	Advance *string `json:"Advance,omitnil" name:"Advance"`

	// URL重定向配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	UrlRedirect *UrlRedirect `json:"UrlRedirect,omitnil" name:"UrlRedirect"`

	// 访问端口配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessPort []*int64 `json:"AccessPort,omitnil" name:"AccessPort"`

	// 标签配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*Tag `json:"Tag,omitnil" name:"Tag"`

	// 时间戳防盗链高级配置，白名单功能
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedAuthentication *AdvancedAuthentication `json:"AdvancedAuthentication,omitnil" name:"AdvancedAuthentication"`

	// 回源鉴权高级配置，白名单功能
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginAuthentication *OriginAuthentication `json:"OriginAuthentication,omitnil" name:"OriginAuthentication"`

	// Ipv6访问配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6Access *Ipv6Access `json:"Ipv6Access,omitnil" name:"Ipv6Access"`

	// 高级配置集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvanceSet []*AdvanceConfig `json:"AdvanceSet,omitnil" name:"AdvanceSet"`

	// 离线缓存（功能灰度中，尚未全量，请等待后续全量发布）
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineCache *OfflineCache `json:"OfflineCache,omitnil" name:"OfflineCache"`

	// 合并回源（白名单功能）
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginCombine *OriginCombine `json:"OriginCombine,omitnil" name:"OriginCombine"`

	// POST上传配置项
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostMaxSize *PostSize `json:"PostMaxSize,omitnil" name:"PostMaxSize"`

	// Quic配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Quic *Quic `json:"Quic,omitnil" name:"Quic"`

	// 回源OSS私有鉴权
	// 注意：此字段可能返回 null，表示取不到有效值。
	OssPrivateAccess *OssPrivateAccess `json:"OssPrivateAccess,omitnil" name:"OssPrivateAccess"`

	// WebSocket配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebSocket *WebSocket `json:"WebSocket,omitnil" name:"WebSocket"`

	// 远程鉴权配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemoteAuthentication *RemoteAuthentication `json:"RemoteAuthentication,omitnil" name:"RemoteAuthentication"`

	// 共享CNAME配置（白名单功能）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShareCname *ShareCname `json:"ShareCname,omitnil" name:"ShareCname"`

	// 规则引擎
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleEngine *RuleEngine `json:"RuleEngine,omitnil" name:"RuleEngine"`

	// 主域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentHost *string `json:"ParentHost,omitnil" name:"ParentHost"`

	// 华为云对象存储回源鉴权
	// 注意：此字段可能返回 null，表示取不到有效值。
	HwPrivateAccess *HwPrivateAccess `json:"HwPrivateAccess,omitnil" name:"HwPrivateAccess"`

	// 七牛云对象存储回源鉴权
	// 注意：此字段可能返回 null，表示取不到有效值。
	QnPrivateAccess *QnPrivateAccess `json:"QnPrivateAccess,omitnil" name:"QnPrivateAccess"`

	// HTTPS服务，缺省时默认开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpsBilling *HttpsBilling `json:"HttpsBilling,omitnil" name:"HttpsBilling"`

	// 其他厂商对象存储回源鉴权
	// 注意：此字段可能返回 null，表示取不到有效值。
	OthersPrivateAccess *OthersPrivateAccess `json:"OthersPrivateAccess,omitnil" name:"OthersPrivateAccess"`
}

type DiagnoseData struct {
	// 诊断报告内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DiagnoseUnit `json:"Data,omitnil" name:"Data"`

	// 当前诊断项是否正常。
	// "ok"：正常
	// "error"：异常
	// "warning"："警告"
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`
}

type DiagnoseInfo struct {
	// 待诊断的URL。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiagnoseUrl *string `json:"DiagnoseUrl,omitnil" name:"DiagnoseUrl"`

	// 由系统生成的诊断链接。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiagnoseLink *string `json:"DiagnoseLink,omitnil" name:"DiagnoseLink"`

	// 诊断创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 诊断链接过期时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireDate *string `json:"ExpireDate,omitnil" name:"ExpireDate"`

	// 诊断链接当前访问次数，一个诊断链接最多可访问10次。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VisitCount *int64 `json:"VisitCount,omitnil" name:"VisitCount"`

	// 访问诊断链接的客户端简易信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientList []*DiagnoseList `json:"ClientList,omitnil" name:"ClientList"`

	// 域名加速区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Area *string `json:"Area,omitnil" name:"Area"`
}

type DiagnoseList struct {
	// 诊断任务标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiagnoseTag *string `json:"DiagnoseTag,omitnil" name:"DiagnoseTag"`

	// 报告ID，用于获取详细诊断报告。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportId *string `json:"ReportId,omitnil" name:"ReportId"`

	// 客户端信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientInfo []*ClientInfo `json:"ClientInfo,omitnil" name:"ClientInfo"`

	// 最终诊断结果。
	// -1：已提交
	// 0  ：检测中
	// 1  ：检测正常
	// 2  ： 检测异常
	// 3  ： 诊断页面异常关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinalDiagnose *int64 `json:"FinalDiagnose,omitnil" name:"FinalDiagnose"`

	// 诊断任务创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`
}

type DiagnoseUnit struct {
	// 内容单元英文名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil" name:"Key"`

	// 内容单元中文名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyText *string `json:"KeyText,omitnil" name:"KeyText"`

	// 报告内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 报告内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueText *string `json:"ValueText,omitnil" name:"ValueText"`
}

// Predefined struct for user
type DisableCachesRequestParams struct {
	// 禁用的 URL 列表（分协议生效，必须包含http://或https://）
	// 每次最多可提交 100 条，每日最多可提交 3000 条
	Urls []*string `json:"Urls,omitnil" name:"Urls"`
}

type DisableCachesRequest struct {
	*tchttp.BaseRequest
	
	// 禁用的 URL 列表（分协议生效，必须包含http://或https://）
	// 每次最多可提交 100 条，每日最多可提交 3000 条
	Urls []*string `json:"Urls,omitnil" name:"Urls"`
}

func (r *DisableCachesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableCachesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Urls")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableCachesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableCachesResponseParams struct {
	// 提交结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheOptResult *CacheOptResult `json:"CacheOptResult,omitnil" name:"CacheOptResult"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisableCachesResponse struct {
	*tchttp.BaseResponse
	Response *DisableCachesResponseParams `json:"Response"`
}

func (r *DisableCachesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableCachesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableClsLogTopicRequestParams struct {
	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil" name:"LogsetId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 接入渠道，cdn或者ecdn，默认值为cdn
	Channel *string `json:"Channel,omitnil" name:"Channel"`
}

type DisableClsLogTopicRequest struct {
	*tchttp.BaseRequest
	
	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil" name:"LogsetId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 接入渠道，cdn或者ecdn，默认值为cdn
	Channel *string `json:"Channel,omitnil" name:"Channel"`
}

func (r *DisableClsLogTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableClsLogTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "TopicId")
	delete(f, "Channel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableClsLogTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableClsLogTopicResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisableClsLogTopicResponse struct {
	*tchttp.BaseResponse
	Response *DisableClsLogTopicResponseParams `json:"Response"`
}

func (r *DisableClsLogTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableClsLogTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DistrictIspInfo struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 协议类型
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// IP协议类型
	IpProtocol *string `json:"IpProtocol,omitnil" name:"IpProtocol"`

	// 起始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 时间间隔，单位为分钟
	Interval *uint64 `json:"Interval,omitnil" name:"Interval"`

	// 指标名称
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 地区ID
	District *int64 `json:"District,omitnil" name:"District"`

	// 运营商ID
	Isp *int64 `json:"Isp,omitnil" name:"Isp"`

	// 指标数据点
	DataPoints []*uint64 `json:"DataPoints,omitnil" name:"DataPoints"`

	// 地区名称
	DistrictName *string `json:"DistrictName,omitnil" name:"DistrictName"`

	// 运营商名称
	IspName *string `json:"IspName,omitnil" name:"IspName"`
}

type DomainAreaConfig struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 地区列表，其中元素可为mainland/overseas
	Area []*string `json:"Area,omitnil" name:"Area"`
}

type DomainBotCount struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// BOT次数
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// Top指标值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 国家/地区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Country *string `json:"Country,omitnil" name:"Country"`

	// 省份
	// 注意：此字段可能返回 null，表示取不到有效值。
	Province *string `json:"Province,omitnil" name:"Province"`

	// 运营商
	// 注意：此字段可能返回 null，表示取不到有效值。
	Isp *string `json:"Isp,omitnil" name:"Isp"`
}

type DomainFilter struct {
	// 过滤字段名，支持的列表如下：
	// - origin：主源站。
	// - domain：域名。
	// - resourceId：域名id。
	// - status：域名状态，online，offline或processing。
	// - serviceType：业务类型，web，download，media，hybrid或dynamic。
	// - projectId：项目ID。
	// - domainType：主源站类型，cname表示自有源，cos表示cos接入，third_party表示第三方对象存储，igtm表示IGTM多活源。
	// - fullUrlCache：全路径缓存，on或off。
	// - https：是否配置https，on，off或processing。
	// - originPullProtocol：回源协议类型，支持http，follow或https。
	// - tagKey：标签键。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 过滤字段值。
	Value []*string `json:"Value,omitnil" name:"Value"`

	// 是否启用模糊查询，仅支持过滤字段名为origin，domain。
	// 模糊查询时，Value长度最大为1，否则Value长度最大为5。
	Fuzzy *bool `json:"Fuzzy,omitnil" name:"Fuzzy"`
}

type DomainLog struct {
	// 日志包起始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 日志包结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 日志包下载链接
	LogPath *string `json:"LogPath,omitnil" name:"LogPath"`

	// 日志包对应加速区域
	// mainland：境内
	// overseas：境外
	Area *string `json:"Area,omitnil" name:"Area"`

	// 日志包文件名
	LogName *string `json:"LogName,omitnil" name:"LogName"`

	// 文件大小，单位: Byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *int64 `json:"FileSize,omitnil" name:"FileSize"`
}

type DownstreamCapping struct {
	// 下行速度配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 下行限速规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	CappingRules []*CappingRule `json:"CappingRules,omitnil" name:"CappingRules"`
}

// Predefined struct for user
type DuplicateDomainConfigRequestParams struct {
	// 新增域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 被拷贝配置的域名
	ReferenceDomain *string `json:"ReferenceDomain,omitnil" name:"ReferenceDomain"`
}

type DuplicateDomainConfigRequest struct {
	*tchttp.BaseRequest
	
	// 新增域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 被拷贝配置的域名
	ReferenceDomain *string `json:"ReferenceDomain,omitnil" name:"ReferenceDomain"`
}

func (r *DuplicateDomainConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DuplicateDomainConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "ReferenceDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DuplicateDomainConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DuplicateDomainConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DuplicateDomainConfigResponse struct {
	*tchttp.BaseResponse
	Response *DuplicateDomainConfigResponseParams `json:"Response"`
}

func (r *DuplicateDomainConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DuplicateDomainConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EdgePackTaskFilter struct {
	// 过滤字段名
	// apk: apk名称
	// status: 母包处理进度 done, failed, processing
	Name *string `json:"Name,omitnil" name:"Name"`

	// 过滤字段值
	Value []*string `json:"Value,omitnil" name:"Value"`

	// 是否启用模糊查询，仅支持过滤字段名为 apk。
	// 模糊查询时，Value长度最大为1。
	Fuzzy *bool `json:"Fuzzy,omitnil" name:"Fuzzy"`
}

type EdgePackTaskStatus struct {
	// APK 名称
	Apk *string `json:"Apk,omitnil" name:"Apk"`

	// 输出目录
	DstDir *string `json:"DstDir,omitnil" name:"DstDir"`

	// 上传时间
	UploadTime *string `json:"UploadTime,omitnil" name:"UploadTime"`

	// 任务状态
	// created: 创建成功
	// processing: 处理中
	// done: 处理完成
	// failed: 处理失败
	Status *string `json:"Status,omitnil" name:"Status"`

	// 上传目录
	SrcDir []*string `json:"SrcDir,omitnil" name:"SrcDir"`

	// 失败任务状态详情
	StatusDesc *string `json:"StatusDesc,omitnil" name:"StatusDesc"`
}

// Predefined struct for user
type EnableCachesRequestParams struct {
	// 解封 URL 列表
	Urls []*string `json:"Urls,omitnil" name:"Urls"`

	// URL封禁日期
	Date *string `json:"Date,omitnil" name:"Date"`
}

type EnableCachesRequest struct {
	*tchttp.BaseRequest
	
	// 解封 URL 列表
	Urls []*string `json:"Urls,omitnil" name:"Urls"`

	// URL封禁日期
	Date *string `json:"Date,omitnil" name:"Date"`
}

func (r *EnableCachesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableCachesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Urls")
	delete(f, "Date")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableCachesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableCachesResponseParams struct {
	// 结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheOptResult *CacheOptResult `json:"CacheOptResult,omitnil" name:"CacheOptResult"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EnableCachesResponse struct {
	*tchttp.BaseResponse
	Response *EnableCachesResponseParams `json:"Response"`
}

func (r *EnableCachesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableCachesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableClsLogTopicRequestParams struct {
	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil" name:"LogsetId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 接入渠道，cdn或者ecdn，默认值为cdn
	Channel *string `json:"Channel,omitnil" name:"Channel"`
}

type EnableClsLogTopicRequest struct {
	*tchttp.BaseRequest
	
	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil" name:"LogsetId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 接入渠道，cdn或者ecdn，默认值为cdn
	Channel *string `json:"Channel,omitnil" name:"Channel"`
}

func (r *EnableClsLogTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableClsLogTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "TopicId")
	delete(f, "Channel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableClsLogTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableClsLogTopicResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EnableClsLogTopicResponse struct {
	*tchttp.BaseResponse
	Response *EnableClsLogTopicResponseParams `json:"Response"`
}

func (r *EnableClsLogTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableClsLogTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ErrorPage struct {
	// 状态码重定向配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 状态码重定向规则配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageRules []*ErrorPageRule `json:"PageRules,omitnil" name:"PageRules"`
}

type ErrorPageRule struct {
	// 状态码
	// 支持 400、403、404、500
	StatusCode *int64 `json:"StatusCode,omitnil" name:"StatusCode"`

	// 重定向状态码设置
	// 支持 301 或 302
	RedirectCode *int64 `json:"RedirectCode,omitnil" name:"RedirectCode"`

	// 重定向 URL
	// 需要为完整跳转路径，如 https://www.test.com/error.html
	RedirectUrl *string `json:"RedirectUrl,omitnil" name:"RedirectUrl"`
}

type EventLogStatsData struct {
	// 时间
	Datetime *string `json:"Datetime,omitnil" name:"Datetime"`

	// 请求数
	Request *uint64 `json:"Request,omitnil" name:"Request"`
}

type ExtraLogset struct {
	// 日志集信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Logset *LogSetInfo `json:"Logset,omitnil" name:"Logset"`

	// 日志主题信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Topics []*TopicInfo `json:"Topics,omitnil" name:"Topics"`
}

type FollowRedirect struct {
	// 回源跟随配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 自定义回源302 follow请求host配置，该功能为白名单功能，需要开启请联系腾讯云工程师。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectConfig *RedirectConfig `json:"RedirectConfig,omitnil" name:"RedirectConfig"`
}

type ForceRedirect struct {
	// 访问强制跳转配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 访问强制跳转类型
	// http：强制 http 跳转
	// https：强制 https 跳转
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectType *string `json:"RedirectType,omitnil" name:"RedirectType"`

	// 强制跳转时返回状态码 
	// 支持 301、302
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectStatusCode *int64 `json:"RedirectStatusCode,omitnil" name:"RedirectStatusCode"`

	// 强制跳转时是否返回增加的头部。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CarryHeaders *string `json:"CarryHeaders,omitnil" name:"CarryHeaders"`
}

// Predefined struct for user
type GetDisableRecordsRequestParams struct {
	// 指定 URL 查询
	Url *string `json:"Url,omitnil" name:"Url"`

	// 开始时间，如：2018-12-12 10:24:00。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间，如：2018-12-14 10:24:00。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// URL 当前状态
	// disable：当前仍为禁用状态，访问返回 403
	// enable：当前为可用状态，已解禁，可正常访问
	Status *string `json:"Status,omitnil" name:"Status"`

	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目，默认为20。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 任务ID，任务ID和起始时间需要至少填写一项。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

type GetDisableRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 指定 URL 查询
	Url *string `json:"Url,omitnil" name:"Url"`

	// 开始时间，如：2018-12-12 10:24:00。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间，如：2018-12-14 10:24:00。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// URL 当前状态
	// disable：当前仍为禁用状态，访问返回 403
	// enable：当前为可用状态，已解禁，可正常访问
	Status *string `json:"Status,omitnil" name:"Status"`

	// 分页查询偏移量，默认为 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目，默认为20。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 任务ID，任务ID和起始时间需要至少填写一项。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *GetDisableRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDisableRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Url")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Status")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDisableRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDisableRecordsResponseParams struct {
	// 封禁历史记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	UrlRecordList []*UrlRecord `json:"UrlRecordList,omitnil" name:"UrlRecordList"`

	// 任务总数，用于分页
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetDisableRecordsResponse struct {
	*tchttp.BaseResponse
	Response *GetDisableRecordsResponseParams `json:"Response"`
}

func (r *GetDisableRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDisableRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GuetzliAdapter struct {
	// 图片优化-GuetzliAdapter配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type HTTPHeader struct {
	// 请求头名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 请求头值
	Value *string `json:"Value,omitnil" name:"Value"`
}

type HeaderKey struct {
	// 组成Cachekey配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 组成CacheKey的header数组，';' 分割
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type HeuristicCache struct {
	// 启发式缓存配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 自定义启发式缓存时间配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheConfig *CacheConfig `json:"CacheConfig,omitnil" name:"CacheConfig"`
}

type Hsts struct {
	// HSTS 配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// MaxAge数值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAge *int64 `json:"MaxAge,omitnil" name:"MaxAge"`

	// 是否包含子域名，on或off。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncludeSubDomains *string `json:"IncludeSubDomains,omitnil" name:"IncludeSubDomains"`
}

type HttpHeaderPathRule struct {
	// http 头部设置方式
	// set：设置。变更指定头部参数的取值为设置后的值；若设置的头部不存在，则会增加该头部；若存在多个重复的头部参数，则会全部变更，同时合并为一个头部。
	// del：删除。删除指定的头部参数
	// add：增加。增加指定的头部参数，默认允许重复添加，即重复添加相同的头部（注：重复添加可能会影响浏览器响应，请优先使用set操作）
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderMode *string `json:"HeaderMode,omitnil" name:"HeaderMode"`

	// http 头部名称，最多可设置 100 个字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderName *string `json:"HeaderName,omitnil" name:"HeaderName"`

	// http 头部值，最多可设置 1000 个字符
	// Mode 为 del 时非必填
	// Mode 为 add/set 时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderValue *string `json:"HeaderValue,omitnil" name:"HeaderValue"`

	// 规则类型：
	// all：所有文件生效
	// file：指定文件后缀生效
	// directory：指定路径生效
	// path：指定绝对路径生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *string `json:"RuleType,omitnil" name:"RuleType"`

	// RuleType 对应类型下的匹配内容：
	// all 时填充 *
	// file 时填充后缀名，如 jpg、txt
	// directory 时填充路径，如 /xxx/test/
	// path 时填充绝对路径，如 /xxx/test.html
	// 注意：此字段可能返回 null，表示取不到有效值。
	RulePaths []*string `json:"RulePaths,omitnil" name:"RulePaths"`
}

type HttpHeaderRule struct {
	// http头部设置方式，支持add，set或del，分别表示新增，设置或删除头部。
	HeaderMode *string `json:"HeaderMode,omitnil" name:"HeaderMode"`

	// http头部名称。
	HeaderName *string `json:"HeaderName,omitnil" name:"HeaderName"`

	// http头部值。
	HeaderValue *string `json:"HeaderValue,omitnil" name:"HeaderValue"`
}

type Https struct {
	// https 配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// http2 配置开关
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Http2 *string `json:"Http2,omitnil" name:"Http2"`

	// OCSP 配置开关
	// on：开启
	// off：关闭
	// 默认为关闭状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcspStapling *string `json:"OcspStapling,omitnil" name:"OcspStapling"`

	// 客户端证书校验功能
	// on：开启
	// off：关闭
	// 默认为关闭状态，开启时需要上传客户端证书信息，该配置项目前在灰度中，尚未全量
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyClient *string `json:"VerifyClient,omitnil" name:"VerifyClient"`

	// 服务端证书配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertInfo *ServerCert `json:"CertInfo,omitnil" name:"CertInfo"`

	// 客户端证书配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientCertInfo *ClientCert `json:"ClientCertInfo,omitnil" name:"ClientCertInfo"`

	// Spdy 配置开关
	// on：开启
	// off：关闭
	// 默认为关闭状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Spdy *string `json:"Spdy,omitnil" name:"Spdy"`

	// https 证书部署状态
	// closed：已关闭
	// deploying：部署中
	// deployed：部署成功
	// failed：部署失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	SslStatus *string `json:"SslStatus,omitnil" name:"SslStatus"`

	// Hsts配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hsts *Hsts `json:"Hsts,omitnil" name:"Hsts"`

	// Tls版本设置，仅支持部分Advance域名，支持设置 TLSv1, TLSv1.1, TLSv1.2, TLSv1.3，修改时必须开启连续的版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	TlsVersion []*string `json:"TlsVersion,omitnil" name:"TlsVersion"`
}

type HttpsBilling struct {
	// HTTPS服务配置开关，取值有：
	// on：开启，缺省时默认开启，会产生计费
	// off：关闭，拦截https请求
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type HttpsPackage struct {
	// HTTPS请求包 Id
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// HTTPS请求包类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// HTTPS请求包大小（单位为：次）
	Size *int64 `json:"Size,omitnil" name:"Size"`

	// 已消耗HTTPS请求包（单位为：次）
	SizeUsed *int64 `json:"SizeUsed,omitnil" name:"SizeUsed"`

	// HTTPS请求包状态
	// enabled：已启用
	// expired：已过期
	// disabled：未启用
	Status *string `json:"Status,omitnil" name:"Status"`

	// HTTPS请求包发放时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// HTTPS请求包生效时间
	EnableTime *string `json:"EnableTime,omitnil" name:"EnableTime"`

	// HTTPS请求包过期时间
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// HTTPS请求包来源
	Channel *string `json:"Channel,omitnil" name:"Channel"`

	// HTTPS请求包生命周期月数
	LifeTimeMonth *int64 `json:"LifeTimeMonth,omitnil" name:"LifeTimeMonth"`

	// HTTPS请求包是否支持退费
	RefundAvailable *bool `json:"RefundAvailable,omitnil" name:"RefundAvailable"`

	// HTTPS请求包类型id
	ConfigId *int64 `json:"ConfigId,omitnil" name:"ConfigId"`

	// HTTPS请求包实际生效时间
	TrueEnableTime *string `json:"TrueEnableTime,omitnil" name:"TrueEnableTime"`

	// HTTPS请求包实际过期时间
	TrueExpireTime *string `json:"TrueExpireTime,omitnil" name:"TrueExpireTime"`

	// HTTPS请求包生效区域 
	// global：全球
	Area *string `json:"Area,omitnil" name:"Area"`

	// HTTPS请求包是否续订
	ContractExtension *bool `json:"ContractExtension,omitnil" name:"ContractExtension"`

	// HTTPS请求包是否支持续订
	ExtensionAvailable *bool `json:"ExtensionAvailable,omitnil" name:"ExtensionAvailable"`

	// HTTPS请求包当前续订模式
	// 0：未续订
	// 1：到期续订
	// 2：用完续订
	// 3：到期或用完续订
	ExtensionMode *uint64 `json:"ExtensionMode,omitnil" name:"ExtensionMode"`

	// HTTPS请求包是否自动续订
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoExtension *bool `json:"AutoExtension,omitnil" name:"AutoExtension"`
}

type HwPrivateAccess struct {
	//  华为云对象存储回源鉴权配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 访问 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessKey *string `json:"AccessKey,omitnil" name:"AccessKey"`

	// 密钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitnil" name:"SecretKey"`

	// bucketname
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitnil" name:"Bucket"`
}

type ImageOptimization struct {
	// WebpAdapter配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebpAdapter *WebpAdapter `json:"WebpAdapter,omitnil" name:"WebpAdapter"`

	// TpgAdapter配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TpgAdapter *TpgAdapter `json:"TpgAdapter,omitnil" name:"TpgAdapter"`

	// GuetzliAdapter配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	GuetzliAdapter *GuetzliAdapter `json:"GuetzliAdapter,omitnil" name:"GuetzliAdapter"`

	// AvifAdapter配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvifAdapter *AvifAdapter `json:"AvifAdapter,omitnil" name:"AvifAdapter"`
}

type IpFilter struct {
	// IP 黑白名单配置开关，取值有
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// IP 黑白名单类型
	// whitelist：白名单
	// blacklist：黑名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterType *string `json:"FilterType,omitnil" name:"FilterType"`

	// IP 黑白名单列表
	// 支持 X.X.X.X 形式 IP，或 /8、 /16、/24 形式网段
	// 最多可填充 50 个白名单或 50 个黑名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	Filters []*string `json:"Filters,omitnil" name:"Filters"`

	// IP 黑白名单分路径配置，白名单功能
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterRules []*IpFilterPathRule `json:"FilterRules,omitnil" name:"FilterRules"`

	// IP 黑白名单验证失败时返回的 code <br><font color=red>已下线，参数失效，不支持自定义状态码，固定返回514</font>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnCode *int64 `json:"ReturnCode,omitnil" name:"ReturnCode"`
}

type IpFilterPathRule struct {
	// IP 黑白名单类型
	// whitelist：白名单
	// blacklist：黑名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterType *string `json:"FilterType,omitnil" name:"FilterType"`

	// IP 黑白名单列表
	// 支持 X.X.X.X 形式 IP，或 /8、 /16、/24 形式网段
	// 最多可填充 50 个白名单或 50 个黑名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	Filters []*string `json:"Filters,omitnil" name:"Filters"`

	// 规则类型：
	// all：所有文件生效
	// file：指定文件后缀生效
	// directory：指定路径生效
	// path：指定绝对路径生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *string `json:"RuleType,omitnil" name:"RuleType"`

	// RuleType 对应类型下的匹配内容：
	// all 时填充 *
	// file 时填充后缀名，如 jpg、txt
	// directory 时填充路径，如 /xxx/test/
	// path 时填充绝对路径，如 /xxx/test.html
	// 注意：此字段可能返回 null，表示取不到有效值。
	RulePaths []*string `json:"RulePaths,omitnil" name:"RulePaths"`
}

type IpFreqLimit struct {
	// IP 限频配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 设置每秒请求数限制
	// 超出限制的请求会直接返回 514
	// 注意：此字段可能返回 null，表示取不到有效值。
	Qps *int64 `json:"Qps,omitnil" name:"Qps"`
}

type IpStatus struct {
	// 节点 IP
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 节点所属区域
	District *string `json:"District,omitnil" name:"District"`

	// 节点所属运营商
	Isp *string `json:"Isp,omitnil" name:"Isp"`

	// 节点所在城市
	City *string `json:"City,omitnil" name:"City"`

	// 节点状态
	// online：上线状态，正常调度服务中
	// offline：下线状态
	Status *string `json:"Status,omitnil" name:"Status"`

	// 节点 IPV6
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6 *string `json:"Ipv6,omitnil" name:"Ipv6"`
}

type Ipv6 struct {
	// 域名开启源站ipv6配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type Ipv6Access struct {
	// 域名开启ipv6访问配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type KeyRule struct {
	// CacheType 对应类型下的匹配内容：
	// file 时填充后缀名，如 jpg、txt
	// directory 时填充路径，如 /xxx/test
	// path 时填充绝对路径，如 /xxx/test.html
	// index 时填充 /
	// 注意：此字段可能返回 null，表示取不到有效值。
	RulePaths []*string `json:"RulePaths,omitnil" name:"RulePaths"`

	// 规则类型：
	// file：指定文件后缀生效
	// directory：指定路径生效
	// path：指定绝对路径生效
	// index：首页
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *string `json:"RuleType,omitnil" name:"RuleType"`

	// 是否开启全路径缓存
	// on：开启全路径缓存（即关闭参数忽略）
	// off：关闭全路径缓存（即开启参数忽略）
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullUrlCache *string `json:"FullUrlCache,omitnil" name:"FullUrlCache"`

	// 是否忽略大小写缓存
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreCase *string `json:"IgnoreCase,omitnil" name:"IgnoreCase"`

	// CacheKey中包含请求参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryString *RuleQueryString `json:"QueryString,omitnil" name:"QueryString"`

	// 路径缓存键标签，传 user
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleTag *string `json:"RuleTag,omitnil" name:"RuleTag"`
}

// Predefined struct for user
type ListClsLogTopicsRequestParams struct {
	// 接入渠道，cdn或者ecdn，默认值为cdn
	Channel *string `json:"Channel,omitnil" name:"Channel"`
}

type ListClsLogTopicsRequest struct {
	*tchttp.BaseRequest
	
	// 接入渠道，cdn或者ecdn，默认值为cdn
	Channel *string `json:"Channel,omitnil" name:"Channel"`
}

func (r *ListClsLogTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListClsLogTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Channel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListClsLogTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListClsLogTopicsResponseParams struct {
	// 上海区域日志集信息
	Logset *LogSetInfo `json:"Logset,omitnil" name:"Logset"`

	// 上海区域日志主题信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Topics []*TopicInfo `json:"Topics,omitnil" name:"Topics"`

	// 其他区域日志集信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraLogset []*ExtraLogset `json:"ExtraLogset,omitnil" name:"ExtraLogset"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListClsLogTopicsResponse struct {
	*tchttp.BaseResponse
	Response *ListClsLogTopicsResponseParams `json:"Response"`
}

func (r *ListClsLogTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListClsLogTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListClsTopicDomainsRequestParams struct {
	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil" name:"LogsetId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 接入渠道，cdn或者ecdn，默认值为cdn
	Channel *string `json:"Channel,omitnil" name:"Channel"`
}

type ListClsTopicDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil" name:"LogsetId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 接入渠道，cdn或者ecdn，默认值为cdn
	Channel *string `json:"Channel,omitnil" name:"Channel"`
}

func (r *ListClsTopicDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListClsTopicDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "TopicId")
	delete(f, "Channel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListClsTopicDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListClsTopicDomainsResponseParams struct {
	// 开发者ID
	AppId *uint64 `json:"AppId,omitnil" name:"AppId"`

	// 渠道
	Channel *string `json:"Channel,omitnil" name:"Channel"`

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil" name:"LogsetId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 域名区域配置，其中可能含有已删除的域名，如果要再传回ManageClsTopicDomains接口，需要结合ListCdnDomains接口排除掉已删除的域名
	DomainAreaConfigs []*DomainAreaConfig `json:"DomainAreaConfigs,omitnil" name:"DomainAreaConfigs"`

	// 日志主题名称
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// 日志主题最近更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListClsTopicDomainsResponse struct {
	*tchttp.BaseResponse
	Response *ListClsTopicDomainsResponseParams `json:"Response"`
}

func (r *ListClsTopicDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListClsTopicDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDiagnoseReportRequestParams struct {
	// 用于搜索诊断URL的关键字，不填时返回用户所有的诊断任务。
	KeyWords *string `json:"KeyWords,omitnil" name:"KeyWords"`

	// 用于搜索诊断系统返回的诊断链接，形如：http://cdn.cloud.tencent.com/self_diagnose/xxxxx
	DiagnoseLink *string `json:"DiagnoseLink,omitnil" name:"DiagnoseLink"`

	// 请求源带协议头，形如：https://console.cloud.tencent.com
	Origin *string `json:"Origin,omitnil" name:"Origin"`
}

type ListDiagnoseReportRequest struct {
	*tchttp.BaseRequest
	
	// 用于搜索诊断URL的关键字，不填时返回用户所有的诊断任务。
	KeyWords *string `json:"KeyWords,omitnil" name:"KeyWords"`

	// 用于搜索诊断系统返回的诊断链接，形如：http://cdn.cloud.tencent.com/self_diagnose/xxxxx
	DiagnoseLink *string `json:"DiagnoseLink,omitnil" name:"DiagnoseLink"`

	// 请求源带协议头，形如：https://console.cloud.tencent.com
	Origin *string `json:"Origin,omitnil" name:"Origin"`
}

func (r *ListDiagnoseReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDiagnoseReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KeyWords")
	delete(f, "DiagnoseLink")
	delete(f, "Origin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDiagnoseReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDiagnoseReportResponseParams struct {
	// 诊断信息。
	Data []*DiagnoseInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListDiagnoseReportResponse struct {
	*tchttp.BaseResponse
	Response *ListDiagnoseReportResponseParams `json:"Response"`
}

func (r *ListDiagnoseReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDiagnoseReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListScdnDomainsRequestParams struct {
	// 分页起始地址
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 列表分页记录条数，最大1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 域名信息
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type ListScdnDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 分页起始地址
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 列表分页记录条数，最大1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 域名信息
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *ListScdnDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListScdnDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListScdnDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListScdnDomainsResponseParams struct {
	// 域名列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainList []*ScdnDomain `json:"DomainList,omitnil" name:"DomainList"`

	// 域名的总条数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListScdnDomainsResponse struct {
	*tchttp.BaseResponse
	Response *ListScdnDomainsResponseParams `json:"Response"`
}

func (r *ListScdnDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListScdnDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListScdnLogTasksRequestParams struct {
	// 产品来源 cdn/ecdn
	Source *string `json:"Source,omitnil" name:"Source"`

	// 地域：mainland 或 overseas 为空表示查询所有地域
	Area *string `json:"Area,omitnil" name:"Area"`
}

type ListScdnLogTasksRequest struct {
	*tchttp.BaseRequest
	
	// 产品来源 cdn/ecdn
	Source *string `json:"Source,omitnil" name:"Source"`

	// 地域：mainland 或 overseas 为空表示查询所有地域
	Area *string `json:"Area,omitnil" name:"Area"`
}

func (r *ListScdnLogTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListScdnLogTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Source")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListScdnLogTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListScdnLogTasksResponseParams struct {
	// 日志下载任务详情
	TaskList []*ScdnLogTaskDetail `json:"TaskList,omitnil" name:"TaskList"`

	// 查询到的下载任务的总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListScdnLogTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListScdnLogTasksResponseParams `json:"Response"`
}

func (r *ListScdnLogTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListScdnLogTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListScdnTopBotDataRequestParams struct {
	// 获取Top量，取值范围[1-10]
	TopCount *int64 `json:"TopCount,omitnil" name:"TopCount"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// mainland 大陆地区 overseas境外地区
	Area *string `json:"Area,omitnil" name:"Area"`

	// session表示查询BOT会话的Top信息
	// ip表示查询BOT客户端IP的Top信息
	// 
	// 不填代表获取会话信息
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 域名，仅当Metric=ip，并且Domain为空时有效，不填写表示获取AppID信息
	Domains []*string `json:"Domains,omitnil" name:"Domains"`
}

type ListScdnTopBotDataRequest struct {
	*tchttp.BaseRequest
	
	// 获取Top量，取值范围[1-10]
	TopCount *int64 `json:"TopCount,omitnil" name:"TopCount"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// mainland 大陆地区 overseas境外地区
	Area *string `json:"Area,omitnil" name:"Area"`

	// session表示查询BOT会话的Top信息
	// ip表示查询BOT客户端IP的Top信息
	// 
	// 不填代表获取会话信息
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 域名，仅当Metric=ip，并且Domain为空时有效，不填写表示获取AppID信息
	Domains []*string `json:"Domains,omitnil" name:"Domains"`
}

func (r *ListScdnTopBotDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListScdnTopBotDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopCount")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Area")
	delete(f, "Metric")
	delete(f, "Domains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListScdnTopBotDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListScdnTopBotDataResponseParams struct {
	// 域名BOT次数列表
	Data []*BotStatisticsCount `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListScdnTopBotDataResponse struct {
	*tchttp.BaseResponse
	Response *ListScdnTopBotDataResponseParams `json:"Response"`
}

func (r *ListScdnTopBotDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListScdnTopBotDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTopBotDataRequestParams struct {
	// 获取Top量，取值范围[1-10]
	TopCount *int64 `json:"TopCount,omitnil" name:"TopCount"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// session表示查询BOT会话的Top信息
	// ip表示查询BOT客户端IP的Top信息
	// 
	// 不填代表获取会话信息
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 域名，仅当Metric=ip时有效，不填写表示使用Domains参数
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名，仅当Metric=ip，并且Domain为空时有效，不填写表示获取AppID信息
	Domains []*string `json:"Domains,omitnil" name:"Domains"`
}

type ListTopBotDataRequest struct {
	*tchttp.BaseRequest
	
	// 获取Top量，取值范围[1-10]
	TopCount *int64 `json:"TopCount,omitnil" name:"TopCount"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// session表示查询BOT会话的Top信息
	// ip表示查询BOT客户端IP的Top信息
	// 
	// 不填代表获取会话信息
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 域名，仅当Metric=ip时有效，不填写表示使用Domains参数
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名，仅当Metric=ip，并且Domain为空时有效，不填写表示获取AppID信息
	Domains []*string `json:"Domains,omitnil" name:"Domains"`
}

func (r *ListTopBotDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTopBotDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopCount")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Metric")
	delete(f, "Domain")
	delete(f, "Domains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTopBotDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTopBotDataResponseParams struct {
	// 域名BOT次数列表
	Data []*DomainBotCount `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListTopBotDataResponse struct {
	*tchttp.BaseResponse
	Response *ListTopBotDataResponseParams `json:"Response"`
}

func (r *ListTopBotDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTopBotDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTopCcDataRequestParams struct {
	// 查询Top数据的开始时间，格式为：2020-01-01 00:00:00
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询Top数据的结束时间，格式为：2020-01-01 23:59:59
	// 支持 90 天内数据查询，不传此参数，表示查当天数据
	// 时间跨度要小于等于7天
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 域名，不传此参数，表示查询账号级别数据
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 统计指标：
	// ip_url : Top IP+URL 默认值
	// ua :  Top UA
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// cdn表示CDN数据，默认值
	// ecdn表示ECDN数据
	Source *string `json:"Source,omitnil" name:"Source"`

	// 域名列表，不传此参数，表示查询账号级别数据
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 执行动作，取值为：intercept/redirect/observe
	// 分别表示：拦截/重定向/观察
	// 为空表示查询所有执行动作数据
	ActionName *string `json:"ActionName,omitnil" name:"ActionName"`

	// 地域：mainland或overseas，表示国内或海外，不填写默认表示国内
	Area *string `json:"Area,omitnil" name:"Area"`
}

type ListTopCcDataRequest struct {
	*tchttp.BaseRequest
	
	// 查询Top数据的开始时间，格式为：2020-01-01 00:00:00
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询Top数据的结束时间，格式为：2020-01-01 23:59:59
	// 支持 90 天内数据查询，不传此参数，表示查当天数据
	// 时间跨度要小于等于7天
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 域名，不传此参数，表示查询账号级别数据
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 统计指标：
	// ip_url : Top IP+URL 默认值
	// ua :  Top UA
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// cdn表示CDN数据，默认值
	// ecdn表示ECDN数据
	Source *string `json:"Source,omitnil" name:"Source"`

	// 域名列表，不传此参数，表示查询账号级别数据
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 执行动作，取值为：intercept/redirect/observe
	// 分别表示：拦截/重定向/观察
	// 为空表示查询所有执行动作数据
	ActionName *string `json:"ActionName,omitnil" name:"ActionName"`

	// 地域：mainland或overseas，表示国内或海外，不填写默认表示国内
	Area *string `json:"Area,omitnil" name:"Area"`
}

func (r *ListTopCcDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTopCcDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Domain")
	delete(f, "Metric")
	delete(f, "Source")
	delete(f, "Domains")
	delete(f, "ActionName")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTopCcDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTopCcDataResponseParams struct {
	// Top数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*CcTopData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListTopCcDataResponse struct {
	*tchttp.BaseResponse
	Response *ListTopCcDataResponseParams `json:"Response"`
}

func (r *ListTopCcDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTopCcDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTopClsLogDataRequestParams struct {
	// 需要查询的日志集ID
	LogsetId *string `json:"LogsetId,omitnil" name:"LogsetId"`

	// 需要查询的日志主题ID组合，多个以逗号分隔
	TopicIds *string `json:"TopicIds,omitnil" name:"TopicIds"`

	// 需要查询的日志的起始时间，格式 YYYY-mm-dd HH:MM:SS
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 需要查询的日志的结束时间，格式 YYYY-mm-dd HH:MM:SS，时间跨度应小于10分钟
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 指定域名查询
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 指定访问的URL查询，支持URL开头和结尾使用\*表示通配
	// 如：
	// /files/* 表示所有以/files/开头的请求
	// *.jpg 表示所有以.jpg结尾的请求
	Url *string `json:"Url,omitnil" name:"Url"`

	// 接入渠道，cdn或者ecdn，默认值为cdn
	Channel *string `json:"Channel,omitnil" name:"Channel"`

	// 要查询的Top条数，最大值为100，默认为10
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 按请求量排序， asc（升序）或者 desc（降序），默认为 desc
	Sort *string `json:"Sort,omitnil" name:"Sort"`
}

type ListTopClsLogDataRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的日志集ID
	LogsetId *string `json:"LogsetId,omitnil" name:"LogsetId"`

	// 需要查询的日志主题ID组合，多个以逗号分隔
	TopicIds *string `json:"TopicIds,omitnil" name:"TopicIds"`

	// 需要查询的日志的起始时间，格式 YYYY-mm-dd HH:MM:SS
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 需要查询的日志的结束时间，格式 YYYY-mm-dd HH:MM:SS，时间跨度应小于10分钟
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 指定域名查询
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 指定访问的URL查询，支持URL开头和结尾使用\*表示通配
	// 如：
	// /files/* 表示所有以/files/开头的请求
	// *.jpg 表示所有以.jpg结尾的请求
	Url *string `json:"Url,omitnil" name:"Url"`

	// 接入渠道，cdn或者ecdn，默认值为cdn
	Channel *string `json:"Channel,omitnil" name:"Channel"`

	// 要查询的Top条数，最大值为100，默认为10
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 按请求量排序， asc（升序）或者 desc（降序），默认为 desc
	Sort *string `json:"Sort,omitnil" name:"Sort"`
}

func (r *ListTopClsLogDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTopClsLogDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "TopicIds")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Domain")
	delete(f, "Url")
	delete(f, "Channel")
	delete(f, "Limit")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTopClsLogDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTopClsLogDataResponseParams struct {
	// 数据列表
	Data []*ClsLogIpData `json:"Data,omitnil" name:"Data"`

	// 获取到Top总记录数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 获取到的不重复IP条数
	IpCount *uint64 `json:"IpCount,omitnil" name:"IpCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListTopClsLogDataResponse struct {
	*tchttp.BaseResponse
	Response *ListTopClsLogDataResponseParams `json:"Response"`
}

func (r *ListTopClsLogDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTopClsLogDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTopDDoSDataRequestParams struct {
	// 查询Top数据的开始时间，格式为：2020-01-01 00:00:00
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询Top数据的结束时间，格式为：2020-01-01 23:59:59
	// 支持 90 天内数据查询，时间跨度要小于等于7天
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询Top的数量，不填默认值为10
	TopCount *uint64 `json:"TopCount,omitnil" name:"TopCount"`

	// AttackIP表示查询攻击ip的top排行，AttackType表示攻击类型的top排行，为空默认为AttackType
	Metric *string `json:"Metric,omitnil" name:"Metric"`
}

type ListTopDDoSDataRequest struct {
	*tchttp.BaseRequest
	
	// 查询Top数据的开始时间，格式为：2020-01-01 00:00:00
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询Top数据的结束时间，格式为：2020-01-01 23:59:59
	// 支持 90 天内数据查询，时间跨度要小于等于7天
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询Top的数量，不填默认值为10
	TopCount *uint64 `json:"TopCount,omitnil" name:"TopCount"`

	// AttackIP表示查询攻击ip的top排行，AttackType表示攻击类型的top排行，为空默认为AttackType
	Metric *string `json:"Metric,omitnil" name:"Metric"`
}

func (r *ListTopDDoSDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTopDDoSDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TopCount")
	delete(f, "Metric")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTopDDoSDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTopDDoSDataResponseParams struct {
	// DDoS 攻击类型的top数据，当Metric=AttackType的时候返回攻击类型的统计数据，IPData为空
	Data []*DDoSTopData `json:"Data,omitnil" name:"Data"`

	// ddos攻击ip的top数据，Metric=AttackIP的时候返回IPData，Data为空
	IPData []*DDoSAttackIPTopData `json:"IPData,omitnil" name:"IPData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListTopDDoSDataResponse struct {
	*tchttp.BaseResponse
	Response *ListTopDDoSDataResponseParams `json:"Response"`
}

func (r *ListTopDDoSDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTopDDoSDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTopDataRequestParams struct {
	// 查询起始时间：yyyy-MM-dd HH:mm:ss
	// 仅支持按分钟粒度的数据查询，按入参抹去秒位作为起始时间，如 StartTime为2018-09-04 10:40:23，返回数据的起始时间为2018-09-04 10:40:00
	// 仅支持 90 天内数据查询
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间：yyyy-MM-dd HH:mm:ss
	// 仅支持按天粒度的数据查询，取入参中的天信息作为结束日期 返回小于等于结束日期当天 23:59:59 产生的数据，如EndTime为2018-09-05 22:40:00，返回数据的结束时间为2018-09-05 23:59:59
	// EndTime 需要大于等于 StartTime
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 排序对象，支持以下几种形式：
	// url：访问 URL 排序（无参数的URL），支持的 Filter 为 flux、request
	// district：省份、国家/地区排序，支持的 Filter 为 flux、request
	// isp：运营商排序，支持的 Filter 为 flux、request
	// host：域名访问数据排序，支持的 Filter 为：flux、request、bandwidth、fluxHitRate、2XX、3XX、4XX、5XX、statusCode
	// originHost：域名回源数据排序，支持的 Filter 为 flux、request、bandwidth、origin_2XX、origin_3XX、origin_4XX、origin_5XX、OriginStatusCode
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 排序使用的指标名称：
	// flux：Metric 为 host 时指代访问流量，originHost 时指代回源流量
	// bandwidth：Metric 为 host 时指代访问带宽，originHost 时指代回源带宽
	// request：Metric 为 host 时指代访问请求数，originHost 时指代回源请求数
	// fluxHitRate：平均流量命中率
	// 2XX：访问 2XX 状态码
	// 3XX：访问 3XX 状态码
	// 4XX：访问 4XX 状态码
	// 5XX：访问 5XX 状态码
	// origin_2XX：回源 2XX 状态码
	// origin_3XX：回源 3XX 状态码
	// origin_4XX：回源 4XX 状态码
	// origin_5XX：回源 5XX 状态码
	// statusCode：指定访问状态码统计，在 Code 参数中填充指定状态码
	// OriginStatusCode：指定回源状态码统计，在 Code 参数中填充指定状态码
	Filter *string `json:"Filter,omitnil" name:"Filter"`

	// 指定查询域名列表，最多可一次性查询 30 个加速域名明细
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 指定要查询的项目 ID，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 未填充域名情况下，指定项目查询，若填充了具体域名信息，以域名为主
	Project *int64 `json:"Project,omitnil" name:"Project"`

	// 多域名查询时，默认（false)返回所有域名汇总排序结果
	// Metric 为 url、path、district、isp，Filter 为 flux、request 时，可设置为 true，返回每一个 Domain 的排序数据
	Detail *bool `json:"Detail,omitnil" name:"Detail"`

	// Filter 为 statusCode、OriginStatusCode 时，填充指定状态码查询排序结果
	Code *string `json:"Code,omitnil" name:"Code"`

	// 指定服务地域查询，不填充表示查询中国境内 CDN 数据
	// mainland：指定查询中国境内 CDN 数据
	// overseas：指定查询中国境外 CDN 数据，支持的 Metric 为 url、district、host、originHost，当 Metric 为 originHost 时仅支持 flux、request、bandwidth Filter
	Area *string `json:"Area,omitnil" name:"Area"`

	// 查询中国境外CDN数据，且仅当 Metric 为 district 或 host 时，可指定地区类型查询，不填充表示查询服务地区数据（仅在 Area 为 overseas，且 Metric 是 district 或 host 时可用）
	// server：指定查询服务地区（腾讯云 CDN 节点服务器所在地区）数据
	// client：指定查询客户端地区（用户请求终端所在地区）数据，当 Metric 为 host 时仅支持 flux、request、bandwidth Filter
	AreaType *string `json:"AreaType,omitnil" name:"AreaType"`

	// 指定查询的产品数据，可选为cdn或者ecdn，默认为cdn
	Product *string `json:"Product,omitnil" name:"Product"`

	// 只返回前N条数据，默认为最大值100，metric=url时默认为最大值1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type ListTopDataRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间：yyyy-MM-dd HH:mm:ss
	// 仅支持按分钟粒度的数据查询，按入参抹去秒位作为起始时间，如 StartTime为2018-09-04 10:40:23，返回数据的起始时间为2018-09-04 10:40:00
	// 仅支持 90 天内数据查询
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间：yyyy-MM-dd HH:mm:ss
	// 仅支持按天粒度的数据查询，取入参中的天信息作为结束日期 返回小于等于结束日期当天 23:59:59 产生的数据，如EndTime为2018-09-05 22:40:00，返回数据的结束时间为2018-09-05 23:59:59
	// EndTime 需要大于等于 StartTime
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 排序对象，支持以下几种形式：
	// url：访问 URL 排序（无参数的URL），支持的 Filter 为 flux、request
	// district：省份、国家/地区排序，支持的 Filter 为 flux、request
	// isp：运营商排序，支持的 Filter 为 flux、request
	// host：域名访问数据排序，支持的 Filter 为：flux、request、bandwidth、fluxHitRate、2XX、3XX、4XX、5XX、statusCode
	// originHost：域名回源数据排序，支持的 Filter 为 flux、request、bandwidth、origin_2XX、origin_3XX、origin_4XX、origin_5XX、OriginStatusCode
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 排序使用的指标名称：
	// flux：Metric 为 host 时指代访问流量，originHost 时指代回源流量
	// bandwidth：Metric 为 host 时指代访问带宽，originHost 时指代回源带宽
	// request：Metric 为 host 时指代访问请求数，originHost 时指代回源请求数
	// fluxHitRate：平均流量命中率
	// 2XX：访问 2XX 状态码
	// 3XX：访问 3XX 状态码
	// 4XX：访问 4XX 状态码
	// 5XX：访问 5XX 状态码
	// origin_2XX：回源 2XX 状态码
	// origin_3XX：回源 3XX 状态码
	// origin_4XX：回源 4XX 状态码
	// origin_5XX：回源 5XX 状态码
	// statusCode：指定访问状态码统计，在 Code 参数中填充指定状态码
	// OriginStatusCode：指定回源状态码统计，在 Code 参数中填充指定状态码
	Filter *string `json:"Filter,omitnil" name:"Filter"`

	// 指定查询域名列表，最多可一次性查询 30 个加速域名明细
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 指定要查询的项目 ID，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 未填充域名情况下，指定项目查询，若填充了具体域名信息，以域名为主
	Project *int64 `json:"Project,omitnil" name:"Project"`

	// 多域名查询时，默认（false)返回所有域名汇总排序结果
	// Metric 为 url、path、district、isp，Filter 为 flux、request 时，可设置为 true，返回每一个 Domain 的排序数据
	Detail *bool `json:"Detail,omitnil" name:"Detail"`

	// Filter 为 statusCode、OriginStatusCode 时，填充指定状态码查询排序结果
	Code *string `json:"Code,omitnil" name:"Code"`

	// 指定服务地域查询，不填充表示查询中国境内 CDN 数据
	// mainland：指定查询中国境内 CDN 数据
	// overseas：指定查询中国境外 CDN 数据，支持的 Metric 为 url、district、host、originHost，当 Metric 为 originHost 时仅支持 flux、request、bandwidth Filter
	Area *string `json:"Area,omitnil" name:"Area"`

	// 查询中国境外CDN数据，且仅当 Metric 为 district 或 host 时，可指定地区类型查询，不填充表示查询服务地区数据（仅在 Area 为 overseas，且 Metric 是 district 或 host 时可用）
	// server：指定查询服务地区（腾讯云 CDN 节点服务器所在地区）数据
	// client：指定查询客户端地区（用户请求终端所在地区）数据，当 Metric 为 host 时仅支持 flux、request、bandwidth Filter
	AreaType *string `json:"AreaType,omitnil" name:"AreaType"`

	// 指定查询的产品数据，可选为cdn或者ecdn，默认为cdn
	Product *string `json:"Product,omitnil" name:"Product"`

	// 只返回前N条数据，默认为最大值100，metric=url时默认为最大值1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *ListTopDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTopDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Metric")
	delete(f, "Filter")
	delete(f, "Domains")
	delete(f, "Project")
	delete(f, "Detail")
	delete(f, "Code")
	delete(f, "Area")
	delete(f, "AreaType")
	delete(f, "Product")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTopDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTopDataResponseParams struct {
	// 各个资源的Top 访问数据详情。
	Data []*TopData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListTopDataResponse struct {
	*tchttp.BaseResponse
	Response *ListTopDataResponseParams `json:"Response"`
}

func (r *ListTopDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTopDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTopWafDataRequestParams struct {
	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 指定域名查询，不填写查询整个AppID下数据
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 指定攻击类型
	// 不填则查询所有攻击类型的数据总和
	// AttackType 映射如下:
	// "webshell" : Webshell检测防护
	// "oa" : 常见OA漏洞防护
	// "xss" : XSS跨站脚本攻击防护
	// "xxe" : XXE攻击防护
	// "webscan" : 扫描器攻击漏洞防护
	// "cms" : 常见CMS漏洞防护
	// "upload" : 恶意文件上传攻击防护
	// "sql" : SQL注入攻击防护
	// "cmd_inject": 命令/代码注入攻击防护
	// "osc" : 开源组件漏洞防护
	// "file_read" : 任意文件读取
	// "ldap" : LDAP注入攻击防护
	// "other" : 其它漏洞防护
	AttackType *string `json:"AttackType,omitnil" name:"AttackType"`

	// 指定防御模式
	// 不填则查询所有防御模式的数据总和
	// DefenceMode 映射如下：
	//   observe = '观察模式'
	//   intercept = '拦截模式'
	DefenceMode *string `json:"DefenceMode,omitnil" name:"DefenceMode"`

	// 排序对象，支持以下几种形式：
	// url：攻击目标 url 排序
	// ip：攻击源 IP 排序
	// attackType：攻击类型排序
	// domain：当查询整个AppID下数据时，按照域名请求量排序
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 地域：mainland 或 overseas
	Area *string `json:"Area,omitnil" name:"Area"`

	// 指定攻击类型列表，取值参考AttackType
	AttackTypes []*string `json:"AttackTypes,omitnil" name:"AttackTypes"`

	// 指定域名列表查询，不填写查询整个AppID下数据
	Domains []*string `json:"Domains,omitnil" name:"Domains"`
}

type ListTopWafDataRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 指定域名查询，不填写查询整个AppID下数据
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 指定攻击类型
	// 不填则查询所有攻击类型的数据总和
	// AttackType 映射如下:
	// "webshell" : Webshell检测防护
	// "oa" : 常见OA漏洞防护
	// "xss" : XSS跨站脚本攻击防护
	// "xxe" : XXE攻击防护
	// "webscan" : 扫描器攻击漏洞防护
	// "cms" : 常见CMS漏洞防护
	// "upload" : 恶意文件上传攻击防护
	// "sql" : SQL注入攻击防护
	// "cmd_inject": 命令/代码注入攻击防护
	// "osc" : 开源组件漏洞防护
	// "file_read" : 任意文件读取
	// "ldap" : LDAP注入攻击防护
	// "other" : 其它漏洞防护
	AttackType *string `json:"AttackType,omitnil" name:"AttackType"`

	// 指定防御模式
	// 不填则查询所有防御模式的数据总和
	// DefenceMode 映射如下：
	//   observe = '观察模式'
	//   intercept = '拦截模式'
	DefenceMode *string `json:"DefenceMode,omitnil" name:"DefenceMode"`

	// 排序对象，支持以下几种形式：
	// url：攻击目标 url 排序
	// ip：攻击源 IP 排序
	// attackType：攻击类型排序
	// domain：当查询整个AppID下数据时，按照域名请求量排序
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 地域：mainland 或 overseas
	Area *string `json:"Area,omitnil" name:"Area"`

	// 指定攻击类型列表，取值参考AttackType
	AttackTypes []*string `json:"AttackTypes,omitnil" name:"AttackTypes"`

	// 指定域名列表查询，不填写查询整个AppID下数据
	Domains []*string `json:"Domains,omitnil" name:"Domains"`
}

func (r *ListTopWafDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTopWafDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Domain")
	delete(f, "AttackType")
	delete(f, "DefenceMode")
	delete(f, "Metric")
	delete(f, "Area")
	delete(f, "AttackTypes")
	delete(f, "Domains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTopWafDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTopWafDataResponseParams struct {
	// 攻击类型统计
	TopTypeData []*ScdnTypeData `json:"TopTypeData,omitnil" name:"TopTypeData"`

	// IP统计
	TopIpData []*ScdnTopData `json:"TopIpData,omitnil" name:"TopIpData"`

	// URL统计
	TopUrlData []*ScdnTopUrlData `json:"TopUrlData,omitnil" name:"TopUrlData"`

	// 域名统计
	TopDomainData []*ScdnTopDomainData `json:"TopDomainData,omitnil" name:"TopDomainData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListTopWafDataResponse struct {
	*tchttp.BaseResponse
	Response *ListTopWafDataResponseParams `json:"Response"`
}

func (r *ListTopWafDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTopWafDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogSetInfo struct {
	// 开发者ID
	AppId *uint64 `json:"AppId,omitnil" name:"AppId"`

	// 渠道
	// 注意：此字段可能返回 null，表示取不到有效值。
	Channel *string `json:"Channel,omitnil" name:"Channel"`

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil" name:"LogsetId"`

	// 日志集名字
	LogsetName *string `json:"LogsetName,omitnil" name:"LogsetName"`

	// 是否默认日志集
	IsDefault *uint64 `json:"IsDefault,omitnil" name:"IsDefault"`

	// 日志保存时间，单位为天
	LogsetSavePeriod *uint64 `json:"LogsetSavePeriod,omitnil" name:"LogsetSavePeriod"`

	// 创建日期
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 区域
	Region *string `json:"Region,omitnil" name:"Region"`

	// cls侧是否已经被删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deleted *string `json:"Deleted,omitnil" name:"Deleted"`

	// 英文区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionEn *string `json:"RegionEn,omitnil" name:"RegionEn"`
}

type MainlandConfig struct {
	// 时间戳防盗链配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Authentication *Authentication `json:"Authentication,omitnil" name:"Authentication"`

	// 带宽封顶配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitnil" name:"BandwidthAlert"`

	// 缓存规则配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cache *Cache `json:"Cache,omitnil" name:"Cache"`

	// 缓存相关配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheKey *CacheKey `json:"CacheKey,omitnil" name:"CacheKey"`

	// 智能压缩配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compression *Compression `json:"Compression,omitnil" name:"Compression"`

	// 下载限速配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownstreamCapping *DownstreamCapping `json:"DownstreamCapping,omitnil" name:"DownstreamCapping"`

	// 错误码重定向配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorPage *ErrorPage `json:"ErrorPage,omitnil" name:"ErrorPage"`

	// 301和302自动回源跟随配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitnil" name:"FollowRedirect"`

	// 访问协议强制跳转配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitnil" name:"ForceRedirect"`

	// Https配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Https *Https `json:"Https,omitnil" name:"Https"`

	// IP黑白名单配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpFilter *IpFilter `json:"IpFilter,omitnil" name:"IpFilter"`

	// IP限频配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitnil" name:"IpFreqLimit"`

	// 浏览器缓存规则配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAge *MaxAge `json:"MaxAge,omitnil" name:"MaxAge"`

	// 源站配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Origin *Origin `json:"Origin,omitnil" name:"Origin"`

	// 跨国优化配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginPullOptimization *OriginPullOptimization `json:"OriginPullOptimization,omitnil" name:"OriginPullOptimization"`

	// Range回源配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitnil" name:"RangeOriginPull"`

	// 防盗链配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Referer *Referer `json:"Referer,omitnil" name:"Referer"`

	// 回源请求头部配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestHeader *RequestHeader `json:"RequestHeader,omitnil" name:"RequestHeader"`

	// 源站响应头部配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitnil" name:"ResponseHeader"`

	// 遵循源站缓存头部配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseHeaderCache *ResponseHeaderCache `json:"ResponseHeaderCache,omitnil" name:"ResponseHeaderCache"`

	// seo优化配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Seo *Seo `json:"Seo,omitnil" name:"Seo"`

	// 域名业务类型，web，download，media分别表示静态加速，下载加速和流媒体加速。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 状态码缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitnil" name:"StatusCodeCache"`

	// 视频拖拽配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoSeek *VideoSeek `json:"VideoSeek,omitnil" name:"VideoSeek"`

	// 回源S3私有鉴权。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AwsPrivateAccess *AwsPrivateAccess `json:"AwsPrivateAccess,omitnil" name:"AwsPrivateAccess"`

	// 回源OSS私有鉴权。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OssPrivateAccess *OssPrivateAccess `json:"OssPrivateAccess,omitnil" name:"OssPrivateAccess"`

	// 华为云对象存储回源鉴权
	// 注意：此字段可能返回 null，表示取不到有效值。
	HwPrivateAccess *HwPrivateAccess `json:"HwPrivateAccess,omitnil" name:"HwPrivateAccess"`

	// 七牛云对象存储回源鉴权
	// 注意：此字段可能返回 null，表示取不到有效值。
	QnPrivateAccess *QnPrivateAccess `json:"QnPrivateAccess,omitnil" name:"QnPrivateAccess"`
}

// Predefined struct for user
type ManageClsTopicDomainsRequestParams struct {
	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil" name:"LogsetId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 接入渠道，cdn或者ecdn，默认值为cdn
	Channel *string `json:"Channel,omitnil" name:"Channel"`

	// 域名区域配置，注意：如果此字段为空，则表示解绑对应主题下的所有域名
	DomainAreaConfigs []*DomainAreaConfig `json:"DomainAreaConfigs,omitnil" name:"DomainAreaConfigs"`
}

type ManageClsTopicDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil" name:"LogsetId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 接入渠道，cdn或者ecdn，默认值为cdn
	Channel *string `json:"Channel,omitnil" name:"Channel"`

	// 域名区域配置，注意：如果此字段为空，则表示解绑对应主题下的所有域名
	DomainAreaConfigs []*DomainAreaConfig `json:"DomainAreaConfigs,omitnil" name:"DomainAreaConfigs"`
}

func (r *ManageClsTopicDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageClsTopicDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "TopicId")
	delete(f, "Channel")
	delete(f, "DomainAreaConfigs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManageClsTopicDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManageClsTopicDomainsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ManageClsTopicDomainsResponse struct {
	*tchttp.BaseResponse
	Response *ManageClsTopicDomainsResponseParams `json:"Response"`
}

func (r *ManageClsTopicDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManageClsTopicDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MapInfo struct {
	// 对象 Id
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 对象名称
	Name *string `json:"Name,omitnil" name:"Name"`
}

type MaxAge struct {
	// 浏览器缓存配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// MaxAge 规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAgeRules []*MaxAgeRule `json:"MaxAgeRules,omitnil" name:"MaxAgeRules"`

	// MaxAge 状态码相关规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAgeCodeRule *MaxAgeCodeRule `json:"MaxAgeCodeRule,omitnil" name:"MaxAgeCodeRule"`
}

type MaxAgeCodeRule struct {
	// 处理动作
	// clear：清除 cache-control 头部
	Action *string `json:"Action,omitnil" name:"Action"`

	// 指定HTTP状态码生效，当前仅支持填写"400-599"
	StatusCodes []*string `json:"StatusCodes,omitnil" name:"StatusCodes"`
}

type MaxAgeRule struct {
	// 规则类型：
	// all：所有文件生效
	// file：指定文件后缀生效
	// directory：指定路径生效
	// path：指定绝对路径生效
	// index: 指定主页生效
	MaxAgeType *string `json:"MaxAgeType,omitnil" name:"MaxAgeType"`

	// MaxAgeType 对应类型下的匹配内容：
	// all 时填充 *
	// file 时填充后缀名，如 jpg、txt
	// directory 时填充路径，如 /xxx/test/
	// path 时填充绝对路径，如 /xxx/test.html
	// index 时填充 /
	// 注意：all规则不可删除，默认遵循源站，可修改。
	MaxAgeContents []*string `json:"MaxAgeContents,omitnil" name:"MaxAgeContents"`

	// MaxAge 时间设置，单位秒
	// 注意：时间为0，即不缓存。
	MaxAgeTime *int64 `json:"MaxAgeTime,omitnil" name:"MaxAgeTime"`

	// 是否遵循源站，on或off，开启时忽略时间设置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowOrigin *string `json:"FollowOrigin,omitnil" name:"FollowOrigin"`
}

// Predefined struct for user
type ModifyDomainConfigRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 配置路径
	Route *string `json:"Route,omitnil" name:"Route"`

	// 配置路径值，使用 json 进行序列化，其中固定 update 作为 key
	Value *string `json:"Value,omitnil" name:"Value"`
}

type ModifyDomainConfigRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 配置路径
	Route *string `json:"Route,omitnil" name:"Route"`

	// 配置路径值，使用 json 进行序列化，其中固定 update 作为 key
	Value *string `json:"Value,omitnil" name:"Value"`
}

func (r *ModifyDomainConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Route")
	delete(f, "Value")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDomainConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainConfigResponseParams `json:"Response"`
}

func (r *ModifyDomainConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPurgeFetchTaskStatusRequestParams struct {
	// 执行时间
	ExecutionTime *string `json:"ExecutionTime,omitnil" name:"ExecutionTime"`

	// 执行状态
	// success: 成功
	// failed: 失败
	ExecutionStatus *string `json:"ExecutionStatus,omitnil" name:"ExecutionStatus"`

	// 任务 ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 执行状态详情
	ExecutionStatusDesc *string `json:"ExecutionStatusDesc,omitnil" name:"ExecutionStatusDesc"`
}

type ModifyPurgeFetchTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 执行时间
	ExecutionTime *string `json:"ExecutionTime,omitnil" name:"ExecutionTime"`

	// 执行状态
	// success: 成功
	// failed: 失败
	ExecutionStatus *string `json:"ExecutionStatus,omitnil" name:"ExecutionStatus"`

	// 任务 ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 执行状态详情
	ExecutionStatusDesc *string `json:"ExecutionStatusDesc,omitnil" name:"ExecutionStatusDesc"`
}

func (r *ModifyPurgeFetchTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPurgeFetchTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExecutionTime")
	delete(f, "ExecutionStatus")
	delete(f, "Id")
	delete(f, "ExecutionStatusDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPurgeFetchTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPurgeFetchTaskStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyPurgeFetchTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPurgeFetchTaskStatusResponseParams `json:"Response"`
}

func (r *ModifyPurgeFetchTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPurgeFetchTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OfflineCache struct {
	// 离线缓存配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type Origin struct {
	// 主源站列表
	// <font color=red>修改源站时，需要同时填充对应的 OriginType</font>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Origins []*string `json:"Origins,omitnil" name:"Origins"`

	// 主源站类型
	// <font color=red>当源站列表 Origins 不为空时必填</font>
	// 入参支持以下几种类型：
	// domain：域名类型
	// domainv6：域名解析V6类型
	// cos：对象存储源站
	// third_party: 第三方存储源站
	// igtm: IGTM多活源
	// ip：IP 列表作为源站
	// ipv6：源站列表为一个单独的 IPv6 地址
	// ip_ipv6：源站列表为多个 IPv4 地址和IPv6 地址
	// ip_domain: 支持IP和域名形式源站混填（白名单功能）
	// ip_domainv6：源站列表为多个 IPv4 地址以及域名解析v6地址
	// ipv6_domain: 源站列表为多个 IPv6 地址以及域名
	// ipv6_domainv6：源站列表为多个 IPv6 地址以及域名解析v6地址
	// domain_domainv6：源站列表为多个域名解析v4 地址以及域名解析v6地址
	// ip_ipv6_domain：源站列表为多个 IPv4 地址IPv6 地址以及域名
	// ip_ipv6_domainv6：源站列表为多个 IPv4 地址IPv6 地址以及域名解析v6地址
	// ip_domain_domainv6：源站列表为多个 IPv4 地址域名解析v4 地址以及域名解析v6地址
	// ipv6_domain_domainv6：源站列表为多个 域名解析v4 地址IPv6 地址以及域名解析v6地址
	// ip_ipv6_domain_domainv6：源站列表为多个 IPv4 地址IPv6 地址 域名解析v4 地址以及域名解析v6地址
	// 出参增加以下几种类型：
	// image：数据万象源站
	// ftp：历史 FTP 托管源源站，现已不维护
	// 修改 Origins 时需要同时填充对应的 OriginType
	// IPv6 功能目前尚未全量，需要先申请试用
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginType *string `json:"OriginType,omitnil" name:"OriginType"`

	// 回主源站时 Host 头部
	// <font color=red>当源站类型为cos或者第三方存储加速时,ServerName字段必填</font>
	// 不填充则默认为加速域名
	// 若接入的是泛域名，则回源 Host 默认为访问时的子域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerName *string `json:"ServerName,omitnil" name:"ServerName"`

	// OriginType 为对象存储（COS）时，可以指定是否允许访问私有 bucket
	// 注意：需要先授权 CDN 访问该私有 Bucket 的权限后，才可开启此配置。取值范围: on/off
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosPrivateAccess *string `json:"CosPrivateAccess,omitnil" name:"CosPrivateAccess"`

	// 回源协议配置
	// http：强制 http 回源
	// follow：协议跟随回源
	// https：强制 https 回源，https 回源时仅支持源站 443 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginPullProtocol *string `json:"OriginPullProtocol,omitnil" name:"OriginPullProtocol"`

	// 备源站列表
	// <font color=red>修改备源站时，需要同时填充对应的 BackupOriginType</font>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupOrigins []*string `json:"BackupOrigins,omitnil" name:"BackupOrigins"`

	// 备源站类型
	// <font color=red>备源站列表BackupOrigins 不为空时必填</font>
	// 支持以下类型：
	// domain：域名类型
	// ip：IP 列表作为源站
	// 以下备源源站类型尚未全量支持，需要申请试用：
	// ipv6_domain: 源站列表为多个 IPv6 地址以及域名
	// ip_ipv6：源站列表为多个 IPv4 地址和IPv6 地址
	// ipv6_domain: 源站列表为多个 IPv6 地址以及域名
	// ip_ipv6_domain：源站列表为多个 IPv4 地址IPv6 地址以及域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupOriginType *string `json:"BackupOriginType,omitnil" name:"BackupOriginType"`

	// 回备源站时 Host 头部，不填充则默认为主源站的 ServerName
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupServerName *string `json:"BackupServerName,omitnil" name:"BackupServerName"`

	// 回源路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	BasePath *string `json:"BasePath,omitnil" name:"BasePath"`

	// 回源路径重写规则配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathRules []*PathRule `json:"PathRules,omitnil" name:"PathRules"`

	// 分路径回源配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathBasedOrigin []*PathBasedOriginRule `json:"PathBasedOrigin,omitnil" name:"PathBasedOrigin"`

	// HTTPS回源SNI配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sni *OriginSni `json:"Sni,omitnil" name:"Sni"`

	// HTTPS回源高级配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvanceHttps *AdvanceHttps `json:"AdvanceHttps,omitnil" name:"AdvanceHttps"`

	// 对象存储回源厂商
	// <font color=red>当源站类型为第三方存储源站(third_party)时必填</font>
	// 可选值包括以下:
	// aws_s3: AWS S3
	// ali_oss: 阿里云 OSS
	// hw_obs: 华为 OBS
	// qiniu_kodo: 七牛云 kodo
	// others: 其它厂商对象存储,仅支持兼容以AWS签名算法的对象存储，如腾讯云金融专区COS
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginCompany *string `json:"OriginCompany,omitnil" name:"OriginCompany"`
}

type OriginAuthentication struct {
	// 回源鉴权高级配置开关，取值有：
	// on：开启
	// off：关闭
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 鉴权类型A配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeA *OriginAuthenticationTypeA `json:"TypeA,omitnil" name:"TypeA"`
}

type OriginAuthenticationTypeA struct {
	// 用于计算签名的密钥，只允许字母和数字，长度6-32字节。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitnil" name:"SecretKey"`
}

type OriginCombine struct {
	// 合并回源配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type OriginIp struct {
	// 回源IP段/回源IP，默认返回IP段信息。
	Ip *string `json:"Ip,omitnil" name:"Ip"`
}

type OriginPullOptimization struct {
	// 跨国回源优化配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 跨国类型
	// OVToCN：境外回源境内
	// CNToOV：境内回源境外
	// 注意：此字段可能返回 null，表示取不到有效值。
	OptimizationType *string `json:"OptimizationType,omitnil" name:"OptimizationType"`
}

type OriginPullTimeout struct {
	// 回源建连超时时间，单位为秒，要求5~60之间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectTimeout *uint64 `json:"ConnectTimeout,omitnil" name:"ConnectTimeout"`

	// 回源接收超时时间，单位为秒，要求10 ~ 300之间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiveTimeout *uint64 `json:"ReceiveTimeout,omitnil" name:"ReceiveTimeout"`
}

type OriginSni struct {
	// 是否开启HTTPS回源SNI。
	// 开启：on，
	// 关闭：off
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 回源域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerName *string `json:"ServerName,omitnil" name:"ServerName"`
}

type OssPrivateAccess struct {
	// oss回源鉴权配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 访问ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessKey *string `json:"AccessKey,omitnil" name:"AccessKey"`

	// 密钥。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitnil" name:"SecretKey"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// Bucketname
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitnil" name:"Bucket"`
}

type OthersPrivateAccess struct {
	// 其他厂商对象存储回源鉴权配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 访问ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessKey *string `json:"AccessKey,omitnil" name:"AccessKey"`

	// 密钥。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitnil" name:"SecretKey"`

	// 地域。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 存储桶名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitnil" name:"Bucket"`
}

type OverseaConfig struct {
	// 时间戳防盗链配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Authentication *Authentication `json:"Authentication,omitnil" name:"Authentication"`

	// 带宽封顶配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitnil" name:"BandwidthAlert"`

	// 缓存规则配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cache *Cache `json:"Cache,omitnil" name:"Cache"`

	// 缓存相关配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheKey *CacheKey `json:"CacheKey,omitnil" name:"CacheKey"`

	// 智能压缩配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compression *Compression `json:"Compression,omitnil" name:"Compression"`

	// 下载限速配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownstreamCapping *DownstreamCapping `json:"DownstreamCapping,omitnil" name:"DownstreamCapping"`

	// 错误码重定向配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorPage *ErrorPage `json:"ErrorPage,omitnil" name:"ErrorPage"`

	// 301和302自动回源跟随配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitnil" name:"FollowRedirect"`

	// 访问协议强制跳转配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitnil" name:"ForceRedirect"`

	// Https配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Https *Https `json:"Https,omitnil" name:"Https"`

	// IP黑白名单配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpFilter *IpFilter `json:"IpFilter,omitnil" name:"IpFilter"`

	// IP限频配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitnil" name:"IpFreqLimit"`

	// 浏览器缓存规则配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAge *MaxAge `json:"MaxAge,omitnil" name:"MaxAge"`

	// 源站配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Origin *Origin `json:"Origin,omitnil" name:"Origin"`

	// 跨国优化配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginPullOptimization *OriginPullOptimization `json:"OriginPullOptimization,omitnil" name:"OriginPullOptimization"`

	// Range回源配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitnil" name:"RangeOriginPull"`

	// 防盗链配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Referer *Referer `json:"Referer,omitnil" name:"Referer"`

	// 回源请求头部配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestHeader *RequestHeader `json:"RequestHeader,omitnil" name:"RequestHeader"`

	// 源站响应头部配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitnil" name:"ResponseHeader"`

	// 遵循源站缓存头部配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseHeaderCache *ResponseHeaderCache `json:"ResponseHeaderCache,omitnil" name:"ResponseHeaderCache"`

	// seo优化配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Seo *Seo `json:"Seo,omitnil" name:"Seo"`

	// 域名业务类型，web，download，media分别表示静态加速，下载加速和流媒体加速。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 状态码缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitnil" name:"StatusCodeCache"`

	// 视频拖拽配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoSeek *VideoSeek `json:"VideoSeek,omitnil" name:"VideoSeek"`

	// 回源S3私有鉴权。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AwsPrivateAccess *AwsPrivateAccess `json:"AwsPrivateAccess,omitnil" name:"AwsPrivateAccess"`

	// 回源OSS私有鉴权。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OssPrivateAccess *OssPrivateAccess `json:"OssPrivateAccess,omitnil" name:"OssPrivateAccess"`

	// 华为云对象存储鉴权
	// 注意：此字段可能返回 null，表示取不到有效值。
	HwPrivateAccess *HwPrivateAccess `json:"HwPrivateAccess,omitnil" name:"HwPrivateAccess"`

	// 七牛云对象存储鉴权
	// 注意：此字段可能返回 null，表示取不到有效值。
	QnPrivateAccess *QnPrivateAccess `json:"QnPrivateAccess,omitnil" name:"QnPrivateAccess"`
}

type PathBasedOriginRule struct {
	// 规则类型：
	// file：指定文件后缀生效
	// directory：指定路径生效
	// path：指定绝对路径生效
	// index: 指定主页生效
	RuleType *string `json:"RuleType,omitnil" name:"RuleType"`

	// RuleType 对应类型下的匹配内容：
	// file 时填充后缀名，如 jpg、txt
	// directory 时填充路径，如 /xxx/test/
	// path 时填充绝对路径，如 /xxx/test.html
	// index 时填充 /
	RulePaths []*string `json:"RulePaths,omitnil" name:"RulePaths"`

	// 源站列表，支持域名或ipv4地址
	Origin []*string `json:"Origin,omitnil" name:"Origin"`
}

type PathRule struct {
	// 是否开启通配符“*”匹配：
	// false：关闭
	// true：开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	Regex *bool `json:"Regex,omitnil" name:"Regex"`

	// 匹配的URL路径，仅支持Url路径，不支持参数。默认完全匹配，开启通配符“*”匹配后，最多支持5个通配符，最大长度为1024个字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil" name:"Path"`

	// 路径匹配时的回源源站。暂不支持开了私有读写的COS源。不填写时沿用默认源站。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Origin *string `json:"Origin,omitnil" name:"Origin"`

	// 路径匹配时回源的Host头部。不填写时沿用默认ServerName。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerName *string `json:"ServerName,omitnil" name:"ServerName"`

	// 源站所属区域，支持CN，OV：
	// CN：中国境内
	// OV：中国境外
	// 默认为CN。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginArea *string `json:"OriginArea,omitnil" name:"OriginArea"`

	// 路径匹配时回源的URI路径，必须以“/”开头，不包含参数部分。最大长度为1024个字符。可使用$1, $2, $3, $4, $5分别捕获匹配路径中的通配符号“*”，最多支持10个捕获值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForwardUri *string `json:"ForwardUri,omitnil" name:"ForwardUri"`

	// 路径匹配时回源的头部设置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestHeaders []*HttpHeaderRule `json:"RequestHeaders,omitnil" name:"RequestHeaders"`

	// 当Regex为false时，Path是否开启完全匹配。
	// false：关闭
	// true：开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullMatch *bool `json:"FullMatch,omitnil" name:"FullMatch"`
}

type PostSize struct {
	// POST请求上传文件流式传输最大限制配置开关，取值有：
	// on：开启，平台默认为32MB
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 最大限制，取值在1MB和200MB之间。
	MaxSize *int64 `json:"MaxSize,omitnil" name:"MaxSize"`
}

// Predefined struct for user
type PurgePathCacheRequestParams struct {
	// 目录列表，需要包含协议头部 http:// 或 https://
	Paths []*string `json:"Paths,omitnil" name:"Paths"`

	// 刷新类型
	// flush：刷新产生更新的资源
	// delete：刷新全部资源
	FlushType *string `json:"FlushType,omitnil" name:"FlushType"`

	// 是否对中文字符进行编码后刷新
	UrlEncode *bool `json:"UrlEncode,omitnil" name:"UrlEncode"`

	// 刷新区域
	// 无此参数时，默认刷新加速域名所在加速区域
	// 填充 mainland 时，仅刷新中国境内加速节点上缓存内容
	// 填充 overseas 时，仅刷新中国境外加速节点上缓存内容
	// 指定刷新区域时，需要与域名加速区域匹配
	Area *string `json:"Area,omitnil" name:"Area"`
}

type PurgePathCacheRequest struct {
	*tchttp.BaseRequest
	
	// 目录列表，需要包含协议头部 http:// 或 https://
	Paths []*string `json:"Paths,omitnil" name:"Paths"`

	// 刷新类型
	// flush：刷新产生更新的资源
	// delete：刷新全部资源
	FlushType *string `json:"FlushType,omitnil" name:"FlushType"`

	// 是否对中文字符进行编码后刷新
	UrlEncode *bool `json:"UrlEncode,omitnil" name:"UrlEncode"`

	// 刷新区域
	// 无此参数时，默认刷新加速域名所在加速区域
	// 填充 mainland 时，仅刷新中国境内加速节点上缓存内容
	// 填充 overseas 时，仅刷新中国境外加速节点上缓存内容
	// 指定刷新区域时，需要与域名加速区域匹配
	Area *string `json:"Area,omitnil" name:"Area"`
}

func (r *PurgePathCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PurgePathCacheRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Paths")
	delete(f, "FlushType")
	delete(f, "UrlEncode")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PurgePathCacheRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PurgePathCacheResponseParams struct {
	// 刷新任务 ID，同一批次提交的目录共用一个任务 ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type PurgePathCacheResponse struct {
	*tchttp.BaseResponse
	Response *PurgePathCacheResponseParams `json:"Response"`
}

func (r *PurgePathCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PurgePathCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PurgeTask struct {
	// 刷新任务 ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 刷新 URL
	Url *string `json:"Url,omitnil" name:"Url"`

	// 刷新任务状态
	// fail：刷新失败
	// done：刷新成功
	// process：刷新中
	Status *string `json:"Status,omitnil" name:"Status"`

	// 刷新类型
	// url：URL 刷新
	// path：目录刷新
	PurgeType *string `json:"PurgeType,omitnil" name:"PurgeType"`

	// 刷新方式
	// flush：刷新更新资源（仅目录刷新时有此类型）
	// delete：刷新全部资源
	FlushType *string `json:"FlushType,omitnil" name:"FlushType"`

	// 刷新任务提交时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`
}

// Predefined struct for user
type PurgeUrlsCacheRequestParams struct {
	// URL 列表，需要包含协议头部 http:// 或 https://
	Urls []*string `json:"Urls,omitnil" name:"Urls"`

	// 刷新区域
	// 无此参数时，默认刷新加速域名所在加速区域
	// 填充 mainland 时，仅刷新中国境内加速节点上缓存内容
	// 填充 overseas 时，仅刷新中国境外加速节点上缓存内容
	// 指定刷新区域时，需要与域名加速区域匹配
	Area *string `json:"Area,omitnil" name:"Area"`

	// 是否对中文字符进行编码后刷新
	UrlEncode *bool `json:"UrlEncode,omitnil" name:"UrlEncode"`
}

type PurgeUrlsCacheRequest struct {
	*tchttp.BaseRequest
	
	// URL 列表，需要包含协议头部 http:// 或 https://
	Urls []*string `json:"Urls,omitnil" name:"Urls"`

	// 刷新区域
	// 无此参数时，默认刷新加速域名所在加速区域
	// 填充 mainland 时，仅刷新中国境内加速节点上缓存内容
	// 填充 overseas 时，仅刷新中国境外加速节点上缓存内容
	// 指定刷新区域时，需要与域名加速区域匹配
	Area *string `json:"Area,omitnil" name:"Area"`

	// 是否对中文字符进行编码后刷新
	UrlEncode *bool `json:"UrlEncode,omitnil" name:"UrlEncode"`
}

func (r *PurgeUrlsCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PurgeUrlsCacheRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Urls")
	delete(f, "Area")
	delete(f, "UrlEncode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PurgeUrlsCacheRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PurgeUrlsCacheResponseParams struct {
	// 刷新任务 ID，同一批次提交的 URL 共用一个任务 ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type PurgeUrlsCacheResponse struct {
	*tchttp.BaseResponse
	Response *PurgeUrlsCacheResponseParams `json:"Response"`
}

func (r *PurgeUrlsCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PurgeUrlsCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PushTask struct {
	// 预热任务 ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 预热 URL
	Url *string `json:"Url,omitnil" name:"Url"`

	// 预热任务状态
	// fail：预热失败
	// done：预热成功
	// process：预热中
	// invalid：预热无效(源站返回4xx或5xx状态码)
	Status *string `json:"Status,omitnil" name:"Status"`

	// 预热进度百分比
	Percent *int64 `json:"Percent,omitnil" name:"Percent"`

	// 预热任务提交时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 预热区域
	// mainland：境内
	// overseas：境外
	// global：全球
	Area *string `json:"Area,omitnil" name:"Area"`

	// 预热任务更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

// Predefined struct for user
type PushUrlsCacheRequestParams struct {
	// URL 列表，需要包含协议头部 http:// 或 https://
	Urls []*string `json:"Urls,omitnil" name:"Urls"`

	// 指定预热请求回源时 HTTP 请求的 User-Agent 头部
	// 默认为 TencentCdn
	UserAgent *string `json:"UserAgent,omitnil" name:"UserAgent"`

	// 预热生效区域
	// mainland：预热至境内节点
	// overseas：预热至境外节点
	// global：预热全球节点
	// 不填充情况下，默认为 mainland， URL 中域名必须在对应区域启用了加速服务才能提交对应区域的预热任务
	Area *string `json:"Area,omitnil" name:"Area"`

	// 中国境内区域默认预热至中间层节点，中国境外区域默认预热至边缘节点。预热至边缘产生的边缘层流量会计入计费流量。
	// 填写"middle"或不填充时，可指定预热至中间层节点。
	Layer *string `json:"Layer,omitnil" name:"Layer"`

	// 是否递归解析m3u8文件中的ts分片预热
	// 注意事项：
	// 1. 该功能要求m3u8索引文件能直接请求获取
	// 2. 当前只支持递归解析一级索引和子索引中的ts分片，递归深度不超过3层
	// 3. 解析获取的ts分片会正常累加每日预热用量，当用量超出配额时，会静默处理，不再执行预热
	ParseM3U8 *bool `json:"ParseM3U8,omitnil" name:"ParseM3U8"`

	// 是否关闭Range回源
	// 注意事项：
	// 此功能灰度发布中，敬请期待
	DisableRange *bool `json:"DisableRange,omitnil" name:"DisableRange"`

	// 自定义 HTTP 请求头。最多定义 20 个，Name 长度不超过 128 字节，Value 长度不超过 1024 字节
	Headers []*HTTPHeader `json:"Headers,omitnil" name:"Headers"`

	// 是否对URL进行编码
	UrlEncode *bool `json:"UrlEncode,omitnil" name:"UrlEncode"`
}

type PushUrlsCacheRequest struct {
	*tchttp.BaseRequest
	
	// URL 列表，需要包含协议头部 http:// 或 https://
	Urls []*string `json:"Urls,omitnil" name:"Urls"`

	// 指定预热请求回源时 HTTP 请求的 User-Agent 头部
	// 默认为 TencentCdn
	UserAgent *string `json:"UserAgent,omitnil" name:"UserAgent"`

	// 预热生效区域
	// mainland：预热至境内节点
	// overseas：预热至境外节点
	// global：预热全球节点
	// 不填充情况下，默认为 mainland， URL 中域名必须在对应区域启用了加速服务才能提交对应区域的预热任务
	Area *string `json:"Area,omitnil" name:"Area"`

	// 中国境内区域默认预热至中间层节点，中国境外区域默认预热至边缘节点。预热至边缘产生的边缘层流量会计入计费流量。
	// 填写"middle"或不填充时，可指定预热至中间层节点。
	Layer *string `json:"Layer,omitnil" name:"Layer"`

	// 是否递归解析m3u8文件中的ts分片预热
	// 注意事项：
	// 1. 该功能要求m3u8索引文件能直接请求获取
	// 2. 当前只支持递归解析一级索引和子索引中的ts分片，递归深度不超过3层
	// 3. 解析获取的ts分片会正常累加每日预热用量，当用量超出配额时，会静默处理，不再执行预热
	ParseM3U8 *bool `json:"ParseM3U8,omitnil" name:"ParseM3U8"`

	// 是否关闭Range回源
	// 注意事项：
	// 此功能灰度发布中，敬请期待
	DisableRange *bool `json:"DisableRange,omitnil" name:"DisableRange"`

	// 自定义 HTTP 请求头。最多定义 20 个，Name 长度不超过 128 字节，Value 长度不超过 1024 字节
	Headers []*HTTPHeader `json:"Headers,omitnil" name:"Headers"`

	// 是否对URL进行编码
	UrlEncode *bool `json:"UrlEncode,omitnil" name:"UrlEncode"`
}

func (r *PushUrlsCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PushUrlsCacheRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Urls")
	delete(f, "UserAgent")
	delete(f, "Area")
	delete(f, "Layer")
	delete(f, "ParseM3U8")
	delete(f, "DisableRange")
	delete(f, "Headers")
	delete(f, "UrlEncode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PushUrlsCacheRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PushUrlsCacheResponseParams struct {
	// 此批提交的任务 ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type PushUrlsCacheResponse struct {
	*tchttp.BaseResponse
	Response *PushUrlsCacheResponseParams `json:"Response"`
}

func (r *PushUrlsCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PushUrlsCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QnPrivateAccess struct {
	// 七牛元对象存储回源鉴权配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 访问 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessKey *string `json:"AccessKey,omitnil" name:"AccessKey"`

	// 密钥
	SecretKey *string `json:"SecretKey,omitnil" name:"SecretKey"`
}

type QueryStringKey struct {
	// CacheKey是否由QueryString组成配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 是否重新排序
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reorder *string `json:"Reorder,omitnil" name:"Reorder"`

	// 使用/排除部分url参数，取值有：
	// includeAll：包含所有
	// excludeAll：排除所有
	// includeCustom：自定义包含
	// excludeCustom：自定义排除
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitnil" name:"Action"`

	// 使用/排除的url参数数组，';' 分割
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type Quic struct {
	// Quic功能配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type Quota struct {
	// 单次批量提交配额上限。
	Batch *int64 `json:"Batch,omitnil" name:"Batch"`

	// 每日提交配额上限。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 每日剩余的可提交配额。
	Available *int64 `json:"Available,omitnil" name:"Available"`

	// 配额的区域。
	Area *string `json:"Area,omitnil" name:"Area"`
}

type RangeOriginPull struct {
	// 分片回源配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 分路径分片回源配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	RangeRules []*RangeOriginPullRule `json:"RangeRules,omitnil" name:"RangeRules"`
}

type RangeOriginPullRule struct {
	// 分片回源配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 规则类型：
	// file：指定文件后缀生效
	// directory：指定路径生效
	// path：指定绝对路径生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *string `json:"RuleType,omitnil" name:"RuleType"`

	// RuleType 对应类型下的匹配内容：
	// file 时填充后缀名，如 jpg、txt
	// directory 时填充路径，如 /xxx/test
	// path 时填充绝对路径，如 /xxx/test.html
	// 注意：此字段可能返回 null，表示取不到有效值。
	RulePaths []*string `json:"RulePaths,omitnil" name:"RulePaths"`
}

type RedirectConfig struct {
	// 自定义回源302 follow请求host配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 主源站follow302请求时带的自定义的host头部
	FollowRedirectHost *string `json:"FollowRedirectHost,omitnil" name:"FollowRedirectHost"`

	// 备份源站follow302请求时带的自定义的host头部
	FollowRedirectBackupHost *string `json:"FollowRedirectBackupHost,omitnil" name:"FollowRedirectBackupHost"`
}

type Referer struct {
	// referer 黑白名单配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// referer 黑白名单配置规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefererRules []*RefererRule `json:"RefererRules,omitnil" name:"RefererRules"`
}

type RefererRule struct {
	// 规则类型：
	// all：所有文件生效
	// file：指定文件后缀生效
	// directory：指定路径生效
	// path：指定绝对路径生效
	RuleType *string `json:"RuleType,omitnil" name:"RuleType"`

	// RuleType 对应类型下的匹配内容：
	// all 时填充 *
	// file 时填充后缀名，如 jpg、txt
	// directory 时填充路径，如 /xxx/test/
	// path 时填充绝对路径，如 /xxx/test.html
	RulePaths []*string `json:"RulePaths,omitnil" name:"RulePaths"`

	// referer 配置类型
	// whitelist：白名单
	// blacklist：黑名单
	RefererType *string `json:"RefererType,omitnil" name:"RefererType"`

	// referer 内容列表
	Referers []*string `json:"Referers,omitnil" name:"Referers"`

	// 是否允许空 referer
	// 防盗链类型为白名单时，true表示允许空 referer，false表示不允许空 referer；
	// 防盗链类型为黑名单时，true表示拒绝空referer，false表示不拒绝空referer；
	AllowEmpty *bool `json:"AllowEmpty,omitnil" name:"AllowEmpty"`
}

type RegionMapRelation struct {
	// 区域ID。
	RegionId *int64 `json:"RegionId,omitnil" name:"RegionId"`

	// 子区域ID列表
	SubRegionIdList []*int64 `json:"SubRegionIdList,omitnil" name:"SubRegionIdList"`
}

type RemoteAuthentication struct {
	// 远程鉴权配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 远程鉴权规则配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemoteAuthenticationRules []*RemoteAuthenticationRule `json:"RemoteAuthenticationRules,omitnil" name:"RemoteAuthenticationRules"`

	// 远程鉴权Server
	// 注意：此字段可能返回 null，表示取不到有效值。
	Server *string `json:"Server,omitnil" name:"Server"`
}

type RemoteAuthenticationRule struct {
	// 远程鉴权Server。
	// 默认值:和上层配置的"Server"一致；
	Server *string `json:"Server,omitnil" name:"Server"`

	// 请求远程鉴权服务器的http方法；取值范围[get,post,head,all]; 
	// all: 表示"遵循终端用户请求方法"
	// 默认值: all
	AuthMethod *string `json:"AuthMethod,omitnil" name:"AuthMethod"`

	// 规则类型：
	// all：所有文件生效
	// file：指定文件后缀生效
	// directory：指定目录生效
	// path：指定文件绝对路径生效
	// 默认值:all
	RuleType *string `json:"RuleType,omitnil" name:"RuleType"`

	// 对应类型下的匹配内容：
	// all 时填充 *
	// file 时填充后缀名，如 jpg、txt
	// directory 时填充路径，如 /xxx/test
	// path 时填充绝对路径，如 /xxx/test.html
	// index 时填充 /
	// 默认值:*
	RulePaths []*string `json:"RulePaths,omitnil" name:"RulePaths"`

	// 请求远程鉴权服务器超时时间，单位毫秒；
	// 取值范围：[1,30 000]
	// 默认值:20000
	AuthTimeout *int64 `json:"AuthTimeout,omitnil" name:"AuthTimeout"`

	// 请求远程鉴权服务器超时后执行拦截或者放行；
	// RETURN_200: 超时后放行；
	// RETURN_403:超时拦截；
	// 默认值:RETURN_200
	AuthTimeoutAction *string `json:"AuthTimeoutAction,omitnil" name:"AuthTimeoutAction"`
}

type ReportData struct {
	// 项目ID/域名ID。
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 项目名称/域名。
	Resource *string `json:"Resource,omitnil" name:"Resource"`

	// 流量总和/带宽最大值，单位分别为bytes，bps。
	Value *int64 `json:"Value,omitnil" name:"Value"`

	// 单个资源占总体百分比。
	Percentage *float64 `json:"Percentage,omitnil" name:"Percentage"`

	// 计费流量总和/计费带宽最大值，单位分别为bytes，bps。
	BillingValue *int64 `json:"BillingValue,omitnil" name:"BillingValue"`

	// 计费数值占总体百分比。
	BillingPercentage *float64 `json:"BillingPercentage,omitnil" name:"BillingPercentage"`
}

type RequestHeader struct {
	// 自定义请求头配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 自定义请求头配置规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderRules []*HttpHeaderPathRule `json:"HeaderRules,omitnil" name:"HeaderRules"`
}

type ResourceBillingData struct {
	// 资源名称，根据查询条件不同分为以下几类：
	// 某一个具体域名：表示该域名明细数据
	// multiDomains：表示多域名汇总明细数据
	// 某一个项目 ID：指定项目查询时，显示为项目 ID
	// all：账号维度数据明细
	Resource *string `json:"Resource,omitnil" name:"Resource"`

	// 计费数据详情
	BillingData []*CdnData `json:"BillingData,omitnil" name:"BillingData"`
}

type ResourceData struct {
	// 资源名称，根据查询条件不同分为以下几类：
	// 单域名：指定单域名查询，表示该域名明细数据，当传入参数 detail 指定为 true 时，显示该域名（ detail 参数默认为 false ）
	// 多域名：指定多个域名查询，表示多域名汇总明细数据，显示 multiDomains
	// 项目 ID：指定项目查询时，表示该项目下的域名汇总明细数据，显示该项目 ID
	// all：账号维度明细数据，即账号下所有域名的汇总明细数据
	Resource *string `json:"Resource,omitnil" name:"Resource"`

	// 资源对应的数据明细
	CdnData []*CdnData `json:"CdnData,omitnil" name:"CdnData"`
}

type ResourceOriginData struct {
	// 资源名称，根据查询条件不同分为以下几类：
	// 具体域名：表示该域名明细数据
	// multiDomains：表示多域名汇总明细数据
	// 项目 ID：指定项目查询时，显示为项目 ID
	// all：账号维度明细数据
	Resource *string `json:"Resource,omitnil" name:"Resource"`

	// 回源数据详情
	OriginData []*CdnData `json:"OriginData,omitnil" name:"OriginData"`
}

type ResponseHeader struct {
	// 自定义响应头配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 自定义响应头规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderRules []*HttpHeaderPathRule `json:"HeaderRules,omitnil" name:"HeaderRules"`
}

type ResponseHeaderCache struct {
	// 源站头部缓存配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type Revalidate struct {
	// 总是回源校验配置开关，取值有：
	// on：开启
	// off：关闭
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 只在特定请求路径回源站校验
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil" name:"Path"`
}

type RuleCache struct {
	// CacheType 对应类型下的匹配内容：
	// all 时填充 *
	// file 时填充后缀名，如 jpg、txt
	// directory 时填充路径，如 /xxx/test
	// path 时填充绝对路径，如 /xxx/test.html
	// index 时填充 /
	// 注意：此字段可能返回 null，表示取不到有效值。
	RulePaths []*string `json:"RulePaths,omitnil" name:"RulePaths"`

	// 规则类型：
	// all：所有文件生效
	// file：指定文件后缀生效
	// directory：指定路径生效
	// path：指定绝对路径生效
	// index：首页
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *string `json:"RuleType,omitnil" name:"RuleType"`

	// 缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheConfig *RuleCacheConfig `json:"CacheConfig,omitnil" name:"CacheConfig"`
}

type RuleCacheConfig struct {
	// 缓存配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cache *CacheConfigCache `json:"Cache,omitnil" name:"Cache"`

	// 不缓存配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoCache *CacheConfigNoCache `json:"NoCache,omitnil" name:"NoCache"`

	// 遵循源站配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowOrigin *CacheConfigFollowOrigin `json:"FollowOrigin,omitnil" name:"FollowOrigin"`
}

type RuleEngine struct {
	// 规则引擎配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil" name:"Content"`
}

type RuleQueryString struct {
	// 路径保留参数配置开关，取值有：
	// on：开启，CacheKey由QueryString组成
	// off：关闭，CacheKey不由QueryString组成
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// includeCustom 包含部分url参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitnil" name:"Action"`

	// 使用/排除的url参数数组，';' 分割
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type ScdnAclConfig struct {
	// SCDN访问控制配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 新版本请使用AdvancedScriptData
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptData []*ScdnAclGroup `json:"ScriptData,omitnil" name:"ScriptData"`

	// 错误页面配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorPage *ScdnErrorPage `json:"ErrorPage,omitnil" name:"ErrorPage"`

	// Acl规则组，switch为on时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedScriptData []*AdvancedScdnAclGroup `json:"AdvancedScriptData,omitnil" name:"AdvancedScriptData"`
}

type ScdnAclGroup struct {
	// 规则名称
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 具体配置
	Configure []*ScdnAclRule `json:"Configure,omitnil" name:"Configure"`

	// 执行动作，intercept|redirect
	Result *string `json:"Result,omitnil" name:"Result"`

	// 规则是否生效，active|inactive
	Status *string `json:"Status,omitnil" name:"Status"`

	// 错误页面配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorPage *ScdnErrorPage `json:"ErrorPage,omitnil" name:"ErrorPage"`
}

type ScdnAclRule struct {
	// 匹配关键字
	MatchKey *string `json:"MatchKey,omitnil" name:"MatchKey"`

	// 逻辑操作符，取值如下
	LogiOperator *string `json:"LogiOperator,omitnil" name:"LogiOperator"`

	// 匹配值。
	MatchValue *string `json:"MatchValue,omitnil" name:"MatchValue"`
}

type ScdnBotConfig struct {
	// Scdn bot配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// Bot cookie策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotCookie []*BotCookie `json:"BotCookie,omitnil" name:"BotCookie"`

	// Bot Js策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotJavaScript []*BotJavaScript `json:"BotJavaScript,omitnil" name:"BotJavaScript"`
}

type ScdnCCRules struct {
	// 规则类型：
	// all：所有文件生效
	// file：指定文件后缀生效
	// directory：指定路径生效
	// path：指定绝对路径生效
	// index：首页
	RuleType *string `json:"RuleType,omitnil" name:"RuleType"`

	// 规则值
	RuleValue []*string `json:"RuleValue,omitnil" name:"RuleValue"`

	// 规则限频
	Qps *uint64 `json:"Qps,omitnil" name:"Qps"`

	// 探测时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectionTime *uint64 `json:"DetectionTime,omitnil" name:"DetectionTime"`

	// 限频阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrequencyLimit *uint64 `json:"FrequencyLimit,omitnil" name:"FrequencyLimit"`

	// IP 惩罚配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	PunishmentSwitch *string `json:"PunishmentSwitch,omitnil" name:"PunishmentSwitch"`

	// IP 惩罚时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	PunishmentTime *uint64 `json:"PunishmentTime,omitnil" name:"PunishmentTime"`

	// 执行动作，intercept|redirect
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitnil" name:"Action"`

	// 动作为 redirect 时，重定向的url
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectUrl *string `json:"RedirectUrl,omitnil" name:"RedirectUrl"`
}

type ScdnConfig struct {
	// scdn cc配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 自定义 cc 防护规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rules []*ScdnCCRules `json:"Rules,omitnil" name:"Rules"`

	// 增强自定义 cc 防护规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedRules []*AdvancedCCRules `json:"AdvancedRules,omitnil" name:"AdvancedRules"`

	// 增强自定义 cc 防护规则， 全局
	// 注意：此字段可能返回 null，表示取不到有效值。
	GlobalAdvancedRules []*AdvancedCCRules `json:"GlobalAdvancedRules,omitnil" name:"GlobalAdvancedRules"`
}

type ScdnDdosConfig struct {
	// Scdn ddos配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type ScdnDomain struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 当前状态，取值online | offline | process
	Status *string `json:"Status,omitnil" name:"Status"`

	// Waf 状态默认为‘/’，取值 close | intercept | observe
	Waf *string `json:"Waf,omitnil" name:"Waf"`

	// Acl 状态默认为‘/’，取值 close | open
	Acl *string `json:"Acl,omitnil" name:"Acl"`

	// CC 状态默认为‘/’，取值 close | open
	CC *string `json:"CC,omitnil" name:"CC"`

	// Ddos 状态默认为‘/’，取值 close | open
	Ddos *string `json:"Ddos,omitnil" name:"Ddos"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// Acl 规则数
	AclRuleNumbers *uint64 `json:"AclRuleNumbers,omitnil" name:"AclRuleNumbers"`

	// Bot 状态默认为‘/’，取值 close | open
	Bot *string `json:"Bot,omitnil" name:"Bot"`

	// 域名加速区域，取值global | mainland |  overseas
	// 注意：此字段可能返回 null，表示取不到有效值。
	Area *string `json:"Area,omitnil" name:"Area"`

	// waf规则等级，可取100|200|300
	// 注意：此字段可能返回 null，表示取不到有效值。
	WafLevel *int64 `json:"WafLevel,omitnil" name:"WafLevel"`
}

type ScdnErrorPage struct {
	// 状态码
	// 执行动作为：intercept 默认传值 403
	// 执行动作为：redirect 默认传值 301
	RedirectCode *int64 `json:"RedirectCode,omitnil" name:"RedirectCode"`

	// 重定向url
	RedirectUrl *string `json:"RedirectUrl,omitnil" name:"RedirectUrl"`
}

type ScdnEventLogConditions struct {
	// 匹配关键字，ip, attack_location
	Key *string `json:"Key,omitnil" name:"Key"`

	// 逻辑操作符，取值 exclude, include
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// 匹配值，允许使用通配符(*)查询，匹配零个、单个、多个字符，例如 1.2.*
	Value *string `json:"Value,omitnil" name:"Value"`
}

type ScdnIpStrategy struct {
	// 域名|global表示全部域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 策略ID
	StrategyId *string `json:"StrategyId,omitnil" name:"StrategyId"`

	// IP白名单列表
	IpList []*string `json:"IpList,omitnil" name:"IpList"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 规则类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *string `json:"RuleType,omitnil" name:"RuleType"`

	// 规则值
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleValue []*string `json:"RuleValue,omitnil" name:"RuleValue"`
}

type ScdnIpStrategyFilter struct {
	// 过滤字段名，支持domain, ip
	Name *string `json:"Name,omitnil" name:"Name"`

	// 过滤字段值
	Value []*string `json:"Value,omitnil" name:"Value"`

	// 是否启用模糊查询，仅支持过滤字段名为domain。
	// 模糊查询时，Value长度最大为1
	Fuzzy *bool `json:"Fuzzy,omitnil" name:"Fuzzy"`
}

type ScdnLogTaskDetail struct {
	// scdn域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 防护类型
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 查询任务开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询任务结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 任务创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 日志包下载链接
	// 成功返回下载链接，其他情况返回'-'
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 任务状态
	// created->任务已经创建
	// processing->任务正在执行
	// done->任务执行成功
	// failed->任务执行失败
	// no-log->没有日志产生
	Status *string `json:"Status,omitnil" name:"Status"`

	// 日志任务唯一id
	TaskID *string `json:"TaskID,omitnil" name:"TaskID"`

	// 攻击类型, 可以为"all"
	// AttackType映射如下:
	//   other = '未知类型'
	//   malicious_scan = "恶意扫描"
	//   sql_inject = "SQL注入攻击"
	//   xss = "XSS攻击"
	//   cmd_inject = "命令注入攻击"
	//   ldap_inject = "LDAP注入攻击"
	//   ssi_inject = "SSI注入攻击"
	//   xml_inject = "XML注入攻击"
	//   web_service = "WEB服务漏洞攻击"
	//   web_app = "WEB应用漏洞攻击"
	//   path_traversal = "路径跨越攻击"
	//   illegal_access_core_file = "核心文件非法访问"
	//   file_upload = "文件上传攻击"
	//   trojan_horse = "木马后门攻击"
	//   csrf = "CSRF攻击"
	//   custom_policy = "自定义策略"
	//   ai_engine= 'AI引擎检出'
	//   malicious_file_upload= '恶意文件上传'
	AttackType *string `json:"AttackType,omitnil" name:"AttackType"`

	// 防御模式,可以为"all"
	// DefenceMode映射如下：
	//   observe = '观察模式'
	//   intercept = '防御模式'
	DefenceMode *string `json:"DefenceMode,omitnil" name:"DefenceMode"`

	// 查询条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Conditions []*ScdnEventLogConditions `json:"Conditions,omitnil" name:"Conditions"`

	// mainland或overseas
	// 注意：此字段可能返回 null，表示取不到有效值。
	Area *string `json:"Area,omitnil" name:"Area"`
}

type ScdnSevenLayerRules struct {
	// 区分大小写
	CaseSensitive *bool `json:"CaseSensitive,omitnil" name:"CaseSensitive"`

	// 规则类型：
	// protocol：协议，填写 HTTP/HTTPS
	// method：请求方法，支持 HEAD、GET、POST、PUT、OPTIONS、TRACE、DELETE、PATCH、CONNECT
	// all：域名 匹配内容固定为"*",不可编辑修改
	// ip：IP 填写 CIDR 表达式
	// directory：路径，以/开头，支持目录和具体路径，128字符以内
	// index：首页 默认固定值：/;/index.html,不可编辑修改
	// path：文件全路径，资源地址，如/acb/test.png，支持通配符，如/abc/*.jpg
	// file：文件扩展名，填写具体扩展名，如 jpg;png;css
	// param：请求参数，填写具体 value 值，512字符以内
	// referer：Referer，填写具体 value 值，512字符以内
	// cookie：Cookie，填写具体 value 值，512字符以内
	// user-agent：User-Agent，填写具体 value 值，512字符以内
	// head：自定义请求头，填写具体value值，512字符以内；内容为空或者不存在时，无匹配内容输入框，填写匹配参数即可
	RuleType *string `json:"RuleType,omitnil" name:"RuleType"`

	// 逻辑操作符，取值 ：
	// 不包含：exclude, 
	// 包含：include, 
	// 不等于：notequal, 
	// 等于：equal, 
	// 前缀匹配：matching
	// 内容为空或不存在：null
	LogicOperator *string `json:"LogicOperator,omitnil" name:"LogicOperator"`

	// 规则值
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleValue []*string `json:"RuleValue,omitnil" name:"RuleValue"`

	// 匹配参数，只有请求参数、Cookie、自定义请求头 有值
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleParam *string `json:"RuleParam,omitnil" name:"RuleParam"`
}

type ScdnTopData struct {
	// 时间
	Time *string `json:"Time,omitnil" name:"Time"`

	// 数值
	Value *uint64 `json:"Value,omitnil" name:"Value"`

	// 运营商
	Isp *string `json:"Isp,omitnil" name:"Isp"`

	// IP地址
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 区域
	District *string `json:"District,omitnil" name:"District"`
}

type ScdnTopDomainData struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 请求量
	Value *uint64 `json:"Value,omitnil" name:"Value"`

	// 百分比
	Percent *float64 `json:"Percent,omitnil" name:"Percent"`
}

type ScdnTopUrlData struct {
	// Top数据的URL
	Url *string `json:"Url,omitnil" name:"Url"`

	// 数值
	Value *uint64 `json:"Value,omitnil" name:"Value"`

	// 时间
	Time *string `json:"Time,omitnil" name:"Time"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type ScdnTypeData struct {
	// 攻击类型
	AttackType *string `json:"AttackType,omitnil" name:"AttackType"`

	// 攻击值
	Value *uint64 `json:"Value,omitnil" name:"Value"`
}

type ScdnWafConfig struct {
	// Scdn waf配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// intercept|observe，默认intercept
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 重定向的错误页面
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorPage *ScdnErrorPage `json:"ErrorPage,omitnil" name:"ErrorPage"`

	// webshell拦截配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebShellSwitch *string `json:"WebShellSwitch,omitnil" name:"WebShellSwitch"`

	// 类型拦截规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rules []*ScdnWafRule `json:"Rules,omitnil" name:"Rules"`

	// waf规则等级，可取100|200|300
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *int64 `json:"Level,omitnil" name:"Level"`

	// waf子规则配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubRuleSwitch []*WafSubRuleStatus `json:"SubRuleSwitch,omitnil" name:"SubRuleSwitch"`
}

type ScdnWafRule struct {
	// 攻击类型
	AttackType *string `json:"AttackType,omitnil" name:"AttackType"`

	// 防护措施，observe
	Operate *string `json:"Operate,omitnil" name:"Operate"`
}

type SchemeKey struct {
	// scheme作为cache key配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

// Predefined struct for user
type SearchClsLogRequestParams struct {
	// 需要查询的日志集ID
	LogsetId *string `json:"LogsetId,omitnil" name:"LogsetId"`

	// 需要查询的日志主题ID组合，以逗号分隔
	TopicIds *string `json:"TopicIds,omitnil" name:"TopicIds"`

	// 需要查询的日志的起始时间，格式 YYYY-mm-dd HH:MM:SS
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 需要查询的日志的结束时间，格式 YYYY-mm-dd HH:MM:SS
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 单次要返回的日志条数，单次返回的最大条数为100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 接入渠道，cdn或者ecdn，默认值为cdn
	Channel *string `json:"Channel,omitnil" name:"Channel"`

	// 需要查询的内容，详情请参考https://cloud.tencent.com/document/product/614/16982
	Query *string `json:"Query,omitnil" name:"Query"`

	// 加载更多使用，透传上次返回的 context 值，获取后续的日志内容，通过游标最多可获取10000条，请尽可能缩小时间范围
	Context *string `json:"Context,omitnil" name:"Context"`

	// 按日志时间排序， asc（升序）或者 desc（降序），默认为 desc
	Sort *string `json:"Sort,omitnil" name:"Sort"`
}

type SearchClsLogRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的日志集ID
	LogsetId *string `json:"LogsetId,omitnil" name:"LogsetId"`

	// 需要查询的日志主题ID组合，以逗号分隔
	TopicIds *string `json:"TopicIds,omitnil" name:"TopicIds"`

	// 需要查询的日志的起始时间，格式 YYYY-mm-dd HH:MM:SS
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 需要查询的日志的结束时间，格式 YYYY-mm-dd HH:MM:SS
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 单次要返回的日志条数，单次返回的最大条数为100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 接入渠道，cdn或者ecdn，默认值为cdn
	Channel *string `json:"Channel,omitnil" name:"Channel"`

	// 需要查询的内容，详情请参考https://cloud.tencent.com/document/product/614/16982
	Query *string `json:"Query,omitnil" name:"Query"`

	// 加载更多使用，透传上次返回的 context 值，获取后续的日志内容，通过游标最多可获取10000条，请尽可能缩小时间范围
	Context *string `json:"Context,omitnil" name:"Context"`

	// 按日志时间排序， asc（升序）或者 desc（降序），默认为 desc
	Sort *string `json:"Sort,omitnil" name:"Sort"`
}

func (r *SearchClsLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClsLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "TopicIds")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Channel")
	delete(f, "Query")
	delete(f, "Context")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchClsLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchClsLogResponseParams struct {
	// 查询结果
	Logs *ClsSearchLogs `json:"Logs,omitnil" name:"Logs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SearchClsLogResponse struct {
	*tchttp.BaseResponse
	Response *SearchClsLogResponseParams `json:"Response"`
}

func (r *SearchClsLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClsLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityConfig struct {
	// scdn 安全配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type Seo struct {
	// SEO 搜索引擎优化配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type ServerCert struct {
	// 服务器证书 ID 在 SSL 证书管理进行证书托管时自动生成
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitnil" name:"CertId"`

	// 服务器证书名称
	// 在 SSL 证书管理进行证书托管时自动生成
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertName *string `json:"CertName,omitnil" name:"CertName"`

	// 服务器证书信息
	// 上传自有证书时必填，需要包含完整的证书链
	// 注意：此字段可能返回 null，表示取不到有效值。
	Certificate *string `json:"Certificate,omitnil" name:"Certificate"`

	// 服务器密钥信息
	// 上传自有证书时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateKey *string `json:"PrivateKey,omitnil" name:"PrivateKey"`

	// 证书过期时间
	// 作为入参配置时无需填充
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 证书颁发时间
	// 作为入参配置时无需填充
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployTime *string `json:"DeployTime,omitnil" name:"DeployTime"`

	// 证书备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 证书来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	From *string `json:"From,omitnil" name:"From"`
}

type ShareCname struct {
	// ShareCname 配置开关, 取值有：
	// on：开启，使用共享CNAME
	// off：关闭，使用默认CNAME
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 设置共享CNAME.
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cname *string `json:"Cname,omitnil" name:"Cname"`
}

type SimpleCache struct {
	// 缓存过期时间规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheRules []*SimpleCacheRule `json:"CacheRules,omitnil" name:"CacheRules"`

	// 遵循源站 Cache-Control: max-age 配置
	// on：开启
	// off：关闭
	// 开启后，未能匹配 CacheRules 规则的资源将根据源站返回的 max-age 值进行节点缓存；匹配了 CacheRules 规则的资源将按照 CacheRules 中设置的缓存过期时间在节点进行缓存
	// 与 CompareMaxAge 冲突，不能同时开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowOrigin *string `json:"FollowOrigin,omitnil" name:"FollowOrigin"`

	// 强制缓存
	// on：开启
	// off：关闭
	// 默认为关闭状态，开启后，源站返回的 no-store、no-cache 资源，也将按照 CacheRules 规则进行缓存
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreCacheControl *string `json:"IgnoreCacheControl,omitnil" name:"IgnoreCacheControl"`

	// 忽略源站的Set-Cookie头部
	// on：开启
	// off：关闭
	// 默认为关闭状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreSetCookie *string `json:"IgnoreSetCookie,omitnil" name:"IgnoreSetCookie"`

	// 高级缓存过期配置，开启时会对比源站返回的 max-age 值与 CacheRules 中设置的缓存过期时间，取最小值在节点进行缓存
	// on：开启
	// off：关闭
	// 默认为关闭状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareMaxAge *string `json:"CompareMaxAge,omitnil" name:"CompareMaxAge"`

	// 总是回源站校验
	// 注意：此字段可能返回 null，表示取不到有效值。
	Revalidate *Revalidate `json:"Revalidate,omitnil" name:"Revalidate"`
}

type SimpleCacheRule struct {
	// 规则类型：
	// all：所有文件生效
	// file：指定文件后缀生效
	// directory：指定路径生效
	// path：指定绝对路径生效
	// index：首页
	CacheType *string `json:"CacheType,omitnil" name:"CacheType"`

	// CacheType 对应类型下的匹配内容：
	// all 时填充 *
	// file 时填充后缀名，如 jpg、txt
	// directory 时填充路径，如 /xxx/test
	// path 时填充绝对路径，如 /xxx/test.html
	// index 时填充 /
	CacheContents []*string `json:"CacheContents,omitnil" name:"CacheContents"`

	// 缓存过期时间设置
	// 单位为秒，最大可设置为 365 天
	CacheTime *int64 `json:"CacheTime,omitnil" name:"CacheTime"`
}

type Sort struct {
	// 排序字段，当前支持：
	// createTime，域名创建时间
	// certExpireTime，证书过期时间
	// 默认createTime。
	Key *string `json:"Key,omitnil" name:"Key"`

	// asc/desc，默认desc。
	Sequence *string `json:"Sequence,omitnil" name:"Sequence"`
}

type SpecificConfig struct {
	// 国内特殊配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mainland *MainlandConfig `json:"Mainland,omitnil" name:"Mainland"`

	// 海外特殊配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Overseas *OverseaConfig `json:"Overseas,omitnil" name:"Overseas"`
}

// Predefined struct for user
type StartCdnDomainRequestParams struct {
	// 域名
	// 域名状态需要为【已停用】
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type StartCdnDomainRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	// 域名状态需要为【已停用】
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *StartCdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartCdnDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartCdnDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartCdnDomainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StartCdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *StartCdnDomainResponseParams `json:"Response"`
}

func (r *StartCdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartCdnDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartScdnDomainRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type StartScdnDomainRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *StartScdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartScdnDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartScdnDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartScdnDomainResponseParams struct {
	// 开启结果，Success表示成功
	Result *string `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StartScdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *StartScdnDomainResponseParams `json:"Response"`
}

func (r *StartScdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartScdnDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StatisticItem struct {
	// 封顶类型，累计用量total，瞬时用量moment
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 自动解封时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnBlockTime *uint64 `json:"UnBlockTime,omitnil" name:"UnBlockTime"`

	// 带宽、流量阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	BpsThreshold *uint64 `json:"BpsThreshold,omitnil" name:"BpsThreshold"`

	// 关闭方式 返回404:RETURN_404, dns回源：RESOLVE_DNS_TO_ORIGIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	CounterMeasure *string `json:"CounterMeasure,omitnil" name:"CounterMeasure"`

	// 触发提醒阈值百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertPercentage *uint64 `json:"AlertPercentage,omitnil" name:"AlertPercentage"`

	// 累计用量封顶告警配置，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertSwitch *string `json:"AlertSwitch,omitnil" name:"AlertSwitch"`

	// 指标类型，流量flux或带宽bandwidth
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 检测周期，单位分钟，60或1440
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cycle *uint64 `json:"Cycle,omitnil" name:"Cycle"`

	// 累计用量封顶配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type StatusCodeCache struct {
	// 状态码缓存过期配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 状态码缓存过期规则明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheRules []*StatusCodeCacheRule `json:"CacheRules,omitnil" name:"CacheRules"`
}

type StatusCodeCacheRule struct {
	// http 状态码
	// 支持 403、404 状态码
	StatusCode *string `json:"StatusCode,omitnil" name:"StatusCode"`

	// 状态码缓存过期时间，单位秒
	CacheTime *int64 `json:"CacheTime,omitnil" name:"CacheTime"`
}

// Predefined struct for user
type StopCdnDomainRequestParams struct {
	// 域名
	// 域名需要为【已启动】状态
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type StopCdnDomainRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	// 域名需要为【已启动】状态
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *StopCdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCdnDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopCdnDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopCdnDomainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StopCdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *StopCdnDomainResponseParams `json:"Response"`
}

func (r *StopCdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCdnDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopScdnDomainRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type StopScdnDomainRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *StopScdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopScdnDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopScdnDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopScdnDomainResponseParams struct {
	// 关闭结果，Success表示成功
	Result *string `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StopScdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *StopScdnDomainResponseParams `json:"Response"`
}

func (r *StopScdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopScdnDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SummarizedData struct {
	// 汇总方式，存在以下几种：
	// sum：累加求和
	// max：最大值，带宽模式下，采用 5 分钟粒度汇总数据，计算峰值带宽
	// avg：平均值
	Name *string `json:"Name,omitnil" name:"Name"`

	// 汇总后的数据值
	Value *float64 `json:"Value,omitnil" name:"Value"`
}

type Tag struct {
	// 标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

type TimestampData struct {
	// 数据统计时间点，采用向前汇总模式
	// 以 5 分钟粒度为例，13:35:00 时间点代表的统计数据区间为 13:35:00 至 13:39:59
	Time *string `json:"Time,omitnil" name:"Time"`

	// 数据值
	Value *float64 `json:"Value,omitnil" name:"Value"`
}

type TopData struct {
	// 资源名称，根据查询条件不同分为以下几类：
	// 具体域名：表示该域名明细数据
	// multiDomains：表示多域名汇总明细数据
	// 项目 ID：指定项目查询时，显示为项目 ID
	// all：账号维度明细数据
	Resource *string `json:"Resource,omitnil" name:"Resource"`

	// 排序结果详情
	DetailData []*TopDetailData `json:"DetailData,omitnil" name:"DetailData"`
}

type TopDataMore struct {
	// 资源名称，根据查询条件不同分为以下几类：
	Resource *string `json:"Resource,omitnil" name:"Resource"`

	// 排序结果详情
	DetailData []*TopDetailDataMore `json:"DetailData,omitnil" name:"DetailData"`
}

type TopDetailData struct {
	// 数据类型的名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 数据值
	Value *float64 `json:"Value,omitnil" name:"Value"`
}

type TopDetailDataMore struct {
	// 数据类型的名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 数据值
	Value *float64 `json:"Value,omitnil" name:"Value"`

	// 数据值在总值中的百分比
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent *float64 `json:"Percent,omitnil" name:"Percent"`
}

type TopicInfo struct {
	// 主题ID
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 主题名字
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// 是否启用投递
	Enabled *int64 `json:"Enabled,omitnil" name:"Enabled"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 归属于cdn或ecdn
	// 注意：此字段可能返回 null，表示取不到有效值。
	Channel *string `json:"Channel,omitnil" name:"Channel"`

	// cls侧是否已经被删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deleted *string `json:"Deleted,omitnil" name:"Deleted"`
}

type TpgAdapter struct {
	// 图片优化-TpgAdapter配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type TrafficPackage struct {
	// 流量包 Id
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 流量包类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 流量包大小（单位为 Byte）
	Bytes *int64 `json:"Bytes,omitnil" name:"Bytes"`

	// 已消耗流量（单位为 Byte）
	BytesUsed *int64 `json:"BytesUsed,omitnil" name:"BytesUsed"`

	// 流量包状态
	// enabled：已启用
	// expired：已过期
	// disabled：未启用
	Status *string `json:"Status,omitnil" name:"Status"`

	// 流量包发放时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 流量包生效时间
	EnableTime *string `json:"EnableTime,omitnil" name:"EnableTime"`

	// 流量包过期时间
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 流量包是否续订
	ContractExtension *bool `json:"ContractExtension,omitnil" name:"ContractExtension"`

	// 流量包是否自动续订
	AutoExtension *bool `json:"AutoExtension,omitnil" name:"AutoExtension"`

	// 流量包来源
	Channel *string `json:"Channel,omitnil" name:"Channel"`

	// 流量包生效区域，mainland或overseas
	Area *string `json:"Area,omitnil" name:"Area"`

	// 流量包生命周期月数
	LifeTimeMonth *int64 `json:"LifeTimeMonth,omitnil" name:"LifeTimeMonth"`

	// 流量包是否支持续订
	ExtensionAvailable *bool `json:"ExtensionAvailable,omitnil" name:"ExtensionAvailable"`

	// 流量包是否支持退费
	RefundAvailable *bool `json:"RefundAvailable,omitnil" name:"RefundAvailable"`

	// 流量包生效区域
	// 0：中国大陆
	// 1：亚太一区
	// 2：亚太二区
	// 3：亚太三区
	// 4：中东
	// 5：北美
	// 6：欧洲
	// 7：南美
	// 8：非洲
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *int64 `json:"Region,omitnil" name:"Region"`

	// 流量包类型id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigId *int64 `json:"ConfigId,omitnil" name:"ConfigId"`

	// 流量包当前续订模式，0 未续订、1到期续订、2用完续订、3到期或用完续订
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtensionMode *uint64 `json:"ExtensionMode,omitnil" name:"ExtensionMode"`

	// 流量包实际生效时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrueEnableTime *string `json:"TrueEnableTime,omitnil" name:"TrueEnableTime"`

	// 流量包实际过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrueExpireTime *string `json:"TrueExpireTime,omitnil" name:"TrueExpireTime"`
}

// Predefined struct for user
type UpdateDomainConfigRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 项目 ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 源站配置
	Origin *Origin `json:"Origin,omitnil" name:"Origin"`

	// IP 黑白名单配置
	IpFilter *IpFilter `json:"IpFilter,omitnil" name:"IpFilter"`

	// IP 限频配置
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitnil" name:"IpFreqLimit"`

	// 状态码缓存配置
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitnil" name:"StatusCodeCache"`

	// 智能压缩配置
	Compression *Compression `json:"Compression,omitnil" name:"Compression"`

	// 带宽封顶配置
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitnil" name:"BandwidthAlert"`

	// Range 回源配置
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitnil" name:"RangeOriginPull"`

	// 301/302 回源跟随配置
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitnil" name:"FollowRedirect"`

	// 错误码重定向配置（功能灰度中，尚未全量）
	ErrorPage *ErrorPage `json:"ErrorPage,omitnil" name:"ErrorPage"`

	// 回源请求头部配置
	RequestHeader *RequestHeader `json:"RequestHeader,omitnil" name:"RequestHeader"`

	// 响应头部配置
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitnil" name:"ResponseHeader"`

	// 下载速度配置
	DownstreamCapping *DownstreamCapping `json:"DownstreamCapping,omitnil" name:"DownstreamCapping"`

	// 节点缓存键配置
	CacheKey *CacheKey `json:"CacheKey,omitnil" name:"CacheKey"`

	// 头部缓存配置
	ResponseHeaderCache *ResponseHeaderCache `json:"ResponseHeaderCache,omitnil" name:"ResponseHeaderCache"`

	// 视频拖拽配置
	VideoSeek *VideoSeek `json:"VideoSeek,omitnil" name:"VideoSeek"`

	// 缓存过期时间配置
	Cache *Cache `json:"Cache,omitnil" name:"Cache"`

	// 跨国链路优化配置（已下线）
	OriginPullOptimization *OriginPullOptimization `json:"OriginPullOptimization,omitnil" name:"OriginPullOptimization"`

	// Https 加速配置
	Https *Https `json:"Https,omitnil" name:"Https"`

	// 时间戳防盗链配置
	Authentication *Authentication `json:"Authentication,omitnil" name:"Authentication"`

	// SEO 优化配置
	Seo *Seo `json:"Seo,omitnil" name:"Seo"`

	// 访问协议强制跳转配置
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitnil" name:"ForceRedirect"`

	// Referer 防盗链配置
	Referer *Referer `json:"Referer,omitnil" name:"Referer"`

	// 浏览器缓存配置（功能灰度中，尚未全量）
	MaxAge *MaxAge `json:"MaxAge,omitnil" name:"MaxAge"`

	// 地域属性特殊配置
	// 适用于域名境内加速、境外加速配置不一致场景
	SpecificConfig *SpecificConfig `json:"SpecificConfig,omitnil" name:"SpecificConfig"`

	// 域名业务类型
	// web：静态加速
	// download：下载加速
	// media：流媒体点播加速
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 域名加速区域
	// mainland：中国境内加速
	// overseas：中国境外加速
	// global：全球加速
	// 从mainland/overseas修改至global时，域名的配置将被同步至overseas/mainland。若域名含有后端特殊配置，此类配置的同步过程有一定延时，请耐心等待
	Area *string `json:"Area,omitnil" name:"Area"`

	// 回源超时配置
	OriginPullTimeout *OriginPullTimeout `json:"OriginPullTimeout,omitnil" name:"OriginPullTimeout"`

	// 回源S3私有鉴权
	AwsPrivateAccess *AwsPrivateAccess `json:"AwsPrivateAccess,omitnil" name:"AwsPrivateAccess"`

	// UA黑白名单配置
	UserAgentFilter *UserAgentFilter `json:"UserAgentFilter,omitnil" name:"UserAgentFilter"`

	// 访问控制
	AccessControl *AccessControl `json:"AccessControl,omitnil" name:"AccessControl"`

	// 访问URL重写配置
	UrlRedirect *UrlRedirect `json:"UrlRedirect,omitnil" name:"UrlRedirect"`

	// 访问端口配置
	AccessPort []*int64 `json:"AccessPort,omitnil" name:"AccessPort"`

	// 时间戳防盗链高级版配置，白名单功能
	AdvancedAuthentication *AdvancedAuthentication `json:"AdvancedAuthentication,omitnil" name:"AdvancedAuthentication"`

	// 回源鉴权高级版配置，白名单功能
	OriginAuthentication *OriginAuthentication `json:"OriginAuthentication,omitnil" name:"OriginAuthentication"`

	// Ipv6 访问配置
	Ipv6Access *Ipv6Access `json:"Ipv6Access,omitnil" name:"Ipv6Access"`

	// 离线缓存
	OfflineCache *OfflineCache `json:"OfflineCache,omitnil" name:"OfflineCache"`

	// 合并回源
	OriginCombine *OriginCombine `json:"OriginCombine,omitnil" name:"OriginCombine"`

	// POST请求传输配置
	PostMaxSize *PostSize `json:"PostMaxSize,omitnil" name:"PostMaxSize"`

	// Quic访问（收费服务，详见计费说明和产品文档）
	Quic *Quic `json:"Quic,omitnil" name:"Quic"`

	// 回源OSS私有鉴权
	OssPrivateAccess *OssPrivateAccess `json:"OssPrivateAccess,omitnil" name:"OssPrivateAccess"`

	// WebSocket配置
	WebSocket *WebSocket `json:"WebSocket,omitnil" name:"WebSocket"`

	// 远程鉴权配置
	RemoteAuthentication *RemoteAuthentication `json:"RemoteAuthentication,omitnil" name:"RemoteAuthentication"`

	// 共享CNAME配置，白名单功能
	ShareCname *ShareCname `json:"ShareCname,omitnil" name:"ShareCname"`

	// 华为云对象存储回源鉴权
	HwPrivateAccess *HwPrivateAccess `json:"HwPrivateAccess,omitnil" name:"HwPrivateAccess"`

	// 七牛云对象存储回源鉴权
	QnPrivateAccess *QnPrivateAccess `json:"QnPrivateAccess,omitnil" name:"QnPrivateAccess"`

	// 其他厂商对象存储回源鉴权
	OthersPrivateAccess *OthersPrivateAccess `json:"OthersPrivateAccess,omitnil" name:"OthersPrivateAccess"`

	// HTTPS服务（收费服务，详见计费说明和产品文档）
	HttpsBilling *HttpsBilling `json:"HttpsBilling,omitnil" name:"HttpsBilling"`
}

type UpdateDomainConfigRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 项目 ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 源站配置
	Origin *Origin `json:"Origin,omitnil" name:"Origin"`

	// IP 黑白名单配置
	IpFilter *IpFilter `json:"IpFilter,omitnil" name:"IpFilter"`

	// IP 限频配置
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitnil" name:"IpFreqLimit"`

	// 状态码缓存配置
	StatusCodeCache *StatusCodeCache `json:"StatusCodeCache,omitnil" name:"StatusCodeCache"`

	// 智能压缩配置
	Compression *Compression `json:"Compression,omitnil" name:"Compression"`

	// 带宽封顶配置
	BandwidthAlert *BandwidthAlert `json:"BandwidthAlert,omitnil" name:"BandwidthAlert"`

	// Range 回源配置
	RangeOriginPull *RangeOriginPull `json:"RangeOriginPull,omitnil" name:"RangeOriginPull"`

	// 301/302 回源跟随配置
	FollowRedirect *FollowRedirect `json:"FollowRedirect,omitnil" name:"FollowRedirect"`

	// 错误码重定向配置（功能灰度中，尚未全量）
	ErrorPage *ErrorPage `json:"ErrorPage,omitnil" name:"ErrorPage"`

	// 回源请求头部配置
	RequestHeader *RequestHeader `json:"RequestHeader,omitnil" name:"RequestHeader"`

	// 响应头部配置
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitnil" name:"ResponseHeader"`

	// 下载速度配置
	DownstreamCapping *DownstreamCapping `json:"DownstreamCapping,omitnil" name:"DownstreamCapping"`

	// 节点缓存键配置
	CacheKey *CacheKey `json:"CacheKey,omitnil" name:"CacheKey"`

	// 头部缓存配置
	ResponseHeaderCache *ResponseHeaderCache `json:"ResponseHeaderCache,omitnil" name:"ResponseHeaderCache"`

	// 视频拖拽配置
	VideoSeek *VideoSeek `json:"VideoSeek,omitnil" name:"VideoSeek"`

	// 缓存过期时间配置
	Cache *Cache `json:"Cache,omitnil" name:"Cache"`

	// 跨国链路优化配置（已下线）
	OriginPullOptimization *OriginPullOptimization `json:"OriginPullOptimization,omitnil" name:"OriginPullOptimization"`

	// Https 加速配置
	Https *Https `json:"Https,omitnil" name:"Https"`

	// 时间戳防盗链配置
	Authentication *Authentication `json:"Authentication,omitnil" name:"Authentication"`

	// SEO 优化配置
	Seo *Seo `json:"Seo,omitnil" name:"Seo"`

	// 访问协议强制跳转配置
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitnil" name:"ForceRedirect"`

	// Referer 防盗链配置
	Referer *Referer `json:"Referer,omitnil" name:"Referer"`

	// 浏览器缓存配置（功能灰度中，尚未全量）
	MaxAge *MaxAge `json:"MaxAge,omitnil" name:"MaxAge"`

	// 地域属性特殊配置
	// 适用于域名境内加速、境外加速配置不一致场景
	SpecificConfig *SpecificConfig `json:"SpecificConfig,omitnil" name:"SpecificConfig"`

	// 域名业务类型
	// web：静态加速
	// download：下载加速
	// media：流媒体点播加速
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`

	// 域名加速区域
	// mainland：中国境内加速
	// overseas：中国境外加速
	// global：全球加速
	// 从mainland/overseas修改至global时，域名的配置将被同步至overseas/mainland。若域名含有后端特殊配置，此类配置的同步过程有一定延时，请耐心等待
	Area *string `json:"Area,omitnil" name:"Area"`

	// 回源超时配置
	OriginPullTimeout *OriginPullTimeout `json:"OriginPullTimeout,omitnil" name:"OriginPullTimeout"`

	// 回源S3私有鉴权
	AwsPrivateAccess *AwsPrivateAccess `json:"AwsPrivateAccess,omitnil" name:"AwsPrivateAccess"`

	// UA黑白名单配置
	UserAgentFilter *UserAgentFilter `json:"UserAgentFilter,omitnil" name:"UserAgentFilter"`

	// 访问控制
	AccessControl *AccessControl `json:"AccessControl,omitnil" name:"AccessControl"`

	// 访问URL重写配置
	UrlRedirect *UrlRedirect `json:"UrlRedirect,omitnil" name:"UrlRedirect"`

	// 访问端口配置
	AccessPort []*int64 `json:"AccessPort,omitnil" name:"AccessPort"`

	// 时间戳防盗链高级版配置，白名单功能
	AdvancedAuthentication *AdvancedAuthentication `json:"AdvancedAuthentication,omitnil" name:"AdvancedAuthentication"`

	// 回源鉴权高级版配置，白名单功能
	OriginAuthentication *OriginAuthentication `json:"OriginAuthentication,omitnil" name:"OriginAuthentication"`

	// Ipv6 访问配置
	Ipv6Access *Ipv6Access `json:"Ipv6Access,omitnil" name:"Ipv6Access"`

	// 离线缓存
	OfflineCache *OfflineCache `json:"OfflineCache,omitnil" name:"OfflineCache"`

	// 合并回源
	OriginCombine *OriginCombine `json:"OriginCombine,omitnil" name:"OriginCombine"`

	// POST请求传输配置
	PostMaxSize *PostSize `json:"PostMaxSize,omitnil" name:"PostMaxSize"`

	// Quic访问（收费服务，详见计费说明和产品文档）
	Quic *Quic `json:"Quic,omitnil" name:"Quic"`

	// 回源OSS私有鉴权
	OssPrivateAccess *OssPrivateAccess `json:"OssPrivateAccess,omitnil" name:"OssPrivateAccess"`

	// WebSocket配置
	WebSocket *WebSocket `json:"WebSocket,omitnil" name:"WebSocket"`

	// 远程鉴权配置
	RemoteAuthentication *RemoteAuthentication `json:"RemoteAuthentication,omitnil" name:"RemoteAuthentication"`

	// 共享CNAME配置，白名单功能
	ShareCname *ShareCname `json:"ShareCname,omitnil" name:"ShareCname"`

	// 华为云对象存储回源鉴权
	HwPrivateAccess *HwPrivateAccess `json:"HwPrivateAccess,omitnil" name:"HwPrivateAccess"`

	// 七牛云对象存储回源鉴权
	QnPrivateAccess *QnPrivateAccess `json:"QnPrivateAccess,omitnil" name:"QnPrivateAccess"`

	// 其他厂商对象存储回源鉴权
	OthersPrivateAccess *OthersPrivateAccess `json:"OthersPrivateAccess,omitnil" name:"OthersPrivateAccess"`

	// HTTPS服务（收费服务，详见计费说明和产品文档）
	HttpsBilling *HttpsBilling `json:"HttpsBilling,omitnil" name:"HttpsBilling"`
}

func (r *UpdateDomainConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDomainConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "ProjectId")
	delete(f, "Origin")
	delete(f, "IpFilter")
	delete(f, "IpFreqLimit")
	delete(f, "StatusCodeCache")
	delete(f, "Compression")
	delete(f, "BandwidthAlert")
	delete(f, "RangeOriginPull")
	delete(f, "FollowRedirect")
	delete(f, "ErrorPage")
	delete(f, "RequestHeader")
	delete(f, "ResponseHeader")
	delete(f, "DownstreamCapping")
	delete(f, "CacheKey")
	delete(f, "ResponseHeaderCache")
	delete(f, "VideoSeek")
	delete(f, "Cache")
	delete(f, "OriginPullOptimization")
	delete(f, "Https")
	delete(f, "Authentication")
	delete(f, "Seo")
	delete(f, "ForceRedirect")
	delete(f, "Referer")
	delete(f, "MaxAge")
	delete(f, "SpecificConfig")
	delete(f, "ServiceType")
	delete(f, "Area")
	delete(f, "OriginPullTimeout")
	delete(f, "AwsPrivateAccess")
	delete(f, "UserAgentFilter")
	delete(f, "AccessControl")
	delete(f, "UrlRedirect")
	delete(f, "AccessPort")
	delete(f, "AdvancedAuthentication")
	delete(f, "OriginAuthentication")
	delete(f, "Ipv6Access")
	delete(f, "OfflineCache")
	delete(f, "OriginCombine")
	delete(f, "PostMaxSize")
	delete(f, "Quic")
	delete(f, "OssPrivateAccess")
	delete(f, "WebSocket")
	delete(f, "RemoteAuthentication")
	delete(f, "ShareCname")
	delete(f, "HwPrivateAccess")
	delete(f, "QnPrivateAccess")
	delete(f, "OthersPrivateAccess")
	delete(f, "HttpsBilling")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDomainConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDomainConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateDomainConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDomainConfigResponseParams `json:"Response"`
}

func (r *UpdateDomainConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDomainConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateImageConfigRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// WebpAdapter配置项
	WebpAdapter *WebpAdapter `json:"WebpAdapter,omitnil" name:"WebpAdapter"`

	// TpgAdapter配置项
	TpgAdapter *TpgAdapter `json:"TpgAdapter,omitnil" name:"TpgAdapter"`

	// GuetzliAdapter配置项
	GuetzliAdapter *GuetzliAdapter `json:"GuetzliAdapter,omitnil" name:"GuetzliAdapter"`

	// AvifAdapter配置项
	AvifAdapter *AvifAdapter `json:"AvifAdapter,omitnil" name:"AvifAdapter"`
}

type UpdateImageConfigRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// WebpAdapter配置项
	WebpAdapter *WebpAdapter `json:"WebpAdapter,omitnil" name:"WebpAdapter"`

	// TpgAdapter配置项
	TpgAdapter *TpgAdapter `json:"TpgAdapter,omitnil" name:"TpgAdapter"`

	// GuetzliAdapter配置项
	GuetzliAdapter *GuetzliAdapter `json:"GuetzliAdapter,omitnil" name:"GuetzliAdapter"`

	// AvifAdapter配置项
	AvifAdapter *AvifAdapter `json:"AvifAdapter,omitnil" name:"AvifAdapter"`
}

func (r *UpdateImageConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateImageConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "WebpAdapter")
	delete(f, "TpgAdapter")
	delete(f, "GuetzliAdapter")
	delete(f, "AvifAdapter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateImageConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateImageConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateImageConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateImageConfigResponseParams `json:"Response"`
}

func (r *UpdateImageConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateImageConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePayTypeRequestParams struct {
	// 计费区域，mainland或overseas。
	Area *string `json:"Area,omitnil" name:"Area"`

	// 计费类型，flux或bandwidth。
	PayType *string `json:"PayType,omitnil" name:"PayType"`
}

type UpdatePayTypeRequest struct {
	*tchttp.BaseRequest
	
	// 计费区域，mainland或overseas。
	Area *string `json:"Area,omitnil" name:"Area"`

	// 计费类型，flux或bandwidth。
	PayType *string `json:"PayType,omitnil" name:"PayType"`
}

func (r *UpdatePayTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePayTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Area")
	delete(f, "PayType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdatePayTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePayTypeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdatePayTypeResponse struct {
	*tchttp.BaseResponse
	Response *UpdatePayTypeResponseParams `json:"Response"`
}

func (r *UpdatePayTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePayTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateScdnDomainRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Web 攻击防护（WAF）配置
	Waf *ScdnWafConfig `json:"Waf,omitnil" name:"Waf"`

	// 自定义防护策略配置
	Acl *ScdnAclConfig `json:"Acl,omitnil" name:"Acl"`

	// CC 防护配置，目前 CC 防护默认开启
	CC *ScdnConfig `json:"CC,omitnil" name:"CC"`

	// DDOS 防护配置，目前 DDoS 防护默认开启
	Ddos *ScdnDdosConfig `json:"Ddos,omitnil" name:"Ddos"`

	// BOT 防护配置
	Bot *ScdnBotConfig `json:"Bot,omitnil" name:"Bot"`
}

type UpdateScdnDomainRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Web 攻击防护（WAF）配置
	Waf *ScdnWafConfig `json:"Waf,omitnil" name:"Waf"`

	// 自定义防护策略配置
	Acl *ScdnAclConfig `json:"Acl,omitnil" name:"Acl"`

	// CC 防护配置，目前 CC 防护默认开启
	CC *ScdnConfig `json:"CC,omitnil" name:"CC"`

	// DDOS 防护配置，目前 DDoS 防护默认开启
	Ddos *ScdnDdosConfig `json:"Ddos,omitnil" name:"Ddos"`

	// BOT 防护配置
	Bot *ScdnBotConfig `json:"Bot,omitnil" name:"Bot"`
}

func (r *UpdateScdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateScdnDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Waf")
	delete(f, "Acl")
	delete(f, "CC")
	delete(f, "Ddos")
	delete(f, "Bot")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateScdnDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateScdnDomainResponseParams struct {
	// 提交结果，Success表示成功
	Result *string `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateScdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *UpdateScdnDomainResponseParams `json:"Response"`
}

func (r *UpdateScdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateScdnDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UrlRecord struct {
	// 状态(disable表示封禁，enable表示解封)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 对应的url
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealUrl *string `json:"RealUrl,omitnil" name:"RealUrl"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type UrlRedirect struct {
	// 访问URL重写配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 访问URL重写规则，当Switch为on时必填，规则数量最大为10个。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathRules []*UrlRedirectRule `json:"PathRules,omitnil" name:"PathRules"`
}

type UrlRedirectRule struct {
	// 重定向状态码，301 | 302
	RedirectStatusCode *int64 `json:"RedirectStatusCode,omitnil" name:"RedirectStatusCode"`

	// 待匹配的Url，仅支持Url路径，不支持参数。默认完全匹配，支持通配符 *，最多支持5个通配符，最大长度1024字符。
	Pattern *string `json:"Pattern,omitnil" name:"Pattern"`

	// 目标URL，必须以“/”开头，不包含参数部分。最大长度1024字符。可使用$1, $2, $3, $4, $5分别捕获匹配路径中的通配符号，最多支持10个捕获值。
	RedirectUrl *string `json:"RedirectUrl,omitnil" name:"RedirectUrl"`

	// 目标host，必须以http://或https://开头，并填写标准格式域名，如果不填写，默认为http:// + 当前域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectHost *string `json:"RedirectHost,omitnil" name:"RedirectHost"`

	// 指定是全路径配置还是任意匹配
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullMatch *bool `json:"FullMatch,omitnil" name:"FullMatch"`
}

type UserAgentFilter struct {
	// UserAgent黑白名单配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// UA黑白名单生效规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterRules []*UserAgentFilterRule `json:"FilterRules,omitnil" name:"FilterRules"`
}

type UserAgentFilterRule struct {
	// 访问路径生效类型
	// all: 所有访问路径生效
	// file: 根据文件后缀类型生效
	// directory: 根据目录生效
	// path: 根据完整访问路径生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *string `json:"RuleType,omitnil" name:"RuleType"`

	// 访问路径生效内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	RulePaths []*string `json:"RulePaths,omitnil" name:"RulePaths"`

	// UserAgent列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAgents []*string `json:"UserAgents,omitnil" name:"UserAgents"`

	// 黑名单或白名单，blacklist或whitelist
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterType *string `json:"FilterType,omitnil" name:"FilterType"`
}

// Predefined struct for user
type VerifyDomainRecordRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 验证方式
	// dns: DNS 解析验证（默认值）
	// file: 文件验证
	VerifyType *string `json:"VerifyType,omitnil" name:"VerifyType"`
}

type VerifyDomainRecordRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 验证方式
	// dns: DNS 解析验证（默认值）
	// file: 文件验证
	VerifyType *string `json:"VerifyType,omitnil" name:"VerifyType"`
}

func (r *VerifyDomainRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyDomainRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "VerifyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyDomainRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyDomainRecordResponseParams struct {
	// 是否验证成功
	Result *bool `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type VerifyDomainRecordResponse struct {
	*tchttp.BaseResponse
	Response *VerifyDomainRecordResponseParams `json:"Response"`
}

func (r *VerifyDomainRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyDomainRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VideoSeek struct {
	// 视频拖拽配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type ViolationUrl struct {
	// ID
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 违规资源原始访问 URL
	RealUrl *string `json:"RealUrl,omitnil" name:"RealUrl"`

	// 快照路径，用于控制台展示违规内容快照
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 违规资源当前状态
	// forbid：已封禁
	// release：已解封
	// delay ： 延迟处理
	// reject ：申诉驳回，状态仍为封禁状态
	// complain：申诉进行中
	UrlStatus *string `json:"UrlStatus,omitnil" name:"UrlStatus"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type WafSubRuleStatus struct {
	// Waf子规则开关状态配置开关，取值有：
	// on：开启
	// off：关闭
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 规则id列表
	SubIds []*int64 `json:"SubIds,omitnil" name:"SubIds"`
}

type WebSocket struct {
	// WebSocket 超时配置开关，取值有：
	// on：开启，可以调整超时时间
	// off：关闭，平台仍支持WebSocket连接，此时超时时间默认为15秒
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 设置超时时间，单位为秒，最大超时时间300秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timeout *int64 `json:"Timeout,omitnil" name:"Timeout"`
}

type WebpAdapter struct {
	// 图片优化-WebpAdapter配置开关，取值有：
	// on：开启
	// off：关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}