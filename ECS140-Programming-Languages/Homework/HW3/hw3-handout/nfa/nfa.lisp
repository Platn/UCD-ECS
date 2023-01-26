;; You may define helper functions here

(defun occurs (elem li)
    (cond  ((null li) nil )
           ((equal elem (car li)) t  )
           ( (occurs elem (cdr li) ) )
    )
)


(defun reachable (transition start final input)
    ;; (let() ))
    ;; (princ(car input))
    ;; (let (x '(funcall transition start (car input))) (princ "Input: ") (write input)(terpri) (princ "Start: ") (write start)(princ " ") (princ "Final: ")(write final)(terpri))
        ;; (write (mylength(funcall transition start (car input)))))
    
    (cond
        ((eql input nil)
            (cond
                ((eql start final) T)
                
            )
        )
        ((occurs T (mapcar #'(lambda(y) (reachable transition y final (cdr input))) (funcall transition start (car input)))) )
    )
)
    
    
    
    
    
    ;; (cond 
    ;; ;; ((and (eql (funcall transition start (car input)) nil)) nil)
    ;; ;; If transition is a list, we need to check, but nil is a list, 
    ;; ;; Okay we need to do a mapping. It might not even need me to do a detection, we have mapcar, transition always returns a list
    ;; ;; ((> (mylength (funcall transition start (car input))) 1 )( reachable transition (car(funcall transition start (car input))) final (cdr input)))
    ;;     ((and (eql input nil) (eql start final))T) ;; If input is nil and start == final, then return true
    ;;     ((eql start nil) nil)
        
    ;;     ((eql nil input) (princ "Input nil") nil) ;; If input is nil, what is this returning? It seems that the standalone defaults to nil
    ;;     ((eql (funcall transition start (car input)) nil)) ;; If 
    
    
    ;; ;; ((eql x nil) nil)
    ;; ;; (t(reachable transition (funcall transition start (car input)) final (cdr input)))
    ;;     (t (princ "Split") (terpri)
    ;;         ;; (funcall 'orfunc(mapcar #'(lambda(y) (reachable transition y final (cdr input))) (funcall transition start (car input)))
    ;;             (terpri)(write(mapcon #'(lambda(y) (reachable transition y final (cdr input))) (funcall transition start (car input))
            
    ;;         ))
    
    ;;     )
    ;; )



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