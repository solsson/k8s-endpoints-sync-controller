// Copyright © 2018 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: MIT

package utils

import (
	c "github.com/vmware/k8s-endpoints-sync-controller/pkg/config"
)

func ContainsInArray(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func CanReplicateNamespace(labels map[string]string) bool {
	if val, ok := labels[c.REPLICATED_LABEL_KEY]; ok {
		if val == "false" {
			return false
		}
	}
	return true
}

func ContainsKeyVal(labels map[string]string, val string) bool {
	if v, ok := labels[c.REPLICATED_LABEL_KEY]; ok {
		if v == val {
			return true
		}
	}
	return false
}
