; Some lisp code that we were using lectures.

(defun count-leaf (tree) 
    (if (atom tree) 1 
        (+ (count-leaf (car tree))
           (or (if (cdr tree)
                   (count-leaf (cdr tree)))
               1))))   

; (count-leaf '((1 2 (3 4)) (5) 6))


(defun foo (a)
    (apply '+ (mapcar (lambda (n) 1) a))
)


(defun prefix (xs li)
    (cond ((or (null li) (null xs))  nil)
          (t (cons (car li)
                   (prefix (cdr xs) (cdr li)    )  )    )
    )
)


(defun count (item li)
    (apply '+ (mapcar (lambda (n) (if (equal item n) 1 0)) li)))

(defun mylength (l)
   (if (null l) 0 
       (+ 1 (mylength (cdr l)))))



(defun app (a b)
    (cond ((null a) b)
          ( t  (cons (car a) 
                    (app (cdr a) b))   )
    )
) 

(defun occurs (elem li)
    (cond  ((null li) nil )
           ((equal elem (car li)) t  )
           ( (occurs elem (cdr li) ) )
    )
)

(defun orfunc (l)
    (cond ((null l) nil)
          ((or (car l) (orfunc (cdr l))))
          
    )
)

(defun newoccurs (elem li)
        (funcall 'orfunc (mapcar (lambda (n) (equal elem n) ) li)))




(defun count-atoms (exp)
    (cond ((null exp) 0)
          ((atom exp) 1)
          (t (+ (count-atoms (first exp))  
                (count-atoms (rest exp))  
             )
          )
    )
)

(defun count-anywhere (item tree)
    (cond ( (eql item tree)  1)
          ( (atom tree) 0)
          (t (+ (count-anywhere item (car tree)) 
                (count-anywhere item (cdr tree)))
          )
    )

)

(defun fname (li)
    (cond  ((base li) (   )   )
           (t         (fname li )    )
    
    )

)

(cond  ((predicate) (body)  )
       ((predicate) (body)  )

)

;; (count-anywhere 'a '(a (b ) b))
;; (count-anywhere 'b '(a (b ) b))


(defun factorial (n)
;; n!
   (cond ((zerop n)  1)
          (t (* n  (factorial (- n 1) ))
          )
   )
)

(defun fib (n)
   (cond 
        ((or (zerop n) (eql n 1))  1         )
        ( t (+ (fib (- n 1)) (fib (- n 2))))
   )
)

(defun power (x n)
   ;; x^n
   (cond ( (zerop n) 1)
         ( (evenp n)  
          (let ((temp (power x (/ n 2))))
             (* temp temp)
          )
           )
         (t (* x (power x (- n 1))))
   )
)

;; (dot-product '(10 20) '(3 4)) =>  10*3 + 20*4 = 110
(defun dot-product (a b)
    ;; (cond 
    ;; ( (and (null  a ) (null b) )  0  )
    ;; ;; ( (and (null (cdr a )) (null (cdr b)) )
    ;; ;;          (* (car a) (car b))
    ;; ;;       )
    ;; ( t (+ (* (car a) (car b) )  (dot-product (cdr a) (cdr b) )  ) )
    ;; )
    (apply #'+ (mapcar #'* a b))

)






