package main
 
import (
	"fmt"
)
 
func main() {
	fmt.Println("hello world")
	sample := []int{1, 2, 5, -3, 0, 4, -101, 21, 55, -3, 0, 4, 10, 2, 5, -3, 0, 4, 1, 2, 5, -30, 0, 4, 1, 20, 5, -3, 0, 4, 1, 2, 5, -3, 0, 4, 1, 2, 5, -3, 0, 4, 1, 2, 5, -3, 10, 4, 1, 2, 5, -3, 0, 4, 1, 2, -5, -3, 0, 9, 1, 2, 5, -3, 8, 4, 1, 2, 5, -3, 0, 4, 1, 2, 5, -3, 0, 4}
	sample = mergesort(sample)
	fmt.Println(sample)

	sample2 := []int{1, 2, 5, -3, 0, 4, -101, 21, 55, -3, 0, 4, 10, 2, 5, -3, 0, 4, 1, 2, 5, -30, 0, 4, 1, 20, 5, -3, 0, 4, 1, 2, 5, -3, 0, 4, 1, 2, 5, -3, 0, 4, 1, 2, 5, -3, 10, 4, 1, 2, 5, -3, 0, 4, 1, 2, -5, -3, 0, 9, 1, 2, 5, -3, 8, 4, 1, 2, 5, -3, 0, 4, 1, 2, 5, -3, 0, 4}
	sample2 = bubblesort(sample2)
	fmt.Println(sample2)

	sample3 := []int{1, 2, 5, -3, 0, 4, -101, 21, 55, -3, 0, 4, 10, 2, 5, -3, 0, 4, 1, 2, 5, -30, 0, 4, 1, 20, 5, -3, 0, 4, 1, 2, 5, -3, 0, 4, 1, 2, 5, -3, 0, 4, 1, 2, 5, -3, 10, 4, 1, 2, 5, -3, 0, 4, 1, 2, -5, -3, 0, 9, 1, 2, 5, -3, 8, 4, 1, 2, 5, -3, 0, 4, 1, 2, 5, -3, 0, 4}
	sample3 = shakesort(sample3)
	fmt.Println(sample3)

	sample4 := []int{1, 2, 5, -3, 0, 4, -101, 21, 55, -3, 0, 4, 10, 2, 5, -3, 0, 4, 1, 2, 5, -30, 0, 4, 1, 20, 5, -3, 0, 4, 1, 2, 5, -3, 0, 4, 1, 2, 5, -3, 0, 4, 1, 2, 5, -3, 10, 4, 1, 2, 5, -3, 0, 4, 1, 2, -5, -3, 0, 9, 1, 2, 5, -3, 8, 4, 1, 2, 5, -3, 0, 4, 1, 2, 5, -3, 0, 4}
	sample4 = insertionsort(sample4)
	fmt.Println(sample4)

	sample5 := []int{1, 2, 5, -3, 0, 4, -101, 21, 55, -3, 0, 4, 10, 2, 5, -3, 0, 4, 1, 2, 5, -30, 0, 4, 1, 20, 5, -3, 0, 4, 1, 2, 5, -3, 0, 4, 1, 2, 5, -3, 0, 4, 1, 2, 5, -3, 10, 4, 1, 2, 5, -3, 0, 4, 1, 2, -5, -3, 0, 9, 1, 2, 5, -3, 8, 4, 1, 2, 5, -3, 0, 4, 1, 2, 5, -3, 0, 4}
	sample5 = gnomesort(sample5)
	fmt.Println(sample5)

	sample6 := []int{1, 2, 5, -3, 0, 4, -101, 21, 55, -3, 0, 4, 10, 2, 5, -3, 0, 4, 1, 2, 5, -30, 0, 4, 1, 20, 5, -3, 0, 4, 1, 2, 5, -3, 0, 4, 1, 2, 5, -3, 0, 4, 1, 2, 5, -3, 10, 4, 1, 2, 5, -3, 0, 4, 1, 2, -5, -3, 0, 9, 1, 2, 5, -3, 8, 4, 1, 2, 5, -3, 0, 4, 1, 2, 5, -3, 0, 4}
	sample6 = treesort(sample6)
	fmt.Println(sample6)
}

func gnomesort(nums []int) []int {
	for i, j := 1,2; i < len(nums); {
		if nums[i-1] > nums[i] {
			nums[i-1], nums[i] = nums[i], nums[i-1]
			i--
			if i > 0 { continue }
		}
		i = j
		j++
	}
	return nums
}

func bubblesort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		f:=0
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				f = 1
			}
		}
		if f == 0 { break }
	}
	return nums
}

func insertionsort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		buffer := nums[i]
		index := i
		for ; index > 0 && nums[index-1] > buffer; index-- {
			nums[index] = nums[index-1]
		}
		nums[index] = buffer
	}
	return nums
}

func shakesort(nums []int) []int {
	var left int
	var right = len(nums) - 1
	for left <= right {
		for i := left; i < right; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
		right--
		for i := right; i > left; i-- {
			if nums[i-1] > nums[i] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
			}
		}
		left++
	}
	return nums
}

func mergesort(nums []int) []int {
	length:=len(nums)
    	if length >= 2{
    		mid := len(nums)/2
		nums = mrg(mergesort(nums[:mid]), mergesort(nums[mid:]))
	}
	return nums
}

func mrg(left []int,right []int) []int{
    //Merge two lists in ascending order.
    lst:=make([]int,0)
    for len(left) > 0 && len(right) > 0{
        if left[0] < right[0]{
            lst = append(lst,left[0])
            left = left[1:]
        }else{
            lst = append(lst,right[0])
            right = right[1:]
        }
    }
    if len(left) > 0{
        lst = append(lst,left...)
    }
    if  len(right) > 0{
        lst = append(lst,right...)
    }

    return lst
}

func mergesort2(nums []int) []int {
	length:=len(nums)
    	if length >= 2{
    		mid := len(nums)/2
		nums = mrg2(mergesort2(nums[:mid]), mergesort2(nums[mid:]))
	}
	return nums
}

func mrg2(left []int,right []int) []int{
    //Merge two lists in ascending order.
	nums:= append(left, right...)
	for i := 1; i < len(nums); i++ {
		buffer := nums[i]
		index := i
		for ; index > 0 && nums[index-1] > buffer; index-- {
			nums[index] = nums[index-1]
		}
		nums[index] = buffer
	}
	return nums
}

type TreeNode struct {
	val int
	left *TreeNode
	right *TreeNode
}

var result []int

func NewTree(val int) *TreeNode {
	return &TreeNode{
		val: val,
	}
}

func(t *TreeNode) GetResult() {
	if t == nil {
		return
	}
	t.left.GetResult()
	result = append(result, t.val)
	t.right.GetResult()
}

func(t *TreeNode) Insert(val int) {
	if val < t.val {
		if t.left == nil {
			t.left = NewTree(val)
		} else {
			t.left.Insert(val)
		}
	} else {
		if t.right == nil {
			t.right = NewTree(val)
		} else {
			t.right.Insert(val)
		}
	}
}

func treesort(nums []int) []int {
	tn := NewTree(nums[0])
	for i := 1; i < len(nums); i ++ {
		tn.Insert(nums[i])
	}
	tn.GetResult()
	return result
}
