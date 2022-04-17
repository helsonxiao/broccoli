package main

import (
	"fmt"
	"testing"
)

const msg = `{"type":"message","event":"new","room_id":262451765,"user":{"user_id":9804407,"username":"迎宾机器猫小小川ゞ","iconurl":"http://static.missevan.com/avatars/202112/08/e40e541df9ca26088e9ddf26f0a30595234301.png","titles":[{"type":"level","level":49},{"type":"username","color":"#FF8686","colors":"#FF8686"},{"type":"noble","name":"新秀","level":2},{"type":"medal","name":"彼岸","super_fan":{"expire_time":1648915200},"level":18}]},"msg_id":"34ea0c75-a3dd-4149-8bf7-4d4f0caec46c","message":"@孙夫人s 来啦~ 小心炸弹！(╯‵□′)╯•••*～●，哈 骗你的啦~　"}`

func TestBroccoli(t *testing.T) {
	encodedData, _ := encode([]byte(msg))
	// encodedData, _ := encode([]byte(msg), brotli.WriterOptions{Quality: 5})
	fmt.Println(encodedData)

	decodedData, _ := decode(encodedData)
	decodedMsg := string(decodedData)
	fmt.Println(decodedMsg)

	if decodedMsg != msg {
		t.Errorf("decodedMsg != msg")
	}
}
