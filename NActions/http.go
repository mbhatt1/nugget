package NActions

import (
	"../NTypes"
	"bufio"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket/tcpassembly/tcpreader"
	"io"
	"log"
	"net/http"
)

type HTMLAction struct {
	executed  bool
	dependsOn BaseAction
	filters   []NTypes.Filter

	results []NTypes.HTTP
}

func (na *HTMLAction) BeenExecuted() bool {
	return na.executed
}

func (na *HTMLAction) DependencySatisfied() bool {
	if na.dependsOn == nil {
		return true
	}
	return na.dependsOn.BeenExecuted()
}

func (na *HTMLAction) SetDependency(action BaseAction) {
	na.dependsOn = action
}

func (na *HTMLAction) Execute() {
	if na.dependsOn != nil {
		//fmt.Println("md5 has a dependency which hasn't been met..")
		if na.dependsOn.BeenExecuted() == false {

			na.dependsOn.Execute()
		}
	}

	// Set up assembly
	streamFactory := &httpStreamFactory{}
	streamPool := tcpassembly.NewStreamPool(streamFactory)
	assembler := tcpassembly.NewAssembler(streamPool)

	operateOn := na.dependsOn.GetResults()
	if _, ok := operateOn.([]NTypes.NPacket); ok {
		var packets []NTypes.NPacket
		packets = operateOn.([]NTypes.NPacket)

		for _, packet := range packets {
			pkt := packet.Pkt

			unusablePrinted := false
			if pkt.NetworkLayer() == nil || pkt.TransportLayer() == nil || pkt.TransportLayer().LayerType() != layers.LayerTypeTCP {
				//log.Println("Unusable packet")
				if unusablePrinted == false {
					fmt.Println("Warning: unusable packets given to HTTP extractor")
					unusablePrinted = true
				}
				continue
			}
			tcp := pkt.TransportLayer().(*layers.TCP)
			assembler.AssembleWithTimestamp(pkt.NetworkLayer().NetworkFlow(), tcp, pkt.Metadata().Timestamp)

		}

	}
	na.executed = true
}

func (na *HTMLAction) GetResults() interface{} {
	if !na.BeenExecuted() {
		na.Execute()
	}
	return na.results
}

func (na *HTMLAction) SetFilters(filters []NTypes.Filter) {
	//TODO: investigate if resetting executed status will be a problem:
	na.executed = false
	na.filters = filters
}

// Packet -> HTTP helpers, taken from https://github.com/google/gopacket/blob/master/examples/httpassembly/main.go
// httpStreamFactory implements tcpassembly.StreamFactory
type httpStreamFactory struct{}

// httpStream will handle the actual decoding of http requests.
type httpStream struct {
	net, transport gopacket.Flow
	r              tcpreader.ReaderStream
}

func (h *httpStreamFactory) New(net, transport gopacket.Flow) tcpassembly.Stream {
	hstream := &httpStream{
		net:       net,
		transport: transport,
		r:         tcpreader.NewReaderStream(),
	}
	go hstream.run() // Important... we must guarantee that data from the reader stream is read.

	// ReaderStream implements tcpassembly.Stream, so we can return a pointer to it.
	return &hstream.r
}

func (h *httpStream) run() {
	buf := bufio.NewReader(&h.r)
	for {
		req, err := http.ReadRequest(buf)
		if err == io.EOF {
			// We must read until we see an EOF... very important!
			return
		} else if err != nil {
			//log.Println("Error reading stream", h.net, h.transport, ":", err)
		} else {
			bodyBytes := tcpreader.DiscardBytesToEOF(req.Body)
			req.Body.Close()
			log.Println("Received request from stream", h.net, h.transport, ":", req, "with", bodyBytes, "bytes in request body")
		}
		fmt.Println("HTTP Request host: " + req.Host)
	}
}
