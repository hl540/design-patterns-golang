package mediator

type StationManager struct {
	// 站台是否可以停靠
	isPlatformFree bool
	// 等待进入站台的列车
	trainQueue []Train
}

func newStationManager() *StationManager {
	return &StationManager{
		// 默认站台可以停靠
		isPlatformFree: true,
	}
}

// 询问站台是否可以停靠，如果不可以停靠将列车放入迭代队列
func (s *StationManager) canArrive(train Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, train)
	return false
}

// 列车离开通知，站台允许等待中的列车停靠
func (s *StationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.permitArrive()
	}
}
