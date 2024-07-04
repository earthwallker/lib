package json

import (
	_"encoding/hex"
	"encoding/json"
	"io/ioutil"
	_"strings"
    _"strconv"

	"fmt"
	"os"
	_ "sort"
)

var(

   
)

func init(){


}


type Config struct {
	Cmds []string `json:"Cmds"`
	Ratio []float64 `json:"Ratio"`

}


// func Test() error{
//    var config Config
//     err := ParseJSON("config.json",&config)
//     if err != nil {
        
//         errMsg:=fmt.Errorf("解析 JSON 数据失败: %v", err)
//         fmt.Println(errMsg,config)
//     }
//     fmt.Println(config)
//     return nil
// }


func ParseJSON(jsonStrFile string, v interface{}) error {
	// 读取 JSON 文件内容
	jsonFile, err := os.Open(jsonStrFile) // 假设JSON文件名为file.json(jsonFile)
	if err != nil {
		return fmt.Errorf("加载 JSON 文件失败: %v", err)
	}
    defer jsonFile.Close()
  
 
    bytes, err := ioutil.ReadAll(jsonFile)
      if err != nil {
          fmt.Println("读取JSON文件时出错:", err)
          return err
      }
	// 解析 JSON 数据到指定类型的结构体数组中
	err = json.Unmarshal(bytes, &v)
	if err != nil {
		return fmt.Errorf("解析 JSON 数据失败: %v", err)
	}
	return nil
}





