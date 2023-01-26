;;;;;;;;;;;;;;;;;;;;;;;;
;;; pre-testing prep ;;;
;;;;;;;;;;;;;;;;;;;;;;;;

(load "../lisp-unit.lisp")

(use-package :lisp-unit)

(load "nfa.lisp")

(remove-tests :all)

(setq *print-failures* t)
(setq *print-errors* t)

(defun dagTransitions (state input)
  (cond
    ((eq state 40) '())
    ((and (eq state 41) (eq input 'F)) '(42))

    ((and (eq state 42) (eq input 'D)) '(43 44 48 49 50))

    ((and (eq state 43) (eq input 'T)) '(45))
    ((and (eq state 43) (eq input 'Q)) '(46))

    ((and (eq state 44) (eq input 'T)) '(45))
    ((and (eq state 44) (eq input 'M)) '(47))

    (t (list nil))))

(defun dfaTransitions (state input)
  (cond
    ((and (eq state 47) (eq input 'P)) '(56))
    ((and (eq state 47) (eq input 'U)) '(65))

    ((and (eq state 56) (eq input 'P)) '(59))
    ((and (eq state 56) (eq input 'U)) '(62))

    ((and (eq state 65) (eq input 'P)) '(68))
    ((and (eq state 65) (eq input 'U)) '(71))

    ((and (eq state 59) (eq input 'P)) '(60))
    ((and (eq state 59) (eq input 'U)) '(61))

    ((and (eq state 62) (eq input 'P)) '(63))
    ((and (eq state 62) (eq input 'U)) '(64))

    ((and (eq state 68) (eq input 'P)) '(69))
    ((and (eq state 68) (eq input 'U)) '(70))

    ((and (eq state 71) (eq input 'P)) '(72))
    ((and (eq state 71) (eq input 'U)) '(73))

    ((and (eq state 60) (eq input 'F)) '(48))

    ((and (eq state 61) (eq input 'F)) '(48))

    ((and (eq state 63) (eq input 'F)) '(48))

    ((and (eq state 64) (eq input 'F)) '(48))

    ((and (eq state 69) (eq input 'F)) '(48))

    ((and (eq state 70) (eq input 'F)) '(48))

    ((and (eq state 72) (eq input 'F)) '(48))

    ((and (eq state 73) (eq input 'F)) '(48))

    ((and (eq state 48) (eq input 'M)) '(47))

    (t (list nil))))

(defun nfaTransitions (state input)
  (cond
    ((and (eq state 33) (eq input 'Y)) '(100 300 33 200 500 34))
    ((and (eq state 33) (eq input 'Z)) '(100 400 200 500 33))

    ((and (eq state 34) (eq input 'Y)) '(100 200 300 400 35))
    ((and (eq state 34) (eq input 'Z)) '(100 200 300 400 35))

    ((and (eq state 35) (eq input 'Y)) '(100 400 500 36 200 300))
    ((and (eq state 35) (eq input 'Z)) '(100 400 500 36 300 200))

    ((and (eq state 36) (eq input 'Y)) '(100 500 300 200 37 400))
    ((and (eq state 36) (eq input 'Z)) '(400 300 200 100 37 500))

    (t (list nil))))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;;; reachable test definitions ;;;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(define-test test-reachable-dag-01
             (assert-equal T (reachable 'dagTransitions 40 40 '())))
(define-test test-reachable-dag-02
             (assert-equal NIL (reachable 'dagTransitions 40 41 '())))
(define-test test-reachable-dag-03
             (assert-equal T (reachable 'dagTransitions 41 42 '(F))))
(define-test test-reachable-dag-04
             (assert-equal NIL (reachable 'dagTransitions 41 42 '(D))))
(define-test test-reachable-dag-05
             (assert-equal T (reachable 'dagTransitions 42 50 '(D))))
(define-test test-reachable-dag-06
             (assert-equal NIL (reachable 'dagTransitions 42 45 '(D T W))))

(define-test test-reachable-dfa-01
             (assert-equal NIL (reachable 'dfaTransitions 47 61 '(P P P U))))
(define-test test-reachable-dfa-02
             (assert-equal T (reachable 'dfaTransitions 47 48 '(P U U F M U P P F))))

(define-test test-reachable-nfa-01
             (assert-equal NIL (reachable 'nfaTransitions 33 37 '(Y Z Z Y Z Z))))
(define-test test-reachable-nfa-02
             (assert-equal T (reachable 'nfaTransitions 33 37 '(Z Y Y Z Y Y))))

;;;;;;;;;;;;;;;;;
;;; run tests ;;;
;;;;;;;;;;;;;;;;;

(run-tests :all)
