package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"sort"
	"time"
)

func Find(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

//双色球
func randSSQ(redNum, blueNum int) {
	var blueBall = make([]int, 0) //蓝球
	var redBall = make([]int, 0)  //红球
	var filterBlueBall = []int{06, 12, 15, 18, 19, 20, 21, 23, 24, 25, 32, 33}
	var filterRedBall = []int{15, 13, 14, 16}

	for i := 1; i <= 33; i++ {
		_, found := Find(filterBlueBall, i)
		if !found {
			blueBall = append(blueBall, i)
		}

	}
	for i := 1; i <= 16; i++ {
		_, found := Find(filterRedBall, i)
		if !found {
			redBall = append(redBall, i)
		}
	}

	ball := make([]int, 0)
	//蓝球随机6个
	for i := 1; i <= redNum; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano())) //随机数种子
		vindex := r.Intn(len(blueBall))
		ball = append(ball, blueBall[vindex])
		blueBall = append(blueBall[:vindex], blueBall[vindex+1:]...) //移除已经选过的数
	}
	//排序
	sort.Ints(ball)

	ball2 := make([]int, 0)
	//红球随机一个
	for i := 1; i <= blueNum; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano())) //随机数种子
		vindex := r.Intn(len(redBall))
		ball2 = append(ball2, redBall[vindex])
		redBall = append(redBall[:vindex], redBall[vindex+1:]...) //移除已经选过的数
	}
	sort.Ints(ball2)
	//打印
	//fmt.Printf("\033[1;31;40m%02d %02d %02d %02d %02d %02d\033[0m | \033[1;34;40m%02d\033[0m \n", ball[0], ball[1], ball[2], ball[3], ball[4], ball[5], ball[6])
	for i := 0; i < len(ball); i++ {
		fmt.Printf("\033[1;31;40m%02d \033[0m", ball[i])
	}
	fmt.Printf(" | ")
	for i := 0; i < len(ball2); i++ {
		fmt.Printf("\033[1;34;40m%02d \033[0m", ball2[i])
	}
	fmt.Println("")

}

//大乐透
func randDLT(redNum, blueNum int) {
	var blueBall = make([]int, 0) //蓝球
	var redBall = make([]int, 0)  //红球
	var filterBlueBall = []int{4, 13, 14, 19, 18, 31, 35}
	var filterRedBall = []int{7, 11, 10}

	for i := 1; i <= 35; i++ {
		_, found := Find(filterBlueBall, i)
		if !found {
			blueBall = append(blueBall, i)
		}

	}
	for i := 1; i <= 12; i++ {
		_, found := Find(filterRedBall, i)
		if !found {
			redBall = append(redBall, i)
		}
	}

	ball := make([]int, 0)
	//蓝球随机5个
	for i := 1; i <= redNum; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano())) //随机数种子
		vindex := r.Intn(len(blueBall))
		ball = append(ball, blueBall[vindex])
		blueBall = append(blueBall[:vindex], blueBall[vindex+1:]...) //移除已经选过的数
	}
	//蓝球排序
	sort.Ints(ball)

	ball2 := make([]int, 0)
	//红球随机2个
	for i := 1; i <= blueNum; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano())) //随机数种子
		vindex := r.Intn(len(redBall))
		ball2 = append(ball2, redBall[vindex])
		redBall = append(redBall[:vindex], redBall[vindex+1:]...) //移除已经选过的数
	}
	//红球排序
	sort.Ints(ball2)

	//打印
	//fmt.Printf("\033[1;31;40m%02d %02d %02d %02d %02d\033[0m |\033[1;34;40m%02d %02d\033[0m \n", ball[0], ball[1], ball[2], ball[3], ball[4], ball[5], ball[6])
	for i := 0; i < len(ball); i++ {
		fmt.Printf("\033[1;31;40m%02d \033[0m", ball[i])
	}
	fmt.Printf(" | ")
	for i := 0; i < len(ball2); i++ {
		fmt.Printf("\033[1;34;40m%02d \033[0m", ball2[i])
	}
	fmt.Println("")

}

//清屏
func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	for {
		var lottType int
		var lottTypeName string
		var num int
		fmt.Println("请选择彩票类型：1 大乐透，2 双色球")
		fmt.Scan(&lottType)
		if lottType == 1 {
			lottTypeName = "大乐透"
		} else if lottType == 2 {
			lottTypeName = "双色球"
		} else {
			fmt.Println("无效的彩票类型，请重新输入！")
			continue
		}
		fmt.Println("请输入需要机选的注数：")
		fmt.Scan(&num)
		clearScreen()
		fmt.Printf("机选%d注%s:\n", num, lottTypeName)

		for i := 0; i < num; i++ {
			//暂停下 否则以时间作为随机种子 有重复
			time.Sleep(100 * time.Microsecond)
			if lottType == 1 {
				randDLT(5, 2)
			} else if lottType == 2 {
				randSSQ(6, 1)
			}
		}
		fmt.Println("=========================================")
	}
}

// 格式：\033[显示方式;前景色;背景色m

// 说明：
// 前景色            背景色           颜色
// ---------------------------------------
// 30                40              黑色
// 31                41              红色
// 32                42              绿色
// 33                43              黃色
// 34                44              蓝色
// 35                45              紫红色
// 36                46              青蓝色
// 37                47              白色
// 显示方式           意义
// -------------------------
// 0                终端默认设置
// 1                高亮显示
// 4                使用下划线
// 5                闪烁
// 7                反白显示
// 8                不可见

// 例子：
// \033[1;31;40m
// \033[0m
