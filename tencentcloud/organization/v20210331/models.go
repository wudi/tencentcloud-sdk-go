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

package v20210331

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AddOrganizationMemberEmailRequestParams struct {
	// 成员Uin
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 邮箱地址
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 国际区号
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`
}

type AddOrganizationMemberEmailRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 邮箱地址
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 国际区号
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`
}

func (r *AddOrganizationMemberEmailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddOrganizationMemberEmailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "Email")
	delete(f, "CountryCode")
	delete(f, "Phone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddOrganizationMemberEmailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddOrganizationMemberEmailResponseParams struct {
	// 绑定Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindId *uint64 `json:"BindId,omitnil,omitempty" name:"BindId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddOrganizationMemberEmailResponse struct {
	*tchttp.BaseResponse
	Response *AddOrganizationMemberEmailResponseParams `json:"Response"`
}

func (r *AddOrganizationMemberEmailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddOrganizationMemberEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddOrganizationNodeRequestParams struct {
	// 父节点ID。可以调用DescribeOrganizationNodes获取
	ParentNodeId *uint64 `json:"ParentNodeId,omitnil,omitempty" name:"ParentNodeId"`

	// 节点名称。最大长度为40个字符，支持英文字母、数字、汉字、符号+@、&._[]-
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type AddOrganizationNodeRequest struct {
	*tchttp.BaseRequest
	
	// 父节点ID。可以调用DescribeOrganizationNodes获取
	ParentNodeId *uint64 `json:"ParentNodeId,omitnil,omitempty" name:"ParentNodeId"`

	// 节点名称。最大长度为40个字符，支持英文字母、数字、汉字、符号+@、&._[]-
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *AddOrganizationNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddOrganizationNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ParentNodeId")
	delete(f, "Name")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddOrganizationNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddOrganizationNodeResponseParams struct {
	// 节点ID。
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddOrganizationNodeResponse struct {
	*tchttp.BaseResponse
	Response *AddOrganizationNodeResponseParams `json:"Response"`
}

func (r *AddOrganizationNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddOrganizationNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddShareUnitMembersRequestParams struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 共享成员列表。最大10个。
	Members []*ShareMember `json:"Members,omitnil,omitempty" name:"Members"`
}

type AddShareUnitMembersRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 共享成员列表。最大10个。
	Members []*ShareMember `json:"Members,omitnil,omitempty" name:"Members"`
}

func (r *AddShareUnitMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddShareUnitMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UnitId")
	delete(f, "Area")
	delete(f, "Members")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddShareUnitMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddShareUnitMembersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddShareUnitMembersResponse struct {
	*tchttp.BaseResponse
	Response *AddShareUnitMembersResponseParams `json:"Response"`
}

func (r *AddShareUnitMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddShareUnitMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddShareUnitRequestParams struct {
	// 共享单元名称。仅支持大小写字母、数字、-、以及_的组合，3-128个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 共享单元地域。可通过接口DescribeShareAreas获取支持共享的地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 共享单元描述。最大128个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type AddShareUnitRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元名称。仅支持大小写字母、数字、-、以及_的组合，3-128个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 共享单元地域。可通过接口DescribeShareAreas获取支持共享的地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 共享单元描述。最大128个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *AddShareUnitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddShareUnitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Area")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddShareUnitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddShareUnitResourcesRequestParams struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 共享资源类型。支持共享的资源类型,请参见[资源共享概述](https://cloud.tencent.com/document/product/850/59489)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 共享资源列表。最大10个。
	Resources []*ProductResource `json:"Resources,omitnil,omitempty" name:"Resources"`
}

type AddShareUnitResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 共享资源类型。支持共享的资源类型,请参见[资源共享概述](https://cloud.tencent.com/document/product/850/59489)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 共享资源列表。最大10个。
	Resources []*ProductResource `json:"Resources,omitnil,omitempty" name:"Resources"`
}

func (r *AddShareUnitResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddShareUnitResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UnitId")
	delete(f, "Area")
	delete(f, "Type")
	delete(f, "Resources")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddShareUnitResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddShareUnitResourcesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddShareUnitResourcesResponse struct {
	*tchttp.BaseResponse
	Response *AddShareUnitResourcesResponseParams `json:"Response"`
}

func (r *AddShareUnitResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddShareUnitResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddShareUnitResponseParams struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddShareUnitResponse struct {
	*tchttp.BaseResponse
	Response *AddShareUnitResponseParams `json:"Response"`
}

func (r *AddShareUnitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddShareUnitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuthNode struct {
	// 互信主体关系ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelationId *int64 `json:"RelationId,omitnil,omitempty" name:"RelationId"`

	// 互信主体名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthName *string `json:"AuthName,omitnil,omitempty" name:"AuthName"`

	// 主体管理员
	// 注意：此字段可能返回 null，表示取不到有效值。
	Manager *MemberMainInfo `json:"Manager,omitnil,omitempty" name:"Manager"`
}

// Predefined struct for user
type BindOrganizationMemberAuthAccountRequestParams struct {
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 策略ID。可以调用DescribeOrganizationMemberPolicies获取
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 组织管理员子账号Uin列表。最大5个
	OrgSubAccountUins []*int64 `json:"OrgSubAccountUins,omitnil,omitempty" name:"OrgSubAccountUins"`
}

type BindOrganizationMemberAuthAccountRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 策略ID。可以调用DescribeOrganizationMemberPolicies获取
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 组织管理员子账号Uin列表。最大5个
	OrgSubAccountUins []*int64 `json:"OrgSubAccountUins,omitnil,omitempty" name:"OrgSubAccountUins"`
}

func (r *BindOrganizationMemberAuthAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindOrganizationMemberAuthAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "PolicyId")
	delete(f, "OrgSubAccountUins")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindOrganizationMemberAuthAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindOrganizationMemberAuthAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindOrganizationMemberAuthAccountResponse struct {
	*tchttp.BaseResponse
	Response *BindOrganizationMemberAuthAccountResponseParams `json:"Response"`
}

func (r *BindOrganizationMemberAuthAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindOrganizationMemberAuthAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelOrganizationMemberAuthAccountRequestParams struct {
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 策略ID。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 组织子账号Uin。
	OrgSubAccountUin *int64 `json:"OrgSubAccountUin,omitnil,omitempty" name:"OrgSubAccountUin"`
}

type CancelOrganizationMemberAuthAccountRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 策略ID。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 组织子账号Uin。
	OrgSubAccountUin *int64 `json:"OrgSubAccountUin,omitnil,omitempty" name:"OrgSubAccountUin"`
}

func (r *CancelOrganizationMemberAuthAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelOrganizationMemberAuthAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "PolicyId")
	delete(f, "OrgSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelOrganizationMemberAuthAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelOrganizationMemberAuthAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelOrganizationMemberAuthAccountResponse struct {
	*tchttp.BaseResponse
	Response *CancelOrganizationMemberAuthAccountResponseParams `json:"Response"`
}

func (r *CancelOrganizationMemberAuthAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelOrganizationMemberAuthAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckAccountDeleteRequestParams struct {
	// 成员uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

type CheckAccountDeleteRequest struct {
	*tchttp.BaseRequest
	
	// 成员uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

func (r *CheckAccountDeleteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAccountDeleteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckAccountDeleteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckAccountDeleteResponseParams struct {
	// 成员是否允许删除。 true-是、false-否
	AllowDelete *bool `json:"AllowDelete,omitnil,omitempty" name:"AllowDelete"`

	// 不允许删除原因。
	NotAllowReason *NotAllowReason `json:"NotAllowReason,omitnil,omitempty" name:"NotAllowReason"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckAccountDeleteResponse struct {
	*tchttp.BaseResponse
	Response *CheckAccountDeleteResponseParams `json:"Response"`
}

func (r *CheckAccountDeleteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAccountDeleteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationIdentityRequestParams struct {
	// 身份名称
	IdentityAliasName *string `json:"IdentityAliasName,omitnil,omitempty" name:"IdentityAliasName"`

	// 身份策略
	IdentityPolicy []*IdentityPolicy `json:"IdentityPolicy,omitnil,omitempty" name:"IdentityPolicy"`

	// 身份描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateOrganizationIdentityRequest struct {
	*tchttp.BaseRequest
	
	// 身份名称
	IdentityAliasName *string `json:"IdentityAliasName,omitnil,omitempty" name:"IdentityAliasName"`

	// 身份策略
	IdentityPolicy []*IdentityPolicy `json:"IdentityPolicy,omitnil,omitempty" name:"IdentityPolicy"`

	// 身份描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateOrganizationIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdentityAliasName")
	delete(f, "IdentityPolicy")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationIdentityResponseParams struct {
	// 身份ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationIdentityResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationIdentityResponseParams `json:"Response"`
}

func (r *CreateOrganizationIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMemberAuthIdentityRequestParams struct {
	// 成员uin列表。最多10个
	MemberUins []*uint64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 身份Id列表。最多5个
	IdentityIds []*uint64 `json:"IdentityIds,omitnil,omitempty" name:"IdentityIds"`
}

type CreateOrganizationMemberAuthIdentityRequest struct {
	*tchttp.BaseRequest
	
	// 成员uin列表。最多10个
	MemberUins []*uint64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 身份Id列表。最多5个
	IdentityIds []*uint64 `json:"IdentityIds,omitnil,omitempty" name:"IdentityIds"`
}

func (r *CreateOrganizationMemberAuthIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberAuthIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUins")
	delete(f, "IdentityIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationMemberAuthIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMemberAuthIdentityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationMemberAuthIdentityResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationMemberAuthIdentityResponseParams `json:"Response"`
}

func (r *CreateOrganizationMemberAuthIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberAuthIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMemberPolicyRequestParams struct {
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 策略名。最大长度为128个字符，支持英文字母、数字、符号+=,.@_-
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 成员访问身份ID。可以调用DescribeOrganizationMemberAuthIdentities获取
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateOrganizationMemberPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 策略名。最大长度为128个字符，支持英文字母、数字、符号+=,.@_-
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 成员访问身份ID。可以调用DescribeOrganizationMemberAuthIdentities获取
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateOrganizationMemberPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "PolicyName")
	delete(f, "IdentityId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationMemberPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMemberPolicyResponseParams struct {
	// 策略ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationMemberPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationMemberPolicyResponseParams `json:"Response"`
}

func (r *CreateOrganizationMemberPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMemberRequestParams struct {
	// 成员名称。最大长度为25个字符，支持英文字母、数字、汉字、符号+@、&._[]-:,
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 关系策略。取值：Financial
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 成员财务权限ID列表。取值：1-查看账单、2-查看余额、3-资金划拨、4-合并出账、5-开票、6-优惠继承、7-代付费，1、2 默认必须
	PermissionIds []*uint64 `json:"PermissionIds,omitnil,omitempty" name:"PermissionIds"`

	// 成员所属部门的节点ID。可以调用DescribeOrganizationNodes获取
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 账号名称。最大长度为25个字符，支持英文字母、数字、汉字、符号+@、&._[]-:,
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 成员创建记录ID。创建异常重试时需要
	RecordId *int64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 代付者Uin。成员代付费时需要
	PayUin *string `json:"PayUin,omitnil,omitempty" name:"PayUin"`

	// 成员访问身份ID列表。可以调用ListOrganizationIdentity获取，1默认支持
	IdentityRoleID []*uint64 `json:"IdentityRoleID,omitnil,omitempty" name:"IdentityRoleID"`

	// 认证主体关系ID。给不同主体创建成员时需要，可以调用DescribeOrganizationAuthNode获取
	AuthRelationId *int64 `json:"AuthRelationId,omitnil,omitempty" name:"AuthRelationId"`
}

type CreateOrganizationMemberRequest struct {
	*tchttp.BaseRequest
	
	// 成员名称。最大长度为25个字符，支持英文字母、数字、汉字、符号+@、&._[]-:,
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 关系策略。取值：Financial
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 成员财务权限ID列表。取值：1-查看账单、2-查看余额、3-资金划拨、4-合并出账、5-开票、6-优惠继承、7-代付费，1、2 默认必须
	PermissionIds []*uint64 `json:"PermissionIds,omitnil,omitempty" name:"PermissionIds"`

	// 成员所属部门的节点ID。可以调用DescribeOrganizationNodes获取
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 账号名称。最大长度为25个字符，支持英文字母、数字、汉字、符号+@、&._[]-:,
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 成员创建记录ID。创建异常重试时需要
	RecordId *int64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 代付者Uin。成员代付费时需要
	PayUin *string `json:"PayUin,omitnil,omitempty" name:"PayUin"`

	// 成员访问身份ID列表。可以调用ListOrganizationIdentity获取，1默认支持
	IdentityRoleID []*uint64 `json:"IdentityRoleID,omitnil,omitempty" name:"IdentityRoleID"`

	// 认证主体关系ID。给不同主体创建成员时需要，可以调用DescribeOrganizationAuthNode获取
	AuthRelationId *int64 `json:"AuthRelationId,omitnil,omitempty" name:"AuthRelationId"`
}

func (r *CreateOrganizationMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "PolicyType")
	delete(f, "PermissionIds")
	delete(f, "NodeId")
	delete(f, "AccountName")
	delete(f, "Remark")
	delete(f, "RecordId")
	delete(f, "PayUin")
	delete(f, "IdentityRoleID")
	delete(f, "AuthRelationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMemberResponseParams struct {
	// 成员Uin。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *int64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationMemberResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationMemberResponseParams `json:"Response"`
}

func (r *CreateOrganizationMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMembersPolicyRequestParams struct {
	// 成员Uin列表。最多10个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 策略名。长度1～128个字符，支持英文字母、数字、符号+=,.@_-
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 成员访问身份ID。
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 策略描述。最大长度为128个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateOrganizationMembersPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin列表。最多10个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 策略名。长度1～128个字符，支持英文字母、数字、符号+=,.@_-
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 成员访问身份ID。
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 策略描述。最大长度为128个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateOrganizationMembersPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMembersPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUins")
	delete(f, "PolicyName")
	delete(f, "IdentityId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationMembersPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMembersPolicyResponseParams struct {
	// 策略ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationMembersPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationMembersPolicyResponseParams `json:"Response"`
}

func (r *CreateOrganizationMembersPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMembersPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationRequestParams struct {

}

type CreateOrganizationRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreateOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationResponseParams struct {
	// 企业组织ID
	OrgId *uint64 `json:"OrgId,omitnil,omitempty" name:"OrgId"`

	// 创建者昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationResponseParams `json:"Response"`
}

func (r *CreateOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccountRequestParams struct {
	// 成员uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

type DeleteAccountRequest struct {
	*tchttp.BaseRequest
	
	// 成员uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

func (r *DeleteAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAccountResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccountResponseParams `json:"Response"`
}

func (r *DeleteAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationIdentityRequestParams struct {
	// 身份ID
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`
}

type DeleteOrganizationIdentityRequest struct {
	*tchttp.BaseRequest
	
	// 身份ID
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`
}

func (r *DeleteOrganizationIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdentityId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationIdentityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrganizationIdentityResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationIdentityResponseParams `json:"Response"`
}

func (r *DeleteOrganizationIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMemberAuthIdentityRequestParams struct {
	// 成员uin。
	MemberUin *uint64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 身份Id。
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`
}

type DeleteOrganizationMemberAuthIdentityRequest struct {
	*tchttp.BaseRequest
	
	// 成员uin。
	MemberUin *uint64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 身份Id。
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`
}

func (r *DeleteOrganizationMemberAuthIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMemberAuthIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "IdentityId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationMemberAuthIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMemberAuthIdentityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrganizationMemberAuthIdentityResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationMemberAuthIdentityResponseParams `json:"Response"`
}

func (r *DeleteOrganizationMemberAuthIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMemberAuthIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMembersPolicyRequestParams struct {
	// 访问策略ID。
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type DeleteOrganizationMembersPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 访问策略ID。
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

func (r *DeleteOrganizationMembersPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMembersPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationMembersPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMembersPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrganizationMembersPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationMembersPolicyResponseParams `json:"Response"`
}

func (r *DeleteOrganizationMembersPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMembersPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMembersRequestParams struct {
	// 被删除成员的UIN列表。
	MemberUin []*int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

type DeleteOrganizationMembersRequest struct {
	*tchttp.BaseRequest
	
	// 被删除成员的UIN列表。
	MemberUin []*int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

func (r *DeleteOrganizationMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMembersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrganizationMembersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationMembersResponseParams `json:"Response"`
}

func (r *DeleteOrganizationMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationNodesRequestParams struct {
	// 节点ID列表。
	NodeId []*int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

type DeleteOrganizationNodesRequest struct {
	*tchttp.BaseRequest
	
	// 节点ID列表。
	NodeId []*int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

func (r *DeleteOrganizationNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationNodesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrganizationNodesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationNodesResponseParams `json:"Response"`
}

func (r *DeleteOrganizationNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationRequestParams struct {

}

type DeleteOrganizationRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DeleteOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationResponseParams `json:"Response"`
}

func (r *DeleteOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteShareUnitMembersRequestParams struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 成员列表。
	Members []*ShareMember `json:"Members,omitnil,omitempty" name:"Members"`
}

type DeleteShareUnitMembersRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 成员列表。
	Members []*ShareMember `json:"Members,omitnil,omitempty" name:"Members"`
}

func (r *DeleteShareUnitMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShareUnitMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UnitId")
	delete(f, "Area")
	delete(f, "Members")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteShareUnitMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteShareUnitMembersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteShareUnitMembersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteShareUnitMembersResponseParams `json:"Response"`
}

func (r *DeleteShareUnitMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShareUnitMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteShareUnitRequestParams struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`
}

type DeleteShareUnitRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`
}

func (r *DeleteShareUnitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShareUnitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UnitId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteShareUnitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteShareUnitResourcesRequestParams struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 资源类型。支持共享的资源类型。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 共享资源列表。最大10个。
	Resources []*ShareResource `json:"Resources,omitnil,omitempty" name:"Resources"`
}

type DeleteShareUnitResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 资源类型。支持共享的资源类型。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 共享资源列表。最大10个。
	Resources []*ShareResource `json:"Resources,omitnil,omitempty" name:"Resources"`
}

func (r *DeleteShareUnitResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShareUnitResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UnitId")
	delete(f, "Area")
	delete(f, "Type")
	delete(f, "Resources")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteShareUnitResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteShareUnitResourcesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteShareUnitResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteShareUnitResourcesResponseParams `json:"Response"`
}

func (r *DeleteShareUnitResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShareUnitResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteShareUnitResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteShareUnitResponse struct {
	*tchttp.BaseResponse
	Response *DeleteShareUnitResponseParams `json:"Response"`
}

func (r *DeleteShareUnitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShareUnitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationAuthNodeRequestParams struct {
	// 偏移量。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。最大50
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 互信主体名称。
	AuthName *string `json:"AuthName,omitnil,omitempty" name:"AuthName"`
}

type DescribeOrganizationAuthNodeRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。最大50
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 互信主体名称。
	AuthName *string `json:"AuthName,omitnil,omitempty" name:"AuthName"`
}

func (r *DescribeOrganizationAuthNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationAuthNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AuthName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationAuthNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationAuthNodeResponseParams struct {
	// 总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 条目详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*AuthNode `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationAuthNodeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationAuthNodeResponseParams `json:"Response"`
}

func (r *DescribeOrganizationAuthNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationAuthNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationFinancialByMemberRequestParams struct {
	// 查询开始月份。格式：yyyy-mm，例如：2021-01。
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 限制数目。取值范围：1~50，默认值：10	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结束月份。格式：yyyy-mm，例如：2021-01,默认值为查询开始月份。
	EndMonth *string `json:"EndMonth,omitnil,omitempty" name:"EndMonth"`

	// 查询成员列表。 最大100个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 查询产品列表。 最大100个
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`
}

type DescribeOrganizationFinancialByMemberRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始月份。格式：yyyy-mm，例如：2021-01。
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 限制数目。取值范围：1~50，默认值：10	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结束月份。格式：yyyy-mm，例如：2021-01,默认值为查询开始月份。
	EndMonth *string `json:"EndMonth,omitnil,omitempty" name:"EndMonth"`

	// 查询成员列表。 最大100个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 查询产品列表。 最大100个
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`
}

func (r *DescribeOrganizationFinancialByMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationFinancialByMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EndMonth")
	delete(f, "MemberUins")
	delete(f, "ProductCodes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationFinancialByMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationFinancialByMemberResponseParams struct {
	// 当月总消耗。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCost *float64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 成员消耗详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*OrgMemberFinancial `json:"Items,omitnil,omitempty" name:"Items"`

	// 总数目。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationFinancialByMemberResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationFinancialByMemberResponseParams `json:"Response"`
}

func (r *DescribeOrganizationFinancialByMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationFinancialByMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationFinancialByMonthRequestParams struct {
	// 查询月数。取值范围：1~6，默认值：6
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询结束月份。格式：yyyy-mm，例如：2021-01
	EndMonth *string `json:"EndMonth,omitnil,omitempty" name:"EndMonth"`

	// 查询成员列表。 最大100个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 查询产品列表。 最大100个
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`
}

type DescribeOrganizationFinancialByMonthRequest struct {
	*tchttp.BaseRequest
	
	// 查询月数。取值范围：1~6，默认值：6
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询结束月份。格式：yyyy-mm，例如：2021-01
	EndMonth *string `json:"EndMonth,omitnil,omitempty" name:"EndMonth"`

	// 查询成员列表。 最大100个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 查询产品列表。 最大100个
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`
}

func (r *DescribeOrganizationFinancialByMonthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationFinancialByMonthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "EndMonth")
	delete(f, "MemberUins")
	delete(f, "ProductCodes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationFinancialByMonthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationFinancialByMonthResponseParams struct {
	// 产品消耗详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*OrgFinancialByMonth `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationFinancialByMonthResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationFinancialByMonthResponseParams `json:"Response"`
}

func (r *DescribeOrganizationFinancialByMonthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationFinancialByMonthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationFinancialByProductRequestParams struct {
	// 查询开始月份。格式：yyyy-mm，例如：2021-01
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 限制数目。取值范围：1~50，默认值：10	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结束月份。格式：yyyy-mm，例如：2021-01,默认值为查询开始月份
	EndMonth *string `json:"EndMonth,omitnil,omitempty" name:"EndMonth"`

	// 查询成员列表。 最大100个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 查询产品列表。 最大100个
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`
}

type DescribeOrganizationFinancialByProductRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始月份。格式：yyyy-mm，例如：2021-01
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 限制数目。取值范围：1~50，默认值：10	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结束月份。格式：yyyy-mm，例如：2021-01,默认值为查询开始月份
	EndMonth *string `json:"EndMonth,omitnil,omitempty" name:"EndMonth"`

	// 查询成员列表。 最大100个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 查询产品列表。 最大100个
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`
}

func (r *DescribeOrganizationFinancialByProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationFinancialByProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EndMonth")
	delete(f, "MemberUins")
	delete(f, "ProductCodes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationFinancialByProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationFinancialByProductResponseParams struct {
	// 当月总消耗。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCost *float64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 产品消耗详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*OrgProductFinancial `json:"Items,omitnil,omitempty" name:"Items"`

	// 总数目。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationFinancialByProductResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationFinancialByProductResponseParams `json:"Response"`
}

func (r *DescribeOrganizationFinancialByProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationFinancialByProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberAuthAccountsRequestParams struct {
	// 偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 策略ID。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type DescribeOrganizationMemberAuthAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 策略ID。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

func (r *DescribeOrganizationMemberAuthAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberAuthAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "MemberUin")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationMemberAuthAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberAuthAccountsResponseParams struct {
	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*OrgMemberAuthAccount `json:"Items,omitnil,omitempty" name:"Items"`

	// 总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationMemberAuthAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationMemberAuthAccountsResponseParams `json:"Response"`
}

func (r *DescribeOrganizationMemberAuthAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberAuthAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberAuthIdentitiesRequestParams struct {
	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50，默认值：10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 组织成员Uin。入参MemberUin与IdentityId至少填写一个
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 身份ID。入参MemberUin与IdentityId至少填写一个
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`
}

type DescribeOrganizationMemberAuthIdentitiesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50，默认值：10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 组织成员Uin。入参MemberUin与IdentityId至少填写一个
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 身份ID。入参MemberUin与IdentityId至少填写一个
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`
}

func (r *DescribeOrganizationMemberAuthIdentitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberAuthIdentitiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "MemberUin")
	delete(f, "IdentityId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationMemberAuthIdentitiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberAuthIdentitiesResponseParams struct {
	// 授权身份列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*OrgMemberAuthIdentity `json:"Items,omitnil,omitempty" name:"Items"`

	// 总数目。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationMemberAuthIdentitiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationMemberAuthIdentitiesResponseParams `json:"Response"`
}

func (r *DescribeOrganizationMemberAuthIdentitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberAuthIdentitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberEmailBindRequestParams struct {
	// 成员Uin
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

type DescribeOrganizationMemberEmailBindRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

func (r *DescribeOrganizationMemberEmailBindRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberEmailBindRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationMemberEmailBindRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberEmailBindResponseParams struct {
	// 绑定ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindId *uint64 `json:"BindId,omitnil,omitempty" name:"BindId"`

	// 申请时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplyTime *string `json:"ApplyTime,omitnil,omitempty" name:"ApplyTime"`

	// 邮箱地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 手机号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 绑定状态    未绑定：Unbound，待激活：Valid，绑定成功：Success，绑定失败：Failed
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindStatus *string `json:"BindStatus,omitnil,omitempty" name:"BindStatus"`

	// 绑定时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindTime *string `json:"BindTime,omitnil,omitempty" name:"BindTime"`

	// 失败说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 安全手机绑定状态  未绑定：0，已绑定：1
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneBind *uint64 `json:"PhoneBind,omitnil,omitempty" name:"PhoneBind"`

	// 国际区号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationMemberEmailBindResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationMemberEmailBindResponseParams `json:"Response"`
}

func (r *DescribeOrganizationMemberEmailBindResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberEmailBindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberPoliciesRequestParams struct {
	// 偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。最大50
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 搜索关键字。可用于策略名或描述搜索
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

type DescribeOrganizationMemberPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。最大50
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 搜索关键字。可用于策略名或描述搜索
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

func (r *DescribeOrganizationMemberPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "MemberUin")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationMemberPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberPoliciesResponseParams struct {
	// 列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*OrgMemberPolicy `json:"Items,omitnil,omitempty" name:"Items"`

	// 总数目。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationMemberPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationMemberPoliciesResponseParams `json:"Response"`
}

func (r *DescribeOrganizationMemberPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMembersRequestParams struct {
	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50，默认值：10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 国际站：en，国内站：zh
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`

	// 成员名称或者成员ID搜索。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 主体名称搜索。
	AuthName *string `json:"AuthName,omitnil,omitempty" name:"AuthName"`

	// 可信服务产品简称。可信服务管理员查询时必须指定
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeOrganizationMembersRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50，默认值：10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 国际站：en，国内站：zh
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`

	// 成员名称或者成员ID搜索。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 主体名称搜索。
	AuthName *string `json:"AuthName,omitnil,omitempty" name:"AuthName"`

	// 可信服务产品简称。可信服务管理员查询时必须指定
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeOrganizationMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Lang")
	delete(f, "SearchKey")
	delete(f, "AuthName")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMembersResponseParams struct {
	// 成员列表。
	Items []*OrgMember `json:"Items,omitnil,omitempty" name:"Items"`

	// 总数目。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationMembersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationMembersResponseParams `json:"Response"`
}

func (r *DescribeOrganizationMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationNodesRequestParams struct {
	// 限制数目。最大50
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeOrganizationNodesRequest struct {
	*tchttp.BaseRequest
	
	// 限制数目。最大50
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeOrganizationNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationNodesResponseParams struct {
	// 总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 列表详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*OrgNode `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationNodesResponseParams `json:"Response"`
}

func (r *DescribeOrganizationNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationRequestParams struct {
	// 国际站：en，国内站：zh
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`

	// 可信服务产品简称。查询是否该可信服务管理员时必须指定
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// 国际站：en，国内站：zh
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`

	// 可信服务产品简称。查询是否该可信服务管理员时必须指定
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Lang")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationResponseParams struct {
	// 企业组织ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgId *int64 `json:"OrgId,omitnil,omitempty" name:"OrgId"`

	// 创建者UIN。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostUin *int64 `json:"HostUin,omitnil,omitempty" name:"HostUin"`

	// 创建者昵称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 企业组织类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgType *int64 `json:"OrgType,omitnil,omitempty" name:"OrgType"`

	// 是否组织管理员。是：true ，否：false
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsManager *bool `json:"IsManager,omitnil,omitempty" name:"IsManager"`

	// 策略类型。财务管理：Financial
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgPolicyType *string `json:"OrgPolicyType,omitnil,omitempty" name:"OrgPolicyType"`

	// 策略名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgPolicyName *string `json:"OrgPolicyName,omitnil,omitempty" name:"OrgPolicyName"`

	// 成员财务权限列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgPermission []*OrgPermission `json:"OrgPermission,omitnil,omitempty" name:"OrgPermission"`

	// 组织根节点ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootNodeId *int64 `json:"RootNodeId,omitnil,omitempty" name:"RootNodeId"`

	// 组织创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 成员加入时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	JoinTime *string `json:"JoinTime,omitnil,omitempty" name:"JoinTime"`

	// 成员是否允许退出。允许：Allow，不允许：Denied
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAllowQuit *string `json:"IsAllowQuit,omitnil,omitempty" name:"IsAllowQuit"`

	// 代付者Uin。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayUin *string `json:"PayUin,omitnil,omitempty" name:"PayUin"`

	// 代付者名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayName *string `json:"PayName,omitnil,omitempty" name:"PayName"`

	// 是否可信服务管理员。是：true，否：false
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAssignManager *bool `json:"IsAssignManager,omitnil,omitempty" name:"IsAssignManager"`

	// 是否实名主体管理员。是：true，否：false
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAuthManager *bool `json:"IsAuthManager,omitnil,omitempty" name:"IsAuthManager"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationResponseParams `json:"Response"`
}

func (r *DescribeOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShareAreasRequestParams struct {
	// 国际站：en，国内站：zh
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`
}

type DescribeShareAreasRequest struct {
	*tchttp.BaseRequest
	
	// 国际站：en，国内站：zh
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`
}

func (r *DescribeShareAreasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShareAreasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Lang")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShareAreasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShareAreasResponseParams struct {
	// 详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*ShareArea `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeShareAreasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShareAreasResponseParams `json:"Response"`
}

func (r *DescribeShareAreasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShareAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShareUnitMembersRequestParams struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键字。支持成员uin搜索。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

type DescribeShareUnitMembersRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键字。支持成员uin搜索。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

func (r *DescribeShareUnitMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShareUnitMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UnitId")
	delete(f, "Area")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShareUnitMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShareUnitMembersResponseParams struct {
	// 总数目。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 共享单元成员列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*ShareUnitMember `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeShareUnitMembersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShareUnitMembersResponseParams `json:"Response"`
}

func (r *DescribeShareUnitMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShareUnitMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShareUnitResourcesRequestParams struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键字。支持产品资源ID搜索。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 共享资源类型。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeShareUnitResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键字。支持产品资源ID搜索。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 共享资源类型。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeShareUnitResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShareUnitResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UnitId")
	delete(f, "Area")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchKey")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShareUnitResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShareUnitResourcesResponseParams struct {
	// 总数目。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 共享单元资源列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*ShareUnitResource `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeShareUnitResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShareUnitResourcesResponseParams `json:"Response"`
}

func (r *DescribeShareUnitResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShareUnitResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShareUnitsRequestParams struct {
	// 共享单元地域。可通过接口DescribeShareAreas获取支持共享的地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键字。支持UnitId和Name搜索。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

type DescribeShareUnitsRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元地域。可通过接口DescribeShareAreas获取支持共享的地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键字。支持UnitId和Name搜索。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

func (r *DescribeShareUnitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShareUnitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Area")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShareUnitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShareUnitsResponseParams struct {
	// 总数目。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 共享单元列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*ManagerShareUnit `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeShareUnitsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShareUnitsResponseParams `json:"Response"`
}

func (r *DescribeShareUnitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShareUnitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IdentityPolicy struct {
	// CAM预设策略ID。PolicyType 为预设策略时有效且必选
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// CAM预设策略名称。PolicyType 为预设策略时有效且必选
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略类型。取值 1-自定义策略  2-预设策略；默认值2
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyType *uint64 `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 自定义策略内容，遵循CAM策略语法。PolicyType 为自定义策略时有效且必选
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyDocument *string `json:"PolicyDocument,omitnil,omitempty" name:"PolicyDocument"`
}

// Predefined struct for user
type ListOrganizationIdentityRequestParams struct {
	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。默认值：10。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 名称搜索关键字。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 身份ID搜索。
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份类型。取值范围 1-预设, 2-自定义
	IdentityType *uint64 `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`
}

type ListOrganizationIdentityRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。默认值：10。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 名称搜索关键字。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 身份ID搜索。
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份类型。取值范围 1-预设, 2-自定义
	IdentityType *uint64 `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`
}

func (r *ListOrganizationIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchKey")
	delete(f, "IdentityId")
	delete(f, "IdentityType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListOrganizationIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationIdentityResponseParams struct {
	// 总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 条目详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*OrgIdentity `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListOrganizationIdentityResponse struct {
	*tchttp.BaseResponse
	Response *ListOrganizationIdentityResponseParams `json:"Response"`
}

func (r *ListOrganizationIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManagerShareUnit struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 共享单元管理员Uin。
	Uin *int64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 共享单元管理员OwnerUin。
	OwnerUin *int64 `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 共享单元资源数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShareResourceNum *int64 `json:"ShareResourceNum,omitnil,omitempty" name:"ShareResourceNum"`

	// 共享单元成员数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShareMemberNum *int64 `json:"ShareMemberNum,omitnil,omitempty" name:"ShareMemberNum"`
}

type MemberIdentity struct {
	// 身份ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityAliasName *string `json:"IdentityAliasName,omitnil,omitempty" name:"IdentityAliasName"`
}

type MemberMainInfo struct {
	// 成员uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 成员名称j
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberName *string `json:"MemberName,omitnil,omitempty" name:"MemberName"`
}

// Predefined struct for user
type MoveOrganizationNodeMembersRequestParams struct {
	// 组织节点ID。
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 成员UIN列表。
	MemberUin []*int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

type MoveOrganizationNodeMembersRequest struct {
	*tchttp.BaseRequest
	
	// 组织节点ID。
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 成员UIN列表。
	MemberUin []*int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

func (r *MoveOrganizationNodeMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MoveOrganizationNodeMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeId")
	delete(f, "MemberUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MoveOrganizationNodeMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MoveOrganizationNodeMembersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type MoveOrganizationNodeMembersResponse struct {
	*tchttp.BaseResponse
	Response *MoveOrganizationNodeMembersResponseParams `json:"Response"`
}

func (r *MoveOrganizationNodeMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MoveOrganizationNodeMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NotAllowReason struct {
	// 是否创建的成员。true-是、false-否；成员不是创建的成员不允许删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCreateMember *bool `json:"IsCreateMember,omitnil,omitempty" name:"IsCreateMember"`

	// 成员删除许可。true-开启、false-关闭；成员删除许可关闭时不允许删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeletionPermission *bool `json:"DeletionPermission,omitnil,omitempty" name:"DeletionPermission"`

	// 是否可信服务委派管理员。true-是、false-否；成员是可信服务委派管理员不允许删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAssignManager *bool `json:"IsAssignManager,omitnil,omitempty" name:"IsAssignManager"`

	// 是否主体管理员。true-是、false-否；成员是主体管理员不允许删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAuthManager *bool `json:"IsAuthManager,omitnil,omitempty" name:"IsAuthManager"`

	// 是否共享资源管理员。true-是、false-否；成员是共享资源管理员不允许删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsShareManager *bool `json:"IsShareManager,omitnil,omitempty" name:"IsShareManager"`

	// 成员是否设置了操作审批。true-是、false-否；成员设置了操作审批时不允许删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateProcess *bool `json:"OperateProcess,omitnil,omitempty" name:"OperateProcess"`

	// 是否允许解除成员财务权限。true-是、false-否；成员不能解除财务权限时不允许删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingPermission *bool `json:"BillingPermission,omitnil,omitempty" name:"BillingPermission"`

	// 存在的资源列表。账号存在资源时不允许删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExistResources []*string `json:"ExistResources,omitnil,omitempty" name:"ExistResources"`

	// 检测失败的资源列表。账号有资源检测失败时不允许删除。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectFailedResources []*string `json:"DetectFailedResources,omitnil,omitempty" name:"DetectFailedResources"`
}

type OrgFinancialByMonth struct {
	// 记录ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 月份，格式：yyyy-mm，示例：2021-01。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 消耗金额，单元：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCost *float64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 比上月增长率%。正数增长，负数下降，空值无法统计。
	// 注意：此字段可能返回 null，表示取不到有效值。
	GrowthRate *string `json:"GrowthRate,omitnil,omitempty" name:"GrowthRate"`
}

type OrgIdentity struct {
	// 身份ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityAliasName *string `json:"IdentityAliasName,omitnil,omitempty" name:"IdentityAliasName"`

	// 描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 身份策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityPolicy []*IdentityPolicy `json:"IdentityPolicy,omitnil,omitempty" name:"IdentityPolicy"`

	// 身份类型。 1-预设、 2-自定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityType *uint64 `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`

	// 更新时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type OrgMember struct {
	// 成员Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 成员名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 成员类型，邀请：Invite， 创建：Create
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberType *string `json:"MemberType,omitnil,omitempty" name:"MemberType"`

	// 关系策略类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgPolicyType *string `json:"OrgPolicyType,omitnil,omitempty" name:"OrgPolicyType"`

	// 关系策略名
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgPolicyName *string `json:"OrgPolicyName,omitnil,omitempty" name:"OrgPolicyName"`

	// 关系策略权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgPermission []*OrgPermission `json:"OrgPermission,omitnil,omitempty" name:"OrgPermission"`

	// 所属节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 所属节点名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 是否允许成员退出。允许：Allow，不允许：Denied。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAllowQuit *string `json:"IsAllowQuit,omitnil,omitempty" name:"IsAllowQuit"`

	// 代付者Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayUin *string `json:"PayUin,omitnil,omitempty" name:"PayUin"`

	// 代付者名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayName *string `json:"PayName,omitnil,omitempty" name:"PayName"`

	// 管理身份
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgIdentity []*MemberIdentity `json:"OrgIdentity,omitnil,omitempty" name:"OrgIdentity"`

	// 安全信息绑定状态  未绑定：Unbound，待激活：Valid，绑定成功：Success，绑定失败：Failed
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindStatus *string `json:"BindStatus,omitnil,omitempty" name:"BindStatus"`

	// 成员权限状态 已确认：Confirmed ，待确认：UnConfirmed
	// 注意：此字段可能返回 null，表示取不到有效值。
	PermissionStatus *string `json:"PermissionStatus,omitnil,omitempty" name:"PermissionStatus"`
}

type OrgMemberAuthAccount struct {
	// 组织子账号Uin。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgSubAccountUin *int64 `json:"OrgSubAccountUin,omitnil,omitempty" name:"OrgSubAccountUin"`

	// 策略ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 身份ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份角色名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityRoleName *string `json:"IdentityRoleName,omitnil,omitempty" name:"IdentityRoleName"`

	// 身份角色别名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityRoleAliasName *string `json:"IdentityRoleAliasName,omitnil,omitempty" name:"IdentityRoleAliasName"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 子账号名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgSubAccountName *string `json:"OrgSubAccountName,omitnil,omitempty" name:"OrgSubAccountName"`
}

type OrgMemberAuthIdentity struct {
	// 身份ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份的角色名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityRoleName *string `json:"IdentityRoleName,omitnil,omitempty" name:"IdentityRoleName"`

	// 身份的角色别名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityRoleAliasName *string `json:"IdentityRoleAliasName,omitnil,omitempty" name:"IdentityRoleAliasName"`

	// 身份描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 首次配置成功的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后一次配置成功的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 身份类型。取值： 1-预设身份  2-自定义身份
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityType *uint64 `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`

	// 配置状态。取值：1-配置完成 2-需重新配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 成员Uin。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 成员名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberName *string `json:"MemberName,omitnil,omitempty" name:"MemberName"`
}

type OrgMemberFinancial struct {
	// 成员Uin。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 成员名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberName *string `json:"MemberName,omitnil,omitempty" name:"MemberName"`

	// 消耗金额，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCost *float64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 占比%。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ratio *string `json:"Ratio,omitnil,omitempty" name:"Ratio"`
}

type OrgMemberPolicy struct {
	// 策略ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 身份ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份角色名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityRoleName *string `json:"IdentityRoleName,omitnil,omitempty" name:"IdentityRoleName"`

	// 身份角色别名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityRoleAliasName *string `json:"IdentityRoleAliasName,omitnil,omitempty" name:"IdentityRoleAliasName"`

	// 描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type OrgNode struct {
	// 组织节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 父节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentNodeId *int64 `json:"ParentNodeId,omitnil,omitempty" name:"ParentNodeId"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type OrgPermission struct {
	// 权限Id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 权限名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type OrgProductFinancial struct {
	// 产品Code。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 产品名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 产品消耗，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCost *float64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 占比%。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ratio *string `json:"Ratio,omitnil,omitempty" name:"Ratio"`
}

type ProductResource struct {
	// 产品资源ID。
	ProductResourceId *string `json:"ProductResourceId,omitnil,omitempty" name:"ProductResourceId"`

	// 资源六段式最后一节
	//
	// Deprecated: ResourceGrantLast is deprecated.
	ResourceGrantLast *string `json:"ResourceGrantLast,omitnil,omitempty" name:"ResourceGrantLast"`
}

// Predefined struct for user
type QuitOrganizationRequestParams struct {
	// 企业组织ID
	OrgId *uint64 `json:"OrgId,omitnil,omitempty" name:"OrgId"`
}

type QuitOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// 企业组织ID
	OrgId *uint64 `json:"OrgId,omitnil,omitempty" name:"OrgId"`
}

func (r *QuitOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QuitOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrgId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QuitOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QuitOrganizationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QuitOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *QuitOrganizationResponseParams `json:"Response"`
}

func (r *QuitOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QuitOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShareArea struct {
	// 地域名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 地域标识。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 地域ID。
	AreaId *int64 `json:"AreaId,omitnil,omitempty" name:"AreaId"`
}

type ShareMember struct {
	// 共享成员Uin。
	ShareMemberUin *int64 `json:"ShareMemberUin,omitnil,omitempty" name:"ShareMemberUin"`
}

type ShareResource struct {
	// 共享资源ID。
	//
	// Deprecated: ResourceId is deprecated.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 产品资源ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductResourceId *string `json:"ProductResourceId,omitnil,omitempty" name:"ProductResourceId"`
}

type ShareUnitMember struct {
	// 共享成员Uin。
	ShareMemberUin *int64 `json:"ShareMemberUin,omitnil,omitempty" name:"ShareMemberUin"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type ShareUnitResource struct {
	// 共享资源ID。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 共享资源类型。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 产品资源ID。
	ProductResourceId *string `json:"ProductResourceId,omitnil,omitempty" name:"ProductResourceId"`

	// 共享单元成员数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SharedMemberNum *uint64 `json:"SharedMemberNum,omitnil,omitempty" name:"SharedMemberNum"`

	// 使用中共享单元成员数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SharedMemberUseNum *uint64 `json:"SharedMemberUseNum,omitnil,omitempty" name:"SharedMemberUseNum"`

	// 共享管理员OwnerUin。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShareManagerUin *int64 `json:"ShareManagerUin,omitnil,omitempty" name:"ShareManagerUin"`
}

// Predefined struct for user
type UpdateOrganizationIdentityRequestParams struct {
	// 身份ID
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 身份策略
	IdentityPolicy []*IdentityPolicy `json:"IdentityPolicy,omitnil,omitempty" name:"IdentityPolicy"`
}

type UpdateOrganizationIdentityRequest struct {
	*tchttp.BaseRequest
	
	// 身份ID
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 身份策略
	IdentityPolicy []*IdentityPolicy `json:"IdentityPolicy,omitnil,omitempty" name:"IdentityPolicy"`
}

func (r *UpdateOrganizationIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdentityId")
	delete(f, "Description")
	delete(f, "IdentityPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOrganizationIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationIdentityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateOrganizationIdentityResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOrganizationIdentityResponseParams `json:"Response"`
}

func (r *UpdateOrganizationIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationMemberEmailBindRequestParams struct {
	// 成员Uin
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 绑定ID
	BindId *int64 `json:"BindId,omitnil,omitempty" name:"BindId"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 国际区号
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`
}

type UpdateOrganizationMemberEmailBindRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 绑定ID
	BindId *int64 `json:"BindId,omitnil,omitempty" name:"BindId"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 国际区号
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`
}

func (r *UpdateOrganizationMemberEmailBindRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationMemberEmailBindRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "BindId")
	delete(f, "Email")
	delete(f, "CountryCode")
	delete(f, "Phone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOrganizationMemberEmailBindRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationMemberEmailBindResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateOrganizationMemberEmailBindResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOrganizationMemberEmailBindResponseParams `json:"Response"`
}

func (r *UpdateOrganizationMemberEmailBindResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationMemberEmailBindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationMemberRequestParams struct {
	// 成员Uin。
	MemberUin *uint64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 成员名称。最大长度为25个字符，支持英文字母、数字、汉字、符号+@、&._[]-:,
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注。最大长度为40个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 关系策略类型。PolicyType不为空，PermissionIds不能为空。取值：Financial
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 成员财务权限ID列表。PermissionIds不为空，PolicyType不能为空。
	// 取值：1-查看账单、2-查看余额、3-资金划拨、4-合并出账、5-开票、6-优惠继承、7-代付费、8-成本分析，如果有值，1、2 默认必须
	PermissionIds []*uint64 `json:"PermissionIds,omitnil,omitempty" name:"PermissionIds"`

	// 是否允许成员退出组织。取值：Allow-允许、Denied-不允许
	IsAllowQuit *string `json:"IsAllowQuit,omitnil,omitempty" name:"IsAllowQuit"`

	// 代付者Uin。成员财务权限有代付费时需要，取值为成员对应主体的主体管理员Uin
	PayUin *string `json:"PayUin,omitnil,omitempty" name:"PayUin"`
}

type UpdateOrganizationMemberRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin。
	MemberUin *uint64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 成员名称。最大长度为25个字符，支持英文字母、数字、汉字、符号+@、&._[]-:,
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注。最大长度为40个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 关系策略类型。PolicyType不为空，PermissionIds不能为空。取值：Financial
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 成员财务权限ID列表。PermissionIds不为空，PolicyType不能为空。
	// 取值：1-查看账单、2-查看余额、3-资金划拨、4-合并出账、5-开票、6-优惠继承、7-代付费、8-成本分析，如果有值，1、2 默认必须
	PermissionIds []*uint64 `json:"PermissionIds,omitnil,omitempty" name:"PermissionIds"`

	// 是否允许成员退出组织。取值：Allow-允许、Denied-不允许
	IsAllowQuit *string `json:"IsAllowQuit,omitnil,omitempty" name:"IsAllowQuit"`

	// 代付者Uin。成员财务权限有代付费时需要，取值为成员对应主体的主体管理员Uin
	PayUin *string `json:"PayUin,omitnil,omitempty" name:"PayUin"`
}

func (r *UpdateOrganizationMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "Name")
	delete(f, "Remark")
	delete(f, "PolicyType")
	delete(f, "PermissionIds")
	delete(f, "IsAllowQuit")
	delete(f, "PayUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOrganizationMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationMemberResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateOrganizationMemberResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOrganizationMemberResponseParams `json:"Response"`
}

func (r *UpdateOrganizationMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationNodeRequestParams struct {
	// 节点ID。
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 节点名称。最大长度为40个字符，支持英文字母、数字、汉字、符号+@、&._[]-
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type UpdateOrganizationNodeRequest struct {
	*tchttp.BaseRequest
	
	// 节点ID。
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 节点名称。最大长度为40个字符，支持英文字母、数字、汉字、符号+@、&._[]-
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *UpdateOrganizationNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeId")
	delete(f, "Name")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOrganizationNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationNodeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateOrganizationNodeResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOrganizationNodeResponseParams `json:"Response"`
}

func (r *UpdateOrganizationNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateShareUnitRequestParams struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元名称。仅支持大小写字母、数字、-、以及_的组合，3-128个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 共享单元描述。最大128个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UpdateShareUnitRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元名称。仅支持大小写字母、数字、-、以及_的组合，3-128个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 共享单元描述。最大128个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *UpdateShareUnitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateShareUnitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UnitId")
	delete(f, "Name")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateShareUnitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateShareUnitResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateShareUnitResponse struct {
	*tchttp.BaseResponse
	Response *UpdateShareUnitResponseParams `json:"Response"`
}

func (r *UpdateShareUnitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateShareUnitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}