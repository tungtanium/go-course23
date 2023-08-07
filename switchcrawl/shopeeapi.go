package main

import (
	"fmt"
	"io"
	"net/http"
)

func IsomniaGeneratedCode() {

	url := "https://shopee.vn/api/v4/pdp/get_pc?shop_id=629426402&item_id=10192100450"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("62350719", "Ct_Pp.qh?-1MK86C#BYidW`DKhJTPeD#@:seE@=+r4=M71<r6[-?oBE'VPe68GZO.6H'@NU+uURB;-[P,L&/6:0RELKLeJ<P15[ojTEW\u0021,?sE':CrE\u0021iibE=f;_+?@lIqV6M>D#T?biP`,>]k.HH5OPhe,Td=^NNT8%QMr?D<?)tt^;[bbfN_EVR67*@@\"KGtS;#W3V]a&\u00215-;/TlcFR<s%#-XPu7m=V)UU]Y3baMDZ-ZL$]3PdI%9oC^Y8<u7bT'h;3@R[\\9#c),Cho?9n\\u<ma>m21>$0,IFlQd+%ahmA1l;Zai0BS`2Lc7PP4m8A\"NXU*$m\\2E]Qjehj&FUUGR##6klks0W^ceO<W6ETZ6q8Ph")
	req.Header.Add("cookie", "REC_T_ID=bda626e2-854d-11eb-a059-1409dcb3ab07; REC_T_ID=bdb3ecdb-854d-11eb-baea-806933898349; SPC_F=utob1YjkuSPYjpfpEQJZZCSiiWy9mL82; SPC_CLIENTID=dXRvYjFZamt1U1BZdsudchxhuyqcajzh; SPC_ST=.eXBCT1FRb2pRTEwxa2t3d9Xqb7Eb5YUkVLo9JP4B0jcPFGkjTRP2ufM/Wf7BFJtYz1f9wmm+W/Yr2MVt8nZvxDWWu4eQIy3wdVTpyhByuzGHK5PFCFyMp48CGyR2hAvsams+gfN69qzf+2qj3rhNkR9y+RGD0PhXGDuu/lTeMV8cnXm544575Os79mffXd9eROJPfpL2385e1ZcjvKB83A==; SPC_U=839803496; SPC_R_T_IV=ZThIellyOEMwTjJCeW15TQ==; SPC_T_ID=/q+1mKepFK0161FGos4AdTZKPEPutonKjJCmI7eKyNWYF+u//7ngU8WPteiobIHhA2kOlsciM0h7NEDRokB8NcOLHiY7N06Mtl+2uuYaCHLy7B25kmH/5luh4rvwlYCmqBpm1xoB5icwF6+bqkPrXMfqLQAw77cbBu/HnuNq9dI=; SPC_T_IV=ZThIellyOEMwTjJCeW15TQ==; SPC_R_T_ID=/q+1mKepFK0161FGos4AdTZKPEPutonKjJCmI7eKyNWYF+u//7ngU8WPteiobIHhA2kOlsciM0h7NEDRokB8NcOLHiY7N06Mtl+2uuYaCHLy7B25kmH/5luh4rvwlYCmqBpm1xoB5icwF6+bqkPrXMfqLQAw77cbBu/HnuNq9dI=; SPC_SI=HsHIZAAAAABRY09IUFEwSOijCgAAAAAAUkVvamxzdks=; _QPWSDCXHZQA=a843f33c-64cc-47d2-efd0-2f6b5efdb083; SPC_SC_TK=5e6a8e890bc294587169864271c976a7; SPC_SC_UD=839803496; SPC_STK=fDMuXcmasw+EEWpBkBmhkSpskV64y+n4x+KZpXr/G2XUYYAPYT6OrteZPdBFn3z8bxKfHPZRCaT9wvWXHKUtRkJr5V80DeFpXfRd0elRa+rmyyXSYSMX7HRwOBExIm2WwrwAz/0VB1p/WiiQu31t29mYpMN6dMFy3iwRRuoSycU=; SC_DFP=OxYCeJnEXbakUrxxfHlJavHcFzQXuTCL; csrftoken=r58GFeXTKywxjBbou9nMnE8cRrzd6be5; SPC_IA=1; SPC_EC=SlZOZDFyRWVvN1dUWjhXNP94exMZ9MP+UERLoknBOWIbUzxtYj4sdQZA0Nr0pqNA5TKcSxaLLAvwpqxdRckrxM00cGhaB+0UwTh7r3TTBhcaZ77qUrp0vDKFhBjOPhOEnn+YkqQZy3JD//+1LucipZXkG7N5N1KcE8exdzAW6bQ=; shopee_webUnique_ccd=hjs%2F9%2Fm%2FCZ669BRi3m50eA%3D%3D%7CJfYntApJsHj2zHrBKZxXBjcaEV4mS%2FphWYqTj7%2BaWpKfR%2FzJowZ8dgiJsC1nPzxVjjWawgdBqv%2FkpQ%3D%3D%7C41%2F%2Fa04Fp4hKDXfd%7C08%7C3; ds=cffd687cd42dd6bd54ad352c0c6d1533")
	req.Header.Add("498b672c", "(5r\u00212dC2s_qEp^P,(;CtI2_\\T")
	req.Header.Add("authority", "shopee.vn")
	req.Header.Add("accept", "application/json")
	req.Header.Add("accept-language", "en-GB,en;q=0.9,en-US;q=0.8,vi;q=0.7")
	req.Header.Add("af-ac-enc-dat", "AAcyLjkuMi0yAAABicVg1YQAAA/uAyAAAAAAAAAAAj+vsCN1JmrO07Nj/3HXJNy9QGsY8t2XUcjkdYeEuFx4oF5PgwahHgRTpavd08DCpHf03QqLA3lGw3O0C8i62TArI+CnNcU4eotcTl87zP8GwJ2b8/e9SLULEl7rs5z/kgGsDT0qf/XbSAhfgh1wAi5B7O1p0DANfYttJLBUsIQH2OaeOeO84wTESRwang2hc2zYva87nk216dH/BYGLgq6XiRrOhSCV9eJiLcYxZuWwUpJa9er74qe510iif7EIkY20MTulwYOE/ES3+T32U4Y3GInDnYbq6NoxU5JSt4VK1GuXxf++tGS9C/ya8h53KhwcDI7+N/b85J63HZXluidukfEplrT70pdqAp/LXAlLC0UwNSZC1hAVZjppuV4eqmBgsXo2NEAZt1bM9xSYJxKrYowJt/KhTukZQsnGJ3BhdNq9t5f/NhMnw8f21UXD2mkb2JvlFQZdSsnqOggu2GJqO7WkUQJfq3K1AG1EpKO616ktiwbY7P9KHML6aQtF4FlKd6wIZ7ugf/+klKjC7cTH7sR8wbX2+obBZ0Ft7Itc/pw0Wp7JqRbQi4qOBfX5l99KLTRetMvPd7B24O0bbJV1UmSWwbX2+obBZ0Ft7Itc/pw0Wp7JqRbQi4qOBfX5l99KLTQax8fiGjC0OB/S2DbFR/RPwsXNNYTyK3HSzCGoV/4BXUckfn/aqZ2EUBMC3PvgG/UYz6Rdg8a/lQzfapMRlwDxYH6oJSbUnFvFBdt+DtZeJ0ckfn/aqZ2EUBMC3PvgG/UOGNnJy3Wnw29vQcKonkqmxQyKBsKdMmAaH144fpIGJkWp+KOsMG49WlUfad0aZT2XUcIvcDzRL2N9M2Ja32QNGhtQHiFBbNAhh9c+IqgMSJf/NhMnw8f21UXD2mkb2JuX/zYTJ8PH9tVFw9ppG9ibp4C441kgdALq5ekJVPxMnA1sO8pQEaDD8zXt8x5oERKmgopehCL+3aPsA6YagwnuHINqdWOAy0MNkx5lkwfOJKnnA5lSo8Wp5u8S/TGXP15IA/TtWmBoI5o7mG3pL/q/foTUxxk6uAZIb9OKIF3Y9w==")
	req.Header.Add("af-ac-enc-sz-token", "hjs/9/m/CZ669BRi3m50eA==|JfYntApJsHj2zHrBKZxXBjcaEV4mS/phWYqTj7+aWpKfR/zJowZ8dgiJsC1nPzxVjjWawgdBqv/kpQ==|41//a04Fp4hKDXfd|08|3")
	req.Header.Add("be873082", "+)b$0p$C5[ru/A\"Fu)O#MF>&4")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("dnt", "1")
	req.Header.Add("referer", "https://shopee.vn/Switch-AKKO-CS-b%C3%A1n-l%E1%BA%BB-c%C3%B4ng-t%E1%BA%AFc-ph%C3%ADm-c%C6%A1-c%C3%B3-lube-s%E1%BA%B5n-i.629426402.10192100450?sp_atk=04a8ed60-7236-4004-93a7-ba9812e6d902&xptdk=04a8ed60-7236-4004-93a7-ba9812e6d902")
	req.Header.Add("sec-ch-ua", "\"Not/A)Brand\";v=\"99\", \"Microsoft Edge\";v=\"115\", \"Chromium\";v=\"115\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sz-token", "hjs/9/m/CZ669BRi3m50eA==|JfYntApJsHj2zHrBKZxXBjcaEV4mS/phWYqTj7+aWpKfR/zJowZ8dgiJsC1nPzxVjjWawgdBqv/kpQ==|41//a04Fp4hKDXfd|08|3")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/116.0")
	// req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.188")
	req.Header.Add("x-api-source", "pc")
	req.Header.Add("x-csrftoken", "r58GFeXTKywxjBbou9nMnE8cRrzd6be5")
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("x-sap-ri", "452dce64d6f8c3f46b24b03f66c9a44bd352a1505b7200e7")
	req.Header.Add("x-shopee-language", "vi")
	req.Header.Add("x-sz-sdk-version", "2.9.2-2&1.4.1")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
