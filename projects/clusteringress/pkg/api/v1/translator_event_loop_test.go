// Code generated by solo-kit. DO NOT EDIT.

// +build solokit

package v1

import (
	"context"
	"sync"
	"time"

	github_com_solo_io_gloo_projects_clusteringress_pkg_api_external_knative "github.com/solo-io/gloo/projects/clusteringress/pkg/api/external/knative"
	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
)

var _ = Describe("TranslatorEventLoop", func() {
	var (
		namespace string
		emitter   TranslatorEmitter
		err       error
	)

	BeforeEach(func() {

		secretClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		secretClient, err := gloo_solo_io.NewSecretClient(secretClientFactory)
		Expect(err).NotTo(HaveOccurred())

		clusterIngressClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		clusterIngressClient, err := github_com_solo_io_gloo_projects_clusteringress_pkg_api_external_knative.NewClusterIngressClient(clusterIngressClientFactory)
		Expect(err).NotTo(HaveOccurred())

		emitter = NewTranslatorEmitter(secretClient, clusterIngressClient)
	})
	It("runs sync function on a new snapshot", func() {
		_, err = emitter.Secret().Write(gloo_solo_io.NewSecret(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		_, err = emitter.ClusterIngress().Write(github_com_solo_io_gloo_projects_clusteringress_pkg_api_external_knative.NewClusterIngress(namespace, "jerry"), clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
		sync := &mockTranslatorSyncer{}
		el := NewTranslatorEventLoop(emitter, sync)
		_, err := el.Run([]string{namespace}, clients.WatchOpts{})
		Expect(err).NotTo(HaveOccurred())
		Eventually(sync.Synced, 5*time.Second).Should(BeTrue())
	})
})

type mockTranslatorSyncer struct {
	synced bool
	mutex  sync.Mutex
}

func (s *mockTranslatorSyncer) Synced() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.synced
}

func (s *mockTranslatorSyncer) Sync(ctx context.Context, snap *TranslatorSnapshot) error {
	s.mutex.Lock()
	s.synced = true
	s.mutex.Unlock()
	return nil
}
