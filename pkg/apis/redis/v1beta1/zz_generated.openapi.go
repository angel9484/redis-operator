// +build !

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1beta1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/ucloud/redis-operator/pkg/apis/redis/v1beta1.RedisCluster":       schema_pkg_apis_redis_v1beta1_RedisCluster(ref),
		"github.com/ucloud/redis-operator/pkg/apis/redis/v1beta1.RedisClusterSpec":   schema_pkg_apis_redis_v1beta1_RedisClusterSpec(ref),
		"github.com/ucloud/redis-operator/pkg/apis/redis/v1beta1.RedisClusterStatus": schema_pkg_apis_redis_v1beta1_RedisClusterStatus(ref),
	}
}

func schema_pkg_apis_redis_v1beta1_RedisCluster(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RedisCluster is the Schema for the redisclusters API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/ucloud/redis-operator/pkg/apis/redis/v1beta1.RedisClusterSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/ucloud/redis-operator/pkg/apis/redis/v1beta1.RedisClusterStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/ucloud/redis-operator/pkg/apis/redis/v1beta1.RedisClusterSpec", "github.com/ucloud/redis-operator/pkg/apis/redis/v1beta1.RedisClusterStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_redis_v1beta1_RedisClusterSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RedisClusterSpec defines the desired state of RedisCluster",
				Properties: map[string]spec.Schema{
					"size": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"image": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"command": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"shutdownConfigMap": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"storage": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/ucloud/redis-operator/pkg/apis/redis/v1beta1.RedisStorage"),
						},
					},
					"password": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"exporter": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/ucloud/redis-operator/pkg/apis/redis/v1beta1.RedisExporter"),
						},
					},
					"affinity": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.Affinity"),
						},
					},
					"securityContext": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.PodSecurityContext"),
						},
					},
					"toleRations": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.Toleration"),
									},
								},
							},
						},
					},
					"config": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"disablePersistence": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"sentinel": {
						SchemaProps: spec.SchemaProps{
							Description: "Sentinel defines its cluster settings",
							Ref:         ref("github.com/ucloud/redis-operator/pkg/apis/redis/v1beta1.SentinelSettings"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/ucloud/redis-operator/pkg/apis/redis/v1beta1.RedisExporter", "github.com/ucloud/redis-operator/pkg/apis/redis/v1beta1.RedisStorage", "github.com/ucloud/redis-operator/pkg/apis/redis/v1beta1.SentinelSettings", "k8s.io/api/core/v1.Affinity", "k8s.io/api/core/v1.PodSecurityContext", "k8s.io/api/core/v1.ResourceRequirements", "k8s.io/api/core/v1.Toleration"},
	}
}

func schema_pkg_apis_redis_v1beta1_RedisClusterStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RedisClusterStatus defines the observed state of RedisCluster",
				Properties: map[string]spec.Schema{
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/ucloud/redis-operator/pkg/apis/redis/v1beta1.Condition"),
									},
								},
							},
						},
					},
					"masterIP": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"sentinelIP": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/ucloud/redis-operator/pkg/apis/redis/v1beta1.Condition"},
	}
}