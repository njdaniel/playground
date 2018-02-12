fmt.Println() - prints to console with '\n'

fmt.Printf() - prints to console can format


log.Print()

log.Fatal() - Print follow by os.Exit(1)

log.Panic() - Print follow by panic() ->
    panic() - 

The Error Type:

    type error interface {
        Error() string
    }
   
   The error type is an interface type. 