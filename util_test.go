package faceeasco

import (
	"log"
	"testing"
)

func TestFileRead(t *testing.T) {
	bytes, err := FileRead("z:\\tmp\\avatar.png")
	log.Println(err)
	result, err1 := ImageToBase64(bytes)
	log.Println(err1)
	log.Println(string(result))
}

func TestNewUUID(t *testing.T) {
	log.Println(NewUUID())
	log.Println(NewUUID())
	log.Println(NewUUID())
	log.Println(NewUUID())
}
