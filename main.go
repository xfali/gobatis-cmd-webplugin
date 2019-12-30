// Copyright (C) 2019, Xiongfa Li.
// All right reserved.
// @author xiongfa.li
// @version V1.0
// Description: 

package main

import (
    "bufio"
    "encoding/json"
    "flag"
    "fmt"
    "github.com/xfali/gobatis-cmd/common"
    "os"
    "strings"
)

func main() {
    method := flag.String(common.MethodFlag, "", "method name")
    flag.Parse()

    if *method == common.OutPutSuffixMethod {
        outputSuffix()
    } else if *method == common.GenerateMethod {
        generate()
    }
}

func outputSuffix() {
    fmt.Println("_handler.go")
    os.Exit(0)
}

func generate() {
    reader := bufio.NewReader(os.Stdin)
    data, _, err := reader.ReadLine()
    if err != nil {
        os.Exit(1)
    }

    ret := common.GenerateInfo{}
    err = json.Unmarshal([]byte(data), &ret)
    if err != nil {
        fmt.Errorf("data format error %s\n", string(data))
        os.Exit(2)
    }

    fmt.Println(writeHandler(ret))
    os.Exit(0)
}

func writeHandler(info common.GenerateInfo) string {
    builder := strings.Builder{}
    ModelName := common.TableName2ModelName(info.Table)

    sessionFunc := "newSession"

    builder.WriteString("package handler")
    builder.WriteString(common.Newline())
    builder.WriteString(common.Newline())

    builder.WriteString("import(")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(`"github.com/gin-gonic/gin"`)
    builder.WriteString(common.Newline())
    builder.WriteString(")")

    builder.WriteString(common.Newline())
    builder.WriteString(common.Newline())
    builder.WriteString(fmt.Sprintf("func Register%sHandler(engine *gin.Engine) {", ModelName))

    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(fmt.Sprintf(`engine.GET("%s", Get%s)`, ModelName, ModelName))
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(fmt.Sprintf(`engine.POST("%s", Create%s)`, ModelName, ModelName))
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(fmt.Sprintf(`engine.PUT("%s", Update%s)`, ModelName, ModelName))
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(fmt.Sprintf(`engine.DELETE("%s", Delete%s)`, ModelName, ModelName))

    builder.WriteString(common.Newline())
    builder.WriteString("}")

    builder.WriteString(common.Newline())
    builder.WriteString(common.Newline())
    builder.WriteString(fmt.Sprintf("func Get%s(ctx *gin.Context) {", ModelName))

    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("var req " + ModelName)

    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(fmt.Sprintf("list, errSelect := %s.Select%s(%s(), req)", info.Package, ModelName, sessionFunc))

    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("if errSelect != nil {")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("ctx.AbortWithStatus(400)")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("return")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("}")

    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("ctx.JSON(200, list)")

    builder.WriteString(common.Newline())
    builder.WriteString("}")

    builder.WriteString(common.Newline())
    builder.WriteString(common.Newline())
    builder.WriteString(fmt.Sprintf("func Create%s(ctx *gin.Context) {", ModelName))
    builder.WriteString(common.Newline())

    builder.WriteString(common.ColumnSpace())
    builder.WriteString("var req " + ModelName)
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("err := ctx.Bind(&req)")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("if err != nil {")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("ctx.AbortWithStatus(400)")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("return")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("}")

    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(fmt.Sprintf("_, _, errInsert := %s.Insert%s(%s(), req)", info.Package, ModelName, sessionFunc))

    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("if errInsert != nil {")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("ctx.AbortWithStatus(400)")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("return")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("}")

    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("ctx.Writer.WriteHeader(201)")

    builder.WriteString(common.Newline())
    builder.WriteString("}")

    builder.WriteString(common.Newline())
    builder.WriteString(common.Newline())
    builder.WriteString(fmt.Sprintf("func Update%s(ctx *gin.Context) {", ModelName))
    builder.WriteString(common.Newline())

    builder.WriteString(common.ColumnSpace())
    builder.WriteString("var req " + ModelName)
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("err := ctx.Bind(&req)")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("if err != nil {")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("ctx.AbortWithStatus(400)")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("return")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("}")

    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(fmt.Sprintf("i, errUpdate := %s.Update%s(%s(), req)", info.Package, ModelName, sessionFunc))

    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("if errUpdate != nil {")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("ctx.AbortWithStatus(400)")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("return")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("}")

    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("if i == 0 {")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("ctx.AbortWithStatus(404)")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("return")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("}")

    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("ctx.Writer.WriteHeader(200)")

    builder.WriteString(common.Newline())
    builder.WriteString("}")

    builder.WriteString(common.Newline())
    builder.WriteString(common.Newline())
    builder.WriteString(fmt.Sprintf("func Delete%s(ctx *gin.Context) {", ModelName))

    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("var req " + ModelName)

    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(fmt.Sprintf("i, errDelete := %s.Delete%s(%s(), req)", info.Package, ModelName, sessionFunc))

    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("if errDelete != nil {")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("ctx.AbortWithStatus(400)")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("return")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("}")

    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("if i == 0 {")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("ctx.AbortWithStatus(404)")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("return")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("}")

    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("ctx.Writer.WriteHeader(200)")

    builder.WriteString(common.Newline())
    builder.WriteString("}")

    return builder.String()
}
