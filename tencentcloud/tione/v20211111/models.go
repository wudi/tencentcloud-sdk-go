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

package v20211111

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type APIConfigDetail struct {
	// 接口id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 接口所属服务组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 接口描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 相对路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelativeUrl *string `json:"RelativeUrl,omitnil,omitempty" name:"RelativeUrl"`

	// 服务类型 HTTP HTTPS
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// GET POST
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpMethod *string `json:"HttpMethod,omitnil,omitempty" name:"HttpMethod"`

	// 请求示例
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpInputExample *string `json:"HttpInputExample,omitnil,omitempty" name:"HttpInputExample"`

	// 回包示例
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpOutputExample *string `json:"HttpOutputExample,omitnil,omitempty" name:"HttpOutputExample"`

	// 更新成员
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedBy *string `json:"UpdatedBy,omitnil,omitempty" name:"UpdatedBy"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 主账号uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 子账号subuin
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`
}

type BatchModelAccTask struct {
	// 模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 模型版本
	ModelVersion *string `json:"ModelVersion,omitnil,omitempty" name:"ModelVersion"`

	// 模型来源(JOB/COS)
	ModelSource *string `json:"ModelSource,omitnil,omitempty" name:"ModelSource"`

	// 模型格式(TORCH_SCRIPT/DETECTRON2/SAVED_MODEL/FROZEN_GRAPH/MMDETECTION/ONNX/HUGGING_FACE)
	ModelFormat *string `json:"ModelFormat,omitnil,omitempty" name:"ModelFormat"`

	// 模型Tensor信息
	TensorInfos []*string `json:"TensorInfos,omitnil,omitempty" name:"TensorInfos"`

	// 加速引擎版本
	AccEngineVersion *string `json:"AccEngineVersion,omitnil,omitempty" name:"AccEngineVersion"`

	// 模型输入cos路径
	ModelInputPath *CosPathInfo `json:"ModelInputPath,omitnil,omitempty" name:"ModelInputPath"`

	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// SavedModel保存时配置的签名
	ModelSignature *string `json:"ModelSignature,omitnil,omitempty" name:"ModelSignature"`

	// 加速引擎对应的框架版本
	FrameworkVersion *string `json:"FrameworkVersion,omitnil,omitempty" name:"FrameworkVersion"`
}

type BatchTaskDetail struct {
	// 批量预测任务ID
	BatchTaskId *string `json:"BatchTaskId,omitnil,omitempty" name:"BatchTaskId"`

	// 批量预测任务名称
	BatchTaskName *string `json:"BatchTaskName,omitnil,omitempty" name:"BatchTaskName"`

	// 主账号uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 子账号uin
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 计费模式
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 包年包月资源组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 包年包月资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// 资源配置
	ResourceConfigInfo *ResourceConfigInfo `json:"ResourceConfigInfo,omitnil,omitempty" name:"ResourceConfigInfo"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 服务对应的模型信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelInfo *ModelInfo `json:"ModelInfo,omitnil,omitempty" name:"ModelInfo"`

	// 自定义镜像信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 代码包
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodePackagePath *CosPathInfo `json:"CodePackagePath,omitnil,omitempty" name:"CodePackagePath"`

	// 启动命令
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartCmd *string `json:"StartCmd,omitnil,omitempty" name:"StartCmd"`

	// 输入数据配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataConfigs []*DataConfig `json:"DataConfigs,omitnil,omitempty" name:"DataConfigs"`

	// 输出数据配置
	Outputs []*DataConfig `json:"Outputs,omitnil,omitempty" name:"Outputs"`

	// 是否上报日志
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// 日志配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// vpc id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 任务状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 运行时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeInSeconds *uint64 `json:"RuntimeInSeconds,omitnil,omitempty" name:"RuntimeInSeconds"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 任务开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 计费状态，取值范围:
	// BILLING: 计费中
	// NOT_BILLING: 未计费
	// WHITELIST_USING: 白名单使用中
	// WHITELIST_STOP: 白名单到期
	// ARREARS_STOP: 欠费停止
	ChargeStatus *string `json:"ChargeStatus,omitnil,omitempty" name:"ChargeStatus"`

	// 最近一次实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestInstanceId *string `json:"LatestInstanceId,omitnil,omitempty" name:"LatestInstanceId"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureReason *string `json:"FailureReason,omitnil,omitempty" name:"FailureReason"`

	// 计费金额信息，eg：2.00元/小时 (for 按量计费)
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingInfo *string `json:"BillingInfo,omitnil,omitempty" name:"BillingInfo"`

	// 运行中的Pod的名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodList []*string `json:"PodList,omitnil,omitempty" name:"PodList"`

	// 模型推理代码信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelInferenceCodeInfo *CosPathInfo `json:"ModelInferenceCodeInfo,omitnil,omitempty" name:"ModelInferenceCodeInfo"`
}

type BatchTaskInstance struct {
	// 任务实例id
	BatchTaskInstanceId *string `json:"BatchTaskInstanceId,omitnil,omitempty" name:"BatchTaskInstanceId"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 运行时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeInSeconds *uint64 `json:"RuntimeInSeconds,omitnil,omitempty" name:"RuntimeInSeconds"`
}

type BatchTaskSetItem struct {
	// 批量预测任务ID
	BatchTaskId *string `json:"BatchTaskId,omitnil,omitempty" name:"BatchTaskId"`

	// 批量预测任务名称
	BatchTaskName *string `json:"BatchTaskName,omitnil,omitempty" name:"BatchTaskName"`

	// 模型信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelInfo *ModelInfo `json:"ModelInfo,omitnil,omitempty" name:"ModelInfo"`

	// 镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 计费模式
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 计费状态，取值范围:
	// BILLING: 计费中
	// NOT_BILLING: 未计费
	// WHITELIST_USING: 白名单使用中
	// WHITELIST_STOP: 白名单到期
	// ARREARS_STOP: 欠费停止
	ChargeStatus *string `json:"ChargeStatus,omitnil,omitempty" name:"ChargeStatus"`

	// 包年包月资源组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 资源配置
	ResourceConfigInfo *ResourceConfigInfo `json:"ResourceConfigInfo,omitnil,omitempty" name:"ResourceConfigInfo"`

	// 标签配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 任务状态, 取值范围:
	// INIT, STARTING, RUNNING, FAILED, STOPPING, STOPPED, SUCCEED
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 运行时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeInSeconds *uint64 `json:"RuntimeInSeconds,omitnil,omitempty" name:"RuntimeInSeconds"`

	// 任务创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务开始运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 输出
	Outputs []*DataConfig `json:"Outputs,omitnil,omitempty" name:"Outputs"`

	// 包年包月资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// 失败原因
	FailureReason *string `json:"FailureReason,omitnil,omitempty" name:"FailureReason"`

	// 按量计费信息
	BillingInfo *string `json:"BillingInfo,omitnil,omitempty" name:"BillingInfo"`
}

type CFSConfig struct {
	// cfs的实例的ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 存储的路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// cfs的挂载类型，可选值为：STORAGE、SOURCE 分别表示存储拓展模式和数据源模式，默认为 STORAGE
	// 注意：此字段可能返回 null，表示取不到有效值。
	MountType *string `json:"MountType,omitnil,omitempty" name:"MountType"`

	// 协议 1: NFS, 2: TURBO
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

type CFSTurbo struct {
	// CFSTurbo实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// CFSTurbo路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

// Predefined struct for user
type ChatCompletionRequestParams struct {
	// 对话的目标模型ID。
	// 自行部署的开源大模型聊天：部署的模型服务组ID，形如ms-xxyyzz。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 输入对话历史。旧的对话在前，数组中最后一项应该为这次的问题。
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 仅当模型为自行部署的开源大模型时生效。采样随机值，默认值为1.0，取值范围[0,2]。较高的值(如0.8)将使输出更加随机，而较低的值(如0.2)将使输出更加确定。建议仅修改此参数或TopP，但不建议两者都修改。
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 仅当模型为自行部署的开源大模型时生效。核采样，默认值为1，取值范围[0,1]。指的是预先设置一个概率界限 p，然后将所有可能生成的token，根据概率大小从高到低排列，依次选取。当这些选取的token的累积概率大于或等于 p 值时停止，然后从已经选取的token中进行采样，生成下一个token。例如top_p为0.1时意味着模型只考虑累积概率为10%的token。建议仅修改此参数或Temperature，不建议两者都修改。
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// 仅当模型为自行部署的开源大模型时生效。最大生成的token数目。默认为无限大。
	MaxTokens *int64 `json:"MaxTokens,omitnil,omitempty" name:"MaxTokens"`
}

type ChatCompletionRequest struct {
	*tchttp.BaseRequest
	
	// 对话的目标模型ID。
	// 自行部署的开源大模型聊天：部署的模型服务组ID，形如ms-xxyyzz。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 输入对话历史。旧的对话在前，数组中最后一项应该为这次的问题。
	Messages []*Message `json:"Messages,omitnil,omitempty" name:"Messages"`

	// 仅当模型为自行部署的开源大模型时生效。采样随机值，默认值为1.0，取值范围[0,2]。较高的值(如0.8)将使输出更加随机，而较低的值(如0.2)将使输出更加确定。建议仅修改此参数或TopP，但不建议两者都修改。
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 仅当模型为自行部署的开源大模型时生效。核采样，默认值为1，取值范围[0,1]。指的是预先设置一个概率界限 p，然后将所有可能生成的token，根据概率大小从高到低排列，依次选取。当这些选取的token的累积概率大于或等于 p 值时停止，然后从已经选取的token中进行采样，生成下一个token。例如top_p为0.1时意味着模型只考虑累积概率为10%的token。建议仅修改此参数或Temperature，不建议两者都修改。
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// 仅当模型为自行部署的开源大模型时生效。最大生成的token数目。默认为无限大。
	MaxTokens *int64 `json:"MaxTokens,omitnil,omitempty" name:"MaxTokens"`
}

func (r *ChatCompletionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatCompletionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Model")
	delete(f, "Messages")
	delete(f, "Temperature")
	delete(f, "TopP")
	delete(f, "MaxTokens")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChatCompletionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChatCompletionResponseParams struct {
	// 部署好的服务Id
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 本次问答的答案。
	Choices []*Choice `json:"Choices,omitnil,omitempty" name:"Choices"`

	// 会话Id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// token统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	Usage *Usage `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChatCompletionResponse struct {
	*tchttp.BaseResponse
	Response *ChatCompletionResponseParams `json:"Response"`
}

func (r *ChatCompletionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatCompletionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Choice struct {
	// 对话结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *Message `json:"Message,omitnil,omitempty" name:"Message"`

	// 结束理由: stop, length, content_filter, null
	FinishReason *string `json:"FinishReason,omitnil,omitempty" name:"FinishReason"`

	// 序号
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`
}

type Container struct {
	// 名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerId *string `json:"ContainerId,omitnil,omitempty" name:"ContainerId"`

	// 镜像地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 容器状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *ContainerStatus `json:"Status,omitnil,omitempty" name:"Status"`
}

type ContainerStatus struct {
	// 重启次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartCount *int64 `json:"RestartCount,omitnil,omitempty" name:"RestartCount"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 是否就绪
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ready *bool `json:"Ready,omitnil,omitempty" name:"Ready"`

	// 状态原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 容器的错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type CosPathInfo struct {
	// 存储桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 路径列表，目前只支持单个
	// 注意：此字段可能返回 null，表示取不到有效值。
	Paths []*string `json:"Paths,omitnil,omitempty" name:"Paths"`
}

// Predefined struct for user
type CreateBatchModelAccTasksRequestParams struct {
	// 模型加速任务名称
	ModelAccTaskName *string `json:"ModelAccTaskName,omitnil,omitempty" name:"ModelAccTaskName"`

	// 批量模型加速任务
	BatchModelAccTasks []*BatchModelAccTask `json:"BatchModelAccTasks,omitnil,omitempty" name:"BatchModelAccTasks"`

	// 模型加速保存路径
	ModelOutputPath *CosPathInfo `json:"ModelOutputPath,omitnil,omitempty" name:"ModelOutputPath"`

	// 标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 优化级别(NO_LOSS/FP16/INT8)，默认FP16
	OptimizationLevel *string `json:"OptimizationLevel,omitnil,omitempty" name:"OptimizationLevel"`

	// GPU卡类型(T4/V100/A10)，默认T4
	GPUType *string `json:"GPUType,omitnil,omitempty" name:"GPUType"`

	// 专业参数设置
	HyperParameter *HyperParameter `json:"HyperParameter,omitnil,omitempty" name:"HyperParameter"`
}

type CreateBatchModelAccTasksRequest struct {
	*tchttp.BaseRequest
	
	// 模型加速任务名称
	ModelAccTaskName *string `json:"ModelAccTaskName,omitnil,omitempty" name:"ModelAccTaskName"`

	// 批量模型加速任务
	BatchModelAccTasks []*BatchModelAccTask `json:"BatchModelAccTasks,omitnil,omitempty" name:"BatchModelAccTasks"`

	// 模型加速保存路径
	ModelOutputPath *CosPathInfo `json:"ModelOutputPath,omitnil,omitempty" name:"ModelOutputPath"`

	// 标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 优化级别(NO_LOSS/FP16/INT8)，默认FP16
	OptimizationLevel *string `json:"OptimizationLevel,omitnil,omitempty" name:"OptimizationLevel"`

	// GPU卡类型(T4/V100/A10)，默认T4
	GPUType *string `json:"GPUType,omitnil,omitempty" name:"GPUType"`

	// 专业参数设置
	HyperParameter *HyperParameter `json:"HyperParameter,omitnil,omitempty" name:"HyperParameter"`
}

func (r *CreateBatchModelAccTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchModelAccTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelAccTaskName")
	delete(f, "BatchModelAccTasks")
	delete(f, "ModelOutputPath")
	delete(f, "Tags")
	delete(f, "OptimizationLevel")
	delete(f, "GPUType")
	delete(f, "HyperParameter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBatchModelAccTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchModelAccTasksResponseParams struct {
	// 模型优化任务ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAccTaskIds []*string `json:"ModelAccTaskIds,omitnil,omitempty" name:"ModelAccTaskIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBatchModelAccTasksResponse struct {
	*tchttp.BaseResponse
	Response *CreateBatchModelAccTasksResponseParams `json:"Response"`
}

func (r *CreateBatchModelAccTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchModelAccTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchTaskRequestParams struct {
	// 批量预测任务名称，不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	BatchTaskName *string `json:"BatchTaskName,omitnil,omitempty" name:"BatchTaskName"`

	// 计费模式，eg：PREPAID 包年包月；POSTPAID_BY_HOUR 按量计费
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 资源配置
	ResourceConfigInfo *ResourceConfigInfo `json:"ResourceConfigInfo,omitnil,omitempty" name:"ResourceConfigInfo"`

	// 结果输出存储信息
	Outputs []*DataConfig `json:"Outputs,omitnil,omitempty" name:"Outputs"`

	// 是否上报日志
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// 工作类型 1:单次 2:周期
	//
	// Deprecated: JobType is deprecated.
	JobType *uint64 `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 任务周期描述
	//
	// Deprecated: CronInfo is deprecated.
	CronInfo *CronInfo `json:"CronInfo,omitnil,omitempty" name:"CronInfo"`

	// 包年包月资源组ID
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 标签配置
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 服务对应的模型信息，有模型文件时需要填写
	ModelInfo *ModelInfo `json:"ModelInfo,omitnil,omitempty" name:"ModelInfo"`

	// 自定义镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 代码包
	CodePackage *CosPathInfo `json:"CodePackage,omitnil,omitempty" name:"CodePackage"`

	// 启动命令
	StartCmd *string `json:"StartCmd,omitnil,omitempty" name:"StartCmd"`

	// 数据配置
	DataConfigs []*DataConfig `json:"DataConfigs,omitnil,omitempty" name:"DataConfigs"`

	// 日志配置
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// VPC Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网Id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 任务执行结果回调URL，仅支持http和https。回调格式&内容详见: [TI-ONE 接口回调说明](https://cloud.tencent.com/document/product/851/84292)
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`
}

type CreateBatchTaskRequest struct {
	*tchttp.BaseRequest
	
	// 批量预测任务名称，不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	BatchTaskName *string `json:"BatchTaskName,omitnil,omitempty" name:"BatchTaskName"`

	// 计费模式，eg：PREPAID 包年包月；POSTPAID_BY_HOUR 按量计费
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 资源配置
	ResourceConfigInfo *ResourceConfigInfo `json:"ResourceConfigInfo,omitnil,omitempty" name:"ResourceConfigInfo"`

	// 结果输出存储信息
	Outputs []*DataConfig `json:"Outputs,omitnil,omitempty" name:"Outputs"`

	// 是否上报日志
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// 工作类型 1:单次 2:周期
	JobType *uint64 `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 任务周期描述
	CronInfo *CronInfo `json:"CronInfo,omitnil,omitempty" name:"CronInfo"`

	// 包年包月资源组ID
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 标签配置
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 服务对应的模型信息，有模型文件时需要填写
	ModelInfo *ModelInfo `json:"ModelInfo,omitnil,omitempty" name:"ModelInfo"`

	// 自定义镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 代码包
	CodePackage *CosPathInfo `json:"CodePackage,omitnil,omitempty" name:"CodePackage"`

	// 启动命令
	StartCmd *string `json:"StartCmd,omitnil,omitempty" name:"StartCmd"`

	// 数据配置
	DataConfigs []*DataConfig `json:"DataConfigs,omitnil,omitempty" name:"DataConfigs"`

	// 日志配置
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// VPC Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网Id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 任务执行结果回调URL，仅支持http和https。回调格式&内容详见: [TI-ONE 接口回调说明](https://cloud.tencent.com/document/product/851/84292)
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`
}

func (r *CreateBatchTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchTaskName")
	delete(f, "ChargeType")
	delete(f, "ResourceConfigInfo")
	delete(f, "Outputs")
	delete(f, "LogEnable")
	delete(f, "JobType")
	delete(f, "CronInfo")
	delete(f, "ResourceGroupId")
	delete(f, "Tags")
	delete(f, "ModelInfo")
	delete(f, "ImageInfo")
	delete(f, "CodePackage")
	delete(f, "StartCmd")
	delete(f, "DataConfigs")
	delete(f, "LogConfig")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Remark")
	delete(f, "CallbackUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBatchTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchTaskResponseParams struct {
	// 批量预测任务ID
	BatchTaskId *string `json:"BatchTaskId,omitnil,omitempty" name:"BatchTaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBatchTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateBatchTaskResponseParams `json:"Response"`
}

func (r *CreateBatchTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatasetRequestParams struct {
	// 数据集名称，不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	DatasetName *string `json:"DatasetName,omitnil,omitempty" name:"DatasetName"`

	// 数据集类型:
	// TYPE_DATASET_TEXT，文本
	// TYPE_DATASET_IMAGE，图片
	// TYPE_DATASET_TABLE，表格
	// TYPE_DATASET_OTHER，其他
	DatasetType *string `json:"DatasetType,omitnil,omitempty" name:"DatasetType"`

	// 数据源cos路径
	StorageDataPath *CosPathInfo `json:"StorageDataPath,omitnil,omitempty" name:"StorageDataPath"`

	// 数据集标签cos存储路径
	StorageLabelPath *CosPathInfo `json:"StorageLabelPath,omitnil,omitempty" name:"StorageLabelPath"`

	// 数据集标签
	DatasetTags []*Tag `json:"DatasetTags,omitnil,omitempty" name:"DatasetTags"`

	// 数据集标注状态:
	// STATUS_NON_ANNOTATED，未标注
	// STATUS_ANNOTATED，已标注
	AnnotationStatus *string `json:"AnnotationStatus,omitnil,omitempty" name:"AnnotationStatus"`

	// 标注类型:
	// ANNOTATION_TYPE_CLASSIFICATION，图片分类
	// ANNOTATION_TYPE_DETECTION，目标检测
	// ANNOTATION_TYPE_SEGMENTATION，图片分割
	// ANNOTATION_TYPE_TRACKING，目标跟踪
	AnnotationType *string `json:"AnnotationType,omitnil,omitempty" name:"AnnotationType"`

	// 标注格式:
	// ANNOTATION_FORMAT_TI，TI平台格式
	// ANNOTATION_FORMAT_PASCAL，Pascal Voc
	// ANNOTATION_FORMAT_COCO，COCO
	// ANNOTATION_FORMAT_FILE，文件目录结构
	AnnotationFormat *string `json:"AnnotationFormat,omitnil,omitempty" name:"AnnotationFormat"`

	// 表头信息
	SchemaInfos []*SchemaInfo `json:"SchemaInfos,omitnil,omitempty" name:"SchemaInfos"`

	// 数据是否存在表头
	IsSchemaExisted *bool `json:"IsSchemaExisted,omitnil,omitempty" name:"IsSchemaExisted"`

	// 导入文件粒度，按行或者按文件
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`
}

type CreateDatasetRequest struct {
	*tchttp.BaseRequest
	
	// 数据集名称，不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	DatasetName *string `json:"DatasetName,omitnil,omitempty" name:"DatasetName"`

	// 数据集类型:
	// TYPE_DATASET_TEXT，文本
	// TYPE_DATASET_IMAGE，图片
	// TYPE_DATASET_TABLE，表格
	// TYPE_DATASET_OTHER，其他
	DatasetType *string `json:"DatasetType,omitnil,omitempty" name:"DatasetType"`

	// 数据源cos路径
	StorageDataPath *CosPathInfo `json:"StorageDataPath,omitnil,omitempty" name:"StorageDataPath"`

	// 数据集标签cos存储路径
	StorageLabelPath *CosPathInfo `json:"StorageLabelPath,omitnil,omitempty" name:"StorageLabelPath"`

	// 数据集标签
	DatasetTags []*Tag `json:"DatasetTags,omitnil,omitempty" name:"DatasetTags"`

	// 数据集标注状态:
	// STATUS_NON_ANNOTATED，未标注
	// STATUS_ANNOTATED，已标注
	AnnotationStatus *string `json:"AnnotationStatus,omitnil,omitempty" name:"AnnotationStatus"`

	// 标注类型:
	// ANNOTATION_TYPE_CLASSIFICATION，图片分类
	// ANNOTATION_TYPE_DETECTION，目标检测
	// ANNOTATION_TYPE_SEGMENTATION，图片分割
	// ANNOTATION_TYPE_TRACKING，目标跟踪
	AnnotationType *string `json:"AnnotationType,omitnil,omitempty" name:"AnnotationType"`

	// 标注格式:
	// ANNOTATION_FORMAT_TI，TI平台格式
	// ANNOTATION_FORMAT_PASCAL，Pascal Voc
	// ANNOTATION_FORMAT_COCO，COCO
	// ANNOTATION_FORMAT_FILE，文件目录结构
	AnnotationFormat *string `json:"AnnotationFormat,omitnil,omitempty" name:"AnnotationFormat"`

	// 表头信息
	SchemaInfos []*SchemaInfo `json:"SchemaInfos,omitnil,omitempty" name:"SchemaInfos"`

	// 数据是否存在表头
	IsSchemaExisted *bool `json:"IsSchemaExisted,omitnil,omitempty" name:"IsSchemaExisted"`

	// 导入文件粒度，按行或者按文件
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`
}

func (r *CreateDatasetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatasetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasetName")
	delete(f, "DatasetType")
	delete(f, "StorageDataPath")
	delete(f, "StorageLabelPath")
	delete(f, "DatasetTags")
	delete(f, "AnnotationStatus")
	delete(f, "AnnotationType")
	delete(f, "AnnotationFormat")
	delete(f, "SchemaInfos")
	delete(f, "IsSchemaExisted")
	delete(f, "ContentType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDatasetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatasetResponseParams struct {
	// 数据集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetId *string `json:"DatasetId,omitnil,omitempty" name:"DatasetId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDatasetResponse struct {
	*tchttp.BaseResponse
	Response *CreateDatasetResponseParams `json:"Response"`
}

func (r *CreateDatasetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatasetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModelServiceRequestParams struct {
	// 新增版本时需要填写
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 不超过60个字，仅支持英文、数字、下划线"_"、短横"-"，只能以英文、数字开头
	ServiceGroupName *string `json:"ServiceGroupName,omitnil,omitempty" name:"ServiceGroupName"`

	// 模型服务的描述
	ServiceDescription *string `json:"ServiceDescription,omitnil,omitempty" name:"ServiceDescription"`

	// 付费模式,有 PREPAID （包年包月）和 POSTPAID_BY_HOUR（按量付费）
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 预付费模式下所属的资源组id，同服务组下唯一
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 模型信息，需要挂载模型时填写
	ModelInfo *ModelInfo `json:"ModelInfo,omitnil,omitempty" name:"ModelInfo"`

	// 镜像信息，配置服务运行所需的镜像地址等信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 环境变量，可选参数，用于配置容器中的环境变量
	Env []*EnvVar `json:"Env,omitnil,omitempty" name:"Env"`

	// 资源描述，指定包年包月模式下的cpu,mem,gpu等信息，后付费无需填写
	Resources *ResourceInfo `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 使用DescribeBillingSpecs接口返回的规格列表中的值，或者参考实例列表:
	// TI.S.MEDIUM.POST	2C4G
	// TI.S.LARGE.POST	4C8G
	// TI.S.2XLARGE16.POST	8C16G
	// TI.S.2XLARGE32.POST	8C32G
	// TI.S.4XLARGE32.POST	16C32G
	// TI.S.4XLARGE64.POST	16C64G
	// TI.S.6XLARGE48.POST	24C48G
	// TI.S.6XLARGE96.POST	24C96G
	// TI.S.8XLARGE64.POST	32C64G
	// TI.S.8XLARGE128.POST 32C128G
	// TI.GN7.LARGE20.POST	4C20G T4*1/4
	// TI.GN7.2XLARGE40.POST	10C40G T4*1/2
	// TI.GN7.2XLARGE32.POST	8C32G T4*1
	// TI.GN7.5XLARGE80.POST	20C80G T4*1
	// TI.GN7.8XLARGE128.POST	32C128G T4*1
	// TI.GN7.10XLARGE160.POST	40C160G T4*2
	// TI.GN7.20XLARGE320.POST	80C320G T4*4
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 扩缩容类型 支持：自动 - "AUTO", 手动 - "MANUAL",默认为MANUAL
	ScaleMode *string `json:"ScaleMode,omitnil,omitempty" name:"ScaleMode"`

	// 实例数量, 不同计费模式和调节模式下对应关系如下
	// PREPAID 和 POSTPAID_BY_HOUR:
	// 手动调节模式下对应 实例数量
	// 自动调节模式下对应 基于时间的默认策略的实例数量
	// HYBRID_PAID:
	// 后付费实例手动调节模式下对应 实例数量
	// 后付费实例自动调节模式下对应 时间策略的默认策略的实例数量
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// 自动伸缩信息
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil,omitempty" name:"HorizontalPodAutoscaler"`

	// 是否开启日志投递，开启后需填写配置投递到指定cls
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// 日志配置，需要投递服务日志到指定cls时填写
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// 是否开启接口鉴权，开启后自动生成token信息，访问需要token鉴权
	AuthorizationEnable *bool `json:"AuthorizationEnable,omitnil,omitempty" name:"AuthorizationEnable"`

	// 腾讯云标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否新增版本
	NewVersion *bool `json:"NewVersion,omitnil,omitempty" name:"NewVersion"`

	// 定时任务配置，使用定时策略时填写
	CronScaleJobs []*CronScaleJob `json:"CronScaleJobs,omitnil,omitempty" name:"CronScaleJobs"`

	// 自动伸缩策略配置 HPA : 通过HPA进行弹性伸缩 CRON 通过定时任务进行伸缩
	ScaleStrategy *string `json:"ScaleStrategy,omitnil,omitempty" name:"ScaleStrategy"`

	// 计费模式[HYBRID_PAID]时生效, 用于标识混合计费模式下的预付费实例数
	HybridBillingPrepaidReplicas *int64 `json:"HybridBillingPrepaidReplicas,omitnil,omitempty" name:"HybridBillingPrepaidReplicas"`

	// [AUTO_ML 自动学习，自动学习正式发布 AUTO_ML_FORMAL, DEFAULT 默认]
	CreateSource *string `json:"CreateSource,omitnil,omitempty" name:"CreateSource"`

	// 是否开启模型的热更新。默认不开启
	ModelHotUpdateEnable *bool `json:"ModelHotUpdateEnable,omitnil,omitempty" name:"ModelHotUpdateEnable"`

	// 定时停止配置
	ScheduledAction *ScheduledAction `json:"ScheduledAction,omitnil,omitempty" name:"ScheduledAction"`

	// 挂载配置，目前只支持CFS
	VolumeMount *VolumeMount `json:"VolumeMount,omitnil,omitempty" name:"VolumeMount"`

	// 服务限速限流相关配置
	ServiceLimit *ServiceLimit `json:"ServiceLimit,omitnil,omitempty" name:"ServiceLimit"`

	// 回调地址，用于回调创建服务状态信息，回调格式&内容详情见：[TI-ONE 接口回调说明](https://cloud.tencent.com/document/product/851/84292)
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 是否开启模型的加速, 仅对StableDiffusion(动态加速)格式的模型有效。
	ModelTurboEnable *bool `json:"ModelTurboEnable,omitnil,omitempty" name:"ModelTurboEnable"`

	// 服务分类
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`

	// 服务的启动命令
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 是否开启TIONE内网访问外部
	ServiceEIP *ServiceEIP `json:"ServiceEIP,omitnil,omitempty" name:"ServiceEIP"`
}

type CreateModelServiceRequest struct {
	*tchttp.BaseRequest
	
	// 新增版本时需要填写
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 不超过60个字，仅支持英文、数字、下划线"_"、短横"-"，只能以英文、数字开头
	ServiceGroupName *string `json:"ServiceGroupName,omitnil,omitempty" name:"ServiceGroupName"`

	// 模型服务的描述
	ServiceDescription *string `json:"ServiceDescription,omitnil,omitempty" name:"ServiceDescription"`

	// 付费模式,有 PREPAID （包年包月）和 POSTPAID_BY_HOUR（按量付费）
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 预付费模式下所属的资源组id，同服务组下唯一
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 模型信息，需要挂载模型时填写
	ModelInfo *ModelInfo `json:"ModelInfo,omitnil,omitempty" name:"ModelInfo"`

	// 镜像信息，配置服务运行所需的镜像地址等信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 环境变量，可选参数，用于配置容器中的环境变量
	Env []*EnvVar `json:"Env,omitnil,omitempty" name:"Env"`

	// 资源描述，指定包年包月模式下的cpu,mem,gpu等信息，后付费无需填写
	Resources *ResourceInfo `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 使用DescribeBillingSpecs接口返回的规格列表中的值，或者参考实例列表:
	// TI.S.MEDIUM.POST	2C4G
	// TI.S.LARGE.POST	4C8G
	// TI.S.2XLARGE16.POST	8C16G
	// TI.S.2XLARGE32.POST	8C32G
	// TI.S.4XLARGE32.POST	16C32G
	// TI.S.4XLARGE64.POST	16C64G
	// TI.S.6XLARGE48.POST	24C48G
	// TI.S.6XLARGE96.POST	24C96G
	// TI.S.8XLARGE64.POST	32C64G
	// TI.S.8XLARGE128.POST 32C128G
	// TI.GN7.LARGE20.POST	4C20G T4*1/4
	// TI.GN7.2XLARGE40.POST	10C40G T4*1/2
	// TI.GN7.2XLARGE32.POST	8C32G T4*1
	// TI.GN7.5XLARGE80.POST	20C80G T4*1
	// TI.GN7.8XLARGE128.POST	32C128G T4*1
	// TI.GN7.10XLARGE160.POST	40C160G T4*2
	// TI.GN7.20XLARGE320.POST	80C320G T4*4
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 扩缩容类型 支持：自动 - "AUTO", 手动 - "MANUAL",默认为MANUAL
	ScaleMode *string `json:"ScaleMode,omitnil,omitempty" name:"ScaleMode"`

	// 实例数量, 不同计费模式和调节模式下对应关系如下
	// PREPAID 和 POSTPAID_BY_HOUR:
	// 手动调节模式下对应 实例数量
	// 自动调节模式下对应 基于时间的默认策略的实例数量
	// HYBRID_PAID:
	// 后付费实例手动调节模式下对应 实例数量
	// 后付费实例自动调节模式下对应 时间策略的默认策略的实例数量
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// 自动伸缩信息
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil,omitempty" name:"HorizontalPodAutoscaler"`

	// 是否开启日志投递，开启后需填写配置投递到指定cls
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// 日志配置，需要投递服务日志到指定cls时填写
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// 是否开启接口鉴权，开启后自动生成token信息，访问需要token鉴权
	AuthorizationEnable *bool `json:"AuthorizationEnable,omitnil,omitempty" name:"AuthorizationEnable"`

	// 腾讯云标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否新增版本
	NewVersion *bool `json:"NewVersion,omitnil,omitempty" name:"NewVersion"`

	// 定时任务配置，使用定时策略时填写
	CronScaleJobs []*CronScaleJob `json:"CronScaleJobs,omitnil,omitempty" name:"CronScaleJobs"`

	// 自动伸缩策略配置 HPA : 通过HPA进行弹性伸缩 CRON 通过定时任务进行伸缩
	ScaleStrategy *string `json:"ScaleStrategy,omitnil,omitempty" name:"ScaleStrategy"`

	// 计费模式[HYBRID_PAID]时生效, 用于标识混合计费模式下的预付费实例数
	HybridBillingPrepaidReplicas *int64 `json:"HybridBillingPrepaidReplicas,omitnil,omitempty" name:"HybridBillingPrepaidReplicas"`

	// [AUTO_ML 自动学习，自动学习正式发布 AUTO_ML_FORMAL, DEFAULT 默认]
	CreateSource *string `json:"CreateSource,omitnil,omitempty" name:"CreateSource"`

	// 是否开启模型的热更新。默认不开启
	ModelHotUpdateEnable *bool `json:"ModelHotUpdateEnable,omitnil,omitempty" name:"ModelHotUpdateEnable"`

	// 定时停止配置
	ScheduledAction *ScheduledAction `json:"ScheduledAction,omitnil,omitempty" name:"ScheduledAction"`

	// 挂载配置，目前只支持CFS
	VolumeMount *VolumeMount `json:"VolumeMount,omitnil,omitempty" name:"VolumeMount"`

	// 服务限速限流相关配置
	ServiceLimit *ServiceLimit `json:"ServiceLimit,omitnil,omitempty" name:"ServiceLimit"`

	// 回调地址，用于回调创建服务状态信息，回调格式&内容详情见：[TI-ONE 接口回调说明](https://cloud.tencent.com/document/product/851/84292)
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 是否开启模型的加速, 仅对StableDiffusion(动态加速)格式的模型有效。
	ModelTurboEnable *bool `json:"ModelTurboEnable,omitnil,omitempty" name:"ModelTurboEnable"`

	// 服务分类
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`

	// 服务的启动命令
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 是否开启TIONE内网访问外部
	ServiceEIP *ServiceEIP `json:"ServiceEIP,omitnil,omitempty" name:"ServiceEIP"`
}

func (r *CreateModelServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceGroupId")
	delete(f, "ServiceGroupName")
	delete(f, "ServiceDescription")
	delete(f, "ChargeType")
	delete(f, "ResourceGroupId")
	delete(f, "ModelInfo")
	delete(f, "ImageInfo")
	delete(f, "Env")
	delete(f, "Resources")
	delete(f, "InstanceType")
	delete(f, "ScaleMode")
	delete(f, "Replicas")
	delete(f, "HorizontalPodAutoscaler")
	delete(f, "LogEnable")
	delete(f, "LogConfig")
	delete(f, "AuthorizationEnable")
	delete(f, "Tags")
	delete(f, "NewVersion")
	delete(f, "CronScaleJobs")
	delete(f, "ScaleStrategy")
	delete(f, "HybridBillingPrepaidReplicas")
	delete(f, "CreateSource")
	delete(f, "ModelHotUpdateEnable")
	delete(f, "ScheduledAction")
	delete(f, "VolumeMount")
	delete(f, "ServiceLimit")
	delete(f, "CallbackUrl")
	delete(f, "ModelTurboEnable")
	delete(f, "ServiceCategory")
	delete(f, "Command")
	delete(f, "ServiceEIP")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateModelServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModelServiceResponseParams struct {
	// 生成的模型服务
	// 注意：此字段可能返回 null，表示取不到有效值。
	Service *Service `json:"Service,omitnil,omitempty" name:"Service"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateModelServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateModelServiceResponseParams `json:"Response"`
}

func (r *CreateModelServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookImageRequestParams struct {
	// 镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// notebook id
	NotebookId *string `json:"NotebookId,omitnil,omitempty" name:"NotebookId"`

	// 要保存的kernel数组
	Kernels []*string `json:"Kernels,omitnil,omitempty" name:"Kernels"`
}

type CreateNotebookImageRequest struct {
	*tchttp.BaseRequest
	
	// 镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// notebook id
	NotebookId *string `json:"NotebookId,omitnil,omitempty" name:"NotebookId"`

	// 要保存的kernel数组
	Kernels []*string `json:"Kernels,omitnil,omitempty" name:"Kernels"`
}

func (r *CreateNotebookImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageInfo")
	delete(f, "NotebookId")
	delete(f, "Kernels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNotebookImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookImageResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNotebookImageResponse struct {
	*tchttp.BaseResponse
	Response *CreateNotebookImageResponseParams `json:"Response"`
}

func (r *CreateNotebookImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookRequestParams struct {
	// 名称。不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 计算资源付费模式 ，可选值为：
	// PREPAID：预付费，即包年包月
	// POSTPAID_BY_HOUR：按小时后付费
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 计算资源配置
	ResourceConf *ResourceConf `json:"ResourceConf,omitnil,omitempty" name:"ResourceConf"`

	// 是否上报日志
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// 是否ROOT权限
	RootAccess *bool `json:"RootAccess,omitnil,omitempty" name:"RootAccess"`

	// 是否自动停止
	AutoStopping *bool `json:"AutoStopping,omitnil,omitempty" name:"AutoStopping"`

	// 是否访问公网
	DirectInternetAccess *bool `json:"DirectInternetAccess,omitnil,omitempty" name:"DirectInternetAccess"`

	// 资源组ID(for预付费)
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// Vpc-Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网Id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 存储的类型。取值包含： 
	// FREE：预付费的免费存储
	// CLOUD_PREMIUM：高性能云硬盘
	// CLOUD_SSD：SSD云硬盘
	// CFS：CFS存储
	// CFS_TURBO：CFS Turbo存储
	// GooseFSx：GooseFSx存储
	VolumeSourceType *string `json:"VolumeSourceType,omitnil,omitempty" name:"VolumeSourceType"`

	// 云硬盘存储卷大小，单位GB
	VolumeSizeInGB *uint64 `json:"VolumeSizeInGB,omitnil,omitempty" name:"VolumeSizeInGB"`

	// CFS存储的配置
	VolumeSourceCFS *CFSConfig `json:"VolumeSourceCFS,omitnil,omitempty" name:"VolumeSourceCFS"`

	// 日志配置
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// 生命周期脚本的ID
	LifecycleScriptId *string `json:"LifecycleScriptId,omitnil,omitempty" name:"LifecycleScriptId"`

	// 默认GIT存储库的ID
	DefaultCodeRepoId *string `json:"DefaultCodeRepoId,omitnil,omitempty" name:"DefaultCodeRepoId"`

	// 其他GIT存储库的ID，最多3个
	AdditionalCodeRepoIds []*string `json:"AdditionalCodeRepoIds,omitnil,omitempty" name:"AdditionalCodeRepoIds"`

	// 自动停止时间，单位小时
	AutomaticStopTime *uint64 `json:"AutomaticStopTime,omitnil,omitempty" name:"AutomaticStopTime"`

	// 标签配置
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 数据配置，只支持WEDATA_HDFS存储类型
	DataConfigs []*DataConfig `json:"DataConfigs,omitnil,omitempty" name:"DataConfigs"`

	// 镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 镜像类型，包括SYSTEM、TCR、CCR
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// SSH配置信息
	SSHConfig *SSHConfig `json:"SSHConfig,omitnil,omitempty" name:"SSHConfig"`

	// GooseFS存储配置
	VolumeSourceGooseFS *GooseFS `json:"VolumeSourceGooseFS,omitnil,omitempty" name:"VolumeSourceGooseFS"`
}

type CreateNotebookRequest struct {
	*tchttp.BaseRequest
	
	// 名称。不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 计算资源付费模式 ，可选值为：
	// PREPAID：预付费，即包年包月
	// POSTPAID_BY_HOUR：按小时后付费
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 计算资源配置
	ResourceConf *ResourceConf `json:"ResourceConf,omitnil,omitempty" name:"ResourceConf"`

	// 是否上报日志
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// 是否ROOT权限
	RootAccess *bool `json:"RootAccess,omitnil,omitempty" name:"RootAccess"`

	// 是否自动停止
	AutoStopping *bool `json:"AutoStopping,omitnil,omitempty" name:"AutoStopping"`

	// 是否访问公网
	DirectInternetAccess *bool `json:"DirectInternetAccess,omitnil,omitempty" name:"DirectInternetAccess"`

	// 资源组ID(for预付费)
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// Vpc-Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网Id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 存储的类型。取值包含： 
	// FREE：预付费的免费存储
	// CLOUD_PREMIUM：高性能云硬盘
	// CLOUD_SSD：SSD云硬盘
	// CFS：CFS存储
	// CFS_TURBO：CFS Turbo存储
	// GooseFSx：GooseFSx存储
	VolumeSourceType *string `json:"VolumeSourceType,omitnil,omitempty" name:"VolumeSourceType"`

	// 云硬盘存储卷大小，单位GB
	VolumeSizeInGB *uint64 `json:"VolumeSizeInGB,omitnil,omitempty" name:"VolumeSizeInGB"`

	// CFS存储的配置
	VolumeSourceCFS *CFSConfig `json:"VolumeSourceCFS,omitnil,omitempty" name:"VolumeSourceCFS"`

	// 日志配置
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// 生命周期脚本的ID
	LifecycleScriptId *string `json:"LifecycleScriptId,omitnil,omitempty" name:"LifecycleScriptId"`

	// 默认GIT存储库的ID
	DefaultCodeRepoId *string `json:"DefaultCodeRepoId,omitnil,omitempty" name:"DefaultCodeRepoId"`

	// 其他GIT存储库的ID，最多3个
	AdditionalCodeRepoIds []*string `json:"AdditionalCodeRepoIds,omitnil,omitempty" name:"AdditionalCodeRepoIds"`

	// 自动停止时间，单位小时
	AutomaticStopTime *uint64 `json:"AutomaticStopTime,omitnil,omitempty" name:"AutomaticStopTime"`

	// 标签配置
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 数据配置，只支持WEDATA_HDFS存储类型
	DataConfigs []*DataConfig `json:"DataConfigs,omitnil,omitempty" name:"DataConfigs"`

	// 镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 镜像类型，包括SYSTEM、TCR、CCR
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// SSH配置信息
	SSHConfig *SSHConfig `json:"SSHConfig,omitnil,omitempty" name:"SSHConfig"`

	// GooseFS存储配置
	VolumeSourceGooseFS *GooseFS `json:"VolumeSourceGooseFS,omitnil,omitempty" name:"VolumeSourceGooseFS"`
}

func (r *CreateNotebookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ChargeType")
	delete(f, "ResourceConf")
	delete(f, "LogEnable")
	delete(f, "RootAccess")
	delete(f, "AutoStopping")
	delete(f, "DirectInternetAccess")
	delete(f, "ResourceGroupId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "VolumeSourceType")
	delete(f, "VolumeSizeInGB")
	delete(f, "VolumeSourceCFS")
	delete(f, "LogConfig")
	delete(f, "LifecycleScriptId")
	delete(f, "DefaultCodeRepoId")
	delete(f, "AdditionalCodeRepoIds")
	delete(f, "AutomaticStopTime")
	delete(f, "Tags")
	delete(f, "DataConfigs")
	delete(f, "ImageInfo")
	delete(f, "ImageType")
	delete(f, "SSHConfig")
	delete(f, "VolumeSourceGooseFS")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNotebookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookResponseParams struct {
	// notebook标志
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNotebookResponse struct {
	*tchttp.BaseResponse
	Response *CreateNotebookResponseParams `json:"Response"`
}

func (r *CreateNotebookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOptimizedModelRequestParams struct {
	// 模型加速任务ID
	ModelAccTaskId *string `json:"ModelAccTaskId,omitnil,omitempty" name:"ModelAccTaskId"`

	// 标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateOptimizedModelRequest struct {
	*tchttp.BaseRequest
	
	// 模型加速任务ID
	ModelAccTaskId *string `json:"ModelAccTaskId,omitnil,omitempty" name:"ModelAccTaskId"`

	// 标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateOptimizedModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOptimizedModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelAccTaskId")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOptimizedModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOptimizedModelResponseParams struct {
	// 模型ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 模型版本ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelVersionId *string `json:"ModelVersionId,omitnil,omitempty" name:"ModelVersionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOptimizedModelResponse struct {
	*tchttp.BaseResponse
	Response *CreateOptimizedModelResponseParams `json:"Response"`
}

func (r *CreateOptimizedModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOptimizedModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrainingModelRequestParams struct {
	// 导入方式
	// MODEL：导入新模型
	// VERSION：导入新版本
	// EXIST：导入现有版本
	ImportMethod *string `json:"ImportMethod,omitnil,omitempty" name:"ImportMethod"`

	// 推理环境来源（SYSTEM/CUSTOM）
	ReasoningEnvironmentSource *string `json:"ReasoningEnvironmentSource,omitnil,omitempty" name:"ReasoningEnvironmentSource"`

	// 模型名称，不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	TrainingModelName *string `json:"TrainingModelName,omitnil,omitempty" name:"TrainingModelName"`

	// 标签配置
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 训练任务名称
	TrainingJobName *string `json:"TrainingJobName,omitnil,omitempty" name:"TrainingJobName"`

	// 模型来源cos目录，以/结尾
	TrainingModelCosPath *CosPathInfo `json:"TrainingModelCosPath,omitnil,omitempty" name:"TrainingModelCosPath"`

	// 算法框架 （PYTORCH/TENSORFLOW/DETECTRON2/PMML/MMDETECTION)
	AlgorithmFramework *string `json:"AlgorithmFramework,omitnil,omitempty" name:"AlgorithmFramework"`

	// 推理环境
	ReasoningEnvironment *string `json:"ReasoningEnvironment,omitnil,omitempty" name:"ReasoningEnvironment"`

	// 训练指标，最多支持1000字符
	TrainingModelIndex *string `json:"TrainingModelIndex,omitnil,omitempty" name:"TrainingModelIndex"`

	// 模型版本
	TrainingModelVersion *string `json:"TrainingModelVersion,omitnil,omitempty" name:"TrainingModelVersion"`

	// 自定义推理环境
	ReasoningImageInfo *ImageInfo `json:"ReasoningImageInfo,omitnil,omitempty" name:"ReasoningImageInfo"`

	// 模型移动方式（CUT/COPY）
	ModelMoveMode *string `json:"ModelMoveMode,omitnil,omitempty" name:"ModelMoveMode"`

	// 训练任务ID
	TrainingJobId *string `json:"TrainingJobId,omitnil,omitempty" name:"TrainingJobId"`

	// 模型ID（导入新模型不需要，导入新版本需要）
	TrainingModelId *string `json:"TrainingModelId,omitnil,omitempty" name:"TrainingModelId"`

	// 模型存储cos目录
	ModelOutputPath *CosPathInfo `json:"ModelOutputPath,omitnil,omitempty" name:"ModelOutputPath"`

	// 模型来源 （JOB/COS）
	TrainingModelSource *string `json:"TrainingModelSource,omitnil,omitempty" name:"TrainingModelSource"`

	// 模型偏好
	TrainingPreference *string `json:"TrainingPreference,omitnil,omitempty" name:"TrainingPreference"`

	// 自动学习任务ID（已废弃）
	AutoMLTaskId *string `json:"AutoMLTaskId,omitnil,omitempty" name:"AutoMLTaskId"`

	// 任务版本
	TrainingJobVersion *string `json:"TrainingJobVersion,omitnil,omitempty" name:"TrainingJobVersion"`

	// 模型版本类型；
	// 枚举值：NORMAL(通用)  ACCELERATE(加速)
	// 注意:  默认为NORMAL
	ModelVersionType *string `json:"ModelVersionType,omitnil,omitempty" name:"ModelVersionType"`

	// 模型格式 （PYTORCH/TORCH_SCRIPT/DETECTRON2/SAVED_MODEL/FROZEN_GRAPH/PMML/MMDETECTION/ONNX/HUGGING_FACE）
	ModelFormat *string `json:"ModelFormat,omitnil,omitempty" name:"ModelFormat"`

	// 推理镜像ID
	ReasoningEnvironmentId *string `json:"ReasoningEnvironmentId,omitnil,omitempty" name:"ReasoningEnvironmentId"`

	// 模型自动清理开关(true/false)，当前版本仅支持SAVED_MODEL格式模型
	AutoClean *string `json:"AutoClean,omitnil,omitempty" name:"AutoClean"`

	// 模型数量保留上限(默认值为24个，上限为24，下限为1，步长为1)
	MaxReservedModels *uint64 `json:"MaxReservedModels,omitnil,omitempty" name:"MaxReservedModels"`

	// 模型清理周期(默认值为1分钟，上限为1440，下限为1分钟，步长为1)
	ModelCleanPeriod *uint64 `json:"ModelCleanPeriod,omitnil,omitempty" name:"ModelCleanPeriod"`

	// 是否QAT模型
	IsQAT *bool `json:"IsQAT,omitnil,omitempty" name:"IsQAT"`
}

type CreateTrainingModelRequest struct {
	*tchttp.BaseRequest
	
	// 导入方式
	// MODEL：导入新模型
	// VERSION：导入新版本
	// EXIST：导入现有版本
	ImportMethod *string `json:"ImportMethod,omitnil,omitempty" name:"ImportMethod"`

	// 推理环境来源（SYSTEM/CUSTOM）
	ReasoningEnvironmentSource *string `json:"ReasoningEnvironmentSource,omitnil,omitempty" name:"ReasoningEnvironmentSource"`

	// 模型名称，不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	TrainingModelName *string `json:"TrainingModelName,omitnil,omitempty" name:"TrainingModelName"`

	// 标签配置
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 训练任务名称
	TrainingJobName *string `json:"TrainingJobName,omitnil,omitempty" name:"TrainingJobName"`

	// 模型来源cos目录，以/结尾
	TrainingModelCosPath *CosPathInfo `json:"TrainingModelCosPath,omitnil,omitempty" name:"TrainingModelCosPath"`

	// 算法框架 （PYTORCH/TENSORFLOW/DETECTRON2/PMML/MMDETECTION)
	AlgorithmFramework *string `json:"AlgorithmFramework,omitnil,omitempty" name:"AlgorithmFramework"`

	// 推理环境
	ReasoningEnvironment *string `json:"ReasoningEnvironment,omitnil,omitempty" name:"ReasoningEnvironment"`

	// 训练指标，最多支持1000字符
	TrainingModelIndex *string `json:"TrainingModelIndex,omitnil,omitempty" name:"TrainingModelIndex"`

	// 模型版本
	TrainingModelVersion *string `json:"TrainingModelVersion,omitnil,omitempty" name:"TrainingModelVersion"`

	// 自定义推理环境
	ReasoningImageInfo *ImageInfo `json:"ReasoningImageInfo,omitnil,omitempty" name:"ReasoningImageInfo"`

	// 模型移动方式（CUT/COPY）
	ModelMoveMode *string `json:"ModelMoveMode,omitnil,omitempty" name:"ModelMoveMode"`

	// 训练任务ID
	TrainingJobId *string `json:"TrainingJobId,omitnil,omitempty" name:"TrainingJobId"`

	// 模型ID（导入新模型不需要，导入新版本需要）
	TrainingModelId *string `json:"TrainingModelId,omitnil,omitempty" name:"TrainingModelId"`

	// 模型存储cos目录
	ModelOutputPath *CosPathInfo `json:"ModelOutputPath,omitnil,omitempty" name:"ModelOutputPath"`

	// 模型来源 （JOB/COS）
	TrainingModelSource *string `json:"TrainingModelSource,omitnil,omitempty" name:"TrainingModelSource"`

	// 模型偏好
	TrainingPreference *string `json:"TrainingPreference,omitnil,omitempty" name:"TrainingPreference"`

	// 自动学习任务ID（已废弃）
	AutoMLTaskId *string `json:"AutoMLTaskId,omitnil,omitempty" name:"AutoMLTaskId"`

	// 任务版本
	TrainingJobVersion *string `json:"TrainingJobVersion,omitnil,omitempty" name:"TrainingJobVersion"`

	// 模型版本类型；
	// 枚举值：NORMAL(通用)  ACCELERATE(加速)
	// 注意:  默认为NORMAL
	ModelVersionType *string `json:"ModelVersionType,omitnil,omitempty" name:"ModelVersionType"`

	// 模型格式 （PYTORCH/TORCH_SCRIPT/DETECTRON2/SAVED_MODEL/FROZEN_GRAPH/PMML/MMDETECTION/ONNX/HUGGING_FACE）
	ModelFormat *string `json:"ModelFormat,omitnil,omitempty" name:"ModelFormat"`

	// 推理镜像ID
	ReasoningEnvironmentId *string `json:"ReasoningEnvironmentId,omitnil,omitempty" name:"ReasoningEnvironmentId"`

	// 模型自动清理开关(true/false)，当前版本仅支持SAVED_MODEL格式模型
	AutoClean *string `json:"AutoClean,omitnil,omitempty" name:"AutoClean"`

	// 模型数量保留上限(默认值为24个，上限为24，下限为1，步长为1)
	MaxReservedModels *uint64 `json:"MaxReservedModels,omitnil,omitempty" name:"MaxReservedModels"`

	// 模型清理周期(默认值为1分钟，上限为1440，下限为1分钟，步长为1)
	ModelCleanPeriod *uint64 `json:"ModelCleanPeriod,omitnil,omitempty" name:"ModelCleanPeriod"`

	// 是否QAT模型
	IsQAT *bool `json:"IsQAT,omitnil,omitempty" name:"IsQAT"`
}

func (r *CreateTrainingModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrainingModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImportMethod")
	delete(f, "ReasoningEnvironmentSource")
	delete(f, "TrainingModelName")
	delete(f, "Tags")
	delete(f, "TrainingJobName")
	delete(f, "TrainingModelCosPath")
	delete(f, "AlgorithmFramework")
	delete(f, "ReasoningEnvironment")
	delete(f, "TrainingModelIndex")
	delete(f, "TrainingModelVersion")
	delete(f, "ReasoningImageInfo")
	delete(f, "ModelMoveMode")
	delete(f, "TrainingJobId")
	delete(f, "TrainingModelId")
	delete(f, "ModelOutputPath")
	delete(f, "TrainingModelSource")
	delete(f, "TrainingPreference")
	delete(f, "AutoMLTaskId")
	delete(f, "TrainingJobVersion")
	delete(f, "ModelVersionType")
	delete(f, "ModelFormat")
	delete(f, "ReasoningEnvironmentId")
	delete(f, "AutoClean")
	delete(f, "MaxReservedModels")
	delete(f, "ModelCleanPeriod")
	delete(f, "IsQAT")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTrainingModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrainingModelResponseParams struct {
	// 模型ID，TrainingModel ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 模型版本ID
	TrainingModelVersionId *string `json:"TrainingModelVersionId,omitnil,omitempty" name:"TrainingModelVersionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTrainingModelResponse struct {
	*tchttp.BaseResponse
	Response *CreateTrainingModelResponseParams `json:"Response"`
}

func (r *CreateTrainingModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrainingModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrainingTaskRequestParams struct {
	// 训练任务名称，不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 计费模式，eg：PREPAID 包年包月（资源组）;
	// POSTPAID_BY_HOUR 按量计费
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 资源配置，需填写对应算力规格ID和节点数量，算力规格ID查询接口为DescribeBillingSpecsPrice，eg：[{"Role":"WORKER", "InstanceType": "TI.S.MEDIUM.POST", "InstanceNum": 1}]
	ResourceConfigInfos []*ResourceConfigInfo `json:"ResourceConfigInfos,omitnil,omitempty" name:"ResourceConfigInfos"`

	// 训练框架名称，通过DescribeTrainingFrameworks接口查询，eg：SPARK、PYSPARK、TENSORFLOW、PYTORCH
	FrameworkName *string `json:"FrameworkName,omitnil,omitempty" name:"FrameworkName"`

	// 训练框架版本，通过DescribeTrainingFrameworks接口查询，eg：1.15、1.9
	FrameworkVersion *string `json:"FrameworkVersion,omitnil,omitempty" name:"FrameworkVersion"`

	// 训练框架环境，通过DescribeTrainingFrameworks接口查询，eg：tf1.15-py3.7-cpu、torch1.9-py3.8-cuda11.1-gpu
	FrameworkEnvironment *string `json:"FrameworkEnvironment,omitnil,omitempty" name:"FrameworkEnvironment"`

	// 预付费专用资源组ID，通过DescribeBillingResourceGroups接口查询
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 标签配置
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 自定义镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// COS代码包路径
	CodePackagePath *CosPathInfo `json:"CodePackagePath,omitnil,omitempty" name:"CodePackagePath"`

	// 启动命令信息，默认为sh start.sh
	StartCmdInfo *StartCmdInfo `json:"StartCmdInfo,omitnil,omitempty" name:"StartCmdInfo"`

	// 训练模式，通过DescribeTrainingFrameworks接口查询，eg：PS_WORKER、DDP、MPI、HOROVOD
	TrainingMode *string `json:"TrainingMode,omitnil,omitempty" name:"TrainingMode"`

	// 数据配置，依赖DataSource字段，数量不超过10个
	DataConfigs []*DataConfig `json:"DataConfigs,omitnil,omitempty" name:"DataConfigs"`

	// VPC Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网Id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// COS训练输出路径
	Output *CosPathInfo `json:"Output,omitnil,omitempty" name:"Output"`

	// CLS日志配置
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// 调优参数，不超过2048个字符
	TuningParameters *string `json:"TuningParameters,omitnil,omitempty" name:"TuningParameters"`

	// 是否上报日志
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// 备注，不超过1024个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 数据来源，eg：DATASET、COS、CFS、CFSTurbo、HDFS、GooseFSx
	DataSource *string `json:"DataSource,omitnil,omitempty" name:"DataSource"`

	// 回调地址，用于创建/启动/停止训练任务的异步回调。回调格式&内容详见：[[TI-ONE接口回调说明]](https://cloud.tencent.com/document/product/851/84292)
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 太极预训练模型ID
	PreTrainModel *PreTrainModel `json:"PreTrainModel,omitnil,omitempty" name:"PreTrainModel"`
}

type CreateTrainingTaskRequest struct {
	*tchttp.BaseRequest
	
	// 训练任务名称，不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 计费模式，eg：PREPAID 包年包月（资源组）;
	// POSTPAID_BY_HOUR 按量计费
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 资源配置，需填写对应算力规格ID和节点数量，算力规格ID查询接口为DescribeBillingSpecsPrice，eg：[{"Role":"WORKER", "InstanceType": "TI.S.MEDIUM.POST", "InstanceNum": 1}]
	ResourceConfigInfos []*ResourceConfigInfo `json:"ResourceConfigInfos,omitnil,omitempty" name:"ResourceConfigInfos"`

	// 训练框架名称，通过DescribeTrainingFrameworks接口查询，eg：SPARK、PYSPARK、TENSORFLOW、PYTORCH
	FrameworkName *string `json:"FrameworkName,omitnil,omitempty" name:"FrameworkName"`

	// 训练框架版本，通过DescribeTrainingFrameworks接口查询，eg：1.15、1.9
	FrameworkVersion *string `json:"FrameworkVersion,omitnil,omitempty" name:"FrameworkVersion"`

	// 训练框架环境，通过DescribeTrainingFrameworks接口查询，eg：tf1.15-py3.7-cpu、torch1.9-py3.8-cuda11.1-gpu
	FrameworkEnvironment *string `json:"FrameworkEnvironment,omitnil,omitempty" name:"FrameworkEnvironment"`

	// 预付费专用资源组ID，通过DescribeBillingResourceGroups接口查询
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 标签配置
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 自定义镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// COS代码包路径
	CodePackagePath *CosPathInfo `json:"CodePackagePath,omitnil,omitempty" name:"CodePackagePath"`

	// 启动命令信息，默认为sh start.sh
	StartCmdInfo *StartCmdInfo `json:"StartCmdInfo,omitnil,omitempty" name:"StartCmdInfo"`

	// 训练模式，通过DescribeTrainingFrameworks接口查询，eg：PS_WORKER、DDP、MPI、HOROVOD
	TrainingMode *string `json:"TrainingMode,omitnil,omitempty" name:"TrainingMode"`

	// 数据配置，依赖DataSource字段，数量不超过10个
	DataConfigs []*DataConfig `json:"DataConfigs,omitnil,omitempty" name:"DataConfigs"`

	// VPC Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网Id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// COS训练输出路径
	Output *CosPathInfo `json:"Output,omitnil,omitempty" name:"Output"`

	// CLS日志配置
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// 调优参数，不超过2048个字符
	TuningParameters *string `json:"TuningParameters,omitnil,omitempty" name:"TuningParameters"`

	// 是否上报日志
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// 备注，不超过1024个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 数据来源，eg：DATASET、COS、CFS、CFSTurbo、HDFS、GooseFSx
	DataSource *string `json:"DataSource,omitnil,omitempty" name:"DataSource"`

	// 回调地址，用于创建/启动/停止训练任务的异步回调。回调格式&内容详见：[[TI-ONE接口回调说明]](https://cloud.tencent.com/document/product/851/84292)
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// 太极预训练模型ID
	PreTrainModel *PreTrainModel `json:"PreTrainModel,omitnil,omitempty" name:"PreTrainModel"`
}

func (r *CreateTrainingTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrainingTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ChargeType")
	delete(f, "ResourceConfigInfos")
	delete(f, "FrameworkName")
	delete(f, "FrameworkVersion")
	delete(f, "FrameworkEnvironment")
	delete(f, "ResourceGroupId")
	delete(f, "Tags")
	delete(f, "ImageInfo")
	delete(f, "CodePackagePath")
	delete(f, "StartCmdInfo")
	delete(f, "TrainingMode")
	delete(f, "DataConfigs")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Output")
	delete(f, "LogConfig")
	delete(f, "TuningParameters")
	delete(f, "LogEnable")
	delete(f, "Remark")
	delete(f, "DataSource")
	delete(f, "CallbackUrl")
	delete(f, "PreTrainModel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTrainingTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrainingTaskResponseParams struct {
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTrainingTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateTrainingTaskResponseParams `json:"Response"`
}

func (r *CreateTrainingTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrainingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CronInfo struct {
	// cron配置
	CronConfig *string `json:"CronConfig,omitnil,omitempty" name:"CronConfig"`

	// 周期开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 周期结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type CronScaleJob struct {
	// Cron表达式，标识任务的执行时间，精确到分钟级
	Schedule *string `json:"Schedule,omitnil,omitempty" name:"Schedule"`

	// 定时任务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 目标实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetReplicas *int64 `json:"TargetReplicas,omitnil,omitempty" name:"TargetReplicas"`

	// 目标min
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinReplicas *int64 `json:"MinReplicas,omitnil,omitempty" name:"MinReplicas"`

	// 目标max
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxReplicas *int64 `json:"MaxReplicas,omitnil,omitempty" name:"MaxReplicas"`

	// 例外时间，Cron表达式，在对应时间内不执行任务。最多支持3条。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcludeDates []*string `json:"ExcludeDates,omitnil,omitempty" name:"ExcludeDates"`
}

type CrossTenantENIInfo struct {
	// Pod IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrimaryIP *string `json:"PrimaryIP,omitnil,omitempty" name:"PrimaryIP"`

	// Pod Port
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`
}

type CustomTrainingData struct {
	// 指标名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 指标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metrics []*CustomTrainingMetric `json:"Metrics,omitnil,omitempty" name:"Metrics"`
}

type CustomTrainingMetric struct {
	// X轴数据类型: TIMESTAMP; EPOCH; STEP
	XType *string `json:"XType,omitnil,omitempty" name:"XType"`

	// 数据点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Points []*CustomTrainingPoint `json:"Points,omitnil,omitempty" name:"Points"`
}

type CustomTrainingPoint struct {
	// X值
	XValue *float64 `json:"XValue,omitnil,omitempty" name:"XValue"`

	// Y值
	YValue *float64 `json:"YValue,omitnil,omitempty" name:"YValue"`
}

type DataConfig struct {
	// 映射路径
	MappingPath *string `json:"MappingPath,omitnil,omitempty" name:"MappingPath"`

	// DATASET、COS、CFS、CFSTurbo、GooseFSx、HDFS、WEDATA_HDFS
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceType *string `json:"DataSourceType,omitnil,omitempty" name:"DataSourceType"`

	// 来自数据集的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSetSource *DataSetConfig `json:"DataSetSource,omitnil,omitempty" name:"DataSetSource"`

	// 来自cos的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	COSSource *CosPathInfo `json:"COSSource,omitnil,omitempty" name:"COSSource"`

	// 来自CFS的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFSSource *CFSConfig `json:"CFSSource,omitnil,omitempty" name:"CFSSource"`

	// 来自HDFS的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	HDFSSource *HDFSConfig `json:"HDFSSource,omitnil,omitempty" name:"HDFSSource"`

	// 配置GooseFS的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	GooseFSSource *GooseFS `json:"GooseFSSource,omitnil,omitempty" name:"GooseFSSource"`

	// 配置TurboFS的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFSTurboSource *CFSTurbo `json:"CFSTurboSource,omitnil,omitempty" name:"CFSTurboSource"`
}

type DataPoint struct {
	// 指标名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 值
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type DataSetConfig struct {
	// 数据集ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DatasetGroup struct {
	// 数据集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetId *string `json:"DatasetId,omitnil,omitempty" name:"DatasetId"`

	// 数据集名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetName *string `json:"DatasetName,omitnil,omitempty" name:"DatasetName"`

	// 创建者
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 数据集版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetVersion *string `json:"DatasetVersion,omitnil,omitempty" name:"DatasetVersion"`

	// 数据集类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetType *string `json:"DatasetType,omitnil,omitempty" name:"DatasetType"`

	// 数据集标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetTags []*Tag `json:"DatasetTags,omitnil,omitempty" name:"DatasetTags"`

	// 数据集标注任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetAnnotationTaskName *string `json:"DatasetAnnotationTaskName,omitnil,omitempty" name:"DatasetAnnotationTaskName"`

	// 数据集标注任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetAnnotationTaskId *string `json:"DatasetAnnotationTaskId,omitnil,omitempty" name:"DatasetAnnotationTaskId"`

	// 处理进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Process *uint64 `json:"Process,omitnil,omitempty" name:"Process"`

	// 数据集状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetStatus *string `json:"DatasetStatus,omitnil,omitempty" name:"DatasetStatus"`

	// 错误详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 外部关联TASKType
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalTaskType *string `json:"ExternalTaskType,omitnil,omitempty" name:"ExternalTaskType"`

	// 数据集大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetSize *string `json:"DatasetSize,omitnil,omitempty" name:"DatasetSize"`

	// 数据集数据量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileNum *uint64 `json:"FileNum,omitnil,omitempty" name:"FileNum"`

	// 数据集源COS路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageDataPath *CosPathInfo `json:"StorageDataPath,omitnil,omitempty" name:"StorageDataPath"`

	// 数据集标签存储路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageLabelPath *CosPathInfo `json:"StorageLabelPath,omitnil,omitempty" name:"StorageLabelPath"`

	// 数据集版本聚合详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetVersions []*DatasetInfo `json:"DatasetVersions,omitnil,omitempty" name:"DatasetVersions"`

	// 数据集标注状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationStatus *string `json:"AnnotationStatus,omitnil,omitempty" name:"AnnotationStatus"`

	// 数据集类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationType *string `json:"AnnotationType,omitnil,omitempty" name:"AnnotationType"`

	// 数据集标注格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationFormat *string `json:"AnnotationFormat,omitnil,omitempty" name:"AnnotationFormat"`

	// 数据集范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetScope *string `json:"DatasetScope,omitnil,omitempty" name:"DatasetScope"`

	// 数据集OCR子场景
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrScene *string `json:"OcrScene,omitnil,omitempty" name:"OcrScene"`

	// 数据集字典修改状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationKeyStatus *string `json:"AnnotationKeyStatus,omitnil,omitempty" name:"AnnotationKeyStatus"`

	// 文本数据集导入方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`
}

type DatasetInfo struct {
	// 数据集id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetId *string `json:"DatasetId,omitnil,omitempty" name:"DatasetId"`

	// 数据集名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetName *string `json:"DatasetName,omitnil,omitempty" name:"DatasetName"`

	// 数据集创建者
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 数据集版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetVersion *string `json:"DatasetVersion,omitnil,omitempty" name:"DatasetVersion"`

	// 数据集类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetType *string `json:"DatasetType,omitnil,omitempty" name:"DatasetType"`

	// 数据集标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetTags []*Tag `json:"DatasetTags,omitnil,omitempty" name:"DatasetTags"`

	// 数据集对应标注任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetAnnotationTaskName *string `json:"DatasetAnnotationTaskName,omitnil,omitempty" name:"DatasetAnnotationTaskName"`

	// 数据集对应标注任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetAnnotationTaskId *string `json:"DatasetAnnotationTaskId,omitnil,omitempty" name:"DatasetAnnotationTaskId"`

	// 处理进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Process *int64 `json:"Process,omitnil,omitempty" name:"Process"`

	// 数据集状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetStatus *string `json:"DatasetStatus,omitnil,omitempty" name:"DatasetStatus"`

	// 错误详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 数据集创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 数据集更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 外部任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalTaskType *string `json:"ExternalTaskType,omitnil,omitempty" name:"ExternalTaskType"`

	// 数据集存储大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetSize *string `json:"DatasetSize,omitnil,omitempty" name:"DatasetSize"`

	// 数据集数据数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileNum *uint64 `json:"FileNum,omitnil,omitempty" name:"FileNum"`

	// 数据集源cos 路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageDataPath *CosPathInfo `json:"StorageDataPath,omitnil,omitempty" name:"StorageDataPath"`

	// 数据集输出cos路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageLabelPath *CosPathInfo `json:"StorageLabelPath,omitnil,omitempty" name:"StorageLabelPath"`

	// 数据集标注状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationStatus *string `json:"AnnotationStatus,omitnil,omitempty" name:"AnnotationStatus"`

	// 数据集类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationType *string `json:"AnnotationType,omitnil,omitempty" name:"AnnotationType"`

	// 数据集标注格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationFormat *string `json:"AnnotationFormat,omitnil,omitempty" name:"AnnotationFormat"`

	// 数据集范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetScope *string `json:"DatasetScope,omitnil,omitempty" name:"DatasetScope"`

	// 数据集OCR子场景
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrScene *string `json:"OcrScene,omitnil,omitempty" name:"OcrScene"`

	// 数据集字典修改状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationKeyStatus *string `json:"AnnotationKeyStatus,omitnil,omitempty" name:"AnnotationKeyStatus"`
}

type DefaultInnerCallInfo struct {
	// 可以进行调用的VPC-ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// 默认内网调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerHttpAddr *string `json:"InnerHttpAddr,omitnil,omitempty" name:"InnerHttpAddr"`
}

type DefaultNginxGatewayCallInfo struct {
	// host
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

// Predefined struct for user
type DeleteBatchTaskRequestParams struct {
	// 批量预测任务ID
	BatchTaskId *string `json:"BatchTaskId,omitnil,omitempty" name:"BatchTaskId"`
}

type DeleteBatchTaskRequest struct {
	*tchttp.BaseRequest
	
	// 批量预测任务ID
	BatchTaskId *string `json:"BatchTaskId,omitnil,omitempty" name:"BatchTaskId"`
}

func (r *DeleteBatchTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBatchTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBatchTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBatchTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteBatchTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBatchTaskResponseParams `json:"Response"`
}

func (r *DeleteBatchTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBatchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDatasetRequestParams struct {
	// 数据集id
	DatasetId *string `json:"DatasetId,omitnil,omitempty" name:"DatasetId"`

	// 是否删除cos标签文件
	DeleteLabelEnable *bool `json:"DeleteLabelEnable,omitnil,omitempty" name:"DeleteLabelEnable"`
}

type DeleteDatasetRequest struct {
	*tchttp.BaseRequest
	
	// 数据集id
	DatasetId *string `json:"DatasetId,omitnil,omitempty" name:"DatasetId"`

	// 是否删除cos标签文件
	DeleteLabelEnable *bool `json:"DeleteLabelEnable,omitnil,omitempty" name:"DeleteLabelEnable"`
}

func (r *DeleteDatasetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDatasetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasetId")
	delete(f, "DeleteLabelEnable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDatasetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDatasetResponseParams struct {
	// 删除的datasetId
	DatasetId *string `json:"DatasetId,omitnil,omitempty" name:"DatasetId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDatasetResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDatasetResponseParams `json:"Response"`
}

func (r *DeleteDatasetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDatasetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelAccelerateTaskRequestParams struct {
	// 模型加速任务ID
	ModelAccTaskId *string `json:"ModelAccTaskId,omitnil,omitempty" name:"ModelAccTaskId"`
}

type DeleteModelAccelerateTaskRequest struct {
	*tchttp.BaseRequest
	
	// 模型加速任务ID
	ModelAccTaskId *string `json:"ModelAccTaskId,omitnil,omitempty" name:"ModelAccTaskId"`
}

func (r *DeleteModelAccelerateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelAccelerateTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelAccTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteModelAccelerateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelAccelerateTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteModelAccelerateTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteModelAccelerateTaskResponseParams `json:"Response"`
}

func (r *DeleteModelAccelerateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelAccelerateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelServiceGroupRequestParams struct {
	// 服务id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`
}

type DeleteModelServiceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 服务id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`
}

func (r *DeleteModelServiceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelServiceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteModelServiceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelServiceGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteModelServiceGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteModelServiceGroupResponseParams `json:"Response"`
}

func (r *DeleteModelServiceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelServiceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelServiceRequestParams struct {
	// 服务id
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 服务分类
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

type DeleteModelServiceRequest struct {
	*tchttp.BaseRequest
	
	// 服务id
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 服务分类
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

func (r *DeleteModelServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ServiceCategory")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteModelServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelServiceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteModelServiceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteModelServiceResponseParams `json:"Response"`
}

func (r *DeleteModelServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNotebookImageRecordRequestParams struct {
	// 记录id
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`
}

type DeleteNotebookImageRecordRequest struct {
	*tchttp.BaseRequest
	
	// 记录id
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`
}

func (r *DeleteNotebookImageRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNotebookImageRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RecordId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNotebookImageRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNotebookImageRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNotebookImageRecordResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNotebookImageRecordResponseParams `json:"Response"`
}

func (r *DeleteNotebookImageRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNotebookImageRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNotebookRequestParams struct {
	// notebook id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteNotebookRequest struct {
	*tchttp.BaseRequest
	
	// notebook id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DeleteNotebookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNotebookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNotebookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNotebookResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNotebookResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNotebookResponseParams `json:"Response"`
}

func (r *DeleteNotebookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNotebookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrainingModelRequestParams struct {
	// 模型ID
	TrainingModelId *string `json:"TrainingModelId,omitnil,omitempty" name:"TrainingModelId"`

	// 是否同步清理cos
	EnableDeleteCos *bool `json:"EnableDeleteCos,omitnil,omitempty" name:"EnableDeleteCos"`

	// 删除模型类型，枚举值：NORMAL 普通，ACCELERATE 加速，不传则删除所有
	ModelVersionType *string `json:"ModelVersionType,omitnil,omitempty" name:"ModelVersionType"`
}

type DeleteTrainingModelRequest struct {
	*tchttp.BaseRequest
	
	// 模型ID
	TrainingModelId *string `json:"TrainingModelId,omitnil,omitempty" name:"TrainingModelId"`

	// 是否同步清理cos
	EnableDeleteCos *bool `json:"EnableDeleteCos,omitnil,omitempty" name:"EnableDeleteCos"`

	// 删除模型类型，枚举值：NORMAL 普通，ACCELERATE 加速，不传则删除所有
	ModelVersionType *string `json:"ModelVersionType,omitnil,omitempty" name:"ModelVersionType"`
}

func (r *DeleteTrainingModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrainingModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrainingModelId")
	delete(f, "EnableDeleteCos")
	delete(f, "ModelVersionType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTrainingModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrainingModelResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTrainingModelResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTrainingModelResponseParams `json:"Response"`
}

func (r *DeleteTrainingModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrainingModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrainingModelVersionRequestParams struct {
	// 模型版本ID
	TrainingModelVersionId *string `json:"TrainingModelVersionId,omitnil,omitempty" name:"TrainingModelVersionId"`

	// 是否同步清理cos
	EnableDeleteCos *bool `json:"EnableDeleteCos,omitnil,omitempty" name:"EnableDeleteCos"`
}

type DeleteTrainingModelVersionRequest struct {
	*tchttp.BaseRequest
	
	// 模型版本ID
	TrainingModelVersionId *string `json:"TrainingModelVersionId,omitnil,omitempty" name:"TrainingModelVersionId"`

	// 是否同步清理cos
	EnableDeleteCos *bool `json:"EnableDeleteCos,omitnil,omitempty" name:"EnableDeleteCos"`
}

func (r *DeleteTrainingModelVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrainingModelVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrainingModelVersionId")
	delete(f, "EnableDeleteCos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTrainingModelVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrainingModelVersionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTrainingModelVersionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTrainingModelVersionResponseParams `json:"Response"`
}

func (r *DeleteTrainingModelVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrainingModelVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrainingTaskRequestParams struct {
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteTrainingTaskRequest struct {
	*tchttp.BaseRequest
	
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DeleteTrainingTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrainingTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTrainingTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrainingTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTrainingTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTrainingTaskResponseParams `json:"Response"`
}

func (r *DeleteTrainingTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrainingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAPIConfigsRequestParams struct {
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序的依据字段， 取值范围 "CreateTime" "UpdateTime"
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 分页参数，支持的分页过滤Name包括：
	// ["ClusterId", "ServiceId", "ServiceGroupName", "ServiceGroupId"]
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeAPIConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序的依据字段， 取值范围 "CreateTime" "UpdateTime"
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 分页参数，支持的分页过滤Name包括：
	// ["ClusterId", "ServiceId", "ServiceGroupName", "ServiceGroupId"]
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeAPIConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAPIConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAPIConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAPIConfigsResponseParams struct {
	// 接口数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 接口详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Details []*APIConfigDetail `json:"Details,omitnil,omitempty" name:"Details"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAPIConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAPIConfigsResponseParams `json:"Response"`
}

func (r *DescribeAPIConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAPIConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchTaskInstancesRequestParams struct {
	// 批量预测任务id
	BatchTaskId *string `json:"BatchTaskId,omitnil,omitempty" name:"BatchTaskId"`
}

type DescribeBatchTaskInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 批量预测任务id
	BatchTaskId *string `json:"BatchTaskId,omitnil,omitempty" name:"BatchTaskId"`
}

func (r *DescribeBatchTaskInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchTaskInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBatchTaskInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchTaskInstancesResponseParams struct {
	// 批量预测实例结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchInstances []*BatchTaskInstance `json:"BatchInstances,omitnil,omitempty" name:"BatchInstances"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBatchTaskInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBatchTaskInstancesResponseParams `json:"Response"`
}

func (r *DescribeBatchTaskInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchTaskInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchTaskRequestParams struct {
	// 批量预测任务ID
	BatchTaskId *string `json:"BatchTaskId,omitnil,omitempty" name:"BatchTaskId"`
}

type DescribeBatchTaskRequest struct {
	*tchttp.BaseRequest
	
	// 批量预测任务ID
	BatchTaskId *string `json:"BatchTaskId,omitnil,omitempty" name:"BatchTaskId"`
}

func (r *DescribeBatchTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBatchTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchTaskResponseParams struct {
	// 批量预测任务详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchTaskDetail *BatchTaskDetail `json:"BatchTaskDetail,omitnil,omitempty" name:"BatchTaskDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBatchTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBatchTaskResponseParams `json:"Response"`
}

func (r *DescribeBatchTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchTasksRequestParams struct {
	// 过滤器
	// 
	// Name支持的取值范围：
	// BatchTaskId, BatchTaskName, Status, ModelVersionId
	// 
	// 其中Status 的有效取值范围:
	// INIT, STARTING, RUNNING, FAILED, STOPPING, STOPPED, SUCCEED
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤器
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目，默认为20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC（升序排列）/ DESC（降序排列），默认为DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段。当前仅支持 "CreateTime"。
	// 不传此字段则按照DB默认排序结果返回
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeBatchTasksRequest struct {
	*tchttp.BaseRequest
	
	// 过滤器
	// 
	// Name支持的取值范围：
	// BatchTaskId, BatchTaskName, Status, ModelVersionId
	// 
	// 其中Status 的有效取值范围:
	// INIT, STARTING, RUNNING, FAILED, STOPPING, STOPPED, SUCCEED
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤器
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目，默认为20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC（升序排列）/ DESC（降序排列），默认为DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段。当前仅支持 "CreateTime"。
	// 不传此字段则按照DB默认排序结果返回
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeBatchTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "TagFilters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBatchTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchTasksResponseParams struct {
	// 批量预测任务总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 批量预测任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchTaskSet []*BatchTaskSetItem `json:"BatchTaskSet,omitnil,omitempty" name:"BatchTaskSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBatchTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBatchTasksResponseParams `json:"Response"`
}

func (r *DescribeBatchTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingResourceGroupRequestParams struct {
	// 资源组id, 取值为创建资源组接口(CreateBillingResourceGroup)响应中的ResourceGroupId
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 过滤条件
	// 注意: 
	// 1. Filter.Name 只支持以下枚举值:
	//     InstanceId (资源组节点id)
	//     InstanceStatus (资源组节点状态)
	// 2. Filter.Values: 长度为1且Filter.Fuzzy=true时，支持模糊查询; 不为1时，精确查询
	// 3. 每次请求的Filters的上限为10，Filter.Values的上限为100
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询起始位置，如：Limit为10，第一页Offset为0，第二页Offset为10....即每页左边为闭区间; 默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询每页大小，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方向; 枚举值: ASC | DESC；默认DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段; 枚举值: CreateTime (创建时间) ｜ ExpireTime (到期时间)；默认CreateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeBillingResourceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 资源组id, 取值为创建资源组接口(CreateBillingResourceGroup)响应中的ResourceGroupId
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 过滤条件
	// 注意: 
	// 1. Filter.Name 只支持以下枚举值:
	//     InstanceId (资源组节点id)
	//     InstanceStatus (资源组节点状态)
	// 2. Filter.Values: 长度为1且Filter.Fuzzy=true时，支持模糊查询; 不为1时，精确查询
	// 3. 每次请求的Filters的上限为10，Filter.Values的上限为100
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询起始位置，如：Limit为10，第一页Offset为0，第二页Offset为10....即每页左边为闭区间; 默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询每页大小，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方向; 枚举值: ASC | DESC；默认DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段; 枚举值: CreateTime (创建时间) ｜ ExpireTime (到期时间)；默认CreateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeBillingResourceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingResourceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceGroupId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillingResourceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingResourceGroupResponseParams struct {
	// 资源组节点总数； 注意接口是分页拉取的，total是指资源组节点总数，不是本次返回中InstanceSet数组的大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 资源组节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceSet []*Instance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// 资源组纳管类型
	ResourceGroupSWType *string `json:"ResourceGroupSWType,omitnil,omitempty" name:"ResourceGroupSWType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillingResourceGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillingResourceGroupResponseParams `json:"Response"`
}

func (r *DescribeBillingResourceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingResourceGroupsRequestParams struct {
	// 资源组类型;
	// 枚举值:
	// 空: 通用, TRAIN: 训练, INFERENCE: 推理
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter.Name: 枚举值: ResourceGroupId (资源组id列表)
	//                     ResourceGroupName (资源组名称列表)
	// Filter.Values: 长度为1且Filter.Fuzzy=true时，支持模糊查询; 不为1时，精确查询
	// 每次请求的Filters的上限为5，Filter.Values的上限为100
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 偏移量，默认为0；分页查询起始位置，如：Limit为100，第一页Offset为0，第二页OffSet为100....即每页左边为闭区间
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询每页大小，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 支持模糊查找资源组id和资源组名
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 是否不展示节点列表; 
	// true: 不展示，false 展示；
	// 默认为false
	DontShowInstanceSet *bool `json:"DontShowInstanceSet,omitnil,omitempty" name:"DontShowInstanceSet"`
}

type DescribeBillingResourceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 资源组类型;
	// 枚举值:
	// 空: 通用, TRAIN: 训练, INFERENCE: 推理
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Filter.Name: 枚举值: ResourceGroupId (资源组id列表)
	//                     ResourceGroupName (资源组名称列表)
	// Filter.Values: 长度为1且Filter.Fuzzy=true时，支持模糊查询; 不为1时，精确查询
	// 每次请求的Filters的上限为5，Filter.Values的上限为100
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 偏移量，默认为0；分页查询起始位置，如：Limit为100，第一页Offset为0，第二页OffSet为100....即每页左边为闭区间
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询每页大小，默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 支持模糊查找资源组id和资源组名
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 是否不展示节点列表; 
	// true: 不展示，false 展示；
	// 默认为false
	DontShowInstanceSet *bool `json:"DontShowInstanceSet,omitnil,omitempty" name:"DontShowInstanceSet"`
}

func (r *DescribeBillingResourceGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingResourceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Filters")
	delete(f, "TagFilters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchWord")
	delete(f, "DontShowInstanceSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillingResourceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingResourceGroupsResponseParams struct {
	// 资源组总数； 注意接口是分页拉取的，total是指资源组总数，不是本次返回中ResourceGroupSet数组的大小
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 资源组详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupSet []*ResourceGroup `json:"ResourceGroupSet,omitnil,omitempty" name:"ResourceGroupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillingResourceGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillingResourceGroupsResponseParams `json:"Response"`
}

func (r *DescribeBillingResourceGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingResourceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingResourceInstanceRunningJobsRequestParams struct {
	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 资源组节点id
	ResourceInstanceId *string `json:"ResourceInstanceId,omitnil,omitempty" name:"ResourceInstanceId"`
}

type DescribeBillingResourceInstanceRunningJobsRequest struct {
	*tchttp.BaseRequest
	
	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 资源组节点id
	ResourceInstanceId *string `json:"ResourceInstanceId,omitnil,omitempty" name:"ResourceInstanceId"`
}

func (r *DescribeBillingResourceInstanceRunningJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingResourceInstanceRunningJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceGroupId")
	delete(f, "ResourceInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillingResourceInstanceRunningJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingResourceInstanceRunningJobsResponseParams struct {
	// 资源组节点运行中的任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceInstanceRunningJobInfos []*ResourceInstanceRunningJobInfo `json:"ResourceInstanceRunningJobInfos,omitnil,omitempty" name:"ResourceInstanceRunningJobInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillingResourceInstanceRunningJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillingResourceInstanceRunningJobsResponseParams `json:"Response"`
}

func (r *DescribeBillingResourceInstanceRunningJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingResourceInstanceRunningJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingSpecsPriceRequestParams struct {
	// 询价参数，支持批量询价
	SpecsParam []*SpecUnit `json:"SpecsParam,omitnil,omitempty" name:"SpecsParam"`
}

type DescribeBillingSpecsPriceRequest struct {
	*tchttp.BaseRequest
	
	// 询价参数，支持批量询价
	SpecsParam []*SpecUnit `json:"SpecsParam,omitnil,omitempty" name:"SpecsParam"`
}

func (r *DescribeBillingSpecsPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingSpecsPriceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpecsParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillingSpecsPriceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingSpecsPriceResponseParams struct {
	// 计费项价格，支持批量返回
	SpecsPrice []*SpecPrice `json:"SpecsPrice,omitnil,omitempty" name:"SpecsPrice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillingSpecsPriceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillingSpecsPriceResponseParams `json:"Response"`
}

func (r *DescribeBillingSpecsPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingSpecsPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingSpecsRequestParams struct {
	// 付费模式：POSTPAID_BY_HOUR按量计费、PREPAID包年包月
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 枚举值：TRAIN、NOTEBOOK、INFERENCE
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 资源类型：CALC 计算资源、CPU CPU资源、GPU GPU资源、CBS云硬盘
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

type DescribeBillingSpecsRequest struct {
	*tchttp.BaseRequest
	
	// 付费模式：POSTPAID_BY_HOUR按量计费、PREPAID包年包月
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 枚举值：TRAIN、NOTEBOOK、INFERENCE
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 资源类型：CALC 计算资源、CPU CPU资源、GPU GPU资源、CBS云硬盘
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

func (r *DescribeBillingSpecsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingSpecsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChargeType")
	delete(f, "TaskType")
	delete(f, "ResourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillingSpecsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingSpecsResponseParams struct {
	// 计费项列表
	Specs []*Spec `json:"Specs,omitnil,omitempty" name:"Specs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillingSpecsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillingSpecsResponseParams `json:"Response"`
}

func (r *DescribeBillingSpecsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingSpecsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasetDetailStructuredRequestParams struct {
	// 数据集ID
	DatasetId *string `json:"DatasetId,omitnil,omitempty" name:"DatasetId"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数据条数，默认20，目前最大支持2000条数据
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDatasetDetailStructuredRequest struct {
	*tchttp.BaseRequest
	
	// 数据集ID
	DatasetId *string `json:"DatasetId,omitnil,omitempty" name:"DatasetId"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数据条数，默认20，目前最大支持2000条数据
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDatasetDetailStructuredRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasetDetailStructuredRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasetId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatasetDetailStructuredRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasetDetailStructuredResponseParams struct {
	// 数据总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 表格头信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnNames []*string `json:"ColumnNames,omitnil,omitempty" name:"ColumnNames"`

	// 表格内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	RowItems []*RowItem `json:"RowItems,omitnil,omitempty" name:"RowItems"`

	// 文本内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	RowTexts []*string `json:"RowTexts,omitnil,omitempty" name:"RowTexts"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatasetDetailStructuredResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatasetDetailStructuredResponseParams `json:"Response"`
}

func (r *DescribeDatasetDetailStructuredResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasetDetailStructuredResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasetDetailUnstructuredRequestParams struct {
	// 数据集ID
	DatasetId *string `json:"DatasetId,omitnil,omitempty" name:"DatasetId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回个数，默认20，目前最大支持2000条数据
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 标签过滤参数，对应标签值
	LabelList []*string `json:"LabelList,omitnil,omitempty" name:"LabelList"`

	// 标注状态过滤参数:
	// STATUS_ANNOTATED，已标注
	// STATUS_NON_ANNOTATED，未标注
	// STATUS_ALL，全部
	// 默认为STATUS_ALL
	AnnotationStatus *string `json:"AnnotationStatus,omitnil,omitempty" name:"AnnotationStatus"`

	// 数据集ID列表
	DatasetIds []*string `json:"DatasetIds,omitnil,omitempty" name:"DatasetIds"`

	// 要筛选的文本分类场景标签信息
	TextClassificationLabels []*TextLabelDistributionInfo `json:"TextClassificationLabels,omitnil,omitempty" name:"TextClassificationLabels"`
}

type DescribeDatasetDetailUnstructuredRequest struct {
	*tchttp.BaseRequest
	
	// 数据集ID
	DatasetId *string `json:"DatasetId,omitnil,omitempty" name:"DatasetId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回个数，默认20，目前最大支持2000条数据
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 标签过滤参数，对应标签值
	LabelList []*string `json:"LabelList,omitnil,omitempty" name:"LabelList"`

	// 标注状态过滤参数:
	// STATUS_ANNOTATED，已标注
	// STATUS_NON_ANNOTATED，未标注
	// STATUS_ALL，全部
	// 默认为STATUS_ALL
	AnnotationStatus *string `json:"AnnotationStatus,omitnil,omitempty" name:"AnnotationStatus"`

	// 数据集ID列表
	DatasetIds []*string `json:"DatasetIds,omitnil,omitempty" name:"DatasetIds"`

	// 要筛选的文本分类场景标签信息
	TextClassificationLabels []*TextLabelDistributionInfo `json:"TextClassificationLabels,omitnil,omitempty" name:"TextClassificationLabels"`
}

func (r *DescribeDatasetDetailUnstructuredRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasetDetailUnstructuredRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasetId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "LabelList")
	delete(f, "AnnotationStatus")
	delete(f, "DatasetIds")
	delete(f, "TextClassificationLabels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatasetDetailUnstructuredRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasetDetailUnstructuredResponseParams struct {
	// 已标注数据量
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotatedTotalCount *uint64 `json:"AnnotatedTotalCount,omitnil,omitempty" name:"AnnotatedTotalCount"`

	// 没有标注数据量
	// 注意：此字段可能返回 null，表示取不到有效值。
	NonAnnotatedTotalCount *uint64 `json:"NonAnnotatedTotalCount,omitnil,omitempty" name:"NonAnnotatedTotalCount"`

	// 过滤数据总量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterTotalCount *uint64 `json:"FilterTotalCount,omitnil,omitempty" name:"FilterTotalCount"`

	// 过滤数据详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterLabelList []*FilterLabelInfo `json:"FilterLabelList,omitnil,omitempty" name:"FilterLabelList"`

	// 数据文本行，默认返回前1000行
	// 注意：此字段可能返回 null，表示取不到有效值。
	RowTexts []*string `json:"RowTexts,omitnil,omitempty" name:"RowTexts"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatasetDetailUnstructuredResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatasetDetailUnstructuredResponseParams `json:"Response"`
}

func (r *DescribeDatasetDetailUnstructuredResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasetDetailUnstructuredResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasetsRequestParams struct {
	// 数据集id列表
	DatasetIds []*string `json:"DatasetIds,omitnil,omitempty" name:"DatasetIds"`

	// 数据集查询过滤条件，多个Filter之间的关系为逻辑与（AND）关系，过滤字段Filter.Name，类型为String
	// DatasetName，数据集名称
	// DatasetScope，数据集范围，SCOPE_DATASET_PRIVATE或SCOPE_DATASET_PUBLIC
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤条件
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 排序值，支持Asc或Desc，默认Desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段，支持CreateTime或UpdateTime，默认CreateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数据个数，默认20，最大支持200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDatasetsRequest struct {
	*tchttp.BaseRequest
	
	// 数据集id列表
	DatasetIds []*string `json:"DatasetIds,omitnil,omitempty" name:"DatasetIds"`

	// 数据集查询过滤条件，多个Filter之间的关系为逻辑与（AND）关系，过滤字段Filter.Name，类型为String
	// DatasetName，数据集名称
	// DatasetScope，数据集范围，SCOPE_DATASET_PRIVATE或SCOPE_DATASET_PUBLIC
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤条件
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 排序值，支持Asc或Desc，默认Desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段，支持CreateTime或UpdateTime，默认CreateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数据个数，默认20，最大支持200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDatasetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasetIds")
	delete(f, "Filters")
	delete(f, "TagFilters")
	delete(f, "Order")
	delete(f, "OrderField")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatasetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasetsResponseParams struct {
	// 数据集总量（名称维度）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 数据集按照数据集名称聚合的分组
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetGroups []*DatasetGroup `json:"DatasetGroups,omitnil,omitempty" name:"DatasetGroups"`

	// 数据集ID总量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetIdNums *uint64 `json:"DatasetIdNums,omitnil,omitempty" name:"DatasetIdNums"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatasetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatasetsResponseParams `json:"Response"`
}

func (r *DescribeDatasetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventsRequestParams struct {
	// 服务类型，TRAIN为任务式建模, NOTEBOOK为Notebook, INFER为在线服务, BATCH为批量预测
	// 枚举值：
	// - TRAIN
	// - NOTEBOOK
	// - INFER
	// - BATCH
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 服务ID，和Service参数对应，不同Service的服务ID获取方式不同，具体如下：
	// - Service类型为TRAIN：
	//   调用[DescribeTrainingTask接口](/document/product/851/75089)查询训练任务详情，ServiceId为接口返回值中Response.TrainingTaskDetail.LatestInstanceId
	// - Service类型为NOTEBOOK：
	//   调用[DescribeNotebook接口](/document/product/851/95662)查询Notebook详情，ServiceId为接口返回值中Response.NotebookDetail.PodName
	// - Service类型为INFER：
	//   调用[DescribeModelServiceGroup接口](/document/product/851/82285)查询服务组详情，ServiceId为接口返回值中Response.ServiceGroup.Services.ServiceId
	// - Service类型为BATCH：
	//   调用[DescribeBatchTask接口](/document/product/851/80180)查询跑批任务详情，ServiceId为接口返回值中Response.BatchTaskDetail.LatestInstanceId
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 查询事件最早发生的时间（RFC3339格式的时间字符串），默认值为当前时间的前一天
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询事件最晚发生的时间（RFC3339格式的时间字符串），默认值为当前时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页Limit，默认值为100，最大值为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页Offset，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排列顺序（可选值为ASC, DESC ），默认为DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序的依据字段（可选值为FirstTimestamp, LastTimestamp），默认值为LastTimestamp
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 过滤条件
	// 注意: 
	// 1. Filter.Name：目前支持ResourceKind（按事件关联的资源类型过滤）；Type（按事件类型过滤）
	// 2. Filter.Values：
	// 对于Name为ResourceKind，Values的可选取值为Deployment, Replicaset, Pod等K8S资源类型；
	// 对于Name为Type，Values的可选取值仅为Normal或者Warning；
	// Values为多个的时候表示同时满足
	// 3. Filter. Negative和Filter. Fuzzy没有使用
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeEventsRequest struct {
	*tchttp.BaseRequest
	
	// 服务类型，TRAIN为任务式建模, NOTEBOOK为Notebook, INFER为在线服务, BATCH为批量预测
	// 枚举值：
	// - TRAIN
	// - NOTEBOOK
	// - INFER
	// - BATCH
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 服务ID，和Service参数对应，不同Service的服务ID获取方式不同，具体如下：
	// - Service类型为TRAIN：
	//   调用[DescribeTrainingTask接口](/document/product/851/75089)查询训练任务详情，ServiceId为接口返回值中Response.TrainingTaskDetail.LatestInstanceId
	// - Service类型为NOTEBOOK：
	//   调用[DescribeNotebook接口](/document/product/851/95662)查询Notebook详情，ServiceId为接口返回值中Response.NotebookDetail.PodName
	// - Service类型为INFER：
	//   调用[DescribeModelServiceGroup接口](/document/product/851/82285)查询服务组详情，ServiceId为接口返回值中Response.ServiceGroup.Services.ServiceId
	// - Service类型为BATCH：
	//   调用[DescribeBatchTask接口](/document/product/851/80180)查询跑批任务详情，ServiceId为接口返回值中Response.BatchTaskDetail.LatestInstanceId
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 查询事件最早发生的时间（RFC3339格式的时间字符串），默认值为当前时间的前一天
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询事件最晚发生的时间（RFC3339格式的时间字符串），默认值为当前时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页Limit，默认值为100，最大值为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页Offset，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排列顺序（可选值为ASC, DESC ），默认为DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序的依据字段（可选值为FirstTimestamp, LastTimestamp），默认值为LastTimestamp
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 过滤条件
	// 注意: 
	// 1. Filter.Name：目前支持ResourceKind（按事件关联的资源类型过滤）；Type（按事件类型过滤）
	// 2. Filter.Values：
	// 对于Name为ResourceKind，Values的可选取值为Deployment, Replicaset, Pod等K8S资源类型；
	// 对于Name为Type，Values的可选取值仅为Normal或者Warning；
	// Values为多个的时候表示同时满足
	// 3. Filter. Negative和Filter. Fuzzy没有使用
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Service")
	delete(f, "ServiceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "OrderField")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventsResponseParams struct {
	// 事件的列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Events []*Event `json:"Events,omitnil,omitempty" name:"Events"`

	// 此次查询的事件的个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEventsResponseParams `json:"Response"`
}

func (r *DescribeEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInferTemplatesRequestParams struct {

}

type DescribeInferTemplatesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInferTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInferTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInferTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInferTemplatesResponseParams struct {
	// 模板列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkTemplates []*InferTemplateGroup `json:"FrameworkTemplates,omitnil,omitempty" name:"FrameworkTemplates"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInferTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInferTemplatesResponseParams `json:"Response"`
}

func (r *DescribeInferTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInferTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLatestTrainingMetricsRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeLatestTrainingMetricsRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeLatestTrainingMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLatestTrainingMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLatestTrainingMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLatestTrainingMetricsResponseParams struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 最近一次上报的训练指标.每个Metric中只有一个点的数据, 即len(Values) = len(Timestamps) = 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metrics []*TrainingMetric `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLatestTrainingMetricsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLatestTrainingMetricsResponseParams `json:"Response"`
}

func (r *DescribeLatestTrainingMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLatestTrainingMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogsRequestParams struct {
	// 服务类型，TRAIN为任务式建模, NOTEBOOK为Notebook, INFER为在线服务, BATCH为批量预测
	// 枚举值：
	// - TRAIN
	// - NOTEBOOK
	// - INFER
	// - BATCH
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 日志查询开始时间（RFC3339格式的时间字符串），默认值为当前时间的前一个小时
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 日志查询结束时间（RFC3339格式的时间字符串），默认值为当前时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 日志查询条数，默认值100，最大值100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 服务ID，和Service参数对应，不同Service的服务ID获取方式不同，具体如下：
	// - Service类型为TRAIN：
	//   调用[DescribeTrainingTask接口](/document/product/851/75089)查询训练任务详情，ServiceId为接口返回值中Response.TrainingTaskDetail.LatestInstanceId
	// - Service类型为NOTEBOOK：
	//   调用[DescribeNotebook接口](/document/product/851/95662)查询Notebook详情，ServiceId为接口返回值中Response.NotebookDetail.PodName
	// - Service类型为INFER：
	//   调用[DescribeModelServiceGroup接口](/document/product/851/82285)查询服务组详情，ServiceId为接口返回值中Response.ServiceGroup.Services.ServiceId
	// - Service类型为BATCH：
	//   调用[DescribeBatchTask接口](/document/product/851/80180)查询跑批任务详情，ServiceId为接口返回值中Response.BatchTaskDetail.LatestInstanceId
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Pod的名称，即需要查询服务对应的Pod，和Service参数对应，不同Service的PodName获取方式不同，具体如下：
	// - Service类型为TRAIN：
	//   调用[DescribeTrainingTaskPods接口](/document/product/851/75088)查询训练任务pod列表，PodName为接口返回值中Response.PodNames
	// - Service类型为NOTEBOOK：
	//   调用[DescribeNotebook接口](/document/product/851/95662)查询Notebook详情，PodName为接口返回值中Response.NotebookDetail.PodName
	// - Service类型为INFER：
	//   调用[DescribeModelService接口](/document/product/851/82287)查询单个服务详情，PodName为接口返回值中Response.Service.ServiceInfo.PodInfos
	// - Service类型为BATCH：
	//   调用[DescribeBatchTask接口](/document/product/851/80180)查询跑批任务详情，PodName为接口返回值中Response.BatchTaskDetail. PodList
	// 注：支持结尾通配符*
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// 排序方向（可选值为ASC, DESC ），默认为DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 按哪个字段排序（可选值为Timestamp），默认值为Timestamp
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 日志查询上下文，查询下一页的时候需要回传这个字段，该字段来自本接口的返回
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 过滤条件
	// 注意: 
	// 1. Filter.Name：目前只支持Key（也就是按关键字过滤日志）
	// 2. Filter.Values：表示过滤日志的关键字；Values为多个的时候表示同时满足
	// 3. Filter. Negative和Filter. Fuzzy没有使用
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeLogsRequest struct {
	*tchttp.BaseRequest
	
	// 服务类型，TRAIN为任务式建模, NOTEBOOK为Notebook, INFER为在线服务, BATCH为批量预测
	// 枚举值：
	// - TRAIN
	// - NOTEBOOK
	// - INFER
	// - BATCH
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 日志查询开始时间（RFC3339格式的时间字符串），默认值为当前时间的前一个小时
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 日志查询结束时间（RFC3339格式的时间字符串），默认值为当前时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 日志查询条数，默认值100，最大值100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 服务ID，和Service参数对应，不同Service的服务ID获取方式不同，具体如下：
	// - Service类型为TRAIN：
	//   调用[DescribeTrainingTask接口](/document/product/851/75089)查询训练任务详情，ServiceId为接口返回值中Response.TrainingTaskDetail.LatestInstanceId
	// - Service类型为NOTEBOOK：
	//   调用[DescribeNotebook接口](/document/product/851/95662)查询Notebook详情，ServiceId为接口返回值中Response.NotebookDetail.PodName
	// - Service类型为INFER：
	//   调用[DescribeModelServiceGroup接口](/document/product/851/82285)查询服务组详情，ServiceId为接口返回值中Response.ServiceGroup.Services.ServiceId
	// - Service类型为BATCH：
	//   调用[DescribeBatchTask接口](/document/product/851/80180)查询跑批任务详情，ServiceId为接口返回值中Response.BatchTaskDetail.LatestInstanceId
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// Pod的名称，即需要查询服务对应的Pod，和Service参数对应，不同Service的PodName获取方式不同，具体如下：
	// - Service类型为TRAIN：
	//   调用[DescribeTrainingTaskPods接口](/document/product/851/75088)查询训练任务pod列表，PodName为接口返回值中Response.PodNames
	// - Service类型为NOTEBOOK：
	//   调用[DescribeNotebook接口](/document/product/851/95662)查询Notebook详情，PodName为接口返回值中Response.NotebookDetail.PodName
	// - Service类型为INFER：
	//   调用[DescribeModelService接口](/document/product/851/82287)查询单个服务详情，PodName为接口返回值中Response.Service.ServiceInfo.PodInfos
	// - Service类型为BATCH：
	//   调用[DescribeBatchTask接口](/document/product/851/80180)查询跑批任务详情，PodName为接口返回值中Response.BatchTaskDetail. PodList
	// 注：支持结尾通配符*
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// 排序方向（可选值为ASC, DESC ），默认为DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 按哪个字段排序（可选值为Timestamp），默认值为Timestamp
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 日志查询上下文，查询下一页的时候需要回传这个字段，该字段来自本接口的返回
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 过滤条件
	// 注意: 
	// 1. Filter.Name：目前只支持Key（也就是按关键字过滤日志）
	// 2. Filter.Values：表示过滤日志的关键字；Values为多个的时候表示同时满足
	// 3. Filter. Negative和Filter. Fuzzy没有使用
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Service")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "ServiceId")
	delete(f, "PodName")
	delete(f, "Order")
	delete(f, "OrderField")
	delete(f, "Context")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogsResponseParams struct {
	// 分页的游标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 日志数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*LogIdentity `json:"Content,omitnil,omitempty" name:"Content"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogsResponseParams `json:"Response"`
}

func (r *DescribeLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelAccEngineVersionsRequestParams struct {

}

type DescribeModelAccEngineVersionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeModelAccEngineVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelAccEngineVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelAccEngineVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelAccEngineVersionsResponseParams struct {
	// 模型加速版本列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAccEngineVersions []*ModelAccEngineVersion `json:"ModelAccEngineVersions,omitnil,omitempty" name:"ModelAccEngineVersions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelAccEngineVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelAccEngineVersionsResponseParams `json:"Response"`
}

func (r *DescribeModelAccEngineVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelAccEngineVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelAccelerateTaskRequestParams struct {
	// 模型加速任务ID
	ModelAccTaskId *string `json:"ModelAccTaskId,omitnil,omitempty" name:"ModelAccTaskId"`
}

type DescribeModelAccelerateTaskRequest struct {
	*tchttp.BaseRequest
	
	// 模型加速任务ID
	ModelAccTaskId *string `json:"ModelAccTaskId,omitnil,omitempty" name:"ModelAccTaskId"`
}

func (r *DescribeModelAccelerateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelAccelerateTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelAccTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelAccelerateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelAccelerateTaskResponseParams struct {
	// 模型加速任务详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAccelerateTask *ModelAccelerateTask `json:"ModelAccelerateTask,omitnil,omitempty" name:"ModelAccelerateTask"`

	// 模型加速时长，单位s
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAccRuntimeInSecond *uint64 `json:"ModelAccRuntimeInSecond,omitnil,omitempty" name:"ModelAccRuntimeInSecond"`

	// 模型加速任务开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAccStartTime *string `json:"ModelAccStartTime,omitnil,omitempty" name:"ModelAccStartTime"`

	// 模型加速任务结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAccEndTime *string `json:"ModelAccEndTime,omitnil,omitempty" name:"ModelAccEndTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelAccelerateTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelAccelerateTaskResponseParams `json:"Response"`
}

func (r *DescribeModelAccelerateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelAccelerateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelAccelerateTasksRequestParams struct {
	// 过滤器
	// ModelAccTaskName 任务名称
	// ModelSource 模型来源
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段，默认CreateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序方式：ASC/DESC，默认DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 标签过滤
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

type DescribeModelAccelerateTasksRequest struct {
	*tchttp.BaseRequest
	
	// 过滤器
	// ModelAccTaskName 任务名称
	// ModelSource 模型来源
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段，默认CreateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序方式：ASC/DESC，默认DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回记录条数，默认10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 标签过滤
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

func (r *DescribeModelAccelerateTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelAccelerateTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "OrderField")
	delete(f, "Order")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelAccelerateTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelAccelerateTasksResponseParams struct {
	// 模型加速任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAccelerateTasks []*ModelAccelerateTask `json:"ModelAccelerateTasks,omitnil,omitempty" name:"ModelAccelerateTasks"`

	// 任务总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelAccelerateTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelAccelerateTasksResponseParams `json:"Response"`
}

func (r *DescribeModelAccelerateTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelAccelerateTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelAccelerateVersionsRequestParams struct {
	// 过滤条件
	//     Filter.Name: 枚举值: ModelJobName (任务名称)|TrainingModelVersionId (模型版本id)
	//     Filter.Values: 当长度为1时，支持模糊查询; 不为1时，精确查询
	// 每次请求的Filters的上限为10，Filter.Values的上限为100
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段; 枚举值: CreateTime (创建时间) ；默认CreateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序方向; 枚举值: ASC | DESC；默认DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 分页查询起始位置，如：Limit为100，第一页Offset为0，第二页Offset为100....即每页左边为闭区间; 默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询每页大小，最大20000; 默认10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模型ID
	TrainingModelId *string `json:"TrainingModelId,omitnil,omitempty" name:"TrainingModelId"`
}

type DescribeModelAccelerateVersionsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件
	//     Filter.Name: 枚举值: ModelJobName (任务名称)|TrainingModelVersionId (模型版本id)
	//     Filter.Values: 当长度为1时，支持模糊查询; 不为1时，精确查询
	// 每次请求的Filters的上限为10，Filter.Values的上限为100
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段; 枚举值: CreateTime (创建时间) ；默认CreateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序方向; 枚举值: ASC | DESC；默认DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 分页查询起始位置，如：Limit为100，第一页Offset为0，第二页Offset为100....即每页左边为闭区间; 默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询每页大小，最大20000; 默认10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模型ID
	TrainingModelId *string `json:"TrainingModelId,omitnil,omitempty" name:"TrainingModelId"`
}

func (r *DescribeModelAccelerateVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelAccelerateVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "OrderField")
	delete(f, "Order")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TrainingModelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelAccelerateVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelAccelerateVersionsResponseParams struct {
	// 优化模型总数； 注意接口是分页拉取的，total是指优化模型节点总数，不是本次返回中ModelAccelerateVersions数组的大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 优化模型列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAccelerateVersions []*ModelAccelerateVersion `json:"ModelAccelerateVersions,omitnil,omitempty" name:"ModelAccelerateVersions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelAccelerateVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelAccelerateVersionsResponseParams `json:"Response"`
}

func (r *DescribeModelAccelerateVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelAccelerateVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceCallInfoRequestParams struct {
	// 服务组id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 服务分类
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

type DescribeModelServiceCallInfoRequest struct {
	*tchttp.BaseRequest
	
	// 服务组id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 服务分类
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

func (r *DescribeModelServiceCallInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceCallInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceGroupId")
	delete(f, "ServiceCategory")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelServiceCallInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceCallInfoResponseParams struct {
	// 服务调用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceCallInfo *ServiceCallInfo `json:"ServiceCallInfo,omitnil,omitempty" name:"ServiceCallInfo"`

	// 升级网关调用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InferGatewayCallInfo *InferGatewayCallInfo `json:"InferGatewayCallInfo,omitnil,omitempty" name:"InferGatewayCallInfo"`

	// 默认nginx网关的调用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultNginxGatewayCallInfo *DefaultNginxGatewayCallInfo `json:"DefaultNginxGatewayCallInfo,omitnil,omitempty" name:"DefaultNginxGatewayCallInfo"`

	// 太极服务的调用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TJCallInfo *TJCallInfo `json:"TJCallInfo,omitnil,omitempty" name:"TJCallInfo"`

	// 内网调用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntranetCallInfo *IntranetCallInfo `json:"IntranetCallInfo,omitnil,omitempty" name:"IntranetCallInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelServiceCallInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelServiceCallInfoResponseParams `json:"Response"`
}

func (r *DescribeModelServiceCallInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceCallInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceGroupRequestParams struct {
	// 服务组ID
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 服务分类
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

type DescribeModelServiceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 服务组ID
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 服务分类
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

func (r *DescribeModelServiceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceGroupId")
	delete(f, "ServiceCategory")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelServiceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceGroupResponseParams struct {
	// 服务组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceGroup *ServiceGroup `json:"ServiceGroup,omitnil,omitempty" name:"ServiceGroup"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelServiceGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelServiceGroupResponseParams `json:"Response"`
}

func (r *DescribeModelServiceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceGroupsRequestParams struct {
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序的依据字段， 取值范围 "CreateTime" "UpdateTime"
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 分页参数，支持的分页过滤Name包括：
	// ["ClusterId", "ServiceId", "ServiceGroupName", "ServiceGroupId","Status","CreatedBy","ModelVersionId"]
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤参数
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 服务分类
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

type DescribeModelServiceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序的依据字段， 取值范围 "CreateTime" "UpdateTime"
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 分页参数，支持的分页过滤Name包括：
	// ["ClusterId", "ServiceId", "ServiceGroupName", "ServiceGroupId","Status","CreatedBy","ModelVersionId"]
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤参数
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 服务分类
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

func (r *DescribeModelServiceGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	delete(f, "Filters")
	delete(f, "TagFilters")
	delete(f, "ServiceCategory")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelServiceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceGroupsResponseParams struct {
	// 推理服务组数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 服务组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceGroups []*ServiceGroup `json:"ServiceGroups,omitnil,omitempty" name:"ServiceGroups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelServiceGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelServiceGroupsResponseParams `json:"Response"`
}

func (r *DescribeModelServiceGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceHistoryRequestParams struct {
	// 服务Id
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`
}

type DescribeModelServiceHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 服务Id
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`
}

func (r *DescribeModelServiceHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelServiceHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceHistoryResponseParams struct {
	// 历史版本总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 服务版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceHistory []*ServiceHistory `json:"ServiceHistory,omitnil,omitempty" name:"ServiceHistory"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelServiceHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelServiceHistoryResponseParams `json:"Response"`
}

func (r *DescribeModelServiceHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceHotUpdatedRequestParams struct {
	// 镜像信息，配置服务运行所需的镜像地址等信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 模型信息，需要挂载模型时填写
	ModelInfo *ModelInfo `json:"ModelInfo,omitnil,omitempty" name:"ModelInfo"`

	// 挂载信息
	VolumeMount *VolumeMount `json:"VolumeMount,omitnil,omitempty" name:"VolumeMount"`
}

type DescribeModelServiceHotUpdatedRequest struct {
	*tchttp.BaseRequest
	
	// 镜像信息，配置服务运行所需的镜像地址等信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 模型信息，需要挂载模型时填写
	ModelInfo *ModelInfo `json:"ModelInfo,omitnil,omitempty" name:"ModelInfo"`

	// 挂载信息
	VolumeMount *VolumeMount `json:"VolumeMount,omitnil,omitempty" name:"VolumeMount"`
}

func (r *DescribeModelServiceHotUpdatedRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceHotUpdatedRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageInfo")
	delete(f, "ModelInfo")
	delete(f, "VolumeMount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelServiceHotUpdatedRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceHotUpdatedResponseParams struct {
	// 模型加速标志位.Allowed 允许模型加速. Forbidden 禁止模型加速
	ModelTurboFlag *string `json:"ModelTurboFlag,omitnil,omitempty" name:"ModelTurboFlag"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelServiceHotUpdatedResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelServiceHotUpdatedResponseParams `json:"Response"`
}

func (r *DescribeModelServiceHotUpdatedResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceHotUpdatedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceRequestParams struct {
	// 服务id
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 服务分类
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

type DescribeModelServiceRequest struct {
	*tchttp.BaseRequest
	
	// 服务id
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 服务分类
	ServiceCategory *string `json:"ServiceCategory,omitnil,omitempty" name:"ServiceCategory"`
}

func (r *DescribeModelServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ServiceCategory")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServiceResponseParams struct {
	// 服务信息
	Service *Service `json:"Service,omitnil,omitempty" name:"Service"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelServiceResponseParams `json:"Response"`
}

func (r *DescribeModelServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServicesRequestParams struct {
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序的依据字段， 取值范围 "CreateTime" "UpdateTime"
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 分页参数，支持的分页过滤Name包括：
	// ["ClusterId", "ServiceId", "ServiceGroupName", "ServiceGroupId","Status","CreatedBy","ModelId"]
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤参数
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

type DescribeModelServicesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序的依据字段， 取值范围 "CreateTime" "UpdateTime"
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 分页参数，支持的分页过滤Name包括：
	// ["ClusterId", "ServiceId", "ServiceGroupName", "ServiceGroupId","Status","CreatedBy","ModelId"]
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤参数
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

func (r *DescribeModelServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	delete(f, "Filters")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelServicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelServicesResponseParams struct {
	// 服务数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Services []*Service `json:"Services,omitnil,omitempty" name:"Services"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelServicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelServicesResponseParams `json:"Response"`
}

func (r *DescribeModelServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookImageKernelsRequestParams struct {
	// notebook id
	NotebookId *string `json:"NotebookId,omitnil,omitempty" name:"NotebookId"`
}

type DescribeNotebookImageKernelsRequest struct {
	*tchttp.BaseRequest
	
	// notebook id
	NotebookId *string `json:"NotebookId,omitnil,omitempty" name:"NotebookId"`
}

func (r *DescribeNotebookImageKernelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookImageKernelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NotebookId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookImageKernelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookImageKernelsResponseParams struct {
	// 镜像kernel数组
	Kernels []*string `json:"Kernels,omitnil,omitempty" name:"Kernels"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNotebookImageKernelsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebookImageKernelsResponseParams `json:"Response"`
}

func (r *DescribeNotebookImageKernelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookImageKernelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookImageRecordsRequestParams struct {
	// notebook id
	NotebookId *string `json:"NotebookId,omitnil,omitempty" name:"NotebookId"`

	// 位移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 日志限制
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 状态筛选
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeNotebookImageRecordsRequest struct {
	*tchttp.BaseRequest
	
	// notebook id
	NotebookId *string `json:"NotebookId,omitnil,omitempty" name:"NotebookId"`

	// 位移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 日志限制
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 状态筛选
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeNotebookImageRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookImageRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NotebookId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookImageRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookImageRecordsResponseParams struct {
	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 镜像保存记录
	NotebookImageRecords []*NotebookImageRecord `json:"NotebookImageRecords,omitnil,omitempty" name:"NotebookImageRecords"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNotebookImageRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebookImageRecordsResponseParams `json:"Response"`
}

func (r *DescribeNotebookImageRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookImageRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookRequestParams struct {
	// notebook id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DescribeNotebookRequest struct {
	*tchttp.BaseRequest
	
	// notebook id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DescribeNotebookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookResponseParams struct {
	// 详情
	NotebookDetail *NotebookDetail `json:"NotebookDetail,omitnil,omitempty" name:"NotebookDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNotebookResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebookResponseParams `json:"Response"`
}

func (r *DescribeNotebookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebooksRequestParams struct {
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页返回的实例数，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列。默认为DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 根据哪个字段排序，如：CreateTime、UpdateTime，默认为UpdateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 过滤器，eg：[{ "Name": "Id", "Values": ["nb-123456789"] }]
	// 
	// 取值范围
	// Name（名称）：notebook1
	// Id（notebook ID）：nb-123456789
	// Status（状态）：Starting / Running / Stopped / Stopping / Failed / SubmitFailed
	// Creator（创建者 uin）：100014761913
	// ChargeType（计费类型）：PREPAID（预付费）/ POSTPAID_BY_HOUR（后付费）
	// ChargeStatus（计费状态）：NOT_BILLING（未开始计费）/ BILLING（计费中）/ BILLING_STORAGE（存储计费中）/ARREARS_STOP（欠费停止）
	// DefaultCodeRepoId（默认代码仓库ID）：cr-123456789
	// AdditionalCodeRepoId（关联代码仓库ID）：cr-123456789
	// LifecycleScriptId（生命周期ID）：ls-12312312311312
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤器，eg：[{ "TagKey": "TagKeyA", "TagValue": ["TagValueA"] }]
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

type DescribeNotebooksRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页返回的实例数，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC：升序排列 DESC：降序排列。默认为DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 根据哪个字段排序，如：CreateTime、UpdateTime，默认为UpdateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 过滤器，eg：[{ "Name": "Id", "Values": ["nb-123456789"] }]
	// 
	// 取值范围
	// Name（名称）：notebook1
	// Id（notebook ID）：nb-123456789
	// Status（状态）：Starting / Running / Stopped / Stopping / Failed / SubmitFailed
	// Creator（创建者 uin）：100014761913
	// ChargeType（计费类型）：PREPAID（预付费）/ POSTPAID_BY_HOUR（后付费）
	// ChargeStatus（计费状态）：NOT_BILLING（未开始计费）/ BILLING（计费中）/ BILLING_STORAGE（存储计费中）/ARREARS_STOP（欠费停止）
	// DefaultCodeRepoId（默认代码仓库ID）：cr-123456789
	// AdditionalCodeRepoId（关联代码仓库ID）：cr-123456789
	// LifecycleScriptId（生命周期ID）：ls-12312312311312
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤器，eg：[{ "TagKey": "TagKeyA", "TagValue": ["TagValueA"] }]
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

func (r *DescribeNotebooksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebooksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	delete(f, "Filters")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebooksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebooksResponseParams struct {
	// 详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotebookSet []*NotebookSetItem `json:"NotebookSet,omitnil,omitempty" name:"NotebookSet"`

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNotebooksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebooksResponseParams `json:"Response"`
}

func (r *DescribeNotebooksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebooksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingFrameworksRequestParams struct {

}

type DescribeTrainingFrameworksRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeTrainingFrameworksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingFrameworksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingFrameworksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingFrameworksResponseParams struct {
	// 框架信息列表
	FrameworkInfos []*FrameworkInfo `json:"FrameworkInfos,omitnil,omitempty" name:"FrameworkInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTrainingFrameworksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingFrameworksResponseParams `json:"Response"`
}

func (r *DescribeTrainingFrameworksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingFrameworksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingMetricsRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeTrainingMetricsRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeTrainingMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingMetricsResponseParams struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 训练指标数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*CustomTrainingData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTrainingMetricsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingMetricsResponseParams `json:"Response"`
}

func (r *DescribeTrainingMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingModelVersionRequestParams struct {
	// 模型版本ID
	TrainingModelVersionId *string `json:"TrainingModelVersionId,omitnil,omitempty" name:"TrainingModelVersionId"`
}

type DescribeTrainingModelVersionRequest struct {
	*tchttp.BaseRequest
	
	// 模型版本ID
	TrainingModelVersionId *string `json:"TrainingModelVersionId,omitnil,omitempty" name:"TrainingModelVersionId"`
}

func (r *DescribeTrainingModelVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingModelVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrainingModelVersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingModelVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingModelVersionResponseParams struct {
	// 模型版本
	TrainingModelVersion *TrainingModelVersionDTO `json:"TrainingModelVersion,omitnil,omitempty" name:"TrainingModelVersion"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTrainingModelVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingModelVersionResponseParams `json:"Response"`
}

func (r *DescribeTrainingModelVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingModelVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingModelVersionsRequestParams struct {
	// 模型ID
	TrainingModelId *string `json:"TrainingModelId,omitnil,omitempty" name:"TrainingModelId"`

	// 过滤条件
	// Filter.Name: 枚举值:
	//     TrainingModelVersionId (模型版本ID)
	//     ModelVersionType (模型版本类型) 其值支持: NORMAL(通用) ACCELERATE (加速)
	//     ModelFormat（模型格式）其值Filter.Values支持：
	// TORCH_SCRIPT/PYTORCH/DETECTRON2/SAVED_MODEL/FROZEN_GRAPH/PMML
	//     AlgorithmFramework (算法框架) 其值Filter.Values支持：TENSORFLOW/PYTORCH/DETECTRON2
	// Filter.Values: 当长度为1时，支持模糊查询; 不为1时，精确查询
	// 每次请求的Filters的上限为10，Filter.Values的上限为100
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTrainingModelVersionsRequest struct {
	*tchttp.BaseRequest
	
	// 模型ID
	TrainingModelId *string `json:"TrainingModelId,omitnil,omitempty" name:"TrainingModelId"`

	// 过滤条件
	// Filter.Name: 枚举值:
	//     TrainingModelVersionId (模型版本ID)
	//     ModelVersionType (模型版本类型) 其值支持: NORMAL(通用) ACCELERATE (加速)
	//     ModelFormat（模型格式）其值Filter.Values支持：
	// TORCH_SCRIPT/PYTORCH/DETECTRON2/SAVED_MODEL/FROZEN_GRAPH/PMML
	//     AlgorithmFramework (算法框架) 其值Filter.Values支持：TENSORFLOW/PYTORCH/DETECTRON2
	// Filter.Values: 当长度为1时，支持模糊查询; 不为1时，精确查询
	// 每次请求的Filters的上限为10，Filter.Values的上限为100
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeTrainingModelVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingModelVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrainingModelId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingModelVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingModelVersionsResponseParams struct {
	// 模型版本列表
	TrainingModelVersions []*TrainingModelVersionDTO `json:"TrainingModelVersions,omitnil,omitempty" name:"TrainingModelVersions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTrainingModelVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingModelVersionsResponseParams `json:"Response"`
}

func (r *DescribeTrainingModelVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingModelVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingModelsRequestParams struct {
	// 过滤器
	// Filter.Name: 枚举值:
	// keyword (模型名称)
	// TrainingModelId (模型ID)
	// ModelVersionType (模型版本类型) 其值Filter.Values支持: NORMAL(通用) ACCELERATE (加速)
	// TrainingModelSource (模型来源) 其值Filter.Values支持： JOB/COS
	// ModelFormat（模型格式）其值Filter.Values支持：
	// PYTORCH/TORCH_SCRIPT/DETECTRON2/SAVED_MODEL/FROZEN_GRAPH/PMML/MMDETECTION/ONNX/HUGGING_FACE
	// Filter.Values: 当长度为1时，支持模糊查询; 不为1时，精确查询
	// 每次请求的Filters的上限为10，Filter.Values的上限为100
	// Filter.Fuzzy取值：true/false，是否支持模糊匹配
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段，默认CreateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序方式，ASC/DESC，默认DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回结果数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 标签过滤
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 是否同时返回模型版本列表
	WithModelVersions *bool `json:"WithModelVersions,omitnil,omitempty" name:"WithModelVersions"`
}

type DescribeTrainingModelsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤器
	// Filter.Name: 枚举值:
	// keyword (模型名称)
	// TrainingModelId (模型ID)
	// ModelVersionType (模型版本类型) 其值Filter.Values支持: NORMAL(通用) ACCELERATE (加速)
	// TrainingModelSource (模型来源) 其值Filter.Values支持： JOB/COS
	// ModelFormat（模型格式）其值Filter.Values支持：
	// PYTORCH/TORCH_SCRIPT/DETECTRON2/SAVED_MODEL/FROZEN_GRAPH/PMML/MMDETECTION/ONNX/HUGGING_FACE
	// Filter.Values: 当长度为1时，支持模糊查询; 不为1时，精确查询
	// 每次请求的Filters的上限为10，Filter.Values的上限为100
	// Filter.Fuzzy取值：true/false，是否支持模糊匹配
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段，默认CreateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序方式，ASC/DESC，默认DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回结果数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 标签过滤
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 是否同时返回模型版本列表
	WithModelVersions *bool `json:"WithModelVersions,omitnil,omitempty" name:"WithModelVersions"`
}

func (r *DescribeTrainingModelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingModelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "OrderField")
	delete(f, "Order")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TagFilters")
	delete(f, "WithModelVersions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingModelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingModelsResponseParams struct {
	// 模型列表
	TrainingModels []*TrainingModelDTO `json:"TrainingModels,omitnil,omitempty" name:"TrainingModels"`

	// 模型总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTrainingModelsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingModelsResponseParams `json:"Response"`
}

func (r *DescribeTrainingModelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingModelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingTaskPodsRequestParams struct {
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DescribeTrainingTaskPodsRequest struct {
	*tchttp.BaseRequest
	
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DescribeTrainingTaskPodsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingTaskPodsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingTaskPodsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingTaskPodsResponseParams struct {
	// pod名称列表
	PodNames []*string `json:"PodNames,omitnil,omitempty" name:"PodNames"`

	// 数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// pod详细信息
	PodInfoList []*PodInfo `json:"PodInfoList,omitnil,omitempty" name:"PodInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTrainingTaskPodsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingTaskPodsResponseParams `json:"Response"`
}

func (r *DescribeTrainingTaskPodsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingTaskPodsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingTaskRequestParams struct {
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DescribeTrainingTaskRequest struct {
	*tchttp.BaseRequest
	
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DescribeTrainingTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingTaskResponseParams struct {
	// 训练任务详情
	TrainingTaskDetail *TrainingTaskDetail `json:"TrainingTaskDetail,omitnil,omitempty" name:"TrainingTaskDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTrainingTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingTaskResponseParams `json:"Response"`
}

func (r *DescribeTrainingTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingTasksRequestParams struct {
	// 过滤器，eg：[{ "Name": "Id", "Values": ["train-23091792777383936"] }]
	// 
	// 取值范围：
	// Name（名称）：task1
	// Id（task ID）：train-23091792777383936
	// Status（状态）：SUBMITTING/PENDING/STARTING / RUNNING / STOPPING / STOPPED / FAILED / SUCCEED / SUBMIT_FAILED
	// ResourceGroupId（资源组 Id）：trsg-kvvfrwl7
	// Creator（创建者 uin）：100014761913
	// ChargeType（计费类型）：PREPAID（预付费）/ POSTPAID_BY_HOUR（后付费）
	// CHARGE_STATUS（计费状态）：NOT_BILLING（未开始计费）/ BILLING（计费中）/ ARREARS_STOP（欠费停止）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤器，eg：[{ "TagKey": "TagKeyA", "TagValue": ["TagValueA"] }]
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为10，最大为50
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC（升序排列）/ DESC（降序排列），默认为DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序的依据字段， 取值范围 "CreateTime" 、"UpdateTime"、"StartTime"，默认为UpdateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeTrainingTasksRequest struct {
	*tchttp.BaseRequest
	
	// 过滤器，eg：[{ "Name": "Id", "Values": ["train-23091792777383936"] }]
	// 
	// 取值范围：
	// Name（名称）：task1
	// Id（task ID）：train-23091792777383936
	// Status（状态）：SUBMITTING/PENDING/STARTING / RUNNING / STOPPING / STOPPED / FAILED / SUCCEED / SUBMIT_FAILED
	// ResourceGroupId（资源组 Id）：trsg-kvvfrwl7
	// Creator（创建者 uin）：100014761913
	// ChargeType（计费类型）：PREPAID（预付费）/ POSTPAID_BY_HOUR（后付费）
	// CHARGE_STATUS（计费状态）：NOT_BILLING（未开始计费）/ BILLING（计费中）/ ARREARS_STOP（欠费停止）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤器，eg：[{ "TagKey": "TagKeyA", "TagValue": ["TagValueA"] }]
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为10，最大为50
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC（升序排列）/ DESC（降序排列），默认为DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序的依据字段， 取值范围 "CreateTime" 、"UpdateTime"、"StartTime"，默认为UpdateTime
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeTrainingTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "TagFilters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingTasksResponseParams struct {
	// 训练任务集
	TrainingTaskSet []*TrainingTaskSetItem `json:"TrainingTaskSet,omitnil,omitempty" name:"TrainingTaskSet"`

	// 数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTrainingTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingTasksResponseParams `json:"Response"`
}

func (r *DescribeTrainingTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetectionLabelInfo struct {
	// 点坐标列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Points []*PointInfo `json:"Points,omitnil,omitempty" name:"Points"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*string `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 类别
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameType *string `json:"FrameType,omitnil,omitempty" name:"FrameType"`
}

type EngineVersion struct {
	// 引擎版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 运行镜像
	// 注意：此字段可能返回 null，表示取不到有效值。
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 是否支持int8量化
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSupportIntEightQuantization *bool `json:"IsSupportIntEightQuantization,omitnil,omitempty" name:"IsSupportIntEightQuantization"`

	// 框架版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkVersion *string `json:"FrameworkVersion,omitnil,omitempty" name:"FrameworkVersion"`
}

type EnvVar struct {
	// 环境变量key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 环境变量value
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Event struct {
	// 事件的id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 事件的具体信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 事件第一次发生的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstTimestamp *string `json:"FirstTimestamp,omitnil,omitempty" name:"FirstTimestamp"`

	// 事件最后一次发生的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTimestamp *string `json:"LastTimestamp,omitnil,omitempty" name:"LastTimestamp"`

	// 事件发生的次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 事件的类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 事件关联的资源的类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceKind *string `json:"ResourceKind,omitnil,omitempty" name:"ResourceKind"`

	// 事件关联的资源的名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`
}

type Filter struct {
	// 过滤字段名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤字段取值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 是否开启反向查询
	Negative *bool `json:"Negative,omitnil,omitempty" name:"Negative"`

	// 是否开启模糊匹配
	Fuzzy *bool `json:"Fuzzy,omitnil,omitempty" name:"Fuzzy"`
}

type FilterLabelInfo struct {
	// 数据集id
	DatasetId *string `json:"DatasetId,omitnil,omitempty" name:"DatasetId"`

	// 文件ID
	FileId *string `json:"FileId,omitnil,omitempty" name:"FileId"`

	// 文件路径
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 分类标签结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassificationLabels []*string `json:"ClassificationLabels,omitnil,omitempty" name:"ClassificationLabels"`

	// 检测标签结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectionLabels []*DetectionLabelInfo `json:"DetectionLabels,omitnil,omitempty" name:"DetectionLabels"`

	// 分割标签结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentationLabels []*SegmentationInfo `json:"SegmentationLabels,omitnil,omitempty" name:"SegmentationLabels"`

	// RGB 图片路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	RGBPath *string `json:"RGBPath,omitnil,omitempty" name:"RGBPath"`

	// 标签模板类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelTemplateType *string `json:"LabelTemplateType,omitnil,omitempty" name:"LabelTemplateType"`

	// 下载url链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 缩略图下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadThumbnailUrl *string `json:"DownloadThumbnailUrl,omitnil,omitempty" name:"DownloadThumbnailUrl"`

	// 分割结果图片下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadRGBUrl *string `json:"DownloadRGBUrl,omitnil,omitempty" name:"DownloadRGBUrl"`

	// OCR场景
	// IDENTITY：识别
	// STRUCTURE：智能结构化
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrScene *string `json:"OcrScene,omitnil,omitempty" name:"OcrScene"`

	// OCR场景标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrLabels []*OcrLabelInfo `json:"OcrLabels,omitnil,omitempty" name:"OcrLabels"`

	// OCR场景标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrLabelInfo *string `json:"OcrLabelInfo,omitnil,omitempty" name:"OcrLabelInfo"`

	// 文本分类场景标签结果，内容是json结构
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextClassificationLabelList *string `json:"TextClassificationLabelList,omitnil,omitempty" name:"TextClassificationLabelList"`

	// 文本内容，返回50字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	RowText *string `json:"RowText,omitnil,omitempty" name:"RowText"`

	// 文本内容是否完全返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContentOmit *bool `json:"ContentOmit,omitnil,omitempty" name:"ContentOmit"`
}

type FrameworkInfo struct {
	// 框架名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 框架版本以及对应的训练模式
	VersionInfos []*FrameworkVersion `json:"VersionInfos,omitnil,omitempty" name:"VersionInfos"`
}

type FrameworkVersion struct {
	// 框架版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 训练模式
	TrainingModes []*string `json:"TrainingModes,omitnil,omitempty" name:"TrainingModes"`

	// 框架运行环境
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`
}

type GooseFS struct {
	// goosefs实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// GooseFS类型，包括GooseFS和GooseFSx
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// GooseFSx实例需要挂载的路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

type GpuDetail struct {
	// GPU 显卡类型；枚举值: V100 A100 T4
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// GPU 显卡数；单位为1/100卡，比如100代表1卡
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *uint64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type GroupResource struct {
	// CPU核数; 单位为1/1000核，比如100表示0.1核
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存；单位为MB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 总卡数；GPUDetail 显卡数之和；单位为1/100卡，比如100代表1卡
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gpu *uint64 `json:"Gpu,omitnil,omitempty" name:"Gpu"`

	// Gpu详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuDetailSet []*GpuDetail `json:"GpuDetailSet,omitnil,omitempty" name:"GpuDetailSet"`
}

type HDFSConfig struct {
	// 集群实例ID,实例ID形如: emr-xxxxxxxx
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

type HorizontalPodAutoscaler struct {
	// 最小实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinReplicas *int64 `json:"MinReplicas,omitnil,omitempty" name:"MinReplicas"`

	// 最大实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxReplicas *int64 `json:"MaxReplicas,omitnil,omitempty" name:"MaxReplicas"`

	// 支持：
	// "gpu-util": GPU利用率。范围{10, 100}      "cpu-util": CPU利用率。范围{10, 100}	      "memory-util": 内存利用率。范围{10, 100}      "service-qps": 单个实例QPS值。范围{1, 5000}
	// "concurrency-util":单个实例请求数量值。范围{1,100000}
	// 注意：此字段可能返回 null，表示取不到有效值。
	HpaMetrics []*Option `json:"HpaMetrics,omitnil,omitempty" name:"HpaMetrics"`
}

type HyperParameter struct {
	// 最大nnz数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxNNZ *string `json:"MaxNNZ,omitnil,omitempty" name:"MaxNNZ"`

	// slot数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlotNum *string `json:"SlotNum,omitnil,omitempty" name:"SlotNum"`

	// gpu cache 使用率
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuCachePercentage *string `json:"CpuCachePercentage,omitnil,omitempty" name:"CpuCachePercentage"`

	// cpu cache 使用率
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuCachePercentage *string `json:"GpuCachePercentage,omitnil,omitempty" name:"GpuCachePercentage"`

	// 是否开启分布式模式(true/false)
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableDistributed *string `json:"EnableDistributed,omitnil,omitempty" name:"EnableDistributed"`

	// TORCH_SCRIPT、MMDETECTION、DETECTRON2、HUGGINGFACE格式在进行优化时切分子图的最小算子数目，一般无需进行改动，默认为3
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinBlockSizePt *string `json:"MinBlockSizePt,omitnil,omitempty" name:"MinBlockSizePt"`

	// FROZEN_GRAPH、SAVED_MODEL格式在进行优化时切分子图的最小算子数目，一般无需进行改动，默认为10
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinBlockSizeTf *string `json:"MinBlockSizeTf,omitnil,omitempty" name:"MinBlockSizeTf"`

	// Stable Diffusion 模型优化参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PipelineArgs *string `json:"PipelineArgs,omitnil,omitempty" name:"PipelineArgs"`

	// Stable Diffusion 模型优化参数，控制Lora模型的影响效果
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoraScale *string `json:"LoraScale,omitnil,omitempty" name:"LoraScale"`
}

type ImageInfo struct {
	// 镜像类型：TCR为腾讯云TCR镜像; CCR为腾讯云TCR个人版镜像，PreSet为平台预置镜像
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// 镜像地址
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// TCR镜像对应的地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryRegion *string `json:"RegistryRegion,omitnil,omitempty" name:"RegistryRegion"`

	// TCR镜像对应的实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 是否允许导出全部内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowSaveAllContent *bool `json:"AllowSaveAllContent,omitnil,omitempty" name:"AllowSaveAllContent"`

	// 镜像名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`
}

type InferCodeInfo struct {
	// 推理代码所在的cos详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosPathInfo *CosPathInfo `json:"CosPathInfo,omitnil,omitempty" name:"CosPathInfo"`
}

type InferGatewayCallInfo struct {
	// 内网http调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcHttpAddr *string `json:"VpcHttpAddr,omitnil,omitempty" name:"VpcHttpAddr"`

	// 内网https调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcHttpsAddr *string `json:"VpcHttpsAddr,omitnil,omitempty" name:"VpcHttpsAddr"`

	// 内网grpc调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcGrpcTlsAddr *string `json:"VpcGrpcTlsAddr,omitnil,omitempty" name:"VpcGrpcTlsAddr"`

	// 可访问的vpcid
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 后端ip对应的子网
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

type InferTemplate struct {
	// 模板ID
	InferTemplateId *string `json:"InferTemplateId,omitnil,omitempty" name:"InferTemplateId"`

	// 模板镜像
	InferTemplateImage *string `json:"InferTemplateImage,omitnil,omitempty" name:"InferTemplateImage"`
}

type InferTemplateGroup struct {
	// 算法框架
	// 注意：此字段可能返回 null，表示取不到有效值。
	Framework *string `json:"Framework,omitnil,omitempty" name:"Framework"`

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkVersion *string `json:"FrameworkVersion,omitnil,omitempty" name:"FrameworkVersion"`

	// 支持的训练框架集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Groups []*string `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 镜像模板参数列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InferTemplates []*InferTemplate `json:"InferTemplates,omitnil,omitempty" name:"InferTemplates"`
}

type IngressPrivateLinkInfo struct {
	// 用户VpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 用户子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 内网http调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerHttpAddr []*string `json:"InnerHttpAddr,omitnil,omitempty" name:"InnerHttpAddr"`

	// 内网https调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerHttpsAddr []*string `json:"InnerHttpsAddr,omitnil,omitempty" name:"InnerHttpsAddr"`

	// 私有连接状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil,omitempty" name:"State"`
}

type Instance struct {
	// 资源组节点id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点已用资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedResource *ResourceInfo `json:"UsedResource,omitnil,omitempty" name:"UsedResource"`

	// 节点总资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalResource *ResourceInfo `json:"TotalResource,omitnil,omitempty" name:"TotalResource"`

	// 节点状态 
	// 注意：此字段为枚举值
	// 说明: 
	// DEPLOYING: 部署中
	// RUNNING: 运行中 
	// DEPLOY_FAILED: 部署失败
	//  RELEASING 释放中 
	// RELEASED：已释放 
	// EXCEPTION：异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 创建人
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// 创建时间: 
	// 注意：北京时间，比如: 2021-12-01 12:00:00
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 到期时间
	// 注意：北京时间，比如：2021-12-11 12:00:00
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 自动续费标识
	// 注意：此字段为枚举值
	// 说明：
	// NOTIFY_AND_MANUAL_RENEW：手动续费(取消自动续费)且到期通知
	// NOTIFY_AND_AUTO_RENEW：自动续费且到期通知
	// DISABLE_NOTIFY_AND_MANUAL_RENEW：手动续费(取消自动续费)且到期不通知
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlag *string `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 计费项ID
	SpecId *string `json:"SpecId,omitnil,omitempty" name:"SpecId"`

	// 计费项别名
	SpecAlias *string `json:"SpecAlias,omitnil,omitempty" name:"SpecAlias"`

	// 计费项特性列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecFeatures []*string `json:"SpecFeatures,omitnil,omitempty" name:"SpecFeatures"`

	// 纳管cvmid
	CvmInstanceId *string `json:"CvmInstanceId,omitnil,omitempty" name:"CvmInstanceId"`
}

type IntranetCallInfo struct {
	// 私有连接通道信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	IngressPrivateLinkInfo *IngressPrivateLinkInfo `json:"IngressPrivateLinkInfo,omitnil,omitempty" name:"IngressPrivateLinkInfo"`

	// 共享弹性网卡信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceEIPInfo []*ServiceEIPInfo `json:"ServiceEIPInfo,omitnil,omitempty" name:"ServiceEIPInfo"`

	// 私有连接信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateLinkInfos []*PrivateLinkInfo `json:"PrivateLinkInfos,omitnil,omitempty" name:"PrivateLinkInfos"`

	// 默认内网调用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultInnerCallInfos []*DefaultInnerCallInfo `json:"DefaultInnerCallInfos,omitnil,omitempty" name:"DefaultInnerCallInfos"`
}

type LogConfig struct {
	// 日志需要投递到cls的日志集
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 日志需要投递到cls的主题
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type LogIdentity struct {
	// 单条日志的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 单条日志的内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 这条日志对应的Pod名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// 日志的时间戳（RFC3339格式的时间字符串）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`
}

type Message struct {
	// 角色名。支持三个角色：system、user、assistant，其中system仅开头可出现一次，也可忽略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 对话输入内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type MetricData struct {
	// 训练任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 时间戳.unix timestamp,单位为秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 本次上报数据所处的训练周期数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Epoch *int64 `json:"Epoch,omitnil,omitempty" name:"Epoch"`

	// 本次上报数据所处的训练迭代次数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Step *int64 `json:"Step,omitnil,omitempty" name:"Step"`

	// 训练停止所需的迭代总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSteps *int64 `json:"TotalSteps,omitnil,omitempty" name:"TotalSteps"`

	// 数据点。数组元素为不同指标的数据。数组长度不超过10。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Points []*DataPoint `json:"Points,omitnil,omitempty" name:"Points"`
}

type ModelAccEngineVersion struct {
	// 模型格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelFormat *string `json:"ModelFormat,omitnil,omitempty" name:"ModelFormat"`

	// 引擎版本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineVersions []*EngineVersion `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`
}

type ModelAccelerateTask struct {
	// 模型加速任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAccTaskId *string `json:"ModelAccTaskId,omitnil,omitempty" name:"ModelAccTaskId"`

	// 模型加速任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAccTaskName *string `json:"ModelAccTaskName,omitnil,omitempty" name:"ModelAccTaskName"`

	// 模型ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 模型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 模型版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelVersion *string `json:"ModelVersion,omitnil,omitempty" name:"ModelVersion"`

	// 模型来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelSource *string `json:"ModelSource,omitnil,omitempty" name:"ModelSource"`

	// 优化级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	OptimizationLevel *string `json:"OptimizationLevel,omitnil,omitempty" name:"OptimizationLevel"`

	// 任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// input节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelInputNum *uint64 `json:"ModelInputNum,omitnil,omitempty" name:"ModelInputNum"`

	// input节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelInputInfos []*ModelInputInfo `json:"ModelInputInfos,omitnil,omitempty" name:"ModelInputInfos"`

	// GPU型号
	// 注意：此字段可能返回 null，表示取不到有效值。
	GPUType *string `json:"GPUType,omitnil,omitempty" name:"GPUType"`

	// 计费模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 加速比
	// 注意：此字段可能返回 null，表示取不到有效值。
	Speedup *string `json:"Speedup,omitnil,omitempty" name:"Speedup"`

	// 模型输入cos路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelInputPath *CosPathInfo `json:"ModelInputPath,omitnil,omitempty" name:"ModelInputPath"`

	// 模型输出cos路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelOutputPath *CosPathInfo `json:"ModelOutputPath,omitnil,omitempty" name:"ModelOutputPath"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 算法框架
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlgorithmFramework *string `json:"AlgorithmFramework,omitnil,omitempty" name:"AlgorithmFramework"`

	// 排队个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaitNumber *uint64 `json:"WaitNumber,omitnil,omitempty" name:"WaitNumber"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskProgress *uint64 `json:"TaskProgress,omitnil,omitempty" name:"TaskProgress"`

	// 模型格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelFormat *string `json:"ModelFormat,omitnil,omitempty" name:"ModelFormat"`

	// 模型Tensor信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TensorInfos []*string `json:"TensorInfos,omitnil,omitempty" name:"TensorInfos"`

	// 模型专业参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	HyperParameter *HyperParameter `json:"HyperParameter,omitnil,omitempty" name:"HyperParameter"`

	// 加速引擎版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccEngineVersion *string `json:"AccEngineVersion,omitnil,omitempty" name:"AccEngineVersion"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 优化模型是否已保存到模型仓库
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSaved *bool `json:"IsSaved,omitnil,omitempty" name:"IsSaved"`

	// SAVED_MODEL保存时配置的签名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelSignature *string `json:"ModelSignature,omitnil,omitempty" name:"ModelSignature"`

	// 是否是QAT模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	QATModel *bool `json:"QATModel,omitnil,omitempty" name:"QATModel"`

	// 加速引擎对应的框架版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkVersion *string `json:"FrameworkVersion,omitnil,omitempty" name:"FrameworkVersion"`
}

type ModelAccelerateVersion struct {
	// 模型id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 优化模型版本id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelVersionId *string `json:"ModelVersionId,omitnil,omitempty" name:"ModelVersionId"`

	// 优化任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelJobId *string `json:"ModelJobId,omitnil,omitempty" name:"ModelJobId"`

	// 优化任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelJobName *string `json:"ModelJobName,omitnil,omitempty" name:"ModelJobName"`

	// 优化后模型版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelVersion *string `json:"ModelVersion,omitnil,omitempty" name:"ModelVersion"`

	// 加速比
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpeedUp *string `json:"SpeedUp,omitnil,omitempty" name:"SpeedUp"`

	// 模型来源/任务名称/任务版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelSource *ModelSource `json:"ModelSource,omitnil,omitempty" name:"ModelSource"`

	// 模型cos路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosPathInfo *CosPathInfo `json:"CosPathInfo,omitnil,omitempty" name:"CosPathInfo"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模型规范
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelFormat *string `json:"ModelFormat,omitnil,omitempty" name:"ModelFormat"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *uint64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// GPU类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GPUType *string `json:"GPUType,omitnil,omitempty" name:"GPUType"`

	// 模型cos路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelCosPath *CosPathInfo `json:"ModelCosPath,omitnil,omitempty" name:"ModelCosPath"`
}

type ModelInfo struct {
	// 模型版本id, DescribeTrainingModelVersion查询模型接口时的id
	// 自动学习类型的模型填写自动学习的任务id
	ModelVersionId *string `json:"ModelVersionId,omitnil,omitempty" name:"ModelVersionId"`

	// 模型id
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 模型名
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 模型版本
	ModelVersion *string `json:"ModelVersion,omitnil,omitempty" name:"ModelVersion"`

	// 模型来源
	ModelSource *string `json:"ModelSource,omitnil,omitempty" name:"ModelSource"`

	// cos路径信息
	CosPathInfo *CosPathInfo `json:"CosPathInfo,omitnil,omitempty" name:"CosPathInfo"`

	// 模型对应的算法框架，预留
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlgorithmFramework *string `json:"AlgorithmFramework,omitnil,omitempty" name:"AlgorithmFramework"`

	// 默认为 NORMAL, 已加速模型: ACCELERATE, 自动学习模型 AUTO_ML
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelType *string `json:"ModelType,omitnil,omitempty" name:"ModelType"`

	// 模型格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelFormat *string `json:"ModelFormat,omitnil,omitempty" name:"ModelFormat"`

	// 是否为私有化大模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPrivateModel *bool `json:"IsPrivateModel,omitnil,omitempty" name:"IsPrivateModel"`
}

type ModelInputInfo struct {
	// input数据类型
	// FIXED：固定
	// RANGE：浮动
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelInputType *string `json:"ModelInputType,omitnil,omitempty" name:"ModelInputType"`

	// input数据尺寸
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelInputDimension []*string `json:"ModelInputDimension,omitnil,omitempty" name:"ModelInputDimension"`
}

type ModelSource struct {
	// 来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 来源任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// 来源任务版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobVersion *string `json:"JobVersion,omitnil,omitempty" name:"JobVersion"`

	// 来源任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 模型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 算法框架
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlgorithmFramework *string `json:"AlgorithmFramework,omitnil,omitempty" name:"AlgorithmFramework"`

	// 训练偏好
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingPreference *string `json:"TrainingPreference,omitnil,omitempty" name:"TrainingPreference"`

	// 推理环境来源，SYSTEM/CUSTOM
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReasoningEnvironmentSource *string `json:"ReasoningEnvironmentSource,omitnil,omitempty" name:"ReasoningEnvironmentSource"`

	// 推理环境
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReasoningEnvironment *string `json:"ReasoningEnvironment,omitnil,omitempty" name:"ReasoningEnvironment"`

	// 推理环境id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReasoningEnvironmentId *string `json:"ReasoningEnvironmentId,omitnil,omitempty" name:"ReasoningEnvironmentId"`

	// 自定义推理环境
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReasoningImageInfo *ImageInfo `json:"ReasoningImageInfo,omitnil,omitempty" name:"ReasoningImageInfo"`
}

// Predefined struct for user
type ModifyModelServicePartialConfigRequestParams struct {
	// 在线推理服务Id，需已存在
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 更新后服务不重启，定时停止的配置
	ScheduledAction *ScheduledAction `json:"ScheduledAction,omitnil,omitempty" name:"ScheduledAction"`

	// 更新后服务不重启，服务对应限流限频配置
	ServiceLimit *ServiceLimit `json:"ServiceLimit,omitnil,omitempty" name:"ServiceLimit"`
}

type ModifyModelServicePartialConfigRequest struct {
	*tchttp.BaseRequest
	
	// 在线推理服务Id，需已存在
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 更新后服务不重启，定时停止的配置
	ScheduledAction *ScheduledAction `json:"ScheduledAction,omitnil,omitempty" name:"ScheduledAction"`

	// 更新后服务不重启，服务对应限流限频配置
	ServiceLimit *ServiceLimit `json:"ServiceLimit,omitnil,omitempty" name:"ServiceLimit"`
}

func (r *ModifyModelServicePartialConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelServicePartialConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ScheduledAction")
	delete(f, "ServiceLimit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModelServicePartialConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelServicePartialConfigResponseParams struct {
	// 被修改后的服务配置
	Service *Service `json:"Service,omitnil,omitempty" name:"Service"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyModelServicePartialConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModelServicePartialConfigResponseParams `json:"Response"`
}

func (r *ModifyModelServicePartialConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelServicePartialConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelServiceRequestParams struct {
	// 服务id
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 模型信息，需要挂载模型时填写
	ModelInfo *ModelInfo `json:"ModelInfo,omitnil,omitempty" name:"ModelInfo"`

	// 镜像信息，配置服务运行所需的镜像地址等信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 环境变量，可选参数，用于配置容器中的环境变量
	Env []*EnvVar `json:"Env,omitnil,omitempty" name:"Env"`

	// 资源描述，指定预付费模式下的cpu,mem,gpu等信息，后付费无需填写
	Resources *ResourceInfo `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 使用DescribeBillingSpecs接口返回的规格列表中的值，或者参考实例列表:
	// TI.S.MEDIUM.POST	2C4G
	// TI.S.LARGE.POST	4C8G
	// TI.S.2XLARGE16.POST	8C16G
	// TI.S.2XLARGE32.POST	8C32G
	// TI.S.4XLARGE32.POST	16C32G
	// TI.S.4XLARGE64.POST	16C64G
	// TI.S.6XLARGE48.POST	24C48G
	// TI.S.6XLARGE96.POST	24C96G
	// TI.S.8XLARGE64.POST	32C64G
	// TI.S.8XLARGE128.POST 32C128G
	// TI.GN7.LARGE20.POST	4C20G T4*1/4
	// TI.GN7.2XLARGE40.POST	10C40G T4*1/2
	// TI.GN7.2XLARGE32.POST	8C32G T4*1
	// TI.GN7.5XLARGE80.POST	20C80G T4*1
	// TI.GN7.8XLARGE128.POST	32C128G T4*1
	// TI.GN7.10XLARGE160.POST	40C160G T4*2
	// TI.GN7.20XLARGE320.POST	80C320G T4*4
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 扩缩容类型 支持：自动 - "AUTO", 手动 - "MANUAL"
	ScaleMode *string `json:"ScaleMode,omitnil,omitempty" name:"ScaleMode"`

	// 实例数量, 不同计费模式和调节模式下对应关系如下
	// PREPAID 和 POSTPAID_BY_HOUR:
	// 手动调节模式下对应 实例数量
	// 自动调节模式下对应 基于时间的默认策略的实例数量
	// HYBRID_PAID:
	// 后付费实例手动调节模式下对应 实例数量
	// 后付费实例自动调节模式下对应 时间策略的默认策略的实例数量
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// 自动伸缩信息
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil,omitempty" name:"HorizontalPodAutoscaler"`

	// 是否开启日志投递，开启后需填写配置投递到指定cls
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// 日志配置，需要投递服务日志到指定cls时填写
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// 特殊更新行为： "STOP": 停止, "RESUME": 重启, "SCALE": 扩缩容, 存在这些特殊更新行为时，会忽略其他更新字段
	ServiceAction *string `json:"ServiceAction,omitnil,omitempty" name:"ServiceAction"`

	// 服务的描述
	ServiceDescription *string `json:"ServiceDescription,omitnil,omitempty" name:"ServiceDescription"`

	// 自动伸缩策略
	ScaleStrategy *string `json:"ScaleStrategy,omitnil,omitempty" name:"ScaleStrategy"`

	// 自动伸缩策略配置 HPA : 通过HPA进行弹性伸缩 CRON 通过定时任务进行伸缩
	CronScaleJobs []*CronScaleJob `json:"CronScaleJobs,omitnil,omitempty" name:"CronScaleJobs"`

	// 计费模式[HYBRID_PAID]时生效, 用于标识混合计费模式下的预付费实例数, 若不填则默认为1
	HybridBillingPrepaidReplicas *int64 `json:"HybridBillingPrepaidReplicas,omitnil,omitempty" name:"HybridBillingPrepaidReplicas"`

	// 是否开启模型的热更新。默认不开启
	ModelHotUpdateEnable *bool `json:"ModelHotUpdateEnable,omitnil,omitempty" name:"ModelHotUpdateEnable"`

	// 定时停止配置
	ScheduledAction *ScheduledAction `json:"ScheduledAction,omitnil,omitempty" name:"ScheduledAction"`

	// 服务限速限流相关配置
	ServiceLimit *ServiceLimit `json:"ServiceLimit,omitnil,omitempty" name:"ServiceLimit"`

	// 挂载配置，目前只支持CFS
	VolumeMount *VolumeMount `json:"VolumeMount,omitnil,omitempty" name:"VolumeMount"`

	// 是否开启模型的加速, 仅对StableDiffusion(动态加速)格式的模型有效。默认不开启
	ModelTurboEnable *bool `json:"ModelTurboEnable,omitnil,omitempty" name:"ModelTurboEnable"`

	// 服务的启动命令
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 是否开启TIONE内网访问外部
	ServiceEIP *ServiceEIP `json:"ServiceEIP,omitnil,omitempty" name:"ServiceEIP"`
}

type ModifyModelServiceRequest struct {
	*tchttp.BaseRequest
	
	// 服务id
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 模型信息，需要挂载模型时填写
	ModelInfo *ModelInfo `json:"ModelInfo,omitnil,omitempty" name:"ModelInfo"`

	// 镜像信息，配置服务运行所需的镜像地址等信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 环境变量，可选参数，用于配置容器中的环境变量
	Env []*EnvVar `json:"Env,omitnil,omitempty" name:"Env"`

	// 资源描述，指定预付费模式下的cpu,mem,gpu等信息，后付费无需填写
	Resources *ResourceInfo `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 使用DescribeBillingSpecs接口返回的规格列表中的值，或者参考实例列表:
	// TI.S.MEDIUM.POST	2C4G
	// TI.S.LARGE.POST	4C8G
	// TI.S.2XLARGE16.POST	8C16G
	// TI.S.2XLARGE32.POST	8C32G
	// TI.S.4XLARGE32.POST	16C32G
	// TI.S.4XLARGE64.POST	16C64G
	// TI.S.6XLARGE48.POST	24C48G
	// TI.S.6XLARGE96.POST	24C96G
	// TI.S.8XLARGE64.POST	32C64G
	// TI.S.8XLARGE128.POST 32C128G
	// TI.GN7.LARGE20.POST	4C20G T4*1/4
	// TI.GN7.2XLARGE40.POST	10C40G T4*1/2
	// TI.GN7.2XLARGE32.POST	8C32G T4*1
	// TI.GN7.5XLARGE80.POST	20C80G T4*1
	// TI.GN7.8XLARGE128.POST	32C128G T4*1
	// TI.GN7.10XLARGE160.POST	40C160G T4*2
	// TI.GN7.20XLARGE320.POST	80C320G T4*4
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 扩缩容类型 支持：自动 - "AUTO", 手动 - "MANUAL"
	ScaleMode *string `json:"ScaleMode,omitnil,omitempty" name:"ScaleMode"`

	// 实例数量, 不同计费模式和调节模式下对应关系如下
	// PREPAID 和 POSTPAID_BY_HOUR:
	// 手动调节模式下对应 实例数量
	// 自动调节模式下对应 基于时间的默认策略的实例数量
	// HYBRID_PAID:
	// 后付费实例手动调节模式下对应 实例数量
	// 后付费实例自动调节模式下对应 时间策略的默认策略的实例数量
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// 自动伸缩信息
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil,omitempty" name:"HorizontalPodAutoscaler"`

	// 是否开启日志投递，开启后需填写配置投递到指定cls
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// 日志配置，需要投递服务日志到指定cls时填写
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// 特殊更新行为： "STOP": 停止, "RESUME": 重启, "SCALE": 扩缩容, 存在这些特殊更新行为时，会忽略其他更新字段
	ServiceAction *string `json:"ServiceAction,omitnil,omitempty" name:"ServiceAction"`

	// 服务的描述
	ServiceDescription *string `json:"ServiceDescription,omitnil,omitempty" name:"ServiceDescription"`

	// 自动伸缩策略
	ScaleStrategy *string `json:"ScaleStrategy,omitnil,omitempty" name:"ScaleStrategy"`

	// 自动伸缩策略配置 HPA : 通过HPA进行弹性伸缩 CRON 通过定时任务进行伸缩
	CronScaleJobs []*CronScaleJob `json:"CronScaleJobs,omitnil,omitempty" name:"CronScaleJobs"`

	// 计费模式[HYBRID_PAID]时生效, 用于标识混合计费模式下的预付费实例数, 若不填则默认为1
	HybridBillingPrepaidReplicas *int64 `json:"HybridBillingPrepaidReplicas,omitnil,omitempty" name:"HybridBillingPrepaidReplicas"`

	// 是否开启模型的热更新。默认不开启
	ModelHotUpdateEnable *bool `json:"ModelHotUpdateEnable,omitnil,omitempty" name:"ModelHotUpdateEnable"`

	// 定时停止配置
	ScheduledAction *ScheduledAction `json:"ScheduledAction,omitnil,omitempty" name:"ScheduledAction"`

	// 服务限速限流相关配置
	ServiceLimit *ServiceLimit `json:"ServiceLimit,omitnil,omitempty" name:"ServiceLimit"`

	// 挂载配置，目前只支持CFS
	VolumeMount *VolumeMount `json:"VolumeMount,omitnil,omitempty" name:"VolumeMount"`

	// 是否开启模型的加速, 仅对StableDiffusion(动态加速)格式的模型有效。默认不开启
	ModelTurboEnable *bool `json:"ModelTurboEnable,omitnil,omitempty" name:"ModelTurboEnable"`

	// 服务的启动命令
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 是否开启TIONE内网访问外部
	ServiceEIP *ServiceEIP `json:"ServiceEIP,omitnil,omitempty" name:"ServiceEIP"`
}

func (r *ModifyModelServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ModelInfo")
	delete(f, "ImageInfo")
	delete(f, "Env")
	delete(f, "Resources")
	delete(f, "InstanceType")
	delete(f, "ScaleMode")
	delete(f, "Replicas")
	delete(f, "HorizontalPodAutoscaler")
	delete(f, "LogEnable")
	delete(f, "LogConfig")
	delete(f, "ServiceAction")
	delete(f, "ServiceDescription")
	delete(f, "ScaleStrategy")
	delete(f, "CronScaleJobs")
	delete(f, "HybridBillingPrepaidReplicas")
	delete(f, "ModelHotUpdateEnable")
	delete(f, "ScheduledAction")
	delete(f, "ServiceLimit")
	delete(f, "VolumeMount")
	delete(f, "ModelTurboEnable")
	delete(f, "Command")
	delete(f, "ServiceEIP")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModelServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelServiceResponseParams struct {
	// 生成的模型服务
	// 注意：此字段可能返回 null，表示取不到有效值。
	Service *Service `json:"Service,omitnil,omitempty" name:"Service"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyModelServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModelServiceResponseParams `json:"Response"`
}

func (r *ModifyModelServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNotebookRequestParams struct {
	// notebook id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 名称。不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// （不允许修改）计算资源付费模式 ，可选值为：
	// PREPAID：预付费，即包年包月
	// POSTPAID_BY_HOUR：按小时后付费
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 计算资源配置
	ResourceConf *ResourceConf `json:"ResourceConf,omitnil,omitempty" name:"ResourceConf"`

	// 是否上报日志
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// 是否自动停止
	AutoStopping *bool `json:"AutoStopping,omitnil,omitempty" name:"AutoStopping"`

	// 是否访问公网
	DirectInternetAccess *bool `json:"DirectInternetAccess,omitnil,omitempty" name:"DirectInternetAccess"`

	// 是否ROOT权限
	RootAccess *bool `json:"RootAccess,omitnil,omitempty" name:"RootAccess"`

	// 资源组ID(for预付费)
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// （不允许修改）Vpc-Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// （不允许修改）子网Id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 存储卷大小，单位GB
	VolumeSizeInGB *uint64 `json:"VolumeSizeInGB,omitnil,omitempty" name:"VolumeSizeInGB"`

	// （不允许修改）存储的类型。取值包含： 
	//     FREE:    预付费的免费存储
	//     CLOUD_PREMIUM： 高性能云硬盘
	//     CLOUD_SSD： SSD云硬盘
	//     CFS:     CFS存储，包含NFS和turbo
	VolumeSourceType *string `json:"VolumeSourceType,omitnil,omitempty" name:"VolumeSourceType"`

	// （不允许修改）CFS存储的配置
	VolumeSourceCFS *CFSConfig `json:"VolumeSourceCFS,omitnil,omitempty" name:"VolumeSourceCFS"`

	// 日志配置
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// 生命周期脚本的ID
	LifecycleScriptId *string `json:"LifecycleScriptId,omitnil,omitempty" name:"LifecycleScriptId"`

	// 默认GIT存储库的ID
	DefaultCodeRepoId *string `json:"DefaultCodeRepoId,omitnil,omitempty" name:"DefaultCodeRepoId"`

	// 其他GIT存储库的ID，最多3个
	AdditionalCodeRepoIds []*string `json:"AdditionalCodeRepoIds,omitnil,omitempty" name:"AdditionalCodeRepoIds"`

	// 自动停止时间，单位小时
	AutomaticStopTime *int64 `json:"AutomaticStopTime,omitnil,omitempty" name:"AutomaticStopTime"`

	// 标签配置
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 数据配置，只支持WEDATA_HDFS
	DataConfigs []*DataConfig `json:"DataConfigs,omitnil,omitempty" name:"DataConfigs"`

	// 镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 镜像类型，包括SYSTEM、TCR、CCR
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// SSH配置
	SSHConfig *SSHConfig `json:"SSHConfig,omitnil,omitempty" name:"SSHConfig"`
}

type ModifyNotebookRequest struct {
	*tchttp.BaseRequest
	
	// notebook id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 名称。不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// （不允许修改）计算资源付费模式 ，可选值为：
	// PREPAID：预付费，即包年包月
	// POSTPAID_BY_HOUR：按小时后付费
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 计算资源配置
	ResourceConf *ResourceConf `json:"ResourceConf,omitnil,omitempty" name:"ResourceConf"`

	// 是否上报日志
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// 是否自动停止
	AutoStopping *bool `json:"AutoStopping,omitnil,omitempty" name:"AutoStopping"`

	// 是否访问公网
	DirectInternetAccess *bool `json:"DirectInternetAccess,omitnil,omitempty" name:"DirectInternetAccess"`

	// 是否ROOT权限
	RootAccess *bool `json:"RootAccess,omitnil,omitempty" name:"RootAccess"`

	// 资源组ID(for预付费)
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// （不允许修改）Vpc-Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// （不允许修改）子网Id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 存储卷大小，单位GB
	VolumeSizeInGB *uint64 `json:"VolumeSizeInGB,omitnil,omitempty" name:"VolumeSizeInGB"`

	// （不允许修改）存储的类型。取值包含： 
	//     FREE:    预付费的免费存储
	//     CLOUD_PREMIUM： 高性能云硬盘
	//     CLOUD_SSD： SSD云硬盘
	//     CFS:     CFS存储，包含NFS和turbo
	VolumeSourceType *string `json:"VolumeSourceType,omitnil,omitempty" name:"VolumeSourceType"`

	// （不允许修改）CFS存储的配置
	VolumeSourceCFS *CFSConfig `json:"VolumeSourceCFS,omitnil,omitempty" name:"VolumeSourceCFS"`

	// 日志配置
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// 生命周期脚本的ID
	LifecycleScriptId *string `json:"LifecycleScriptId,omitnil,omitempty" name:"LifecycleScriptId"`

	// 默认GIT存储库的ID
	DefaultCodeRepoId *string `json:"DefaultCodeRepoId,omitnil,omitempty" name:"DefaultCodeRepoId"`

	// 其他GIT存储库的ID，最多3个
	AdditionalCodeRepoIds []*string `json:"AdditionalCodeRepoIds,omitnil,omitempty" name:"AdditionalCodeRepoIds"`

	// 自动停止时间，单位小时
	AutomaticStopTime *int64 `json:"AutomaticStopTime,omitnil,omitempty" name:"AutomaticStopTime"`

	// 标签配置
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 数据配置，只支持WEDATA_HDFS
	DataConfigs []*DataConfig `json:"DataConfigs,omitnil,omitempty" name:"DataConfigs"`

	// 镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 镜像类型，包括SYSTEM、TCR、CCR
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// SSH配置
	SSHConfig *SSHConfig `json:"SSHConfig,omitnil,omitempty" name:"SSHConfig"`
}

func (r *ModifyNotebookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNotebookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Name")
	delete(f, "ChargeType")
	delete(f, "ResourceConf")
	delete(f, "LogEnable")
	delete(f, "AutoStopping")
	delete(f, "DirectInternetAccess")
	delete(f, "RootAccess")
	delete(f, "ResourceGroupId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "VolumeSizeInGB")
	delete(f, "VolumeSourceType")
	delete(f, "VolumeSourceCFS")
	delete(f, "LogConfig")
	delete(f, "LifecycleScriptId")
	delete(f, "DefaultCodeRepoId")
	delete(f, "AdditionalCodeRepoIds")
	delete(f, "AutomaticStopTime")
	delete(f, "Tags")
	delete(f, "DataConfigs")
	delete(f, "ImageInfo")
	delete(f, "ImageType")
	delete(f, "SSHConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNotebookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNotebookResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNotebookResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNotebookResponseParams `json:"Response"`
}

func (r *ModifyNotebookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNotebookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNotebookTagsRequestParams struct {
	// Notebook Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Notebook修改标签集合
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ModifyNotebookTagsRequest struct {
	*tchttp.BaseRequest
	
	// Notebook Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Notebook修改标签集合
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *ModifyNotebookTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNotebookTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNotebookTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNotebookTagsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNotebookTagsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNotebookTagsResponseParams `json:"Response"`
}

func (r *ModifyNotebookTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNotebookTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceGroupWeightsRequestParams struct {
	// 服务组id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 权重设置
	Weights []*WeightEntry `json:"Weights,omitnil,omitempty" name:"Weights"`
}

type ModifyServiceGroupWeightsRequest struct {
	*tchttp.BaseRequest
	
	// 服务组id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 权重设置
	Weights []*WeightEntry `json:"Weights,omitnil,omitempty" name:"Weights"`
}

func (r *ModifyServiceGroupWeightsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceGroupWeightsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceGroupId")
	delete(f, "Weights")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyServiceGroupWeightsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceGroupWeightsResponseParams struct {
	// 更新权重后的服务组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceGroup *ServiceGroup `json:"ServiceGroup,omitnil,omitempty" name:"ServiceGroup"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyServiceGroupWeightsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyServiceGroupWeightsResponseParams `json:"Response"`
}

func (r *ModifyServiceGroupWeightsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceGroupWeightsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NotebookDetail struct {
	// notebook  ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// notebook 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 生命周期脚本
	// 注意：此字段可能返回 null，表示取不到有效值。
	LifecycleScriptId *string `json:"LifecycleScriptId,omitnil,omitempty" name:"LifecycleScriptId"`

	// Pod-Name
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// Update-Time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 是否访问公网
	DirectInternetAccess *bool `json:"DirectInternetAccess,omitnil,omitempty" name:"DirectInternetAccess"`

	// 预付费专用资源组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 标签配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否自动停止
	AutoStopping *bool `json:"AutoStopping,omitnil,omitempty" name:"AutoStopping"`

	// 其他GIT存储库，最多3个，单个
	// 长度不超过512字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdditionalCodeRepoIds []*string `json:"AdditionalCodeRepoIds,omitnil,omitempty" name:"AdditionalCodeRepoIds"`

	// 自动停止时间，单位小时
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutomaticStopTime *int64 `json:"AutomaticStopTime,omitnil,omitempty" name:"AutomaticStopTime"`

	// 资源配置
	ResourceConf *ResourceConf `json:"ResourceConf,omitnil,omitempty" name:"ResourceConf"`

	// 默认GIT存储库，长度不超过512字符
	DefaultCodeRepoId *string `json:"DefaultCodeRepoId,omitnil,omitempty" name:"DefaultCodeRepoId"`

	// 训练输出
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 是否上报日志
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// 日志配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// VPC ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 任务状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 运行时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeInSeconds *uint64 `json:"RuntimeInSeconds,omitnil,omitempty" name:"RuntimeInSeconds"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 训练开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 计费状态，eg：BILLING计费中，ARREARS_STOP欠费停止，NOT_BILLING不在计费中
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeStatus *string `json:"ChargeStatus,omitnil,omitempty" name:"ChargeStatus"`

	// 是否ROOT权限
	RootAccess *bool `json:"RootAccess,omitnil,omitempty" name:"RootAccess"`

	// 计贺金额信息，eg:2.00元/小时
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingInfos []*string `json:"BillingInfos,omitnil,omitempty" name:"BillingInfos"`

	// 存储卷大小 （单位时GB，最小10GB，必须是10G的倍数）
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeSizeInGB *uint64 `json:"VolumeSizeInGB,omitnil,omitempty" name:"VolumeSizeInGB"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureReason *string `json:"FailureReason,omitnil,omitempty" name:"FailureReason"`

	// 计算资源付费模式 (- PREPAID：预付费，即包年包月 - POSTPAID_BY_HOUR：按小时后付费)
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 后付费资源规格说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTypeAlias *string `json:"InstanceTypeAlias,omitnil,omitempty" name:"InstanceTypeAlias"`

	// 预付费资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// 存储的类型。取值包含： 
	//     FREE:        预付费的免费存储
	//     CLOUD_PREMIUM： 高性能云硬盘
	//     CLOUD_SSD： SSD云硬盘
	//     CFS:     CFS存储，包含NFS和turbo
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeSourceType *string `json:"VolumeSourceType,omitnil,omitempty" name:"VolumeSourceType"`

	// CFS存储的配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeSourceCFS *CFSConfig `json:"VolumeSourceCFS,omitnil,omitempty" name:"VolumeSourceCFS"`

	// 数据配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataConfigs []*DataConfig `json:"DataConfigs,omitnil,omitempty" name:"DataConfigs"`

	// notebook 信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 数据源来源，eg：WeData_HDFS
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSource *string `json:"DataSource,omitnil,omitempty" name:"DataSource"`

	// 镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 镜像类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// SSH配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	SSHConfig *SSHConfig `json:"SSHConfig,omitnil,omitempty" name:"SSHConfig"`

	// GooseFS存储配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeSourceGooseFS *GooseFS `json:"VolumeSourceGooseFS,omitnil,omitempty" name:"VolumeSourceGooseFS"`
}

type NotebookImageRecord struct {
	// 保存记录ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 镜像地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 状态。eg：creating导出中/success已完成/stopped已停止/fail异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// kernel数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kernels []*string `json:"Kernels,omitnil,omitempty" name:"Kernels"`
}

type NotebookSetItem struct {
	// notebook ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// notebook 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 计费模式
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 资源配置
	ResourceConf *ResourceConf `json:"ResourceConf,omitnil,omitempty" name:"ResourceConf"`

	// 预付费资源组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 存储卷大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeSizeInGB *uint64 `json:"VolumeSizeInGB,omitnil,omitempty" name:"VolumeSizeInGB"`

	// 计费金额信息，eg：2.00元/小时 (for后付费)
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingInfos []*string `json:"BillingInfos,omitnil,omitempty" name:"BillingInfos"`

	// 标签配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeInSeconds *uint64 `json:"RuntimeInSeconds,omitnil,omitempty" name:"RuntimeInSeconds"`

	// 计费状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeStatus *string `json:"ChargeStatus,omitnil,omitempty" name:"ChargeStatus"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureReason *string `json:"FailureReason,omitnil,omitempty" name:"FailureReason"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Pod名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// 后付费资源规格名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTypeAlias *string `json:"InstanceTypeAlias,omitnil,omitempty" name:"InstanceTypeAlias"`

	// 预付费资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// 是否自动终止
	AutoStopping *bool `json:"AutoStopping,omitnil,omitempty" name:"AutoStopping"`

	// 自动停止时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutomaticStopTime *uint64 `json:"AutomaticStopTime,omitnil,omitempty" name:"AutomaticStopTime"`

	// 存储的类型。取值包含： 
	//     FREE:        预付费的免费存储
	//     CLOUD_PREMIUM： 高性能云硬盘
	//     CLOUD_SSD： SSD云硬盘
	//     CFS:     CFS存储，包含NFS和turbo
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeSourceType *string `json:"VolumeSourceType,omitnil,omitempty" name:"VolumeSourceType"`

	// CFS存储的配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeSourceCFS *CFSConfig `json:"VolumeSourceCFS,omitnil,omitempty" name:"VolumeSourceCFS"`

	// notebook 信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// notebook用户类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserTypes []*string `json:"UserTypes,omitnil,omitempty" name:"UserTypes"`

	// SSH配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	SSHConfig *SSHConfig `json:"SSHConfig,omitnil,omitempty" name:"SSHConfig"`

	// GooseFS存储配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeSourceGooseFS *GooseFS `json:"VolumeSourceGooseFS,omitnil,omitempty" name:"VolumeSourceGooseFS"`
}

type OcrLabelInfo struct {
	// 坐标点围起来的框
	// 注意：此字段可能返回 null，表示取不到有效值。
	Points []*PointInfo `json:"Points,omitnil,omitempty" name:"Points"`

	// 框的形状：
	// FRAME_TYPE_RECTANGLE
	// FRAME_TYPE_POLYGON
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameType *string `json:"FrameType,omitnil,omitempty" name:"FrameType"`

	// 智能结构化：key区域对应的内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 智能结构化：上述key的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// 识别：框区域的内容
	// 智能结构化：value区域对应的内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 智能结构化：value区域所关联的key 区域的keyID的集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyIdsForValue []*string `json:"KeyIdsForValue,omitnil,omitempty" name:"KeyIdsForValue"`

	// key或者value区域内容的方向：
	// DIRECTION_VERTICAL
	// DIRECTION_HORIZONTAL
	// 注意：此字段可能返回 null，表示取不到有效值。
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type Option struct {
	// 指标名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 指标值
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type Pod struct {
	// pod名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// pod的唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 服务付费模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// pod的状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phase *string `json:"Phase,omitnil,omitempty" name:"Phase"`

	// pod的IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// pod的创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 容器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Containers *Container `json:"Containers,omitnil,omitempty" name:"Containers"`

	// 容器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerInfos []*Container `json:"ContainerInfos,omitnil,omitempty" name:"ContainerInfos"`

	// 容器调用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrossTenantENIInfo *CrossTenantENIInfo `json:"CrossTenantENIInfo,omitnil,omitempty" name:"CrossTenantENIInfo"`
}

type PodInfo struct {
	// pod名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// pod的IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// pod状态。eg：SUBMITTING提交中、PENDING排队中、RUNNING运行中、SUCCEEDED已完成、FAILED异常、TERMINATING停止中、TERMINATED已停止
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// pod启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// pod结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// pod资源配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceConfigInfo *ResourceConfigInfo `json:"ResourceConfigInfo,omitnil,omitempty" name:"ResourceConfigInfo"`
}

type PointInfo struct {
	// X坐标值
	// 注意：此字段可能返回 null，表示取不到有效值。
	X *float64 `json:"X,omitnil,omitempty" name:"X"`

	// Y坐标值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Y *float64 `json:"Y,omitnil,omitempty" name:"Y"`
}

type PreTrainModel struct {
	// 模型ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 模型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`
}

type PrivateLinkInfo struct {
	// 私有连接所在的VPCID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有连接所在的子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// HTTP内网调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerHttpAddr []*string `json:"InnerHttpAddr,omitnil,omitempty" name:"InnerHttpAddr"`

	// HTTPS内网调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerHttpsAddr []*string `json:"InnerHttpsAddr,omitnil,omitempty" name:"InnerHttpsAddr"`

	// 私有连接状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil,omitempty" name:"State"`
}

// Predefined struct for user
type PushTrainingMetricsRequestParams struct {
	// 指标数据
	Data []*MetricData `json:"Data,omitnil,omitempty" name:"Data"`
}

type PushTrainingMetricsRequest struct {
	*tchttp.BaseRequest
	
	// 指标数据
	Data []*MetricData `json:"Data,omitnil,omitempty" name:"Data"`
}

func (r *PushTrainingMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PushTrainingMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PushTrainingMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PushTrainingMetricsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PushTrainingMetricsResponse struct {
	*tchttp.BaseResponse
	Response *PushTrainingMetricsResponseParams `json:"Response"`
}

func (r *PushTrainingMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PushTrainingMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RDMAConfig struct {
	// 是否开启RDMA
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type ResourceConf struct {
	// cpu 处理器资源, 单位为1/1000核 (for预付费)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// memory 内存资源, 单位为1M (for预付费)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// gpu Gpu卡资源，单位为1/100卡，例如GpuType=T4时，1 Gpu = 1/100 T4卡 (for预付费)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gpu *uint64 `json:"Gpu,omitnil,omitempty" name:"Gpu"`

	// GpuType 卡类型，参考资源组上对应的卡类型。eg: H800,A800,A100,T4,P4,V100等 (for预付费)
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuType *string `json:"GpuType,omitnil,omitempty" name:"GpuType"`

	// 计算规格 (for后付费)，可选值如下：
	// TI.S.LARGE.POST: 4C8G 
	// TI.S.2XLARGE16.POST:  8C16G 
	// TI.S.2XLARGE32.POST:  8C32G 
	// TI.S.4XLARGE32.POST:  16C32G
	// TI.S.4XLARGE64.POST:  16C64G
	// TI.S.6XLARGE48.POST:  24C48G
	// TI.S.6XLARGE96.POST:  24C96G
	// TI.S.8XLARGE64.POST:  32C64G
	// TI.S.8XLARGE128.POST : 32C128G
	// TI.GN10.2XLARGE40.POST: 8C40G V100*1 
	// TI.GN10.5XLARGE80.POST:  18C80G V100*2 
	// TI.GN10.10XLARGE160.POST :  32C160G V100*4
	// TI.GN10.20XLARGE320.POST :  72C320G V100*8
	// TI.GN7.8XLARGE128.POST: 32C128G T4*1 
	// TI.GN7.10XLARGE160.POST: 40C160G T4*2 
	// TI.GN7.20XLARGE320.POST: 80C320G T4*4
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

type ResourceConfigInfo struct {
	// 角色，eg：PS、WORKER、DRIVER、EXECUTOR
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// cpu核数，使用资源组时需配置。单位：1/1000，即1000=1核
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存，使用资源组时需配置。单位为MB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// gpu卡类型，使用资源组时需配置
	GpuType *string `json:"GpuType,omitnil,omitempty" name:"GpuType"`

	// gpu卡数，使用资源组时需配置。单位：1/100，即100=1卡
	Gpu *uint64 `json:"Gpu,omitnil,omitempty" name:"Gpu"`

	// 算力规格ID
	// 计算规格 (for后付费)，可选值如下：
	// TI.S.LARGE.POST: 4C8G 
	// TI.S.2XLARGE16.POST:  8C16G 
	// TI.S.2XLARGE32.POST:  8C32G 
	// TI.S.4XLARGE32.POST:  16C32G
	// TI.S.4XLARGE64.POST:  16C64G
	// TI.S.6XLARGE48.POST:  24C48G
	// TI.S.6XLARGE96.POST:  24C96G
	// TI.S.8XLARGE64.POST:  32C64G
	// TI.S.8XLARGE128.POST : 32C128G
	// TI.GN10.2XLARGE40.POST: 8C40G V100*1 
	// TI.GN10.5XLARGE80.POST:  18C80G V100*2 
	// TI.GN10.10XLARGE160.POST :  32C160G V100*4
	// TI.GN10.20XLARGE320.POST :  72C320G V100*8
	// TI.GN7.8XLARGE128.POST: 32C128G T4*1 
	// TI.GN7.10XLARGE160.POST: 40C160G T4*2 
	// TI.GN7.20XLARGE320.POST: 80C32
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 计算节点数
	InstanceNum *uint64 `json:"InstanceNum,omitnil,omitempty" name:"InstanceNum"`

	// 算力规格名称
	// 计算规格 (for后付费)，可选值如下：
	// 4C8G 
	// 8C16G 
	// 8C32G 
	// 16C32G
	// 6C64G
	// 24C48G
	// 24C96G
	// 32C64G
	// 32C128G
	// 8C40G V100*1 
	// 8C80G V100*2 
	// 32C160G V100*4
	// 72C320G V100*8
	// 32C128G T4*1 
	// 40C160G T4*2 
	// 80C32
	InstanceTypeAlias *string `json:"InstanceTypeAlias,omitnil,omitempty" name:"InstanceTypeAlias"`

	// RDMA配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	RDMAConfig *RDMAConfig `json:"RDMAConfig,omitnil,omitempty" name:"RDMAConfig"`
}

type ResourceGroup struct {
	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 资源组名称
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// 可用节点个数(运行中的节点)
	FreeInstance *uint64 `json:"FreeInstance,omitnil,omitempty" name:"FreeInstance"`

	// 总节点个数(所有节点)
	TotalInstance *uint64 `json:"TotalInstance,omitnil,omitempty" name:"TotalInstance"`

	// 资资源组已用的资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedResource *GroupResource `json:"UsedResource,omitnil,omitempty" name:"UsedResource"`

	// 资源组总资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalResource *GroupResource `json:"TotalResource,omitnil,omitempty" name:"TotalResource"`

	// 节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceSet []*Instance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*Tag `json:"TagSet,omitnil,omitempty" name:"TagSet"`
}

type ResourceInfo struct {
	// 处理器资源, 单位为1/1000核
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存资源, 单位为1M
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Gpu卡个数资源, 单位为0.01单位的GpuType.
	// Gpu=100表示使用了“一张”gpu卡, 但此处的“一张”卡有可能是虚拟化后的1/4卡, 也有可能是整张卡. 取决于实例的机型
	// 例1 实例的机型带有1张虚拟gpu卡, 每张虚拟gpu卡对应1/4张实际T4卡, 则此时 GpuType=T4, Gpu=100, RealGpu=25.
	// 例2 实例的机型带有4张gpu整卡, 每张卡对应1张实际T4卡, 则 此时 GpuType=T4, Gpu=400, RealGpu=400.
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gpu *uint64 `json:"Gpu,omitnil,omitempty" name:"Gpu"`

	// Gpu卡型号 T4或者V100。仅展示当前 GPU 卡型号，若存在多类型同时使用，则参考 RealGpuDetailSet 的值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuType *string `json:"GpuType,omitnil,omitempty" name:"GpuType"`

	// 创建或更新时无需填写，仅展示需要关注
	// 后付费非整卡实例对应的实际的Gpu卡资源, 表示gpu资源对应实际的gpu卡个数.
	// RealGpu=100表示实际使用了一张gpu卡, 对应实际的实例机型, 有可能代表带有1/4卡的实例4个, 或者带有1/2卡的实例2个, 或者带有1卡的实力1个.
	RealGpu *uint64 `json:"RealGpu,omitnil,omitempty" name:"RealGpu"`

	// 创建或更新时无需填写，仅展示需要关注。详细的GPU使用信息。
	RealGpuDetailSet []*GpuDetail `json:"RealGpuDetailSet,omitnil,omitempty" name:"RealGpuDetailSet"`
}

type ResourceInstanceRunningJobInfo struct {
	// pod名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务自定义名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`
}

// Predefined struct for user
type RestartModelAccelerateTaskRequestParams struct {
	// 模型加速任务ID
	ModelAccTaskId *string `json:"ModelAccTaskId,omitnil,omitempty" name:"ModelAccTaskId"`

	// 模型加速任务名称
	ModelAccTaskName *string `json:"ModelAccTaskName,omitnil,omitempty" name:"ModelAccTaskName"`

	// 模型来源（JOB/COS）
	ModelSource *string `json:"ModelSource,omitnil,omitempty" name:"ModelSource"`

	// 算法框架（废弃）
	AlgorithmFramework *string `json:"AlgorithmFramework,omitnil,omitempty" name:"AlgorithmFramework"`

	// 模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 模型版本
	ModelVersion *string `json:"ModelVersion,omitnil,omitempty" name:"ModelVersion"`

	// 模型输入cos路径
	ModelInputPath *CosPathInfo `json:"ModelInputPath,omitnil,omitempty" name:"ModelInputPath"`

	// 优化级别（NO_LOSS/FP16/INT8），默认FP16
	OptimizationLevel *string `json:"OptimizationLevel,omitnil,omitempty" name:"OptimizationLevel"`

	// input节点个数（废弃）
	ModelInputNum *uint64 `json:"ModelInputNum,omitnil,omitempty" name:"ModelInputNum"`

	// input节点信息（废弃）
	ModelInputInfos []*ModelInputInfo `json:"ModelInputInfos,omitnil,omitempty" name:"ModelInputInfos"`

	// 模型输出cos路径
	ModelOutputPath *CosPathInfo `json:"ModelOutputPath,omitnil,omitempty" name:"ModelOutputPath"`

	// 模型格式（TORCH_SCRIPT/DETECTRON2/SAVED_MODEL/FROZEN_GRAPH/MMDETECTION/ONNX/HUGGING_FACE）
	ModelFormat *string `json:"ModelFormat,omitnil,omitempty" name:"ModelFormat"`

	// 模型Tensor信息
	TensorInfos []*string `json:"TensorInfos,omitnil,omitempty" name:"TensorInfos"`

	// GPU类型（T4/V100/A10），默认T4
	GPUType *string `json:"GPUType,omitnil,omitempty" name:"GPUType"`

	// 模型专业参数
	HyperParameter *HyperParameter `json:"HyperParameter,omitnil,omitempty" name:"HyperParameter"`

	// 加速引擎版本
	AccEngineVersion *string `json:"AccEngineVersion,omitnil,omitempty" name:"AccEngineVersion"`

	// 标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// SavedModel保存时配置的签名
	ModelSignature *string `json:"ModelSignature,omitnil,omitempty" name:"ModelSignature"`

	// 加速引擎对应的框架版本
	FrameworkVersion *string `json:"FrameworkVersion,omitnil,omitempty" name:"FrameworkVersion"`
}

type RestartModelAccelerateTaskRequest struct {
	*tchttp.BaseRequest
	
	// 模型加速任务ID
	ModelAccTaskId *string `json:"ModelAccTaskId,omitnil,omitempty" name:"ModelAccTaskId"`

	// 模型加速任务名称
	ModelAccTaskName *string `json:"ModelAccTaskName,omitnil,omitempty" name:"ModelAccTaskName"`

	// 模型来源（JOB/COS）
	ModelSource *string `json:"ModelSource,omitnil,omitempty" name:"ModelSource"`

	// 算法框架（废弃）
	AlgorithmFramework *string `json:"AlgorithmFramework,omitnil,omitempty" name:"AlgorithmFramework"`

	// 模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 模型版本
	ModelVersion *string `json:"ModelVersion,omitnil,omitempty" name:"ModelVersion"`

	// 模型输入cos路径
	ModelInputPath *CosPathInfo `json:"ModelInputPath,omitnil,omitempty" name:"ModelInputPath"`

	// 优化级别（NO_LOSS/FP16/INT8），默认FP16
	OptimizationLevel *string `json:"OptimizationLevel,omitnil,omitempty" name:"OptimizationLevel"`

	// input节点个数（废弃）
	ModelInputNum *uint64 `json:"ModelInputNum,omitnil,omitempty" name:"ModelInputNum"`

	// input节点信息（废弃）
	ModelInputInfos []*ModelInputInfo `json:"ModelInputInfos,omitnil,omitempty" name:"ModelInputInfos"`

	// 模型输出cos路径
	ModelOutputPath *CosPathInfo `json:"ModelOutputPath,omitnil,omitempty" name:"ModelOutputPath"`

	// 模型格式（TORCH_SCRIPT/DETECTRON2/SAVED_MODEL/FROZEN_GRAPH/MMDETECTION/ONNX/HUGGING_FACE）
	ModelFormat *string `json:"ModelFormat,omitnil,omitempty" name:"ModelFormat"`

	// 模型Tensor信息
	TensorInfos []*string `json:"TensorInfos,omitnil,omitempty" name:"TensorInfos"`

	// GPU类型（T4/V100/A10），默认T4
	GPUType *string `json:"GPUType,omitnil,omitempty" name:"GPUType"`

	// 模型专业参数
	HyperParameter *HyperParameter `json:"HyperParameter,omitnil,omitempty" name:"HyperParameter"`

	// 加速引擎版本
	AccEngineVersion *string `json:"AccEngineVersion,omitnil,omitempty" name:"AccEngineVersion"`

	// 标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// SavedModel保存时配置的签名
	ModelSignature *string `json:"ModelSignature,omitnil,omitempty" name:"ModelSignature"`

	// 加速引擎对应的框架版本
	FrameworkVersion *string `json:"FrameworkVersion,omitnil,omitempty" name:"FrameworkVersion"`
}

func (r *RestartModelAccelerateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartModelAccelerateTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelAccTaskId")
	delete(f, "ModelAccTaskName")
	delete(f, "ModelSource")
	delete(f, "AlgorithmFramework")
	delete(f, "ModelId")
	delete(f, "ModelName")
	delete(f, "ModelVersion")
	delete(f, "ModelInputPath")
	delete(f, "OptimizationLevel")
	delete(f, "ModelInputNum")
	delete(f, "ModelInputInfos")
	delete(f, "ModelOutputPath")
	delete(f, "ModelFormat")
	delete(f, "TensorInfos")
	delete(f, "GPUType")
	delete(f, "HyperParameter")
	delete(f, "AccEngineVersion")
	delete(f, "Tags")
	delete(f, "ModelSignature")
	delete(f, "FrameworkVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartModelAccelerateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartModelAccelerateTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartModelAccelerateTaskResponse struct {
	*tchttp.BaseResponse
	Response *RestartModelAccelerateTaskResponseParams `json:"Response"`
}

func (r *RestartModelAccelerateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartModelAccelerateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RowItem struct {
	// rowValue 数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*RowValue `json:"Values,omitnil,omitempty" name:"Values"`
}

type RowValue struct {
	// 列名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 列值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type SSHConfig struct {
	// 是否开启ssh
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 公钥信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicKey *string `json:"PublicKey,omitnil,omitempty" name:"PublicKey"`

	// 端口号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 登录命令
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoginCommand *string `json:"LoginCommand,omitnil,omitempty" name:"LoginCommand"`
}

type ScheduledAction struct {
	// 是否要定时停止服务，true or false。true 则 ScheduleStopTime 必填， false 则 ScheduleStopTime 不生效
	ScheduleStop *bool `json:"ScheduleStop,omitnil,omitempty" name:"ScheduleStop"`

	// 要执行定时停止的时间，格式：“2022-01-26T19:46:22+08:00”
	ScheduleStopTime *string `json:"ScheduleStopTime,omitnil,omitempty" name:"ScheduleStopTime"`
}

type SchemaInfo struct {
	// 长度30字符内
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type SegmentationInfo struct {
	// 点坐标数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Points []*PointInfo `json:"Points,omitnil,omitempty" name:"Points"`

	// 分割标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 灰度值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gray *uint64 `json:"Gray,omitnil,omitempty" name:"Gray"`

	// 颜色
	// 注意：此字段可能返回 null，表示取不到有效值。
	Color *string `json:"Color,omitnil,omitempty" name:"Color"`
}

// Predefined struct for user
type SendChatMessageRequestParams struct {
	// 会话id，标识一组对话的唯一id，id变更则重置会话
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 问题描述
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 会话模型版本。
	// 金融大模型：填写sn-finllm-13b-chat-v1。
	// 默认为sn-finllm-13b-chat-v1，即金融大模型。
	ModelVersion *string `json:"ModelVersion,omitnil,omitempty" name:"ModelVersion"`

	// 使用模式。
	// 通用问答：填写General。
	// 搜索增强问答：填写WithSearchPlugin。
	// 默认为General，即通用问答。
	// 当前可体验模型仅支持General。
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 搜索来源。仅当Mode为WithSearchPlugin时生效。
	// 预置文稿库：填写Preset。自定义：填写Custom。
	SearchSource *string `json:"SearchSource,omitnil,omitempty" name:"SearchSource"`
}

type SendChatMessageRequest struct {
	*tchttp.BaseRequest
	
	// 会话id，标识一组对话的唯一id，id变更则重置会话
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 问题描述
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 会话模型版本。
	// 金融大模型：填写sn-finllm-13b-chat-v1。
	// 默认为sn-finllm-13b-chat-v1，即金融大模型。
	ModelVersion *string `json:"ModelVersion,omitnil,omitempty" name:"ModelVersion"`

	// 使用模式。
	// 通用问答：填写General。
	// 搜索增强问答：填写WithSearchPlugin。
	// 默认为General，即通用问答。
	// 当前可体验模型仅支持General。
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 搜索来源。仅当Mode为WithSearchPlugin时生效。
	// 预置文稿库：填写Preset。自定义：填写Custom。
	SearchSource *string `json:"SearchSource,omitnil,omitempty" name:"SearchSource"`
}

func (r *SendChatMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendChatMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "Question")
	delete(f, "ModelVersion")
	delete(f, "Mode")
	delete(f, "SearchSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendChatMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendChatMessageResponseParams struct {
	// 答案
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 会话id,返回请求的会话id
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SendChatMessageResponse struct {
	*tchttp.BaseResponse
	Response *SendChatMessageResponseParams `json:"Response"`
}

func (r *SendChatMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendChatMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Service struct {
	// 服务组id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 服务id
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 服务组名
	ServiceGroupName *string `json:"ServiceGroupName,omitnil,omitempty" name:"ServiceGroupName"`

	// 服务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceDescription *string `json:"ServiceDescription,omitnil,omitempty" name:"ServiceDescription"`

	// 服务的详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceInfo *ServiceInfo `json:"ServiceInfo,omitnil,omitempty" name:"ServiceInfo"`

	// 集群id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 付费类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 包年包月服务的资源组id，按量计费的服务为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 包年包月服务对应的资源组名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// 服务的标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 服务所在的 ingress 的 name
	// 注意：此字段可能返回 null，表示取不到有效值。
	IngressName *string `json:"IngressName,omitnil,omitempty" name:"IngressName"`

	// 创建者
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedBy *string `json:"CreatedBy,omitnil,omitempty" name:"CreatedBy"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 主账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 子账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// app_id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 服务的业务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessStatus *string `json:"BusinessStatus,omitnil,omitempty" name:"BusinessStatus"`

	// 已废弃,以ServiceInfo中的对应为准
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ServiceLimit is deprecated.
	ServiceLimit *ServiceLimit `json:"ServiceLimit,omitnil,omitempty" name:"ServiceLimit"`

	// 已废弃,以ServiceInfo中的对应为准
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ScheduledAction is deprecated.
	ScheduledAction *ScheduledAction `json:"ScheduledAction,omitnil,omitempty" name:"ScheduledAction"`

	// 服务创建失败的原因，创建成功后该字段为默认值 CREATE_SUCCEED
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateFailedReason *string `json:"CreateFailedReason,omitnil,omitempty" name:"CreateFailedReason"`

	// 服务状态
	// CREATING 创建中
	// CREATE_FAILED 创建失败
	// Normal	正常运行中
	// Stopped  已停止
	// Stopping 停止中
	// Abnormal 异常
	// Pending 启动中
	// Waiting 就绪中
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 费用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingInfo *string `json:"BillingInfo,omitnil,omitempty" name:"BillingInfo"`

	// 模型权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 服务的创建来源
	// AUTO_ML: 来自自动学习的一键发布
	// DEFAULT: 其他来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateSource *string `json:"CreateSource,omitnil,omitempty" name:"CreateSource"`

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 服务组下服务的最高版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestVersion *string `json:"LatestVersion,omitnil,omitempty" name:"LatestVersion"`

	// 资源组类别 托管 NORMAL，纳管 SW
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupSWType *string `json:"ResourceGroupSWType,omitnil,omitempty" name:"ResourceGroupSWType"`
}

type ServiceCallInfo struct {
	// 服务组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 内网http调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerHttpAddr *string `json:"InnerHttpAddr,omitnil,omitempty" name:"InnerHttpAddr"`

	// 内网https调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerHttpsAddr *string `json:"InnerHttpsAddr,omitnil,omitempty" name:"InnerHttpsAddr"`

	// 内网http调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	OuterHttpAddr *string `json:"OuterHttpAddr,omitnil,omitempty" name:"OuterHttpAddr"`

	// 内网https调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	OuterHttpsAddr *string `json:"OuterHttpsAddr,omitnil,omitempty" name:"OuterHttpsAddr"`

	// 调用key
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// 调用secret
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppSecret *string `json:"AppSecret,omitnil,omitempty" name:"AppSecret"`

	// 鉴权是否开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorizationEnable *bool `json:"AuthorizationEnable,omitnil,omitempty" name:"AuthorizationEnable"`
}

type ServiceEIP struct {
	// 是否开启TIONE内网到外部的访问
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableEIP *bool `json:"EnableEIP,omitnil,omitempty" name:"EnableEIP"`

	// 用户VpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 用户subnetId
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

type ServiceEIPInfo struct {
	// 服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 用户VpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 用户子网Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

type ServiceGroup struct {
	// 服务组id
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`

	// 服务组名
	ServiceGroupName *string `json:"ServiceGroupName,omitnil,omitempty" name:"ServiceGroupName"`

	// 创建者
	CreatedBy *string `json:"CreatedBy,omitnil,omitempty" name:"CreatedBy"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 主账号
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 服务组下服务总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceCount *uint64 `json:"ServiceCount,omitnil,omitempty" name:"ServiceCount"`

	// 服务组下在运行的服务数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningServiceCount *uint64 `json:"RunningServiceCount,omitnil,omitempty" name:"RunningServiceCount"`

	// 服务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Services []*Service `json:"Services,omitnil,omitempty" name:"Services"`

	// 服务组状态，与服务一致
	//  CREATING 创建中
	//      CREATE_FAILED 创建失败
	//      Normal	正常运行中
	//      Stopped  已停止
	//      Stopping 停止中
	//      Abnormal 异常
	//      Pending 启动中
	//      Waiting 就绪中
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 服务组标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 服务组下最高版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestVersion *string `json:"LatestVersion,omitnil,omitempty" name:"LatestVersion"`

	// 服务的业务状态
	// CREATING 创建中
	//      CREATE_FAILED 创建失败
	//      ARREARS_STOP 因欠费被强制停止
	//      BILLING 计费中
	//      WHITELIST_USING 白名单试用中
	//      WHITELIST_STOP 白名单额度不足
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessStatus *string `json:"BusinessStatus,omitnil,omitempty" name:"BusinessStatus"`

	// 服务的计费信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingInfo *string `json:"BillingInfo,omitnil,omitempty" name:"BillingInfo"`

	// 服务的创建来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateSource *string `json:"CreateSource,omitnil,omitempty" name:"CreateSource"`

	// 服务组的权重更新状态 
	// UPDATING 更新中
	//      UPDATED 更新成功
	//      UPDATE_FAILED 更新失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeightUpdateStatus *string `json:"WeightUpdateStatus,omitnil,omitempty" name:"WeightUpdateStatus"`

	// 服务组下运行的pod数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplicasCount *uint64 `json:"ReplicasCount,omitnil,omitempty" name:"ReplicasCount"`

	// 服务组下期望的pod数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableReplicasCount *uint64 `json:"AvailableReplicasCount,omitnil,omitempty" name:"AvailableReplicasCount"`
}

type ServiceHistory struct {
	// 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Revision *string `json:"Revision,omitnil,omitempty" name:"Revision"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 镜像
	// 注意：此字段可能返回 null，表示取不到有效值。
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 模型文件
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelFile *string `json:"ModelFile,omitnil,omitempty" name:"ModelFile"`

	// 原始数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	RawData *string `json:"RawData,omitnil,omitempty" name:"RawData"`
}

type ServiceInfo struct {
	// 期望运行的Pod数量，停止状态是0
	// 不同计费模式和调节模式下对应关系如下
	// PREPAID 和 POSTPAID_BY_HOUR:
	// 手动调节模式下对应 实例数量
	// 自动调节模式下对应 基于时间的默认策略的实例数量
	// HYBRID_PAID:
	// 后付费实例手动调节模式下对应 实例数量
	// 后付费实例自动调节模式下对应 时间策略的默认策略的实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// 镜像信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 环境变量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Env []*EnvVar `json:"Env,omitnil,omitempty" name:"Env"`

	// 资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resources *ResourceInfo `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 后付费实例对应的机型规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 模型信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelInfo *ModelInfo `json:"ModelInfo,omitnil,omitempty" name:"ModelInfo"`

	// 是否启用日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// 日志配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// 是否开启鉴权
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorizationEnable *bool `json:"AuthorizationEnable,omitnil,omitempty" name:"AuthorizationEnable"`

	// hpa配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	HorizontalPodAutoscaler *HorizontalPodAutoscaler `json:"HorizontalPodAutoscaler,omitnil,omitempty" name:"HorizontalPodAutoscaler"`

	// 服务的状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *WorkloadStatus `json:"Status,omitnil,omitempty" name:"Status"`

	// 权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 资源总量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceTotal *ResourceInfo `json:"ResourceTotal,omitnil,omitempty" name:"ResourceTotal"`

	// 历史实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldReplicas *int64 `json:"OldReplicas,omitnil,omitempty" name:"OldReplicas"`

	// 计费模式[HYBRID_PAID]时生效, 用于标识混合计费模式下的预付费实例数, 若不填则默认为1
	// 注意：此字段可能返回 null，表示取不到有效值。
	HybridBillingPrepaidReplicas *int64 `json:"HybridBillingPrepaidReplicas,omitnil,omitempty" name:"HybridBillingPrepaidReplicas"`

	// 历史 HYBRID_PAID 时的实例数，用户恢复服务
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldHybridBillingPrepaidReplicas *int64 `json:"OldHybridBillingPrepaidReplicas,omitnil,omitempty" name:"OldHybridBillingPrepaidReplicas"`

	// 是否开启模型的热更新。默认不开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelHotUpdateEnable *bool `json:"ModelHotUpdateEnable,omitnil,omitempty" name:"ModelHotUpdateEnable"`

	// 实例数量调节方式,默认为手动
	// 支持：自动 - "AUTO", 手动 - "MANUAL"
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScaleMode *string `json:"ScaleMode,omitnil,omitempty" name:"ScaleMode"`

	// 定时伸缩任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	CronScaleJobs []*CronScaleJob `json:"CronScaleJobs,omitnil,omitempty" name:"CronScaleJobs"`

	// 定时伸缩策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScaleStrategy *string `json:"ScaleStrategy,omitnil,omitempty" name:"ScaleStrategy"`

	// 定时停止的配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduledAction *string `json:"ScheduledAction,omitnil,omitempty" name:"ScheduledAction"`

	// 实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: PodList is deprecated.
	PodList []*string `json:"PodList,omitnil,omitempty" name:"PodList"`

	// Pod列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Pods is deprecated.
	Pods *Pod `json:"Pods,omitnil,omitempty" name:"Pods"`

	// Pod列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodInfos []*Pod `json:"PodInfos,omitnil,omitempty" name:"PodInfos"`

	// 服务限速限流相关配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceLimit *ServiceLimit `json:"ServiceLimit,omitnil,omitempty" name:"ServiceLimit"`

	// 是否开启模型的加速, 仅对StableDiffusion(动态加速)格式的模型有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelTurboEnable *bool `json:"ModelTurboEnable,omitnil,omitempty" name:"ModelTurboEnable"`

	// 挂载
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeMount *VolumeMount `json:"VolumeMount,omitnil,omitempty" name:"VolumeMount"`

	// 推理代码信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InferCodeInfo *InferCodeInfo `json:"InferCodeInfo,omitnil,omitempty" name:"InferCodeInfo"`

	// 服务的启动命令
	// 注意：此字段可能返回 null，表示取不到有效值。
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 开启TIONE内网访问外部设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceEIP *ServiceEIP `json:"ServiceEIP,omitnil,omitempty" name:"ServiceEIP"`
}

type ServiceLimit struct {
	// 是否开启实例层面限流限速，true or false。true 则 InstanceRpsLimit 必填， false 则 InstanceRpsLimit 不生效
	EnableInstanceRpsLimit *bool `json:"EnableInstanceRpsLimit,omitnil,omitempty" name:"EnableInstanceRpsLimit"`

	// 每个服务实例的 request per second 限速, 0 为不限流
	InstanceRpsLimit *int64 `json:"InstanceRpsLimit,omitnil,omitempty" name:"InstanceRpsLimit"`

	// 是否开启单实例最大并发数限制，true or false。true 则 InstanceReqLimit 必填， false 则 InstanceReqLimit 不生效
	EnableInstanceReqLimit *bool `json:"EnableInstanceReqLimit,omitnil,omitempty" name:"EnableInstanceReqLimit"`

	// 每个服务实例的最大并发
	InstanceReqLimit *int64 `json:"InstanceReqLimit,omitnil,omitempty" name:"InstanceReqLimit"`
}

type Spec struct {
	// 计费项标签
	SpecId *string `json:"SpecId,omitnil,omitempty" name:"SpecId"`

	// 计费项名称
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 计费项显示名称
	SpecAlias *string `json:"SpecAlias,omitnil,omitempty" name:"SpecAlias"`

	// 是否售罄
	Available *bool `json:"Available,omitnil,omitempty" name:"Available"`

	// 当前资源售罄时，可用的区域有哪些
	AvailableRegion []*string `json:"AvailableRegion,omitnil,omitempty" name:"AvailableRegion"`

	// 当前计费项支持的特性
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecFeatures []*string `json:"SpecFeatures,omitnil,omitempty" name:"SpecFeatures"`

	// 计费项类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecType *string `json:"SpecType,omitnil,omitempty" name:"SpecType"`

	// GPU类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuType *string `json:"GpuType,omitnil,omitempty" name:"GpuType"`

	// 计费项CategoryId
	// 注意：此字段可能返回 null，表示取不到有效值。
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`
}

type SpecPrice struct {
	// 计费项名称
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 原价，单位：分。最大值42亿，超过则返回0
	TotalCost *uint64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 优惠后的价格，单位：分
	RealTotalCost *uint64 `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 计费项数量
	SpecCount *uint64 `json:"SpecCount,omitnil,omitempty" name:"SpecCount"`
}

type SpecUnit struct {
	// 计费项名称
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 计费项数量,建议不超过100万
	SpecCount *uint64 `json:"SpecCount,omitnil,omitempty" name:"SpecCount"`
}

type StartCmdInfo struct {
	// 启动命令
	StartCmd *string `json:"StartCmd,omitnil,omitempty" name:"StartCmd"`

	// ps启动命令
	PsStartCmd *string `json:"PsStartCmd,omitnil,omitempty" name:"PsStartCmd"`

	// worker启动命令
	WorkerStartCmd *string `json:"WorkerStartCmd,omitnil,omitempty" name:"WorkerStartCmd"`
}

// Predefined struct for user
type StartNotebookRequestParams struct {
	// notebook id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type StartNotebookRequest struct {
	*tchttp.BaseRequest
	
	// notebook id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *StartNotebookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartNotebookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartNotebookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartNotebookResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartNotebookResponse struct {
	*tchttp.BaseResponse
	Response *StartNotebookResponseParams `json:"Response"`
}

func (r *StartNotebookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartNotebookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartTrainingTaskRequestParams struct {
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type StartTrainingTaskRequest struct {
	*tchttp.BaseRequest
	
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *StartTrainingTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartTrainingTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartTrainingTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartTrainingTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartTrainingTaskResponse struct {
	*tchttp.BaseResponse
	Response *StartTrainingTaskResponseParams `json:"Response"`
}

func (r *StartTrainingTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartTrainingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StatefulSetCondition struct {
	// 信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// Status of the condition, one of True, False, Unknown.
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 上次更新的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTransitionTime *string `json:"LastTransitionTime,omitnil,omitempty" name:"LastTransitionTime"`

	// 上次更新的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`
}

// Predefined struct for user
type StopBatchTaskRequestParams struct {
	// 批量预测任务ID
	BatchTaskId *string `json:"BatchTaskId,omitnil,omitempty" name:"BatchTaskId"`
}

type StopBatchTaskRequest struct {
	*tchttp.BaseRequest
	
	// 批量预测任务ID
	BatchTaskId *string `json:"BatchTaskId,omitnil,omitempty" name:"BatchTaskId"`
}

func (r *StopBatchTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopBatchTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopBatchTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopBatchTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopBatchTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopBatchTaskResponseParams `json:"Response"`
}

func (r *StopBatchTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopBatchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopCreatingImageRequestParams struct {
	// 镜像保存记录ID
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`
}

type StopCreatingImageRequest struct {
	*tchttp.BaseRequest
	
	// 镜像保存记录ID
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`
}

func (r *StopCreatingImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCreatingImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RecordId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopCreatingImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopCreatingImageResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopCreatingImageResponse struct {
	*tchttp.BaseResponse
	Response *StopCreatingImageResponseParams `json:"Response"`
}

func (r *StopCreatingImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCreatingImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopModelAccelerateTaskRequestParams struct {
	// 模型加速任务ID
	ModelAccTaskId *string `json:"ModelAccTaskId,omitnil,omitempty" name:"ModelAccTaskId"`
}

type StopModelAccelerateTaskRequest struct {
	*tchttp.BaseRequest
	
	// 模型加速任务ID
	ModelAccTaskId *string `json:"ModelAccTaskId,omitnil,omitempty" name:"ModelAccTaskId"`
}

func (r *StopModelAccelerateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopModelAccelerateTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelAccTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopModelAccelerateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopModelAccelerateTaskResponseParams struct {
	// 模型加速任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAccTaskId *string `json:"ModelAccTaskId,omitnil,omitempty" name:"ModelAccTaskId"`

	// 异步任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncTaskId *string `json:"AsyncTaskId,omitnil,omitempty" name:"AsyncTaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopModelAccelerateTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopModelAccelerateTaskResponseParams `json:"Response"`
}

func (r *StopModelAccelerateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopModelAccelerateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopNotebookRequestParams struct {
	// notebook id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type StopNotebookRequest struct {
	*tchttp.BaseRequest
	
	// notebook id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *StopNotebookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopNotebookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopNotebookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopNotebookResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopNotebookResponse struct {
	*tchttp.BaseResponse
	Response *StopNotebookResponseParams `json:"Response"`
}

func (r *StopNotebookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopNotebookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopTrainingTaskRequestParams struct {
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type StopTrainingTaskRequest struct {
	*tchttp.BaseRequest
	
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *StopTrainingTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopTrainingTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopTrainingTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopTrainingTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopTrainingTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopTrainingTaskResponseParams `json:"Response"`
}

func (r *StopTrainingTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopTrainingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TJCallInfo struct {
	// 调用地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpAddr *string `json:"HttpAddr,omitnil,omitempty" name:"HttpAddr"`

	// token
	// 注意：此字段可能返回 null，表示取不到有效值。
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 调用示例
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallExample *string `json:"CallExample,omitnil,omitempty" name:"CallExample"`
}

type Tag struct {
	// 标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TagFilter struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 多个标签值
	TagValues []*string `json:"TagValues,omitnil,omitempty" name:"TagValues"`
}

type TextLabelDistributionDetailInfoFifthClass struct {
	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelValue *string `json:"LabelValue,omitnil,omitempty" name:"LabelValue"`

	// 标签个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelCount *uint64 `json:"LabelCount,omitnil,omitempty" name:"LabelCount"`

	// 标签占比
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelPercentage *float64 `json:"LabelPercentage,omitnil,omitempty" name:"LabelPercentage"`
}

type TextLabelDistributionDetailInfoFirstClass struct {
	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelValue *string `json:"LabelValue,omitnil,omitempty" name:"LabelValue"`

	// 标签个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelCount *uint64 `json:"LabelCount,omitnil,omitempty" name:"LabelCount"`

	// 标签占比
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelPercentage *float64 `json:"LabelPercentage,omitnil,omitempty" name:"LabelPercentage"`

	// 子标签分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChildLabelList []*TextLabelDistributionDetailInfoSecondClass `json:"ChildLabelList,omitnil,omitempty" name:"ChildLabelList"`
}

type TextLabelDistributionDetailInfoFourthClass struct {
	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelValue *string `json:"LabelValue,omitnil,omitempty" name:"LabelValue"`

	// 标签个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelCount *uint64 `json:"LabelCount,omitnil,omitempty" name:"LabelCount"`

	// 标签占比
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelPercentage *float64 `json:"LabelPercentage,omitnil,omitempty" name:"LabelPercentage"`

	// 子标签分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChildLabelList []*TextLabelDistributionDetailInfoFifthClass `json:"ChildLabelList,omitnil,omitempty" name:"ChildLabelList"`
}

type TextLabelDistributionDetailInfoSecondClass struct {
	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelValue *string `json:"LabelValue,omitnil,omitempty" name:"LabelValue"`

	// 标签个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelCount *uint64 `json:"LabelCount,omitnil,omitempty" name:"LabelCount"`

	// 标签占比
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelPercentage *float64 `json:"LabelPercentage,omitnil,omitempty" name:"LabelPercentage"`

	// 子标签分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChildLabelList []*TextLabelDistributionDetailInfoThirdClass `json:"ChildLabelList,omitnil,omitempty" name:"ChildLabelList"`
}

type TextLabelDistributionDetailInfoThirdClass struct {
	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelValue *string `json:"LabelValue,omitnil,omitempty" name:"LabelValue"`

	// 标签个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelCount *uint64 `json:"LabelCount,omitnil,omitempty" name:"LabelCount"`

	// 标签占比
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelPercentage *float64 `json:"LabelPercentage,omitnil,omitempty" name:"LabelPercentage"`

	// 子标签分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChildLabelList []*TextLabelDistributionDetailInfoFourthClass `json:"ChildLabelList,omitnil,omitempty" name:"ChildLabelList"`
}

type TextLabelDistributionInfo struct {
	// 文本分类题目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Theme *string `json:"Theme,omitnil,omitempty" name:"Theme"`

	// 一级标签分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassLabelList []*TextLabelDistributionDetailInfoFirstClass `json:"ClassLabelList,omitnil,omitempty" name:"ClassLabelList"`
}

type TrainingDataPoint struct {
	// 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *uint64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 训练上报的值。可以为训练指标（双精度浮点数，也可以为Epoch/Step（两者皆保证是整数）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type TrainingMetric struct {
	// 指标名
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 数据值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*TrainingDataPoint `json:"Values,omitnil,omitempty" name:"Values"`

	// 上报的Epoch. 可能为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Epochs []*TrainingDataPoint `json:"Epochs,omitnil,omitempty" name:"Epochs"`

	// 上报的Step. 可能为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Steps []*TrainingDataPoint `json:"Steps,omitnil,omitempty" name:"Steps"`

	// 上报的TotalSteps. 可能为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSteps []*TrainingDataPoint `json:"TotalSteps,omitnil,omitempty" name:"TotalSteps"`
}

type TrainingModelDTO struct {
	// 模型id
	TrainingModelId *string `json:"TrainingModelId,omitnil,omitempty" name:"TrainingModelId"`

	// 模型名称
	TrainingModelName *string `json:"TrainingModelName,omitnil,omitempty" name:"TrainingModelName"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 模型创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模型版本列表。默认不返回，仅在指定请求参数开启时返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingModelVersions []*TrainingModelVersionDTO `json:"TrainingModelVersions,omitnil,omitempty" name:"TrainingModelVersions"`
}

type TrainingModelVersionDTO struct {
	// 模型id
	TrainingModelId *string `json:"TrainingModelId,omitnil,omitempty" name:"TrainingModelId"`

	// 模型版本id
	TrainingModelVersionId *string `json:"TrainingModelVersionId,omitnil,omitempty" name:"TrainingModelVersionId"`

	// 模型版本
	TrainingModelVersion *string `json:"TrainingModelVersion,omitnil,omitempty" name:"TrainingModelVersion"`

	// 模型来源
	TrainingModelSource *string `json:"TrainingModelSource,omitnil,omitempty" name:"TrainingModelSource"`

	// 创建时间
	TrainingModelCreateTime *string `json:"TrainingModelCreateTime,omitnil,omitempty" name:"TrainingModelCreateTime"`

	// 创建人uin
	TrainingModelCreator *string `json:"TrainingModelCreator,omitnil,omitempty" name:"TrainingModelCreator"`

	// 算法框架
	AlgorithmFramework *string `json:"AlgorithmFramework,omitnil,omitempty" name:"AlgorithmFramework"`

	// 推理环境
	ReasoningEnvironment *string `json:"ReasoningEnvironment,omitnil,omitempty" name:"ReasoningEnvironment"`

	// 推理环境来源
	ReasoningEnvironmentSource *string `json:"ReasoningEnvironmentSource,omitnil,omitempty" name:"ReasoningEnvironmentSource"`

	// 模型指标
	TrainingModelIndex *string `json:"TrainingModelIndex,omitnil,omitempty" name:"TrainingModelIndex"`

	// 训练任务名称
	TrainingJobName *string `json:"TrainingJobName,omitnil,omitempty" name:"TrainingJobName"`

	// 模型cos路径
	TrainingModelCosPath *CosPathInfo `json:"TrainingModelCosPath,omitnil,omitempty" name:"TrainingModelCosPath"`

	// 模型名称
	TrainingModelName *string `json:"TrainingModelName,omitnil,omitempty" name:"TrainingModelName"`

	// 训练任务id
	TrainingJobId *string `json:"TrainingJobId,omitnil,omitempty" name:"TrainingJobId"`

	// 自定义推理环境
	ReasoningImageInfo *ImageInfo `json:"ReasoningImageInfo,omitnil,omitempty" name:"ReasoningImageInfo"`

	// 模型版本创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 模型处理状态
	// STATUS_SUCCESS：导入成功，STATUS_FAILED：导入失败 ，STATUS_RUNNING：导入中
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingModelStatus *string `json:"TrainingModelStatus,omitnil,omitempty" name:"TrainingModelStatus"`

	// 模型处理进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingModelProgress *uint64 `json:"TrainingModelProgress,omitnil,omitempty" name:"TrainingModelProgress"`

	// 模型错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingModelErrorMsg *string `json:"TrainingModelErrorMsg,omitnil,omitempty" name:"TrainingModelErrorMsg"`

	// 模型格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingModelFormat *string `json:"TrainingModelFormat,omitnil,omitempty" name:"TrainingModelFormat"`

	// 模型版本类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionType *string `json:"VersionType,omitnil,omitempty" name:"VersionType"`

	// GPU类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GPUType *string `json:"GPUType,omitnil,omitempty" name:"GPUType"`

	// 模型自动清理开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoClean *string `json:"AutoClean,omitnil,omitempty" name:"AutoClean"`

	// 模型清理周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelCleanPeriod *uint64 `json:"ModelCleanPeriod,omitnil,omitempty" name:"ModelCleanPeriod"`

	// 模型数量保留上限
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxReservedModels *uint64 `json:"MaxReservedModels,omitnil,omitempty" name:"MaxReservedModels"`

	// 模型热更新目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelHotUpdatePath *CosPathInfo `json:"ModelHotUpdatePath,omitnil,omitempty" name:"ModelHotUpdatePath"`

	// 推理环境id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReasoningEnvironmentId *string `json:"ReasoningEnvironmentId,omitnil,omitempty" name:"ReasoningEnvironmentId"`

	// 训练任务版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingJobVersion *string `json:"TrainingJobVersion,omitnil,omitempty" name:"TrainingJobVersion"`

	// 训练偏好
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingPreference *string `json:"TrainingPreference,omitnil,omitempty" name:"TrainingPreference"`

	// 自动学习任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoMLTaskId *string `json:"AutoMLTaskId,omitnil,omitempty" name:"AutoMLTaskId"`

	// 是否QAT模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsQAT *bool `json:"IsQAT,omitnil,omitempty" name:"IsQAT"`
}

type TrainingTaskDetail struct {
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 训练任务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 主账号uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 子账号uin
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 训练框架名称，eg：SPARK、PYSARK、TENSORFLOW、PYTORCH
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkName *string `json:"FrameworkName,omitnil,omitempty" name:"FrameworkName"`

	// 训练框架版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkVersion *string `json:"FrameworkVersion,omitnil,omitempty" name:"FrameworkVersion"`

	// 框架运行环境
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkEnvironment *string `json:"FrameworkEnvironment,omitnil,omitempty" name:"FrameworkEnvironment"`

	// 计费模式
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 预付费专用资源组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 资源配置
	ResourceConfigInfos []*ResourceConfigInfo `json:"ResourceConfigInfos,omitnil,omitempty" name:"ResourceConfigInfos"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 训练模式，eg：PS_WORKER、DDP、MPI、HOROVOD
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingMode *string `json:"TrainingMode,omitnil,omitempty" name:"TrainingMode"`

	// 代码包
	CodePackagePath *CosPathInfo `json:"CodePackagePath,omitnil,omitempty" name:"CodePackagePath"`

	// 启动命令信息
	StartCmdInfo *StartCmdInfo `json:"StartCmdInfo,omitnil,omitempty" name:"StartCmdInfo"`

	// 数据来源，eg：DATASET、COS
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSource *string `json:"DataSource,omitnil,omitempty" name:"DataSource"`

	// 数据配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataConfigs []*DataConfig `json:"DataConfigs,omitnil,omitempty" name:"DataConfigs"`

	// 调优参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TuningParameters *string `json:"TuningParameters,omitnil,omitempty" name:"TuningParameters"`

	// 训练输出
	Output *CosPathInfo `json:"Output,omitnil,omitempty" name:"Output"`

	// 是否上报日志
	LogEnable *bool `json:"LogEnable,omitnil,omitempty" name:"LogEnable"`

	// 日志配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogConfig *LogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// VPC ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 自定义镜像信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 运行时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeInSeconds *uint64 `json:"RuntimeInSeconds,omitnil,omitempty" name:"RuntimeInSeconds"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 训练开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 计费状态，eg：BILLING计费中，ARREARS_STOP欠费停止，NOT_BILLING不在计费中
	ChargeStatus *string `json:"ChargeStatus,omitnil,omitempty" name:"ChargeStatus"`

	// 最近一次实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestInstanceId *string `json:"LatestInstanceId,omitnil,omitempty" name:"LatestInstanceId"`

	// TensorBoard ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TensorBoardId *string `json:"TensorBoardId,omitnil,omitempty" name:"TensorBoardId"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureReason *string `json:"FailureReason,omitnil,omitempty" name:"FailureReason"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 训练结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 计费金额信息，eg：2.00元/小时 (按量计费)
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingInfo *string `json:"BillingInfo,omitnil,omitempty" name:"BillingInfo"`

	// 预付费专用资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// 任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 任务状态，eg：STARTING启动中、RUNNING运行中、STOPPING停止中、STOPPED已停止、FAILED异常、SUCCEED已完成
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 回调地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`
}

type TrainingTaskSetItem struct {
	// 训练任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 训练任务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 框架名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkName *string `json:"FrameworkName,omitnil,omitempty" name:"FrameworkName"`

	// 训练框架版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkVersion *string `json:"FrameworkVersion,omitnil,omitempty" name:"FrameworkVersion"`

	// 框架运行环境
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkEnvironment *string `json:"FrameworkEnvironment,omitnil,omitempty" name:"FrameworkEnvironment"`

	// 计费模式
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 计费状态，eg：BILLING计费中，ARREARS_STOP欠费停止，NOT_BILLING不在计费中
	ChargeStatus *string `json:"ChargeStatus,omitnil,omitempty" name:"ChargeStatus"`

	// 预付费专用资源组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 资源配置
	ResourceConfigInfos []*ResourceConfigInfo `json:"ResourceConfigInfos,omitnil,omitempty" name:"ResourceConfigInfos"`

	// 训练模式eg：PS_WORKER、DDP、MPI、HOROVOD
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingMode *string `json:"TrainingMode,omitnil,omitempty" name:"TrainingMode"`

	// 任务状态，eg：SUBMITTING提交中、PENDING排队中、
	// STARTING启动中、RUNNING运行中、STOPPING停止中、STOPPED已停止、FAILED异常、SUCCEED已完成
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 运行时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeInSeconds *uint64 `json:"RuntimeInSeconds,omitnil,omitempty" name:"RuntimeInSeconds"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 训练开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 训练结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 训练输出
	Output *CosPathInfo `json:"Output,omitnil,omitempty" name:"Output"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureReason *string `json:"FailureReason,omitnil,omitempty" name:"FailureReason"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 计费金额信息，eg：2.00元/小时 (按量计费)
	BillingInfo *string `json:"BillingInfo,omitnil,omitempty" name:"BillingInfo"`

	// 预付费专用资源组名称
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// 自定义镜像信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageInfo *ImageInfo `json:"ImageInfo,omitnil,omitempty" name:"ImageInfo"`

	// 任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 标签配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 回调地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`
}

type Usage struct {
	// 生成的token数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompletionTokens *int64 `json:"CompletionTokens,omitnil,omitempty" name:"CompletionTokens"`

	// 输入的token数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	PromptTokens *int64 `json:"PromptTokens,omitnil,omitempty" name:"PromptTokens"`

	// 总共token数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalTokens *int64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}

type VolumeMount struct {
	// cfs的配置信息
	CFSConfig *CFSConfig `json:"CFSConfig,omitnil,omitempty" name:"CFSConfig"`

	// 挂载源类型，CFS、COS，默认为CFS
	VolumeSourceType *string `json:"VolumeSourceType,omitnil,omitempty" name:"VolumeSourceType"`
}

type WeightEntry struct {
	// 服务id
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 流量权重值，同 ServiceGroup 下 总和应为 100
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type WorkloadStatus struct {
	// 当前实例数
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// 更新的实例数
	UpdatedReplicas *int64 `json:"UpdatedReplicas,omitnil,omitempty" name:"UpdatedReplicas"`

	// 就绪的实例数
	ReadyReplicas *int64 `json:"ReadyReplicas,omitnil,omitempty" name:"ReadyReplicas"`

	// 可用的实例数
	AvailableReplicas *int64 `json:"AvailableReplicas,omitnil,omitempty" name:"AvailableReplicas"`

	// 不可用的实例数
	UnavailableReplicas *int64 `json:"UnavailableReplicas,omitnil,omitempty" name:"UnavailableReplicas"`

	// Normal	正常运行中
	// Abnormal	服务异常，例如容器启动失败等
	// Waiting	服务等待中，例如容器下载镜像过程等
	// Stopped   已停止 
	// Pending 启动中
	// Stopping 停止中
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 工作负载的状况信息
	//
	// Deprecated: StatefulSetCondition is deprecated.
	StatefulSetCondition []*StatefulSetCondition `json:"StatefulSetCondition,omitnil,omitempty" name:"StatefulSetCondition"`

	// 工作负载历史的状况信息
	Conditions []*StatefulSetCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`

	// 状态异常时，展示原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}