package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"net"
)

// 该程序分析数据包中http的请求响应信息，方便做测试
// 注意：
// 1. 包要去重
// 2. 只能处理http 1.1
// todo:
// 1. 最好写成一个可以复用的小程序

type endpoint struct {
	ip   net.IP
	port layers.TCPPort
}

func (e *endpoint) String() string {
	return e.ip.String() + ":" + e.port.String()
}

type httpMsg struct {
	request     []byte
	response    []byte
	src         endpoint
	dst         endpoint
	finished    bool
	seqRequest  uint32
	seqResponse uint32
}

func (h *httpMsg) Print() {
	fmt.Printf("%s:%s => %s:%s\n", h.src.ip.String(), h.src.port.String(), h.dst.ip.String(), h.dst.port.String())
	fmt.Printf("%s\n", string(h.request))
	fmt.Printf("%s:%s => %s:%s\n", h.dst.ip.String(), h.dst.port.String(), h.src.ip.String(), h.src.port.String())
	fmt.Printf("%s\n", string(h.response))
}
func (h *httpMsg) IsFinished() bool {
	return h.finished
}
func main() {

	//handle, err := pcap.OpenLive(
	//	"en0", // device
	//	int32(65535), // snapshot length
	//	false, // promiscuous mode?
	//	-1 * time.Second, // timeout 负数表示不缓存，直接输出
	//)
	// todo: 这个包里面有重发的包，会导致不正确的结果； 其中一种做法是，专门写一个包处理工具，负责去除重复的包
	handle, err := pcap.OpenOffline("/tmp/a.tcpdump")
	if err != nil {
		panic(err)
	}
	defer handle.Close()
	var filter string = "tcp and port 80"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	flow := map[string]*httpMsg{}
	for packet := range packetSource.Packets() {
		iplayer, ok := packet.NetworkLayer().(*layers.IPv4)
		if !ok {
			continue
		}
		tcpLayer, ok := packet.TransportLayer().(*layers.TCP)
		if !ok {
			continue
		}
		streamId := ""

		src := endpoint{ip: iplayer.SrcIP, port: tcpLayer.SrcPort}
		dst := endpoint{iplayer.DstIP, tcpLayer.DstPort}
		streamId = src.String() + "=>" + dst.String()
		if tcpLayer.SYN && !tcpLayer.ACK {
			httpmsg := &httpMsg{}
			httpmsg.src = src
			httpmsg.dst = dst
			flow[streamId] = httpmsg
		}
		seq := tcpLayer.Seq
		fmt.Printf("streamId: %s seq: %d, len: %d\n", streamId, seq, len(tcpLayer.LayerPayload()))

		if _, ok := flow[streamId]; !ok {
			streamId = dst.String() + "=>" + src.String()
			if _, ok := flow[streamId]; !ok {
				continue
			}
		}
		httpmsg := flow[streamId]
		isRequest := iplayer.SrcIP.Equal(httpmsg.src.ip)
		isResponse := iplayer.SrcIP.Equal(httpmsg.dst.ip)

		content := tcpLayer.LayerPayload()

		// 扔掉重传的包
		if isRequest {
			if seq <= httpmsg.seqRequest {
				continue
			}
			httpmsg.seqRequest = seq
		}
		if isResponse {
			if seq <= httpmsg.seqResponse {
				continue
			}
			httpmsg.seqResponse = seq
		}

		if len(content) > 0 {
			if isRequest {
				httpmsg.request = append(httpmsg.request, content...)
			}
			if isResponse {
				httpmsg.response = append(httpmsg.response, content...)
				httpmsg.finished = true
			}
			//fmt.Printf("%s => %s: \n%s", src, dst, string(content))
		}
		if tcpLayer.FIN && isResponse {
			//if httpmsg.IsFinished() {
			httpmsg.Print()
			delete(flow, streamId)
		}

	}
}

func findAllDev() {
	// Find all devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	// Print device information
	fmt.Println("Devices found:")
	for _, device := range devices {
		fmt.Println("\nName: ", device.Name)
		fmt.Println("Description: ", device.Description)
		fmt.Println("Devices addresses: ", device.Description)
		for _, address := range device.Addresses {
			fmt.Println("- IP address: ", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)
		}
	}
}
