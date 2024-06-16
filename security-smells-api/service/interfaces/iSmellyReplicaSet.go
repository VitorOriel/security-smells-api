package interfaces

type ISmellyReplicaSet interface {
	ISmellyKubernetes
}

type SmellyReplicaSet struct {
	ISmellyReplicaSet ISmellyReplicaSet
}

func (smellyReplicaSet *SmellyReplicaSet) CheckSmelly() {
	smellyReplicaSet.ISmellyReplicaSet.SmellyResourceAndLimit()
	smellyReplicaSet.ISmellyReplicaSet.SmellySecurityContextRunAsUser()
	smellyReplicaSet.ISmellyReplicaSet.SmellySecurityContextCapabilities()
	smellyReplicaSet.ISmellyReplicaSet.SmellySecurityContextAllowPrivilegeEscalation()
	smellyReplicaSet.ISmellyReplicaSet.SmellySecurityContextReadOnlyRootFilesystem()
	smellyReplicaSet.ISmellyReplicaSet.SmellySecurityContextPrivileged()
}
