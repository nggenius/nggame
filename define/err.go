package define

const (
	ERR_CREATE_TIMEOUT          = iota + 100 // 创建角色超时
	ERR_CHOOSE_ROLE                          // 选择角色出错
	ERR_CHOOSE_TIMEOUT                       // 选择角色超时
	ERR_REGION_NOT_FOUND                     // 场景未找到
	ERR_ENTER_REGION_FAILED                  // 进入场景失败
	ERR_REGION_OBJECT_NOT_FOUND              // 场景对象没有找天
)

const (
	ERR_REGION_NONE          = 12000 + iota
	ERR_REGION_CREATE_FAILED // 创建场景失败
)
