Go Fence
========

Fence is a golang fencing library pluggable with different fence providers.

Fencing is the ability to isolate a node.
In a cluster, when you detect that a node isn't responding, you cannot know its state. It can be shutted down, but it can also simply have lost its network connectivity but still working and writing on a shared storage. Before restarting the "bad" node services on another node you have to be sure that the "bad" node is really isolated to avoid data corruptions and other various problems.

You can fence it powering it down, isolating its access to the storage and network etc...


## Fence Providers

At the moment there's one fencing provider that uses the fence agents provided by the redhat fence agents project (used by redhat cluster/pacemaker, oVirt and other projects).

[RedHat Fence Agents Provider](https://github.com/sgotti/go-fence-rhprovider)

## Api

Tha API isn't stable but should cover the various possibilities.


## Example

This example will use ipmilan protocol to connect to the integrated management console of a typical x86 based Server and force a reboot.

```go
package main

import (
        "log"
        "time"

        "github.com/sgotti/go-fence"
        "github.com/sgotti/go-fence-rhprovider"
)

func main() {
        f := fence.New()
        provider := rhprovider.New(nil)
        f.RegisterProvider("redhat", provider)

        err := f.LoadAgents(10 * time.Second)
        if err != nil {
                log.Print("error loading agents:", err)
                return
        }

        ac, err := f.NewAgentConfig("redhat", "fence_ipmilan")
        if err != nil {
                log.Print("error:", err)
                return
        }

        ac.SetParameter("ipaddr", "badhost.mydomain.org")
        ac.SetParameter("login", "username")
        ac.SetParameter("passwd", "password")

        err = f.Run(ac, fence.Reboot, 10*time.Second)
        if err != nil {
                log.Print("error: ", err)
                return
        }
        log.Print("Bad node fenced!")
}
```
