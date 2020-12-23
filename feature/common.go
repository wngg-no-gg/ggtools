//package feature
package main

import (
    "fmt"
    "log"
    "os"
    "os/exec"
    "strings"
)

/******************** PrettyJson ********************
 * import:
 *     "strings"
 */
func appendNewLine(strBuilder *strings.Builder, count int, padding string) {
    (*strBuilder).WriteByte('\n')
    for i := 0; i < count; i++ {
        (*strBuilder).WriteString(padding)
    }
}

/**
 * PrettyJson：美化json.Marshal的输出
 * 输入：encoding/json.Marshal输出的结果
 * 输出：比较可观的json字符串
 */
func PrettyJson(jsonByte []byte, padding string) string {
    depth := 0
    arrayComma := false /* Determine if it is a comma in an array */
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
/******************** PrettyJson ********************/

/******************** CodeRunC ********************
 * import:
 *     "fmt"
 *     "os"
 *     "os/exec"
 */

const (
    CODE_TYPE_C        int = 0 * 0x00001000
    CODE_TYPE_CPP      int = 1 * 0x00001000
)

const (
    CODE_MAKE_TYPE_C_GCC    int = 0x1 + CODE_TYPE_C
)

func writeCodeFile(code string) {}

func CodeRunC(code string) {
    file, err := os.OpenFile("../temp/code-run/code-run-c.c", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
    if err != nil {
        log.Fatal(err)
    }

    _, err = file.WriteString(code)
    if err != nil {
        log.Fatal(err)
    }
    file.Close()


    cmd := exec.Command("gcc", "../temp/code-run/code-run-c.c", "-o", "../temp/code-run/code-run-c.out")
    out, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("error out:\n%s\n", string(out))
        log.Fatalf("cmd.Run() failed with %s\n", err)
    }

    cmd = exec.Command("../temp/code-run/code-run-c.out")
    out, err = cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("error out:\n%s\n", string(out))
        log.Fatalf("cmd.Run() failed with %s\n", err)
    }
    fmt.Printf("%s", string(out))
}
/******************** CodeRunC ********************/


/******************** TestMain ********************/
func main() {
    CodeRunC("#include <stdio.h>\n\nint main()\n{\n    printf(\"test\\n1\\n2\\n\");\n    return 0;\n}")
}
/******************** TestMain ********************/