package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/coc1961/gotranslate/internal/translate"
)

var validLanguage = `
Valid Language:

	Afrikaans	af
	Albanian	sq
	Amharic		am
	Arabic		ar
	Armenian	hy
	Azerbaijani	az
	Basque		eu
	Belarusian	be
	Bengali		bn
	Bosnian		bs
	Bulgarian	bg
	Catalan		ca
	Cebuano		ceb 
	Chinese (Simplified)	zh-CN or zh 
	Chinese (Traditional)	zh-TW 
	Corsican	co
	Croatian	hr
	Czech		cs
	Danish		da
	Dutch		nl
	English		en
	Esperanto	eo
	Estonian	et
	Finnish		fi
	French		fr
	Frisian		fy
	Galician	gl
	Georgian	ka
	German		de
	Greek		el
	Gujarati	gu
	Haitian Creole	ht
	Hausa		ha
	Hawaiian	haw 
	Hebrew		he or iw
	Hindi		hi
	Hmong		hmn (ISO-639-2)
	Hungarian	hu
	Icelandic	is
	Igbo		ig
	Indonesian	id
	Irish		ga
	Italian		it
	Japanese	ja
	Javanese	jv
	Kannada		kn
	Kazakh		kk
	Khmer		km
	Korean		ko
	Kurdish		ku
	Kyrgyz		ky
	Lao			lo
	Latin		la
	Latvian		lv
	Lithuanian	lt
	Luxembourgish	lb
	Macedonian	mk
	Malagasy	mg
	Malay		ms
	Malayalam	ml
	Maltese		mt
	Maori		mi
	Marathi		mr
	Mongolian	mn
	Myanmar (Burmese)	my
	Nepali		ne
	Norwegian	no
	Nyanja (Chichewa)	ny
	Pashto		ps
	Persian		fa
	Polish		pl
	Portuguese (Portugal, Brazil)	pt
	Punjabi		pa
	Romanian	ro
	Russian		ru
	Samoan		sm
	Scots Gaelic	gd
	Serbian		sr
	Sesotho		st
	Shona		sn
	Sindhi		sd
	Sinhala (Sinhalese)	si
	Slovak		sk
	Slovenian	sl
	Somali		so
	Spanish		es
	Sundanese	su
	Swahili		sw
	Swedish		sv
	Tagalog (Filipino)	tl
	Tajik		tg
	Tamil		ta
	Telugu		te
	Thai		th
	Turkish		tr
	Ukrainian	uk
	Urdu		ur
	Uzbek		uz
	Vietnamese	vi
	Welsh		cy
	Xhosa		xh
	Yiddish		yi
	Yoruba		yo
	Zulu		zu`

func main() {
	language := flag.String("l", "en", "language to translate, default (en)")
	sourceLanguage := flag.String("s", "es", "actual source code comments language, default (es)")
	inputSource := flag.String("i", "", "input source dir")
	outputSource := flag.String("o", "", "output source dir")
	flag.Parse()

	if *inputSource == "" || *outputSource == "" {
		fmt.Println(validLanguage)
		fmt.Println("\n\nTranslate the source code comments of go project to the specified language\n\n ")
		fmt.Println("input and output source dir are mandatory")
		flag.PrintDefaults()
		os.Exit(1)
	}
	translate, err := translate.New(*language, *sourceLanguage, *inputSource, *outputSource)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = translate.Translate()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
