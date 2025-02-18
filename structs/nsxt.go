package structs

import (
	"fmt"
	"net/netip"
	"strconv"
	"strings"
)

type EnforcementPoint struct {
	Id   string
	Path string
}

type TransportZone struct {
	Id   string
	Name string
	Type string
}

type EdgeClusterMember struct {
	Index int    `json:"member_index"`
	Id    string `json:"transport_node_id"`
}

type EdgeCluster struct {
	Id             string              `json:"id"`
	Name           string              `json:"display_name"`
	Type           string              `json:"deployment_type"`
	Members        []EdgeClusterMember `json:"members"`
	MemberNodeType string              `json:"member_node_type"`
}

type EdgeClusters []EdgeCluster

func (ecs *EdgeClusters) GetClusterById(Id string) *EdgeCluster {
	for _, ec := range *ecs {
		if ec.Id == Id {
			return &ec
		}
	}
	return &EdgeCluster{}
}

type RouteEntry struct {
	Type       string `json:"route_type"`
	Network    string `json:"network"`
	NextHop    string `json:"next_hop"`
	Ad         uint8  `json:"admin_distance"`
	RouterId   string `json:"lr_component_id"`
	RouterType string `json:"lr_component_type"`
}

type EdgeRoute struct {
	NodePath string       `json:"edge_node"`
	Entries  []RouteEntry `json:"route_entries"`
}

func (er *EdgeRoute) Print() {
	for _, e := range er.Entries {
		var routeType string
		switch e.Type {
		case "t0c":
			routeType = "C"
		case "t0s":
			routeType = "S"
		case "b":
			routeType = "B"
		case "t0n":
			routeType = "N"
		case "t1c":
			routeType = "c"
		case "t1s":
			routeType = "s"
		case "t1n":
			routeType = "n"
		case "t1l":
			routeType = "l"
		case "t1ls":
			routeType = "ln"
		case "t1d":
			routeType = "d"
		case "t1ipsec":
			routeType = "p"
		case "isr":
			routeType = "i"
		}
		if routeType == "C" {
			fmt.Printf("%v> %v is directly connected\n", routeType, e.Network)
		} else if routeType == "i" && e.NextHop == "" {
			fmt.Printf("%v> %v [%v] blackhole\n", routeType, e.Network, e.Ad)
		} else {
			fmt.Printf("%v> %v [%v] via %v\n", routeType, e.Network, e.Ad, e.NextHop)
		}
	}
	fmt.Println()
}

func (er *EdgeRoute) GetEdgeClusterId() string {
	path := strings.Split(er.NodePath, "/")
	return path[7]
}

func (er *EdgeRoute) GetEdgeClusterNodeIdx() int {
	path := strings.Split(er.NodePath, "/")
	idx, _ := strconv.Atoi(path[9])
	return idx
}

func (er *EdgeRoute) GetEntries(version int) RouteEntries {
	var entries []RouteEntry
	var bitLen int
	if version == 6 {
		bitLen = 128
	} else {
		bitLen = 32
	}
	// filter whether IPv4 or IPv6
	for _, e := range er.Entries {
		eip, _ := netip.ParsePrefix(e.Network)
		if eip.Addr().BitLen() != bitLen {
			continue
		}
		entries = append(entries, e)
	}
	// check addresssing order
	nthSmall := make([]int, len(entries))
	for idx, en := range entries {
		small := 0
		for _, e := range entries {
			en_prefix, _ := netip.ParsePrefix(en.Network)
			e_prefix, _ := netip.ParsePrefix(e.Network)
			if e_prefix.Addr().Compare(en_prefix.Addr()) < 0 {
				small += 1
			}
		}
		nthSmall[idx] = small
	}
	sorted_entries := make([]RouteEntry, len(entries))
	for idx, se := range nthSmall {
		sorted_entries[se] = entries[idx]
	}
	return RouteEntries(sorted_entries)
}

type RouteEntries []RouteEntry

func (res *RouteEntries) Print() {
	for _, e := range *res {
		var routeType string
		switch e.Type {
		case "t0c":
			routeType = "C"
		case "t0s":
			routeType = "S"
		case "b":
			routeType = "B"
		case "t0n":
			routeType = "N"
		case "t1c":
			routeType = "c"
		case "t1s":
			routeType = "s"
		case "t1n":
			routeType = "n"
		case "t1l":
			routeType = "l"
		case "t1ls":
			routeType = "ln"
		case "t1d":
			routeType = "d"
		case "t1ipsec":
			routeType = "p"
		case "isr":
			routeType = "i"
		}
		if routeType == "C" {
			fmt.Printf("%v> %v is directly connected\n", routeType, e.Network)
		} else if routeType == "i" && e.NextHop == "" {
			fmt.Printf("%v> %v [%v] blackhole\n", routeType, e.Network, e.Ad)
		} else {
			fmt.Printf("%v> %v [%v] via %v\n", routeType, e.Network, e.Ad, e.NextHop)
		}
	}
	fmt.Println()
}

type ComputeManager struct {
	Id     string
	Name   string
	Type   string
	Server string
	Detail string
	Status *ComputeManagerStatus
}

func (cm *ComputeManager) Print() {
	fmt.Println("Name:    ", cm.Name)
	fmt.Println("ID:      ", cm.Id)
	fmt.Println("Type:    ", cm.Type)
	fmt.Println("FQDN/IP: ", cm.Server)
	fmt.Println("Version: ", cm.Detail)
	if cm.Status != nil {
		s := fmt.Sprintf("Status:   %s, %s", cm.Status.Connection, cm.Status.Registration)
		fmt.Println(s)
	}
}

type ComputeManagerStatus struct {
	Connection   string
	Registration string
}

type TransportNode struct {
	Id   string `json:"id"`
	Name string `json:"display_name"`
}

type PerNodeStatisticsRx struct {
	TotalBytes                           uint64 `json:"total_bytes"`
	TotalPackets                         uint64 `json:"total_packets"`
	DroppedPackets                       uint64 `json:"dropped_packets"`
	BlockedPackets                       uint64 `json:"blocked_packets"`
	DestinationUnsupportedDroppedPackets uint64 `json:"destination_unsupported_dropped_packets"`
	FirewallDroppedPackets               uint64 `json:"firewall_dropped_packets"`
	IpsecDroppedPackets                  uint64 `json:"ipsec_dropped_packets"`
	IpsecNoSaDroppedPackets              uint64 `json:"ipsec_no_sa_dropped_packets"`
	IpsecNoVtiDroppedPackets             uint64 `json:"ipsec_no_vti_dropped_packets"`
	Ipv6DroppedPackets                   uint64 `json:"ipv6_dropped_packets"`
	KniDroppedPackets                    uint64 `json:"kni_dropped_packets"`
	L4portUnsupportedDroppedPackets      uint64 `json:"l4port_unsupported_dropped_packets"`
	MalformedDroppedPackets              uint64 `json:"malformed_dropped_packets"`
	NoReceiverDroppedPackets             uint64 `json:"no_receiver_dropped_packets"`
	NoRouteDroppedPackets                uint64 `json:"no_route_dropped_packets"`
	ProtoUnsupportedDroppedPackets       uint64 `json:"proto_unsupported_dropped_packets"`
	RedirectDroppedPackets               uint64 `json:"redirect_dropped_packets"`
	RpfCheckDroppedPackets               uint64 `json:"rpf_check_dropped_packets"`
	TtlExceededDroppedPackets            uint64 `json:"ttl_exceeded_dropped_packets"`
}

type PerNodeStatisticsTx struct {
	TotalBytes                  uint64 `json:"total_bytes"`
	TotalPackets                uint64 `json:"total_packets"`
	DroppedPackets              uint64 `json:"dropped_packets"`
	BlockedPackets              uint64 `json:"blocked_packets"`
	FirewallDroppedPackets      uint64 `json:"firewall_dropped_packets"`
	IpsecDroppedPackets         uint64 `json:"ipsec_dropped_packets"`
	IpsecNoSaDroppedPackets     uint64 `json:"ipsec_no_sa_dropped_packets"`
	IpsecNoVtiDroppedPackets    uint64 `json:"ipsec_no_vti_dropped_packets"`
	DadDroppedPackets           uint64 `json:"dad_dropped_packets"`
	FragNeededDroppedPackets    uint64 `json:"frag_needed_dropped_packets"`
	IpSecPolBlockDroppedPackets uint64 `json:"ipsec_pol_block_dropped_packets"`
	IpSecPolErrDroppedPackets   uint64 `json:"ipsec_pol_err_dropped_packets"`
	NoArpDroppedPackets         uint64 `json:"no_arp_dropped_packets"`
	NoLinkedDroppedPackets      uint64 `json:"no_linked_dropped_packets"`
	NoMemDroppedPackets         uint64 `json:"no_mem_dropped_packets"`
	NonIpDroppedPackets         uint64 `json:"non_ip_dropped_packets"`
	ServiceInsertDroppedPackets uint64 `json:"service_insert_dropped_packets"`
}

type PerNodeStatistics struct {
	LastUpdate uint64              `json:"last_update_timestamp"`
	Rx         PerNodeStatisticsRx `json:"rx"`
	Tx         PerNodeStatisticsTx `json:"tx"`
}

type RouterStats struct {
	PortId            string              `json:"logical_router_port_id"`
	PerNodeStatistics []PerNodeStatistics `json:"per_node_statistics"`
}

type Tier0Gateway struct {
	Id           string `json:"id"`
	HaMode       string `json:"ha_mode"`
	Name         string `json:"display_name"`
	FailoverMode string `json:"failover_mode"`
}

type Tier0Gateways []Tier0Gateway

func (gws Tier0Gateways) Print(output string) {
	if output == "json" {
	} else {
		fmt.Printf("%-8s	%-8s	%-8s	%-8s\n", "ID", "Name", "HA Mode", "Failover Mode")
		for _, gw := range gws {
			fmt.Printf("%-8s	%8s	%8s	%8s\n", gw.Id, gw.Name, gw.HaMode, gw.FailoverMode)
		}
	}

}
