package controllers

import (
	"bcsdemo/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// set -> key:value
func (this *MainController) SetValue() {
	key := this.GetString("key")
	value := this.GetString("value")

	if key == ""||value == ""{
		this.Ctx.ResponseWriter.WriteHeader(500)
		this.Ctx.ResponseWriter.Write([]byte("Parameters key(or value) is empty!"))
	}

	beego.Info(key+":"+value)
	//  此处需要将 key:value 上链,需要请求 models 层的 fabric-sdk-go,此处打击可以理解为访问数据库
	var (
		channelID = beego.AppConfig.String("channelID")
		ueseID = beego.AppConfig.String("userID")
		chaincodeID = beego.AppConfig.String("chaincodeID")
	)
	ccs,err :=models.Init(channelID,ueseID,chaincodeID)
	if err!=nil{
		this.Ctx.ResponseWriter.WriteHeader(500)
		this.Ctx.ResponseWriter.Write([]byte("models.Init err!"))

	}
	var args [][]byte
	args = append(args,[]byte(key))
	args = append(args,[]byte(value))

	response,err := ccs.ChainCodeInvoke("set",args)
	if err!=nil{
		this.Ctx.ResponseWriter.WriteHeader(500)
		this.Ctx.ResponseWriter.Write([]byte("ccs.ChainCodeInvoke err!"))
	}

	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Ctx.ResponseWriter.Write(response)
}

// get -> key
func (this *MainController) GetValue() {
	key := this.GetString("key")

	if key == ""{
		this.Ctx.ResponseWriter.WriteHeader(500)
		this.Ctx.ResponseWriter.Write([]byte("Parameters key is empty!"))
	}

	beego.Info("key = ",key)
	//  此处需要将 key:value 上链,需要请求 models 层的 fabric-sdk-go,此处打击可以理解为访问数据库
	var (
		channelID = beego.AppConfig.String("channelID")
		ueseID = beego.AppConfig.String("userID")
		chaincodeID = beego.AppConfig.String("chaincodeID")
	)

	ccs,err :=models.Init(channelID,ueseID,chaincodeID)
	if err!=nil{
		this.Ctx.ResponseWriter.WriteHeader(500)
		this.Ctx.ResponseWriter.Write([]byte("models.Init err!"))
	}
	var args [][]byte
	args = append(args,[]byte(key))
	
	response,err := ccs.ChainCodeQuery("get",args)
	if err!=nil{
		this.Ctx.ResponseWriter.WriteHeader(500)
		this.Ctx.ResponseWriter.Write([]byte("ccs.ChainCodeQuery err!"))
	}

	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Ctx.ResponseWriter.Write(response)
}


