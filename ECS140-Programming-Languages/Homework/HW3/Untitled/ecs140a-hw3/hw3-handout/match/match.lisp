; You may define helper functions here

(defun match (pattern assertion)
  (cond ((and (null pattern) (null assertion)) T)
        ((xor (null pattern) (null assertion)) nil)
        ((equal pattern assertion) T)
        ((eql (car pattern) (car assertion)) (match (cdr pattern) (cdr assertion)))
        ((eql (car pattern) '?) (match (cdr pattern) (cdr assertion)))
        ((eql (car pattern) '!)
                (cond ((match (cdr pattern) (cdr assertion)) T)
                      ((null assertion) T)
                      (t (match pattern (cdr assertion)))
                )
        )
        (t nil)
  )
)
