package webgpu

import (
	"syscall/js"
)

var navigator = js.Global().Get("navigator")

func getPlatformInfo() string {
	platform := navigator.Get("platform")
	return platform.String()
}

func GetWebGPUSupport() bool {
	gpu := navigator.Get("gpu")
	return gpu.Truthy()
}
