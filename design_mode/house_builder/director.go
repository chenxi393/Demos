package house_builder

type Director struct {
	builder IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b IBuilder) {
	d.builder = b
}
// 使用场景


// - Tabs 有多种：精选、大促、活动、商品……
// 每种 tab 的生成步骤复杂，具体的字段内容繁多，改版之前各个字段的赋值散落在各处不易寻找
// 考虑使用 builder + director 生成器 + 主管 的设计模式重构

// Director 主管：收敛管理生成器的步骤顺序，输出最终的产品
func (d *Director) BuildHouse() House {
	// 并发执行所有tab 的 Build 步骤， 每个步骤中间可加打点
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.GetHouse()
}

// - 由 home_director 主管来收敛用哪些 tab 的 builder，怎么执行 builder 的步骤