;; You may define helper functions here

(defun mylength (l)
   (if (null l) 0 (+ 1 (mylength (cdr l)))
    
    )
)

(defun listIter (l)
    (cond
        ((and eql(mylength(l) 0) nil)) ;; If the length of the list is 0, then it does not work
    )
)

(defun reachable (transition start final input)
    ;; (let() ))
    (let (x '(funcall transition start (car input))) (princ "Start: ") (write start) (princ " Input: ") (write input)(terpri))
    ;; (let (y '(funcall transition start (car input))))
    ;; (princ(start))
    (cond 
    ;; ((and (eql (funcall transition start (car input)) nil)) nil)
    ;; ((and ()))
    ((and (eql input nil) (eql 'start 'final))T) ;; End of symbols
    ((and (eql input nil))) 
    (t(reachable transition (funcall transition start (car input)) final (cdr input)))
    )
   
    

)


    ;; (cond
    ;;     ((eq (mylength input) 0) 
    ;;     (cond
    ;;         ((eq start final) nil)
    ;;         (t T)
    ;;     ))
    ;;     (t (reachable transition (funcall transition start (car input)) final (cdr input)))
        
    ;; )
    ;; (reachable transition (funcall transition start (car input)) final (cdr input))
    ;; Okay we have it able to do all of these but in 
    ;; (reachable transition (funcall transition start (car input)) final (cdr input))
    ;; (reachable 'fooTransitions 0 0 NIL)
    ;; (mylength input)
    ;; (funcall transition '0 'B) ;; This works! Return 2
    ;; (funcall transition start (car input)) 
 
    ;; (cond 
    ;;     ((and (eq start 0) (eq input)))
          
    ;; )

    ;; (cond ((> x 5) (- x 1)) ;; If x > 5, then subtract 1 from x return
    ;;      ((eql x 5) x) ;; If equal to 5, return x
    ;;      ((< x 5) (+ x 1))) ;; If x less than 5, then add to 1 return
    ;; So what we need is 
    ;; (car input) ;; Gets us the first of input

    ;; (( car input ) T) ;; example of working condition
    
    
    ;; nil ;; placeholder
    ;; (apply 'transition '(input))
    ;; TODO: Incomplete function
    ;; (transition(car(input) cdr(input)))
    ;; transition('0 'B) ;; transitions('start car(input))