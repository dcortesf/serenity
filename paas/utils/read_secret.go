// (c) 2016 Daniel Cortés. Licensed under the BSD license (see LICENSE).

// Package serenity.paas.utils implements handling several features related with
// Kubernetes objects.

package utils

import (
    "bufio"
    "fmt"
    "os"
    "io/ioutil"
  )

  var secretData map[string]string

  func main(){
    ReadSecret("/Users/dani/test/secrets")
    fmt.Println(GetValue("user"))
  }


/**
  Read a secret from mount path. Mount path shouldn't finalize in /
*/
  func ReadSecret(path string){

    secretData = map[string]string{}
    listDir,_ := ioutil.ReadDir(path)

    for i:=0; i<len(listDir);i++{

      f, _ := os.Open(path+"/"+listDir[i].Name())
      scanner := bufio.NewScanner(f)

      for scanner.Scan(){
        ucl := scanner.Text()
        secretData[listDir[i].Name()] = ucl
      }
    }
  }

  /**
    Get a data from a Secret
  **/

  func GetValue(key string) string{
    return secretData[key]
  }
