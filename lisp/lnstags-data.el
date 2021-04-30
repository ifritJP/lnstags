;; lnstags-projs
(defvar lnstags-projs nil)
(defun lnstags-projs-add (projDir relative-path)
  (add-to-list 'lnstags-projs (list :projDir projDir :relative relative-path))
  )
(defun lnstags-projs-get-from-projDir (projDir)
  (car (delq nil (mapcar (lambda (X)
			   (when (equal (plist-get X :projDir) projDir)
			     X))
			 lnstags-projs))))

;; lnstags-proj
(defun lnstags-proj-get-projDir (proj)
  (plist-get proj :projDir))
(defun lnstags-proj-get-relative-path (proj path)
  (let ((relative (plist-get proj :relative)))
    (concat relative
	    (if (string-match "/$" relative)
		""
	      "/")
	    (file-name-nondirectory path))))

;; lnstags-file-info
(defun lnstags-file-info-get-projDir (info)
  (plist-get info :projDir))
(defun lnstags-file-info-get-path (info)
  (plist-get info :path))



(provide 'lnstags-data)
