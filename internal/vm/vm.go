package vm

import (
	"github.com/robertkrimen/otto"
	"github.com/sebvautour/event-manager/internal/service"
	"github.com/sebvautour/event-manager/internal/vmmethods"
	"github.com/sebvautour/event-manager/pkg/config"
)

type VM struct {
	s          *service.Service
	scriptList config.ScriptList
	scripts    map[string]config.Script
	vm         *otto.Otto
}

func New(s *service.Service) *VM {
	return &VM{
		s: s,
		scriptList: config.ScriptList{
			Scripts: []string{"one", "two", "three"},
		},
		scripts: map[string]config.Script{
			"one": {
				Script: `m.LogInfo("hello from script one");`,
			},
			"two": {
				Script: `m.LogInfo("msg: " + m.Message);`,
			},
			"three": {
				Script: `m.LogInfo("goodby from script three");`,
			},
		},
		vm: otto.New(),
	}
}

func (vm *VM) Run(msg interface{}) {

	for _, scriptName := range vm.scriptList.Scripts {
		log := vm.s.Log.WithField("script", scriptName)

		m := vmmethods.New(vm.s, log, msg)
		vm.vm.Set("m", m)

		n, ok := vm.scripts[scriptName]
		if !ok {
			log.Error("No script found")
			continue
		}

		if m.IsDropped() {
			log.Debug("message dropped")
			return
		}

		vm.vm.Set("log", log)

		_, err := vm.vm.Run(n.Script)
		if err != nil {
			log.Error(err.Error())
		}
	}
}
