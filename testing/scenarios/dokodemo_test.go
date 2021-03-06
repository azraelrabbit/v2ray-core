package scenarios

import (
	"net"
	"testing"

	v2net "v2ray.com/core/common/net"
	"v2ray.com/core/testing/assert"
	"v2ray.com/core/testing/servers/tcp"
)

func TestDokodemoTCP(t *testing.T) {
	assert := assert.On(t)

	tcpServer := &tcp.Server{
		Port: v2net.Port(50016),
		MsgProcessor: func(data []byte) []byte {
			buffer := make([]byte, 0, 2048)
			buffer = append(buffer, []byte("Processed: ")...)
			buffer = append(buffer, data...)
			return buffer
		},
	}
	_, err := tcpServer.Start()
	assert.Error(err).IsNil()
	defer tcpServer.Close()

	assert.Error(InitializeServerSetOnce("test_2")).IsNil()

	dokodemoPortStart := v2net.Port(50011)
	dokodemoPortEnd := v2net.Port(50015)

	for port := dokodemoPortStart; port <= dokodemoPortEnd; port++ {
		conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
			IP:   []byte{127, 0, 0, 1},
			Port: int(port),
		})

		payload := "dokodemo request."
		nBytes, err := conn.Write([]byte(payload))
		assert.Error(err).IsNil()
		assert.Int(nBytes).Equals(len(payload))

		conn.CloseWrite()

		response := make([]byte, 1024)
		nBytes, err = conn.Read(response)
		assert.Error(err).IsNil()
		assert.String("Processed: " + payload).Equals(string(response[:nBytes]))
		conn.Close()
	}

	CloseAllServers()
}
