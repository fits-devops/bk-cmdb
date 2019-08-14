package metadata


// 定义创建多个班级时需要传入的数据结构
type CreateManyClass struct {
	Data []Class `json:"datas"`
}

// 定义创建单个班级时需要传入的数据结构
type CreateOneClass struct {
	Data Class `json:"data"`
}

// 定义设置多个班级时需要传入的数据结构
type SetManyClass CreateManyClass

// 定义设置单个班级时需要传入的数据结构
type SetOneClass CreateOneClass


// 定义查询班级信息需返回的数据结构
type QueryClassDataResult struct {
	Count int64            `json:"count"`
	Info  []Class		 `json:"info"`
}

// 定义删除班级信息返回的数据结构
type DeleteClassResult struct {
	BaseResp `json:",inline"`
	Data     DeletedCount `json:"data"`
}



// 定义创建学生时需要传入的数据结构
type CreateStudent struct {
	Std       Student      `json:"std"`
	Cls 	  Class 	   `json:"class"`
}

// 定义设置学生信息时需要传入的数据结构
type SetStudent CreateStudent


// 定义查询学生信息需返回的数据结构
type QueryStudentDataResult struct {
	Count int64    `json:"count"`
	Info  []Student `json:"info"`
}

// 定义查询学生信息及班级信息需返回的数据结构
type QueryStudentWithClassDataResult struct {
	Count int64             `json:"count"`
	Info  []SearchStudentInfo `json:"info"`
}

type SearchStudentInfo struct {
	Stu       Student       `json:"student"`
	Cls 	  Class 		`json:"class"`
}
