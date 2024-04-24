package interfaces

type ISmellyPod interface {
	SmellyResourceAndLimit()
	SmellySecurityContextRunAsUser()
	SmellySecurityContextCapabilities()
	SmellySecurityContextAllowPrivilegeEscalation()
	SmellySecurityContextReadOnlyRootFilesystem()
}

type SmellyPod struct {
	ISmellyPod ISmellyDeployment
}

func (smellyDeployment *SmellyPod) checkSmellyPod() {
	smellyDeployment.ISmellyPod.SmellyResourceAndLimit()
	smellyDeployment.ISmellyPod.SmellySecurityContextRunAsUser()
	smellyDeployment.ISmellyPod.SmellySecurityContextCapabilities()
	smellyDeployment.ISmellyPod.SmellySecurityContextAllowPrivilegeEscalation()
	smellyDeployment.ISmellyPod.SmellySecurityContextReadOnlyRootFilesystem()
}
