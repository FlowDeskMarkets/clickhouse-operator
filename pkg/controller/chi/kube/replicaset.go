// Copyright 2019 Altinity Ltd and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kube

import (
	"context"

	apps "k8s.io/api/apps/v1"
	kube "k8s.io/client-go/kubernetes"

	"github.com/altinity/clickhouse-operator/pkg/controller"
)

type ReplicaSet struct {
	kubeClient kube.Interface
}

func NewReplicaSet(kubeClient kube.Interface) *ReplicaSet {
	return &ReplicaSet{
		kubeClient: kubeClient,
	}
}

func (c *ReplicaSet) Get(ctx context.Context, namespace, name string) (*apps.ReplicaSet, error) {
	ctx = k8sCtx(ctx)
	return c.kubeClient.AppsV1().ReplicaSets(namespace).Get(ctx, name, controller.NewGetOptions())
}

func (c *ReplicaSet) Update(ctx context.Context, replicaSet *apps.ReplicaSet) (*apps.ReplicaSet, error) {
	ctx = k8sCtx(ctx)
	return c.kubeClient.AppsV1().ReplicaSets(replicaSet.Namespace).Update(ctx, replicaSet, controller.NewUpdateOptions())
}
