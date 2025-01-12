package house_builder

type IglooBuilder struct {
   WindowType string
   DoorType   string
   Floor      int
}

func newIglooBuilder() *IglooBuilder {
   return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType() {
   b.WindowType = "Snow Window"
}

func (b *IglooBuilder) setDoorType() {
   b.DoorType = "Snow Door"
}

func (b *IglooBuilder) setNumFloor() {
   b.Floor = 1
}

func (b *IglooBuilder) GetHouse() House {
   return House{
      DoorType:   b.DoorType,
      WindowType: b.WindowType,
      Floor:      b.Floor,
   }
}