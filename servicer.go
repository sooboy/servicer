package servicer

type Function = func();

type Servicer  interface {
	Run()
}

// 基础函数
type Service struct {
	 name string 
	 fn   Function
}

func (this *Fun)Run(){
	 this.fn()
}

// LogService
type LogService struct {
	service Servicer
}

func (this *LogService)Run(){
	
}


