package main

import (
	"bytes"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	greq "github.com/syhlion/greq"
	requestwork "github.com/syhlion/requestwork.v2"
)

var wg sync.WaitGroup

func main() {

	const workNum = 100    // worker 數量: 同時呼叫API的數量
	const caseNum = 10000 // request 總量

	tStart := time.Now()
	testAPI(workNum, caseNum)
	tEnd := time.Now().Sub(tStart)
	log.Println("test done ....cost time:", tEnd)

}

// testAPI 測試 API
func testAPI(workNum, caseNum int) {

	// 建立workNum個worker
	worker := requestwork.New(workNum)

	// client 設定timeout 時間
	OnionConn := greq.New(worker, 3*time.Second, false)
	for i := 0; i < caseNum; i++ {
		wg.Add(1)
		go func() {

			// postRaw := map[string]interface{}{
			// 	"game_id": 105,
			// 	"ip":      "127.0.0.1",
			// 	"session": "7923ab5adfcf0c03fa4f6599af1c99815d946da81eee7c2f802fee8a241732df",
			// }

			postRaw := map[string]interface{}{
				"amount":     1,
				"opcode":     10003,
				"partner_id": 9,
			}

			// io buffer
			para := bytes.NewBuffer(httpJSONRawBuild(postRaw))
			tS := time.Now()
			_, _, err := OnionConn.PostRaw("http://127.0.0.1/api/play/cash/increase", para)
			// _, _, err := OnionConn.PostRaw("http://127.0.0.1/api/play/user/session", para)
			tE := time.Now().Sub(tS)

			log.Println("花費時間：", tE)
			if tE > 2*time.Second {
				log.Println("大於2秒")
			}

			if err != nil {
				log.Println(err)
			}
			// if s != http.StatusOK {
			// 	log.Println("http status", s)
			// }

			// log.Println(string(data))
			wg.Done()
		}()

		// 想逼緊一點，使用外面的wg.Wait()
		// 想鬆一點，讓每輪的request做完，才做下一輪
		if i%workNum == 0 {
			log.Println("wait!!!")
			wg.Wait()
		}
	}
	// wg.Wait()
}

// httpJSONRawBuild 將map組成json格式
func httpJSONRawBuild(a map[string]interface{}) []byte {
	var retString = "{"
	for k, v := range a {

		switch t := v.(type) {
		case string:
			sub := (`"` + k + `"` + `:"` + v.(string) + `"`)
			retString += (sub + ",")
		case int:
			v = strconv.Itoa(t)
			sub := (`"` + k + `"` + `:` + v.(string))

			retString += (sub + ",")
		case []string:
			// string array裡面的東西有沒有雙引號會不會有問題，待驗證
			sub := strings.Join(v.([]string), ",")
			sub = strings.TrimSuffix(sub, ",")
			retString += "[" + sub + "]"
		}
	}
	// 最後的,移除掉
	retString = strings.TrimSuffix(retString, ",")
	retString += "}"

	// log.Println(retString)

	return []byte(retString)
}
