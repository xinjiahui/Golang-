package main

import "fmt"

// 有一个物件：
// 1.它保存了一些数据 --->结构体的字段
// 2.它有三个功能 --->结构体的方法

type student struct {
	id   int64
	name string
}

//造一个学生的管理者

type studentMgr struct {
	allStudent map[int64]student
}

var (
	allStudent map[string]student
)

//查看学生
func (s studentMgr) showStudent() {
	//从s.allstudent  这个map中把所有的学生逐个拎出来
	for _, stu := range s.allStudent { //stu是具体每个学生
		fmt.Printf("学号: %d 姓名： %s\n", stu.id, stu.name)
	}
}

//添加学生
func (s studentMgr) addStudent() {

	//1.根据用户的输入的内容，创建一个新的学生
	//把新的学生放到s.allstudent的map中
	var (
		stuID   int64
		stuName string
	)
	//获取用户输入：
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&stuName)
	//根据用户输入，创建结构体对象
	newStu := student{
		id:   stuID,
		name: stuName,
	}
	//2.把新的学生放到s.allstudent这个map中
	s.allStudent[newStu.id] = newStu

}

//修改学生
func (s studentMgr) editStudent() {
	//1.获取用户输入的学号
	var stuID int64
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)
	//2.展示该学号对应的学生信息
	stuObj, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("你要修改的学生信息如下：学号:%d  姓名:%s\n", stuObj.id, stuObj.name)
	//3.输入修改后的学生姓名
	fmt.Println("请输入学生的新名字：")
	var newName string
	fmt.Scanln(&newName)
	//4.更新学生的姓名
	stuObj.name = newName
	s.allStudent[stuID] = stuObj

}

//删除学生
func (s studentMgr) deleteStudent() {
	// 1.请用户输入要删除的学生id
	var stuID int64
	fmt.Print("请输入要删除的学生学号：")
	fmt.Scanln(&stuID)

	// 2.去map中找有没有学生，如果没有，打印查无此人提示
	_, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	// 3.有的话，删除
	delete(s.allStudent, stuID)
	fmt.Println("删除成功")
}
