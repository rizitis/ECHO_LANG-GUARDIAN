;;; echo-mode.el - Major mode for ECHO-LANG
;; Maintainer: You
;; Version: 1.0

(defgroup echo nil
  "ECHO-LANG mode"
  :group 'languages)

;; Define syntax highlighting
(defconst echo-font-lock-keywords
  (list
   ;; EXECUTE in red (must come first!)
   '("\\<EXECUTE\\>" . font-lock-warning-face)

   ;; ON RESONANCE, ON HEARTBEAT, ON SYSTEM, ON RESPONSE
   '("\\<ON\\>\\s-+\\(RESONANCE\\|HEARTBEAT\\|SYSTEM\\|SIGNAL\\|RESPONSE\\|INIT\\)" 1 font-lock-keyword-face t)

   ;; Other keywords
   '("\\<\\(ON\\|RESONANCE\\|SIGNAL\\|INIT\\|BLOCK\\|PRINT\\|MODIFY\\|IF\\|THEN\\|ELSE\\|ELSEIF\\|ENDIF\\|SET\\|TO\\|EXISTS\\|CONTAINS\\|STARTS_WITH\\|MIN\\|MAX\\|APPLY\\|WRITE_FILE\\|DELETE\\|PROMPT\\)\\>" . font-lock-keyword-face)

   ;; Variables: $awareness, $focus
   '("\\$\\w+" . font-lock-variable-name-face)

   ;; Strings: "..."
   '("\"[^\"]*\"" . font-lock-string-face)

   ;; Single quotes: '...'
   '("'[^']*'" . font-lock-string-face)

   ;; Numbers
   '("\\b[0-9]+\\b" . font-lock-number-face)
   '("\\b[0-9]*\\.[0-9]+\\b" . font-lock-number-face)

   ;; Operators
   '("=" . font-lock-operator-face)

   ;; Comments
   '("#.*" . font-lock-comment-face)
   )
  "Highlighting expressions for ECHO-LANG mode")

;; Define the mode
(define-derived-mode echo-mode prog-mode "ECHO"
  "Major mode for editing ECHO-LANG files."
  :syntax-table nil
  (setq font-lock-defaults '(echo-font-lock-keywords))
  (setq-local comment-start "# ")
  (setq-local comment-end "")
  (setq-local indent-tabs-mode nil)
  (setq-local tab-width 4))

;; Auto-detect .echo files
(add-to-list 'auto-mode-alist '("\\.echo\\'" . echo-mode))

(provide 'echo-mode)
