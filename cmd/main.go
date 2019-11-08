package main

import (
	"fmt"
	"github.com/welcome112s/apollo-go-sdk"
	"github.com/welcome112s/apollo-go-sdk/core"
)

func main() {
	apollosdk.SetDebug(false)
	apollosdk.Start("digitalkey-service","default","http://10.160.2.153:8083","")
	//apollosdk.StartWithCusConfig("./util/config.json")

	configNew := apollosdk.GetConfig("application")

	//t.Log(configNew.GetStringProperty("mats", ""))

	//定义变量来监听
	var varFunc core.OnChangeFunc =  func (changeEvent core.ConfigChangeEvent)  {
		fmt.Println("variable onChange",changeEvent)
	}
	configNew.AddChangeListenerFunc(varFunc)

	//定义结构实现接口来监听
	var s SomeThing = "s"
	configNew.AddChangeListener(s)

	//定义普通函数来监听
	configNew.AddChangeListenerFunc(onTestFunc)


	//block Test
	chan1 :=make(chan int)
	<-chan1






}


type SomeThing string

func (s SomeThing)OnChange (changeEvent core.ConfigChangeEvent)  {
	fmt.Println("struct onChange",changeEvent)
}



//定义普通的函数来接收
func onTestFunc(changeEvent core.ConfigChangeEvent) {
	fmt.Println("func onChange",changeEvent)
}
