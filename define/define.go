package define

type Border struct {
	Left   float64
	Top    float64
	Width  float64
	Height float64
}

type Region struct {
	Id        int
	MaxPlayer int
	Border    Border
	Dummy     bool // 是否是复制场景
	Prototype int  // 原型
	Diversion bool // 分流
	Flexible  bool // 动态伸缩
}
