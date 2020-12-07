package feature

import (
	"strings"
)

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