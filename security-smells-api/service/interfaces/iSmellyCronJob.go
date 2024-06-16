package interfaces

type ISmellyCronJob interface {
	ISmellyKubernetes
}

type SmellyCronJob struct {
	ISmellyCronJob ISmellyCronJob
}

func (smellyCronJob *SmellyCronJob) CheckSmelly() {
	smellyCronJob.ISmellyCronJob.SmellyResourceAndLimit()
	smellyCronJob.ISmellyCronJob.SmellySecurityContextRunAsUser()
	smellyCronJob.ISmellyCronJob.SmellySecurityContextCapabilities()
	smellyCronJob.ISmellyCronJob.SmellySecurityContextAllowPrivilegeEscalation()
	smellyCronJob.ISmellyCronJob.SmellySecurityContextReadOnlyRootFilesystem()
	smellyCronJob.ISmellyCronJob.SmellySecurityContextPrivileged()
}
