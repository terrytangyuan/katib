/*

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

package util

import (
	//v1 "k8s.io/api/core/v1"

	trialsv1alpha2 "github.com/kubeflow/katib/pkg/api/operators/apis/trial/v1alpha2"
)

func CreateTrialInDB(instance *trialsv1alpha2.Trial) error {

	return nil
}

func UpdateTrialStatusInDB(instance *trialsv1alpha2.Trial) error {

	return nil
}

func GetTrialObservation(instance *trialsv1alpha2.Trial) error {

	return nil
}
