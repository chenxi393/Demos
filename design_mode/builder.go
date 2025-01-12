package main

// // 待构造的产品
// type Product struct {
// 	CollectResult  string
// 	AssembleResult string
// 	GenerateResult string
//  }

//  // 建造分步骤流程
//  type ProductBuilder interface {
// 	CollectMaterial()
// 	Assemble()
// 	Generate()
// 	Return() Product
//  }

//  // 玩具工厂建造者
//  type ToyProductBuilder struct {
// 	Pro Product
//  }

//  func (tb *ToyProductBuilder) CollectMaterial() {
// 	tb.Pro.CollectResult = "Collect Toy Success"
//  }
//  func (tb *ToyProductBuilder) Assemble() {
// 	tb.Pro.AssembleResult = "Assemble Toy Success"
//  }
//  func (tb *ToyProductBuilder) Generate() {
// 	tb.Pro.GenerateResult = "Generate Toy Success"
//  }
//  func (tb *ToyProductBuilder) Return() Product {
// 	return tb.Pro
//  }

//  // 食品工厂建造者
//  type FoodProductBuilder struct {
// 	Pro Product
//  }

//  func (fb *FoodProductBuilder) CollectMaterial() {
// 	fb.Pro.CollectResult = "Collect Food Success"
//  }
//  func (fb *FoodProductBuilder) Assemble() {
// 	fb.Pro.AssembleResult = "Assemble Food Success"
//  }
//  func (fb *FoodProductBuilder) Generate() {
// 	fb.Pro.GenerateResult = "Generate Food Success"
//  }
//  func (fb *FoodProductBuilder) Return() Product {
// 	return fb.Pro
//  }

//  // 建造者统一调用,外部调用其build方法，传入对应的建造者就能得到想要的产品
//  var Manager ProductBuilderManager

//  type ProductBuilderManager struct {
// 	Builder ProductBuilder
//  }

//  func (pbm *ProductBuilderManager) Build(builder ProductBuilder) Product {
// 	builder.CollectMaterial()
// 	builder.Generate()
// 	builder.Generate()
// 	return builder.Return()
//  }
