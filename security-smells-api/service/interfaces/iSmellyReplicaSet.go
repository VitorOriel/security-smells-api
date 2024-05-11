package interfaces

type ISmellyReplicaSet interface {
	SmellyResourceAndLimit()
	SmellySecurityContextRunAsUser()
	SmellySecurityContextCapabilities()
	SmellySecurityContextAllowPrivilegeEscalation()
	SmellySecurityContextReadOnlyRootFilesystem()
}

type SmellyReplicaSet struct {
	ISmellyReplicaSet ISmellyReplicaSet
}

func (smellyReplicaSet *SmellyReplicaSet) checkSmellyReplicaSet() {
	smellyReplicaSet.ISmellyReplicaSet.SmellyResourceAndLimit()
	smellyReplicaSet.ISmellyReplicaSet.SmellySecurityContextRunAsUser()
	smellyReplicaSet.ISmellyReplicaSet.SmellySecurityContextCapabilities()
	smellyReplicaSet.ISmellyReplicaSet.SmellySecurityContextAllowPrivilegeEscalation()
	smellyReplicaSet.ISmellyReplicaSet.SmellySecurityContextReadOnlyRootFilesystem()
}
