package main

const FOUR_MILION int = 4000000

func solution() int{
    prev := 0; curr := 1; sum := 0
    for curr <= FOUR_MILION {
        temp := curr
        curr += prev
        prev = temp

        if curr%2 == 0 {
            sum += curr
        }
    }
    return sum
}

//[mine] fib recursion 1,1,2,3...
func fibRec(num int) int {
    if num == 1 || num == 2 {
        return 1
    }

    return fib(num - 1) + fib(num - 2)
}

//[mine] fib loop 1,2,3,5...
func fib(num int) int {
    prev := 0; curr := 1
    for i:=0; i<num; i++ {
        temp := curr
        curr += prev
        prev = temp
    }
    return curr
}

func main(){
    println(solution())
}