# signum
A easy bcrypt wrapper


**Usage**

```go
package main
import "github.com/mattb2401/signum"

func main() {
  // Generate a bcyrpt hashed string
  password := "boom@s1gn@m"
  // Use strong salt rounds. Bcrypt is an adaptative password hashing function
  rounds := 12
  hash, err := signum.MakeHash(password, rounds)
  if err != nil {
    fmt.Println(err)
  }
  // Check if hash is valid 
  if signum.Attempt(password, hash) {
      return true
  } else {
      return false
  }
  // Getting the salt rounds in a hash
  rounds, err := signum.GetRounds(hash)
  if err != nil {
    fmt.Println(err)
  }
}

````
