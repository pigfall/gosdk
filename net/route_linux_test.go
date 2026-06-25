package net

import (
	nl "github.com/vishvananda/netlink"
	"testing"
)

func TestRouteList(t *testing.T) {
	rules, err := nl.RouteList(nil, nl.FAMILY_ALL)
	if err != nil {
		t.Fatal(err)
	}
	for _, rule := range rules {
		t.Log(rule.String())
	}
}
