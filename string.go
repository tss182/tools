package tools

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Slug(str string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]")
	var slug = reg.ReplaceAllString(str, "-")
	reg, _ = regexp.Compile("-+")
	slug = reg.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	return strings.ToLower(slug)
}

func PhoneNumberFormat(phoneNumber string) string {
	reg, _ := regexp.Compile("[^0-9]+")
	phoneNumber = reg.ReplaceAllString(phoneNumber, "")
	reg, _ = regexp.Compile("^(08|8)")
	phoneNumber = reg.ReplaceAllString(phoneNumber, "628")
	return phoneNumber
}

func InArray(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("SliceExists() given a non-slice type")
	}
	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}
	return false
}

func Password(plaintext string) string {
	h512 := sha512.New()
	h512.Write([]byte(plaintext))
	return hex.EncodeToString(h512.Sum(nil))
}

func PasswordCheck(plaintext, hash string) bool {
	if Password(plaintext) == hash {
		return true
	}
	return false
}

func ValidationEmail(email string) bool {
	regexEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !regexEmail.MatchString(email) {
		return false
	}
	return true
}

func GenerateOtp(str string) string {
	var date = time.Now().UnixNano()
	var t = strconv.Itoa(int(date))
	var code = Password(str + t)
	var reg, _ = regexp.Compile("[^0-9]+")
	code = reg.ReplaceAllString(code, "")
	var codes = SplitN(code, 4)
	return strings.ToUpper(codes[0])
}

func SplitN(String string, n int) []string {
	number := strconv.Itoa(n)
	re := regexp.MustCompile(".{0," + number + "}")
	data := re.FindAllString(String, -1)
	return data
}

func JoinUrlPath(paths ...string) string {
	for i := range paths {
		paths[i] = strings.TrimSpace(paths[i])
		paths[i] = strings.TrimLeft(paths[i], "/")
		paths[i] = strings.TrimRight(paths[i], "/")
	}
	return strings.Join(paths, "/")
}

func PrivacyContact(str string) string {
	var result string
	isEmail := ValidationEmail(str)
	if !isEmail {
		end := 0
		front := 0
		if len(str) < 6 {
			front = 2
			end = 2
		} else {
			end = 4
			front = 3
		}
		str = PhoneNumberFormat(str)
		splitStr := strings.Split(str, "")
		for i, v := range splitStr {
			if i <= front || len(splitStr)-end < i+1 {
				result += v
			} else {
				result += "*"
			}
		}

	} else {
		show := 0
		wordSlice := strings.Split(str, "@")
		if len(wordSlice[0]) <= 5 {
			show = 2
		} else {
			show = 4
		}
		splitStr := strings.Split(wordSlice[0], "")
		for i, v := range splitStr {
			if i <= show {
				result += v
			} else {
				result += "*"
			}
		}
		result += "@" + wordSlice[1]
	}

	return result
}

func StructToJson(data interface{}) string {
	b, _ := json.Marshal(data)
	return string(b)
}

func StructToMap(data interface{}) map[string]interface{} {
	r := StructToJson(data)
	var result map[string]interface{}
	_ = json.Unmarshal([]byte(r), &result)
	return result
}

func StringToInt(data string) int {
	r, _ := strconv.Atoi(data)
	return r
}

func AddressSplit(str string, split int) (string, string, string) {
	shipperAddress := SplitN(str, split)
	var address1, address2, address3 string
	for i, v := range shipperAddress {
		switch i {
		case 0:
			address1 = v
		case 1:
			address2 = v
		case 2:
			address3 = v
		default:
			break
		}
	}
	return address1, address2, address3
}

func DecimalSeparator(i int) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d", i)
}

func RandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}
