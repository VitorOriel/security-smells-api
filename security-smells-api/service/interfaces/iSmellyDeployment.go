package interfaces

type ISmellyDeployment interface {
	SmellyResourceAndLimit()
	SmellySecurityContextRunAsUser()
	SmellySecurityContextCapabilities()
	SmellySecurityContextAllowPrivilegeEscalation()
	SmellySecurityContextReadOnlyRootFilesystem()
}

type SmellyDeployment struct {
	ISmellyDeployment ISmellyDeployment
}

func (smellyDeployment *SmellyDeployment) checkSmellyDeployment() {
	smellyDeployment.ISmellyDeployment.SmellyResourceAndLimit()
	smellyDeployment.ISmellyDeployment.SmellySecurityContextRunAsUser()
	smellyDeployment.ISmellyDeployment.SmellySecurityContextCapabilities()
	smellyDeployment.ISmellyDeployment.SmellySecurityContextAllowPrivilegeEscalation()
	smellyDeployment.ISmellyDeployment.SmellySecurityContextReadOnlyRootFilesystem()
}
