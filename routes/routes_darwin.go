package routes

import (
	"fmt"
	"net"
	"os/exec"
)

func routeAdd(dst interface{}, gw net.IP, priority int, iface string) error {
	// an implementation of "replace"
	routeDel(dst, gw, priority, iface)
	args := []string{
		"-n",
		"add",
		"-net",
		getNet(dst).String(),
	}
	if gw == nil {
		args = append(args, "-interface", iface)
	} else {
		args = append(args, gw.String())
	}
	v, err := exec.Command("route", args...).Output()
	if err != nil {
		return fmt.Errorf("failed to add %s route to %s interface: %s: %s", dst, iface, v, err)
	}
	return nil
}

func routeDel(dst interface{}, gw net.IP, priority int, iface string) error {
	args := []string{
		"-n",
		"delete",
		"-net",
		getNet(dst).String(),
	}
	if gw == nil {
		args = append(args, "-interface", iface)
	} else {
		args = append(args, gw.String())
	}
	v, err := exec.Command("route", args...).Output()
	if err != nil {
		return fmt.Errorf("failed to delete %s route from %s interface: %s: %s", dst, iface, v, err)
	}
	return nil
}