package themes

var (
	// builtinCodeStyles is a map of the themes names corresponding to chroma styles
	// See the Chroma Style Gallery for more styles: https://xyproto.github.io/splash/docs/
	// "default" currently points to "material"
	builtinCodeStyles = map[string]string{"material": "lovelace", "gray": "manni", "dark": "dracula", "redbox": "fruity", "bw": "bw", "wing": "tango", "neon": "api", "light": "monokailight", "werc": "dracula"}

	// Themes is a map over the available built-in CSS themes. Corresponds with the font themes below.
	// "default" and "gray" are equal. "default" should never be used directly, but is here as a safeguard.
	builtinThemes = map[string][]byte{
		"default":  []byte(`body { background-color: #e7eaed; color: #0b0b0b; font-family: 'Trebuchet MS', Helvetica, sans-serif; font-weight: 300; margin: 1em; font-size: 1em; } a { color: #401010; font-family: Lato, courier; } a:hover { color: #801010; } a:active { color: yellow; } h1, h2, h3, h4, h5, h6 { color: #101010; margin-left: 0; margin-right: 0; } img { max-width: 100%; } table { font-family: Helvetica, sans-serif; border-collapse: collapse; width: 100%; font-weight: bold; } table td, table th { border: 1px solid #ccc; padding: 0.45em 0.55em 0.45em 0.55em; } table tr:nth-child(even){ background-color: #efefef; } table tr:nth-child(odd) { background-color: #fff; } table tr:hover { background-color: #e0f0ff; color: black; } table th { padding-top: 12px; padding-bottom: 12px; text-align: left; background-color: #666; color: white; } .chroma { white-space: pre-wrap; border-radius: 0.3em; padding: 0.1em 0.4em 0.1em 0.4em; word-wrap: break-word; }`),
		"gray":     []byte(`body { background-color: #e7eaed; color: #0b0b0b; font-family: 'Trebuchet MS', Helvetica, sans-serif; font-weight: 300; margin: 4.5em; font-size: 1em; } a { color: #4a4a4a; font-family: 'Trebuchet MS', Helvetica, sans-serif; } a:hover { color: #6a6a6a; } a:active { color: #0b0b0b; } h1 { color: #101010; } img { max-width: 100%; } table { font-family: Helvetica, sans-serif; border-collapse: collapse; width: 100%; font-weight: bold; } table td, table th { border: 1px solid #ccc; padding: 0.45em 0.55em 0.45em 0.55em; } table tr:nth-child(even){ background-color: #efefef; } table tr:nth-child(odd) { background-color: #fff; } table tr:hover { background-color: #e0f0ff; color: black; } table th { padding-top: 12px; padding-bottom: 12px; text-align: left; background-color: #666; color: white; } .chroma { white-space: pre-wrap; border-radius: 0.3em; padding: 0.1em 0.4em 0.1em 0.4em; word-wrap: break-word; }`),
		"light":    []byte(`body { background-color: #eaeaea; color: #0b0b0b; font-family: 'Trebuchet MS', Helvetica, sans-serif; font-weight: 300; margin: 4.5em; font-size: 1em; } a { color: #0b0b0b; font-family: 'Trebuchet MS', Helvetica, sans-serif; text-decoration: underline; } a:hover { color: #2a2a2a; } a:active { color: #0b0b0b; } h1 { color: #101010; } img { max-width: 100%; } table { font-family: Helvetica, sans-serif; border-collapse: collapse; width: 100%; font-weight: bold; } table td, table th { border: 1px solid #ccc; padding: 0.45em 0.55em 0.45em 0.55em; } table tr:nth-child(even){ background-color: #efefef; } table tr:nth-child(odd) { background-color: #fff; } table tr:hover { background-color: #e0f0ff; color: black; } table th { padding-top: 12px; padding-bottom: 12px; text-align: left; background-color: #666; color: white; } .chroma { white-space: pre-wrap; border-radius: 0.3em; padding: 0.1em 0.4em 0.1em 0.4em; word-wrap: break-word; }`),
		"dark":     []byte(`body { background-color: #101010; color: #f0f0f0; font-family: 'Trebuchet MS', Helvetica, sans-serif; font-weight: 400; margin: 4.5em; font-size: 1em; } a { color: #9db5b2; font-family: 'Trebuchet MS', Helvetica, sans-serif; } a:hover { color: #cad2d3; } a:active { color: #f0f0f0; } h1 { color: #f0f0f0; } img { max-width: 100%; } table { font-family: Helvetica, sans-serif; border-collapse: collapse; width: 100%; font-weight: bold; } table td, table th { border: 1px solid #999; padding: 0.45em 0.55em 0.45em 0.55em; } table tr:nth-child(even){ background-color: #222; } table tr:nth-child(odd) { background-color: #333; } table tr:hover { background-color: #333; color: #ff8060; } table th { padding-top: 12px; padding-bottom: 12px; text-align: left; background-color: #d7d7d7; color: #292929; } .chroma { white-space: pre-wrap; border-radius: 0.3em; padding: 0.1em 0.4em 0.1em 0.4em; word-wrap: break-word; }`),
		"redbox":   []byte("html{background-color:#222;}body{color:#111;background-color:#999;font-family:Impact,'Arial Black',sans-serif;font-size:1.7em;margin:2.7em;padding:0 5em 1em 2em;border-radius:50px;border:solid 10px #a00;box-shadow:10px 10px 16px black, 6px 6px 8px #222 inset;}h2{font-family:monospace;color:black;}ul{margin-left:1em;}a{text-decoration:none;color:#900;}a:hover{color:#dc0;}code{font-family:Monofett,'Courier New',Courier,monospace;} img { max-width: 100%; } .chroma { white-space: pre-wrap; border-radius: 0.3em; padding: 0.1em 0.4em 0.1em 0.4em; word-wrap: break-word; }"),
		"bw":       []byte("body { font-family: BlinkMacSystemFont, \"Segoe UI\", Helvetica, sans-serif; margin: 1px; padding: 1px; } button { background: black; border-radius: 4px; color: #ffffff; font-size: 20px; padding: 10px 20px 10px 20px; border: none; outline: none; margin-left: 2px; margin-right: 2px; } button:hover { background: orange; text-decoration: none; } button:disabled { background: gray; color: #eee; } input { font-size: 16px; padding: 10px; border: 1px solid #999; } a { padding: 1px; text-decoration: none; color: black; } a:hover { padding: 1px; border-bottom: 1px solid black; } a.button { font-size: 1em; border: 1px solid black; background-color: white; color: black; width: 150px; padding: 10px 0px; margin-right: 10px; margin-top: 5px; margin-bottom: 5px; text-align: center; display: inline-block; } a.button:hover { text-decoration: none; color: white; background-color: black;  } a.button:active { color: white; background-color: black; } a.go-back { position: absolute; right: 10px; top: 10px; padding: 10px; font-size: 0.8em; background-color: black; color: white; } .ha { box-sizing: border-box; margin: 0 150px; padding: 20px; line-height: 1.5; font-size: 15.5px; line-height: 1.55; } .ha ul { list-style: none; padding: 0 !important; column-count: 2; max-width: 0; column-gap: 20px; min-width: 350px; } .ha ul li { margin-left: 0px; } .ha h1 { margin-top: 0; margin-bottom: 0; } .ha h1 { font-size: 32px; font-weight: 200; margin-top: 0px; margin-bottom: 16px; } .ha h5 { margin-bottom: 2px; } .ha hr { padding: 0; margin: 24px 0; background-color: #e7e7e7; border: 0; height: 0.09em; } .ha li { margin-top: 0.25em; } .gh-btns { float: right; margin-top: -4px; }"),                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  // from MIT licensed HyperApp: hyperapp.glitch.me/style.css
		"wing":     []byte(`body,html{margin:0;padding:0}html{box-sizing:border-box;font-size:62.5%}body{line-height:1.6;font-size:1.5rem;font-weight:400;font-family:-apple-system,BlinkMacSystemFont,Avenir,"Avenir Next","Segoe UI",Roboto,Oxygen,Ubuntu,Cantarell,"Fira Sans","Droid Sans","Helvetica Neue",sans-serif}h1,h2,h3,h4,h5,h6{margin-top:0;margin-bottom:2rem;font-weight:400;font-family:-apple-system,BlinkMacSystemFont,Avenir,"Avenir Next","Segoe UI",Roboto,Oxygen,Ubuntu,Cantarell,"Fira Sans","Droid Sans","Helvetica Neue",sans-serif}h1{font-size:5rem;line-height:1.2}h2{font-size:4.2rem;line-height:1.25}h3{font-size:3.6rem;line-height:1.3}h4{font-size:3rem;line-height:1.35}h5{font-size:2.4rem;line-height:1.5}h6{font-size:1.8rem;line-height:1.6}p{margin-top:0;margin-bottom:2rem;font-size:1.5rem}@media (max-width:400px){h1{font-size:4rem}h2{font-size:3.5rem}h3{font-size:3rem}h4{font-size:2.6rem}h5{font-size:2.2rem}h6{font-size:1.8rem}}a{color:#104cfb;transition:color .1s ease}a:hover{cursor:pointer;color:#222}.button,[type=submit],button{padding:1.1rem 3.5rem;margin-top:0;margin-bottom:2rem;color:#f5f5f5;font-size:1.5rem;background:#111;border:1.5px solid #111;border-radius:2px;transition:all .2s ease}.button.outline,[type=submit].outline,button.outline{color:#111;background:0 0;border:1.5px solid #111}.button:hover,[type=submit]:hover,button:hover{color:#f5f5f5;background:#222}.button.outline:hover,[type=submit].outline:hover,button.outline:hover{background:0 0;border:1.5px solid #222;color:#222}.button:focus,[type=submit]:focus,button:focus{outline:0}.button:active,[type=submit]:active,button:active{transform:scale(.99)}.button.disabled,.button:disabled,[type=submit]:disabled,button:disabled{opacity:.8;cursor:not-allowed}input[type=email],input[type=file],input[type=number],input[type=password],input[type=search],input[type=tel],input[type=text],input[type=url],select,textarea,textarea[type=text]{width:100%;height:45px;padding:10px 10px;margin-top:1rem;margin-bottom:1rem;font-size:1.3rem;background:#f1f1f1;border:1px solid #a4a4a4;border-radius:2px;box-sizing:border-box;transition:all .2s ease}input[type=email]:hover,input[type=file]:hover,input[type=number]:hover,input[type=password]:hover,input[type=search]:hover,input[type=tel]:hover,input[type=text]:hover,input[type=url]:hover,select:hover,textarea:hover,textarea[type=text]:hover{border:1px solid #111}input[type=email]:focus,input[type=file]:focus,input[type=number]:focus,input[type=password]:focus,input[type=search]:focus,input[type=tel]:focus,input[type=text]:focus,input[type=url]:focus,select:focus,textarea:focus,textarea[type=text]:focus{outline:0;border:1px solid #104cfb}textarea,textarea[type=text]{min-height:7rem}.container{max-width:960px;margin:0 auto;width:80%}.row{display:flex;flex-flow:row wrap;justify-content:space-between;width:100%;margin:1rem 0}.row>:first-child{margin-left:0}.col{flex:1 1 0px}.col,[class*=" col-"],[class^=col-]{margin-left:4%}.col-1{width:4.333333333333332%}.col-2{width:12.666666666666664%}.col-3{width:21%}.col-4{width:29.33333333333333%}.col-5{width:37.66666666666667%}.col-6{width:46%}.col-7{width:54.333333333333336%}.col-8{width:62.66666666666666%}.col-9{width:71%}.col-10{width:79.33333333333334%}.col-11{width:87.66666666666666%}.col-12{width:96%}@media screen and (max-width:768px){.col,[class*=" col-"],[class^=col-]{margin:1rem 0;flex:0 0 100%}}ol,ul{padding-left:0;margin-top:0;margin-bottom:2rem;list-style-position:inside}ol ol,ol ul,ul ol,ul ul{margin:1rem 0 1rem 2rem;font-size:95%}li{margin-bottom:1rem}.table{width:100%;border:none;border-collapse:collapse;border-spacing:0;text-align:left}.table td,.table th{vertical-align:middle;padding:12px 4px}.table thead{border-bottom:2px solid #333030}@media screen and (max-width:768px){.table.responsive{position:relative;display:block}.table.responsive td,.table.responsive th{margin:0}.table.responsive thead{display:block;float:left;border:0}.table.responsive thead tr{display:block;padding:0 10px 0 0;border-right:2px solid #333030}.table.responsive th{display:block;text-align:right}.table.responsive tbody{display:block;overflow-x:auto;white-space:nowrap}.table.responsive tbody tr{display:inline-block}.table.responsive td{display:block;min-height:16px;text-align:left}.table.responsive tr{padding:0 10px}}img{margin-top:0;margin-bottom:2rem}.nav{position:relative;display:flex;flex-wrap:wrap;align-items:center;padding-top:1rem;padding-bottom:1rem}.nav-item{padding:1rem 2rem}.cards{display:flex;flex-wrap:wrap;justify-content:space-between;overflow:hidden;margin-bottom:2rem}.card{flex-direction:column;flex:0 1 calc(50% - .5rem);box-shadow:0 5px 15px rgba(0,0,0,.1);margin-bottom:2rem}.card-image{display:block;position:relative}.card-image img{display:block;height:auto;width:100%}.card-header{font-weight:600;margin:0;padding:2rem 3rem 1rem}.card-body{padding:0 3rem 2rem 3rem;min-height:100px}.card-footer{display:flex;align-items:stretch;border-top:1px solid #dbdbdb;flex:1}.card-footer .card-footer-item{display:flex;flex-basis:0;flex-grow:1;flex-shrink:0;align-items:center;justify-content:center;margin:0;padding:1rem}.card-footer-item:not(:first-child){border-left:1px solid #dbdbdb}@media (max-width:768px){.cards>.card{flex:auto}}.pull-right{float:right}.pull-left{float:left}.text-center{text-align:center}.text-left{text-align:left}.text-right{text-align:right}.full-screen{width:100%;min-height:100vh}.full-width{width:100%}.vertical-align{display:flex;align-items:center}.horizontal-align{display:flex;justify-content:center}.center{display:flex;align-items:center;justify-content:center;flex-direction:column}.right{display:flex;align-items:center;justify-content:flex-end}.left{display:flex;align-items:center;justify-content:flex-start}.fixed{position:fixed;width:100%}@media screen and (max-width:400px){.hide-phone{display:none}}@media screen and (max-width:768px){.hide-tablet{display:none}}pre{margin-top:0;margin-bottom:2rem}code{padding:.2rem .5rem;margin:0 .2rem;font-size:1.3rem;white-space:nowrap;background:#f1f1f1;border:1px solid #dbdbdb;border-radius:4px;font-family:Consolas,Monaco,Menlo,monospace}pre>code{display:block;padding:1rem 1.5rem;white-space:pre-wrap;word-wrap:break-word}`), // From https://github.com/kbrsh/wing/blob/master/dist/wing.min.css
		"material": []byte("body { margin: 3em; } pre{margin-top:0;margin-bottom:2rem}code{padding:.2rem .5rem;margin:0 .2rem;font-size:1.3rem;white-space:nowrap;background:#f1f1f1;border:1px solid #dbdbdb;border-radius:4px;font-family:Consolas,Monaco,Menlo,monospace}pre>code{display:block;padding:1rem 1.5rem;white-space:pre-wrap;word-wrap:break-word} img { max-width: 100%; } table { font-family: Helvetica, sans-serif; border-collapse: collapse; width: 100%; font-weight: bold; } table td, table th { border: 1px solid #ccc; padding: 0.45em 0.55em 0.45em 0.55em; } table tr:nth-child(even){background-color: #efefef; } table tr:nth-child(odd) { background-color: #fff; } table tr:hover { background-color: #e0f0ff; color: black; } table th { padding-top: 12px; padding-bottom: 12px; text-align: left; background-color: #666; color: white; } .chroma { white-space: pre-wrap; border-radius: 0.3em; padding: 0.1em 0.4em 0.1em 0.4em; word-wrap: break-word; } code { font-size: 0.8rem; }"),
		"neon":     []byte("html{line-height:1}body{align-items:center;background-color:#111;display:flex;font-family:Helvetica Neue,sans-serif;height:90vh;justify-content:center;margin:0;padding:0;text-align:center}h1{color:#00caff;font-weight:100;font-size:8em;margin:0;padding-bottom:15px}button{background:#111;border:1px solid #00caff;color:#00caff;font-size:2em;font-weight:100;margin:0;outline:none;padding:5px 15px;transition:background .2s;&:hover,&:active,&:disabled{background:#00caff;color:#111}&:active{outline:2px solid #00caff}&:focus{border:1px solid #00caff}}button + button{margin-left:3px}pre{margin-top:0;margin-bottom:2rem}code{padding:.2rem .5rem;margin:0 .2rem;font-size:1.3rem;white-space:nowrap;background:#f1f1f1;border:1px solid #dbdbdb;border-radius:4px;font-family:Consolas,Monaco,Menlo,monospace}pre>code{display:block;padding:1rem 1.5rem;white-space:pre-wrap;word-wrap:break-word}*{box-sizing:border-box;margin:0;padding:0}body{background:#111}.container{align-items:center;display:flex;height:100vh;justify-content:center;text-align:center}input{background:rgba(0,0,0,.2);color:#00caff;border:1px solid #111;font:4em Helvetica Neue,sans-serif;font-weight:300;height:100vh;width:100vw;outline:0;padding:0 40px;position:absolute;&:hover{background:rgba(0,0,0,0)}}img{width:90vw;height:90vh}"),
		"werc":     []byte(`#side-bar a,ul.sitemap-list a{text-transform:capitalize}body{color:#000;background-color:#fff;font-family:Helvetica,Verdana,Arial,'Liberation Sans',FreeSans,sans-serif;font-size:84%;margin:0;padding:0}.superHeader{color:#fff;background-color:#6487dc;height:1.6em}.superHeader img{vertical-align:bottom}.superHeader a{color:#fff;background-color:transparent;font-size:91%;margin:0;padding:0 .5ex 0 .25ex;Xcolor:#000}a{text-decoration:none}a:hover{text-decoration:underline}#main-copy .topOfPage,#side-bar a:hover,.headerTitle a:hover{text-decoration:none}.superHeader div{position:absolute;top:.4ex}.superHeader .left{left:.4em}.superHeader .right{right:.4em}.midHeader{color:#274e90;border:0 solid #000;border-width:2px 0}table,td,th{border:1px solid #000}.headerTitle{color:#000;font-size:233%;font-weight:400;margin:0 0 0 4mm;padding:.25ex 0}#headerSubTitle{font-size:50%;font-style:italic;margin-left:1em}.headerTitle a{color:#000}.subHeader{display:none;color:#fff;background-color:#039;margin:0;padding:1ex 1ex 1ex 1.5mm}#side-bar a,.subHeader a{background-color:transparent}.subHeader a{color:#fff;font-weight:700;margin:0;padding:0 .75ex 0 .5ex}.subHeader .highlight,.superHeader .highlight{color:#fda05b;background-color:transparent}#side-bar{width:16em;float:left;clear:left;border-right:1px solid #ddd}#side-bar div{border-bottom:1px solid #ddd}.sideBarTitle{font-weight:700;margin:0 0 .5em 2mm;padding:1em 0 0}#side-bar ul{list-style-type:none;list-style-position:outside;margin:0;padding:0 0 .3em}li ul{padding-left:.6em!important}#side-bar li{margin:0;padding:.1ex 0}#side-bar a{color:#06c;margin:0;padding:.25em 1ex .25em 2mm;display:block;font-weight:700!important;font-size:102%;border-left:#fff solid .2em}.thisPage,.thisPage a{color:#000!important;background-color:#fff;padding-left:5mm}#main-copy,#main-copy a{background-color:transparent}#side-bar a:hover{color:#fff;background-color:#6487dc;border-left:#000 solid .2em}.sideBarText{line-height:1.5em;margin:0 0 1em;padding:0 1.5ex 0 2.5mm;display:block}#side-bar .sideBarText a{margin:0;padding:0;display:inline}#side-bar .sideBarText a:hover{color:#06c;background-color:transparent;text-decoration:none}#main-copy{max-width:70em;color:#000;text-align:justify;line-height:1.5em;margin:0 0 0 16em;padding:.5mm 5mm 5mm;border-left:1px solid #ddd}#bodyText{margin:0 0 0 15.5em;padding:2mm 5mm}#main-copy p{margin:1em 1ex!important;padding:0}#main-copy a{color:#06c}#main-copy a:hover{color:#6487dc}#main-copy .topOfPage,#main-copy h1,#main-copy h2{color:#06c;background-color:transparent;font-weight:700}#main-copy h1,#main-copy h2{font-size:145.5%;margin:2em 0 0;padding:.5ex 0 .5ex .6ex;border-bottom:2px solid #06c}#main-copy h2{font-size:115.5%;border-bottom:1px solid #06c}#main-copy .topOfPage{font-size:91%;margin:3ex 1ex 0 0;padding:0;float:right}dl{margin:1em 1ex 2em;padding:0}dt{font-weight:700;margin:0;padding:0}dd{margin:0 0 2em 2em;padding:0}#footer{color:#fff;background-color:#6487dc;padding:1em;clear:both}#footer .left{text-align:left;line-height:1.55em;float:left;clear:left}#footer .right{text-align:right;line-height:1.45em}#footer a{color:#fff;background-color:transparent}th{background-color:#abc;text-align:center}td{background-color:#def}hr{border-width:0 0 .1em;border-color:#000}.titleTip,acronym{border-bottom:1px solid #ddd;cursor:help;margin:0;padding:0 0 .4px}pre{margin-left:2em;font-size:1.2em}blockquote{border-left:1px solid #00f;font-style:italic}.smallCaps{font-size:110%;font-variant:small-caps}.doNotDisplay{display:none}.notify_errors,.notify_notes,.notify_success{padding:.8em;margin-bottom:1em;border:2px solid #ddd}.notify_errors{background:#FBE3E4;color:#8a1f11;border-color:#FBC2C4}.notify_notes{background:#FFF6BF;color:#514721;border-color:#FFD324}.notify_success{background:#E6EFC2;color:#264409;border-color:#C6D880}.notify_errors a{color:#8a1f11}.notify_notes a{color:#514721}.notify_success a{color:#264409}h1.dir-list-head,ul.dir-list{text-transform:capitalize;font-weight:700}.midHeader{background-color:#b4e1ff}`),
	}
)
