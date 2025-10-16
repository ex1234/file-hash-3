package main
import ("crypto/sha256";"encoding/hex";"fmt";"io";"os")
func main(){
  if len(os.Args)<2{fmt.Println("Usage: hash <file>");return}
  f,_:=os.Open(os.Args[1]);defer f.Close()
  h:=sha256.New(); io.Copy(h,f)
  fmt.Println(hex.EncodeToString(h.Sum(nil)))
}
