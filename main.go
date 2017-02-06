package main

import (
	"fmt"
	"flag"
	"bufio"
	"os"
	"errors"
	"strings"
	"strconv"
	"sort"
)

var (
	ps1List  map[int]string
	charList map[int]string
	PS1      string
)

func init() {
	//todo 添加更多的提示符
	//todo 直接通过命令行参数配置来输出
	//todo 预置一些提示符样式
	//todo 颜色, powerline
	ps1List = map[int]string{0: "空格", 1: "日期", 2: "时间", 3: "Hostname", 4: "当前工作目录名", 5: "当前用户名", 6: "$"}
	charList = map[int]string{0: ` `, 1: `\d`, 2: `\t`, 3: `\H`, 4: `\w`, 5: `\u`, 6: `\$`}
}

func main() {
	test := flag.String("test", "", "test info")
	flag.Parse()
	fmt.Println(*test)
	fmt.Println("PS1 Edit")
	fmt.Println("Switch what you want:")
	var keys []int
	for k := range ps1List {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Println(k, ps1List[k], charList[k])
	}
	fmt.Println("Input like: 5 @ 3 : 4 6")
	fmt.Println(`And return: export PS1='\u@\H:\w\$ '`)
	fmt.Println("")

	tmpOrder := getInput()
	PS1 = parseInput(tmpOrder) //shell提示符属于私有环境变量, 库函数只能修改`env`中的, `set`中的私有部分无法修改
	//FIXME 自动化修改
	ps1 := os.Getenv("PS1")
	info := fmt.Sprintf("export PS1='%s'", PS1)
	if ps1 == "" {
		fmt.Println("Can't find PS1 in env")
		fmt.Println("Try to run: ")
		fmt.Println(info)
		fmt.Println("OR persistent it in your .bashrc")
	} else {
		os.Setenv("PS1", PS1)
		fmt.Println(info)
		fmt.Println("You can persistent it in your .bashrc")
	}
	fmt.Println()

}

func getInput() (order []string) {
	inputErr := errors.New("Nothing")
	input := ""
	for inputErr != nil {
		inputReader := bufio.NewReader(os.Stdin)
		input, inputErr = inputReader.ReadString('\n')
	}
	input = strings.TrimSpace(input)
	order = strings.Split(input, " ")
	return
}

func parseInput(order []string) (PS1 string) {
	for _, v := range order {
		num, err := strconv.Atoi(v)
		if err != nil {
			PS1 += v
		} else {
			if str, ok := charList[num]; ok {
				PS1 += str
			} else {
				errInfo := fmt.Errorf("Can't find key %d", num)
				fmt.Println(errInfo)
				os.Exit(-1)
			}
		}
	}
	if !strings.HasSuffix(PS1, " ") {
		PS1 += " "
	}
	return
}
