/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package event

import (
	"github.com/apache/incubator-servicecomb-service-center/pkg/log"
	"github.com/apache/incubator-servicecomb-service-center/server/core"
	"github.com/apache/incubator-servicecomb-service-center/server/core/backend"
	pb "github.com/apache/incubator-servicecomb-service-center/server/core/proto"
	"github.com/apache/incubator-servicecomb-service-center/server/plugin/pkg/discovery"
	"github.com/apache/incubator-servicecomb-service-center/server/service/cache"
)

type DependencyRuleEventHandler struct {
}

func (h *DependencyRuleEventHandler) Type() discovery.Type {
	return backend.DEPENDENCY_RULE
}

func (h *DependencyRuleEventHandler) OnEvent(evt discovery.KvEvent) {
	action := evt.Type
	if action != pb.EVT_UPDATE && action != pb.EVT_DELETE {
		return
	}
	t, providerKey := core.GetInfoFromDependencyRuleKV(evt.KV.Key)
	if t != core.DEPS_PROVIDER {
		return
	}
	log.Debugf("caught [%s] provider rule[%s/%s/%s/%s] event",
		action, providerKey.Environment, providerKey.AppId, providerKey.ServiceName, providerKey.Version)
	cache.DependencyRule.Remove(providerKey)
}

func NewDependencyRuleEventHandler() *DependencyRuleEventHandler {
	return &DependencyRuleEventHandler{}
}
