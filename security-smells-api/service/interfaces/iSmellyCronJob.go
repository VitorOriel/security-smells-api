package interfaces

package interfaces

type ISmellyCronJob interface {
	SmellyResourceAndLimit()
	SmellySecurityContextRunAsUser()
	SmellySecurityContextCapabilities()
	SmellySecurityContextAllowPrivilegeEscalation()
	SmellySecurityContextReadOnlyRootFilesystem()
}

type SmellyCronJob struct {
	ISmellyCronJob ISmellyCronJob
}

func (smellyCronJob *SmellyCronJob) checkSmellyCronJob() {
	smellyCronJob.ISmellyCronJob.SmellyResourceAndLimit()
	smellyCronJob.ISmellyCronJob.SmellySecurityContextRunAsUser()
	smellyCronJob.ISmellyCronJob.SmellySecurityContextCapabilities()
	smellyCronJob.ISmellyCronJob.SmellySecurityContextAllowPrivilegeEscalation()
	smellyCronJob.ISmellyCronJob.SmellySecurityContextReadOnlyRootFilesystem()
}
