package main

import (
	"fmt"
	"leecode/leetcode"
	"leecode/linkList"
	"leecode/recursion"
	"leecode/search"
	"leecode/sort"
	"leecode/stack"
)

func main() {
	//排序
	//Sort()
	//List()
	//Stack()
	//Recursion()
	//Search()
	Leetcode()
}

func Recursion() {
	n := 10
	fmt.Printf("阶梯【%d】：共有%d种走法", n, recursion.Stage(n))
	fmt.Printf("阶梯【%d】：共有%d种走法", n, recursion.Stage2(n))
}

func Sort() {
	fmt.Println(sort.BubbleSort(slice()))
	fmt.Println(sort.InsertSort(slice()))
	fmt.Println(sort.QuickSort(slice()))
	fmt.Println(sort.ShellSort(slice()))
	fmt.Println(sort.SelectSort(slice()))
	fmt.Println(sort.MergeSort(slice()))
}

func List() {

	fmt.Println("pre", linkList.PrintListNode(head()))
	// 判断是否有环
	fmt.Println("has circle", linkList.Circle(head()))
	node := head().Next
	// 交换相邻节点
	fmt.Println("swap", linkList.PrintListNode(linkList.SwapPair(node)))
	fmt.Println("swap", linkList.PrintListNode(linkList.SwapPair2(head().Next)))
	// 反转
	fmt.Println("after", linkList.PrintListNode(linkList.Reverses(head().Next)))
}

func Stack() {
	fmt.Println(stack.IsValid("{aa[cc(aaa)dd]bb}"))
}

func slice() []int {
	return []int{1, 9, 2, 8, 3, 7, 4, 6, 5}
}

func head() *linkList.ListNode {
	head := new(linkList.ListNode)
	// 链表
	linkList.CreateListNode(head, 10)

	return head
}

func Search() {
	i := search.BinarySearch2(sort.BubbleSort(slice()), 9)
	fmt.Println("index:", i)
}

func Leetcode() {
	nums := []int{
		-1, 0, 1, 2, -1, -4,
	}
	fmt.Print(leetcode.ThreeSum(nums))
}

func calculateSum() {

}
