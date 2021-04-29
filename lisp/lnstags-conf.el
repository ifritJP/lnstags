(add-hook 'lnstags-mode-hook
      '(lambda ()
         (local-set-key (kbd "M-t") 'lnstags-def-at)
         (local-set-key (kbd "M-r") 'lnstags-ref-at)
	 (local-set-key (kbd "M-m") 'lnstags-tags-resume)
         (local-set-key "\C-t" 'lnstags-history-pop)))

(require 'lnstags)
(require 'lnstags-helm)

(add-hook 'lns-mode-hook
          '(lambda()
	     (lnstags-mode 1)))

(provide 'lnstags-conf)
