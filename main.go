package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/wangtuanjie/ip17mon"
	"net"
)

type SmartIP struct {
	IP string
	ip17mon.LocationInfo
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	addr := this.Ctx.Request.RemoteAddr
	ip, _, _ := net.SplitHostPort(addr)
	loc, _ := ip17mon.Find(ip)
	sip := SmartIP{
		IP:           ip,
		LocationInfo: *loc,
	}
	data, _ := json.Marshal(sip)
	this.Ctx.WriteString(string(data))

	hostName := this.Ctx.Request.Header.Get("HostName")
	ifName := this.Ctx.Request.Header.Get("ifName")
	ifAddr := this.Ctx.Request.Header.Get("ifAddr")

	fmt.Printf("[%s, %s, %s, %s]\n", hostName, ifName, ifAddr, string(data))
}

func main() {
	if err := ip17mon.Init("./smart.db"); err != nil {
		panic(err)
	}
	beego.SetStaticPath("/download", "download")
	beego.Router("/smartip", &MainController{})
	beego.Run(":6001")
}
