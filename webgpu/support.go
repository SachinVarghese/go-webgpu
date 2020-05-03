package webgpu

import "log"

func CheckBrowserSupport() bool {
	var platform = getPlatformInfo()
	var IsWebgpuSupported = GetWebGPUSupport()
	if IsWebgpuSupported {
		log.Println("Webgpu supported on " + platform)
	} else {
		log.Println("Webgpu not supported on " + platform)
	}
	return IsWebgpuSupported
}
