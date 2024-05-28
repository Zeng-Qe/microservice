package webs

import (
	"fmt"
	"github.com/pion/webrtc/v3"
	"log"
)

// 定义全局的客户端连接池
//var clients = make(map[chan string]struct{})

//// 广播消息给所有连接的客户端
//func broadcast(message string) {
//	for client := range clients {
//		client <- message
//	}
//}

// 处理客户端连接
//func handleClient(connChan chan string) {
//	// 将客户端连接加入连接池
//	clients[connChan] = struct{}{}
//
//	// 客户端连接处理循环
//	for {
//		// 从连接读取消息
//		msg, ok := <-connChan
//		if !ok {
//			break
//		}
//		// 广播消息到所有客户端
//		broadcast(msg)
//	}
//
//	// 移除客户端连接
//	delete(clients, connChan)
//}
//
//// WebSocketHandler 服务器
//func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
//	connChan := make(chan string)
//	go handleClient(connChan)
//
//	// 升级HTTP连接为WebSocket连接
//	upgrader := websocket.Upgrader{}
//	ws, err := upgrader.Upgrade(w, r, nil)
//	if err != nil {
//		log.Println("Upgrade error:", err)
//		return
//	}
//	defer ws.Close()
//
//	// 处理WebSocket消息
//	for {
//		_, msg, err := ws.ReadMessage()
//		if err != nil {
//			log.Println("Read error:", err)
//			break
//		}
//		connChan <- string(msg) // 将消息转发到客户端连接处理
//		ws.WriteMessage(websocket.TextMessage, msg)
//	}
//}

func AddRTC() {
	// 创建一个新的WebRTC对象
	api := webrtc.NewAPI(nil)

	// 设置WebRTC配置
	cfg := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{URLs: []string{"stun:stun.l.google.com:19302"}},
		},
	}
	//api.SetConfiguration(cfg)

	// 创建一个新的PeerConnection
	pc, err := api.NewPeerConnection(cfg)
	if err != nil {
		log.Fatalf("Error creating PeerConnection: %v", err)
	}
	defer pc.Close()

	//// 处理远程流
	//pc.OnTrack(func(track *webrtc.TrackRemote, stream *webrtc.MediaStream) {
	//	fmt.Println("Track added:", track)
	//})

	// 创建一个Offer
	offer, err := pc.CreateOffer(nil)
	if err != nil {
		log.Fatalf("Error creating offer: %v", err)
	}

	// 设置本地描述
	err = pc.SetLocalDescription(offer)
	if err != nil {
		log.Fatalf("Error setting local description: %v", err)
	}

	// 通过信令服务器交换信令，然后连接PeerConnection

	// ... 信令交换和连接处理 ...

	fmt.Println("WebRTC server started")
	// 在这里添加更多的逻辑来处理音视频流...
}
