package test

import (
	"fmt"
	"testing"
)

// 声明单个变量
func TestVar(t *testing.T) {
	var age int // 变量声明
	fmt.Println("my age is", age)
	age = 29 // 赋值
	fmt.Println("my age is", age)
	age = 54 // 赋值
	fmt.Println("my new age is", age)
}

// 声明变量并初始化
func TestInitVar(t *testing.T) {
	var age int = 29 // 声明变量并初始化
	fmt.Println("my age is", age)
}

// 类型推断
func TestTypeInference(t *testing.T) {
	var age = 29 // 可以推断类型
	fmt.Println("my age is", age)
}

// 声明多个变量并初始化变量
func TestMultipleVar(t *testing.T) {
	var width, height int = 100, 50 // 声明多个变量
	fmt.Println("width is", width, "height is", height)
}

// 声明多个变量,省略了初始化,它们的初始值将赋值为 0
func TestMultipleVar1(t *testing.T) {
	var width, height int
	fmt.Println("width is", width, "height is", height)
	width = 100
	height = 50
	fmt.Println("new width is", width, "new height is ", height)
}



