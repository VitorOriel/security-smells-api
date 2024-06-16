package interfaces

type ISmellyDaemonSet interface {
	ISmellyKubernetes
}

type SmellyDaemonSet struct {
	ISmellyDaemonSet ISmellyDaemonSet
}

func (smellyDaemonSet *SmellyDaemonSet) CheckSmelly() {
	smellyDaemonSet.ISmellyDaemonSet.SmellyResourceAndLimit()
	smellyDaemonSet.ISmellyDaemonSet.SmellySecurityContextRunAsUser()
	smellyDaemonSet.ISmellyDaemonSet.SmellySecurityContextCapabilities()
	smellyDaemonSet.ISmellyDaemonSet.SmellySecurityContextAllowPrivilegeEscalation()
	smellyDaemonSet.ISmellyDaemonSet.SmellySecurityContextReadOnlyRootFilesystem()
	smellyDaemonSet.ISmellyDaemonSet.SmellySecurityContextPrivileged()
}
