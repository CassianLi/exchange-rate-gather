package utils

import (
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"log"
	"strings"
)

// XmlToStruct xml解析为struct 并指定encoding
func XmlToStruct(xmlStr string, v interface{}, encoding string) error {
	decoder := xml.NewDecoder(strings.NewReader(xmlStr))
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(v)
	if err != nil {
		log.Println("XmlToStruct err:", err)
		return err
	}
	return nil
}
