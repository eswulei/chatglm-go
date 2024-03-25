package main

import (
	"fmt"
	"github.com/eswulei/chatglm-go"
)

func main() {
	//apiKey := os.Getenv("API_KEY")
	apiKey := "2d7ec71ead2fe3312f6609a948d70ec3.XMno5bZYT2AXYGOA"
	//apiKey := "命令区\n｛\n1.当提问XX公司是否可以申请省科技领军企业时，你需要做出以下判断并回答：\n1)如果你知道这家企业的营业收入，且该企业营业收入超过5亿元，那么该企业具有较好的发展水平，有申请省科技领军企业的潜力，需要企业法人至科技引擎平台完善企业相关信息，以便你能更有效的协助企业申请省科技领军企业；\n2)如果你知道这家企业的营业收入，但该企业营业收入不足5亿元，那么该企业暂时不具备申请省科技领军企业的资质；\n3)如果你不知道这家企业的营业收入，那么需要该企业法人至科技引擎平台完善企业相关信息。\n4)嘉兴金领电子有限公司去年营收为60000万元\n2.按照以下例子来回答：\n｛\n火星人厨具股份有限公司营业收入22.77亿元，超过5亿元，具有较好的发展水平，有申请省科技领军企业的潜力，请企业法人至科技引擎平台完善企业相关信息，以便我能更有效的协助企业申请省科技领军企业。\n｝\n｛\n浙江焜腾红外科技有限公司营业收入7037.94万元，没有达到5亿元，暂时不具备申请省科技领军企业的资质。\n｝\n｛\n浙江骏马智创科技股份有限公司相关数据缺失，请在企业补充信息后尝试提问。\n｝\n｝\n我的问题是：嘉兴金领电子有限公司是否具备申报省科技领军企业资格"
	m := chatglm.ModelAPI{
		//Model:       chatglm.ChatGLMLite,
		Model:       "glm-3-turbo",
		TopP:        0.7,
		Temperature: 0.9,
		Messages: []map[string]interface{}{
			{"role": "user", "content": "俄乌战争发生在什么时候"},
			//{"role": "assistant", "content": "我是人工智能助手"},
			//{"role": "user", "content": "你叫什么名字"},
			//{"role": "assistant", "content": "我叫chatGLM"},
			//{"role": "user", "content": "你都可以做些什么事"},
		},
	}
	resp, err := m.Invoke(apiKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)

}
