/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package konnector

import (
	"context"
	"fmt"

	kubebindv1alpha1 "go.bytebuilders.dev/kube-bind/apis/kubebind/v1alpha1"

	"k8s.io/klog/v2"
	"kmodules.xyz/client-go/apiextensions"
)

type Server struct {
	Config     *Config
	Controller *Controller
}

func NewServer(config *Config) (*Server, error) {
	// construct controllers
	k, err := New(
		config.ClientConfig,
		config.BindInformers.KubeBind().V1alpha1().APIServiceBindings(),
		config.KubeInformers.Core().V1().Secrets(), // TODO(sttts): watch individual secrets for security and memory consumption
		config.KubeInformers.Core().V1().Namespaces(),
		config.ApiextensionsInformers.Apiextensions().V1().CustomResourceDefinitions(),
	)
	if err != nil {
		return nil, err
	}

	s := &Server{
		Config:     config,
		Controller: k,
	}

	return s, nil
}

type prepared struct {
	Server
}

type Prepared struct {
	*prepared
}

func (s *Server) PrepareRun(ctx context.Context) (Prepared, error) {
	// install/upgrade CRDs
	if err := apiextensions.RegisterCRDs(s.Config.ApiextensionsClient, []*apiextensions.CustomResourceDefinition{
		kubebindv1alpha1.APIServiceBinding{}.CustomResourceDefinition(),
	}); err != nil {
		return Prepared{}, err
	}
	return Prepared{
		prepared: &prepared{
			Server: *s,
		},
	}, nil
}

func (s *Prepared) OptionallyStartInformers(ctx context.Context) {
	logger := klog.FromContext(ctx)

	// start informer factories
	logger.Info("starting informers")
	s.Config.KubeInformers.Start(ctx.Done())
	s.Config.BindInformers.Start(ctx.Done())
	s.Config.ApiextensionsInformers.Start(ctx.Done())
	kubeSynced := s.Config.KubeInformers.WaitForCacheSync(ctx.Done())
	kubeBindSynced := s.Config.BindInformers.WaitForCacheSync(ctx.Done())
	apiextensionsSynced := s.Config.ApiextensionsInformers.WaitForCacheSync(ctx.Done())

	logger.Info("local informers are synced",
		"kubeSynced", fmt.Sprintf("%v", kubeSynced),
		"kubeBindSynced", fmt.Sprintf("%v", kubeBindSynced),
		"apiextensionsSynced", fmt.Sprintf("%v", apiextensionsSynced),
	)
}

func (s Prepared) Run(ctx context.Context) error {
	s.Controller.Start(ctx, 2)
	return nil
}
