# Capitulo: Pacotes Importantes
## Aula 2: Realizando chamadas HTTP

Nesta aula vamos aprender a realizar chamadas HTTP, isso é muito utilizado para comunicação entre sistemas.

Vamos fazer um exemplo simples, e incrementar em aulas futuras.

Neste exemlos vamos fazer uma request no site do GOOGLE e vamos imprimir o código da pagina no console

```go
package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://www.google.com.br")
	if err != nil {
		panic(err)
	}
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))

	err = req.Body.Close()
	if err != nil {
		return
	}
}
```
output:
```shell
<!doctype html><html itemscope="" itemtype="http://schema.org/WebPage" lang="pt-BR"><head><meta content="text/html; charset=UTF-8" http-equiv="Content-Type"><meta content="/images/branding/googleg/1x/googleg_standard_color_128dp.png" itemprop="image"><title>Google</title><script nonce="H7wFK39CVWjzapgv9UZ1wQ">(function(){window.google={kEI:'WFaFY8ibN9q31sQPqvaz2Ac',kEXPI:'0,1359409,6059,206,4804,2316,383,246,5,5367,1123753,1197754,647,380090,16114,28684,22430,1362,12313,17586,4998,13228,3847,10626,22737,5081,887,708,1277,2742,149,1103,840,6297,3514,606,2023,2297,6342,8328,3227,2845,8,33769,1851,15756,3,346,230,6460,148,13975,4,1528,2304,7039,17819,2490,1714,5708,4164,3193,13658,4437,16786,5827,2530,4094,17,4035,3,3541,1,42154,2,14022,2373,23366,5679,1020,2381,28741,4568,6256,23419,1248,5841,14967,4333,7484,445,2,2,1,7171,10135,9326,8155,6682,699,1480,14489,874,19634,6,1922,5784,3995,7782,11348,1509,15515,23803,2700,5797,11,9798,384,4147,14,82,950,2940,751,2070,642,298,11869,3,683,217,841,564,960,819,668,1186,1997,1119,6,8121,3,7,859,242,3248,2411,1741,814,1542,551,1767,84,600,253,564,565,403,588,1391,87,52,565,350,4,86,495,59,368,450,473,110,36,6,291,3918,400,401,893,136,3,785,2,318,1106,1017,324,57,645,2,3,1115,246,65,367,1,57,727,229,1103,14,89,890,57,556,4,6,291,311,252,420,245,53,483,1740,160,50,25,4,3,252,192,55,1099,115,179,280,4,778,86,121,633,220,542,231,63,619,355,65,664,44,4,125,150,347,294,444,198,86,108,22,175,67,328,36,250,5,48,125,5,197,131,315,147,557,274,485,193,138,230,1045,1,495,17,39,11,1087,5289751,6043,51,5995698,2803331,3311,141,801,19729,1,1,346,2759,41,169,140,23947376,487,25,18,2861496,1180116,1964,3094,13578,3406,13968,14',kBL:'-1VY'};google.sn='webhp';google.kHL='pt-BR';})();(function(){```
var f=this||self;var h,k=[];function l(a){for(var b;a&&(!a.getAttribute||!(b=a.getAttribute("eid")));)a=a.parentNode;return b||h}function m(a){for(var b=null;a&&(!a.getAttribute||!(b=a.getAttribute("leid")));)a=a.parentNode;return b} 
```

Para isso vamos utilizar o pacote `http`
- Realizamos a requisição com a função `Get()`
- Lemos o resultado com a função `io.ReadAll()`
- Convertemos o resultado que é um slice de bytes para string e imprimimos
- E por fim fechamos o `Body` com `req.Body.Close()`

