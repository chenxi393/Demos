package house_builder

type NormalBuilder struct {
   WindowType string
   DoorType   string
   Floor      int
}

func newNormalBuilder() *NormalBuilder {
   return &NormalBuilder{}
}

func (b *NormalBuilder) setWindowType() {
   b.WindowType = "Wooden Window"
}

func (b *NormalBuilder) setDoorType() {
   b.DoorType = "Wooden Door"
}

func (b *NormalBuilder) setNumFloor() {
   b.Floor = 2
}

func (b *NormalBuilder) GetHouse() House {
   return House{
      DoorType:   b.DoorType,
      WindowType: b.WindowType,
      Floor:      b.Floor,
   }
}