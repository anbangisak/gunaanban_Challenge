package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/karlseguin/jsonwriter"

	"github.com/spf13/viper"
)

var jsonInput *viper.Viper

type numberVal struct {
	name  string
	value float64
}
type stringVal struct {
	name       string
	value      string
	valueInNum int64
	isEpoch    bool
}

type keyVals struct {
	numbersVal []numberVal
	stringVal  []stringVal
}

func (k keyVals) displayJson() {
	buffer := new(bytes.Buffer)
	writer := jsonwriter.New(buffer)
	writer.RootObject(func() {
		for _, numVals := range k.numbersVal {
			// numVals.name
			writer.KeyValue(numVals.name, numVals.value)
		}
		for _, strVal := range k.stringVal {
			// numVals.name
			if strVal.isEpoch {
				writer.KeyValue(strVal.name, strVal.valueInNum)
			} else {
				writer.KeyValue(strVal.name, strVal.value)
			}
		}
	})
	fmt.Println(writer.W)
}

func main() {
	fmt.Println("hi")
	var err error
	jsonReader := viper.New()
	jsonReader.SetConfigType("json")
	jsonReader.SetConfigName("inputfile")
	jsonReader.AddConfigPath("../challenge2/")
	jsonReader.AddConfigPath("challenge2/")
	err = jsonReader.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}

	keyItem := keyVals{}

	for _, key := range jsonReader.AllKeys() {
		// fmt.Println("key & value: ", key, jsonReader.GetString(key))
		if strings.HasSuffix(key, ".n") {
			value := jsonReader.GetString(key)
			nValStr := strings.TrimSpace(value)
			nValfloat, err := strconv.ParseFloat(nValStr, 8)
			if err != nil {
				fmt.Println("unable to convent number string to float and ignoring")
				continue
			}
			keySplitArr := strings.Split(key, ".")
			numValStruct := numberVal{name: keySplitArr[0], value: nValfloat}
			keyItem.numbersVal = append(keyItem.numbersVal, numValStruct)
		}
		if strings.HasSuffix(key, ".s") {
			value := jsonReader.GetString(key)
			nValStr := strings.TrimSpace(value)
			keySplitArr := strings.Split(key, ".")
			tt, err := time.Parse(time.RFC3339, nValStr)
			if err != nil {
				fmt.Println(err)
				stringValStruct := stringVal{name: keySplitArr[0], value: nValStr}
				keyItem.stringVal = append(keyItem.stringVal, stringValStruct)
			} else {
				epoch := tt.Unix()
				stringValStruct := stringVal{name: keySplitArr[0], valueInNum: epoch, isEpoch: true}
				keyItem.stringVal = append(keyItem.stringVal, stringValStruct)
			}

		}
	}

	keyItem.displayJson()

}
