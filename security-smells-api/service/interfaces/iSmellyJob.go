package interfaces

type ISmellyJob interface {
	ISmellyKubernetes
}

type SmellyJob struct {
	ISmellyJob ISmellyJob
}

func (smellyJob *SmellyJob) CheckSmelly() {
	smellyJob.ISmellyJob.SmellyResourceAndLimit()
	smellyJob.ISmellyJob.SmellySecurityContextRunAsUser()
	smellyJob.ISmellyJob.SmellySecurityContextCapabilities()
	smellyJob.ISmellyJob.SmellySecurityContextAllowPrivilegeEscalation()
	smellyJob.ISmellyJob.SmellySecurityContextReadOnlyRootFilesystem()
	smellyJob.ISmellyJob.SmellySecurityContextPrivileged()
}
