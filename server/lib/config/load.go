package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func Load() (config *Config, err error) {
	config = &Config{}

	// config型のフィールドを舐める
	fieldsInfo := reflect.TypeOf(*config)
	valuesInfo := reflect.ValueOf(config)
	for i := 0; i < fieldsInfo.NumField(); i++ {
		// フィールドにくっついてるタグ
		fieldTag := fieldsInfo.Field(i).Tag

		// 環境変数のキーをタグから引く
		envKey := fieldTag.Get("environment_key")

		// 取得したキーで環境変数の値を取得
		rawValue := os.Getenv(envKey)
		if rawValue == "" {
			errMsg := fmt.Sprintf("設定の値が空です : キー名 %v", envKey)
			return nil, errors.New(errMsg)
		}

		// 本来の型をタグから引く
		envType := fieldTag.Get("type")

		// 本来の型別にパース処理
		switch envType {
		case "int":
			result, err := strconv.Atoi(rawValue)
			if err != nil {
				errMsg := fmt.Sprintf("設定の数値が解釈できませんでした : キー名 %v 値 %v", envKey, rawValue)
				return nil, errors.New(errMsg)
			}
			valuesInfo.Elem().Field(i).SetInt(int64(result))
			continue
		case "string":
			result := rawValue
			valuesInfo.Elem().Field(i).SetString(result)
			continue
		default:
			errMsg := fmt.Sprintf("設定として対応していない型が指定されています : キー名 %v 型名 %v", envKey, envType)
			return nil, errors.New(errMsg)
		}
	}

	return config, nil
}
