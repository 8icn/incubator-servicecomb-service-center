// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package backend

import (
	"github.com/apache/incubator-servicecomb-service-center/server/core"
	pb "github.com/apache/incubator-servicecomb-service-center/server/core/proto"
	"github.com/apache/incubator-servicecomb-service-center/server/plugin/pkg/discovery"
	"time"
)

const (
	leaseProfTimeFmt           = "15:04:05.000"
	eventBlockSize             = 1000
	deferCheckWindow           = 2 * time.Second // instance DELETE event will be delay.
	selfPreservationPercentage = 0.8
	selfPreservationMaxTTL     = 10 * 60 // 10min
	selfPreservationInitCount  = 5
)

var (
	DOMAIN           discovery.Type
	PROJECT          discovery.Type
	SERVICE          discovery.Type
	SERVICE_INDEX    discovery.Type
	SERVICE_ALIAS    discovery.Type
	SERVICE_TAG      discovery.Type
	RULE             discovery.Type
	RULE_INDEX       discovery.Type
	DEPENDENCY_RULE  discovery.Type
	DEPENDENCY_QUEUE discovery.Type
	SCHEMA           discovery.Type
	SCHEMA_SUMMARY   discovery.Type
	INSTANCE         discovery.Type
	LEASE            discovery.Type
)

func registerInnerTypes() {
	SERVICE = Store().MustInstall(NewAddOn("SERVICE",
		discovery.Configure().WithPrefix(core.GetServiceRootKey("")).
			WithInitSize(500).WithParser(pb.ServiceParser)))
	INSTANCE = Store().MustInstall(NewAddOn("INSTANCE",
		discovery.Configure().WithPrefix(core.GetInstanceRootKey("")).
			WithInitSize(1000).WithParser(pb.InstanceParser).
			WithDeferHandler(NewInstanceEventDeferHandler())))
	DOMAIN = Store().MustInstall(NewAddOn("DOMAIN",
		discovery.Configure().WithPrefix(core.GetDomainRootKey()+core.SPLIT).
			WithInitSize(100).WithParser(pb.StringParser)))
	SCHEMA = Store().MustInstall(NewAddOn("SCHEMA",
		discovery.Configure().WithPrefix(core.GetServiceSchemaRootKey("")).
			WithInitSize(0)))
	SCHEMA_SUMMARY = Store().MustInstall(NewAddOn("SCHEMA_SUMMARY",
		discovery.Configure().WithPrefix(core.GetServiceSchemaSummaryRootKey("")).
			WithInitSize(100).WithParser(pb.StringParser)))
	RULE = Store().MustInstall(NewAddOn("RULE",
		discovery.Configure().WithPrefix(core.GetServiceRuleRootKey("")).
			WithInitSize(100).WithParser(pb.RuleParser)))
	LEASE = Store().MustInstall(NewAddOn("LEASE",
		discovery.Configure().WithPrefix(core.GetInstanceLeaseRootKey("")).
			WithInitSize(1000).WithParser(pb.StringParser)))
	SERVICE_INDEX = Store().MustInstall(NewAddOn("SERVICE_INDEX",
		discovery.Configure().WithPrefix(core.GetServiceIndexRootKey("")).
			WithInitSize(500).WithParser(pb.StringParser)))
	SERVICE_ALIAS = Store().MustInstall(NewAddOn("SERVICE_ALIAS",
		discovery.Configure().WithPrefix(core.GetServiceAliasRootKey("")).
			WithInitSize(100).WithParser(pb.StringParser)))
	SERVICE_TAG = Store().MustInstall(NewAddOn("SERVICE_TAG",
		discovery.Configure().WithPrefix(core.GetServiceTagRootKey("")).
			WithInitSize(100).WithParser(pb.MapParser)))
	RULE_INDEX = Store().MustInstall(NewAddOn("RULE_INDEX",
		discovery.Configure().WithPrefix(core.GetServiceRuleIndexRootKey("")).
			WithInitSize(100).WithParser(pb.StringParser)))
	DEPENDENCY_RULE = Store().MustInstall(NewAddOn("DEPENDENCY_RULE",
		discovery.Configure().WithPrefix(core.GetServiceDependencyRuleRootKey("")).
			WithInitSize(100).WithParser(pb.DependencyRuleParser)))
	DEPENDENCY_QUEUE = Store().MustInstall(NewAddOn("DEPENDENCY_QUEUE",
		discovery.Configure().WithPrefix(core.GetServiceDependencyQueueRootKey("")).
			WithInitSize(100).WithParser(pb.DependencyQueueParser)))
	PROJECT = Store().MustInstall(NewAddOn("PROJECT",
		discovery.Configure().WithPrefix(core.GetProjectRootKey("")).
			WithInitSize(100).WithParser(pb.StringParser)))
}
