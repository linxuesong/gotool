package file

import (
	"bufio"
	"os"
)

func FileToStringSliceByLine(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
 	scanner := bufio.NewScanner(file);

	ret := make([]string, 0)
	for scanner.Scan() { // Scanner默认按照/n分隔符扫描
		ret = append(ret, scanner.Text())
	}
	return ret,nil
}

func FileToStringBySpace(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) //指定分隔方式，为每次扫描一个单词，即以空格为分隔符   共有四种分隔模式 ScanLines按行、ScanWords按空格、ScanRunes读取utf-8、ScanBytes读取byte

	ret := make([]string, 0)
	for scanner.Scan() { // Scanner默认按照/n分隔符扫描
		ret = append(ret, scanner.Text())
	}
	return ret,nil
}