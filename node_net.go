// Copyright 2011 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !js

package uuid

import "net"

var interfaces []net.Interface // cached list of interfaces

// getHardwareInterface returns the name and hardware address of interface name.
// If name is "" then the name and hardware address of one of the system's
// interfaces is returned.  If no interfaces are found (name does not exist or
// there are no interfaces) then "", nil is returned.
//
// Only addresses of at least 6 bytes are returned.
func getHardwareInterface(name string) (string, []byte) {
	if interfaces == nil {
		var err error
		interfaces, err = net.Interfaces()
		if err != nil && name != "" {
			return "", nil
		}
	}
	for _, ifs := range interfaces {
		if len(ifs.HardwareAddr) >= 6 && (name == "" || name == ifs.Name) {
			if setNodeID(ifs.HardwareAddr) {
				ifname = ifs.Name
				return ifname, nodeID
			}
		}
	}
	return "", nil
}
