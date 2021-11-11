package network

import (
	"xlddz/core/log"
	"net"
	"strings"
	"sync"
	"time"
)

//TCPClient 客户端连接
type TCPClient struct {
	sync.Mutex
	Addr            string
	ConnectInterval time.Duration
	PendingWriteNum int
	AutoReconnect   bool
	NewAgent        func(*TCPConn) Agent
	conns           ConnSet
	wg              sync.WaitGroup
	closeFlag       bool

	// msg parser
	LenMsgLen    int
	MinMsgLen    uint32
	MaxMsgLen    uint32
	LittleEndian bool
	msgParser    *MsgParser
}

func (client *TCPClient) Start() {
	client.init()

	go func() {
	reconnect:
		conn := client.dial()
		if conn == nil {
			return
		}

		client.Lock()
		if client.closeFlag {
			client.Unlock()
			conn.Close()
			return
		}
		client.conns[conn] = struct{}{}
		client.Unlock()

		tcpConn := newTCPConn(conn, client.PendingWriteNum, client.msgParser)
		agent := client.NewAgent(tcpConn)
		agent.Run()

		// cleanup
		tcpConn.Close()
		client.Lock()
		delete(client.conns, conn)
		client.Unlock()
		agent.OnClose()

		if client.AutoReconnect {
			time.Sleep(client.ConnectInterval)
			goto reconnect
		}
	}()
}

func (client *TCPClient) init() {
	client.Lock()
	defer client.Unlock()

	if client.ConnectInterval <= 0 {
		client.ConnectInterval = 3 * time.Second
		// log.Release("invalid ConnectInterval, reset to %v", client.ConnectInterval)
	}
	if client.PendingWriteNum <= 0 {
		client.PendingWriteNum = 1000
		// log.Release("invalid PendingWriteNum, reset to %v", client.PendingWriteNum)
	}
	if client.conns != nil {
		log.Fatal("tcpclient", "client is running")
	}
	if client.NewAgent == nil {
		log.Fatal("tcpclient", "NewAgent is nil")
	}
	if client.Addr == "" {
		log.Fatal("tcpclient", "client.Addr为空")
	}

	client.AutoReconnect = true
	client.conns = make(ConnSet)
	client.closeFlag = false

	// msg parser
	msgParser := NewMsgParser()
	msgParser.SetMsgLen(client.LenMsgLen, client.MinMsgLen, client.MaxMsgLen)
	msgParser.SetByteOrder(client.LittleEndian)
	client.msgParser = msgParser
}

func (client *TCPClient) dial() net.Conn {
	addr := strings.Split(client.Addr, "|")
	index := 0
	for {
		curConnAddr := addr[index%len(addr)]
		conn, err := net.Dial("tcp", curConnAddr)
		if err == nil || client.closeFlag {
			return conn
		}

		log.Info("TCPClient", "connect to %v error: %v,index=%v", curConnAddr, err, index)
		index++
		if index >= len(addr) {
			index = 0
		}
		time.Sleep(client.ConnectInterval)
		continue
	}
}

func (client *TCPClient) Close() {
	client.Lock()
	client.closeFlag = true
	for conn := range client.conns {
		conn.Close()
	}
	client.conns = nil
	client.Unlock()

	client.wg.Wait()
}

func (client *TCPClient) IsRunning() bool {
	if client.conns != nil {
		return true
	}

	return false
}
