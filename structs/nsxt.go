package structs

import "fmt"

type EnforcementPoint struct {
	Id   string
	Path string
}

type TransportZone struct {
	Id   string
	Name string
	Type string
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
