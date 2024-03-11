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

package v20220501

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Annotation struct {
	// map表中的Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// map表中的Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type AutoscalingAdded struct {
	// 正在加入中的节点数量
	Joining *int64 `json:"Joining,omitnil,omitempty" name:"Joining"`

	// 初始化中的节点数量
	Initializing *int64 `json:"Initializing,omitnil,omitempty" name:"Initializing"`

	// 正常的节点数量
	Normal *int64 `json:"Normal,omitnil,omitempty" name:"Normal"`

	// 节点总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}

// Predefined struct for user
type DescribeClusterInstancesRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件列表:
	// InstanceIds(实例ID),InstanceType(实例类型：Regular，Native，Virtual，External),VagueIpAddress(模糊匹配IP),Labels(k8s节点label),NodePoolNames(节点池名称),VagueInstanceName(模糊匹配节点名),InstanceStates(节点状态),Unschedulable(是否封锁),NodePoolIds(节点池ID)
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序信息
	SortBy *SortBy `json:"SortBy,omitnil,omitempty" name:"SortBy"`
}

type DescribeClusterInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件列表:
	// InstanceIds(实例ID),InstanceType(实例类型：Regular，Native，Virtual，External),VagueIpAddress(模糊匹配IP),Labels(k8s节点label),NodePoolNames(节点池名称),VagueInstanceName(模糊匹配节点名),InstanceStates(节点状态),Unschedulable(是否封锁),NodePoolIds(节点池ID)
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序信息
	SortBy *SortBy `json:"SortBy,omitnil,omitempty" name:"SortBy"`
}

func (r *DescribeClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "SortBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterInstancesResponseParams struct {
	// 集群中实例总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集群中实例列表
	InstanceSet []*Instance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// 错误信息集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Errors []*string `json:"Errors,omitnil,omitempty" name:"Errors"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterInstancesResponseParams `json:"Response"`
}

func (r *DescribeClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodePoolsRequestParams struct {
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 查询过滤条件：
	// ·  NodePoolsName
	//     按照【节点池名】进行过滤。
	//     类型：String
	//     必选：否
	// 
	// ·  NodePoolsId
	//     按照【节点池id】进行过滤。
	//     类型：String
	//     必选：否
	// 
	// ·  tags
	//     按照【标签键值对】进行过滤。
	//     类型：String
	//     必选：否
	// 
	// ·  tag:tag-key
	//     按照【标签键值对】进行过滤。
	//     类型：String
	//     必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条数，默认20，最大为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeNodePoolsRequest struct {
	*tchttp.BaseRequest
	
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 查询过滤条件：
	// ·  NodePoolsName
	//     按照【节点池名】进行过滤。
	//     类型：String
	//     必选：否
	// 
	// ·  NodePoolsId
	//     按照【节点池id】进行过滤。
	//     类型：String
	//     必选：否
	// 
	// ·  tags
	//     按照【标签键值对】进行过滤。
	//     类型：String
	//     必选：否
	// 
	// ·  tag:tag-key
	//     按照【标签键值对】进行过滤。
	//     类型：String
	//     必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条数，默认20，最大为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeNodePoolsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodePoolsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNodePoolsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodePoolsResponseParams struct {
	// 节点池列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodePools []*NodePool `json:"NodePools,omitnil,omitempty" name:"NodePools"`

	// 资源总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNodePoolsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNodePoolsResponseParams `json:"Response"`
}

func (r *DescribeNodePoolsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodePoolsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExternalNodeInfo struct {
	// 第三方节点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// CPU核数，单位：核
	// 注意：此字段可能返回 null，表示取不到有效值。
	CPU *uint64 `json:"CPU,omitnil,omitempty" name:"CPU"`

	// 节点内存容量，单位：`GB`
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 第三方节点kubelet版本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	K8SVersion *string `json:"K8SVersion,omitnil,omitempty" name:"K8SVersion"`
}

type ExternalNodePoolInfo struct {
	// 第三方节点Runtime配置
	RuntimeConfig *RuntimeConfig `json:"RuntimeConfig,omitnil,omitempty" name:"RuntimeConfig"`

	// 节点数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodesNum *uint64 `json:"NodesNum,omitnil,omitempty" name:"NodesNum"`
}

type Filter struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type Instance struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点角色, MASTER, WORKER, ETCD, MASTER_ETCD,ALL, 默认为WORKER
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// 实例异常(或者处于初始化中)的原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedReason *string `json:"FailedReason,omitnil,omitempty" name:"FailedReason"`

	// 实例的状态
	// - initializing创建中
	// - running 运行中
	// - failed 异常
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// 是否不可调度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unschedulable *bool `json:"Unschedulable,omitnil,omitempty" name:"Unschedulable"`

	// 添加时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 节点内网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	LanIP *string `json:"LanIP,omitnil,omitempty" name:"LanIP"`

	// 资源池ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 原生节点参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Native *NativeNodeInfo `json:"Native,omitnil,omitempty" name:"Native"`

	// 普通节点参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Regular *RegularNodeInfo `json:"Regular,omitnil,omitempty" name:"Regular"`

	// 超级节点参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Super *SuperNodeInfo `json:"Super,omitnil,omitempty" name:"Super"`

	// 第三方节点参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	External *ExternalNodeInfo `json:"External,omitnil,omitempty" name:"External"`

	// 节点类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`
}

type InstanceAdvancedSettings struct {
	// 该节点属于podCIDR大小自定义模式时，可指定节点上运行的pod数量上限
	// 注意：此字段可能返回 null，表示取不到有效值。
	DesiredPodNumber *int64 `json:"DesiredPodNumber,omitnil,omitempty" name:"DesiredPodNumber"`

	// base64 编码的用户脚本，在初始化节点之前执行，目前只对添加已有节点生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreStartUserScript *string `json:"PreStartUserScript,omitnil,omitempty" name:"PreStartUserScript"`

	// 运行时描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeConfig *RuntimeConfig `json:"RuntimeConfig,omitnil,omitempty" name:"RuntimeConfig"`

	// base64 编码的用户脚本, 此脚本会在 k8s 组件运行后执行, 需要用户保证脚本的可重入及重试逻辑, 脚本及其生成的日志文件可在节点的 /data/ccs_userscript/ 路径查看, 如果要求节点需要在进行初始化完成后才可加入调度, 可配合 unschedulable 参数使用, 在 userScript 最后初始化完成后, 添加 kubectl uncordon nodename --kubeconfig=/root/.kube/config 命令使节点加入调度
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserScript *string `json:"UserScript,omitnil,omitempty" name:"UserScript"`

	// 节点相关的自定义参数信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraArgs *InstanceExtraArgs `json:"ExtraArgs,omitnil,omitempty" name:"ExtraArgs"`
}

type InstanceExtraArgs struct {
	// kubelet自定义参数，参数格式为["k1=v1", "k1=v2"]， 例如["root-dir=/var/lib/kubelet","feature-gates=PodShareProcessNamespace=true,DynamicKubeletConfig=true"]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kubelet []*string `json:"Kubelet,omitnil,omitempty" name:"Kubelet"`
}

type InternetAccessible struct {
	// 带宽
	MaxBandwidthOut *int64 `json:"MaxBandwidthOut,omitnil,omitempty" name:"MaxBandwidthOut"`

	// 网络计费方式
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 带宽包 ID
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`
}

type Label struct {
	// map表中的Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// map表中的Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ManuallyAdded struct {
	// 加入中的节点数量
	Joining *int64 `json:"Joining,omitnil,omitempty" name:"Joining"`

	// 初始化中的节点数量
	Initializing *int64 `json:"Initializing,omitnil,omitempty" name:"Initializing"`

	// 正常的节点数量
	Normal *int64 `json:"Normal,omitnil,omitempty" name:"Normal"`

	// 节点总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type NativeNodeInfo struct {
	// 节点名称
	MachineName *string `json:"MachineName,omitnil,omitempty" name:"MachineName"`

	// Machine 状态
	MachineState *string `json:"MachineState,omitnil,omitempty" name:"MachineState"`

	// Machine 所在可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 节点计费类型。PREPAID：包年包月；POSTPAID_BY_HOUR：按量计费（默认）；
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Machine 登录状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoginStatus *string `json:"LoginStatus,omitnil,omitempty" name:"LoginStatus"`

	// 是否开启缩容保护
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsProtectedFromScaleIn *bool `json:"IsProtectedFromScaleIn,omitnil,omitempty" name:"IsProtectedFromScaleIn"`

	// Machine 名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// CPU核数，单位：核
	CPU *uint64 `json:"CPU,omitnil,omitempty" name:"CPU"`

	// GPU核数，单位：核
	// 注意：此字段可能返回 null，表示取不到有效值。
	GPU *uint64 `json:"GPU,omitnil,omitempty" name:"GPU"`

	// 自动续费标识
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 节点计费模式（已弃用）
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 节点内存容量，单位：`GB`
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 公网带宽相关信息设置
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 机型所属机型族
	InstanceFamily *string `json:"InstanceFamily,omitnil,omitempty" name:"InstanceFamily"`

	// 节点内网 IP
	LanIp *string `json:"LanIp,omitnil,omitempty" name:"LanIp"`

	// 机型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 包年包月节点计费过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 安全组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroupIDs []*string `json:"SecurityGroupIDs,omitnil,omitempty" name:"SecurityGroupIDs"`

	// VPC 唯一 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网唯一 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// OS的名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsImage *string `json:"OsImage,omitnil,omitempty" name:"OsImage"`
}

type NativeNodePoolInfo struct {
	// 子网列表
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 安全组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type NodeCountSummary struct {
	// 手动管理的节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManuallyAdded *ManuallyAdded `json:"ManuallyAdded,omitnil,omitempty" name:"ManuallyAdded"`

	// 自动管理的节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoscalingAdded *AutoscalingAdded `json:"AutoscalingAdded,omitnil,omitempty" name:"AutoscalingAdded"`
}

type NodePool struct {
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池 ID
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 节点污点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Taints []*Taint `json:"Taints,omitnil,omitempty" name:"Taints"`

	// 是否开启删除保护
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`

	// 节点池类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 节点  Labels
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 节点池状态
	LifeState *string `json:"LifeState,omitnil,omitempty" name:"LifeState"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 节点池名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 原生节点池参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Native *NativeNodePoolInfo `json:"Native,omitnil,omitempty" name:"Native"`

	// 节点 Annotation 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Annotations []*Annotation `json:"Annotations,omitnil,omitempty" name:"Annotations"`

	// 超级节点池参数，在Type等于Super该字段才有值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Super *SuperNodePoolInfo `json:"Super,omitnil,omitempty" name:"Super"`

	// 普通节点池参数，在Type等于Regular该字段才有值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Regular *RegularNodePoolInfo `json:"Regular,omitnil,omitempty" name:"Regular"`

	// 第三方节点池参数，在Type等于External该字段才有值
	// 注意：此字段可能返回 null，表示取不到有效值。
	External *ExternalNodePoolInfo `json:"External,omitnil,omitempty" name:"External"`
}

type RegularNodeInfo struct {
	// 节点配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitnil,omitempty" name:"InstanceAdvancedSettings"`

	// 自动伸缩组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoscalingGroupId *string `json:"AutoscalingGroupId,omitnil,omitempty" name:"AutoscalingGroupId"`
}

type RegularNodePoolInfo struct {
	// LaunchConfigurationId 配置
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// AutoscalingGroupId 分组id
	AutoscalingGroupId *string `json:"AutoscalingGroupId,omitnil,omitempty" name:"AutoscalingGroupId"`

	// NodeCountSummary 节点列表
	NodeCountSummary *NodeCountSummary `json:"NodeCountSummary,omitnil,omitempty" name:"NodeCountSummary"`

	// 状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoscalingGroupStatus *string `json:"AutoscalingGroupStatus,omitnil,omitempty" name:"AutoscalingGroupStatus"`

	// 最大节点数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxNodesNum *int64 `json:"MaxNodesNum,omitnil,omitempty" name:"MaxNodesNum"`

	// 最小节点数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinNodesNum *int64 `json:"MinNodesNum,omitnil,omitempty" name:"MinNodesNum"`

	// 期望的节点数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DesiredNodesNum *int64 `json:"DesiredNodesNum,omitnil,omitempty" name:"DesiredNodesNum"`

	// 节点池osName
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodePoolOs *string `json:"NodePoolOs,omitnil,omitempty" name:"NodePoolOs"`

	// 节点配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitnil,omitempty" name:"InstanceAdvancedSettings"`
}

type RuntimeConfig struct {
	// 运行时类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeType *string `json:"RuntimeType,omitnil,omitempty" name:"RuntimeType"`

	// 运行时版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeVersion *string `json:"RuntimeVersion,omitnil,omitempty" name:"RuntimeVersion"`

	// 运行时根目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeRootDir *string `json:"RuntimeRootDir,omitnil,omitempty" name:"RuntimeRootDir"`
}

type SortBy struct {
	// 排序指标
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// 排序方式
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

type SuperNodeInfo struct {
	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 自动续费标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 节点的 CPU 规格，单位：核。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CPU *float64 `json:"CPU,omitnil,omitempty" name:"CPU"`

	// 节点上 Pod 的 CPU总和，单位：核。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedCPU *float64 `json:"UsedCPU,omitnil,omitempty" name:"UsedCPU"`

	// 节点的内存规格，单位：Gi。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *float64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 节点上 Pod 的内存总和，单位：Gi。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedMemory *float64 `json:"UsedMemory,omitnil,omitempty" name:"UsedMemory"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// VPC 唯一 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网唯一 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 生效时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveAt *string `json:"ActiveAt,omitnil,omitempty" name:"ActiveAt"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireAt *string `json:"ExpireAt,omitnil,omitempty" name:"ExpireAt"`

	// 可调度的单 Pod 最大 CPU 规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxCPUScheduledPod *int64 `json:"MaxCPUScheduledPod,omitnil,omitempty" name:"MaxCPUScheduledPod"`

	// 实例属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceAttribute *string `json:"InstanceAttribute,omitnil,omitempty" name:"InstanceAttribute"`
}

type SuperNodePoolInfo struct {
	// 子网列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 安全组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type Taint struct {
	// Taint的Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Taint的Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Taint的Effect
	Effect *string `json:"Effect,omitnil,omitempty" name:"Effect"`
}