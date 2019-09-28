package comm

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Gzip(data []byte) {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	w.Write(data)
	w.Flush()
	//fmt.Println("gzip size:", len(b.Bytes()))
}

func UnGzip(byte []byte) []byte {
	b := bytes.NewBuffer(byte)
	r, _ := gzip.NewReader(b)
	defer r.Close()
	undatas, _ := ioutil.ReadAll(r)
	//fmt.Println("ungzip size:", len(undatas))
	return undatas
}

func JsonDecodeString(String string) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	json.Unmarshal([]byte(String), &jsonMap)
	return jsonMap
}

func JsonDecodeByte(bytes []byte) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	json.Unmarshal(bytes, &jsonMap)
	return jsonMap
}
func JsonEncodeMapToByte(stringMap map[string]interface{}) []byte {
	jsonBytes, err := json.Marshal(stringMap)
	if err != nil {
		return nil
	}
	return jsonBytes
}

/**
 * @note
 * 获取md5 hash串
 * @param string text 源串
 *
 * @return string
 */
func MD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func JsonEncode(data interface{}) string {
	s, e := json.Marshal(data)
	if e != nil {
		return ""
	}
	return string(s)
}

func JsonDecode(data string, inter interface{}) error {
	return json.Unmarshal([]byte(data), inter)
}

func StructToMap(data interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	j, _ := json.Marshal(data)
	_ = json.Unmarshal(j, &m)
	return m
}
func SendDDRobot(robotStruct DDRobotStruct, robotWebHook string) {
	jsonParams, err := json.Marshal(robotStruct)
	if err != nil {
		return
	}
	jsonString := string(jsonParams)
	payload := strings.NewReader(jsonString)
	req, _ := http.NewRequest("POST", robotWebHook, payload)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	res, _ := http.DefaultClient.Do(req)
	if res == nil {
		return
	}
	defer func() {
		if res != nil && res.Body != nil {
			defer res.Body.Close()
		}
	}()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))
}
