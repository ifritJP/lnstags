;;-*- coding:utf-8; mode:emacs-lisp-*-


(require 'lnstags-data)
(require 'lnstags-helm)

(defvar lnstags-command
  "lnstags"
  "lnstags command")

(defvar lnstags-mode-map (make-sparse-keymap)
  "lnstags-mode Keymap.")

(defvar lnstags-skip-error t)
(defun lnstags-err (mess)
  (error mess))



(defun lnstags-get-line ()
  (interactive)
  (if (eq (point) (point-max))
      (if (eq (current-column) 0)
	  (1+ (count-lines 1 (point)))
	(count-lines 1 (point)))
    (count-lines 1 (1+ (point)))))

(defun lnstags-get-column ()
  (interactive)
  (1+ (string-bytes (buffer-substring (point-at-bol) (point)))))


(defvar lnstags-token-pattern "[a-zA-Z0-9_]")
(defvar lnstags-no-token-pattern "[^a-zA-Z0-9_]")

(defun lnstags-get-current-token ()
  (interactive)
  (save-excursion
    (let (symstart symend)
      (setq symstart (re-search-backward lnstags-no-token-pattern nil t))
      (forward-char)
      (setq symend (re-search-forward lnstags-no-token-pattern nil t))
      (when (and symstart symend)
	(buffer-substring (1+ symstart) (1- symend))))))

(defun lnstags-match-token ( pos )
  (buffer-substring (match-beginning pos ) (match-end pos)))


(define-minor-mode lnstags-mode
  "LibClang Tag System mode.
"
  ;; 初期値
  nil
  ;; mode-line 文字列
  " lnstags"
  ;; マイナモード用キーマップの初期値
  lnstags-mode-map
  ;; body
  (run-hooks 'lnstags-mode-hook)
  )

(defun lnstags-get-parentDir (path)
  (cond ((file-name-directory path)
	 (directory-file-name (file-name-directory path)))
	(t
	 "./")))

(defun lnstags-get-path-from-projDir (buffer projdir)
  (with-current-buffer buffer
    (let (projInfo)
      (cond ((boundp 'lnstags-file-info)
	     (lnstags-file-info-get-path lnstags-file-info))
	    ((setq projInfo (lnstags-projs-get-from-projDir default-directory))
	     (lnstags-proj-get-relative-path projInfo buffer-file-name))
	    (t
	     (file-relative-name buffer-file-name projdir))))))

(defun lnstags-get-projDir (buffer)
  (with-current-buffer buffer
    (let ((dir default-directory)
	  projInfo
	  projdir)
      (cond ((boundp 'lnstags-file-info)
	     (setq projdir (lnstags-file-info-get-projDir lnstags-file-info)))
	    ((setq projInfo (lnstags-projs-get-from-projDir default-directory))
	     (setq projdir (lnstags-proj-get-projDir projInfo)))
	    (t
	     (while (and (not (string= dir "/"))
			 (not projdir))
	       (if (or (file-exists-p (expand-file-name "lune.js" dir))
		       (file-exists-p (expand-file-name "lnstags.sqlite3" dir)))
		   (setq projdir dir)
		 (setq dir (lnstags-get-parentDir dir))))
	     (when (not projdir)
	       (setq projdir default-directory))
	     projdir)))))

(defun lnstags-execute-op (src-buf lnstags-buf input async lnstags-opts)
  (let* (command dir exit-code)
    (setq dir (lnstags-get-projDir src-buf))
    (with-temp-buffer
      (setq default-directory dir)
      (setq command
	    (delq nil
		  (append lnstags-opts
			  (list "--log" "error")
			  )))
      (when input
	(insert input))
      (if async
	  (let (process)
	    (setq command
		  (append (list "lnstags-process" lnstags-buf lnstags-command)
			  command))
	    (setq process (apply 'start-process command))
	    (setq exit-code process)
	    (process-send-string process (buffer-string))
	    (run-with-timer 1 nil 'lnstags-process-scroll lnstags-buf) )
	(setq command
	      (append (list (point-min) (point-max) lnstags-command
			    nil lnstags-buf nil )
		      command))
	(setq exit-code (apply 'call-process-region command))))
    (with-current-buffer lnstags-buf
      (setq default-directory dir)
      (goto-char (point-min))
      )
    exit-code))

(defun lnstags-execute (src-buf lnstags-buf input &rest lnstags-opts)
  (lnstags-execute-op src-buf lnstags-buf input nil lnstags-opts))

(defun lnstags-process-scroll (lnstags-buf)
  (with-current-buffer lnstags-buf
    (goto-char (point-max))
    (when (get-buffer-window lnstags-buf)
      (recenter))
    ))

(defun lnstags-select-tags (buffer header-name projdir
				   select-func decide-func
				   &optional create-candidate-list)
  (with-current-buffer buffer
    ;;(message (buffer-string))
    (let ((lineNum (count-lines (point-min) (point-max))))
      (if (and (not lnstags-skip-error)
	       (string-match "^lnstags:" (buffer-string)))
	  (switch-to-buffer buffer)
	(cond
	 ((= lineNum 0)
	  (message "not found")
	  (lnstags-history-pop t)
	  (kill-buffer buffer)
	  nil
	  )
	 (t
	  (funcall select-func decide-func
		   (if create-candidate-list
		       create-candidate-list
		     'lnstags-tags-create-candidate-list)
		   header-name projdir)
	  t)
	 )))
    ))

(defun lnstags-pos-at ( mode &optional tag &rest lnstags-opt-list)
  (let ((save (current-buffer))
	(line (lnstags-get-line))
	(column (lnstags-get-column)) 
	(projdir (lnstags-get-projDir (current-buffer)))
	path buffer select-name lnstags-opt lnstags-opt2 opt-list input )
    (setq path (lnstags-get-path-from-projDir (current-buffer) projdir))
    (cond
     ((or (equal mode "def-at")
	  (equal mode "ref-at"))
      (setq lnstags-opt
	    (cond ((string= mode "def-at")
		   "def")
		  ((string= mode "ref-at")
		   "ref")))
      (setq input (buffer-string))
      (setq lnstags-opt2 '( "-i" ))
      (setq select-name
	    (format "(%s)%s<%s:%d:%d>"
		    (cond ((equal mode "def-at") "D")
			  ((equal mode "ref-at") "R"))
		    (lnstags-get-current-token)
		    (file-name-nondirectory buffer-file-name) line column)))
     )
    (setq buffer (generate-new-buffer lnstags-helm-buffer-name))
    ;;(setq default-directory projdir)
    (setq opt-list
	  (append (list save buffer input "inq-at" lnstags-opt
			path
			(number-to-string line)
			(number-to-string column)
			)
		  lnstags-opt2
		  lnstags-opt-list
		  ))
    (cond ((eq (apply 'lnstags-execute opt-list) 0)
	   (lnstags-select-tags buffer select-name projdir
				'lnstags-tags-select-mode 'lnstags-tags-select))
	  (t
	   (when (not lnstags-skip-error)
	     (switch-to-buffer buffer)))
	  )
    ))

(defun lnstags-ref-at ()
  (interactive)
  (lnstags-history-push)
  (lnstags-pos-at "ref-at"))

(defun lnstags-def-at ()
  (interactive)
  (lnstags-history-push)
  (lnstags-pos-at "def-at"))

(defvar lnstags-history nil)
(defun lnstags-history-push ()
  (add-to-list 'lnstags-history (list :buf (current-buffer) :mark (point-marker)) nil
	       (lambda (X X) nil)))

(defun lnstags-history-pop (&optional only-pop)
  (interactive)
  (let ((info (car lnstags-history))
	buf)
    (setq lnstags-history (cdr lnstags-history))
    (when (not only-pop)
      (while info
	(setq buf (plist-get info :buf))
	(cond ((buffer-live-p buf)
	       (setq marker (plist-get info :mark))
	       (switch-to-buffer buf)
	       (goto-char marker)
	       (setq info nil))
	      (t
	       (setq info (car lnstags-history))
	       (setq lnstags-history (cdr lnstags-history))))))
  ))


(provide 'lnstags)
