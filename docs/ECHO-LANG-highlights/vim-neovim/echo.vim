" Vim syntax file
" Language: ECHO-LANG
" Maintainer: You
" Last Change: 2025

" Clear any existing syntax
if exists("b:current_syntax")
  unlet b:current_syntax
endif

" Highlight groups
hi def link echoKeyword Keyword
hi def link echoDanger Function          " EXECUTE in red
hi def link echoString String
hi def link echoVariable Identifier
hi def link echoComment Comment
hi def link echoNumber Number
hi def link echoOperator Operator
hi def link echoEscape Special

" Colors (for 256-color or truecolor terminals)
hi echoKeyword      guifg=#00FFFF ctermfg=87
hi echoDanger       guifg=#FF0000 ctermfg=196 bold
hi echoString       guifg=#00FF00 ctermfg=46
hi echoVariable     guifg=#FFA500 ctermfg=214
hi echoComment      guifg=#808080 ctermfg=244
hi echoNumber       guifg=#FF6347 ctermfg=203
hi echoOperator     guifg=#FFFF00 ctermfg=226
hi echoEscape       guifg=#FF69B4 ctermfg=219

" Keywords
syn keyword echoKeyword ON RESONANCE SIGNAL INIT BLOCK PRINT MODIFY IF THEN ELSE ELSEIF ENDIF SET TO EXISTS CONTAINS STARTS_WITH MIN MAX APPLY DELETE PROMPT

" EXECUTE in red (must come before general keywords)
syn keyword echoDanger EXECUTE

" Compound ON events
syn match echoKeyword /\bON\s\+HEARTBEAT\b/
syn match echoKeyword /\bON\s\+SYSTEM\b/
syn match echoKeyword /\bON\s\+RESPONSE\b/

" Variables: $awareness, $focus
syn match echoVariable /\$\w\+/

" Strings: "..." and '...'
syn region echoString start=+"+ end=+"+ skip=+\\\\\|\\"+ contains=echoEscape,echoVariable
syn region echoString start=+'+ end=+'+ skip=+\\\\\|\\'+ contains=echoEscape,echoVariable

" Escape sequences in strings
syn match echoEscape /\\[nrtb\\"]/

" Numbers
syn match echoNumber /\d\+\.\d\+/
syn match echoNumber /\d\+/

" Operators
syn match echoOperator /=/

" Comments
syn match echoComment /#.*/

" Delimiters (braces, semicolons)
syn match echoOperator /[{;}]/

" Default link if no colorscheme overrides
if !exists("did_echo_syn_inits")
  let did_echo_syn_inits = 1
  hi def link echoKeyword Keyword
  hi def link echoString String
  hi def link echoVariable Identifier
  hi def link echoComment Comment
  hi def link echoNumber Number
  hi def link echoOperator Operator
  hi def link echoEscape Special
endif
