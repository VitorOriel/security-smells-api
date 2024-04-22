package interfaces

type ISmellyDeployment interface {
	SmellyResourceAndLimit()
	SmellySecurityContext()
}

type SmellyDeployment struct {
	ISmellyDeployment ISmellyDeployment
}

func (smellyDeployment *SmellyDeployment) checkSmellyDeployment() {
	smellyDeployment.ISmellyDeployment.SmellyResourceAndLimit()
	smellyDeployment.ISmellyDeployment.SmellySecurityContext()
}
