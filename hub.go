package gologger

import "sync"

var (
	__loggers          map[string]*Logger = map[string]*Logger{}
	__loggers_mux      sync.Mutex
	__global_behaviors map[string]IBehavior = make(map[string]IBehavior, 0)
	__global_mux       sync.Mutex
	__behaviors        map[string]map[string]IBehavior = map[string]map[string]IBehavior{}
	__behaviors_mux    sync.Mutex
)

func GetGlobalLogger(agent string) *Logger {

	if logger, ok := __loggers[agent]; ok {

		return logger
	}
	logger := NewLogger(agent)
	behaviors := map[string]IBehavior{}
	__global_mux.Lock()
	for name, behave := range __global_behaviors {
		behaviors[name] = behave
	}
	__global_mux.Unlock()
	__behaviors_mux.Lock()
	if agentBehaviors, ok := __behaviors[agent]; ok {

		for name, behave := range agentBehaviors {

			behaviors[name] = behave
		}
	}
	__behaviors_mux.Unlock()
	__loggers_mux.Lock()
	for _, behave := range behaviors {

		logger.AddBehavior(behave)
	}
	__loggers[agent] = logger
	__loggers_mux.Unlock()
	return logger
}

//AddGlobalBehavior add global behavior using for all logger
func AddGlobalBehavior(behavior IBehavior) {
	__global_mux.Lock()

	__global_behaviors[behavior.GetName()] = behavior
	__global_mux.Unlock()

	__loggers_mux.Lock()
	for _, logger := range __loggers {
		logger.AddBehavior(behavior)
	}
	__loggers_mux.Unlock()

}
func RemoveGlobalBehavior(behavior IBehavior) {
	behaviorName := behavior.GetName()
	__global_mux.Lock()

	delete(__global_behaviors, behaviorName)

	__global_mux.Unlock()
	__behaviors_mux.Lock()

	__loggers_mux.Lock()
	for agentName, agentBehaviors := range __behaviors {
		if _, has := agentBehaviors[behaviorName]; !has {
			if agent, hasAgent := __loggers[agentName]; hasAgent {
				agent.RemoveBehavior(behavior)
			}
		}
	}
	__loggers_mux.Unlock()
	__behaviors_mux.Unlock()
}

//AddBehavior add a behavior to agent
func AddBehavior(agent string, behavior IBehavior) {
	__behaviors_mux.Lock()
	if _, ok := __behaviors[agent]; !ok {

		__behaviors[agent] = map[string]IBehavior{}
	}
	__behaviors[agent][behavior.GetName()] = behavior
	__behaviors_mux.Unlock()

	__loggers_mux.Lock()
	if logger, ok := __loggers[agent]; ok {

		logger.AddBehavior(behavior)
	}
	__loggers_mux.Unlock()
}

func RemoveBehavior(agent string, behavior IBehavior) {
	__behaviors_mux.Lock()
	if _, ok := __behaviors[agent]; ok {

		delete(__behaviors[agent], behavior.GetName())
	}
	__behaviors_mux.Unlock()
	__global_mux.Lock()
	globalBehave, hasGlobal := __global_behaviors[behavior.GetName()]

	__global_mux.Unlock()
	__loggers_mux.Lock()

	if logger, ok := __loggers[agent]; ok {
		if hasGlobal {
			logger.AddBehavior(globalBehave)
		} else {
			logger.RemoveBehavior(behavior)
		}
	}
	__loggers_mux.Unlock()
}
