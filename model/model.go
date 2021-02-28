package model

type EHttp struct {
	ID       int64  `gorm:"primary_key;auto_increment" json:"id,omitempty"`                    //
	AppName  string `gorm:"column:appname;size:20;default:''" json:"appName"`                  //
	Method   string `gorm:"column:method;size:20;default:''" json:"method"`                    //
	Ip       string `gorm:"column:ip;size:20;default:''" json:"ip"`                            //
	Path     string `gorm:"column:path;size:50;default:''" json:"path"`                        //不需要ip地址
	Nowtime  string `gorm:"column:ntime;type:datetime;default:current_timestamp" json:"ntime"` //
	Duration int    `gorm:"column:duration;default:0" json:"duration"`                         //请求时长
	Code     int    `gorm:"column:code;default:0" json:"code"`                                 //状态吗
	// Succ     int    `gorm:"column:succ;default:0" json:"succ"`                                 //状态吗 0失败，1成功
}

type ETarget struct {
	ID          int64   `gorm:"primary_key;auto_increment" json:"id,,omitempty"`                   //
	AppName     string  `gorm:"column:appname;size:20;default:''" json:"appName"`                  //
	Mem         int     `gorm:"column:mem;size:20;default:0" json:"mem"`                           //
	UMemPercent float64 `gorm:"column:umem_percent;size:20;default:0" json:"uMemPercent"`          //
	Cpu         string  `gorm:"column:cpu;size:50;default:''" json:"cpu"`                          //
	UCpuPercent int     `gorm:"column:ucpu_perent;default:0" json:"uCpuPercent"`                   //
	Nowtime     string  `gorm:"column:ntime;type:datetime;default:current_timestamp" json:"ntime"` //
	Goroutine   int     `gorm:"column:goroutine;default:0" json:"goroutine"`                       //
}
