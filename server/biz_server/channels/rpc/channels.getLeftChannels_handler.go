/*
 *  Copyright (c) 2018, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rpc

import (
    "github.com/golang/glog"
    "github.com/nebulaim/telegramd/proto/mtproto"
    "golang.org/x/net/context"
    "fmt"
    "github.com/nebulaim/telegramd/baselib/grpc_util"
    "github.com/nebulaim/telegramd/baselib/logger"
)

// channels.getLeftChannels#8341ecc0 offset:int = messages.Chats;
func (s *ChannelsServiceImpl) ChannelsGetLeftChannels(ctx context.Context, request *mtproto.TLChannelsGetLeftChannels) (*mtproto.Messages_Chats, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("ChannelsGetLeftChannels - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    // TODO(@benqi): Impl ChannelsGetLeftChannels logic

    return nil, fmt.Errorf("Not impl ChannelsGetLeftChannels")
}