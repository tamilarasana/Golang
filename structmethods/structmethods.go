package main

import "fmt"

type student struct{
	name  string
	grades []int
	age int
}

func (s *student) setAge(age int){
	s.age = age
}

func (s student) getAverageGrade()float32  {
	sum := 0
	for _, v :=range  s.grades{
		sum  += v
	}
	return float32(sum) / float32(len(s.grades))
	
}

func (s student) getMaxGrade()int{
	curMax := 0
	for _, v :=range s.grades{
		if curMax < v{
			curMax = v
		}
	}
	return  curMax
}

func main()  {
	s1 := student{"tamil", []int{20,40,50},24}
	s2:= student{"sam", []int {20,40,50},24}
	s3:= student{"vicky", []int {20,40,50},24}
	s1.setAge(25)
	average := s2.getAverageGrade()
	max :=s.getMaxGrade()
	fmt.Println(s1)
	fmt.Println(average)
	fmt.Println(max)
	//s1.setage(25)
  // fmt.Println(s1)
	
}
	
