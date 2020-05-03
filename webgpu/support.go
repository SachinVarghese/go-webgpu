package webgpu

import "log"

func CheckBrowserSupport() bool {
	var platform = getPlatformInfo()
	var IsWebgpuSupported = GetWebGPUSupport()
	if IsWebgpuSupported {
		log.Println("Webgpu is supported on " + platform)
	} else {
		log.Println("Webgpu is not supported on " + platform)
	}
	return IsWebgpuSupported
}
