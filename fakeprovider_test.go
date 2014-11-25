package fence

import (
	"fmt"
	"log"
	"time"
)

type FakeProvider struct {
	agents Agents
}

func NewFakeProvider() FenceProvider {
	return &FakeProvider{agents: make(Agents)}
}

func (p *FakeProvider) LoadAgents(timeout time.Duration) error {
	a := &Agent{
		Name: "agent01",
		Parameters: map[string]*Parameter{
			"param01": &Parameter{Name: "param01", ContentType: String},
			"param02": &Parameter{Name: "param03", ContentType: Boolean},
			"param03": &Parameter{Name: "param03",
				ContentType: String,
				HasOptions:  true,
				Options: []interface{}{
					"option01",
					"option02",
				},
			},
		},
		MultiplePorts:   true,
		DefaultAction:   Reboot,
		UnfenceAction:   On,
		UnfenceOnTarget: false,
		Actions: []Action{
			On,
			Off,
			Reboot,
		},
	}
	p.agents[a.Name] = a

	return nil
}

func (p *FakeProvider) GetAgents() (Agents, error) {
	return p.agents, nil
}

func (p *FakeProvider) GetAgent(name string) (*Agent, error) {
	a, ok := p.agents[name]
	if !ok {
		return nil, fmt.Errorf("Unknown agent: %s", name)
	}
	return a, nil
}

func (p *FakeProvider) Status(ac *AgentConfig, timeout time.Duration) (DeviceStatus, error) {
	return Ok, nil
}
func (p *FakeProvider) Monitor(ac *AgentConfig, timeout time.Duration) (DeviceStatus, error) {
	return Ok, nil
}
func (p *FakeProvider) List(ac *AgentConfig, timeout time.Duration) (PortList, error) {
	return PortList{
		PortName{Name: "device01", Alias: "alias01"},
		PortName{Name: "device02", Alias: "alias02"},
	}, nil
}
func (p *FakeProvider) Run(ac *AgentConfig, action Action, timeout time.Duration) error {
	_, err := p.GetAgent(ac.Name)
	if err != nil {
		log.Print("error: ", err)
		return err
	}
	return nil
}
