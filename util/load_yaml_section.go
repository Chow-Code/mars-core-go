package util

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
)

// LoadYAMLSection 读取yaml文件某个特定部分
func LoadYAMLSection(path string, section string, out interface{}) error {
	// 读取 YAML 文件
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("[yaml] 读取yaml文件失败: %w", err)
	}

	// 创建一个 map 来存储整个 YAML 文件的数据
	var config map[string]interface{}

	// 解析 YAML 文件
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return fmt.Errorf("[yaml] 解析yaml失败: %w", err)
	}

	// 提取特定部分
	sectionData, ok := config[section]
	if !ok {
		return fmt.Errorf("[yaml] 部分 %s 无法在YAML中找到", section)
	}

	// 将 sectionData 转换为字节切片，以便再次解析
	sectionBytes, err := yaml.Marshal(sectionData)
	if err != nil {
		return fmt.Errorf("[yaml] 无法解析该部分的数据: %w", err)
	}

	// 确保 out 是一个指针
	outValue := reflect.ValueOf(out)
	if outValue.Kind() != reflect.Ptr || outValue.IsNil() {
		return fmt.Errorf("[yaml] 输出值必须是一个指针才行")
	}

	// 解析特定部分到目标结构体
	err = yaml.Unmarshal(sectionBytes, out)
	if err != nil {
		return fmt.Errorf("[yaml] 无法解析部分数据: %w", err)
	}

	return nil
}
