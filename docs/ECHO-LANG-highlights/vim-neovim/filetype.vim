" Detect .echo files
augroup filetypedetect
  au! BufNewFile,BufRead *.echo setfiletype echo
augroup END
