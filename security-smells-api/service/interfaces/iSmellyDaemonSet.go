package interfaces

type ISmellyDaemonSet interface {
	SmellyResourceAndLimit()
	SmellySecurityContextRunAsUser()
	SmellySecurityContextCapabilities()
	SmellySecurityContextAllowPrivilegeEscalation()
	SmellySecurityContextReadOnlyRootFilesystem()
}

type SmellyDaemonSet struct {
	ISmellyDaemonSet ISmellyDaemonSet
}

func (smellyDaemonSet *SmellyDaemonSet) checkSmellyDemonSet() {
	smellyDaemonSet.ISmellyDaemonSet.SmellyResourceAndLimit()
	smellyDaemonSet.ISmellyDaemonSet.SmellySecurityContextRunAsUser()
	smellyDaemonSet.ISmellyDaemonSet.SmellySecurityContextCapabilities()
	smellyDaemonSet.ISmellyDaemonSet.SmellySecurityContextAllowPrivilegeEscalation()
	smellyDaemonSet.ISmellyDaemonSet.SmellySecurityContextReadOnlyRootFilesystem()
}
