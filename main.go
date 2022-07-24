package main

import (
	"encoding/binary"
	"fmt"

	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/codec/format"
	"github.com/go-netty/go-netty/codec/frame"
)

func main() {
	var childInitializer = func(channel netty.Channel) {
		channel.Pipeline().
			//frame.LengthFieldCodec(binary.BigEndian, int(^uint(0)>>1), 0, 2, 0, 2)
			//frame.LengthFieldPrepender(binary.BigEndian, 2, 0, false)
			AddLast(frame.LengthFieldCodec(binary.BigEndian, int(^uint(0)>>1), 0, 2, 0, 2)).
			AddLast(format.TextCodec()).
			AddLast(EchoHandler{})
	}
	netty.NewBootstrap(netty.WithChildInitializer(childInitializer)).
		Listen(":29100").Sync()
}

type EchoHandler struct{}

func (EchoHandler) HandleActive(ctx netty.ActiveContext) {
	fmt.Println("go-netty:", "->", "active:", ctx.Channel().RemoteAddr())
}

func (EchoHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	fmt.Println("go-netty:", "->", "handle read:", []byte(message.(string)))
	ctx.HandleRead(message)
}

func (EchoHandler) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {
	fmt.Println("go-netty:", "->", "inactive:", ctx.Channel().RemoteAddr(), ex)
	ctx.HandleInactive(ex)
}
