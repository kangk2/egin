package lib

import (
    "fmt"
    "log"
    "reflect"
    "strconv"
    "strings"
    "sync"
)

// 反射调用结构体方法
func Invoke(any interface{}, name string, args ...interface{}) ([]reflect.Value, error) {
    method := reflect.ValueOf(any).MethodByName(name)
    var _result []reflect.Value
    notExist := method == reflect.Value{}
    if notExist {
        return _result, fmt.Errorf("Method %s not found ", name)
    }
    methodType := method.Type()
    numIn := methodType.NumIn()
    if numIn > len(args) {
        return _result, fmt.Errorf("Method %s must have minimum %d params. Have %d ", name, numIn, len(args))
    }
    if numIn != len(args) && !methodType.IsVariadic() {
        return _result, fmt.Errorf("Method %s must have %d params. Have %d ", name, numIn, len(args))
    }
    in := make([]reflect.Value, len(args))
    for i := 0; i < len(args); i++ {
        var inType reflect.Type
        if methodType.IsVariadic() && i >= numIn-1 {
            inType = methodType.In(numIn - 1).Elem()
        } else {
            inType = methodType.In(i)
        }
        argValue := reflect.ValueOf(args[i])
        if !argValue.IsValid() {
            return _result, fmt.Errorf("Method %s. Param[%d] must be %s. Have %s ", name, i, inType, argValue.String())
        }
        argType := argValue.Type()
        if argType.ConvertibleTo(inType) {
            in[i] = argValue.Convert(inType)
        } else {
            return _result, fmt.Errorf("Method %s. Param[%d] must be %s. Have %s ", name, i, inType, argType)
        }
    }
    return method.Call(in), nil
}

// 获取结构体中字段的名称
func GetStructFieldsName(structName interface{}) []string {
    t := reflect.TypeOf(structName)
    if t.Kind() == reflect.Ptr {
        t = t.Elem()
    }
    if t.Kind() != reflect.Struct {
        log.Println("Check type error not Struct")
        return nil
    }
    fieldNum := t.NumField()
    result := make([]string, 0, fieldNum)
    for i := 0; i < fieldNum; i++ {
        result = append(result, t.Field(i).Name)
    }
    return result
}

// 获取结构体中Tag的值，如果没有tag则返回字段值
func GetStructTags(structName interface{}) map[string]map[string]string {
    t := reflect.TypeOf(structName)
    if t.Kind() == reflect.Ptr {
        t = t.Elem()
    }
    if t.Kind() != reflect.Struct {
        log.Println("Check type error not Struct")
        return nil
    }
    fieldNum := t.NumField()
    result := make(map[string]map[string]string)
    for i := 0; i < fieldNum; i++ {
        fieldName := t.Field(i).Name
        tagStr := string(t.Field(i).Tag)
        if tagStr != "" {
            tokens := strings.Split(tagStr, " ")
            part := make(map[string]string)
            for i := range tokens {
                tagInfo := strings.Split(strings.Replace(tokens[i], "\"", "", -1), ":")
                if len(tagInfo) > 1 {
                    part[tagInfo[0]] = tagInfo[1]
                }
            }
            result[fieldName] = part
        }
    }
    return result
}

func WaitGo(number int, fu func()) {
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func(index int) {
            wg.Done()
        }(i)
    }
    wg.Wait()
}

func Find(slice []string, val string) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}

func UpdateStructByTagMap(result interface{}, tagName string, tagMap map[string]interface{}) error {
    t := reflect.TypeOf(result)
    if t.Kind() != reflect.Ptr {
        return fmt.Errorf("result have to be a pointer")
    }
    t = t.Elem()
    if t.Kind() != reflect.Struct {
        return fmt.Errorf("result pointer not struct")
    }
    v := reflect.ValueOf(result).Elem()
    fieldNum := v.NumField()
    for i := 0; i < fieldNum; i++ {
        fieldInfo := v.Type().Field(i)
        tag := fieldInfo.Tag.Get(tagName)
        if tag == "" {
            continue
        }
        f := v.FieldByName(fieldInfo.Name)
        if !f.IsValid() || !f.CanSet() {
            continue
        }
        value, ok := tagMap[tag]
        if !ok {
            continue
        }

        valueRealType := reflect.TypeOf(value).Kind()
        targetType := f.Kind()
        if valueRealType == targetType {
            f.Set(reflect.ValueOf(value))
            continue
        }
        expr := fmt.Sprintf("%s-to-%s", valueRealType, targetType)
        switch expr {
        case "int-to-string":
            f.SetString(strconv.Itoa(value.(int)))
        case "string-to-string":
            f.SetString(value.(string))
        case "string-to-int":
            _v, _ := strconv.Atoi(value.(string))
            f.SetInt(int64(_v))
        case "float64-to-int":
            f.SetInt(int64(value.(float64)))
            // TODO more case
        }
    }
    return nil
}
