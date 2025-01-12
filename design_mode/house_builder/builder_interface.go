package house_builder

// House 具体的产品
type House struct {
   WindowType string
   DoorType   string
   Floor      int
}

type IBuilder interface {
   setWindowType()
   setDoorType()
   setNumFloor()
   GetHouse() House
}

func GetBuilder(builderType string) IBuilder {
   if builderType == "normal" {
      return newNormalBuilder()
   }

   if builderType == "igloo" {
      return newIglooBuilder()
   }
   return nil
}