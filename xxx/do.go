package xxx

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func do(method string, url string, body io.Reader) io.Writer {
	// Request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Request Header
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Cookie", "CLIENTKEY=9135-3124-0718; Hm_lvt_f8f6a0064a3e891522bdf044119d462a=1678603948,1680138999; CLIENTKEY_ShowLogin=2171-5775-3983; .ADWASPX7A5C561934E_PCEGGS=64B53F15AC41509FEB7686CED24224D456DF21952B173961795E574C7202113A45EFAA8820AABD244A21C306DA0A7118808BBE46A70FD152B9A77C65C9D4945281A93D1A983B74990A3D7C58981B59654E5512A43A5DBD5FBC18DE70152B082DC52D4C3DD8DB030EA939782B064733E560548EC461C5CD4002C2B98B2FE187D9A19B7903; ckurl.pceggs.com=ckurl=http://www.pceggs.com/game/gameindex/gameindex.aspx?gameid=4; Hm_lpvt_f8f6a0064a3e891522bdf044119d462a=1680139006")
	req.Header.Set("Origin", "http://manorapp.pceggs.com")
	req.Header.Set("Pragma", "http://manorapp.pceggs.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")

	// Response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	// Buffer
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		log.Fatal(err.Error())
	}

	return buf
}
