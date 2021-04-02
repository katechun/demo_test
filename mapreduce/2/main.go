package main

import "fmt"

type Employee struct {
	Name string
	Age int
	Vacation int
	Salary int
}

var list = []Employee{
	{"Hao", 44, 0, 8000},
	{"Bob", 34, 10, 5000},
	{"Alice", 23, 5, 9000},
	{"Jack", 26, 0, 4000},
	{"Tom", 48, 9, 7500},
	{"Marry", 29, 0, 6000},
	{"Mike", 32, 8, 4000},
}

func EmployeeCountIf(list []Employee,fn func(e *Employee) bool)int{
	count := 0
	for i,_:=range list {
		if fn(&list[i]){
			count += 1
		}

	}
	return count
}

func EmployeeFilterIn(list []Employee,fn func(e *Employee)bool)[]Employee{
	var newList []Employee
	for i,_:=range list {
		if fn(&list[i]){
			newList = append(newList,list[i])
		}
	}
	return newList
}

func EmployeeSumIf(list []Employee,fn func(e *Employee)int)int{
	var sum = 0
	for i,_:=range list{
		sum += fn(&list[i])
	}

	return sum
}

func main(){
	//老年人员工
	old := EmployeeCountIf(list,func(e *Employee)bool{
		return e.Age > 40
	})
	fmt.Println("old people:",old)

	//工资大于6000的员工
	highPay := EmployeeCountIf(list,func(e *Employee)bool{
		return e.Salary >= 6000
	})

	fmt.Println("High Salary people:",highPay)

	//没有休假的员工
	noVecation := EmployeeFilterIn(list,func(e *Employee)bool{
		return e.Vacation == 0
	})

	fmt.Println("People no vacation:",noVecation)

	//统计所有员工工资总和
	totalPay := EmployeeSumIf(list,func(e *Employee)int{
		return e.Salary
	})

	fmt.Println("Total Salary:",totalPay)

	//年龄小于30的员工工资总和
	youngerPay := EmployeeSumIf(list,func(e *Employee)int{
		if e.Age < 30{
			return e.Salary
		}
		return 0
	})

	fmt.Println("The age < 30 Total salaey:",youngerPay)
}

