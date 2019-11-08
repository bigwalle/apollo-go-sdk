package apollosdk

import (
	"testing"
	"github.com/welcome112s/apollo-go-sdk/core"
	"fmt"
)


func TestGetConfig(t *testing.T) {

	fmt.Println("TestGetConfig")

	//Start("aaa","aaa","aaa","aa","aaa")

	//config := GetAppConfig()
	//t.Log(config.GetStringProperty("test",""))

	//chanEvent := (*config).GetChangeKeyNotify()

	configNew := GetConfig("app.tc.mat.disable")

	t.Log(configNew.GetStringProperty("mats", ""))

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
