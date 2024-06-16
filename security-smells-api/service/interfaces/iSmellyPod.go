package interfaces

type ISmellyPod interface {
	ISmellyKubernetes
}

type SmellyPod struct {
	ISmellyPod ISmellyDeployment
}

func (smellyPod *SmellyPod) CheckSmelly() {
	smellyPod.ISmellyPod.SmellyResourceAndLimit()
	smellyPod.ISmellyPod.SmellySecurityContextRunAsUser()
	smellyPod.ISmellyPod.SmellySecurityContextCapabilities()
	smellyPod.ISmellyPod.SmellySecurityContextAllowPrivilegeEscalation()
	smellyPod.ISmellyPod.SmellySecurityContextReadOnlyRootFilesystem()
	smellyPod.ISmellyPod.SmellySecurityContextPrivileged()
}
