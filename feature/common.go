//package feature
package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type S3 struct {
	S3_a [7]uint8
	S3_b [7]uint8
}

type S2 struct {
	S2_a [4]uint8
	S2_b S3
	S2_c S3
	S2_d [4]uint8
	S2_e [32]uint8
}

type S1 struct {
	S1_a [8]uint8
	S1_b S2
	S1_c [16]uint8
}

/******************** PrintJson ********************/
func appendNewLine(strBuilder *strings.Builder, count int, padding string) {
	(*strBuilder).WriteByte('\n')
	for i := 0; i < count; i++ {
		(*strBuilder).WriteString(padding)
	}
}

/**
 * PrintJson：美化json.Marshal的输出
 * 输入：encoding/json.Marshal输出的结果
 * 输出：比较可观的json字符串
 */
func PrintJson(jsonByte []byte, padding string) string {
	depth := 0
	arrayComma := false
	var prettyStr strings.Builder
	for i := 0; i < len(jsonByte); i++ {
		switch jsonByte[i] {
		case '{':
			depth++
			prettyStr.WriteByte('{')
			appendNewLine(&prettyStr, depth, padding)
		case '}':
			depth--
			appendNewLine(&prettyStr, depth, padding)
			prettyStr.WriteByte('}')
		case '[':
			arrayComma = true
			prettyStr.WriteByte('[')
		case ']':
			arrayComma = false
			prettyStr.WriteByte(']')
		case ',':
			prettyStr.WriteByte(',')
			if !arrayComma {
				appendNewLine(&prettyStr, depth, padding)
			}
		default:
			prettyStr.WriteByte(jsonByte[i])
		}
	}

	return prettyStr.String()
}
/******************** PrintJson ********************/

func main() {
	var s1 S1
	json, _ := json.Marshal(s1)
	str:= PrintJson(json, "    ")
	fmt.Println(string(str))
}