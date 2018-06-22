package counter

// import (
// 	"fmt"
// 	"os"
// 	"time"

// 	"github.com/octo-tentacle/pkg/messaging"

// 	"github.com/stianeikeland/go-rpio"
// )

// const pin = rpio.Pin(18)

// func StartCounter(messenger messaging.Messenger){
// 	go func(){
// 		// Open and map memory to access gpio, check for errors
// 		if err := rpio.Open(); err != nil {
// 				fmt.Println(err)
// 				os.Exit(1)
// 		}

// 		// Set pin to output mode
// 		pin.Output()
// 		// Unmap gpio memory when done
// 		defer rpio.Close()

// 		for i := 0; true; i++ {
// 			oof := i%2
// 			if(oof > 0){
// 				pin.Write(rpio.High)
// 			}else{
// 				pin.Write(rpio.Low)
// 			}
// 			messenger.WriteToChannel("counter", fmt.Sprintf("State: %d", oof))
//   			time.Sleep(time.Second)
// 		}
// 	}()
// }
