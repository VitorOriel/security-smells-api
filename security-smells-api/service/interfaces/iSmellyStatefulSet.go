package interfaces

type ISmellyStatefulSet interface {
	ISmellyKubernetes
}

type SmellyStatefulSet struct {
	ISmellyStatefulSet ISmellyStatefulSet
}

func (smellyStatefulSet *SmellyStatefulSet) CheckSmelly() {
	smellyStatefulSet.ISmellyStatefulSet.SmellyResourceAndLimit()
	smellyStatefulSet.ISmellyStatefulSet.SmellySecurityContextRunAsUser()
	smellyStatefulSet.ISmellyStatefulSet.SmellySecurityContextCapabilities()
	smellyStatefulSet.ISmellyStatefulSet.SmellySecurityContextAllowPrivilegeEscalation()
	smellyStatefulSet.ISmellyStatefulSet.SmellySecurityContextReadOnlyRootFilesystem()
	smellyStatefulSet.ISmellyStatefulSet.SmellySecurityContextPrivileged()
}
