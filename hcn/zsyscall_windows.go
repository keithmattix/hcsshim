//go:build windows

// Code generated by 'go generate' using "github.com/Microsoft/go-winio/tools/mkwinsyscall"; DO NOT EDIT.

package hcn

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	return e
}

var (
	modcomputenetwork = windows.NewLazySystemDLL("computenetwork.dll")
	modiphlpapi       = windows.NewLazySystemDLL("iphlpapi.dll")
	modvmcompute      = windows.NewLazySystemDLL("vmcompute.dll")

	procHcnCloseEndpoint               = modcomputenetwork.NewProc("HcnCloseEndpoint")
	procHcnCloseLoadBalancer           = modcomputenetwork.NewProc("HcnCloseLoadBalancer")
	procHcnCloseNamespace              = modcomputenetwork.NewProc("HcnCloseNamespace")
	procHcnCloseNetwork                = modcomputenetwork.NewProc("HcnCloseNetwork")
	procHcnCloseSdnRoute               = modcomputenetwork.NewProc("HcnCloseSdnRoute")
	procHcnCreateEndpoint              = modcomputenetwork.NewProc("HcnCreateEndpoint")
	procHcnCreateLoadBalancer          = modcomputenetwork.NewProc("HcnCreateLoadBalancer")
	procHcnCreateNamespace             = modcomputenetwork.NewProc("HcnCreateNamespace")
	procHcnCreateNetwork               = modcomputenetwork.NewProc("HcnCreateNetwork")
	procHcnCreateSdnRoute              = modcomputenetwork.NewProc("HcnCreateSdnRoute")
	procHcnDeleteEndpoint              = modcomputenetwork.NewProc("HcnDeleteEndpoint")
	procHcnDeleteLoadBalancer          = modcomputenetwork.NewProc("HcnDeleteLoadBalancer")
	procHcnDeleteNamespace             = modcomputenetwork.NewProc("HcnDeleteNamespace")
	procHcnDeleteNetwork               = modcomputenetwork.NewProc("HcnDeleteNetwork")
	procHcnDeleteSdnRoute              = modcomputenetwork.NewProc("HcnDeleteSdnRoute")
	procHcnEnumerateEndpoints          = modcomputenetwork.NewProc("HcnEnumerateEndpoints")
	procHcnEnumerateLoadBalancers      = modcomputenetwork.NewProc("HcnEnumerateLoadBalancers")
	procHcnEnumerateNamespaces         = modcomputenetwork.NewProc("HcnEnumerateNamespaces")
	procHcnEnumerateNetworks           = modcomputenetwork.NewProc("HcnEnumerateNetworks")
	procHcnEnumerateSdnRoutes          = modcomputenetwork.NewProc("HcnEnumerateSdnRoutes")
	procHcnModifyEndpoint              = modcomputenetwork.NewProc("HcnModifyEndpoint")
	procHcnModifyLoadBalancer          = modcomputenetwork.NewProc("HcnModifyLoadBalancer")
	procHcnModifyNamespace             = modcomputenetwork.NewProc("HcnModifyNamespace")
	procHcnModifyNetwork               = modcomputenetwork.NewProc("HcnModifyNetwork")
	procHcnModifySdnRoute              = modcomputenetwork.NewProc("HcnModifySdnRoute")
	procHcnOpenEndpoint                = modcomputenetwork.NewProc("HcnOpenEndpoint")
	procHcnOpenLoadBalancer            = modcomputenetwork.NewProc("HcnOpenLoadBalancer")
	procHcnOpenNamespace               = modcomputenetwork.NewProc("HcnOpenNamespace")
	procHcnOpenNetwork                 = modcomputenetwork.NewProc("HcnOpenNetwork")
	procHcnOpenSdnRoute                = modcomputenetwork.NewProc("HcnOpenSdnRoute")
	procHcnQueryEndpointProperties     = modcomputenetwork.NewProc("HcnQueryEndpointProperties")
	procHcnQueryLoadBalancerProperties = modcomputenetwork.NewProc("HcnQueryLoadBalancerProperties")
	procHcnQueryNamespaceProperties    = modcomputenetwork.NewProc("HcnQueryNamespaceProperties")
	procHcnQueryNetworkProperties      = modcomputenetwork.NewProc("HcnQueryNetworkProperties")
	procHcnQuerySdnRouteProperties     = modcomputenetwork.NewProc("HcnQuerySdnRouteProperties")
	procGetCurrentThreadCompartmentId  = modiphlpapi.NewProc("GetCurrentThreadCompartmentId")
	procSetCurrentThreadCompartmentId  = modiphlpapi.NewProc("SetCurrentThreadCompartmentId")
	procHNSCall                        = modvmcompute.NewProc("HNSCall")
)

func hcnCloseEndpoint(endpoint hcnEndpoint) (hr error) {
	hr = procHcnCloseEndpoint.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnCloseEndpoint.Addr(), uintptr(endpoint))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnCloseLoadBalancer(loadBalancer hcnLoadBalancer) (hr error) {
	hr = procHcnCloseLoadBalancer.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnCloseLoadBalancer.Addr(), uintptr(loadBalancer))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnCloseNamespace(namespace hcnNamespace) (hr error) {
	hr = procHcnCloseNamespace.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnCloseNamespace.Addr(), uintptr(namespace))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnCloseNetwork(network hcnNetwork) (hr error) {
	hr = procHcnCloseNetwork.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnCloseNetwork.Addr(), uintptr(network))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnCloseRoute(route hcnRoute) (hr error) {
	hr = procHcnCloseSdnRoute.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnCloseSdnRoute.Addr(), uintptr(route))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnCreateEndpoint(network hcnNetwork, id *_guid, settings string, endpoint *hcnEndpoint, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(settings)
	if hr != nil {
		return
	}
	return _hcnCreateEndpoint(network, id, _p0, endpoint, result)
}

func _hcnCreateEndpoint(network hcnNetwork, id *_guid, settings *uint16, endpoint *hcnEndpoint, result **uint16) (hr error) {
	hr = procHcnCreateEndpoint.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnCreateEndpoint.Addr(), uintptr(network), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(endpoint)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnCreateLoadBalancer(id *_guid, settings string, loadBalancer *hcnLoadBalancer, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(settings)
	if hr != nil {
		return
	}
	return _hcnCreateLoadBalancer(id, _p0, loadBalancer, result)
}

func _hcnCreateLoadBalancer(id *_guid, settings *uint16, loadBalancer *hcnLoadBalancer, result **uint16) (hr error) {
	hr = procHcnCreateLoadBalancer.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnCreateLoadBalancer.Addr(), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(loadBalancer)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnCreateNamespace(id *_guid, settings string, namespace *hcnNamespace, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(settings)
	if hr != nil {
		return
	}
	return _hcnCreateNamespace(id, _p0, namespace, result)
}

func _hcnCreateNamespace(id *_guid, settings *uint16, namespace *hcnNamespace, result **uint16) (hr error) {
	hr = procHcnCreateNamespace.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnCreateNamespace.Addr(), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(namespace)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnCreateNetwork(id *_guid, settings string, network *hcnNetwork, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(settings)
	if hr != nil {
		return
	}
	return _hcnCreateNetwork(id, _p0, network, result)
}

func _hcnCreateNetwork(id *_guid, settings *uint16, network *hcnNetwork, result **uint16) (hr error) {
	hr = procHcnCreateNetwork.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnCreateNetwork.Addr(), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(network)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnCreateRoute(id *_guid, settings string, route *hcnRoute, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(settings)
	if hr != nil {
		return
	}
	return _hcnCreateRoute(id, _p0, route, result)
}

func _hcnCreateRoute(id *_guid, settings *uint16, route *hcnRoute, result **uint16) (hr error) {
	hr = procHcnCreateSdnRoute.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnCreateSdnRoute.Addr(), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(route)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnDeleteEndpoint(id *_guid, result **uint16) (hr error) {
	hr = procHcnDeleteEndpoint.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnDeleteEndpoint.Addr(), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnDeleteLoadBalancer(id *_guid, result **uint16) (hr error) {
	hr = procHcnDeleteLoadBalancer.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnDeleteLoadBalancer.Addr(), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnDeleteNamespace(id *_guid, result **uint16) (hr error) {
	hr = procHcnDeleteNamespace.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnDeleteNamespace.Addr(), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnDeleteNetwork(id *_guid, result **uint16) (hr error) {
	hr = procHcnDeleteNetwork.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnDeleteNetwork.Addr(), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnDeleteRoute(id *_guid, result **uint16) (hr error) {
	hr = procHcnDeleteSdnRoute.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnDeleteSdnRoute.Addr(), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnEnumerateEndpoints(query string, endpoints **uint16, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(query)
	if hr != nil {
		return
	}
	return _hcnEnumerateEndpoints(_p0, endpoints, result)
}

func _hcnEnumerateEndpoints(query *uint16, endpoints **uint16, result **uint16) (hr error) {
	hr = procHcnEnumerateEndpoints.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnEnumerateEndpoints.Addr(), uintptr(unsafe.Pointer(query)), uintptr(unsafe.Pointer(endpoints)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnEnumerateLoadBalancers(query string, loadBalancers **uint16, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(query)
	if hr != nil {
		return
	}
	return _hcnEnumerateLoadBalancers(_p0, loadBalancers, result)
}

func _hcnEnumerateLoadBalancers(query *uint16, loadBalancers **uint16, result **uint16) (hr error) {
	hr = procHcnEnumerateLoadBalancers.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnEnumerateLoadBalancers.Addr(), uintptr(unsafe.Pointer(query)), uintptr(unsafe.Pointer(loadBalancers)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnEnumerateNamespaces(query string, namespaces **uint16, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(query)
	if hr != nil {
		return
	}
	return _hcnEnumerateNamespaces(_p0, namespaces, result)
}

func _hcnEnumerateNamespaces(query *uint16, namespaces **uint16, result **uint16) (hr error) {
	hr = procHcnEnumerateNamespaces.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnEnumerateNamespaces.Addr(), uintptr(unsafe.Pointer(query)), uintptr(unsafe.Pointer(namespaces)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnEnumerateNetworks(query string, networks **uint16, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(query)
	if hr != nil {
		return
	}
	return _hcnEnumerateNetworks(_p0, networks, result)
}

func _hcnEnumerateNetworks(query *uint16, networks **uint16, result **uint16) (hr error) {
	hr = procHcnEnumerateNetworks.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnEnumerateNetworks.Addr(), uintptr(unsafe.Pointer(query)), uintptr(unsafe.Pointer(networks)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnEnumerateRoutes(query string, routes **uint16, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(query)
	if hr != nil {
		return
	}
	return _hcnEnumerateRoutes(_p0, routes, result)
}

func _hcnEnumerateRoutes(query *uint16, routes **uint16, result **uint16) (hr error) {
	hr = procHcnEnumerateSdnRoutes.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnEnumerateSdnRoutes.Addr(), uintptr(unsafe.Pointer(query)), uintptr(unsafe.Pointer(routes)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnModifyEndpoint(endpoint hcnEndpoint, settings string, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(settings)
	if hr != nil {
		return
	}
	return _hcnModifyEndpoint(endpoint, _p0, result)
}

func _hcnModifyEndpoint(endpoint hcnEndpoint, settings *uint16, result **uint16) (hr error) {
	hr = procHcnModifyEndpoint.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnModifyEndpoint.Addr(), uintptr(endpoint), uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnModifyLoadBalancer(loadBalancer hcnLoadBalancer, settings string, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(settings)
	if hr != nil {
		return
	}
	return _hcnModifyLoadBalancer(loadBalancer, _p0, result)
}

func _hcnModifyLoadBalancer(loadBalancer hcnLoadBalancer, settings *uint16, result **uint16) (hr error) {
	hr = procHcnModifyLoadBalancer.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnModifyLoadBalancer.Addr(), uintptr(loadBalancer), uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnModifyNamespace(namespace hcnNamespace, settings string, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(settings)
	if hr != nil {
		return
	}
	return _hcnModifyNamespace(namespace, _p0, result)
}

func _hcnModifyNamespace(namespace hcnNamespace, settings *uint16, result **uint16) (hr error) {
	hr = procHcnModifyNamespace.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnModifyNamespace.Addr(), uintptr(namespace), uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnModifyNetwork(network hcnNetwork, settings string, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(settings)
	if hr != nil {
		return
	}
	return _hcnModifyNetwork(network, _p0, result)
}

func _hcnModifyNetwork(network hcnNetwork, settings *uint16, result **uint16) (hr error) {
	hr = procHcnModifyNetwork.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnModifyNetwork.Addr(), uintptr(network), uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnModifyRoute(route hcnRoute, settings string, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(settings)
	if hr != nil {
		return
	}
	return _hcnModifyRoute(route, _p0, result)
}

func _hcnModifyRoute(route hcnRoute, settings *uint16, result **uint16) (hr error) {
	hr = procHcnModifySdnRoute.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnModifySdnRoute.Addr(), uintptr(route), uintptr(unsafe.Pointer(settings)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnOpenEndpoint(id *_guid, endpoint *hcnEndpoint, result **uint16) (hr error) {
	hr = procHcnOpenEndpoint.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnOpenEndpoint.Addr(), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(endpoint)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnOpenLoadBalancer(id *_guid, loadBalancer *hcnLoadBalancer, result **uint16) (hr error) {
	hr = procHcnOpenLoadBalancer.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnOpenLoadBalancer.Addr(), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(loadBalancer)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnOpenNamespace(id *_guid, namespace *hcnNamespace, result **uint16) (hr error) {
	hr = procHcnOpenNamespace.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnOpenNamespace.Addr(), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(namespace)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnOpenNetwork(id *_guid, network *hcnNetwork, result **uint16) (hr error) {
	hr = procHcnOpenNetwork.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnOpenNetwork.Addr(), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(network)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnOpenRoute(id *_guid, route *hcnRoute, result **uint16) (hr error) {
	hr = procHcnOpenSdnRoute.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnOpenSdnRoute.Addr(), uintptr(unsafe.Pointer(id)), uintptr(unsafe.Pointer(route)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnQueryEndpointProperties(endpoint hcnEndpoint, query string, properties **uint16, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(query)
	if hr != nil {
		return
	}
	return _hcnQueryEndpointProperties(endpoint, _p0, properties, result)
}

func _hcnQueryEndpointProperties(endpoint hcnEndpoint, query *uint16, properties **uint16, result **uint16) (hr error) {
	hr = procHcnQueryEndpointProperties.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnQueryEndpointProperties.Addr(), uintptr(endpoint), uintptr(unsafe.Pointer(query)), uintptr(unsafe.Pointer(properties)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnQueryLoadBalancerProperties(loadBalancer hcnLoadBalancer, query string, properties **uint16, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(query)
	if hr != nil {
		return
	}
	return _hcnQueryLoadBalancerProperties(loadBalancer, _p0, properties, result)
}

func _hcnQueryLoadBalancerProperties(loadBalancer hcnLoadBalancer, query *uint16, properties **uint16, result **uint16) (hr error) {
	hr = procHcnQueryLoadBalancerProperties.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnQueryLoadBalancerProperties.Addr(), uintptr(loadBalancer), uintptr(unsafe.Pointer(query)), uintptr(unsafe.Pointer(properties)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnQueryNamespaceProperties(namespace hcnNamespace, query string, properties **uint16, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(query)
	if hr != nil {
		return
	}
	return _hcnQueryNamespaceProperties(namespace, _p0, properties, result)
}

func _hcnQueryNamespaceProperties(namespace hcnNamespace, query *uint16, properties **uint16, result **uint16) (hr error) {
	hr = procHcnQueryNamespaceProperties.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnQueryNamespaceProperties.Addr(), uintptr(namespace), uintptr(unsafe.Pointer(query)), uintptr(unsafe.Pointer(properties)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnQueryNetworkProperties(network hcnNetwork, query string, properties **uint16, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(query)
	if hr != nil {
		return
	}
	return _hcnQueryNetworkProperties(network, _p0, properties, result)
}

func _hcnQueryNetworkProperties(network hcnNetwork, query *uint16, properties **uint16, result **uint16) (hr error) {
	hr = procHcnQueryNetworkProperties.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnQueryNetworkProperties.Addr(), uintptr(network), uintptr(unsafe.Pointer(query)), uintptr(unsafe.Pointer(properties)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func hcnQueryRouteProperties(route hcnRoute, query string, properties **uint16, result **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(query)
	if hr != nil {
		return
	}
	return _hcnQueryRouteProperties(route, _p0, properties, result)
}

func _hcnQueryRouteProperties(route hcnRoute, query *uint16, properties **uint16, result **uint16) (hr error) {
	hr = procHcnQuerySdnRouteProperties.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHcnQuerySdnRouteProperties.Addr(), uintptr(route), uintptr(unsafe.Pointer(query)), uintptr(unsafe.Pointer(properties)), uintptr(unsafe.Pointer(result)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func GetCurrentThreadCompartmentId() (compartmentId uint32) {
	r0, _, _ := syscall.SyscallN(procGetCurrentThreadCompartmentId.Addr())
	compartmentId = uint32(r0)
	return
}

func SetCurrentThreadCompartmentId(compartmentId uint32) (hr error) {
	r0, _, _ := syscall.SyscallN(procSetCurrentThreadCompartmentId.Addr(), uintptr(compartmentId))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}

func _hnsCall(method string, path string, object string, response **uint16) (hr error) {
	var _p0 *uint16
	_p0, hr = syscall.UTF16PtrFromString(method)
	if hr != nil {
		return
	}
	var _p1 *uint16
	_p1, hr = syscall.UTF16PtrFromString(path)
	if hr != nil {
		return
	}
	var _p2 *uint16
	_p2, hr = syscall.UTF16PtrFromString(object)
	if hr != nil {
		return
	}
	return __hnsCall(_p0, _p1, _p2, response)
}

func __hnsCall(method *uint16, path *uint16, object *uint16, response **uint16) (hr error) {
	hr = procHNSCall.Find()
	if hr != nil {
		return
	}
	r0, _, _ := syscall.SyscallN(procHNSCall.Addr(), uintptr(unsafe.Pointer(method)), uintptr(unsafe.Pointer(path)), uintptr(unsafe.Pointer(object)), uintptr(unsafe.Pointer(response)))
	if int32(r0) < 0 {
		if r0&0x1fff0000 == 0x00070000 {
			r0 &= 0xffff
		}
		hr = syscall.Errno(r0)
	}
	return
}
