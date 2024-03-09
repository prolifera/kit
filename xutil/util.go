package xutil

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"github.com/google/uuid"
	"math"
	"reflect"
	"strings"
	"time"

	"github.com/tristan-club/kit/log"
	"golang.org/x/crypto/sha3"
)

func First(f interface{}, second interface{}) interface{} {
	return f
}

func MGetS(mm interface{}, key string) (string, bool) {
	m, ok := mm.(map[string]interface{})
	if !ok {
		return "", false
	}

	if v, ok := m[key]; ok {
		if r, ok := v.(string); ok {
			return r, true
		}
	}
	return "", false
}

func MGetSDefault(m map[string]interface{}, key string, def string) string {
	s, ok := MGetS(m, key)

	if !ok {
		return def
	}

	return s
}

func MGetF(m map[string]interface{}, key string) (float64, bool) {
	if v, ok := m[key]; ok {
		if r, ok := v.(float64); ok {
			return r, true
		}
	}
	return 0, false
}

func MGetB(m map[string]interface{}, key string) (bool, bool) {
	if v, ok := m[key]; ok {
		if r, ok := v.(bool); ok {
			return r, true
		}
	}
	return false, false
}

func MGetFDefault(m map[string]interface{}, key string, def float64) float64 {
	f, ok := MGetF(m, key)

	if !ok {
		return def
	}

	return f
}

func MaxInt(a int32, b int32) int32 {
	if a >= b {
		return a
	}

	return b
}

func MinInt(a int32, b int32) int32 {
	if a <= b {
		return a
	}

	return b
}

func MinDuration(a time.Duration, b time.Duration) time.Duration {
	if a.Seconds() <= b.Seconds() {
		return a
	}
	return b
}

func Distance(a, b, oa, ob int) int {
	return int(math.Abs(float64(a-oa)) + math.Abs(float64(b-ob)))
}

func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}

func MapToStruct(m interface{}, out interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		log.Error().Msgf("MapToStruct Error %s", err)
		return err
	}

	err = json.Unmarshal(data, out)
	if err != nil {
		log.Error().Msgf("MapToStruct Error %s", err)
		return err
	}

	return nil
}

func FastMarshal(input interface{}) string {
	b, _ := json.Marshal(input)
	return string(b)
}

func HashStr(input string) []byte {
	h := md5.New()
	h.Write([]byte(input))
	return h.Sum(nil)
}

func IsNil(c interface{}) bool {
	return c == nil || (reflect.ValueOf(c).Kind() == reflect.Ptr && reflect.ValueOf(c).IsNil())
}

func ParsePayload(input interface{}) map[string]interface{} {
	resp := map[string]interface{}{}
	if IsNil(input) {
		resp["payload"] = "empty"
		return resp
	}
	b, err := json.Marshal(input)
	if err != nil {
		resp["parse error"] = err.Error()
	}
	if err := json.Unmarshal(b, &resp); err != nil {
		resp["parse error"] = err.Error()
	}
	resp["pin_code"] = ""
	return resp
}

func EIP55Checksum(unchecksummed string) (string, error) {
	v := []byte(Remove0x(strings.ToLower(unchecksummed)))

	_, err := hex.DecodeString(string(v))
	if err != nil {
		return unchecksummed, err
	}

	sha := sha3.NewLegacyKeccak256()
	_, err = sha.Write(v)
	if err != nil {
		return unchecksummed, err
	}
	hash := sha.Sum(nil)

	result := v
	for i := 0; i < len(result); i++ {
		hashByte := hash[i/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if result[i] > '9' && hashByte > 7 {
			result[i] -= 32
		}
	}
	val := string(result)
	return "0x" + val, nil
}

func Remove0x(input string) string {
	if strings.HasPrefix(input, "0x") {
		return input[2:]
	}
	return input
}

func InArrayString(param string, array []string) bool {
	for _, v := range array {
		if param == v {
			return true
		}
	}
	return false
}

func InArrayInt(param int64, array []int64) bool {
	for _, v := range array {
		if param == v {
			return true
		}
	}
	return false
}

// Different 求数组差集 属于arr1不属于arr2
func Different(arr1, arr2 []uint32) []uint32 {
	arr2Map := map[uint32]bool{}
	differentArr1 := []uint32{}

	for _, num := range arr2 {
		arr2Map[num] = true
	}

	for _, num := range arr1 {
		if !arr2Map[num] {
			differentArr1 = append(differentArr1, num)
		}
	}
	return differentArr1
}

// Intersect 求数组交集
func Intersect(arr1, arr2 []uint32) []uint32 {
	arr1Map := map[uint32]bool{}
	intersectArr := []uint32{}

	for _, num := range arr1 {
		arr1Map[num] = true
	}

	for _, num := range arr2 {
		if arr1Map[num] {
			intersectArr = append(intersectArr, num)
		}
	}
	return intersectArr
}

func GetBinBitIsOneArr(bit, maxBit int) []int {
	//todo：暂时枚举 maxBit=2 的情况
	if maxBit == 2 {
		switch bit {
		case 1:
			return []int{1, 3}
		case 2:
			return []int{2, 3}
		}
	}
	return nil
}

func UniqueArrayUint(arr []uint32) []uint32 {
	arrMap := map[uint32]bool{}
	result := []uint32{}
	for _, i := range arr {
		if !arrMap[i] {
			arrMap[i] = true
			result = append(result, i)
		}
	}
	return result
}

func GetTheMinutes(t *time.Time, ts int64) time.Time {
	if t == nil {
		ti := time.Unix(ts, 0)
		t = &ti
	}

	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
}

func ParseArrStr(s string) []uint {
	arr := []uint{}
	_ = json.Unmarshal([]byte("["+s+"]"), &arr)
	return arr
}

func GenerateUuid(format bool) string {
	uuidValue := uuid.New().String()
	if format {
		uuidValue = strings.Replace(uuidValue, "-", "", -1)
	}
	return uuidValue
}
