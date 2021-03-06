// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package ecs

// The network is defined as the communication path over which a host or
// network event happens.
// The network.* fields should be populated with details about the network
// activity associated with an event.
type Network struct {
	// Name given by operators to sections of their network.
	Name string `ecs:"name"`

	// In the OSI Model this would be the Network Layer. ipv4, ipv6, ipsec,
	// pim, etc
	// The field value must be normalized to lowercase for querying. See the
	// documentation section "Implementing ECS".
	Type string `ecs:"type"`

	// IANA Protocol Number
	// (https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).
	// Standardized list of protocols. This aligns well with NetFlow and sFlow
	// related logs which use the IANA Protocol Number.
	IANANumber string `ecs:"iana_number"`

	// Same as network.iana_number, but instead using the Keyword name of the
	// transport layer (udp, tcp, ipv6-icmp, etc.)
	// The field value must be normalized to lowercase for querying. See the
	// documentation section "Implementing ECS".
	Transport string `ecs:"transport"`

	// A name given to an application level protocol. This can be arbitrarily
	// assigned for things like microservices, but also apply to things like
	// skype, icq, facebook, twitter. This would be used in situations where
	// the vendor or service can be decoded such as from the source/dest IP
	// owners, ports, or wire format.
	// The field value must be normalized to lowercase for querying. See the
	// documentation section "Implementing ECS".
	Application string `ecs:"application"`

	// L7 Network protocol name. ex. http, lumberjack, transport protocol.
	// The field value must be normalized to lowercase for querying. See the
	// documentation section "Implementing ECS".
	Protocol string `ecs:"protocol"`

	// Direction of the network traffic.
	// Recommended values are:
	//   * ingress
	//   * egress
	//   * inbound
	//   * outbound
	//   * internal
	//   * external
	//   * unknown
	//
	// When mapping events from a host-based monitoring context, populate this
	// field from the host's point of view, using the values "ingress" or
	// "egress".
	// When mapping events from a network or perimeter-based monitoring
	// context, populate this field from the point of view of the network
	// perimeter, using the values "inbound", "outbound", "internal" or
	// "external".
	// Note that "internal" is not crossing perimeter boundaries, and is meant
	// to describe communication between two hosts within the perimeter. Note
	// also that "external" is meant to describe traffic between two hosts that
	// are external to the perimeter. This could for example be useful for ISPs
	// or VPN service providers.
	Direction string `ecs:"direction"`

	// Host IP address when the source IP address is the proxy.
	ForwardedIP string `ecs:"forwarded_ip"`

	// A hash of source and destination IPs and ports, as well as the protocol
	// used in a communication. This is a tool-agnostic standard to identify
	// flows.
	// Learn more at https://github.com/corelight/community-id-spec.
	CommunityID string `ecs:"community_id"`

	// Total bytes transferred in both directions.
	// If `source.bytes` and `destination.bytes` are known, `network.bytes` is
	// their sum.
	Bytes int64 `ecs:"bytes"`

	// Total packets transferred in both directions.
	// If `source.packets` and `destination.packets` are known,
	// `network.packets` is their sum.
	Packets int64 `ecs:"packets"`

	// Network.inner fields are added in addition to network.vlan fields to
	// describe the innermost VLAN when q-in-q VLAN tagging is present. Allowed
	// fields include vlan.id and vlan.name. Inner vlan fields are typically
	// used when sending traffic with multiple 802.1q encapsulations to a
	// network sensor (e.g. Zeek, Wireshark.)
	Inner map[string]interface{} `ecs:"inner"`
}
