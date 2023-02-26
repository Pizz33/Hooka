package hooka

import "github.com/D3Ext/Hooka/core"

func CreateRemoteThread(shellcode []byte) (error) {
  return core.CreateRemoteThread(shellcode)
}

func CreateProcess(shellcode []byte) (error) {
  return core.CreateProcess(shellcode)
}

func Fibers(shellcode []byte) (error) {
  return core.Fibers(shellcode)
}

func EarlyBirdApc(shellcode []byte) (error) {
  return core.EarlyBirdApc(shellcode)
}

func UuidFromString(shellcode []byte) (error) {
  return core.UuidFromString(shellcode)
}

/*

Hell's Gate + Halo's Gate functions (WIP)

*/

func CreateProcessHalos(shellcode []byte) (error) {
  return core.CreateProcessHalos(shellcode)
}


