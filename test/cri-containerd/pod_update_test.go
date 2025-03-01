//go:build windows && functional
// +build windows,functional

package cri_containerd

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/Microsoft/hcsshim/internal/memory"
	"github.com/Microsoft/hcsshim/osversion"
	"github.com/Microsoft/hcsshim/pkg/annotations"
	"github.com/Microsoft/hcsshim/test/pkg/definitions/cpugroup"
	"github.com/Microsoft/hcsshim/test/pkg/definitions/processorinfo"
	"github.com/Microsoft/hcsshim/test/pkg/require"
	runtime "k8s.io/cri-api/pkg/apis/runtime/v1"
)

func Test_Pod_UpdateResources_Memory(t *testing.T) {
	requireAnyFeature(t, featureWCOWHypervisor)

	type config struct {
		name             string
		requiredFeatures []string
		runtimeHandler   string
		sandboxImage     string
	}
	tests := []config{
		{
			name:             "WCOW_Hypervisor",
			requiredFeatures: []string{featureWCOWHypervisor},
			runtimeHandler:   wcowHypervisorRuntimeHandler,
			sandboxImage:     imageWindowsNanoserver,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			requireFeatures(t, test.requiredFeatures...)
			pullRequiredImages(t, []string{test.sandboxImage})
			var startingMemorySize int64 = 768 * memory.MiB
			podRequest := getRunPodSandboxRequest(
				t,
				test.runtimeHandler,
				WithSandboxAnnotations(map[string]string{
					annotations.ContainerMemorySizeInMB: fmt.Sprintf("%d", startingMemorySize),
				}),
			)

			client := newTestRuntimeClient(t)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			podID := runPodSandbox(t, client, ctx, podRequest)
			defer removePodSandbox(t, client, ctx, podID)
			defer stopPodSandbox(t, client, ctx, podID)

			// make request for shrinking memory size
			newMemorySize := startingMemorySize / 2
			updateReq := &runtime.UpdateContainerResourcesRequest{
				ContainerId: podID,
			}

			updateReq.Windows = &runtime.WindowsContainerResources{
				MemoryLimitInBytes: newMemorySize,
			}

			if _, err := client.UpdateContainerResources(ctx, updateReq); err != nil {
				t.Fatalf("updating container resources for %s with %v", podID, err)
			}
		})
	}
}

func Test_Pod_UpdateResources_Memory_PA(t *testing.T) {
	requireAnyFeature(t, featureWCOWHypervisor)

	type config struct {
		name             string
		requiredFeatures []string
		runtimeHandler   string
		sandboxImage     string
	}
	tests := []config{
		{
			name:             "WCOW_Hypervisor",
			requiredFeatures: []string{featureWCOWHypervisor},
			runtimeHandler:   wcowHypervisorRuntimeHandler,
			sandboxImage:     imageWindowsNanoserver,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			requireFeatures(t, test.requiredFeatures...)
			pullRequiredImages(t, []string{test.sandboxImage})

			var startingMemorySize int64 = 200 * memory.MiB
			podRequest := getRunPodSandboxRequest(
				t,
				test.runtimeHandler,
				WithSandboxAnnotations(map[string]string{
					annotations.FullyPhysicallyBacked:   "true",
					annotations.ContainerMemorySizeInMB: fmt.Sprintf("%d", startingMemorySize),
				}),
			)

			client := newTestRuntimeClient(t)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			podID := runPodSandbox(t, client, ctx, podRequest)
			defer removePodSandbox(t, client, ctx, podID)
			defer stopPodSandbox(t, client, ctx, podID)

			// make request for shrinking memory size
			newMemorySize := startingMemorySize / 2
			updateReq := &runtime.UpdateContainerResourcesRequest{
				ContainerId: podID,
			}

			updateReq.Windows = &runtime.WindowsContainerResources{
				MemoryLimitInBytes: newMemorySize,
			}

			if _, err := client.UpdateContainerResources(ctx, updateReq); err != nil {
				t.Fatalf("updating container resources for %s with %v", podID, err)
			}
		})
	}
}

func Test_Pod_UpdateResources_CPUShares(t *testing.T) {
	requireAnyFeature(t, featureWCOWHypervisor)
	require.Build(t, osversion.V20H2)

	type config struct {
		name             string
		requiredFeatures []string
		runtimeHandler   string
		sandboxImage     string
	}
	tests := []config{
		{
			name:             "WCOW_Hypervisor",
			requiredFeatures: []string{featureWCOWHypervisor},
			runtimeHandler:   wcowHypervisorRuntimeHandler,
			sandboxImage:     imageWindowsNanoserver,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			requireFeatures(t, test.requiredFeatures...)
			pullRequiredImages(t, []string{test.sandboxImage})

			podRequest := getRunPodSandboxRequest(t, test.runtimeHandler)

			client := newTestRuntimeClient(t)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			podID := runPodSandbox(t, client, ctx, podRequest)
			defer removePodSandbox(t, client, ctx, podID)
			defer stopPodSandbox(t, client, ctx, podID)

			updateReq := &runtime.UpdateContainerResourcesRequest{
				ContainerId: podID,
			}

			updateReq.Windows = &runtime.WindowsContainerResources{
				CpuShares: 2000,
			}

			if _, err := client.UpdateContainerResources(ctx, updateReq); err != nil {
				t.Fatalf("updating container resources for %s with %v", podID, err)
			}
		})
	}
}

func Test_Pod_UpdateResources_CPUGroup(t *testing.T) {
	t.Skip("Skipping for now")
	requireAnyFeature(t, featureWCOWHypervisor)

	ctx := context.Background()

	processorTopology, err := processorinfo.HostProcessorInfo(ctx)
	if err != nil {
		t.Fatalf("failed to get host processor information: %s", err)
	}
	lpIndices := make([]uint32, processorTopology.LogicalProcessorCount)
	for i, p := range processorTopology.LogicalProcessors {
		lpIndices[i] = p.LpIndex
	}

	startCPUGroupID := "FA22A12C-36B3-486D-A3E9-BC526C2B450B"
	if err := cpugroup.Create(ctx, startCPUGroupID, lpIndices); err != nil {
		t.Fatalf("failed to create test cpugroup with: %v", err)
	}

	defer func() {
		err := cpugroup.Delete(ctx, startCPUGroupID)
		if err != nil && !errors.Is(err, cpugroup.ErrHVStatusInvalidCPUGroupState) {
			t.Fatalf("failed to clean up test cpugroup with: %v", err)
		}
	}()

	updateCPUGroupID := "FA22A12C-36B3-486D-A3E9-BC526C2B450C"
	if err := cpugroup.Create(ctx, updateCPUGroupID, lpIndices); err != nil {
		t.Fatalf("failed to create test cpugroup with: %v", err)
	}

	defer func() {
		err := cpugroup.Delete(ctx, updateCPUGroupID)
		if err != nil && !errors.Is(err, cpugroup.ErrHVStatusInvalidCPUGroupState) {
			t.Fatalf("failed to clean up test cpugroup with: %v", err)
		}
	}()

	type config struct {
		name             string
		requiredFeatures []string
		runtimeHandler   string
		sandboxImage     string
	}

	tests := []config{
		{
			name:             "WCOW_Hypervisor",
			requiredFeatures: []string{featureWCOWHypervisor},
			runtimeHandler:   wcowHypervisorRuntimeHandler,
			sandboxImage:     imageWindowsNanoserver,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			requireFeatures(t, test.requiredFeatures...)
			pullRequiredImages(t, []string{test.sandboxImage})

			podRequest := getRunPodSandboxRequest(t, test.runtimeHandler, WithSandboxAnnotations(map[string]string{
				annotations.CPUGroupID: startCPUGroupID,
			}))
			client := newTestRuntimeClient(t)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			podID := runPodSandbox(t, client, ctx, podRequest)
			defer removePodSandbox(t, client, ctx, podID)
			defer stopPodSandbox(t, client, ctx, podID)

			updateReq := &runtime.UpdateContainerResourcesRequest{
				ContainerId: podID,
				Annotations: map[string]string{
					annotations.CPUGroupID: updateCPUGroupID,
				},
			}
			updateReq.Windows = &runtime.WindowsContainerResources{}

			if _, err := client.UpdateContainerResources(ctx, updateReq); err != nil {
				t.Fatalf("updating container resources for %s with %v", podID, err)
			}
		})
	}
}
