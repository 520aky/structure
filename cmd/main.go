package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {

	cmd := exec.Command("ping", "www.baidu.com")
	// 获取子进程标准输出
	stdout, _ := cmd.StdoutPipe()

	// 执行命令
	cmd.Start()

	// 读取子进程
	reader := bufio.NewReader(stdout)
	for {
		_, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		// 转换CMD的编码为GBK
		//reader := transform.NewReader(
		//	bytes.NewReader([]byte(line)),
		//	simplifiedchinese.GBK.NewDecoder(),
		//)
		d, _ := ioutil.ReadAll(reader)

		// 将子进程的内容输出
		print(string(d))
	}

	// 模拟CMD暂停
	bufio.NewReader(os.Stdin).ReadLine()
}

/*
 */
