package interfaces

type ISmellyStatefulSet interface {
	SmellyResourceAndLimit()
	SmellySecurityContextRunAsUser()
	SmellySecurityContextCapabilities()
	SmellySecurityContextAllowPrivilegeEscalation()
	SmellySecurityContextReadOnlyRootFilesystem()
}

type SmellyStatefulSet struct {
	ISmellyStatefulSet ISmellyStatefulSet
}

func (smellyStatefulSet *SmellyStatefulSet) checkSmellyStatefulSet() {
	smellyStatefulSet.ISmellyStatefulSet.SmellyResourceAndLimit()
	smellyStatefulSet.ISmellyStatefulSet.SmellySecurityContextRunAsUser()
	smellyStatefulSet.ISmellyStatefulSet.SmellySecurityContextCapabilities()
	smellyStatefulSet.ISmellyStatefulSet.SmellySecurityContextAllowPrivilegeEscalation()
	smellyStatefulSet.ISmellyStatefulSet.SmellySecurityContextReadOnlyRootFilesystem()
}
