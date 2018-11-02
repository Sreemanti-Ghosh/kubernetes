/*Copyright 2015 The Kubernetes Authors.
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
 
package topologymanager

import (
    "k8s.io/kubernetes/pkg/kubelet/lifecycle"
)

type strictPolicy struct {}

var _ Policy = &strictPolicy{}

const PolicyStrict policyName = "strict"

func NewStrictPolicy() Policy {
    return &strictPolicy{}
}

func (p *strictPolicy) Name() string {
    return string(PolicyStrict)
}

func (p *strictPolicy) CanAdmitPodResult (result TopologyHints) lifecycle.PodAdmitResult {
        socketMask := result.SocketAffinity.Mask
        affinity := false
        for _, outerMask := range socketMask {      
            for _, innerMask := range outerMask {
                if innerMask == 0 {
                    affinity = true
                    break
                }
            }
        }
        if !affinity {
            return lifecycle.PodAdmitResult{
                        Admit:  false,
                        Reason:	 "Topology Affinity Error",
                        Message: "Resources cannot be allocated with Topology Locality",
            }
        }  else {
            return lifecycle.PodAdmitResult{
                Admit:   true,          
            }
        }     
}
