// Package approval code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package larkapproval

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

func NewService(config *larkcore.Config) *ApprovalService {
	a := &ApprovalService{config: config}
	a.Approval = &approval{service: a}
	a.ExternalApproval = &externalApproval{service: a}
	a.ExternalInstance = &externalInstance{service: a}
	a.ExternalTask = &externalTask{service: a}
	a.Instance = &instance{service: a}
	a.InstanceComment = &instanceComment{service: a}
	a.Task = &task{service: a}
	return a
}

type ApprovalService struct {
	config           *larkcore.Config
	Approval         *approval         // 事件
	ExternalApproval *externalApproval // 三方审批定义
	ExternalInstance *externalInstance // 三方审批实例
	ExternalTask     *externalTask     // 三方审批任务
	Instance         *instance         // 原生审批实例
	InstanceComment  *instanceComment  // 原生审批评论
	Task             *task             // 原生审批任务
}

type approval struct {
	service *ApprovalService
}
type externalApproval struct {
	service *ApprovalService
}
type externalInstance struct {
	service *ApprovalService
}
type externalTask struct {
	service *ApprovalService
}
type instance struct {
	service *ApprovalService
}
type instanceComment struct {
	service *ApprovalService
}
type task struct {
	service *ApprovalService
}

// 创建审批定义
//
// - 用于通过接口创建简单的审批定义，可以灵活指定定义的基础信息、表单和流程等。创建成功后，不支持从审批管理后台删除该定义。不推荐企业自建应用使用，如有需要尽量联系管理员在审批管理后台创建定义。
//
// - 接口谨慎调用，创建后的审批定义无法停用/删除
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval/create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/create_approval.go
func (a *approval) Create(ctx context.Context, req *CreateApprovalReq, options ...larkcore.RequestOptionFunc) (*CreateApprovalResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/approvals"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateApprovalResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 查看审批定义
//
// - 根据 Approval Code 获取某个审批定义的详情，用于构造创建审批实例的请求。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval/get
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/get_approval.go
func (a *approval) Get(ctx context.Context, req *GetApprovalReq, options ...larkcore.RequestOptionFunc) (*GetApprovalResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/approvals/:approval_code"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetApprovalResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 查询审批定义列表
//
// - 查询当前用户可发起的审批定义列表。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval/list
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/list_approval.go
func (a *approval) List(ctx context.Context, req *ListApprovalReq, options ...larkcore.RequestOptionFunc) (*ListApprovalResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/approvals"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListApprovalResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (a *approval) ListByIterator(ctx context.Context, req *ListApprovalReq, options ...larkcore.RequestOptionFunc) (*ListApprovalIterator, error) {
	return &ListApprovalIterator{
		ctx:      ctx,
		req:      req,
		listFunc: a.List,
		options:  options,
		limit:    req.Limit}, nil
}

// 订阅审批事件
//
// - 应用订阅 approval_code 后，该应用就可以收到该审批定义对应实例的事件通知。同一应用只需要订阅一次，无需重复订阅。;;当应用不希望再收到审批事件时，可以使用取消订阅接口进行取消，取消后将不再给应用推送消息。;;订阅和取消订阅都是应用维度的，多个应用可以同时订阅同一个 approval_code，每个应用都能收到审批事件。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval/subscribe
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/subscribe_approval.go
func (a *approval) Subscribe(ctx context.Context, req *SubscribeApprovalReq, options ...larkcore.RequestOptionFunc) (*SubscribeApprovalResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/approvals/:approval_code/subscribe"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SubscribeApprovalResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 取消订阅审批事件
//
// - 取消订阅 approval_code 后，无法再收到该审批定义对应实例的事件通知
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval/unsubscribe
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/unsubscribe_approval.go
func (a *approval) Unsubscribe(ctx context.Context, req *UnsubscribeApprovalReq, options ...larkcore.RequestOptionFunc) (*UnsubscribeApprovalResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/approvals/:approval_code/unsubscribe"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UnsubscribeApprovalResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 三方审批定义创建
//
// - 审批定义是审批的描述，包括审批名称、图标、描述等基础信息。创建好审批定义，用户就可以在审批应用的发起页中看到审批，如果用户点击发起，则会跳转到配置的发起三方系统地址去发起审批。;;另外，审批定义还配置了审批操作时的回调地址：审批人在待审批列表中进行【同意】【拒绝】操作时，审批中心会调用回调地址通知三方系统。
//
// - 注意，审批中心不负责审批流程的流转，只负责展示、操作、消息通知。因此审批定义创建时没有审批流程的信息。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/external_approval/create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/create_externalApproval.go
func (e *externalApproval) Create(ctx context.Context, req *CreateExternalApprovalReq, options ...larkcore.RequestOptionFunc) (*CreateExternalApprovalResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/external_approvals"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, e.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateExternalApprovalResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, e.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 三方审批实例校验
//
// - 校验三方审批实例数据，用于判断服务端数据是否为最新的。用户提交实例最新更新时间，如果服务端不存在该实例，或者服务端实例更新时间不是最新的，则返回对应实例 id。;;例如，用户可以每隔5分钟，将最近5分钟产生的实例使用该接口进行对比。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/external_instance/check
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/check_externalInstance.go
func (e *externalInstance) Check(ctx context.Context, req *CheckExternalInstanceReq, options ...larkcore.RequestOptionFunc) (*CheckExternalInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/external_instances/check"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, e.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CheckExternalInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, e.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 三方审批实例同步
//
// - 审批中心不负责审批的流转，审批的流转在三方系统，三方系统在审批流转后生成的审批实例、审批任务、审批抄送数据同步到审批中心。;;用户可以在审批中心中浏览三方系统同步过来的实例、任务、抄送信息，并且可以跳转回三方系统进行更详细的查看和操作，其中实例信息在【已发起】列表，任务信息在【待审批】和【已审批】列表，抄送信息在【抄送我】列表;;:::html;<img src="//sf3-cn.feishucdn.com/obj/open-platform-opendoc/9dff4434afbeb0ef69de7f36b9a6e995_z5iwmTzEgg.png" alt="" style="zoom:17%;" />;;;<img src="//sf3-cn.feishucdn.com/obj/open-platform-opendoc/ca6e0e984a7a6d64e1b16a0bac4bf868_tfqjCiaJQM.png" alt="" style="zoom:17%;" />;;;<img src="//sf3-cn.feishucdn.com/obj/open-platform-opendoc/529377e238df78d391bbd22e962ad195_T7eefLI1GA.png" alt="" style="zoom:17%;" />;:::;;对于审批任务，三方系统也可以配置审批任务的回调接口，这样审批人可以在审批中心中直接进行审批操作，审批中心会回调三方系统，三方系统收到回调后更新任务信息，并将新的任务信息同步回审批中心，形成闭环。;;:::html;<img src="//sf3-cn.feishucdn.com/obj/open-platform-opendoc/721c35428bc1187db3318c572f9979ad_je75QpElcg.png" alt=""  style="zoom:25%;" />;:::;<br>
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/external_instance/create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/create_externalInstance.go
func (e *externalInstance) Create(ctx context.Context, req *CreateExternalInstanceReq, options ...larkcore.RequestOptionFunc) (*CreateExternalInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/external_instances"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, e.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateExternalInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, e.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取三方审批任务状态
//
// - 该接口用于获取三方审批的状态。用户传入查询条件，接口返回满足条件的审批实例的状态。该接口支持多种参数的组合，包括如下组合：;;1.通过 instance_ids 获取指定实例的任务状态;;2.通过 user_ids 获取指定用户的任务状态;;3.通过 status 获取指定状态的所有任务;;4.通过page_token获取下一批数据
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/external_task/list
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/list_externalTask.go
func (e *externalTask) List(ctx context.Context, req *ListExternalTaskReq, options ...larkcore.RequestOptionFunc) (*ListExternalTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/external_tasks"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, e.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListExternalTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, e.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *externalTask) ListByIterator(ctx context.Context, req *ListExternalTaskReq, options ...larkcore.RequestOptionFunc) (*ListExternalTaskIterator, error) {
	return &ListExternalTaskIterator{
		ctx:      ctx,
		req:      req,
		listFunc: e.List,
		options:  options,
		limit:    req.Limit}, nil
}

//
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/ukTM5UjL5ETO14SOxkTN/approval-task-addsign
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/addSign_instance.go
func (i *instance) AddSign(ctx context.Context, req *AddSignInstanceReq, options ...larkcore.RequestOptionFunc) (*AddSignInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/add_sign"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &AddSignInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 审批实例撤回
//
// - 对于状态为“审批中”的单个审批实例进行撤销操作，撤销后审批流程结束
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance/cancel
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/cancel_instance.go
func (i *instance) Cancel(ctx context.Context, req *CancelInstanceReq, options ...larkcore.RequestOptionFunc) (*CancelInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/cancel"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CancelInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 审批实例抄送
//
// - 通过接口可以将当前审批实例抄送给其他人。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance/cc
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/cc_instance.go
func (i *instance) Cc(ctx context.Context, req *CcInstanceReq, options ...larkcore.RequestOptionFunc) (*CcInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/cc"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CcInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 创建审批实例
//
// - 创建一个审批实例，调用方需对审批定义的表单有详细了解，将按照定义的表单结构，将表单 Value 通过接口传入
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance/create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/create_instance.go
func (i *instance) Create(ctx context.Context, req *CreateInstanceReq, options ...larkcore.RequestOptionFunc) (*CreateInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取单个审批实例详情
//
// - 通过审批实例 Instance Code  获取审批实例详情。Instance Code 由 [批量获取审批实例](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance/list) 接口获取。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance/get
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/get_instance.go
func (i *instance) Get(ctx context.Context, req *GetInstanceReq, options ...larkcore.RequestOptionFunc) (*GetInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/:instance_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 批量获取审批实例ID
//
// - 根据 approval_code 批量获取审批实例的 instance_code，用于拉取租户下某个审批定义的全部审批实例。默认以审批创建时间先后顺序排列
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance/list
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/list_instance.go
func (i *instance) List(ctx context.Context, req *ListInstanceReq, options ...larkcore.RequestOptionFunc) (*ListInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *instance) ListByIterator(ctx context.Context, req *ListInstanceReq, options ...larkcore.RequestOptionFunc) (*ListInstanceIterator, error) {
	return &ListInstanceIterator{
		ctx:      ctx,
		req:      req,
		listFunc: i.List,
		options:  options,
		limit:    req.Limit}, nil
}

//
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/ukTM5UjL5ETO14SOxkTN/approval-preview
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/preview_instance.go
func (i *instance) Preview(ctx context.Context, req *PreviewInstanceReq, options ...larkcore.RequestOptionFunc) (*PreviewInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/preview"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PreviewInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 查询实例列表
//
// - 该接口通过不同条件查询审批系统中符合条件的审批实例列表。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance/query
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/query_instance.go
func (i *instance) Query(ctx context.Context, req *QueryInstanceReq, options ...larkcore.RequestOptionFunc) (*QueryInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/query"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *instance) QueryByIterator(ctx context.Context, req *QueryInstanceReq, options ...larkcore.RequestOptionFunc) (*QueryInstanceIterator, error) {
	return &QueryInstanceIterator{
		ctx:      ctx,
		req:      req,
		listFunc: i.Query,
		options:  options,
		limit:    req.Limit}, nil
}

// 查询抄送列表
//
// - 该接口通过不同条件查询审批系统中符合条件的审批抄送列表。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance/search_cc
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/searchCc_instance.go
func (i *instance) SearchCc(ctx context.Context, req *SearchCcInstanceReq, options ...larkcore.RequestOptionFunc) (*SearchCcInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/search_cc"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SearchCcInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 审批任务退回
//
// - 从当前审批任务，退回到已审批的一个或多个任务节点。退回后，已审批节点重新生成审批任务
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance/specified_rollback
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/specifiedRollback_instance.go
func (i *instance) SpecifiedRollback(ctx context.Context, req *SpecifiedRollbackInstanceReq, options ...larkcore.RequestOptionFunc) (*SpecifiedRollbackInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/specified_rollback"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SpecifiedRollbackInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 创建评论
//
// - 在某审批实例下创建、修改评论或评论回复（不包含审批同意、拒绝、转交等附加的理由或意见）。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance-comment/create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/create_instanceComment.go
func (i *instanceComment) Create(ctx context.Context, req *CreateInstanceCommentReq, options ...larkcore.RequestOptionFunc) (*CreateInstanceCommentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/:instance_id/comments"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateInstanceCommentResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 删除评论
//
// - 逻辑删除某审批实例下的一条评论或评论回复（不包含审批同意、拒绝、转交等附加的理由或意见）。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance-comment/delete
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/delete_instanceComment.go
func (i *instanceComment) Delete(ctx context.Context, req *DeleteInstanceCommentReq, options ...larkcore.RequestOptionFunc) (*DeleteInstanceCommentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/:instance_id/comments/:comment_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteInstanceCommentResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取评论
//
// - 根据 Instance Code 获取某个审批实例下的全部评论与评论回复（不包含审批同意、拒绝、转交等附加的理由或意见）。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance-comment/list
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/list_instanceComment.go
func (i *instanceComment) List(ctx context.Context, req *ListInstanceCommentReq, options ...larkcore.RequestOptionFunc) (*ListInstanceCommentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/:instance_id/comments"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListInstanceCommentResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 清空评论
//
// - 删除某审批实例下的全部评论与评论回复。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance-comment/remove
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/remove_instanceComment.go
func (i *instanceComment) Remove(ctx context.Context, req *RemoveInstanceCommentReq, options ...larkcore.RequestOptionFunc) (*RemoveInstanceCommentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/:instance_id/comments/remove"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RemoveInstanceCommentResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 审批任务同意
//
// - 对于单个审批任务进行同意操作。同意后审批流程会流转到下一个审批人。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/task/approve
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/approve_task.go
func (t *task) Approve(ctx context.Context, req *ApproveTaskReq, options ...larkcore.RequestOptionFunc) (*ApproveTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/tasks/approve"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ApproveTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 用户角度列出任务
//
// - 根据用户和任务分组查询任务列表
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/task/query
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/query_task.go
func (t *task) Query(ctx context.Context, req *QueryTaskReq, options ...larkcore.RequestOptionFunc) (*QueryTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/tasks/query"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *task) QueryByIterator(ctx context.Context, req *QueryTaskReq, options ...larkcore.RequestOptionFunc) (*QueryTaskIterator, error) {
	return &QueryTaskIterator{
		ctx:      ctx,
		req:      req,
		listFunc: t.Query,
		options:  options,
		limit:    req.Limit}, nil
}

// 审批任务拒绝
//
// - 对于单个审批任务进行拒绝操作。拒绝后审批流程结束。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/task/reject
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/reject_task.go
func (t *task) Reject(ctx context.Context, req *RejectTaskReq, options ...larkcore.RequestOptionFunc) (*RejectTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/tasks/reject"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RejectTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 审批任务重新提交
//
// - 对于单个退回到发起人的审批任务进行重新发起操作。发起后审批流程会流转到下一个审批人。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/task/resubmit
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/resubmit_task.go
func (t *task) Resubmit(ctx context.Context, req *ResubmitTaskReq, options ...larkcore.RequestOptionFunc) (*ResubmitTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/tasks/resubmit"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ResubmitTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 查询任务列表
//
// - 该接口通过不同条件查询审批系统中符合条件的审批任务列表
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/task/search
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/search_task.go
func (t *task) Search(ctx context.Context, req *SearchTaskReq, options ...larkcore.RequestOptionFunc) (*SearchTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/tasks/search"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SearchTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 审批任务转交
//
// - 对于单个审批任务进行转交操作。转交后审批流程流转给被转交人。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/task/transfer
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/approvalv4/transfer_task.go
func (t *task) Transfer(ctx context.Context, req *TransferTaskReq, options ...larkcore.RequestOptionFunc) (*TransferTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/tasks/transfer"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &TransferTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
