/*
Copyright The Kubeform Authors.

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

package framework

import (
	"time"

	base "kubeform.dev/kubeform/apis/base/v1alpha1"
	"kubeform.dev/kubeform/apis/digitalocean/v1alpha1"

	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (i *Invocation) Droplets(name string) *v1alpha1.Droplet {
	return &v1alpha1.Droplet{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: i.Namespace(),
			Labels: map[string]string{
				"app": i.app,
			},
		},
		Spec: v1alpha1.DropletSpec{
			ProviderRef: corev1.LocalObjectReference{
				Name: DigitalOceanProviderRef,
			},
			Name:   name,
			Region: "nyc1",
			Size:   "s-1vcpu-1gb",
			Image:  "ubuntu-18-04-x64",
		},
	}
}

func (f *Framework) CreateDroplet(obj *v1alpha1.Droplet) error {
	_, err := f.kubeformClient.DigitaloceanV1alpha1().Droplets(obj.Namespace).Create(obj)
	return err
}

func (f *Framework) DeleteDroplet(meta metav1.ObjectMeta) error {
	return f.kubeformClient.DigitaloceanV1alpha1().Droplets(meta.Namespace).Delete(meta.Name, deleteInForeground())
}

func (f *Framework) EventuallyDropletRunning(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			droplet, err := f.kubeformClient.DigitaloceanV1alpha1().Droplets(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			return droplet.Status.Phase == base.PhaseRunning
		},
		time.Minute*15,
		time.Second*10,
	)
}

func (f *Framework) EventuallyDropletDeleted(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			_, err := f.kubeformClient.DigitaloceanV1alpha1().Droplets(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
			return errors.IsNotFound(err)
		},
		time.Minute*15,
		time.Second*10,
	)
}
