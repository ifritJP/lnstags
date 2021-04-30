(defvar lnstags-path-length 60)


(require 'cl)
(require 'helm)
(require 'lnstags-data)

(defvar lnstags-helm-buffer-name "*lnstags SEL*")

;; [C-return]
(defvar lnstags-heml-map
  (let (map)
    (setq map (copy-keymap helm-map))
    (define-key map (kbd "C-z") 'lnstags-helm-execute-persistent-action)
    (delq nil map))
  "Keymap")

(defun lnstags-helm-execute-persistent-action (&optional attr onewindow)
  (interactive)
  (helm-execute-persistent-action attr onewindow))



(defface lnstags-candidate-face
  '((t
     :foreground "gold"))
  "candidate face")
(defvar lnstags-candidate-face 'lnstags-candidate-face)

(defface lnstags-candidate-path-face
  '((t
     :foreground "green"))
  "candidate path face")
(defvar lnstags-candidate-path-face 'lnstags-candidate-path-face)


(defun lnstags-tags-select (item)
  (setq line (plist-get item :line))
  (setq file (plist-get item :path))
  (setq relative (plist-get item :relative))
  (setq projDir (plist-get item :projDir))
  (let ((prev-buffer (current-buffer)))
    (find-file file)
    (set (make-local-variable 'lnstags-file-info)
	 `(:path ,relative :projDir ,projDir))
    (lnstags-projs-add projDir (lnstags-get-parentDir relative))
    (goto-line line)
    (recenter)
    ))

(defun lnstags-conv-disp-path (path omit &optional base-dir)
  (when (string-match "^!" path)
    (setq path (substring path 1)))
  (let* ((relative-path (file-relative-name path
					    (or base-dir default-directory)))
	 (abs-path (expand-file-name path))
	 (disp-path relative-path))
    (when (< (length abs-path) (length relative-path))
      (setq disp-path abs-path))
    (if (or (not omit)
	    (not lnstags-path-length)
	    (> lnstags-path-length (length disp-path)))
	disp-path
      (concat "..."
	      (substring disp-path
			 (+ (* -1 lnstags-path-length) 3))))))

(defun lnstags-tags-create-candidate-list (projDir &optional output-symbol-flag)
  (let (candidate-list)
    (while (not (eobp))
      (beginning-of-line)
      (cond ((or (looking-at "^lnstags:	1")
		 (looking-at ".*<not found>"))
	     (message (lnstags-match-token 0)))
	    (t
	     (when (not (looking-at "^\\([^ \t]+\\)[ \t]+\\([0-9]+\\)[ \t]\\([^ \t]+\\)[ \t]\\(.*\\)$"))
	       (lnstags-err "illegal format"))
	     (let* ((symbol (lnstags-match-token 1))
		    (line (string-to-number (lnstags-match-token 2)))
		    (path (lnstags-match-token 3))
		    (txt (lnstags-match-token 4)))
	       (setq candidate-list
		     (cons (cons (concat
				  (when output-symbol-flag
				    (concat (propertize symbol 'face lnstags-candidate-face)
					    ":\t"))
				  (format "%s:%d:%s"
					  (propertize
					   (lnstags-conv-disp-path path t projDir)
					   'face lnstags-candidate-path-face)
					  line txt))
				 (list :line line
 				       :path (expand-file-name path projDir)
				       :relative path
				       :projDir projDir
				       ))
			   candidate-list)))))
      (next-line))
    candidate-list
    ))

(defvar lnstags-helm-history nil)
(defvar lnstags-helm-history-cur nil)

(defun lnstags-helm-history-add-tail (item)
  (let ((time (format-time-string "%m/%d-%T" (current-time))))
    (setq lnstags-helm-history-cur time)
    (add-to-list 'lnstags-helm-history
		 (list (format "%s %s" time (cdr (assoc 'name item)))
		       :helm item
		       :time time)
		 t)))
(defun lnstags-helm-history-del-top ()
  (let ((item (car lnstags-helm-history)))
    (setq lnstags-helm-history (cdr lnstags-helm-history))
    item))
(defun lnstags-helm-history-clear ()
  (interactive)
  (dolist (buf (buffer-list))
    (when (string-match (regexp-quote lnstags-helm-buffer-name)
			(buffer-name buf))
      (kill-buffer buf)))
  (setq lnstags-helm-history nil)
  )
(defun lnstags-helm-history-select (item)
  (setq lnstags-helm-history-cur (plist-get item :time))
  (helm :sources (plist-get item :helm)
	:keymap lnstags-heml-map
	:buffer lnstags-helm-buffer-name)
  )
(defun lnstags-helm-history-show ()
  (interactive)
  (helm :sources `((name . "lnstags-history")
		   (candidates . ,(reverse lnstags-helm-history))
		   (action . lnstags-helm-history-select))
	:buffer lnstags-helm-buffer-name
	:preselect lnstags-helm-history-cur
	)
  )
  

(defun lnstags-tags-select-mode (select-func create-candidate-list-func
					     header-name projDir)
  (let (candidate-list lnstags-params)
    (setq candidate-list (funcall create-candidate-list-func projDir))
    (if (eq (length candidate-list) 1)
	(funcall select-func (cdr (car candidate-list)))
      (setq candidate-list
	    (sort candidate-list
		  (lambda (X Y)
		    (let* ((infoX (cdr X))
			   (infoY (cdr Y))
			   (pathX (plist-get infoX :path))
			   (pathY (plist-get infoY :path)))
		      (if (string< pathX pathY)
			  t
			(if (string= pathX pathY)
			    (if (< (plist-get infoX :line) (plist-get infoY :line))
				t
			      nil)))))))
      (setq lnstags-params
	    `((name . ,(concat lnstags-helm-buffer-name header-name))
	      (candidates . ,candidate-list)
	      (action . ,select-func)))
      (lnstags-helm-history-add-tail lnstags-params)
      (let ((helm-candidate-number-limit 9999))
	(helm :sources lnstags-params
	      :keymap lnstags-heml-map
	      :buffer lnstags-helm-buffer-name)
	))))

(defun lnstags-tags-resume ()
  (interactive)
  (lnstags-helm-history-show))


(provide 'lnstags-helm)
