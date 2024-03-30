package utils

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/samber/lo"
)

func GetArrayItemByIndex[T any](array []T, index int, defaultValue T) T {
	result, err := lo.Nth(array, index)
	if err != nil {
		return defaultValue
	}
	return result
}

func ConsoleConfirm(message string, waitingTime time.Duration) bool {
	fmt.Println(message)
	duration := waitingTime * time.Second
	resultChan := make(chan bool)
	restartChan := make(chan bool)
	stopChan := make(chan bool)
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		go func() {
			_ = keyboard.Close()
		}()
	}()

	//倒计时
	go func() {
		for remaining := duration; remaining > 0; remaining -= time.Second {
			select {
			case <-stopChan:
				return
			default:
				fmt.Printf("\r[Y/N](回车默认值 Y, 倒计时%ds结束值 N): ", remaining.Round(time.Second)/time.Second)
				time.Sleep(time.Second)
			}
		}
		resultChan <- false
		stopChan <- true
	}()

	//用户输入
	go func() {
		choice, _, _ := keyboard.GetKey()
		fmt.Print(string(choice))
		if choice == 'y' || choice == 'Y' || choice == 0 {
			resultChan <- true
		} else if choice == 'n' || choice == 'N' {
			resultChan <- false
		} else {
			fmt.Println(" 输入的值无效，请重新输入")
			restartChan <- true
		}
		stopChan <- true
	}()

	select {
	case result := <-resultChan:
		fmt.Println()
		return result
	case <-restartChan:
		return ConsoleConfirm(message, waitingTime)
	}
}
