###### #7 В какой последовательности будут выведены элементы map[int]int?

    В любой

###### #8 В чем разница make и new?

    Различие можно увидеть на следующем примере:  
    p := new(chan int)   // p has type: *chan int   
    c := make(chan int)  // c has type: chan int
    
    make - создает срезы, мапы и каналы.  
    new - возваращает указатель на инициализированную память.  

###### #9 Сколько существует способов задать переменную типа slice или map?

    1. a := make(map[int]int)
    2. a := map[int]int{}
    3. var a = map[int]int{}
    4. a := new(map[int]int)
       *a = map[int]int{} 
    5. var a map[int]int
       a = map[int]int{}

    Аналогично с slice

###### #10 Что выведет данная программа и почему?
    
    

