/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package prefetch

const (
	ConfigKind     = "PrefetchConfig"
	ControllerName = "kueue.x-k8s.io/prefetch-request"
	// DeprecatedConsumesAnnotationKey  = "cluster-autoscaler.kubernetes.io/consume-provisioning-request"
	// DeprecatedClassNameAnnotationKey = "cluster-autoscaler.kubernetes.io/provisioning-class-name"
	ConsumesAnnotationKey  = "autoscaling.x-k8s.io/consume-prefetch-request"
	ClassNameAnnotationKey = "autoscaling.x-k8s.io/prefetch-class-name"

	CheckInactiveMessage = "the check is not active"
	NoRequestNeeded      = "the prefetch is not needed"
)
