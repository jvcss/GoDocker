package main

// We are simulating, but not giving images bcz we are using the own Linux Kernel as image through GoLang Namespaces.
// 
// docker           run image <cmd> <params>
// go run main.go   run       <cmd> <params>
// 

func main(){
  switch os.Args[1]{
    case "run":
      run()
    default:
      panic("command not implemented")
  }
}

func run(){
  fmt.Printf("Executing %v\n", os.Args[2:])
  
}

func must(err error){
  if err != nil{
    panic(err)
  }
}

