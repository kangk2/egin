package lib

import (
    "fmt"
    "log"
    "reflect"
    "strings"
    "sync"
)

// 反射调用结构体方法
func Invoke(any interface{}, name string, args ...interface{}) (reflect.Value, error) {
    method := reflect.ValueOf(any).MethodByName(name)
    notExist := method == reflect.Value{}
    if notExist {
        return reflect.ValueOf(nil), fmt.Errorf("Method %s not found ", name)
    }
    methodType := method.Type()
    numIn := methodType.NumIn()
    if numIn > len(args) {
        return reflect.ValueOf(nil), fmt.Errorf("Method %s must have minimum %d params. Have %d ", name, numIn, len(args))
    }
    if numIn != len(args) && !methodType.IsVariadic() {
        return reflect.ValueOf(nil), fmt.Errorf("Method %s must have %d params. Have %d ", name, numIn, len(args))
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
            return reflect.ValueOf(nil), fmt.Errorf("Method %s. Param[%d] must be %s. Have %s ", name, i, inType, argValue.String())
        }
        argType := argValue.Type()
        if argType.ConvertibleTo(inType) {
            in[i] = argValue.Convert(inType)
        } else {
            return reflect.ValueOf(nil), fmt.Errorf("Method %s. Param[%d] must be %s. Have %s ", name, i, inType, argType)
        }
    }
    return method.Call(in)[0], nil
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

// 通过反射, 将map中的val更新到struct对应的field中
func UpdateStructByMap(structName interface{}, defaultVal map[string]interface{}) {
    t := reflect.TypeOf(structName)
    if t.Kind() == reflect.Ptr {
        t = t.Elem()
    }
    if t.Kind() != reflect.Struct {
        log.Println("Check type error not Struct")
        return
    }

    ps := reflect.ValueOf(structName)
    // struct
    s := ps.Elem()

    for k, v := range defaultVal {
        f := s.FieldByName(k)
        if f.IsValid() {
            if f.CanSet() {
                switch f.Kind() {
                case reflect.Int:
                    x := v.(int64)
                    if !f.OverflowInt(x) {
                        f.SetInt(x)
                    }
                case reflect.String:
                    f.SetString(v.(string))
                }
            }
        }
    }
}

func CompareDefaultTag(any interface{}) {
    t := reflect.TypeOf(any)
    if t.Kind() == reflect.Ptr {
        t = t.Elem()
    }
    if t.Kind() != reflect.Struct {
        log.Println("Check type error not Struct")
        return
    }
    fields := GetStructFieldsName(any)
    tags := GetStructTags(any)
    defaultVal := make(map[string]interface{})
    for i := range fields {
        fieldName := fields[i]
        if fieldDefaultVal, hasKey := tags[fieldName]["default"]; hasKey {
            defaultVal[fieldName] = fieldDefaultVal
        }
    }
    UpdateStructByMap(any, defaultVal)
    fmt.Println(any)
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
