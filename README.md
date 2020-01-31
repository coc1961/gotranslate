# gotranslate

> Translate the source code comments of go project to the specified language

## install

```sh
go get -u github.com/coc1961/gotranslate/...
```


## Help

```sh
Valid Language:

        Afrikaans       af
        Albanian        sq
        Amharic         am
        Arabic          ar
        Armenian        hy
        Azerbaijani     az
        Basque          eu
        Belarusian      be
        Bengali         bn
        Bosnian         bs
        Bulgarian       bg
        Catalan         ca
        Cebuano         ceb 
        Chinese (Simplified)    zh-CN or zh 
        Chinese (Traditional)   zh-TW 
        Corsican        co
        Croatian        hr
        Czech           cs
        Danish          da
        Dutch           nl
        English         en
        Esperanto       eo
        Estonian        et
        Finnish         fi
        French          fr
        Frisian         fy
        Galician        gl
        Georgian        ka
        German          de
        Greek           el
        Gujarati        gu
        Haitian Creole  ht
        Hausa           ha
        Hawaiian        haw 
        Hebrew          he or iw
        Hindi           hi
        Hmong           hmn (ISO-639-2)
        Hungarian       hu
        Icelandic       is
        Igbo            ig
        Indonesian      id
        Irish           ga
        Italian         it
        Japanese        ja
        Javanese        jv
        Kannada         kn
        Kazakh          kk
        Khmer           km
        Korean          ko
        Kurdish         ku
        Kyrgyz          ky
        Lao                     lo
        Latin           la
        Latvian         lv
        Lithuanian      lt
        Luxembourgish   lb
        Macedonian      mk
        Malagasy        mg
        Malay           ms
        Malayalam       ml
        Maltese         mt
        Maori           mi
        Marathi         mr
        Mongolian       mn
        Myanmar (Burmese)       my
        Nepali          ne
        Norwegian       no
        Nyanja (Chichewa)       ny
        Pashto          ps
        Persian         fa
        Polish          pl
        Portuguese (Portugal, Brazil)   pt
        Punjabi         pa
        Romanian        ro
        Russian         ru
        Samoan          sm
        Scots Gaelic    gd
        Serbian         sr
        Sesotho         st
        Shona           sn
        Sindhi          sd
        Sinhala (Sinhalese)     si
        Slovak          sk
        Slovenian       sl
        Somali          so
        Spanish         es
        Sundanese       su
        Swahili         sw
        Swedish         sv
        Tagalog (Filipino)      tl
        Tajik           tg
        Tamil           ta
        Telugu          te
        Thai            th
        Turkish         tr
        Ukrainian       uk
        Urdu            ur
        Uzbek           uz
        Vietnamese      vi
        Welsh           cy
        Xhosa           xh
        Yiddish         yi
        Yoruba          yo
        Zulu            zu


Translate the source code comments of go project to the specified language

 
input and output source dir are mandatory
  -i string
        input source dir
  -o string
        output source dir
  -s string
        actual source code comments language, default (es) (default "es")
  -l string
        language to translate, default (en) (default "en")
``` 

## Example

> This example translates the source code with comments in English (-s en) to Japanese (-l ja).
> source code of the project in /home/user/go/src/github.com/testproject 
> and save the result in /tmp/testproyect

```sh

gotranslate -i /home/user/go/src/github.com/testproject -o /tmp/testproyect -s en -l ja

```