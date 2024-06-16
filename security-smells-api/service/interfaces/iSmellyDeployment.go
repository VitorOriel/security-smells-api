package interfaces

type ISmellyDeployment interface {
	ISmellyKubernetes
}

type SmellyDeployment struct {
	ISmellyDeployment ISmellyDeployment
}

func (smellyDeployment *SmellyDeployment) CheckSmelly() {
	smellyDeployment.ISmellyDeployment.SmellyResourceAndLimit()
	smellyDeployment.ISmellyDeployment.SmellySecurityContextRunAsUser()
	smellyDeployment.ISmellyDeployment.SmellySecurityContextCapabilities()
	smellyDeployment.ISmellyDeployment.SmellySecurityContextAllowPrivilegeEscalation()
	smellyDeployment.ISmellyDeployment.SmellySecurityContextReadOnlyRootFilesystem()
	smellyDeployment.ISmellyDeployment.SmellySecurityContextPrivileged()
}
