package metrics

type Monitor interface {
	ObserverDuration(eventName string, startNs, endNs int64)
}

var localMonitor Monitor

func SetMonitor(m Monitor) {
	localMonitor = m
}

func ObserverDuration(eventName string, startNs, endNs int64) {
	if localMonitor != nil {
		localMonitor.ObserverDuration(eventName, startNs, endNs)
	}
}
