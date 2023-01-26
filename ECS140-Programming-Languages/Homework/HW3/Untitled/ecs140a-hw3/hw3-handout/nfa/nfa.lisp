;; You may define helper functions here

(defun occurs (elem li)
    (cond  ((null li) nil )
           ((equal elem (car li)) t  )
           ( (occurs elem (cdr li) ) )
    )
)


(defun reachable (transition start final input)

    
    (cond
        ((eql input nil)
            (cond
                ((eql start final) T)
                
            )
        )
        ((occurs T (mapcar #'(lambda(y) (reachable transition y final (cdr input))) (funcall transition start (car input)))) )
    )
)