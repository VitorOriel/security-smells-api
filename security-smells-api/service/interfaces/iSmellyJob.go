package interfaces

type ISmellyJob interface {
	SmellyResourceAndLimit()
	SmellySecurityContextRunAsUser()
	SmellySecurityContextCapabilities()
	SmellySecurityContextAllowPrivilegeEscalation()
	SmellySecurityContextReadOnlyRootFilesystem()
}

type SmellyJob struct {
	ISmellyJob ISmellyJob
}

func (smellyJob *SmellyJob) checkSmellyJob() {
	smellyJob.ISmellyJob.SmellyResourceAndLimit()
	smellyJob.ISmellyJob.SmellySecurityContextRunAsUser()
	smellyJob.ISmellyJob.SmellySecurityContextCapabilities()
	smellyJob.ISmellyJob.SmellySecurityContextAllowPrivilegeEscalation()
	smellyJob.ISmellyJob.SmellySecurityContextReadOnlyRootFilesystem()
}
