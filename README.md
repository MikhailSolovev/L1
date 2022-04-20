#8

Различие можно увидеть на следующем примере:  
p := new(chan int)   // p has type: *chan int   
c := make(chan int)  // c has type: chan int

make - создает срезы, мапы и каналы.  
new - возваращает указатель на инициализированную память.  

