// (c) 2016 Daniel Cortés. Licensed under the BSD license (see LICENSE).

// Package serenity.paas.utils implements handling several features related with
// Kubernetes objects.

package utils

import (
    "bufio"
    "os"
    "io/ioutil"
    "fmt"
  )

  var secretData map[string]string

//  Read a secret from mount path. Mount path shouldn't finalize in /

  func ReadSecret(path string){

    if len(path) == 0{
      fmt.Println("ERROR - Secret path is null")
      os.Exit(0)
    }

    secretData = map[string]string{}
    listDir,_ := ioutil.ReadDir(path)

    for i:=0; i<len(listDir);i++{

      if (listDir[i].IsDir()){ continue }

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
