package fence

import (
	"errors"
	"testing"
)

func TestVerifyAgentConfig(t *testing.T) {
	f := New()
	provider := NewFakeProvider()
	err := provider.LoadAgents(0)
	if err != nil {
		t.Error("error:", err)
	}
	f.RegisterProvider("fakeprovider", provider)

	ac := NewAgentConfig("fakeprovider", "missingagent01")
	err = f.VerifyAgentConfig(ac, false)
	if err == nil {
		t.Error(err)
	}

	ac = NewAgentConfig("fakeprovider", "agent01")
	ac.SetParameter("missingparam", "bla")
	err = f.VerifyAgentConfig(ac, false)
	if err == nil {
		t.Error(err)
	}

	ac = NewAgentConfig("fakeprovider", "agent01")
	ac.SetParameter("param01", "bla")
	err = f.VerifyAgentConfig(ac, false)
	if err != nil {
		t.Error(err)
	}

	ac = NewAgentConfig("fakeprovider", "agent01")
	ac.SetParameter("param02", "bla")
	err = f.VerifyAgentConfig(ac, false)
	if err == nil {
		t.Error(err)
	}

	ac = NewAgentConfig("fakeprovider", "agent01")
	ac.SetParameter("param03", "option02")
	err = f.VerifyAgentConfig(ac, false)
	if err != nil {
		t.Error(err)
	}

	ac = NewAgentConfig("fakeprovider", "agent01")
	ac.SetParameter("param03", "bla")
	err = f.VerifyAgentConfig(ac, false)
	if err == nil {
		t.Error(err)
	}

	ac = NewAgentConfig("fakeprovider", "agent01")
	ac.SetParameter("param03", 1)
	err = f.VerifyAgentConfig(ac, false)
	want := errors.New("Parameter \"param03\" not of string type")
	if !ErrorEquals(err, want) {
		t.Errorf("Expecting \"%s\" error, found \"%s\"", err, want)
	}

	ac = NewAgentConfig("fakeprovider", "agent01")
	ac.SetParameter("param01", "bla")
	err = f.VerifyAgentConfig(ac, true)
	want = errors.New("Port name required")
	if !ErrorEquals(err, want) {
		t.Errorf("Expecting \"%s\" error, found \"%s\"", err, want)
	}

	ac = NewAgentConfig("fakeprovider", "agent01")
	ac.SetParameter("param01", "bla")
	ac.SetPort("port01")
	err = f.VerifyAgentConfig(ac, true)
	if err != nil {
		t.Error(err)
	}

}

func TestRun(t *testing.T) {
	f := New()
	provider := NewFakeProvider()
	err := provider.LoadAgents(0)
	if err != nil {
		t.Error("error:", err)
	}
	f.RegisterProvider("fakeprovider", provider)

	ac := NewAgentConfig("fakeprovider", "agent01")
	ac.SetPort("port01")

	err = f.Run(ac, On, 0)
	if err != nil {
		t.Error(err)
	}

	err = f.Run(ac, Off, 0)
	if err != nil {
		t.Error(err)
	}
}

func ErrorEquals(err1 error, err2 error) bool {
	if err1 == nil || err2 == nil {
		return err1 == err2
	}
	return err1.Error() == err2.Error()
}
