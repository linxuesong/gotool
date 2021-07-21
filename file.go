package gotool

import "github.com/linxuesong/gotool/file"

func Add(a, b int) int {
	return a + b
}
// file文件操作类

/**
 * @Author linxuesong
 * @Description //TODO
 * @Date 8:12 下午 2021/7/12
 * @Param  * @Param: null
 * @return
 **/
func  FileToStringSliceByLine(filePath string) ([]string, error)  {
	return file.FileToStringSliceByLine(filePath)
}

/**
 * @Author linxuesong
 * @Description //TODO 
 * @Date 8:11 下午 2021/7/12
 * @Param  * @Param: null
 * @return 
 **/
func FileToStringBySpace(filePath string) ([]string, error) {
	return FileToStringBySpace(filePath)
}