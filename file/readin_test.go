package file

import (
	"os"
	"testing"
)

// 测内容
func TestFileToStringSliceByLine(t *testing.T) {
	str, _ := os.Getwd()
	filePath := str+"/../testfile/test1.txt"

	strs, err := FileToStringSliceByLine(filePath)
	if err != nil {
		t.Errorf("tranfrom failed, err:%v",err)
	}
	t.Logf("ret:%v", strs)
}

func TestFileToStringBySpace(t *testing.T) {
	str, _ := os.Getwd()
	filePath := str+"/../testfile/test2.txt"

	strs, err := FileToStringBySpace(filePath)
	if err != nil {
		t.Errorf("tranfrom failed, err:%v",err)
	}
	t.Logf("ret:%v", strs)
}
