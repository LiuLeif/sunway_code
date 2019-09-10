@0xa93c02274c32a676;

struct Data {
    x @0 :Int32;
    y @1 :Int32;
    enum Color {
         red @0;
         green @1;
         blue @2;
    }
    struct InnerData {
           x @0 :Text;
           y @1 :Int32; 
    }
    innerData @2 : InnerData;
    listData @3 :List(Int32);
    color @4 :Color;
}