package main

//================================ 版本 version  1.1 ===================================================
// GOlang V1.20 版本,安装与配置 ，请提问Ai
//  GOlang  V1.20 version, installation and configuration, please ask AI.
//MIT 开源授权许可 MIT open source license
//==================================  设置与目录 Configuration and Directories    ===================
// 程序各项配置 ： = "英语English"
// Program configuration ： = "英语English"
// 程序响应的功能目录 ： "初始:")
//Program Response Function Directory  ： "初始:")
//===================================================================================


import (
	"encoding/json"

	"fmt"

	"path/filepath"

	"strings"
	"syscall"
	"time"

	"crypto/md5"

	"unicode"

	"io/ioutil"
	"os"
	"net"
	"net/http"

	"log"

	"os/exec"

	"github.com/atotto/clipboard"
	"sort"
	"strconv"

	"sync"
	"runtime"

	"encoding/hex"

	"github.com/axgle/mahonia"
)

type T1T struct {
	Q11Q	string
	Q8Q	string
	Q10Q	string
	Q6Q	int
	Q3Q	string

	Q1Q	string
	Q9Q	string

	Q4Q	map[string]string

	Q7Q	[]string

	Q5Q	[]string
	Q2Q	string
	Q14Q		sync.Mutex
}

var S2S = new(T1T)

type T2T struct {
	A0_str	string

	L1_str	string
	L2_str	string
	L3_str	string
	L5_str	string
	I1_str	string
	T2_str	string

	T1_str	string
	T3_str	string

	C2_str	string
	C3_str	string
	C5_str	string
	C6_str	string

	M9_str	string
}

var S4S = new(T2T)

type T3T struct {
	W1_strs	[]string
	W2_map	map[string]int
}

type V6V struct {
	Q13Q	string
	Q12Q	int
}

var S7S = new(T3T)

func main() {

	F0F()

	go F19F()

	select {}

}

func F0F() {

	S2S.Q8Q, _ = os.Getwd()

	S2S.Q11Q = "汉语"
	S2S.Q11Q = "英语English"

	var V9V = make(map[string]string)

	if S2S.Q11Q == "汉语" {

		V9V["语法查询"] = "--语法查询-V1.2"
		V9V["防火墙放行-手动打开"] = "防火墙放行-手动打开语法查询窗口.html"
		V9V["记录文件名"] = "0-记录.ini"
	} else {

		V9V["语法查询"] = "--Grammar Query-V1.2"
		V9V["防火墙放行-手动打开"] = "Firewall Allow - Manually Open Grammar Query Window.html"
		V9V["记录文件名"] = "0-Record.ini"

	}

	S2S.Q10Q = filepath.Base(S2S.Q8Q) + V9V["语法查询"]

	S2S.Q1Q = filepath.Join(S2S.Q8Q, V9V["防火墙放行-手动打开"])

	S7S.W2_map = make(map[string]int)

	S2S.Q2Q = fmt.Sprintf("%s/%s", S2S.Q8Q, V9V["记录文件名"])
	var _, V10V = os.Stat(S2S.Q2Q)

	if V10V != nil {

		S7S.W1_strs = F2F()
	} else {

		if V11V := F11F(S2S.Q2Q, &S7S); V11V != nil {

		}
		if len(S7S.W1_strs) == 0 {
			S7S.W1_strs = F2F()
		}
		S7S.W2_map = F12F(S7S.W2_map)

	}

	S2S.Q4Q = make(map[string]string)

	F3F(S2S.Q8Q, 0)

	return

}

func F1F(V30V string) {

	V30V = strings.TrimSpace(V30V)

	var V14V string

	_ = V14V
	var V16V = make(map[string]string)

	if S2S.Q11Q == "汉语" {
		V16V["添加后缀名说明"] = "在此添加可搜索的语法文件后缀名..."

	} else {
		V16V["添加后缀名说明"] = "Add Searchable Syntax File Extensions Here..."

	}

	defer func() {

	}()

	if strings.HasPrefix(V30V, "初始:") {
		S2S.Q6Q++

		S4S.T2_str = strings.Replace(F34F(), "|换行|", "<br />", -1)
		S4S.I1_str = V16V["添加后缀名说明"] + "\n" + strings.Join(S7S.W1_strs, "")

		S4S.L3_str = F31F(S7S.W2_map)
		S2S.Q3Q = "前端"

	} else if strings.HasPrefix(V30V, "搜索词语:") {
		_, V30V, _ = strings.Cut(V30V, ":")

		var V19V int
		S4S.L1_str, V19V = F30F(V30V)

		if S2S.Q11Q == "汉语" {
			V16V["搜索完成语"] = `搜索 <span style="color: rgb(194, 173, 214);">【%s】</span> 词语完成,共计文件数:<span style="color: rgb(194, 173, 214);">&nbsp;&nbsp;%d</span>`
		} else {
			V16V["搜索完成语"] = `search <span style="color: rgb(194, 173, 214);">【%s】</span> Words completed, total number of documents:<span style="color: rgb(194, 173, 214);">&nbsp;&nbsp;%d</span>`

		}
		S4S.T2_str = fmt.Sprintf(V16V["搜索完成语"], V30V, V19V)

	} else if strings.HasPrefix(V30V, "复制文件内容:") {
		_, V30V, _ = strings.Cut(V30V, ":")
		V30V = strings.TrimSpace(V30V)
		go F5F(V30V)
		if S2S.Q11Q == "汉语" {
			V16V["提取文件并复制完成语"] = `已提取文件的语法内容并复制， <span style="color: rgb(194, 173, 214);">Ctrl+V 粘贴</span>。文件:`
		} else {
			V16V["提取文件并复制完成语"] = `The syntax content of the file has been extracted and copied. <span style="color: rgb(194, 173, 214);">Ctrl+V paste</span>. document:`

		}
		S4S.T2_str = V16V["提取文件并复制完成语"] + filepath.Base(V30V)

	} else if strings.HasPrefix(V30V, "编辑文件内容:") {
		_, V30V, _ = strings.Cut(V30V, ":")
		V30V = strings.TrimSpace(V30V)

		go F8F(V30V)
		if S2S.Q11Q == "汉语" {
			V16V["打开语法文件语"] = `<span style="color: rgb(194, 173, 214);">已经打开语法文件：</span>`
		} else {
			V16V["打开语法文件语"] = `<span style="color: rgb(194, 173, 214);"> The grammar file has been opened.：</span>`

		}
		S4S.T2_str = V16V["打开语法文件语"] + filepath.Base(V30V)

	} else if strings.HasPrefix(V30V, "添加文件后缀名:") {
		_, V30V, _ = strings.Cut(V30V, ":")
		V30V = strings.TrimSpace(V30V)

		F33F(V30V, V16V["添加后缀名说明"])
		if S2S.Q11Q == "汉语" {
			S4S.T2_str = `添加文件后缀名,<span style="color: rgb(194, 173, 214);">完成.</span>`

		} else {
			S4S.T2_str = `Add file extension, <span style="color: rgb(194, 173, 214);">Done.</span>`

		}
		S4S.I1_str = V16V["添加后缀名说明"] + "\n" + strings.Join(S7S.W1_strs, "")

	} else if strings.HasPrefix(V30V, "重新载入:") {
		S2S.Q4Q = make(map[string]string)

		F3F(S2S.Q8Q, 0)
		if S2S.Q11Q == "汉语" {
			S4S.T2_str = `重新载入,语法目录搜索子文件名,<span style="color: rgb(194, 173, 214);">完成.</span>`
		} else {
			S4S.T2_str = `Reload, search sub-file names in syntax directory, <span style="color: rgb(194, 173, 214);">Done.</span>`
		}

	} else if strings.HasPrefix(V30V, "打开语法目录:") {
		_, V30V, _ = strings.Cut(V30V, ":")
		V30V = strings.TrimSpace(V30V)
		go F8F(V30V)

		var V29V = filepath.Base(V30V)

		if S2S.Q11Q == "汉语" {
			V16V["打开语法目录语"] = `打开语法目录,<span style="color: rgb(194, 173, 214);">|替换目录名|</span>  完成.`
		} else {
			V16V["打开语法目录语"] = `Open grammar directory,<span style="color: rgb(194, 173, 214);">|替换目录名|</span>  Done.`

		}
		S4S.T2_str = strings.Replace(V16V["打开语法目录语"], "|替换目录名|", V29V, -1)

	} else if strings.HasPrefix(V30V, "PC-Gui框架说明:") {
		_, V30V, _ = strings.Cut(V30V, ":")
		S4S.L5_str = F35F()

	} else if strings.HasPrefix(V30V, "退出:") {
		if S2S.Q11Q == "汉语" {
			V16V["已经退出"] = `已经退出.`
		} else {
			V16V["已经退出"] = `Has quit.`

		}

		S4S.T2_str = V16V["已经退出"]
		S4S.L3_str = V16V["已经退出"]
		S4S.L2_str = V16V["已经退出"]
		S4S.L1_str = V16V["已经退出"]

		go func() {
			time.Sleep(1 * time.Second)
			log.Fatal("退出程序  ：之前有 选择弹出窗口，须强制退出 ")
			os.Exit(0)

		}()

	} else {

	}

	return

}

func F2F() (V34V []string) {

	var V33V = []string{
		".txt", ".html", ".css",
		".py",
		".java",
		".js",
		".cs",
		".go",
		".c",
		".cpp",
		".php",
		".ts",
		".swift",
		".m",
		".kt",
		".rs",
		".sql",
		".rb",
		".r",
		".m",
		".pl",
		".scala",
		".dart",
		".lua",
		".groovy",
		".hs",
		".ps1",
		".sh",
		".asm",
		".bas",
		".pas",
		".vb",
		".jl",
	}
	V34V = V33V
	return

}

func F3F(V35V string, V36V int) {

	var V44V, V47V, V43V string
	var V46V string
	V41V, _ := ioutil.ReadDir(V35V)
	if len(V41V) == 0 {
		return
	}
	var V42V os.FileInfo
	for _, V42V = range V41V {
		V43V = V42V.Name()
		if V42V.IsDir() {
			if V43V == "app" || V43V == "go-file" || V43V == "beiFen" || V43V == "vendor" {
				continue
			}

			V44V = filepath.Join(V35V, V43V)

			if V36V >= 2 {
				continue
			} else if V36V == 0 {
				go F3F(V44V, V36V+1)

			} else {
				F3F(V44V, V36V+1)
			}

		} else {

			if F4F(V43V) == false {
				continue
			}

			V46V = filepath.Base(V35V)

			V46V = fmt.Sprintf("%s/%s", V46V, V43V)

			V47V = filepath.Join(V35V, V43V)

			S2S.Q14Q.Lock()
			S2S.Q4Q[V46V] = V47V
			S2S.Q14Q.Unlock()

		}

	}

}

func F4F(V48V string) (V51V bool) {

	var V50V string
	for _, V50V = range S7S.W1_strs {

		if strings.HasSuffix(V48V, V50V) {
			V51V = true
			return
		}
	}

	return

}

func F5F(V52V string) {

	var V54V = F16F(V52V)

	clipboard.WriteAll(V54V)

	F14F()

	return

}

func F6F(V55V string) {

	var V58V = `<html> <head><meta http-equiv="Content-Type" content="text/html; charset=utf8" /><title></title><meta http-equiv="refresh" content="0;URL=替换网址"></head><body></body></html>`

	V58V = strings.Replace(V58V, "替换网址", V55V, -1)
	F18F(S2S.Q1Q, V58V, "保存")

	return

}

func F7F(V59V string) {
	time.Sleep(300 * time.Millisecond)

	V60V := exec.Command("rundll32", "url.dll,FileProtocolHandler", V59V).Run()
	if V60V == nil {
		return
	}

	V61V := exec.Command("cmd", "/c", "start", V59V)

	V61V.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	var V63V = V61V.Start()
	if V63V != nil {

		_ = V63V
	}
	return
}

func F8F(V68V string) {
	V68V = strings.Replace(V68V, `/`, `\`, -1)
	V68V = strings.Replace(V68V, `\\`, ``, -1)

	V68V = strings.TrimSpace(V68V)
	V68V = strings.Trim(V68V, `\`)

	V69V := exec.Command("cmd", "/c", "start", "", V68V)

	V71V := V69V.Run()

	if V71V == nil {

		return
	}

	V71V = exec.Command(`cmd`, `/c`, `explorer`, V68V).Start()

	if V71V != nil {

		_ = V71V
	}

	return

}

func F9F(V72V string) string {
	V73V, V74V := ioutil.ReadFile(V72V)
	if V74V != nil {

		log.Fatal("读取本地主页网页文件并返回源码 ：错误", V74V)
		return ""
	}

	return string(V73V)
}

func F10F(V75V string, V76V interface{}) error {

	if V77V := os.MkdirAll(filepath.Dir(V75V), 0755); V77V != nil {
		return fmt.Errorf("创建目录失败: %w", V77V)
	}

	V78V, V79V := os.Create(V75V)
	if V79V != nil {
		return fmt.Errorf("文件创建失败: %w", V79V)
	}
	defer V78V.Close()

	V80V := json.NewEncoder(V78V)
	V80V.SetIndent("", "  ")

	if V81V := V80V.Encode(V76V); V81V != nil {
		return fmt.Errorf("JSON编码失败: %w", V81V)
	}

	if V82V := V78V.Sync(); V82V != nil {
		return fmt.Errorf("数据同步失败: %w", V82V)
	}

	return nil
}

func F11F(V83V string, V84V interface{}) error {
	V85V, V86V := os.Open(V83V)
	if V86V != nil {
		return V86V
	}
	defer V85V.Close()

	return json.NewDecoder(V85V).Decode(V84V)
}

func F12F(V87V map[string]int) (V90V map[string]int) {

	V90V = make(map[string]int)

	if len(V87V) < 51 {
		V90V = V87V
		return
	}

	var V93V = F32F(V87V)

	if len(V93V) >= 50 {
		V93V = V93V[:50]
	}
	for _, 词语与词频数一结构体 := range V93V {

		V90V[词语与词频数一结构体.Q13Q] = 词语与词频数一结构体.Q12Q
	}
	return

}

func F13F(V96V string) (V103V bool) {

	V96V = strings.TrimSpace(V96V)

	var V98V string
	for _, 每个字符 := range V96V {
		V98V = string(每个字符)

		if unicode.Is(unicode.Scripts["Han"], 每个字符) {
			V103V = true
			return

		} else if V98V == "'" {
			V103V = true
			return
		} else if V98V == "=" || V98V == "~" || V98V == "$" || V98V == "^" || V98V == "+" || V98V == "|" || V98V == "<" || V98V == ">" {
			V103V = true
			return
		} else if unicode.IsPunct(每个字符) {
			V103V = true
			return
		} else if unicode.IsSymbol(每个字符) {
			V103V = true
			return
		} else {

			continue

		}
	}

	return
}

func F14F() {

	V104V := syscall.MustLoadDLL("user32.dll")
	defer V104V.Release()

	V105V := V104V.MustFindProc("GetForegroundWindow")

	V106V := V104V.MustFindProc("ShowWindow")

	V107V, _, _ := V105V.Call()

	if V107V == 0 {

		return
	}

	V108V, _, V109V := V106V.Call(V107V, 6)

	if V108V == 0 {

		_ = V109V

	} else {

	}
}

func F15F() {

	var V110V runtime.MemStats

	runtime.ReadMemStats(&V110V)

	runtime.GC()

	return

}

func F16F(V111V string) (V115V string) {
	var _, V113V = os.Stat(V111V)
	if V113V != nil {

		return
	}

	var V136V []string

	defer func() {
		V115V = strings.Join(V136V, "")
	}()

	var V117V = mahonia.NewDecoder("GBK")

	var V119V = mahonia.NewEncoder("UTF8")

	var V122V, V130V = os.Open(V111V)
	if V130V != nil {

		return
	}
	defer V122V.Close()

	V123V := bufio.NewReader(V122V)

	var V135V, V132V string

	var V126V int
	var V133V = "检测"
	for {
		V135V, V130V = V123V.ReadString('\n')

		V126V++

		if V126V < 100 && V133V == "检测" {
			if F17F([]byte(V135V)) {
				V132V = "UTF8"
			} else {
				V132V = "GBK"
				V133V = "确认为WIN系统文件 不再检测"

			}

		}

		if V132V == "UTF8" {
			V135V = V119V.ConvertString(V135V)
		} else {

			V135V = V117V.ConvertString(V135V)
		}

		V136V = append(V136V, V135V)
		if V130V != nil {
			return
		}
	}

	return
}

func F17F(V137V []byte) bool {

	V138V := 0
	for V139V := 0; V139V < len(V137V); V139V++ {
		if V138V == 0 {
			if (V137V[V139V] & 0x80) != 0 {
				for (V137V[V139V] & 0x80) != 0 {
					V137V[V139V] <<= 1
					V138V++
				}

				if V138V < 2 || V138V > 6 {
					return false
				}

				V138V--
			}
		} else {
			if V137V[V139V]&0xc0 != 0x80 {
				return false
			}
			V138V--
		}
	}
	return V138V == 0
}

func F18F(V140V, V141V, V142V string) {
	var V144V = filepath.Dir(V140V)
	var _, V150V = os.Stat(V144V)
	if V150V != nil {

		_ = os.MkdirAll(V144V, os.ModePerm)

	}

	var V149V *os.File
	if V142V == "追加" {
		V149V, V150V = os.OpenFile(V140V, os.O_APPEND|os.O_CREATE, 0666)
	} else {
		V149V, V150V = os.OpenFile(V140V, os.O_TRUNC|os.O_CREATE, 0666)

	}

	if V150V != nil {

		return
	}
	defer V149V.Close()
	var V152V = bufio.NewWriter(V149V)
	V152V.WriteString(V141V)
	V152V.Flush()

}

type V153V struct {
	Message string `json:"message"`
}

func F19F() {

	var V155V = 32658 + F26F(S2S.Q8Q)
	for {
		if F20F(V155V) {

			break
		} else {

			V155V++
		}

	}

	var V157V = fmt.Sprintf(":%d", V155V)

	http.HandleFunc("/", F21F)

	http.HandleFunc("/chat", F22F)

	http.HandleFunc("/bc", F23F)

	S2S.Q9Q = fmt.Sprintf("http://localhost%s", V157V)

	go func() {

		F6F(S2S.Q9Q)
		F7F(S2S.Q9Q)

		time.Sleep(200 * time.Millisecond)

		F29F()

	}()

	if V158V := http.ListenAndServe(V157V, nil); V158V != nil {

		V155V++

		log.Fatal(V158V)
	}

}

func F20F(V159V int) bool {

	V160V, V161V := net.Listen("tcp", fmt.Sprintf(":%d", V159V))
	if V161V != nil {

		return false
	}

	_ = V160V.Close()
	return true
}

func F21F(V162V http.ResponseWriter, V163V *http.Request) {
	var V176V string

	V176V = F27F()

	V176V = strings.Replace(V176V, "|替换换行|", "\n", -1)

	V176V = F24F(V176V)

	V176V = strings.Replace(V176V, "|替换页面标题名|", S2S.Q10Q, -1)
	V176V = strings.Replace(V176V, "|替换CSS样式|", F28F(), -1)
	var V171V = make(map[string]string)

	if S2S.Q11Q == "汉语" {

		V171V["IE浏览器过时"] = "您的IE浏览器过时，推荐使用 Chrome/Firefox/Edge 等现代浏览器."
	} else {

		V171V["IE浏览器过时"] = "Your IE browser is outdated. We recommend using modern browsers like Chrome, Firefox, or Edge."

	}
	V176V = strings.Replace(V176V, "|替换IE浏览器过时|", V171V["IE浏览器过时"], -1)
	V176V = strings.Replace(V176V, "<!--替换各类AI模型接入-->", V171V["AI模型接入按钮"], -1)

	V176V = strings.Replace(V176V, "|替换程序名|", S2S.Q10Q, -1)
	V176V = strings.Replace(V176V, "|替换语法路径目录| ", S2S.Q8Q, -1)

	V176V = strings.Replace(V176V, "<!--替换各国语言-->", F25F(), -1)

	V162V.Header().Set("Content-Type", "text/html; charset=utf-8")

	V162V.Write([]byte(V176V))
	return
}

func F22F(V177V http.ResponseWriter, V178V *http.Request) {
	defer func() {

	}()

	if V178V.Method != http.MethodPost {
		http.Error(V177V, "只支持 POST 请求", http.StatusMethodNotAllowed)
		return
	}

	var V179V V153V
	V180V := json.NewDecoder(V178V.Body).Decode(&V179V)
	if V180V != nil {
		http.Error(V177V, "解析 JSON 数据失败", http.StatusBadRequest)
		return
	}
	defer V178V.Body.Close()

	*S4S = T2T{}
	F1F(V179V.Message)

	V177V.Header().Set("Content-Type", "application/json")

	json.NewEncoder(V177V).Encode(S4S)

	time.Sleep(200 * time.Millisecond)

	V178V.Body.Close()

	F15F()

}

func F23F(V181V http.ResponseWriter, V182V *http.Request) {
	V181V.Header().Set("Content-Type", "text/event-stream")
	V181V.Header().Set("Cache-Control", "no-cache")
	V181V.Header().Set("Connection", "keep-alive")
	V181V.Header().Set("Access-Control-Allow-Origin", "*")

	V183V, V184V := V181V.(http.Flusher)
	if !V184V {
		http.Error(V181V, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	var V186V = S2S.Q6Q
	for {
		for {
			if S2S.Q6Q != V186V {

				return
			}

			if S2S.Q3Q != "" {

				S2S.Q3Q = ""
				break
			}
			time.Sleep(100 * time.Millisecond)
		}

		V187V, _ := json.Marshal(S4S)

		fmt.Fprintf(V181V, "data: %s\n\n", V187V)
		V183V.Flush()

	}

}

func F24F(V208V string) string {
	var V190V = make(map[string]string)
	if S2S.Q11Q == "汉语" {

		V190V["网页语言"] = "zh-CN"
		V190V["切换亮色"] = "切换亮色"
		V190V["切换暗色"] = "切换暗色"

		V190V["切换主题"] = "切换主题"
		V190V["最近搜索记录"] = "最近搜索记录"
		V190V["搜索语法文件名或词语"] = "搜索语法文件名或词语"
		V190V["搜索语法词语"] = "搜索语法词语"
		V190V["语法目录"] = "语法目录"

		V190V["点击复制文件内容"] = "点击复制文件内容"
		V190V["编辑"] = "编辑"

		V190V["添加文件后缀名"] = "添加文件后缀名"
		V190V["新增加子文件说明"] = "新增加子文件,重新载入"
		V190V["增加子文件提示"] = "语法的目录新增加子文件,须重新载入"

		V190V["打开语法目录"] = "打开语法目录"
		V190V["PC-Gui框架说明"] = "PC-Gui框架说明"

		V190V["退出程序"] = "退出程序"
		V190V["语法目录行名"] = "语法目录:"
		V190V["常用记录"] = "常用记录"

	} else {

		V190V["网页语言"] = "en-US"

		V190V["切换亮色"] = "Toggle Light Mode"
		V190V["切换暗色"] = "Toggle Dark Mode"

		V190V["切换主题"] = "Toggle Theme"
		V190V["最近搜索记录"] = "Recent"
		V190V["搜索语法文件名或词语"] = "Search Syntax Filename or Term"
		V190V["搜索语法词语"] = "Search Syntax Term"
		V190V["语法目录"] = "Syntax Directory"

		V190V["点击复制文件内容"] = "Click to Copy File Content"
		V190V["编辑"] = "Edit"

		V190V["添加文件后缀名"] = "Add File Extension"
		V190V["新增加子文件说明"] = "Add subfile"
		V190V["增加子文件提示"] = "When a subfile is newly added to the grammar's directory, it must be reloaded."

		V190V["打开语法目录"] = "Open Syntax Directory"
		V190V["PC-Gui框架说明"] = "PC-Gui description"

		V190V["退出程序"] = "Exit Program"
		V190V["语法目录行名"] = "Syntax Directory:"
		V190V["常用记录"] = "most"

	}

	V208V = strings.Replace(V208V, "|替换语言|", V190V["网页语言"], -1)
	V208V = strings.Replace(V208V, "|替换切换亮色|", V190V["切换亮色"], -1)
	V208V = strings.Replace(V208V, "|替换切换暗色|", V190V["切换暗色"], -1)

	V208V = strings.Replace(V208V, "|替换切换主题|", V190V["切换主题"], -1)
	V208V = strings.Replace(V208V, "|替换最近搜索记录|", V190V["最近搜索记录"], -1)
	V208V = strings.Replace(V208V, "|替换搜索语法文件名或词语|", V190V["搜索语法文件名或词语"], -1)
	V208V = strings.Replace(V208V, "|替换搜索语法词语|", V190V["搜索语法词语"], -1)

	V208V = strings.Replace(V208V, "|替换表头语法目录|", V190V["语法目录"], -1)
	V208V = strings.Replace(V208V, "|替换点击复制文件内容|", V190V["点击复制文件内容"], -1)
	V208V = strings.Replace(V208V, "|替换编辑|", V190V["编辑"], -1)

	V208V = strings.Replace(V208V, "|替换添加文件后缀名|", V190V["添加文件后缀名"], -1)
	V208V = strings.Replace(V208V, "|替换新增加子文件说明|", V190V["新增加子文件说明"], -1)

	V208V = strings.Replace(V208V, "|替换新增加子文件提示|", V190V["增加子文件提示"], -1)

	V208V = strings.Replace(V208V, "|替换打开语法目录|", V190V["打开语法目录"], -1)
	V208V = strings.Replace(V208V, "|替换PC-Gui框架说明|", V190V["PC-Gui框架说明"], -1)

	V208V = strings.Replace(V208V, "|替换退出|", V190V["退出程序"], -1)
	V208V = strings.Replace(V208V, "|替换行名的语法目录|", V190V["语法目录行名"], -1)
	V208V = strings.Replace(V208V, "|替换常用记录|", V190V["常用记录"], -1)

	return V208V
}

func F25F() (V210V string) {

	if S2S.Q11Q == "汉语" {

		return
	}
	V210V = ` 	   <div class="dropdown-container">
	<span class="dropdown-L0" > English - Please use the browser's built-in translation <i class="fas fa-chevron-down"></i></span>
	<div class="dropdown-menu" style="top:100%;min-height:550px;">
	<span  class="chat-item-C">English - Please use the browser's built-in translation</span>
	<span  class="chat-item-C" >简体.请用浏览器内置的翻译功能来翻译程序的网页界面。</span>
	<span  class="chat-item-C">Español - Utilice la traducción del navegador</span>
	<span  class="chat-item-C" >Français - Utilisez la traduction du navigateur</span>
	<span  class="chat-item-C" >Deutsch - Browserübersetzung verwenden</span>
	<span  class="chat-item-C" >日本語 - ブラウザの翻訳機能をご利用ください</span>
	<span  class="chat-item-C" >繁體.請用瀏覽器自帶的翻譯，翻譯程序網頁界面。</span>
	<span  class="chat-item-C">한국어 - 브라우저 번역 기능을 사용하세요</span>
	<span  class="chat-item-C" >Русский - Используйте встроенный перевод</span>
	<span  class="chat-item-C">العربية - استخدم ترجمة المتصفح المدمجة</span>
	<span  class="chat-item-C" >Português - Use a tradução do navegador</span>
	<span  class="chat-item-C" >Italiano - Usa la traduzione del browser</span>
	<span  class="chat-item-C">Nederlands - Gebruik browsertranslatie</span>
	<span  class="chat-item-C">Svenska - Använd webbläsaröversättning</span>
	<span  class="chat-item-C">Türkçe - Tarayıcı çevirisini kullanın</span>
	<span  class="chat-item-C" >हिन्दी - ब्राउज़र अनुवाद का उपयोग करें</span>
	<span  class="chat-item-C">ไทย - ใช้การแปลของเบราว์เซอร์</span>
	<span  class="chat-item-C" >Tiếng Việt - Dùng tính năng dịch trình duyệt</span>
	<span  class="chat-item-C">Polski - Użyj tłumaczenia przeglądarki</span>
	<span  class="chat-item-C">Українська - Використовуйте переклад браузера</span>
	<span  class="chat-item-C">Ελληνικά - Χρησιμοποιήστε μετάφραση προγράμματος περιήγησης</span>
	<span  class="chat-item-C">עברית - השתמש בתרגום דפדפן מובנה</span>
	<span  class="chat-item-C">فارسی - از ترجمه مرورگر استفاده کنید</span>
	<span  class="chat-item-C">اردو - براؤزر ترجمہ کا استعمال کریں</span>
	<span  class="chat-item-C">Čeština - Použijte překlad prohlížeče</span>
	<span  class="chat-item-C" >Magyar - Használja a böngészőfordítót</span>
	<span  class="chat-item-C">Suomi - Käytä selaimen käännöstä</span>
	<span  class="chat-item-C">Norsk - Bruk nettleseroversettelse</span>
	<span  class="chat-item-C">Dansk - Brug browseroversættelse</span>
	<span  class="chat-item-C" >Română - Utilizați traducerea browserului</span>
	<span  class="chat-item-C">Bahasa Indonesia - Gunakan terjemahan browser</span>
	<span  class="chat-item-C">Bahasa Melayu - Gunakan terjemahan pelayar</span>
	<span  class="chat-item-C" >Filipino - Gamitin ang browser translation</span>
	<span  class="chat-item-C">বাংলা - ব্রাউজার অনুবাদ ব্যবহার করুন</span>
	<span  class="chat-item-C">தமிழ் - உலாவி மொழிபெயர்ப்பைப் பயன்படுத்தவும்</span>
	<span  class="chat-item-C">తెలుగు - బ్రౌజర్ అనువాదాన్ని ఉపయోగించండి</span>
	<span  class="chat-item-C" >Kiswahili - Tumia tafsiri ya kivinjari</span>
	<span  class="chat-item-C" >नेपाली - ब्राउजर अनुवाद प्रयोग गर्नुहोस्</span>
	<span  class="chat-item-C">සිංහල - බ්රවුසර් පරිවර්තනය භාවිතා කරන්න</span>
	<span  class="chat-item-C" >ខ្មែរ - ប្រើការបកប្រែរបស់កម្មវិធីរុករក</span>
	<span  class="chat-item-C">မြန်မာ - ဘရောက်ဆာ၏ဘာသာပြန်ချက်ကိုသုံးပါ</span>
	<span  class="chat-item-C" >አማርኛ - የማሰሻገሪያ አሞሌ ይጠቀሙ</span>
	<span  class="chat-item-C">Zulu - Sebenzisa ukuhumusha kwebrowser</span>
	<span  class="chat-item-C" >Xhosa - Sebenzisa uthelekiso lwebrowser</span>
	<span  class="chat-item-C">Chichewa - Gwiritsani ntchito kumasulira browser</span>
	<span  class="chat-item-C">Shona - Shandisa browser translation</span>
	<span  class="chat-item-C" >Sesotho - Sebedisa phetolelo ya browser</span>
	
	</div>
		</div>

	`

	return
}

func F26F(V211V string) (V221V int) {

	V213V := md5.Sum([]byte(V211V))
	V214V := hex.EncodeToString(V213V[:])

	V217V := make([]byte, 0, 4)
	for _, c := range V214V {
		if c >= '0' && c <= '9' {
			V217V = append(V217V, byte(c))
			if len(V217V) >= 4 {
				break
			}
		}
	}

	for len(V217V) < 4 {
		V217V = append(V217V, '0')
	}

	V218V := string(V217V)
	V219V, V220V := strconv.Atoi(V218V)
	if V220V != nil {
		fmt.Println("转换失败:", V220V)
		return
	}
	V221V = V219V

	return

}

func F27F() (V224V string) {

	V224V = `<!DOCTYPE html>
	<html lang="|替换语言|">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="IE=edge"> 
		<!--[if IE]>
		<script>
		alert('|替换IE浏览器过时|');
		</script>
		<![endif]-->
		<title>|替换页面标题名|</title>
		<link rel="shortcut icon" href="./app/10.ico" type="image/x-icon">

		|替换CSS样式|  
		
	</head>
	<body>
    <button class="theme-toggle" onclick="toggleTheme()">🌙 |替换切换主题|</button>
    <div class="container">
        <div class="sidebar">
            <h3>🕒 |替换最近搜索记录|</h3>
            <div   id="L2_str">
                </div>
        </div>
		<div class="main-content">
			<div class="title-banner">📁 |替换页面标题名|</div>		
			<!--替换各国语言-->

            <div class="search-area">
                <input type="text"  class="input-field KeyOn"  placeholder="🔍 |替换搜索语法文件名或词语|...">
                <button  class="Dataset-Text" id="In1_str" data-text="搜索词语:" style="min-width: 200px;">🔍 |替换搜索语法词语|</button>
            </div>
            <div id="T2_str"></div>
            <div class="table-container">
                <table class="list-table">
                    <thead>
                        <tr>
                            <th>|替换表头语法目录|</th>
                            <th>|替换点击复制文件内容|</th>
                            <th>|替换编辑|</th>
                        </tr>
                    </thead>
                    <tbody  id="L1_str">
                    </tbody>
                </table>
            </div>
            <div class="bottom-area">
                <textarea id="I1_str" rows="2"  class="input-field" placeholder=""></textarea>
                <button class="Dataset-Text"  data-text="添加文件后缀名:" >📎 |替换添加文件后缀名|</button>
            </div>
            <div class="bottom-area">
             <button class="Dataset-Text"  title="&nbsp;|替换新增加子文件提示|"  data-text="重新载入:">🔄|替换新增加子文件说明|</button>
			<button class="Dataset-Text"  data-text="打开语法目录:"> 📁|替换打开语法目录|</button>
			<button class="Dataset-Text"  data-text="PC-Gui框架说明:">💻|替换PC-Gui框架说明|</button>
    
        <button class="Dataset-Text"  data-text="退出:">🚫 |替换退出|</button>
		</div>
     
        <div class="bottom-area">
            <span style="color: rgb(194, 173, 214);">|替换行名的语法目录|</span>|替换语法路径目录|  
		</div>
		
		<div  id="L5_str" >  </div>

        </div>
        <div class="sidebar">
            <h3  >⭐ |替换常用记录|</h3>
            <div  id="L3_str">
        </div>
        </div>
	</div>
	

		
		<script>
	
	
		const L1_str = document.getElementById('L1_str');
		const L2_str = document.getElementById('L2_str');
		const L3_str = document.getElementById('L3_str');
		const L5_str = document.getElementById('L5_str');

		
		const I1_str = document.getElementById('I1_str');
		const T2_str = document.getElementById('T2_str');
		 const KeyOn = document.querySelector('.KeyOn');
	 const In1_str = document.getElementById('In1_str');
	  window.onload = function() {   
		 console.log("初始 "); 
		PostB("初始:").then((reTxt) => {
		});
		initTheme()
	 };
		function initTheme() {
		console.log("初始化主题设置");
		const body = document.body;
		const button = document.querySelector('.theme-toggle');
		const savedTheme = localStorage.getItem('theme');
		const systemDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
		if (savedTheme === 'dark' || (!savedTheme && systemDark)) {
			body.classList.add('dark-mode');
			button.textContent = '☀️ |替换切换亮色|';
		} else {
			body.classList.remove('dark-mode');
			button.textContent = '🌙 |替换切换暗色|';
		}
		}
		function toggleTheme() {
		console.log("执行主题切换");
		const body = document.body;
		const button = document.querySelector('.theme-toggle');
		if (body.classList.contains('dark-mode')) {
			body.classList.remove('dark-mode');
			button.textContent = '🌙 |替换切换暗色|';
			localStorage.setItem('theme', 'light');
		} else {
			body.classList.add('dark-mode');
			button.textContent = '☀️ |替换切换亮色|';
			localStorage.setItem('theme', 'dark');
		}
		}
		window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', e => {
		if (!localStorage.getItem('theme')) { 
			initTheme();
			console.log("系统主题变更，自动同步");
		}
		});
			document.querySelectorAll('.sidebar-item').forEach(item => {
				item.addEventListener('click', function() {
					console.log('点击了：' + this.textContent);
				});
			});
		KeyOn.addEventListener('keypress', function(e) {
			if(e.key === 'Enter' || e.keyCode === 13) {
				In1_str.click();
			}
		});
		document.body.addEventListener('click', function(event) {
		console.log('统一响应 点击:');
		const target = event.target.closest('.Dataset-Text');
		if (!target) return;
		let  dataText =''
		if (target) {
			dataText = target.dataset.text
		}
		if (target.matches('button')) { 
			const inputElement = target.previousElementSibling;
			if (inputElement) {
				if (inputElement.tagName === 'INPUT' || inputElement.tagName === 'TEXTAREA') {
					const inputText = inputElement.value;
					dataText += inputText;  
				}
			}
		} 
			if (dataText.length < 2) {
				return
			}  
			console.log(dataText);
				PostB(dataText).then((data) => {
			});
		});
		function UPpage(data) {
		console.log("data数据"); 
		console.log("data"); 
		if (data.L1_str !== "") {
			L1_str.innerHTML = data.L1_str; 
		}
		if (data.L2_str !== "") {
			L2_str.innerHTML = data.L2_str; 
		}
		if (data.L3_str !== "") {
			L3_str.innerHTML = data.L3_str; 
		}

			L5_str.innerHTML = data.L5_str; 
	

			if (data.I1_str !== "") {
				I1_str.textContent = data.I1_str; 
		}
		if (data.T2_str !== "") {
				T2_str.innerHTML = data.T2_str; 
				T2_str.style.display = 'flex';
			} else { 
				T2_str.style.display = 'none';
			}
		}
		async function PostB(reTxt) {
		const message = reTxt.trim();
		if (!message) {
			console.warn('消息内容为空');
			return null;
		}
		console.log(reTxt);
		try {
		const response = await fetch('/chat', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({message})
		});
		if (!response.ok) {
			throw new Error(|替换单点号|HTTP error! status: ${response.status}|替换单点号|);
		}
		const data = await response.json();
			console.log('服务器响应:');
			console.log(data);
			UPpage(data) 
			return null;
		} catch (error) {
		return null;
		}
		}
		function connect() {
		eventSource = new EventSource('/bc');
		eventSource.onopen = function() {
		};
		eventSource.onmessage = function(event) {
			try {
				UPpage(JSON.parse(event.data)); 
			} catch (e) {
			}
		};
		}
		function showTemporaryAlert(reTxt) {
		M3.innerHTML = reTxt;
		M3.classList.add('active');
		setTimeout(() => {
			M3.innerHTML =''
			M3.classList.remove('active');              
		}, 3500);
		}
		
		</script>
    
	</body>
	</html>
	`
	V224V = strings.Replace(V224V, "|替换单点号|", "`", -1)
	return

}

func F28F() (V226V string) {
	V226V = `		 				  <style>
	* {
		margin: 0;
		padding: 0;
		box-sizing: border-box;
		font-family: "微软雅黑", "苹方", sans-serif;
		
	}
	:root {
		--bg-primary: #f5f6f0;      
		--bg-secondary: #e8ede2;    
		--bg-card: #fafbf8;         
		--text-primary: #4a5568;    
		--text-secondary: #718096;  
		--border-color: #d4ddd0;    
		--accent-color: #8fbc8f;    
		--hover-color: #a8d5a8;     
		--shadow: 0 2px 8px rgba(0,0,0,0.06);
	}
	body.dark-mode {
		--bg-primary: #1e2a1e;      
		--bg-secondary: #2d3a2d;    
		--bg-card: #253325;         
		--text-primary: #d4ddd4;    
		--text-secondary: #e2ece2;  
		--border-color: #3d4a3d;    
		--accent-color: #6b9b6b;    
		--hover-color: #7fb07f;     
		--shadow: 0 2px 8px rgba(0,0,0,0.3);
	}
	body {
		background-color: var(--bg-primary);
		color: var(--text-primary);
		padding: 20px;
		line-height: 20px;
		transition: all 0.3s ease;
	}
	.theme-toggle {
		position: fixed;
		top: 20px;
		right: 20px;
		padding: 10px 20px;
		background: var(--accent-color);
		color: white;
		border: none;
		border-radius: 20px;
		cursor: pointer;
		font-size: 14px;
		box-shadow: var(--shadow);
		transition: all 0.3s ease;
		z-index: 1000;
	}
	.theme-toggle:hover {
		background: var(--hover-color);
		transform: translateY(-2px);
	}
	.container {
		display: flex;
		gap: 20px;
		max-width: 1400px;
		margin: 0 auto;
		padding-top: 20px;
	}
	.sidebar {
		width: 220px;
		background-color: var(--bg-card);
		border-radius: 12px;
		padding: 20px;
		min-height: 600px;
		box-shadow: var(--shadow);
		transition: all 0.3s ease;
		line-height: 8px;
	}
	.sidebar h3 {
		color: var(--text-primary);
		font-size: 16px;
		margin-bottom: 15px;
		padding-bottom: 10px;
		border-bottom: 2px solid var(--accent-color);
	}
	.sidebar-item {
		padding: 5px;
		margin: 2px 0;
		background: var(--bg-secondary);
		border-radius: 1px;
		cursor: pointer;
		transition: all 0.2s ease;
		font-size: 14px;
		color: var(--text-secondary);
	}
	.sidebar-item:hover {
		background: var(--accent-color);
		color: white;
		transform: translateX(5px);
	}
	.main-content {
		flex: 1;
		display: flex;
		flex-direction: column;
		gap: 20px;
	}
	.title-banner {
		background: linear-gradient(135deg, var(--accent-color), var(--hover-color));
		border-radius: 12px;
		padding: 25px;
		text-align: center;
		font-size: 20px;
		font-weight: 500;
		color: white;
		box-shadow: var(--shadow);
	}
	.search-area {
		display: flex;
		gap: 15px;
		align-items: center;
		flex-wrap: wrap;
	}
	input, button, textarea {
		padding: 12px 18px;
		border: 1px solid var(--border-color);
		border-radius: 8px;
		background-color: var(--bg-card);
		color: var(--text-primary);
		font-size: 14px;
		transition: all 0.3s ease;
	}
	input:focus, textarea:focus {
		outline: none;
		border-color: var(--accent-color);
		box-shadow: 0 0 0 3px rgba(143, 188, 143, 0.1);
	}
	.search-area input {
		flex: 1;
		min-width: 150px;
		width: 70%;
	}
	button {
		cursor: pointer;
		background-color: var(--accent-color);
		color: white;
		border: none;
		font-weight: 500;
		transition: all 0.3s ease;
	}
	button:hover {
		background-color: var(--hover-color);
		transform: translateY(-2px);
		box-shadow: var(--shadow);
	}
	button:active {
		transform: translateY(0);
	}
	.table-container {
		background-color: var(--bg-card);
		border-radius: 12px;
		overflow: hidden;
		box-shadow: var(--shadow);
	}
	.list-table {
		width: 100%;
		border-collapse: collapse;
	}
	.list-table th {
		padding: 15px 20px;
		border-bottom: 1px solid var(--border-color);
		text-align: left;
		background-color: var(--bg-secondary);
		font-weight: 500;
		color: var(--text-primary);
		font-size: 18px;
	}
	.list-table td {
		color: var(--text-secondary);
		font-size: 13px;
		padding: -5px;
		line-height: 8px;
		border-bottom: 1px solid var(--border-color);
		text-align: left;
	}
	.list-table tr:hover {
		background-color: var(--bg-secondary);
		transition: background-color 0.2s ease;
	}
	.list-table tr:last-child td {
		border-bottom: none;
	}
	.bottom-area {
		display: flex;
		gap: 15px;
		align-items: flex-start;
	}
	textarea {
		flex: 1;
		min-height: 50px;
		resize: vertical;
		line-height: 1.6;
	}
	.bottom-area button {
		white-space: nowrap;
		height: fit-content;
	}
	@media (max-width: 1024px) {
		.container {
			flex-direction: column;
		}
		.sidebar {
			width: 100%;
			min-height: auto;
		}
	}
	@media (max-width: 768px) {
		.search-area {
			flex-direction: column;
		}
		.search-area input {
			width: 70%;
		}
		.bottom-area {
			flex-direction: column;
		}
		.bottom-area button {
			width: 100%;
		}
		.theme-toggle {
			top: 10px;
			right: 10px;
			padding: 8px 15px;
			font-size: 12px;
		}
	}
	::-webkit-scrollbar {
		width: 8px;
		height: 8px;
	}
	::-webkit-scrollbar-track {
		background: var(--bg-secondary);
	}
	::-webkit-scrollbar-thumb {
		background: var(--accent-color);
		border-radius: 4px;
	}
	::-webkit-scrollbar-thumb:hover {
		background: var(--hover-color);
	}

	.dropdown-container {
		position: relative;
	}
	.dropdown-container:hover .dropdown-menu{
		opacity: 1;
		visibility: visible;
		transform: translateY(0);
	}
	 .dropdown-container:hover .dropdown-menu2{
		opacity: 1;
		visibility: visible;
		transform: translateY(0);
	}
	.dropdown-menu {
		bottom: 100%;
		min-Height : 250px;  
		min-width: 510px;
		position : absolute;  
		 left: 0;
		border-radius: 12px;
		opacity: 0;
		visibility: hidden;
		transform: translateY(10px);
		transition: all 0.3s ease;
		box-shadow: 0 10px 30px rgba(0,0,0,0.15);
		z-index: 10;
		max-height: 550px;
		overflow-y: auto;
		 background: rgba(219, 231, 217, 0.8);
	border: 1px solid rgba(160, 200, 155, 0.8);
	backdrop-filter: blur(5px);
	-webkit-backdrop-filter: blur(5px);
	}
	 .dropdown-L0 {
		width: 100%;
		padding: 5px 8px;
		border: 1px solid rgba(141, 136, 136, 0.2);
		border-radius: 8px;
		color: rgb(53, 47, 47);
		font-size: 15px;
		font-weight: 300;
		cursor: pointer;
		display: flex;
		justify-content: space-between;
		align-items: center;
		transition: all 0.3s ease;
		backdrop-filter: blur(5px);
		background: linear-gradient(135deg, #b1ccb1 0%, #a4dfa4 100%);
	color: #2d4d2a; 
	border-bottom: 1px solid rgba(160, 200, 155, 0.8);
	}
	.dropdown-L0:hover {
		background: rgba(255, 255, 255, 0.25);
		transform: translateY(-2px);
		box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
	}
	.dropdown-L0 i {
		transition: transform 0.3s ease;
	}
	.dropdown-container:hover .dropdown-L0 i {
		transform: rotate(180deg);
	}
	.chat-item-C {
	width: 150% !important; 
	box-sizing: border-box; 
	padding: 1px 3px;
	font-size: 13px;
	margin: 5px 0;
	border-radius: 8px;
	min-width: 250px; 
	cursor: pointer;
	transition: all 0.3s ease;
	border: 1px solid rgba(150, 180, 145, 0.1);
	color: #333;
	display: block; 
	max-width: 100%; 
	}
	.chat-item-C:hover {
	background: rgba(5, 88, 243, 0.87);
	color: rgb(235, 227, 227);
	transform: translateX(3px);
	}


	</style>
		`

	return

}

func F29F() {

	S4S.T2_str = fmt.Sprintf("查询语法的文件数:%d", len(S2S.Q4Q))
	S2S.Q3Q = "前端"
	return

}

func F30F(V227V string) (V256V string, V229V int) {

	var V230V, V231V, V254V, V252V, V247V, V250V, V251V string
	var V238V = "<tr><td>📁%s</td><td>📄%s</td><td>%s</td></tr>"
	var V240V = `<span class="sidebar-item Dataset-Text" data-text="打开语法目录:%s" >%s</span>`
	var V242V = `<span class="sidebar-item Dataset-Text" data-text="复制文件内容:%s" >%s</span>`
	var V244V = `<span class="sidebar-item Dataset-Text" data-text="编辑文件内容:%s" >🖉</span>`

	_ = V231V
	var V253V, V249V string
	S2S.Q7Q = S2S.Q7Q[0:0]
	for V230V, V231V = range S2S.Q4Q {

		V247V, V254V, _ = strings.Cut(V230V, "/")

		if V254V == V227V || strings.Contains(V254V, V227V) {

			V229V++

			V249V = fmt.Sprintf(V240V, filepath.Dir(V231V), V247V)

			V250V = fmt.Sprintf(V242V, V231V, V254V)
			V251V = fmt.Sprintf(V244V, V231V)

			V252V = fmt.Sprintf(V238V, V249V, V250V, V251V)

		} else {
			continue
		}

		V253V = filepath.Ext(V254V)

		V254V = strings.Replace(V254V, V253V, "", -1)

		if V254V == V227V {

			S2S.Q7Q = append([]string{V252V}, S2S.Q7Q...)
		} else if strings.Contains(V254V, V227V) {
			S2S.Q7Q = append(S2S.Q7Q, V252V)
		}

	}

	if len(S2S.Q7Q) == 0 {
		V256V = "搜索结果为空."
		return
	}
	S2S.Q7Q = append(S2S.Q7Q, "<tr><td></td><td></td><td></td></tr>")

	V256V = strings.Join(S2S.Q7Q, "<br />")

	var V262V []string
	var V261V, V259V string
	for _, V259V = range S2S.Q5Q {

		_, V261V, _ = strings.Cut(V259V, "搜索词语:")
		V261V, _, _ = strings.Cut(V261V, `"`)
		if V261V == V227V {
			continue
		}
		V262V = append(V262V, V259V)
	}
	var V265V = `<h4  class="sidebar-item Dataset-Text" data-text="搜索词语:%s" >%s</h4>`
	V265V = fmt.Sprintf(V265V, V227V, V227V)
	S2S.Q5Q = append([]string{V265V}, V262V...)
	S4S.L2_str = strings.Join(S2S.Q5Q, "<br />")

	S7S.W2_map[V227V]++

	S4S.L3_str = F31F(S7S.W2_map)

	go func() {

		if V266V := F10F(S2S.Q2Q, &S7S); V266V != nil {

		}
	}()

	return

}

func F31F(V267V map[string]int) (V277V string) {

	if len(V267V) == 0 {
		return
	}

	var V270V = F32F(V267V)
	var V276V []string
	var V275V string

	var V274V = `<h4  class="sidebar-item Dataset-Text" data-text="搜索词语:%s" >%s</h4>`

	for _, 词语与词频数一结构体 := range V270V {

		V275V = fmt.Sprintf(V274V, 词语与词频数一结构体.Q13Q, 词语与词频数一结构体.Q13Q)
		V276V = append(V276V, V275V)

	}

	V277V = strings.Join(V276V, "<br />")

	return

}

func F32F(V278V map[string]int) (V281V []V6V) {

	V281V = make([]V6V, 0, len(V278V))

	for 词语一文本, 词频统计数一整数 := range V278V {

		V281V = append(V281V, V6V{词语一文本, 词频统计数一整数})
	}

	sort.Slice(V281V, func(i, j int) bool {

		if V281V[i].Q12Q != V281V[j].Q12Q {
			return V281V[i].Q12Q > V281V[j].Q12Q
		}

		return V281V[i].Q13Q < V281V[j].Q13Q
	})

	return V281V
}

func F33F(V287V, V284V string) {

	V284V = strings.Replace(V284V, ".", "", -1)

	V287V = strings.Replace(V287V, V284V, "", -1)

	V287V = strings.Replace(V287V, "\n", "", -1)
	V287V = strings.Replace(V287V, "\r", "", -1)

	var V289V string
	S7S.W1_strs = S7S.W1_strs[0:0]
	for _, V289V = range strings.Split(V287V, ".") {

		V289V = strings.TrimSpace(V289V)

		if len(V289V) < 1 {
			continue
		} else if F13F(V289V) == true {
			continue
		}
		S7S.W1_strs = append(S7S.W1_strs, fmt.Sprintf(".%s", V289V))

	}

	go func() {

		if V290V := F10F(S2S.Q2Q, &S7S); V290V != nil {

		}
	}()

	return

}

func F34F() (V293V string) {

	if S2S.Q11Q == "汉语" {
		V293V = `程序使用方法
		<br />0. '移动到语法的目录'：把'查询语法程序'移动到编程语法的文件夹里。
		<br />1. '启动程序时'：程序会自动在程序当前文件夹以及最多两层子文件夹中，查找所有语法文本的文件名（比如 '.txt' '.html'等）。
		<br />2. '搜索词语'：在输入框中输入你想查找的词，程序会列出所有包含这个词的文本文件名称。
		<br />3. '复制文件内容'：点击任意一个文件名，程序会自动把该文件里的全部内容复制到电脑的剪贴板中。然后你只需在编程编辑器（比如写代码的软件）里按下“粘贴”快捷键，就能把内容贴进去。
		<br />4. '自定义搜索范围'：程序下方还有一个输入框，你可以在这里增加要搜索的文件后缀名（例如 '.css'、'.py'），程序就会增加查找这些类型的文件。
		<br />5. '右侧的 ⭐ 常用记录栏'：这里会自动记录你搜过的词语。用得越多的词排得越靠前（从高到低排序），你每次打开就能一眼看到最常用的词，越用越顺手。`

	} else {

		V293V = `Program Usage Instructions  
			<br />0. 'Move to the syntax directory': Move the "Query Syntax Program" into the programming syntax folder.  
			<br />1. 'When starting the program': The program will automatically search the current folder and up to two subfolders for all syntax text files (e.g., '.txt', '.html', etc.).  
			<br />2. 'Search for a word': Enter the word you want to find in the input box, and the program will list all text file names that contain that word.  
			<br />3. 'Copy file content': Click on any file name, and the program will automatically copy the entire content of that file to your computer's clipboard. Then, simply press the "Paste" shortcut in your programming editor (e.g., coding software) to paste the content.  
			<br />4. 'Customize the search scope': There is another input box at the bottom of the program where you can add additional file extensions to search (e.g., '.css', '.py'), and the program will then search for those types of files as well.  
			<br />5. '⭐ most': This will automatically record the words you have searched for. The more frequently you use a word, the higher it appears (sorted from highest to lowest). Every time you open the program, you can see your most commonly used words at a glance, making it more convenient the more you use it.`

	}

	return

}

func F35F() (V299V string) {

	defer func() {
		var V296V string

		V296V = `<a href="https://github.com/jiqi136/PC-Gui" target="_blank" style="color: #7d98c7;">https://github.com/jiqi136/PC-Gui</a>`

		V299V = strings.Replace(V299V, "|替换源码下载网址|", V296V, -1)

	}()

	if S2S.Q11Q == "汉语" {

		V299V = `			<hr>	本程序采用桌面 GUI 框架
		<div style="max-width: 110%;">
		<h2 style="font-weight: 700; background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text; margin: 0.5em 0; text-align: center;">PC-Gui: 一款 MIT 开源的轻量级桌面 GUI 框架 🎉</h2>
		
		<blockquote style="background: linear-gradient(135deg, #f5f7fa 0%, #e4e8ec 100%); border-left: 5px solid #667eea; padding: 1em 1.5em; margin: 1.5em 0; border-radius: 0 8px 8px 0; color: #4a5568; font-style: italic;"><p style="margin: 0; line-height: 1.6;">💡 <strong style="color: #7d98c7; font-weight: 700;">核心理念：极速开发 · 极致体积 · 原生性能 · 助力打造用户愿意付费的优质工具</strong></p></blockquote>
		
		
		<p style="line-height: 1.8; color: #d3dae6; margin: 0.8em 0; text-align: justify;">
		在桌面端，用户对高效、实用工具的需求从未减弱，并且拥有强烈的付费意愿。<br />		 
		PC-Gui 旨在帮助开发者快速响应这一市场需求，用最简单、最稳定的技术，构建出小巧而强大的商业级桌面应用。
		</p>
		
		<hr>
		
		<h3 style="font-size: 1.4em; font-weight: 600; color: #c4cee0; margin-top: 1.2em; margin-bottom: 0.6em;">核心技术栈</h3>
		
		<p style="line-height: 1.8; color: #d3dae6; margin: 0.8em 0; text-align: justify;">
		摒弃了复杂的依赖，臃肿的第三方 GUI 库，回归编程的本质：<strong style="color: #d3dae6; font-weight: 700;">用后端思维构建桌面应用</strong>。
		<br />通过一个稳定的 Go 后端提供 Web 服务，驱动一个灵活的 Web 前端界面，实现了无与伦比的轻量化与性能。
		</p>
		
		<table style="width: 100%; border-collapse: collapse; margin: 1.5em 0; box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1); border-radius: 8px; overflow: hidden;">
		  <thead style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);">
			<tr>
			  <th style="padding: 1em; text-align: left; color: white; font-weight: 600; border-bottom: 2px solid rgba(255,255,255,0.2);">组件</th>
			  <th style="padding: 1em; text-align: left; color: white; font-weight: 600; border-bottom: 2px solid rgba(255,255,255,0.2);">技术详情</th>
			</tr>
		  </thead>
		  <tbody>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">后端服务</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Go 语言，基于标准库 <code style="background: #edf2f7; color: #e53e3e; padding: 0.2em 0.4em; border-radius: 4px; font-family: \'Consolas\', monospace; font-size: 0.9em;">net/http</code> 提供本地 Web 服务。</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">前端界面</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">HTML, JavaScript, CSS 标准 Web 技术。</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">数据存储</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">本地加密的 SQLite 数据库，轻量、可靠。</td>
			</tr>
		  </tbody>
		</table>
		
		
		<p style="line-height: 1.8; color: #d3dae6; margin: 0.8em 0; text-align: justify;">
		<br>
		</p>
		
		<h3 style="font-size: 1.4em; font-weight: 600; color: #c4cee0; margin-top: 1.2em; margin-bottom: 0.6em;">核心优势 & 多方案对比</h3>
		
		<table style="width: 100%; border-collapse: collapse; margin: 1.5em 0; box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1); border-radius: 8px; overflow: hidden;">
		  <thead style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);">
			<tr>
			  <th style="padding: 1em; text-align: left; color: white; font-weight: 600; border-bottom: 2px solid rgba(255,255,255,0.2);">类别</th>
			  <th style="padding: 1em; text-align: left; color: white; font-weight: 600; border-bottom: 2px solid rgba(255,255,255,0.2);">✅ PC-Gui 优势</th>
			  <th style="padding: 1em; text-align: left; color: white; font-weight: 600; border-bottom: 2px solid rgba(255,255,255,0.2);">⚠️ 其他方案对比</th>
			</tr>
		  </thead>
		  <tbody>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">🚀 零依赖运行</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">后端Go 语言</strong>极速开发，强类型易于维护；交叉编译，生成单一可执行文件，无需用户安装任何运行时或依赖库，双击即可运行。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">⚠️需要用户预装并配置  WebView2, Python、C++, Node.js 等复杂的环境和依赖。</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">🎨界面技术 (HTML)</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">HTML</strong> 前端界面可借助海量模板与 AI 工具快速生成，不仅效率极高，还能轻松打造出精美、现代的视觉风格。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">传统 GUI 库界面通常较为陈旧，自定义难度高。</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">AI 流式输出</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">仅需简单的异步处理，即可实现 AI 内容的流式输出，提升用户体验。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">实现流式输出通常需要处理复杂的回调或多线程。</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">Markdown 渲染</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">完美渲染 AI 返回的 Markdown 格式，并支持各类语言的语法高亮。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Chatbox、Cherry等对 Markdown 渲染及代码高亮效果较为朴素。</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">单文件部署</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">利用 Go 标准库中的 embed，可以将所有静态资源（如图片、CSS 等）直接打包到单一可执行文件中，并复用 HTML 服务进行直接访问。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">依赖臃肿:需借助外部工具打包，产物体积庞大或文件零散，部署复杂。</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">📦 极致体积</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">应用打包后体积仅 <strong style="color: #d3dae6; font-weight: 700;">10-25MB</strong>，分发和下载毫无压力。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">⚠️ 基于 Electron / WebView2 的应用体积普遍在 <strong style="color: #d3dae6; font-weight: 700;">100MB</strong> 以上。</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">🧠 超低内存占用</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">运行时内存占用仅约 <strong style="color: #d3dae6; font-weight: 700;">8MB</strong>，CPU 开销近乎为零，轻快如飞。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">⚠️ Electron / WebView2 应用内存占用轻松达到 <strong style="color: #d3dae6; font-weight: 700;">500MB</strong> 以上。</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">代码安全性</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Go 编译后的二进制文件,结合 garble 混淆技术，有效防止逻辑被反编译。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">易泄露:Python、Node.js 脚本语言极易被反编译、扒光，毫无商业机密。</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">💻跨平台兼容</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Go 语言原生支持 Windows 7/10/11, Linux, macOS，覆盖最广泛的用户群体。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Webview2 等方案不支持 Windows 7 等旧版系统。</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">运行稳定性</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">核心仅依赖 Go 官方标准库，可长期稳定运行不崩溃。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">依赖大量第三方库，版本兼容性和稳定性存在风险。</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">💯 完全掌控</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">核心代码仅依赖 Go 官方标准库，<strong style="color: #d3dae6; font-weight: 700;">无任何第三方 GUI 框架黑盒</strong>，代码完全自主可控，便于长期维护与排查问题。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">⚠️ 依赖大型第三方框架，代码冗余，遇到疑难杂症时排查困难。</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">🌐 全球化支持</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">界面基于标准网页，可直接利用浏览器的<strong style="color: #d3dae6; font-weight: 700;">内置翻译功能</strong>，轻松支持全球数百种语言。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">需要内置多语言文本库，工作量巨大。</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">💡跨语言复用</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">框架思路清晰，任何支持 HTTP 服务的语言（如 C#, Python, Rust）均可借鉴。</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">框架与特定语言或平台深度绑定，难以迁移。</td>
			</tr>
		  </tbody>
		</table>
		
		
		<p style="line-height: 1.8; color: #d3dae6; margin: 0.8em 0; text-align: justify;">
		<br>
		</p>
		
		<h2 style="font-size: 1.8em; font-weight: 600; color: #d3dae6; border-bottom: 3px solid #667eea; padding-bottom: 0.3em; margin-top: 1.5em; margin-bottom: 0.8em;">致开发者</h2>
		
		<p style="line-height: 1.8; color: #d3dae6; margin: 0.8em 0; text-align: justify;">
		在 AI 浪潮席卷全球、就业市场面临挑战的今天，掌握一门能够快速创造价值的技能至关重要。
		</p>
		
		<p style="line-height: 1.8; color: #d3dae6; margin: 0.8em 0; text-align: justify;">
		希望 PC-Gui 这套轻量、高效的框架，能成为您手中的利器，帮助您快速将创意变为现实，开发出用户愿意付费的桌面实用工具，最终实现实现个人价值与商业创收。
		</p>
		
		<hr>
		

		
		<h2 style="font-size: 1.8em; font-weight: 600; color: #d3dae6; border-bottom: 3px solid #667eea; padding-bottom: 0.3em; margin-top: 1.5em; margin-bottom: 0.8em;">📜 MIT 开源授权许可</h2>
		
		<p style="line-height: 1.8; color: #d3dae6; margin: 0.8em 0; text-align: justify;">
		本项目采用 <strong style="color: #d3dae6; font-weight: 700;">MIT 许可证</strong>。这意味着您可以完全自由地使用框架源码。
		</p>
		
		<p style="line-height: 1.8; color: #d3dae6; margin: 0.8em 0; text-align: justify;">
		您可以：(通俗解释 )<br />
		 
		-   ✅ <strong style="color: #d3dae6; font-weight: 700;">商业使用</strong>：允许将本作品及其衍生品用于商业目的，并进行销售。<br />
		 
		-   ✅ <strong style="color: #d3dae6; font-weight: 700;">修改分发</strong>：允许修改代码，并以开源或闭源的形式重新分发。<br />
		 
		-   ✅ <strong style="color: #d3dae6; font-weight: 700;">私人使用</strong>：允许在私人场合使用和修改。<br />
		</p>
		


		
		

		公开源码下载网址:|替换源码下载网址| 

		<p style="line-height: 1.8; color: #d3dae6; margin: 0.8em 0; text-align: justify;">
		
		</p>
				<hr>
			</div>
			<hr>
			`

	} else {

		V299V = `		<hr>	This program uses a desktop GUI framework.
									<div style="max-width: 110%;">
		<h2 style="font-weight: 700; background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text; margin: 0.5em 0; text-align: center;">PC-Gui: An MIT-licensed open-source lightweight desktop GUI framework 🎉</h2>
		
		
		
		<blockquote style="background: linear-gradient(135deg, #f5f7fa 0%, #e4e8ec 100%); border-left: 5px solid #667eea; padding: 1em 1.5em; margin: 1.5em 0; border-radius: 0 8px 8px 0; color: #4a5568; font-style: italic;"><p style="margin: 0; line-height: 1.6;">💡 <strong style="color: #7d98c7; font-weight: 700;">Core Philosophy: Rapid Development · Minimal Size · Native Performance · Empowering Developers to Build Premium Tools Users Are Willing to Pay For</strong></p></blockquote>
		
		
		<p style="line-height: 1.8; color: #d3dae6; margin: 0.8em 0; text-align: justify;">
		On the desktop, users' demand for efficient, practical tools has never waned, and they have a strong willingness to pay.<br />		 
		PC-Gui aims to help developers quickly respond to this market need, using the simplest and most stable technologies to build compact yet powerful commercial-grade desktop applications.
		</p>
		
		<hr>
		
		<h3 style="font-size: 1.4em; font-weight: 600; color: #4a5568; margin-top: 1.2em; margin-bottom: 0.6em;">Core Tech Stack</h3>
		
		<p style="line-height: 1.8; color: #d3dae6; margin: 0.8em 0; text-align: justify;">
		Abandoning complex dependencies and bloated third-party GUI libraries, it returns to the essence of programming: <strong style="color: #d3dae6; font-weight: 700;">Building desktop applications with a backend mindset</strong>.
		<br />By providing web services through a stable Go backend that drives a flexible web frontend, it achieves unparalleled lightweight and performance.
		</p>
		
		<table style="width: 100%; border-collapse: collapse; margin: 1.5em 0; box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1); border-radius: 8px; overflow: hidden;">
		  <thead style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);">
			 <tr>
			  <th style="padding: 1em; text-align: left; color: white; font-weight: 600; border-bottom: 2px solid rgba(255,255,255,0.2);">Component</th>
			  <th style="padding: 1em; text-align: left; color: white; font-weight: 600; border-bottom: 2px solid rgba(255,255,255,0.2);">Technical Details</th>
			 </tr>
		  </thead>
		  <tbody>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">Backend Service</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Go language, providing local web services based on the standard library <code style="background: #edf2f7; color: #e53e3e; padding: 0.2em 0.4em; border-radius: 4px; font-family: \'Consolas\', monospace; font-size: 0.9em;">net/http</code>.</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">Frontend Interface</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Standard web technologies: HTML, JavaScript, CSS.</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">Data Storage</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Locally encrypted SQLite database, lightweight and reliable.</td>
			</tr>
		  </tbody>
		</table>
		
		
		<p style="line-height: 1.8; color: #d3dae6; margin: 0.8em 0; text-align: justify;">
		<br>
		</p>
		
		<h3 style="font-size: 1.4em; font-weight: 600; color: #4a5568; margin-top: 1.2em; margin-bottom: 0.6em;">Core Advantages & Multi-Solution Comparison</h3>
		
		<table style="width: 100%; border-collapse: collapse; margin: 1.5em 0; box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1); border-radius: 8px; overflow: hidden;">
		  <thead style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);">
			 <tr>
			  <th style="padding: 1em; text-align: left; color: white; font-weight: 600; border-bottom: 2px solid rgba(255,255,255,0.2);">Category</th>
			  <th style="padding: 1em; text-align: left; color: white; font-weight: 600; border-bottom: 2px solid rgba(255,255,255,0.2);">✅ PC-Gui Advantages</th>
			  <th style="padding: 1em; text-align: left; color: white; font-weight: 600; border-bottom: 2px solid rgba(255,255,255,0.2);">⚠️ Other Solutions Comparison</th>
			 </tr>
		  </thead>
		  <tbody>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">🚀 Zero-Dependency Runtime</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">Backend Go language</strong> enables rapid development, strong typing for easy maintenance; cross-compilation generates a single executable file, requiring no runtime or dependency installation from users—just double-click to run.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">⚠️ Requires users to pre-install and configure complex environments and dependencies such as  WebView2, Python, C++, Node.js, etc.</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">🎨 Interface Technology (HTML)</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">HTML</strong> frontend interface can be rapidly generated using numerous templates and AI tools, offering extremely high efficiency and making it easy to create beautiful, modern visual styles.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Traditional GUI library interfaces are often outdated and difficult to customize.</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">AI Streaming Output</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">With simple asynchronous handling, AI content streaming can be achieved, enhancing user experience.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Implementing streaming output typically requires dealing with complex callbacks or multithreading.</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">Markdown Rendering</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Perfectly renders Markdown format returned by AI and supports syntax highlighting for various languages.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Chatbox, Cherry, etc., have relatively basic Markdown rendering and code highlighting.</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">Single-File Deployment</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">By using the embed package in the GoLang standard library, all static resources (such as images, CSS, etc.) can be directly packaged into a single executable file, while reusing the HTML service for direct access.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Bloated dependencies: requires external tools for packaging, resulting in large output size or scattered files, making deployment complex.</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">📦 Minimal Size</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">The packaged application size is only <strong style="color: #d3dae6; font-weight: 700;">10-25MB</strong>, making distribution and download effortless.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">⚠️ Applications based on Electron / WebView2 typically exceed <strong style="color: #d3dae6; font-weight: 700;">100MB</strong>.</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">🧠 Ultra-Low Memory Usage</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Runtime memory usage is only about <strong style="color: #d3dae6; font-weight: 700;">8MB</strong>, with near-zero CPU overhead, fast and light.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">⚠️ Electron / WebView2 applications easily exceed <strong style="color: #d3dae6; font-weight: 700;">500MB</strong> of memory usage.</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">Code Security</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Go compiled binaries, combined with garble obfuscation technology, effectively prevent logic from being decompiled.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Vulnerable to leaks: scripting languages like Python and Node.js are easily decompiled and exposed, with no trade secret protection.</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">💻 Cross-Platform Compatibility</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Go language natively supports Windows 7/10/11, Linux, macOS, covering the widest range of users.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Solutions like Webview2 do not support older systems such as Windows 7.</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">Runtime Stability</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">The core only depends on the Go official standard library, ensuring long-term stable operation without crashes.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Relies on a large number of third-party libraries, posing risks in version compatibility and stability.</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">💯 Full Control</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Core code only depends on the Go official standard library, <strong style="color: #d3dae6; font-weight: 700;">with no third-party GUI framework black boxes</strong>, making the code fully autonomous and controllable, facilitating long-term maintenance and troubleshooting.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">⚠️ Relies on large third-party frameworks, leading to code redundancy and difficulty in troubleshooting complex issues.</td>
			</tr>
			<tr style="background-color: #ffffff; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">🌐 Globalization Support</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">The interface is based on standard web pages and can directly leverage the browser's <strong style="color: #d3dae6; font-weight: 700;">built-in translation feature</strong>, easily supporting hundreds of languages worldwide.</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Requires built-in multilingual text libraries, a massive amount of work.</td>
			</tr>
			<tr style="background-color: #f7fafc; transition: background-color 0.2s;">
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;"><strong style="color: #d3dae6; font-weight: 700;">💡 Cross-Language Reusability</strong></td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">The framework concept is clear and can be adopted by any language that supports HTTP services (e.g., C#, Python, Rust).</td>
			  <td style="padding: 1em; border-bottom: 1px solid #e2e8f0; color: #4a5568; line-height: 1.6;">Frameworks deeply tied to specific languages or platforms are difficult to migrate.</td>
			</tr>
		  </tbody>
		</table>
		
		
		<p style="line-height: 1.8; color: #d3dae6; margin: 0.8em 0; text-align: justify;">
		<br>
		</p>
		
		<h2 style="font-size: 1.8em; font-weight: 600; color: #d3dae6; border-bottom: 3px solid #667eea; padding-bottom: 0.3em; margin-top: 1.5em; margin-bottom: 0.8em;">To Developers</h2>
		
		<p style="line-height: 1.8; color: #d3dae6; margin: 0.8em 0; text-align: justify;">
		In today's world, where the AI wave is sweeping the globe and the job market faces challenges, it is crucial to master a skill that can quickly create value.
		</p>
		
		<p style="line-height: 1.8; color: #d3dae6; margin: 0.8em 0; text-align: justify;">
		We hope that PC-Gui, this lightweight and efficient framework, will become a powerful tool in your hands, helping you quickly turn ideas into reality, develop desktop utilities that users are willing to pay for, and ultimately achieve personal value and commercial revenue.
		</p>
		
		<hr>
		
		
		
		<h2 style="font-size: 1.8em; font-weight: 600; color: #d3dae6; border-bottom: 3px solid #667eea; padding-bottom: 0.3em; margin-top: 1.5em; margin-bottom: 0.8em;">📜 MIT Open Source License</h2>
		
		<p style="line-height: 1.8; color: #d3dae6; margin: 0.8em 0; text-align: justify;">
		This project is licensed under the <strong style="color: #d3dae6; font-weight: 700;">MIT License</strong>. This means you are completely free to use the framework's source code.
		</p>
		
		<p style="line-height: 1.8; color: #d3dae6; margin: 0.8em 0; text-align: justify;">
		You can: (plain English explanation)<br />
		 
		-   ✅ <strong style="color: #d3dae6; font-weight: 700;">Commercial Use</strong>: Use the work and its derivatives for commercial purposes and sell them.<br />
		 
		-   ✅ <strong style="color: #d3dae6; font-weight: 700;">Modify and Distribute</strong>: Modify the code and redistribute it under open-source or closed-source licenses.<br />
		 
		-   ✅ <strong style="color: #d3dae6; font-weight: 700;">Private Use</strong>: Use and modify the code for personal purposes.<br />
		</p>
		
		<hr>
		
			
		Open source download URL:|替换源码下载网址| 

		<p style="line-height: 1.8; color: #d3dae6; margin: 0.8em 0; text-align: justify;">
		
		</p>
				<hr>
		
		
			</div>
			<hr>`

	}

	return

}
