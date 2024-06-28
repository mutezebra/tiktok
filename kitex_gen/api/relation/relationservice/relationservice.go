// Code generated by Kitex v0.9.1. DO NOT EDIT.

package relationservice

import (
	"context"
	"errors"

	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"

	relation "github.com/Mutezebra/tiktok/kitex_gen/api/relation"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Follow": kitex.NewMethodInfo(
		followHandler,
		newRelationServiceFollowArgs,
		newRelationServiceFollowResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetFollowList": kitex.NewMethodInfo(
		getFollowListHandler,
		newRelationServiceGetFollowListArgs,
		newRelationServiceGetFollowListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetFansList": kitex.NewMethodInfo(
		getFansListHandler,
		newRelationServiceGetFansListArgs,
		newRelationServiceGetFansListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetFriendsList": kitex.NewMethodInfo(
		getFriendsListHandler,
		newRelationServiceGetFriendsListArgs,
		newRelationServiceGetFriendsListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	relationServiceServiceInfo                = NewServiceInfo()
	relationServiceServiceInfoForClient       = NewServiceInfoForClient()
	relationServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return relationServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return relationServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return relationServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "RelationService"
	handlerType := (*relation.RelationService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "relation",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func followHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceFollowArgs)
	realResult := result.(*relation.RelationServiceFollowResult)
	success, err := handler.(relation.RelationService).Follow(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceFollowArgs() interface{} {
	return relation.NewRelationServiceFollowArgs()
}

func newRelationServiceFollowResult() interface{} {
	return relation.NewRelationServiceFollowResult()
}

func getFollowListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceGetFollowListArgs)
	realResult := result.(*relation.RelationServiceGetFollowListResult)
	success, err := handler.(relation.RelationService).GetFollowList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceGetFollowListArgs() interface{} {
	return relation.NewRelationServiceGetFollowListArgs()
}

func newRelationServiceGetFollowListResult() interface{} {
	return relation.NewRelationServiceGetFollowListResult()
}

func getFansListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceGetFansListArgs)
	realResult := result.(*relation.RelationServiceGetFansListResult)
	success, err := handler.(relation.RelationService).GetFansList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceGetFansListArgs() interface{} {
	return relation.NewRelationServiceGetFansListArgs()
}

func newRelationServiceGetFansListResult() interface{} {
	return relation.NewRelationServiceGetFansListResult()
}

func getFriendsListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceGetFriendsListArgs)
	realResult := result.(*relation.RelationServiceGetFriendsListResult)
	success, err := handler.(relation.RelationService).GetFriendsList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceGetFriendsListArgs() interface{} {
	return relation.NewRelationServiceGetFriendsListArgs()
}

func newRelationServiceGetFriendsListResult() interface{} {
	return relation.NewRelationServiceGetFriendsListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Follow(ctx context.Context, req *relation.FollowReq) (r *relation.FollowResp, err error) {
	var _args relation.RelationServiceFollowArgs
	_args.Req = req
	var _result relation.RelationServiceFollowResult
	if err = p.c.Call(ctx, "Follow", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFollowList(ctx context.Context, req *relation.GetFollowListReq) (r *relation.GetFollowListResp, err error) {
	var _args relation.RelationServiceGetFollowListArgs
	_args.Req = req
	var _result relation.RelationServiceGetFollowListResult
	if err = p.c.Call(ctx, "GetFollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFansList(ctx context.Context, req *relation.GetFansListReq) (r *relation.GetFansListResp, err error) {
	var _args relation.RelationServiceGetFansListArgs
	_args.Req = req
	var _result relation.RelationServiceGetFansListResult
	if err = p.c.Call(ctx, "GetFansList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFriendsList(ctx context.Context, req *relation.GetFriendsListReq) (r *relation.GetFriendsListResp, err error) {
	var _args relation.RelationServiceGetFriendsListArgs
	_args.Req = req
	var _result relation.RelationServiceGetFriendsListResult
	if err = p.c.Call(ctx, "GetFriendsList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
